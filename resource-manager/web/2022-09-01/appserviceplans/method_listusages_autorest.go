package appserviceplans

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

type ListUsagesOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]CsmUsageQuota

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListUsagesOperationResponse, error)
}

type ListUsagesCompleteResult struct {
	Items []CsmUsageQuota
}

func (r ListUsagesOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListUsagesOperationResponse) LoadMore(ctx context.Context) (resp ListUsagesOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListUsagesOperationOptions struct {
	Filter *string
}

func DefaultListUsagesOperationOptions() ListUsagesOperationOptions {
	return ListUsagesOperationOptions{}
}

func (o ListUsagesOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListUsagesOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListUsages ...
func (c AppServicePlansClient) ListUsages(ctx context.Context, id ServerFarmId, options ListUsagesOperationOptions) (resp ListUsagesOperationResponse, err error) {
	req, err := c.preparerForListUsages(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListUsages", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListUsages", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListUsages(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListUsages", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListUsages prepares the ListUsages request.
func (c AppServicePlansClient) preparerForListUsages(ctx context.Context, id ServerFarmId, options ListUsagesOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListUsagesWithNextLink prepares the ListUsages request with the given nextLink token.
func (c AppServicePlansClient) preparerForListUsagesWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListUsages handles the response to the ListUsages request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForListUsages(resp *http.Response) (result ListUsagesOperationResponse, err error) {
	type page struct {
		Values   []CsmUsageQuota `json:"value"`
		NextLink *string         `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListUsagesOperationResponse, err error) {
			req, err := c.preparerForListUsagesWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListUsages", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListUsages", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListUsages(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListUsages", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListUsagesComplete retrieves all of the results into a single object
func (c AppServicePlansClient) ListUsagesComplete(ctx context.Context, id ServerFarmId, options ListUsagesOperationOptions) (ListUsagesCompleteResult, error) {
	return c.ListUsagesCompleteMatchingPredicate(ctx, id, options, CsmUsageQuotaOperationPredicate{})
}

// ListUsagesCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AppServicePlansClient) ListUsagesCompleteMatchingPredicate(ctx context.Context, id ServerFarmId, options ListUsagesOperationOptions, predicate CsmUsageQuotaOperationPredicate) (resp ListUsagesCompleteResult, err error) {
	items := make([]CsmUsageQuota, 0)

	page, err := c.ListUsages(ctx, id, options)
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

	out := ListUsagesCompleteResult{
		Items: items,
	}
	return out, nil
}
