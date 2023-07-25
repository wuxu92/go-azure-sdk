package resourceproviders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ValidateResponse
}

// Validate ...
func (c ResourceProvidersClient) Validate(ctx context.Context, id commonids.ResourceGroupId, input ValidateRequest) (result ValidateOperationResponse, err error) {
	req, err := c.preparerForValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "Validate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "Validate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForValidate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "Validate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForValidate prepares the Validate request.
func (c ResourceProvidersClient) preparerForValidate(ctx context.Context, id commonids.ResourceGroupId, input ValidateRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Web/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForValidate handles the response to the Validate request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForValidate(resp *http.Response) (result ValidateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
