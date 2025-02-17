package jobruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobRunsClient struct {
	Client *resourcemanager.Client
}

func NewJobRunsClientWithBaseURI(api environments.Api) (*JobRunsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jobruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobRunsClient: %+v", err)
	}

	return &JobRunsClient{
		Client: client,
	}, nil
}
