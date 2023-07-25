package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = NetworkFeatureId{}

// NetworkFeatureId is a struct representing the Resource ID for a Network Feature
type NetworkFeatureId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SiteName           string
	NetworkFeatureName string
}

// NewNetworkFeatureID returns a new NetworkFeatureId struct
func NewNetworkFeatureID(subscriptionId string, resourceGroupName string, siteName string, networkFeatureName string) NetworkFeatureId {
	return NetworkFeatureId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SiteName:           siteName,
		NetworkFeatureName: networkFeatureName,
	}
}

// ParseNetworkFeatureID parses 'input' into a NetworkFeatureId
func ParseNetworkFeatureID(input string) (*NetworkFeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(NetworkFeatureId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NetworkFeatureId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.NetworkFeatureName, ok = parsed.Parsed["networkFeatureName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "networkFeatureName", *parsed)
	}

	return &id, nil
}

// ParseNetworkFeatureIDInsensitively parses 'input' case-insensitively into a NetworkFeatureId
// note: this method should only be used for API response data and not user input
func ParseNetworkFeatureIDInsensitively(input string) (*NetworkFeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(NetworkFeatureId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NetworkFeatureId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.NetworkFeatureName, ok = parsed.Parsed["networkFeatureName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "networkFeatureName", *parsed)
	}

	return &id, nil
}

// ValidateNetworkFeatureID checks that 'input' can be parsed as a Network Feature ID
func ValidateNetworkFeatureID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNetworkFeatureID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Network Feature ID
func (id NetworkFeatureId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/networkFeatures/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.NetworkFeatureName)
}

// Segments returns a slice of Resource ID Segments which comprise this Network Feature ID
func (id NetworkFeatureId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticNetworkFeatures", "networkFeatures", "networkFeatures"),
		resourceids.UserSpecifiedSegment("networkFeatureName", "networkFeatureValue"),
	}
}

// String returns a human-readable description of this Network Feature ID
func (id NetworkFeatureId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Network Feature Name: %q", id.NetworkFeatureName),
	}
	return fmt.Sprintf("Network Feature (%s)", strings.Join(components, "\n"))
}
