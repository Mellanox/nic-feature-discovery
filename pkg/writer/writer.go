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

package writer

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-logr/logr"
	"github.com/google/renameio/v2"

	"github.com/Mellanox/nic-feature-discovery/pkg/label"
)

type LabelWriter interface {
	// Write labels to feature file. if labels are the same, no write is performed
	Write(labels []label.Label) error
}

// labelWriter implements LabelWriter
type labelWriter struct {
	featureFilePath string
	log             logr.Logger
}

func NewLabelWriter(path string, log logr.Logger) LabelWriter {
	return &labelWriter{
		featureFilePath: path,
		log:             log,
	}
}

// Write labels to feature file. if labels are the same, no write is performed
func (lr *labelWriter) Write(labels []label.Label) error {
	curLabels, err := lr.getCurrentLabels()
	if err != nil {
		lr.log.Error(err, "failed to get current labels, will not attempt to determine if labels remained the same")
	}

	if err == nil {
		// check if labels changed
		if label.NewSet(labels...).Equal(label.NewSet(curLabels...)) {
			// labels equal no need to write
			lr.log.V(2).Info("current and new labels are equal, nothing to write.")

			return nil
		}
	}

	return lr.writeLabels(labels)
}

// getCurrentLabels returns current labels set in  feature file
func (lr *labelWriter) getCurrentLabels() ([]label.Label, error) {
	curLabels := make([]label.Label, 0)

	_, err := os.Lstat(lr.featureFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return curLabels, nil
		}

		return nil, err
	}

	data, err := os.ReadFile(lr.featureFilePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			// this is unexpected, return error
			return nil, fmt.Errorf("unexpected line format. %q", line)
		}
		curLabels = append(curLabels, label.NewLabel(kv[0], kv[1]))
	}

	return curLabels, nil
}

// writeLabels writes given labels to featureFilePath
func (lr *labelWriter) writeLabels(labels []label.Label) error {
	lines := make([]string, 0, len(labels))

	for _, l := range labels {
		lines = append(lines, l.String())
	}
	data := strings.Join(lines, "\n")
	data += "\n"

	err := renameio.WriteFile(lr.featureFilePath, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write labels to file(%s). %w", lr.featureFilePath, err)
	}

	return nil
}
