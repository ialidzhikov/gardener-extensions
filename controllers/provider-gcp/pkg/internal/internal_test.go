// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInternal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GCP Internal Suite")
}

var _ = Describe("Infrastructure", func() {
	var (
		projectID          string
		serviceAccountData []byte
		serviceAccount     *ServiceAccount
	)

	BeforeEach(func() {
		projectID = "project"
		serviceAccountData = []byte(fmt.Sprintf(`{"project_id": "%s"}`, projectID))
		serviceAccount = &ServiceAccount{ProjectID: projectID, Raw: serviceAccountData}
	})

	Describe("#ComputeTerraformerVariablesEnvironment", func() {
		It("should correctly compute the terraformer variables environment", func() {
			variablesEnvironment, err := TerraformerVariablesEnvironmentFromServiceAccount(serviceAccount)
			Expect(err).NotTo(HaveOccurred())

			Expect(variablesEnvironment).To(Equal(map[string]string{
				TerraformVarServiceAccount: fmt.Sprintf(`{"project_id":"%s"}`, projectID),
			}))
		})
	})
})