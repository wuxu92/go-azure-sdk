package itemlevelrecoveryconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionOperationResponse struct {
	HttpResponse *http.Response
}

// Provision ...
func (c ItemLevelRecoveryConnectionsClient) Provision(ctx context.Context, id RecoveryPointId, input ILRRequestResource) (result ProvisionOperationResponse, err error) {
	req, err := c.preparerForProvision(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient", "Provision", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient", "Provision", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProvision(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "itemlevelrecoveryconnections.ItemLevelRecoveryConnectionsClient", "Provision", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProvision prepares the Provision request.
func (c ItemLevelRecoveryConnectionsClient) preparerForProvision(ctx context.Context, id RecoveryPointId, input ILRRequestResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/provisionInstantItemRecovery", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProvision handles the response to the Provision request. The method always
// closes the http.Response Body.
func (c ItemLevelRecoveryConnectionsClient) responderForProvision(resp *http.Response) (result ProvisionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
