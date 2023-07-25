package appserviceenvironments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWorkerPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkerPoolResource
}

// UpdateWorkerPool ...
func (c AppServiceEnvironmentsClient) UpdateWorkerPool(ctx context.Context, id WorkerPoolId, input WorkerPoolResource) (result UpdateWorkerPoolOperationResponse, err error) {
	req, err := c.preparerForUpdateWorkerPool(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateWorkerPool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateWorkerPool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateWorkerPool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateWorkerPool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateWorkerPool prepares the UpdateWorkerPool request.
func (c AppServiceEnvironmentsClient) preparerForUpdateWorkerPool(ctx context.Context, id WorkerPoolId, input WorkerPoolResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateWorkerPool handles the response to the UpdateWorkerPool request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForUpdateWorkerPool(resp *http.Response) (result UpdateWorkerPoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
