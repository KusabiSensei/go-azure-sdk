package instancefailovergroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstanceFailoverGroupsClient struct {
	Client *resourcemanager.Client
}

func NewInstanceFailoverGroupsClientWithBaseURI(api environments.Api) (*InstanceFailoverGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "instancefailovergroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InstanceFailoverGroupsClient: %+v", err)
	}

	return &InstanceFailoverGroupsClient{
		Client: client,
	}, nil
}
