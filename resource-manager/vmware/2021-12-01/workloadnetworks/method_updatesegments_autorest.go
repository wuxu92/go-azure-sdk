package workloadnetworks

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

type UpdateSegmentsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateSegments ...
func (c WorkloadNetworksClient) UpdateSegments(ctx context.Context, id SegmentId, input WorkloadNetworkSegment) (result UpdateSegmentsOperationResponse, err error) {
	req, err := c.preparerForUpdateSegments(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateSegments", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateSegments(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateSegments", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateSegmentsThenPoll performs UpdateSegments then polls until it's completed
func (c WorkloadNetworksClient) UpdateSegmentsThenPoll(ctx context.Context, id SegmentId, input WorkloadNetworkSegment) error {
	result, err := c.UpdateSegments(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateSegments: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateSegments: %+v", err)
	}

	return nil
}

// preparerForUpdateSegments prepares the UpdateSegments request.
func (c WorkloadNetworksClient) preparerForUpdateSegments(ctx context.Context, id SegmentId, input WorkloadNetworkSegment) (*http.Request, error) {
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

// senderForUpdateSegments sends the UpdateSegments request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForUpdateSegments(ctx context.Context, req *http.Request) (future UpdateSegmentsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
