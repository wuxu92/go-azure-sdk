package fleetupdatestrategies

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

type ListByFleetOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FleetUpdateStrategy

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByFleetOperationResponse, error)
}

type ListByFleetCompleteResult struct {
	Items []FleetUpdateStrategy
}

func (r ListByFleetOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByFleetOperationResponse) LoadMore(ctx context.Context) (resp ListByFleetOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByFleet ...
func (c FleetUpdateStrategiesClient) ListByFleet(ctx context.Context, id FleetId) (resp ListByFleetOperationResponse, err error) {
	req, err := c.preparerForListByFleet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fleetupdatestrategies.FleetUpdateStrategiesClient", "ListByFleet", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "fleetupdatestrategies.FleetUpdateStrategiesClient", "ListByFleet", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByFleet(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fleetupdatestrategies.FleetUpdateStrategiesClient", "ListByFleet", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByFleet prepares the ListByFleet request.
func (c FleetUpdateStrategiesClient) preparerForListByFleet(ctx context.Context, id FleetId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/updateStrategies", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByFleetWithNextLink prepares the ListByFleet request with the given nextLink token.
func (c FleetUpdateStrategiesClient) preparerForListByFleetWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByFleet handles the response to the ListByFleet request. The method always
// closes the http.Response Body.
func (c FleetUpdateStrategiesClient) responderForListByFleet(resp *http.Response) (result ListByFleetOperationResponse, err error) {
	type page struct {
		Values   []FleetUpdateStrategy `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByFleetOperationResponse, err error) {
			req, err := c.preparerForListByFleetWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "fleetupdatestrategies.FleetUpdateStrategiesClient", "ListByFleet", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "fleetupdatestrategies.FleetUpdateStrategiesClient", "ListByFleet", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByFleet(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "fleetupdatestrategies.FleetUpdateStrategiesClient", "ListByFleet", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByFleetComplete retrieves all of the results into a single object
func (c FleetUpdateStrategiesClient) ListByFleetComplete(ctx context.Context, id FleetId) (ListByFleetCompleteResult, error) {
	return c.ListByFleetCompleteMatchingPredicate(ctx, id, FleetUpdateStrategyOperationPredicate{})
}

// ListByFleetCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c FleetUpdateStrategiesClient) ListByFleetCompleteMatchingPredicate(ctx context.Context, id FleetId, predicate FleetUpdateStrategyOperationPredicate) (resp ListByFleetCompleteResult, err error) {
	items := make([]FleetUpdateStrategy, 0)

	page, err := c.ListByFleet(ctx, id)
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

	out := ListByFleetCompleteResult{
		Items: items,
	}
	return out, nil
}
