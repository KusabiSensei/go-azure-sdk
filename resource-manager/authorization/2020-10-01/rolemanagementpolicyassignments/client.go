package rolemanagementpolicyassignments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyAssignmentsClient struct {
	Client *resourcemanager.Client
}

func NewRoleManagementPolicyAssignmentsClientWithBaseURI(api environments.Api) (*RoleManagementPolicyAssignmentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "rolemanagementpolicyassignments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementPolicyAssignmentsClient: %+v", err)
	}

	return &RoleManagementPolicyAssignmentsClient{
		Client: client,
	}, nil
}
