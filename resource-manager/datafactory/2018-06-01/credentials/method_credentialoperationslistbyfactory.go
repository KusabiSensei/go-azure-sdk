package credentials

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialOperationsListByFactoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedIdentityCredentialResource
}

type CredentialOperationsListByFactoryCompleteResult struct {
	Items []ManagedIdentityCredentialResource
}

// CredentialOperationsListByFactory ...
func (c CredentialsClient) CredentialOperationsListByFactory(ctx context.Context, id FactoryId) (result CredentialOperationsListByFactoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/credentials", id.ID()),
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
		Values *[]ManagedIdentityCredentialResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CredentialOperationsListByFactoryComplete retrieves all the results into a single object
func (c CredentialsClient) CredentialOperationsListByFactoryComplete(ctx context.Context, id FactoryId) (CredentialOperationsListByFactoryCompleteResult, error) {
	return c.CredentialOperationsListByFactoryCompleteMatchingPredicate(ctx, id, ManagedIdentityCredentialResourceOperationPredicate{})
}

// CredentialOperationsListByFactoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CredentialsClient) CredentialOperationsListByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate ManagedIdentityCredentialResourceOperationPredicate) (result CredentialOperationsListByFactoryCompleteResult, err error) {
	items := make([]ManagedIdentityCredentialResource, 0)

	resp, err := c.CredentialOperationsListByFactory(ctx, id)
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

	result = CredentialOperationsListByFactoryCompleteResult{
		Items: items,
	}
	return
}
