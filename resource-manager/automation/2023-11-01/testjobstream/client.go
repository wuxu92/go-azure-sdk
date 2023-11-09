package testjobstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestJobStreamClient struct {
	Client *resourcemanager.Client
}

func NewTestJobStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*TestJobStreamClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "testjobstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TestJobStreamClient: %+v", err)
	}

	return &TestJobStreamClient{
		Client: client,
	}, nil
}
