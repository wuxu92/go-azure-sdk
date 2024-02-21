package v2023_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/actions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/alertrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/alertruletemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/automationrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/bookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/dataconnectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/entitytypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/incidentalerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/incidentbookmarks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/incidentcomments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/incidententities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/incidentrelations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/incidents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/securitymlanalyticssettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/sentinelonboardingstates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/threatintelligence"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/watchlistitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-02-01/watchlists"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Actions                     *actions.ActionsClient
	AlertRuleTemplates          *alertruletemplates.AlertRuleTemplatesClient
	AlertRules                  *alertrules.AlertRulesClient
	AutomationRules             *automationrules.AutomationRulesClient
	Bookmarks                   *bookmarks.BookmarksClient
	DataConnectors              *dataconnectors.DataConnectorsClient
	EntityTypes                 *entitytypes.EntityTypesClient
	IncidentAlerts              *incidentalerts.IncidentAlertsClient
	IncidentBookmarks           *incidentbookmarks.IncidentBookmarksClient
	IncidentComments            *incidentcomments.IncidentCommentsClient
	IncidentEntities            *incidententities.IncidentEntitiesClient
	IncidentRelations           *incidentrelations.IncidentRelationsClient
	Incidents                   *incidents.IncidentsClient
	Metadata                    *metadata.MetadataClient
	SecurityMLAnalyticsSettings *securitymlanalyticssettings.SecurityMLAnalyticsSettingsClient
	SentinelOnboardingStates    *sentinelonboardingstates.SentinelOnboardingStatesClient
	ThreatIntelligence          *threatintelligence.ThreatIntelligenceClient
	WatchlistItems              *watchlistitems.WatchlistItemsClient
	Watchlists                  *watchlists.WatchlistsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	actionsClient, err := actions.NewActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Actions client: %+v", err)
	}
	configureFunc(actionsClient.Client)

	alertRuleTemplatesClient, err := alertruletemplates.NewAlertRuleTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRuleTemplates client: %+v", err)
	}
	configureFunc(alertRuleTemplatesClient.Client)

	alertRulesClient, err := alertrules.NewAlertRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertRules client: %+v", err)
	}
	configureFunc(alertRulesClient.Client)

	automationRulesClient, err := automationrules.NewAutomationRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutomationRules client: %+v", err)
	}
	configureFunc(automationRulesClient.Client)

	bookmarksClient, err := bookmarks.NewBookmarksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Bookmarks client: %+v", err)
	}
	configureFunc(bookmarksClient.Client)

	dataConnectorsClient, err := dataconnectors.NewDataConnectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataConnectors client: %+v", err)
	}
	configureFunc(dataConnectorsClient.Client)

	entityTypesClient, err := entitytypes.NewEntityTypesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntityTypes client: %+v", err)
	}
	configureFunc(entityTypesClient.Client)

	incidentAlertsClient, err := incidentalerts.NewIncidentAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentAlerts client: %+v", err)
	}
	configureFunc(incidentAlertsClient.Client)

	incidentBookmarksClient, err := incidentbookmarks.NewIncidentBookmarksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentBookmarks client: %+v", err)
	}
	configureFunc(incidentBookmarksClient.Client)

	incidentCommentsClient, err := incidentcomments.NewIncidentCommentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentComments client: %+v", err)
	}
	configureFunc(incidentCommentsClient.Client)

	incidentEntitiesClient, err := incidententities.NewIncidentEntitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentEntities client: %+v", err)
	}
	configureFunc(incidentEntitiesClient.Client)

	incidentRelationsClient, err := incidentrelations.NewIncidentRelationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncidentRelations client: %+v", err)
	}
	configureFunc(incidentRelationsClient.Client)

	incidentsClient, err := incidents.NewIncidentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Incidents client: %+v", err)
	}
	configureFunc(incidentsClient.Client)

	metadataClient, err := metadata.NewMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Metadata client: %+v", err)
	}
	configureFunc(metadataClient.Client)

	securityMLAnalyticsSettingsClient, err := securitymlanalyticssettings.NewSecurityMLAnalyticsSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityMLAnalyticsSettings client: %+v", err)
	}
	configureFunc(securityMLAnalyticsSettingsClient.Client)

	sentinelOnboardingStatesClient, err := sentinelonboardingstates.NewSentinelOnboardingStatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SentinelOnboardingStates client: %+v", err)
	}
	configureFunc(sentinelOnboardingStatesClient.Client)

	threatIntelligenceClient, err := threatintelligence.NewThreatIntelligenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreatIntelligence client: %+v", err)
	}
	configureFunc(threatIntelligenceClient.Client)

	watchlistItemsClient, err := watchlistitems.NewWatchlistItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WatchlistItems client: %+v", err)
	}
	configureFunc(watchlistItemsClient.Client)

	watchlistsClient, err := watchlists.NewWatchlistsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Watchlists client: %+v", err)
	}
	configureFunc(watchlistsClient.Client)

	return &Client{
		Actions:                     actionsClient,
		AlertRuleTemplates:          alertRuleTemplatesClient,
		AlertRules:                  alertRulesClient,
		AutomationRules:             automationRulesClient,
		Bookmarks:                   bookmarksClient,
		DataConnectors:              dataConnectorsClient,
		EntityTypes:                 entityTypesClient,
		IncidentAlerts:              incidentAlertsClient,
		IncidentBookmarks:           incidentBookmarksClient,
		IncidentComments:            incidentCommentsClient,
		IncidentEntities:            incidentEntitiesClient,
		IncidentRelations:           incidentRelationsClient,
		Incidents:                   incidentsClient,
		Metadata:                    metadataClient,
		SecurityMLAnalyticsSettings: securityMLAnalyticsSettingsClient,
		SentinelOnboardingStates:    sentinelOnboardingStatesClient,
		ThreatIntelligence:          threatIntelligenceClient,
		WatchlistItems:              watchlistItemsClient,
		Watchlists:                  watchlistsClient,
	}, nil
}
