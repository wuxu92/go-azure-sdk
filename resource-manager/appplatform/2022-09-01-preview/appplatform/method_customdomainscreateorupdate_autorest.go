package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDomainsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CustomDomainsCreateOrUpdate ...
func (c AppPlatformClient) CustomDomainsCreateOrUpdate(ctx context.Context, id DomainId, input CustomDomainResource) (result CustomDomainsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForCustomDomainsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "CustomDomainsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCustomDomainsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "CustomDomainsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CustomDomainsCreateOrUpdateThenPoll performs CustomDomainsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) CustomDomainsCreateOrUpdateThenPoll(ctx context.Context, id DomainId, input CustomDomainResource) error {
	result, err := c.CustomDomainsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CustomDomainsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CustomDomainsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForCustomDomainsCreateOrUpdate prepares the CustomDomainsCreateOrUpdate request.
func (c AppPlatformClient) preparerForCustomDomainsCreateOrUpdate(ctx context.Context, id DomainId, input CustomDomainResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCustomDomainsCreateOrUpdate sends the CustomDomainsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForCustomDomainsCreateOrUpdate(ctx context.Context, req *http.Request) (future CustomDomainsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
