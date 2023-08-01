package recommendations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HostingEnvironmentRecommendationId{}

// HostingEnvironmentRecommendationId is a struct representing the Resource ID for a Hosting Environment Recommendation
type HostingEnvironmentRecommendationId struct {
	SubscriptionId         string
	ResourceGroupName      string
	HostingEnvironmentName string
	RecommendationName     string
}

// NewHostingEnvironmentRecommendationID returns a new HostingEnvironmentRecommendationId struct
func NewHostingEnvironmentRecommendationID(subscriptionId string, resourceGroupName string, hostingEnvironmentName string, recommendationName string) HostingEnvironmentRecommendationId {
	return HostingEnvironmentRecommendationId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		HostingEnvironmentName: hostingEnvironmentName,
		RecommendationName:     recommendationName,
	}
}

// ParseHostingEnvironmentRecommendationID parses 'input' into a HostingEnvironmentRecommendationId
func ParseHostingEnvironmentRecommendationID(input string) (*HostingEnvironmentRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostingEnvironmentRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostingEnvironmentRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HostingEnvironmentName, ok = parsed.Parsed["hostingEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", *parsed)
	}

	if id.RecommendationName, ok = parsed.Parsed["recommendationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", *parsed)
	}

	return &id, nil
}

// ParseHostingEnvironmentRecommendationIDInsensitively parses 'input' case-insensitively into a HostingEnvironmentRecommendationId
// note: this method should only be used for API response data and not user input
func ParseHostingEnvironmentRecommendationIDInsensitively(input string) (*HostingEnvironmentRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostingEnvironmentRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostingEnvironmentRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HostingEnvironmentName, ok = parsed.Parsed["hostingEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", *parsed)
	}

	if id.RecommendationName, ok = parsed.Parsed["recommendationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", *parsed)
	}

	return &id, nil
}

// ValidateHostingEnvironmentRecommendationID checks that 'input' can be parsed as a Hosting Environment Recommendation ID
func ValidateHostingEnvironmentRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHostingEnvironmentRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hosting Environment Recommendation ID
func (id HostingEnvironmentRecommendationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/hostingEnvironments/%s/recommendations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HostingEnvironmentName, id.RecommendationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hosting Environment Recommendation ID
func (id HostingEnvironmentRecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticHostingEnvironments", "hostingEnvironments", "hostingEnvironments"),
		resourceids.UserSpecifiedSegment("hostingEnvironmentName", "hostingEnvironmentValue"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationName", "recommendationValue"),
	}
}

// String returns a human-readable description of this Hosting Environment Recommendation ID
func (id HostingEnvironmentRecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Hosting Environment Name: %q", id.HostingEnvironmentName),
		fmt.Sprintf("Recommendation Name: %q", id.RecommendationName),
	}
	return fmt.Sprintf("Hosting Environment Recommendation (%s)", strings.Join(components, "\n"))
}