package manageddatabasemoveoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseMoveOperationsClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseMoveOperationsClientWithBaseURI(api environments.Api) (*ManagedDatabaseMoveOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "manageddatabasemoveoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseMoveOperationsClient: %+v", err)
	}

	return &ManagedDatabaseMoveOperationsClient{
		Client: client,
	}, nil
}
