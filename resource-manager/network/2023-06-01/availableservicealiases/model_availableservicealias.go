package availableservicealiases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableServiceAlias struct {
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	Type         *string `json:"type,omitempty"`
}
