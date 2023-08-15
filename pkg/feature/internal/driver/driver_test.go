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

package driver_test

import (
	"context"
	"errors"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Mellanox/nic-feature-discovery/pkg/dependencies"
	depMocks "github.com/Mellanox/nic-feature-discovery/pkg/dependencies/mocks"
	"github.com/Mellanox/nic-feature-discovery/pkg/feature"
	"github.com/Mellanox/nic-feature-discovery/pkg/feature/internal/driver"
	"github.com/Mellanox/nic-feature-discovery/pkg/label"
)

var _ = Describe("Driver Source Tests", func() {
	Context("Name()", func() {
		It("returns expected Source name", func() {
			ds := driver.NewDriverSource()
			Expect(ds.Name()).To(Equal(driver.SourceNameDriver))
		})
	})

	Context("Discover()", func() {
		var (
			mockOS *depMocks.Os
			ds     feature.Source
		)

		BeforeEach(func() {
			mockOS = depMocks.NewOs(GinkgoT())
			dependencies.OS = mockOS
			ds = driver.NewDriverSource()
		})

		It("discovers features successfully", func() {
			mockOS.On("Lstat", driver.Mlx5ModulePath).Return(nil, nil)
			mockOS.On("Lstat", driver.Mlx5ModuleVersion).Return(nil, nil)
			mockOS.On("ReadFile", driver.Mlx5ModuleVersion).Return([]byte("23.04-0.5.5.0"), nil)

			features, err := ds.Discover(context.TODO())

			Expect(err).ToNot(HaveOccurred())
			Expect(features).To(HaveLen(1))
			Expect(features[0].Name()).To(Equal(driver.FeatureNameMofedVersion))
			Expect(features[0].Labels()).To(HaveLen(1))
			Expect(features[0].Labels()[0]).To(
				BeEquivalentTo(label.NewLabel("nvidia.com/mofed.version", "23.04-0.5.5.0")))
		})

		It("discovers no features if driver is not loaded", func() {
			mockOS.On("Lstat", driver.Mlx5ModulePath).Return(nil, os.ErrNotExist)

			features, err := ds.Discover(context.TODO())

			Expect(err).ToNot(HaveOccurred())
			Expect(features).To(HaveLen(0))
		})

		It("discovers no features if version is not advertised by driver", func() {
			mockOS.On("Lstat", driver.Mlx5ModulePath).Return(nil, nil)
			mockOS.On("Lstat", driver.Mlx5ModuleVersion).Return(nil, os.ErrNotExist)

			features, err := ds.Discover(context.TODO())

			Expect(err).ToNot(HaveOccurred())
			Expect(features).To(HaveLen(0))
		})

		It("fails if reading driver version file fails", func() {
			mockOS.On("Lstat", driver.Mlx5ModulePath).Return(nil, nil)
			mockOS.On("Lstat", driver.Mlx5ModuleVersion).Return(nil, nil)
			mockOS.On("ReadFile", driver.Mlx5ModuleVersion).Return(nil, errors.New("some error"))

			_, err := ds.Discover(context.TODO())

			Expect(err).To(HaveOccurred())
		})

		It("fails if driver version file empty", func() {
			mockOS.On("Lstat", driver.Mlx5ModulePath).Return(nil, nil)
			mockOS.On("Lstat", driver.Mlx5ModuleVersion).Return(nil, nil)
			mockOS.On("ReadFile", driver.Mlx5ModuleVersion).Return([]byte(" "), nil)

			_, err := ds.Discover(context.TODO())

			Expect(err).To(HaveOccurred())
		})
	})
})
