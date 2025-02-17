package databaseencryptionprotectorrevert

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

type DatabaseEncryptionProtectorsRevertOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// DatabaseEncryptionProtectorsRevert ...
func (c DatabaseEncryptionProtectorRevertClient) DatabaseEncryptionProtectorsRevert(ctx context.Context, id DatabaseId) (result DatabaseEncryptionProtectorsRevertOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/encryptionProtector/current/revert", id.ID()),
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

// DatabaseEncryptionProtectorsRevertThenPoll performs DatabaseEncryptionProtectorsRevert then polls until it's completed
func (c DatabaseEncryptionProtectorRevertClient) DatabaseEncryptionProtectorsRevertThenPoll(ctx context.Context, id DatabaseId) error {
	result, err := c.DatabaseEncryptionProtectorsRevert(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DatabaseEncryptionProtectorsRevert: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DatabaseEncryptionProtectorsRevert: %+v", err)
	}

	return nil
}
