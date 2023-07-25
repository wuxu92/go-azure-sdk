package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopContinuousWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
}

// StopContinuousWebJobSlot ...
func (c WebAppsClient) StopContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (result StopContinuousWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForStopContinuousWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopContinuousWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopContinuousWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStopContinuousWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "StopContinuousWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStopContinuousWebJobSlot prepares the StopContinuousWebJobSlot request.
func (c WebAppsClient) preparerForStopContinuousWebJobSlot(ctx context.Context, id SlotContinuousWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStopContinuousWebJobSlot handles the response to the StopContinuousWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForStopContinuousWebJobSlot(resp *http.Response) (result StopContinuousWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
