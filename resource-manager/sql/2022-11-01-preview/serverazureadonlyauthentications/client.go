package serverazureadonlyauthentications

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerAzureADOnlyAuthenticationsClient struct {
	Client *resourcemanager.Client
}

func NewServerAzureADOnlyAuthenticationsClientWithBaseURI(api environments.Api) (*ServerAzureADOnlyAuthenticationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serverazureadonlyauthentications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerAzureADOnlyAuthenticationsClient: %+v", err)
	}

	return &ServerAzureADOnlyAuthenticationsClient{
		Client: client,
	}, nil
}
