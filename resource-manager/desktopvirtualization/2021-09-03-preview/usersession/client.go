package usersession

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSessionClient struct {
	Client *resourcemanager.Client
}

func NewUserSessionClientWithBaseURI(api environments.Api) (*UserSessionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "usersession", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSessionClient: %+v", err)
	}

	return &UserSessionClient{
		Client: client,
	}, nil
}
