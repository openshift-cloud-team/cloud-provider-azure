// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package routetableclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/metrics"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

type Client struct {
	*armnetwork.RouteTablesClient
	subscriptionID string
	tracer         tracing.Tracer
}

func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}
	tr := options.TracingProvider.NewTracer(utils.ModuleName, utils.ModuleVersion)

	client, err := armnetwork.NewRouteTablesClient(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{
		RouteTablesClient: client,
		subscriptionID:    subscriptionID,
		tracer:            tr,
	}, nil
}

const GetOperationName = "RouteTablesClient.Get"

// Get gets the RouteTable
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string) (result *armnetwork.RouteTable, err error) {

	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "RouteTable", "get")
	defer func() { metricsCtx.Observe(ctx, err) }()
	ctx, endSpan := runtime.StartSpan(ctx, GetOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := client.RouteTablesClient.Get(ctx, resourceGroupName, resourceName, nil)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.RouteTable, nil
}

const CreateOrUpdateOperationName = "RouteTablesClient.Create"

// CreateOrUpdate creates or updates a RouteTable.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, resource armnetwork.RouteTable) (result *armnetwork.RouteTable, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "RouteTable", "create_or_update")
	defer func() { metricsCtx.Observe(ctx, err) }()
	ctx, endSpan := runtime.StartSpan(ctx, CreateOrUpdateOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := utils.NewPollerWrapper(client.RouteTablesClient.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.RouteTable, nil
	}
	return nil, nil
}

const DeleteOperationName = "RouteTablesClient.Delete"

// Delete deletes a RouteTable by name.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, resourceName string) (err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "RouteTable", "delete")
	defer func() { metricsCtx.Observe(ctx, err) }()
	ctx, endSpan := runtime.StartSpan(ctx, DeleteOperationName, client.tracer, nil)
	defer endSpan(err)
	_, err = utils.NewPollerWrapper(client.BeginDelete(ctx, resourceGroupName, resourceName, nil)).WaitforPollerResp(ctx)
	return err
}

const ListOperationName = "RouteTablesClient.List"

// List gets a list of RouteTable in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string) (result []*armnetwork.RouteTable, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "RouteTable", "list")
	defer func() { metricsCtx.Observe(ctx, err) }()
	ctx, endSpan := runtime.StartSpan(ctx, ListOperationName, client.tracer, nil)
	defer endSpan(err)
	pager := client.RouteTablesClient.NewListPager(resourceGroupName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}
