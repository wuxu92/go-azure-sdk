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

type GetProcessDumpOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetProcessDump ...
func (c WebAppsClient) GetProcessDump(ctx context.Context, id ProcessId) (result GetProcessDumpOperationResponse, err error) {
	req, err := c.preparerForGetProcessDump(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessDump", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessDump", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetProcessDump(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetProcessDump", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetProcessDump prepares the GetProcessDump request.
func (c WebAppsClient) preparerForGetProcessDump(ctx context.Context, id ProcessId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dump", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetProcessDump handles the response to the GetProcessDump request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetProcessDump(resp *http.Response) (result GetProcessDumpOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
