package expressroutecircuits

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitsClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCircuitsClientWithBaseURI(api environments.Api) (*ExpressRouteCircuitsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecircuits", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCircuitsClient: %+v", err)
	}

	return &ExpressRouteCircuitsClient{
		Client: client,
	}, nil
}
