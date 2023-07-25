package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteSiteExtensionOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteSiteExtension ...
func (c WebAppsClient) DeleteSiteExtension(ctx context.Context, id SiteExtensionId) (result DeleteSiteExtensionOperationResponse, err error) {
	req, err := c.preparerForDeleteSiteExtension(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSiteExtension", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSiteExtension", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteSiteExtension(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteSiteExtension", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteSiteExtension prepares the DeleteSiteExtension request.
func (c WebAppsClient) preparerForDeleteSiteExtension(ctx context.Context, id SiteExtensionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeleteSiteExtension handles the response to the DeleteSiteExtension request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteSiteExtension(resp *http.Response) (result DeleteSiteExtensionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
