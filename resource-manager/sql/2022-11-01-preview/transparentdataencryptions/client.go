package transparentdataencryptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransparentDataEncryptionsClient struct {
	Client *resourcemanager.Client
}

func NewTransparentDataEncryptionsClientWithBaseURI(api environments.Api) (*TransparentDataEncryptionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "transparentdataencryptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransparentDataEncryptionsClient: %+v", err)
	}

	return &TransparentDataEncryptionsClient{
		Client: client,
	}, nil
}
