package volumegroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkRule struct {
	Action *Action `json:"action,omitempty"`
	Id     string  `json:"id"`
	State  *State  `json:"state,omitempty"`
}
