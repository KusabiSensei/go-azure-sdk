package models

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelsClient struct {
	Client *resourcemanager.Client
}

func NewModelsClientWithBaseURI(api environments.Api) (*ModelsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "models", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ModelsClient: %+v", err)
	}

	return &ModelsClient{
		Client: client,
	}, nil
}
