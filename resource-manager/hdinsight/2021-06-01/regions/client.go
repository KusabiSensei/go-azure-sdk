package regions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionsClient struct {
	Client *resourcemanager.Client
}

func NewRegionsClientWithBaseURI(api environments.Api) (*RegionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "regions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegionsClient: %+v", err)
	}

	return &RegionsClient{
		Client: client,
	}, nil
}
