package managementgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UbscriptionsDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UbscriptionsDeleteOperationOptions struct {
	CacheControl *string
}

func DefaultUbscriptionsDeleteOperationOptions() UbscriptionsDeleteOperationOptions {
	return UbscriptionsDeleteOperationOptions{}
}

func (o UbscriptionsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o UbscriptionsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o UbscriptionsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UbscriptionsDelete ...
func (c ManagementGroupsClient) UbscriptionsDelete(ctx context.Context, id SubscriptionId, options UbscriptionsDeleteOperationOptions) (result UbscriptionsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		Path:          id.ID(),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
