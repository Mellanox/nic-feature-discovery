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

package common

import (
	"github.com/Mellanox/nic-feature-discovery/pkg/label"
)

// NewGenericFeature creates a new GenericFeature
func NewGenericFeature(name string) *GenericFeature {
	return &GenericFeature{
		name: name,
	}
}

// GenericFeature is a generic implementation of Feature
type GenericFeature struct {
	name   string
	labels []label.Label
}

// Name of Feature
func (gf *GenericFeature) Name() string {
	return gf.name
}

// Labels returns the list of Labels for Feature
func (gf *GenericFeature) Labels() []label.Label {
	return gf.labels
}

// AddLabel adds a key=value label to GenericFeature
func (gf *GenericFeature) AddLabel(k, v string) *GenericFeature {
	gf.labels = append(gf.labels, label.Label{Key: k, Value: v})

	return gf
}
