package appserviceenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostingEnvironmentDiagnostics struct {
	DiagnosticsOutput *string `json:"diagnosticsOutput,omitempty"`
	Name              *string `json:"name,omitempty"`
}
