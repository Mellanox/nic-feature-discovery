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

package app

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"

	// register json format for logger
	_ "k8s.io/component-base/logs/json/register"

	"github.com/Mellanox/nic-feature-discovery/cmd/nic-feature-discovery/app/options"
	"github.com/Mellanox/nic-feature-discovery/pkg/utils/signals"
	"github.com/Mellanox/nic-feature-discovery/pkg/utils/version"
)

// NewNICFeatureDiscoveryCommand creates a new command
func NewNICFeatureDiscoveryCommand() *cobra.Command {
	opts := options.New()
	ctx := signals.SetupShutdownSignals()

	cmd := &cobra.Command{
		Use:          "nic-feature-discovery",
		Long:         `NVIDIA NIC Feature Discovery`,
		SilenceUsage: true,
		Version:      version.GetVersionString(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Validate(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}
			klog.EnableContextualLogging(true)

			return RunNICFeatureDiscovery(klog.NewContext(ctx, klog.NewKlogr()), opts)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	sharedFS := cliflag.NamedFlagSets{}
	opts.AddNamedFlagSets(&sharedFS)

	cmdFS := cmd.PersistentFlags()
	for _, f := range sharedFS.FlagSets {
		cmdFS.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, sharedFS, cols)

	return cmd
}

// RunNICFeatureDiscovery runs nic feature discovery daemon
func RunNICFeatureDiscovery(ctx context.Context, opts *options.Options) error {
	logger := logr.FromContextOrDiscard(ctx)
	logger.Info("start NIC feature discovery", "Options", opts)
	<-ctx.Done()

	return nil
}
