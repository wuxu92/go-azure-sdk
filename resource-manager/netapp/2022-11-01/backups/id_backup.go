package backups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = BackupId{}

// BackupId is a struct representing the Resource ID for a Backup
type BackupId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	CapacityPoolName  string
	VolumeName        string
	BackupName        string
}

// NewBackupID returns a new BackupId struct
func NewBackupID(subscriptionId string, resourceGroupName string, netAppAccountName string, capacityPoolName string, volumeName string, backupName string) BackupId {
	return BackupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		CapacityPoolName:  capacityPoolName,
		VolumeName:        volumeName,
		BackupName:        backupName,
	}
}

// ParseBackupID parses 'input' into a BackupId
func ParseBackupID(input string) (*BackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BackupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBackupIDInsensitively parses 'input' case-insensitively into a BackupId
// note: this method should only be used for API response data and not user input
func ParseBackupIDInsensitively(input string) (*BackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(BackupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BackupId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BackupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetAppAccountName, ok = input.Parsed["netAppAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", input)
	}

	if id.CapacityPoolName, ok = input.Parsed["capacityPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", input)
	}

	if id.VolumeName, ok = input.Parsed["volumeName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "volumeName", input)
	}

	if id.BackupName, ok = input.Parsed["backupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "backupName", input)
	}

	return nil
}

// ValidateBackupID checks that 'input' can be parsed as a Backup ID
func ValidateBackupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBackupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Backup ID
func (id BackupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/capacityPools/%s/volumes/%s/backups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.CapacityPoolName, id.VolumeName, id.BackupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Backup ID
func (id BackupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountValue"),
		resourceids.StaticSegment("staticCapacityPools", "capacityPools", "capacityPools"),
		resourceids.UserSpecifiedSegment("capacityPoolName", "capacityPoolValue"),
		resourceids.StaticSegment("staticVolumes", "volumes", "volumes"),
		resourceids.UserSpecifiedSegment("volumeName", "volumeValue"),
		resourceids.StaticSegment("staticBackups", "backups", "backups"),
		resourceids.UserSpecifiedSegment("backupName", "backupValue"),
	}
}

// String returns a human-readable description of this Backup ID
func (id BackupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Capacity Pool Name: %q", id.CapacityPoolName),
		fmt.Sprintf("Volume Name: %q", id.VolumeName),
		fmt.Sprintf("Backup Name: %q", id.BackupName),
	}
	return fmt.Sprintf("Backup (%s)", strings.Join(components, "\n"))
}
