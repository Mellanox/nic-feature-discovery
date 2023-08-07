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

package label

import (
	"fmt"
	"reflect"
)

// Label is a key=value Label
type Label struct {
	Key   string
	Value string
}

// String implements fmt.Stringer interface
func (l *Label) String() string {
	return fmt.Sprintf("%s=%s", l.Key, l.Value)
}

// NewLabel creates a new *Label
func NewLabel(key, value string) Label {
	return Label{Key: key, Value: value}
}

type Set interface {
	Equal(other Set) bool
	AddLabel(l Label)
	AsLabels() []Label
	AsMap() map[string]string
}

func NewSet(labels ...Label) Set {
	lm := make(map[string]string)

	for _, l := range labels {
		lm[l.Key] = l.Value
	}

	return &labelSet{
		labelMap: lm,
	}
}

type labelSet struct {
	labelMap map[string]string
}

func (ls labelSet) Equal(other Set) bool {
	return reflect.DeepEqual(ls.labelMap, other.AsMap())
}

func (ls labelSet) AddLabel(l Label) {
	ls.labelMap[l.Key] = l.Value
}

func (ls labelSet) AsLabels() []Label {
	labels := make([]Label, 0, len(ls.labelMap))
	for k, v := range ls.labelMap {
		labels = append(labels, NewLabel(k, v))
	}

	return labels
}

func (ls labelSet) AsMap() map[string]string {
	m := make(map[string]string)

	for k, v := range ls.labelMap {
		m[k] = v
	}

	return m
}
