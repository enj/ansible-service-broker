//
// Copyright (c) 2018 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package adapters

import (
	"testing"

	ft "github.com/openshift/ansible-service-broker/pkg/fusortest"
)

func TestSpecLabel(t *testing.T) {
	ft.AssertEqual(t, BundleSpecLabel, "com.redhat.apb.spec", "spec label does not match dockerhub")
}
