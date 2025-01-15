package fleets

import (
	"encoding/json"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetOSProfile struct {
	AdminPassword               *string               `json:"adminPassword,omitempty"`
	AdminUsername               *string               `json:"adminUsername,omitempty"`
	AllowExtensionOperations    *bool                 `json:"allowExtensionOperations,omitempty"`
	ComputerNamePrefix          *string               `json:"computerNamePrefix,omitempty"`
	CustomData                  *string               `json:"customData,omitempty"`
	LinuxConfiguration          *LinuxConfiguration   `json:"linuxConfiguration"`
	RequireGuestProvisionSignal *bool                 `json:"requireGuestProvisionSignal,omitempty"`
	Secrets                     *[]VaultSecretGroup   `json:"secrets,omitempty"`
	WindowsConfiguration        *WindowsConfiguration `json:"windowsConfiguration"`
}

var _ json.Marshaler = VirtualMachineScaleSetOSProfile{}

func (s VirtualMachineScaleSetOSProfile) MarshalJSON() ([]byte, error) {
	return nullable.MarshalNullableStruct(s)
}
