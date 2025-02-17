package elasticsanskus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticSanSkusClient struct {
	Client *resourcemanager.Client
}

func NewElasticSanSkusClientWithBaseURI(api environments.Api) (*ElasticSanSkusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "elasticsanskus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ElasticSanSkusClient: %+v", err)
	}

	return &ElasticSanSkusClient{
		Client: client,
	}, nil
}
