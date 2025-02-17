package managedinstancelongtermretentionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceLongTermRetentionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstanceLongTermRetentionPoliciesClientWithBaseURI(api environments.Api) (*ManagedInstanceLongTermRetentionPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedinstancelongtermretentionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstanceLongTermRetentionPoliciesClient: %+v", err)
	}

	return &ManagedInstanceLongTermRetentionPoliciesClient{
		Client: client,
	}, nil
}
