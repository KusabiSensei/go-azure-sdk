package guestagents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestAgentsClient struct {
	Client *resourcemanager.Client
}

func NewGuestAgentsClientWithBaseURI(api environments.Api) (*GuestAgentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "guestagents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GuestAgentsClient: %+v", err)
	}

	return &GuestAgentsClient{
		Client: client,
	}, nil
}
