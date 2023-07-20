package packetcorecontrolplanes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlanesClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreControlPlanesClientWithBaseURI(api environments.Api) (*PacketCoreControlPlanesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "packetcorecontrolplanes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreControlPlanesClient: %+v", err)
	}

	return &PacketCoreControlPlanesClient{
		Client: client,
	}, nil
}