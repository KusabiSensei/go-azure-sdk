package commitmentplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateAssociationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateOrUpdateAssociation ...
func (c CommitmentPlansClient) CreateOrUpdateAssociation(ctx context.Context, id AccountAssociationId, input CommitmentPlanAccountAssociation) (result CreateOrUpdateAssociationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// CreateOrUpdateAssociationThenPoll performs CreateOrUpdateAssociation then polls until it's completed
func (c CommitmentPlansClient) CreateOrUpdateAssociationThenPoll(ctx context.Context, id AccountAssociationId, input CommitmentPlanAccountAssociation) error {
	result, err := c.CreateOrUpdateAssociation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateAssociation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateAssociation: %+v", err)
	}

	return nil
}
