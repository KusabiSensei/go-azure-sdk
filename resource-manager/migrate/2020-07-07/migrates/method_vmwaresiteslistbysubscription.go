package migrates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareSitesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VMwareSite
}

type VMwareSitesListBySubscriptionCompleteResult struct {
	Items []VMwareSite
}

// VMwareSitesListBySubscription ...
func (c MigratesClient) VMwareSitesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result VMwareSitesListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.OffAzure/vmwareSites", id.ID()),
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
		Values *[]VMwareSite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VMwareSitesListBySubscriptionComplete retrieves all the results into a single object
func (c MigratesClient) VMwareSitesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (VMwareSitesListBySubscriptionCompleteResult, error) {
	return c.VMwareSitesListBySubscriptionCompleteMatchingPredicate(ctx, id, VMwareSiteOperationPredicate{})
}

// VMwareSitesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MigratesClient) VMwareSitesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate VMwareSiteOperationPredicate) (result VMwareSitesListBySubscriptionCompleteResult, err error) {
	items := make([]VMwareSite, 0)

	resp, err := c.VMwareSitesListBySubscription(ctx, id)
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

	result = VMwareSitesListBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
