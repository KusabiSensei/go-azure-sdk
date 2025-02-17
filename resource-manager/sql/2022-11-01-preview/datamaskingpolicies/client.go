package datamaskingpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewDataMaskingPoliciesClientWithBaseURI(api environments.Api) (*DataMaskingPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "datamaskingpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataMaskingPoliciesClient: %+v", err)
	}

	return &DataMaskingPoliciesClient{
		Client: client,
	}, nil
}
