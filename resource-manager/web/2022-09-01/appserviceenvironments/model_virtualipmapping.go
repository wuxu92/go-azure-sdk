package appserviceenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualIPMapping struct {
	InUse             *bool   `json:"inUse,omitempty"`
	InternalHTTPPort  *int64  `json:"internalHttpPort,omitempty"`
	InternalHTTPSPort *int64  `json:"internalHttpsPort,omitempty"`
	ServiceName       *string `json:"serviceName,omitempty"`
	VirtualIP         *string `json:"virtualIP,omitempty"`
}
