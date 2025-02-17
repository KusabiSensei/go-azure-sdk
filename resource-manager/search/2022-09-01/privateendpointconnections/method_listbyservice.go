package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type ListByServiceCompleteResult struct {
	Items []PrivateEndpointConnection
}

type ListByServiceOperationOptions struct {
	XMsClientRequestId *string
}

func DefaultListByServiceOperationOptions() ListByServiceOperationOptions {
	return ListByServiceOperationOptions{}
}

func (o ListByServiceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsClientRequestId != nil {
		out.Append("x-ms-client-request-id", fmt.Sprintf("%v", *o.XMsClientRequestId))
	}
	return &out
}

func (o ListByServiceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByServiceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ListByService ...
func (c PrivateEndpointConnectionsClient) ListByService(ctx context.Context, id SearchServiceId, options ListByServiceOperationOptions) (result ListByServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
		OptionsObject: options,
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
		Values *[]PrivateEndpointConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByServiceComplete retrieves all the results into a single object
func (c PrivateEndpointConnectionsClient) ListByServiceComplete(ctx context.Context, id SearchServiceId, options ListByServiceOperationOptions) (ListByServiceCompleteResult, error) {
	return c.ListByServiceCompleteMatchingPredicate(ctx, id, options, PrivateEndpointConnectionOperationPredicate{})
}

// ListByServiceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateEndpointConnectionsClient) ListByServiceCompleteMatchingPredicate(ctx context.Context, id SearchServiceId, options ListByServiceOperationOptions, predicate PrivateEndpointConnectionOperationPredicate) (result ListByServiceCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.ListByService(ctx, id, options)
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

	result = ListByServiceCompleteResult{
		Items: items,
	}
	return
}
