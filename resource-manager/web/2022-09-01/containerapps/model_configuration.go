package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Configuration struct {
	ActiveRevisionsMode *ActiveRevisionsMode   `json:"activeRevisionsMode,omitempty"`
	Ingress             *Ingress               `json:"ingress,omitempty"`
	Registries          *[]RegistryCredentials `json:"registries,omitempty"`
	Secrets             *[]Secret              `json:"secrets,omitempty"`
}
