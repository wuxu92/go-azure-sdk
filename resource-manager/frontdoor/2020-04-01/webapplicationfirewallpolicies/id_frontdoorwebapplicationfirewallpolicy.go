package webapplicationfirewallpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = FrontDoorWebApplicationFirewallPolicyId{}

// FrontDoorWebApplicationFirewallPolicyId is a struct representing the Resource ID for a Front Door Web Application Firewall Policy
type FrontDoorWebApplicationFirewallPolicyId struct {
	SubscriptionId    string
	ResourceGroupName string
	PolicyName        string
}

// NewFrontDoorWebApplicationFirewallPolicyID returns a new FrontDoorWebApplicationFirewallPolicyId struct
func NewFrontDoorWebApplicationFirewallPolicyID(subscriptionId string, resourceGroupName string, policyName string) FrontDoorWebApplicationFirewallPolicyId {
	return FrontDoorWebApplicationFirewallPolicyId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PolicyName:        policyName,
	}
}

// ParseFrontDoorWebApplicationFirewallPolicyID parses 'input' into a FrontDoorWebApplicationFirewallPolicyId
func ParseFrontDoorWebApplicationFirewallPolicyID(input string) (*FrontDoorWebApplicationFirewallPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(FrontDoorWebApplicationFirewallPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FrontDoorWebApplicationFirewallPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PolicyName, ok = parsed.Parsed["policyName"]; !ok {
		return nil, fmt.Errorf("the segment 'policyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseFrontDoorWebApplicationFirewallPolicyIDInsensitively parses 'input' case-insensitively into a FrontDoorWebApplicationFirewallPolicyId
// note: this method should only be used for API response data and not user input
func ParseFrontDoorWebApplicationFirewallPolicyIDInsensitively(input string) (*FrontDoorWebApplicationFirewallPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(FrontDoorWebApplicationFirewallPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FrontDoorWebApplicationFirewallPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PolicyName, ok = parsed.Parsed["policyName"]; !ok {
		return nil, fmt.Errorf("the segment 'policyName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateFrontDoorWebApplicationFirewallPolicyID checks that 'input' can be parsed as a Front Door Web Application Firewall Policy ID
func ValidateFrontDoorWebApplicationFirewallPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFrontDoorWebApplicationFirewallPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Front Door Web Application Firewall Policy ID
func (id FrontDoorWebApplicationFirewallPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Front Door Web Application Firewall Policy ID
func (id FrontDoorWebApplicationFirewallPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticFrontDoorWebApplicationFirewallPolicies", "frontDoorWebApplicationFirewallPolicies", "frontDoorWebApplicationFirewallPolicies"),
		resourceids.UserSpecifiedSegment("policyName", "policyValue"),
	}
}

// String returns a human-readable description of this Front Door Web Application Firewall Policy ID
func (id FrontDoorWebApplicationFirewallPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Policy Name: %q", id.PolicyName),
	}
	return fmt.Sprintf("Front Door Web Application Firewall Policy (%s)", strings.Join(components, "\n"))
}
