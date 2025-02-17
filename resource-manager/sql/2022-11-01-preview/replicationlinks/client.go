package replicationlinks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationLinksClient struct {
	Client *resourcemanager.Client
}

func NewReplicationLinksClientWithBaseURI(api environments.Api) (*ReplicationLinksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "replicationlinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationLinksClient: %+v", err)
	}

	return &ReplicationLinksClient{
		Client: client,
	}, nil
}
