// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package metricsscope_test

import (
	"context"

	metricsscope "cloud.google.com/go/monitoring/metricsscope/apiv1"
	metricsscopepb "google.golang.org/genproto/googleapis/monitoring/metricsscope/v1"
)

func ExampleNewMetricsScopesClient() {
	ctx := context.Background()
	c, err := metricsscope.NewMetricsScopesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleMetricsScopesClient_GetMetricsScope() {
	ctx := context.Background()
	c, err := metricsscope.NewMetricsScopesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metricsscopepb.GetMetricsScopeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetMetricsScope(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleMetricsScopesClient_ListMetricsScopesByMonitoredProject() {
	ctx := context.Background()
	c, err := metricsscope.NewMetricsScopesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metricsscopepb.ListMetricsScopesByMonitoredProjectRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListMetricsScopesByMonitoredProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleMetricsScopesClient_CreateMonitoredProject() {
	ctx := context.Background()
	c, err := metricsscope.NewMetricsScopesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metricsscopepb.CreateMonitoredProjectRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateMonitoredProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleMetricsScopesClient_DeleteMonitoredProject() {
	ctx := context.Background()
	c, err := metricsscope.NewMetricsScopesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &metricsscopepb.DeleteMonitoredProjectRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteMonitoredProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}