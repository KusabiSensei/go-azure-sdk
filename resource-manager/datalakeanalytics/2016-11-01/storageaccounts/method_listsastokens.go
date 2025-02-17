package storageaccounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSasTokensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SasTokenInformation
}

type ListSasTokensCompleteResult struct {
	Items []SasTokenInformation
}

// ListSasTokens ...
func (c StorageAccountsClient) ListSasTokens(ctx context.Context, id ContainerId) (result ListSasTokensOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listSasTokens", id.ID()),
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
		Values *[]SasTokenInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSasTokensComplete retrieves all the results into a single object
func (c StorageAccountsClient) ListSasTokensComplete(ctx context.Context, id ContainerId) (ListSasTokensCompleteResult, error) {
	return c.ListSasTokensCompleteMatchingPredicate(ctx, id, SasTokenInformationOperationPredicate{})
}

// ListSasTokensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StorageAccountsClient) ListSasTokensCompleteMatchingPredicate(ctx context.Context, id ContainerId, predicate SasTokenInformationOperationPredicate) (result ListSasTokensCompleteResult, err error) {
	items := make([]SasTokenInformation, 0)

	resp, err := c.ListSasTokens(ctx, id)
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

	result = ListSasTokensCompleteResult{
		Items: items,
	}
	return
}
