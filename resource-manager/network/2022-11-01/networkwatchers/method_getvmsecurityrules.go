package networkwatchers

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

type GetVMSecurityRulesOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetVMSecurityRules ...
func (c NetworkWatchersClient) GetVMSecurityRules(ctx context.Context, id NetworkWatcherId, input SecurityGroupViewParameters) (result GetVMSecurityRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/securityGroupView", id.ID()),
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

// GetVMSecurityRulesThenPoll performs GetVMSecurityRules then polls until it's completed
func (c NetworkWatchersClient) GetVMSecurityRulesThenPoll(ctx context.Context, id NetworkWatcherId, input SecurityGroupViewParameters) error {
	result, err := c.GetVMSecurityRules(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GetVMSecurityRules: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GetVMSecurityRules: %+v", err)
	}

	return nil
}
