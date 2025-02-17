package sapsupportedsku

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SAPSupportedSkuClient struct {
	Client *resourcemanager.Client
}

func NewSAPSupportedSkuClientWithBaseURI(api environments.Api) (*SAPSupportedSkuClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "sapsupportedsku", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SAPSupportedSkuClient: %+v", err)
	}

	return &SAPSupportedSkuClient{
		Client: client,
	}, nil
}
