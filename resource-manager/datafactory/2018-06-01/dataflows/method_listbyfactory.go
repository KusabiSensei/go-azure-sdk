package dataflows

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByFactoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataFlowResource
}

type ListByFactoryCompleteResult struct {
	Items []DataFlowResource
}

// ListByFactory ...
func (c DataFlowsClient) ListByFactory(ctx context.Context, id FactoryId) (result ListByFactoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/dataflows", id.ID()),
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
		Values *[]DataFlowResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByFactoryComplete retrieves all the results into a single object
func (c DataFlowsClient) ListByFactoryComplete(ctx context.Context, id FactoryId) (ListByFactoryCompleteResult, error) {
	return c.ListByFactoryCompleteMatchingPredicate(ctx, id, DataFlowResourceOperationPredicate{})
}

// ListByFactoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataFlowsClient) ListByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate DataFlowResourceOperationPredicate) (result ListByFactoryCompleteResult, err error) {
	items := make([]DataFlowResource, 0)

	resp, err := c.ListByFactory(ctx, id)
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

	result = ListByFactoryCompleteResult{
		Items: items,
	}
	return
}
