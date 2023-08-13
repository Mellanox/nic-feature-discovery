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

package label_test

import (
	"github.com/Mellanox/nic-feature-discovery/pkg/label"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("label tests", func() {
	Context("Label", func() {
		It("works", func() {
			l := label.NewLabel("foo", "bar")
			Expect(l.Key).To(Equal("foo"))
			Expect(l.Value).To(Equal("bar"))
			Expect(l.String()).To(Equal("foo=bar"))
		})
	})

	Context("Set", func() {
		var labels = []label.Label{label.NewLabel("foo", "bar"), label.NewLabel("baz", "fuz")}
		var otherLabels = []label.Label{label.NewLabel("foo", "bar"), label.NewLabel("baz", "raz")}

		Context("Equal()", func() {
			It("returns true if sets are equal", func() {
				s1 := label.NewSet(labels...)
				s2 := label.NewSet(labels...)
				Expect(s1.Equal(s2)).To(BeTrue())
				Expect(s1.Equal(s1)).To(BeTrue())
			})

			It("returns false if sets are not equal", func() {
				s1 := label.NewSet(labels...)
				s2 := label.NewSet(otherLabels...)
				Expect(s1.Equal(s2)).To(BeFalse())
			})
		})

		Context("AsLabels()", func() {
			It("returns expected list of labels", func() {
				set := label.NewSet(labels...)
				ls := set.AsLabels()
				Expect(ls).To(HaveLen(len(labels)))
				Expect(label.NewSet(ls...).Equal(set)).To(BeTrue())
			})
		})

		Context("AsMap()", func() {
			It("returns expected map", func() {
				set := label.NewSet(labels...)
				m := set.AsMap()
				Expect(m).To(HaveLen(len(labels)))
				Expect(m["foo"]).To(Equal("bar"))
				Expect(m["baz"]).To(Equal("fuz"))
			})
		})
	})
})
