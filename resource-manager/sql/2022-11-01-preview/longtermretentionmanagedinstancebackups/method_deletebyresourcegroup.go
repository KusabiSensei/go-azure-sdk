package longtermretentionmanagedinstancebackups

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

type DeleteByResourceGroupOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// DeleteByResourceGroup ...
func (c LongTermRetentionManagedInstanceBackupsClient) DeleteByResourceGroup(ctx context.Context, id LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId) (result DeleteByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.ID(),
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

// DeleteByResourceGroupThenPoll performs DeleteByResourceGroup then polls until it's completed
func (c LongTermRetentionManagedInstanceBackupsClient) DeleteByResourceGroupThenPoll(ctx context.Context, id LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId) error {
	result, err := c.DeleteByResourceGroup(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteByResourceGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DeleteByResourceGroup: %+v", err)
	}

	return nil
}
