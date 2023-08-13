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

package writer_test

import (
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"

	"github.com/Mellanox/nic-feature-discovery/pkg/label"
	"github.com/Mellanox/nic-feature-discovery/pkg/writer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// labelsFromFile returns Labels that are encoded as k=v formatted lines in file
func labelsFromFile(file string) []label.Label {
	curLabels := make([]label.Label, 0)

	data, err := os.ReadFile(file)
	Expect(err).ToNot(HaveOccurred())

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		kv := strings.Split(line, "=")
		Expect(kv).To(HaveLen(2))
		curLabels = append(curLabels, label.NewLabel(kv[0], kv[1]))
	}

	return curLabels
}

// fileExists returns true if file exists else returns false
func fileExists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		GinkgoT().Fatalf("unexpected Stat error. %w", err)
	}
	return true
}

var _ = Describe("writer tests", func() {
	var (
		tmpFeatureDir   string
		tmpFeatureFile  string
		writerUnderTest writer.LabelWriter
		emptyLabels     []label.Label
		firstLabelList  []label.Label
		secondLabelList []label.Label
	)

	BeforeEach(func() {
		tmpFeatureDir = GinkgoT().TempDir()
		tmpFeatureFile = filepath.Join(tmpFeatureDir, "test-feature-file")
		writerUnderTest = writer.NewLabelWriter(tmpFeatureFile, klog.NewKlogr())
		emptyLabels = make([]label.Label, 0)
		firstLabelList = []label.Label{label.NewLabel("foo", "bar"), label.NewLabel("baz", "bat")}
		secondLabelList = []label.Label{label.NewLabel("foo", "bar"), label.NewLabel("three", "four")}
	})

	Context("Write() tests", func() {
		It("writes feature file if no feature file exists", func() {
			Expect(writerUnderTest.Write(firstLabelList)).ToNot(HaveOccurred())
			writtenLabels := labelsFromFile(tmpFeatureFile)
			Expect(writtenLabels).To(Equal(firstLabelList))
		})

		It("writes feature file if labels change", func() {
			Expect(writerUnderTest.Write(firstLabelList)).ToNot(HaveOccurred())
			Expect(writerUnderTest.Write(secondLabelList)).ToNot(HaveOccurred())
			writtenLabels := labelsFromFile(tmpFeatureFile)
			Expect(writtenLabels).To(Equal(secondLabelList))
		})

		It("removes feature file if no labels specified", func() {
			Expect(writerUnderTest.Write(firstLabelList)).ToNot(HaveOccurred())
			ExpectWithOffset(1, fileExists(tmpFeatureFile)).To(BeTrue())
			Expect(writerUnderTest.Write(emptyLabels)).ToNot(HaveOccurred())
			ExpectWithOffset(1, fileExists(tmpFeatureFile)).To(BeFalse())
		})

		It("labels dont change in file if they are the same", func() {
			Expect(writerUnderTest.Write(firstLabelList)).ToNot(HaveOccurred())
			Expect(writerUnderTest.Write(firstLabelList)).ToNot(HaveOccurred())
			writtenLabels := labelsFromFile(tmpFeatureFile)
			Expect(writtenLabels).To(Equal(firstLabelList))
		})

		It("file does not exist if no labels specified", func() {
			Expect(writerUnderTest.Write(emptyLabels)).ToNot(HaveOccurred())
			ExpectWithOffset(1, fileExists(tmpFeatureFile)).To(BeFalse())
		})

		It("does not fail if current file content is not as expected", func() {
			Expect(os.WriteFile(tmpFeatureFile, []byte("foobarbaz\n"), 0644)).ToNot(HaveOccurred())
			Expect(writerUnderTest.Write(firstLabelList)).ToNot(HaveOccurred())
			writtenLabels := labelsFromFile(tmpFeatureFile)
			Expect(writtenLabels).To(Equal(firstLabelList))
		})
	})
})
