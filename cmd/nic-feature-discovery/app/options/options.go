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

package options

import (
	"fmt"
	"time"

	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"

	"github.com/Mellanox/nic-feature-discovery/pkg/utils/filesystem"
)

const (
	defaultNFDFeaturePath  = "/etc/node-feature-discovery/features.d/"
	defaultFeatureFileName = "nvidia-com-nic-feature-discovery.features"
)

// New creates new Options
func New() *Options {
	return &Options{
		NFDFeaturesPath:     defaultNFDFeaturePath,
		FeatureFileName:     defaultFeatureFileName,
		FeatureScanInterval: 1 * time.Minute,
		LogConfig:           logsapi.NewLoggingConfiguration(),
	}
}

// Options contains application options
type Options struct {
	NFDFeaturesPath     string
	FeatureFileName     string
	FeatureScanInterval time.Duration
	LogConfig           *logsapi.LoggingConfiguration
}

// GetFlagSets returns FlagSet for Options
func (o *Options) AddNamedFlagSets(sharedFS *cliflag.NamedFlagSets) {
	daemonFS := sharedFS.FlagSet("Feature Discovery Daemon")
	daemonFS.StringVar(
		&o.NFDFeaturesPath, "nfd-features-path", o.NFDFeaturesPath, "node feature discovery local features path")
	daemonFS.StringVar(&o.FeatureFileName, "features-file-name", o.FeatureFileName, "features file name")
	daemonFS.DurationVar(&o.FeatureScanInterval, "features-scan-interval", o.FeatureScanInterval, "features scan interval")

	logFS := sharedFS.FlagSet("Logging")
	logsapi.AddFlags(o.LogConfig, logFS)
	logs.AddFlags(logFS, logs.SkipLoggingConfigurationFlags())

	generalFS := sharedFS.FlagSet("General")
	_ = generalFS.Bool("version", false, "print version and exit")
	_ = generalFS.BoolP("help", "h", false, "print help and exit")
}

// Validate registered options
func (o *Options) Validate() error {
	var err error

	if err = logsapi.ValidateAndApply(o.LogConfig, nil); err != nil {
		return fmt.Errorf("failed to validate logging flags. %w", err)
	}

	if err = filesystem.FolderExist(o.NFDFeaturesPath); err != nil {
		return fmt.Errorf("failed to validate NFD features path. %w", err)
	}

	if o.FeatureFileName == "" {
		return fmt.Errorf("feature file name cannot be empty")
	}

	return err
}
