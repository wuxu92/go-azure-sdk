package officeconsents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeConsentsClient struct {
	Client *resourcemanager.Client
}

func NewOfficeConsentsClientWithBaseURI(sdkApi sdkEnv.Api) (*OfficeConsentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "officeconsents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OfficeConsentsClient: %+v", err)
	}

	return &OfficeConsentsClient{
		Client: client,
	}, nil
}
