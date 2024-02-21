package securitymlanalyticssettings

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SecurityMLAnalyticsSetting
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SecurityMLAnalyticsSetting
}

// List ...
func (c SecurityMLAnalyticsSettingsClient) List(ctx context.Context, id WorkspaceId) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]SecurityMLAnalyticsSetting, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalSecurityMLAnalyticsSettingImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for SecurityMLAnalyticsSetting (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListComplete retrieves all the results into a single object
func (c SecurityMLAnalyticsSettingsClient) ListComplete(ctx context.Context, id WorkspaceId) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, SecurityMLAnalyticsSettingOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SecurityMLAnalyticsSettingsClient) ListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate SecurityMLAnalyticsSettingOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]SecurityMLAnalyticsSetting, 0)

	resp, err := c.List(ctx, id)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
