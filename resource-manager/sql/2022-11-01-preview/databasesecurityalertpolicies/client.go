package databasesecurityalertpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseSecurityAlertPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseSecurityAlertPoliciesClientWithBaseURI(api environments.Api) (*DatabaseSecurityAlertPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "databasesecurityalertpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseSecurityAlertPoliciesClient: %+v", err)
	}

	return &DatabaseSecurityAlertPoliciesClient{
		Client: client,
	}, nil
}
