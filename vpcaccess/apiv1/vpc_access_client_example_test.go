// Copyright 2022 Google LLC
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

package vpcaccess_test

import (
	"context"

	vpcaccess "cloud.google.com/go/vpcaccess/apiv1"
	"google.golang.org/api/iterator"
	vpcaccesspb "google.golang.org/genproto/googleapis/cloud/vpcaccess/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := vpcaccess.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateConnector() {
	ctx := context.Background()
	c, err := vpcaccess.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vpcaccesspb.CreateConnectorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vpcaccess/v1#CreateConnectorRequest.
	}
	op, err := c.CreateConnector(ctx, req)
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

func ExampleClient_GetConnector() {
	ctx := context.Background()
	c, err := vpcaccess.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vpcaccesspb.GetConnectorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vpcaccess/v1#GetConnectorRequest.
	}
	resp, err := c.GetConnector(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListConnectors() {
	ctx := context.Background()
	c, err := vpcaccess.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vpcaccesspb.ListConnectorsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vpcaccess/v1#ListConnectorsRequest.
	}
	it := c.ListConnectors(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_DeleteConnector() {
	ctx := context.Background()
	c, err := vpcaccess.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &vpcaccesspb.DeleteConnectorRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/vpcaccess/v1#DeleteConnectorRequest.
	}
	op, err := c.DeleteConnector(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
