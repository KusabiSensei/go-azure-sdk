package virtualnetworkgateways

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

type GetVpnclientIPsecParametersOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetVpnclientIPsecParameters ...
func (c VirtualNetworkGatewaysClient) GetVpnclientIPsecParameters(ctx context.Context, id VirtualNetworkGatewayId) (result GetVpnclientIPsecParametersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/getvpnclientipsecparameters", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// GetVpnclientIPsecParametersThenPoll performs GetVpnclientIPsecParameters then polls until it's completed
func (c VirtualNetworkGatewaysClient) GetVpnclientIPsecParametersThenPoll(ctx context.Context, id VirtualNetworkGatewayId) error {
	result, err := c.GetVpnclientIPsecParameters(ctx, id)
	if err != nil {
		return fmt.Errorf("performing GetVpnclientIPsecParameters: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GetVpnclientIPsecParameters: %+v", err)
	}

	return nil
}
