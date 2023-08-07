/*
 Copyright 2023, NVIDIA CORPORATION & AFFILIATES
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

 SPDX-License-Identifier: Apache-2.0
 SPDX-FileCopyrightText: Copyright 2023, NVIDIA CORPORATION & AFFILIATES
*/

package driver

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-logr/logr"

	"github.com/Mellanox/nic-feature-discovery/pkg/feature"
	"github.com/Mellanox/nic-feature-discovery/pkg/feature/internal/common"
)

const (
	mlxverLabelName   = "mofed.version"
	mlx5ModulePath    = "/sys/module/mlx5_core"
	mlx5ModuleVersion = "/sys/module/mlx5_core/version"
)

var (
	errFeatureNotExist = errors.New("feature does not exist")
)

func init() {
	feature.AddToSources(NewDriverSource())
}

type driverSource struct{}

// NewDriverSource creates a new driver source
func NewDriverSource() feature.Source {
	return &driverSource{}
}

// Name of Source
func (ds *driverSource) Name() string {
	return "driver"
}

// Features returns list of Features this Source supports
func (ds *driverSource) Discover(ctx context.Context) ([]feature.Feature, error) {
	log := logr.FromContextOrDiscard(ctx)
	log = log.WithName("driver-source")

	log.Info("discovering features")

	var fs []feature.Feature
	if f, err := ds.discoverMlxVerFeature(log); err == nil {
		fs = append(fs, f)
	} else {
		if errors.Is(err, errFeatureNotExist) {
			return fs, nil
		}

		return nil, fmt.Errorf("failed to discover mlx version. %w", err)
	}

	return fs, nil
}

func (ds *driverSource) discoverMlxVerFeature(log logr.Logger) (feature.Feature, error) {
	log.V(5).Info("discovering mlx version")
	_, err := os.Lstat(mlx5ModulePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.V(3).Info("mlx5 driver is not loaded", "path", mlx5ModulePath)

			return nil, errFeatureNotExist
		}

		return nil, err
	}

	_, err = os.Lstat(mlx5ModuleVersion)
	if err != nil {
		if os.IsNotExist(err) {
			log.V(3).Info("mlx5 driver has no version file, its most likely inbox", "path", mlx5ModuleVersion)

			return nil, errFeatureNotExist
		}

		return nil, err
	}

	// get version
	data, err := os.ReadFile(mlx5ModuleVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to read driver version. %w", err)
	}

	ver := strings.TrimSpace(string(data))
	if ver == "" {
		return nil, fmt.Errorf("unexpected driver version(%q)", ver)
	}

	return common.NewGenericFeature("mofed-version").AddLabel(common.DefaultPrefixedKey(mlxverLabelName), ver), nil
}
