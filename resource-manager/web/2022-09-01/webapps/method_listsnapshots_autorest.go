package webapps

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSnapshotsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Snapshot

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListSnapshotsOperationResponse, error)
}

type ListSnapshotsCompleteResult struct {
	Items []Snapshot
}

func (r ListSnapshotsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListSnapshotsOperationResponse) LoadMore(ctx context.Context) (resp ListSnapshotsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListSnapshots ...
func (c WebAppsClient) ListSnapshots(ctx context.Context, id SiteId) (resp ListSnapshotsOperationResponse, err error) {
	req, err := c.preparerForListSnapshots(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshots", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshots", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListSnapshots(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshots", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListSnapshots prepares the ListSnapshots request.
func (c WebAppsClient) preparerForListSnapshots(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/snapshots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListSnapshotsWithNextLink prepares the ListSnapshots request with the given nextLink token.
func (c WebAppsClient) preparerForListSnapshotsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSnapshots handles the response to the ListSnapshots request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSnapshots(resp *http.Response) (result ListSnapshotsOperationResponse, err error) {
	type page struct {
		Values   []Snapshot `json:"value"`
		NextLink *string    `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListSnapshotsOperationResponse, err error) {
			req, err := c.preparerForListSnapshotsWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshots", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshots", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListSnapshots(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSnapshots", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListSnapshotsComplete retrieves all of the results into a single object
func (c WebAppsClient) ListSnapshotsComplete(ctx context.Context, id SiteId) (ListSnapshotsCompleteResult, error) {
	return c.ListSnapshotsCompleteMatchingPredicate(ctx, id, SnapshotOperationPredicate{})
}

// ListSnapshotsCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c WebAppsClient) ListSnapshotsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate SnapshotOperationPredicate) (resp ListSnapshotsCompleteResult, err error) {
	items := make([]Snapshot, 0)

	page, err := c.ListSnapshots(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListSnapshotsCompleteResult{
		Items: items,
	}
	return out, nil
}
