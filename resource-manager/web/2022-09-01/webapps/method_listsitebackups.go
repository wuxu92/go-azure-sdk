package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteBackupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BackupItem
}

type ListSiteBackupsCompleteResult struct {
	Items []BackupItem
}

// ListSiteBackups ...
func (c WebAppsClient) ListSiteBackups(ctx context.Context, id SiteId) (result ListSiteBackupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listbackups", id.ID()),
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
		Values *[]BackupItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteBackupsComplete retrieves all the results into a single object
func (c WebAppsClient) ListSiteBackupsComplete(ctx context.Context, id SiteId) (ListSiteBackupsCompleteResult, error) {
	return c.ListSiteBackupsCompleteMatchingPredicate(ctx, id, BackupItemOperationPredicate{})
}

// ListSiteBackupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListSiteBackupsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate BackupItemOperationPredicate) (result ListSiteBackupsCompleteResult, err error) {
	items := make([]BackupItem, 0)

	resp, err := c.ListSiteBackups(ctx, id)
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

	result = ListSiteBackupsCompleteResult{
		Items: items,
	}
	return
}