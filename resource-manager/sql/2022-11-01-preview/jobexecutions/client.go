package jobexecutions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobExecutionsClient struct {
	Client *resourcemanager.Client
}

func NewJobExecutionsClientWithBaseURI(api environments.Api) (*JobExecutionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jobexecutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobExecutionsClient: %+v", err)
	}

	return &JobExecutionsClient{
		Client: client,
	}, nil
}
