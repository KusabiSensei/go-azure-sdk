package loadbalancers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoadBalancersClient struct {
	Client *resourcemanager.Client
}

func NewLoadBalancersClientWithBaseURI(api environments.Api) (*LoadBalancersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "loadbalancers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LoadBalancersClient: %+v", err)
	}

	return &LoadBalancersClient{
		Client: client,
	}, nil
}
