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

package daemon_test

import (
	"context"
	"sync"
	"time"

	"github.com/Mellanox/nic-feature-discovery/pkg/daemon"
	"github.com/Mellanox/nic-feature-discovery/pkg/feature"
	"github.com/Mellanox/nic-feature-discovery/pkg/label"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	featureMocks "github.com/Mellanox/nic-feature-discovery/pkg/feature/mocks"
	writerMocks "github.com/Mellanox/nic-feature-discovery/pkg/writer/mocks"
)

var _ = Describe("Daemon test", func() {
	Context("basic Run", func() {
		It("runs successfully", func() {
			mockWriter := writerMocks.NewLabelWriter(GinkgoT())
			mockSource := featureMocks.NewSource(GinkgoT())
			mockFeature := featureMocks.NewFeature(GinkgoT())
			testLabels := []label.Label{label.NewLabel("foo", "bar")}
			mockSource.On("Name").Return("mockSource")
			mockSource.On("Discover", mock.Anything).Return([]feature.Feature{mockFeature}, nil)
			mockFeature.On("Labels").Return(testLabels)
			mockWriter.On("Write", testLabels).Return(nil)

			d := daemon.New(50*time.Millisecond, mockWriter, []feature.Source{mockSource})
			ctx, cFunc := context.WithCancel(context.Background())

			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer GinkgoRecover()
				defer wg.Done()
				d.Run(ctx)
			}()

			Eventually(func() int {
				return len(mockWriter.Calls)
			}).WithTimeout(200 * time.Millisecond).Should(BeNumerically(">", 1))

			cFunc()
			wg.Wait()
		})
	})
})
