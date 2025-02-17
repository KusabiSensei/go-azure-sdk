package managedinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOutboundNetworkDependenciesByManagedInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OutboundEnvironmentEndpoint
}

type ListOutboundNetworkDependenciesByManagedInstanceCompleteResult struct {
	Items []OutboundEnvironmentEndpoint
}

// ListOutboundNetworkDependenciesByManagedInstance ...
func (c ManagedInstancesClient) ListOutboundNetworkDependenciesByManagedInstance(ctx context.Context, id ManagedInstanceId) (result ListOutboundNetworkDependenciesByManagedInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/outboundNetworkDependenciesEndpoints", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]OutboundEnvironmentEndpoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutboundNetworkDependenciesByManagedInstanceComplete retrieves all the results into a single object
func (c ManagedInstancesClient) ListOutboundNetworkDependenciesByManagedInstanceComplete(ctx context.Context, id ManagedInstanceId) (ListOutboundNetworkDependenciesByManagedInstanceCompleteResult, error) {
	return c.ListOutboundNetworkDependenciesByManagedInstanceCompleteMatchingPredicate(ctx, id, OutboundEnvironmentEndpointOperationPredicate{})
}

// ListOutboundNetworkDependenciesByManagedInstanceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedInstancesClient) ListOutboundNetworkDependenciesByManagedInstanceCompleteMatchingPredicate(ctx context.Context, id ManagedInstanceId, predicate OutboundEnvironmentEndpointOperationPredicate) (result ListOutboundNetworkDependenciesByManagedInstanceCompleteResult, err error) {
	items := make([]OutboundEnvironmentEndpoint, 0)

	resp, err := c.ListOutboundNetworkDependenciesByManagedInstance(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListOutboundNetworkDependenciesByManagedInstanceCompleteResult{
		Items: items,
	}
	return
}
