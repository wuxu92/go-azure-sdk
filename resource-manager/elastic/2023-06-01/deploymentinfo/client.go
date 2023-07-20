package deploymentinfo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentInfoClient struct {
	Client *resourcemanager.Client
}

func NewDeploymentInfoClientWithBaseURI(api environments.Api) (*DeploymentInfoClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "deploymentinfo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeploymentInfoClient: %+v", err)
	}

	return &DeploymentInfoClient{
		Client: client,
	}, nil
}