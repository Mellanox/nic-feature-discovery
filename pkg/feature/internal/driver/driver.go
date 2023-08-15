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

	deps "github.com/Mellanox/nic-feature-discovery/pkg/dependencies"
	"github.com/Mellanox/nic-feature-discovery/pkg/feature"
	"github.com/Mellanox/nic-feature-discovery/pkg/feature/internal/common"
)

const (
	// SourceNameDriver is the name of driver feature source
	SourceNameDriver = "driver"
	// FeatureNameMofedVersion is the name of mofed version feature
	FeatureNameMofedVersion = "mofed-version"
	// LableNameMOFEDVersion is the unprefixed label name to be used for mofed version feature label
	LableNameMOFEDVersion = "mofed.version"
	// Mlx5ModulePath is the path for mlx5_core module if loaded in kernel
	Mlx5ModulePath = "/sys/module/mlx5_core"
	// Mlx5ModuleVersion is the path to mlx5_core version file if exists.
	Mlx5ModuleVersion = "/sys/module/mlx5_core/version"
)

var (
	// errFeatureNotExist is an internal error to indicate a feature does not exist
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
	return SourceNameDriver
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
	_, err := deps.OS.Lstat(Mlx5ModulePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.V(3).Info("mlx5 driver is not loaded", "path", Mlx5ModulePath)

			return nil, errFeatureNotExist
		}

		return nil, err
	}

	_, err = deps.OS.Lstat(Mlx5ModuleVersion)
	if err != nil {
		if os.IsNotExist(err) {
			log.V(3).Info("mlx5 driver has no version file, its most likely inbox", "path", Mlx5ModulePath)

			return nil, errFeatureNotExist
		}

		return nil, err
	}

	// get version
	data, err := deps.OS.ReadFile(Mlx5ModuleVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to read driver version. %w", err)
	}

	ver := strings.TrimSpace(string(data))
	if ver == "" {
		return nil, fmt.Errorf("unexpected driver version(%q)", ver)
	}

	return common.NewGenericFeature(FeatureNameMofedVersion).
		AddLabel(common.DefaultPrefixedKey(LableNameMOFEDVersion), ver), nil
}
