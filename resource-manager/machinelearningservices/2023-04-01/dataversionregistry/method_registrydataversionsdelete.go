package dataversionregistry

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryDataVersionsDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// RegistryDataVersionsDelete ...
func (c DataVersionRegistryClient) RegistryDataVersionsDelete(ctx context.Context, id VersionId) (result RegistryDataVersionsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// RegistryDataVersionsDeleteThenPoll performs RegistryDataVersionsDelete then polls until it's completed
func (c DataVersionRegistryClient) RegistryDataVersionsDeleteThenPoll(ctx context.Context, id VersionId) error {
	result, err := c.RegistryDataVersionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryDataVersionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after RegistryDataVersionsDelete: %+v", err)
	}

	return nil
}