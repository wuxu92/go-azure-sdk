package staticsites

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

type GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StaticSiteUserProvidedFunctionAppARMResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse, error)
}

type GetUserProvidedFunctionAppsForStaticSiteBuildCompleteResult struct {
	Items []StaticSiteUserProvidedFunctionAppARMResource
}

func (r GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse) LoadMore(ctx context.Context) (resp GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// GetUserProvidedFunctionAppsForStaticSiteBuild ...
func (c StaticSitesClient) GetUserProvidedFunctionAppsForStaticSiteBuild(ctx context.Context, id BuildId) (resp GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse, err error) {
	req, err := c.preparerForGetUserProvidedFunctionAppsForStaticSiteBuild(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSiteBuild", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSiteBuild", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForGetUserProvidedFunctionAppsForStaticSiteBuild(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSiteBuild", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForGetUserProvidedFunctionAppsForStaticSiteBuild prepares the GetUserProvidedFunctionAppsForStaticSiteBuild request.
func (c StaticSitesClient) preparerForGetUserProvidedFunctionAppsForStaticSiteBuild(ctx context.Context, id BuildId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/userProvidedFunctionApps", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForGetUserProvidedFunctionAppsForStaticSiteBuildWithNextLink prepares the GetUserProvidedFunctionAppsForStaticSiteBuild request with the given nextLink token.
func (c StaticSitesClient) preparerForGetUserProvidedFunctionAppsForStaticSiteBuildWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForGetUserProvidedFunctionAppsForStaticSiteBuild handles the response to the GetUserProvidedFunctionAppsForStaticSiteBuild request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForGetUserProvidedFunctionAppsForStaticSiteBuild(resp *http.Response) (result GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse, err error) {
	type page struct {
		Values   []StaticSiteUserProvidedFunctionAppARMResource `json:"value"`
		NextLink *string                                        `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result GetUserProvidedFunctionAppsForStaticSiteBuildOperationResponse, err error) {
			req, err := c.preparerForGetUserProvidedFunctionAppsForStaticSiteBuildWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSiteBuild", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSiteBuild", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForGetUserProvidedFunctionAppsForStaticSiteBuild(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "GetUserProvidedFunctionAppsForStaticSiteBuild", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// GetUserProvidedFunctionAppsForStaticSiteBuildComplete retrieves all of the results into a single object
func (c StaticSitesClient) GetUserProvidedFunctionAppsForStaticSiteBuildComplete(ctx context.Context, id BuildId) (GetUserProvidedFunctionAppsForStaticSiteBuildCompleteResult, error) {
	return c.GetUserProvidedFunctionAppsForStaticSiteBuildCompleteMatchingPredicate(ctx, id, StaticSiteUserProvidedFunctionAppARMResourceOperationPredicate{})
}

// GetUserProvidedFunctionAppsForStaticSiteBuildCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StaticSitesClient) GetUserProvidedFunctionAppsForStaticSiteBuildCompleteMatchingPredicate(ctx context.Context, id BuildId, predicate StaticSiteUserProvidedFunctionAppARMResourceOperationPredicate) (resp GetUserProvidedFunctionAppsForStaticSiteBuildCompleteResult, err error) {
	items := make([]StaticSiteUserProvidedFunctionAppARMResource, 0)

	page, err := c.GetUserProvidedFunctionAppsForStaticSiteBuild(ctx, id)
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

	out := GetUserProvidedFunctionAppsForStaticSiteBuildCompleteResult{
		Items: items,
	}
	return out, nil
}
