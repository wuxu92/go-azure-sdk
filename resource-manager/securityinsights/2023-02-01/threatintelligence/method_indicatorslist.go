package threatintelligence

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

type IndicatorsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ThreatIntelligenceInformation
}

type IndicatorsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ThreatIntelligenceInformation
}

type IndicatorsListOperationOptions struct {
	Filter  *string
	Orderby *string
	Top     *int64
}

func DefaultIndicatorsListOperationOptions() IndicatorsListOperationOptions {
	return IndicatorsListOperationOptions{}
}

func (o IndicatorsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o IndicatorsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o IndicatorsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// IndicatorsList ...
func (c ThreatIntelligenceClient) IndicatorsList(ctx context.Context, id WorkspaceId, options IndicatorsListOperationOptions) (result IndicatorsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators", id.ID()),
		OptionsObject: options,
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

	temp := make([]ThreatIntelligenceInformation, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalThreatIntelligenceInformationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for ThreatIntelligenceInformation (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// IndicatorsListComplete retrieves all the results into a single object
func (c ThreatIntelligenceClient) IndicatorsListComplete(ctx context.Context, id WorkspaceId, options IndicatorsListOperationOptions) (IndicatorsListCompleteResult, error) {
	return c.IndicatorsListCompleteMatchingPredicate(ctx, id, options, ThreatIntelligenceInformationOperationPredicate{})
}

// IndicatorsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreatIntelligenceClient) IndicatorsListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options IndicatorsListOperationOptions, predicate ThreatIntelligenceInformationOperationPredicate) (result IndicatorsListCompleteResult, err error) {
	items := make([]ThreatIntelligenceInformation, 0)

	resp, err := c.IndicatorsList(ctx, id, options)
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

	result = IndicatorsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
