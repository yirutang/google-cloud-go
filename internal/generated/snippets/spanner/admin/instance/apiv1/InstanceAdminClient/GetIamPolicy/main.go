// Copyright 2021 Google LLC
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

// [START spanner_generated_spanner_admin_instance_apiv1_InstanceAdminClient_GetIamPolicy]

package main

import (
	"context"

	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func main() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END spanner_generated_spanner_admin_instance_apiv1_InstanceAdminClient_GetIamPolicy]