package blobauditing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerBlobAuditingPoliciesListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServerBlobAuditingPolicy
}

type ServerBlobAuditingPoliciesListByServerCompleteResult struct {
	Items []ServerBlobAuditingPolicy
}

// ServerBlobAuditingPoliciesListByServer ...
func (c BlobAuditingClient) ServerBlobAuditingPoliciesListByServer(ctx context.Context, id ServerId) (result ServerBlobAuditingPoliciesListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/auditingSettings", id.ID()),
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
		Values *[]ServerBlobAuditingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ServerBlobAuditingPoliciesListByServerComplete retrieves all the results into a single object
func (c BlobAuditingClient) ServerBlobAuditingPoliciesListByServerComplete(ctx context.Context, id ServerId) (ServerBlobAuditingPoliciesListByServerCompleteResult, error) {
	return c.ServerBlobAuditingPoliciesListByServerCompleteMatchingPredicate(ctx, id, ServerBlobAuditingPolicyOperationPredicate{})
}

// ServerBlobAuditingPoliciesListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BlobAuditingClient) ServerBlobAuditingPoliciesListByServerCompleteMatchingPredicate(ctx context.Context, id ServerId, predicate ServerBlobAuditingPolicyOperationPredicate) (result ServerBlobAuditingPoliciesListByServerCompleteResult, err error) {
	items := make([]ServerBlobAuditingPolicy, 0)

	resp, err := c.ServerBlobAuditingPoliciesListByServer(ctx, id)
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

	result = ServerBlobAuditingPoliciesListByServerCompleteResult{
		Items: items,
	}
	return
}
