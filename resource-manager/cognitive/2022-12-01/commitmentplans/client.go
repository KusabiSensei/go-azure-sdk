package commitmentplans

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentPlansClient struct {
	Client *resourcemanager.Client
}

func NewCommitmentPlansClientWithBaseURI(api environments.Api) (*CommitmentPlansClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "commitmentplans", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommitmentPlansClient: %+v", err)
	}

	return &CommitmentPlansClient{
		Client: client,
	}, nil
}
