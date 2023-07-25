package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDomainOwnershipIdentifierOperationResponse struct {
	HttpResponse *http.Response
	Model        *Identifier
}

// UpdateDomainOwnershipIdentifier ...
func (c WebAppsClient) UpdateDomainOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId, input Identifier) (result UpdateDomainOwnershipIdentifierOperationResponse, err error) {
	req, err := c.preparerForUpdateDomainOwnershipIdentifier(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDomainOwnershipIdentifier", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDomainOwnershipIdentifier", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateDomainOwnershipIdentifier(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateDomainOwnershipIdentifier", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateDomainOwnershipIdentifier prepares the UpdateDomainOwnershipIdentifier request.
func (c WebAppsClient) preparerForUpdateDomainOwnershipIdentifier(ctx context.Context, id DomainOwnershipIdentifierId, input Identifier) (*http.Request, error) {
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

// responderForUpdateDomainOwnershipIdentifier handles the response to the UpdateDomainOwnershipIdentifier request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateDomainOwnershipIdentifier(resp *http.Response) (result UpdateDomainOwnershipIdentifierOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
