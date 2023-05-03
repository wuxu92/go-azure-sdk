package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcesListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ResourcesListByHostPoolCompleteResult struct {
	Items []PrivateLinkResource
}

// ResourcesListByHostPool ...
func (c PrivateLinkClient) ResourcesListByHostPool(ctx context.Context, id HostPoolId) (result ResourcesListByHostPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateLinkResources", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourcesListByHostPoolComplete retrieves all the results into a single object
func (c PrivateLinkClient) ResourcesListByHostPoolComplete(ctx context.Context, id HostPoolId) (ResourcesListByHostPoolCompleteResult, error) {
	return c.ResourcesListByHostPoolCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// ResourcesListByHostPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) ResourcesListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, predicate PrivateLinkResourceOperationPredicate) (result ResourcesListByHostPoolCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ResourcesListByHostPool(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ResourcesListByHostPoolCompleteResult{
		Items: items,
	}
	return
}