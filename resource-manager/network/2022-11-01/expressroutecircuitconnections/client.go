package expressroutecircuitconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCircuitConnectionsClientWithBaseURI(api environments.Api) (*ExpressRouteCircuitConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecircuitconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCircuitConnectionsClient: %+v", err)
	}

	return &ExpressRouteCircuitConnectionsClient{
		Client: client,
	}, nil
}
