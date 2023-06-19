package backupprotectableitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadProtectableItem = AzureVMWorkloadSAPHanaHSRProtectableItem{}

type AzureVMWorkloadSAPHanaHSRProtectableItem struct {
	IsAutoProtectable       *bool                `json:"isAutoProtectable,omitempty"`
	IsAutoProtected         *bool                `json:"isAutoProtected,omitempty"`
	IsProtectable           *bool                `json:"isProtectable,omitempty"`
	ParentName              *string              `json:"parentName,omitempty"`
	ParentUniqueName        *string              `json:"parentUniqueName,omitempty"`
	Prebackupvalidation     *PreBackupValidation `json:"prebackupvalidation,omitempty"`
	ServerName              *string              `json:"serverName,omitempty"`
	Subinquireditemcount    *int64               `json:"subinquireditemcount,omitempty"`
	Subprotectableitemcount *int64               `json:"subprotectableitemcount,omitempty"`

	// Fields inherited from WorkloadProtectableItem
	BackupManagementType *string           `json:"backupManagementType,omitempty"`
	FriendlyName         *string           `json:"friendlyName,omitempty"`
	ProtectionState      *ProtectionStatus `json:"protectionState,omitempty"`
	WorkloadType         *string           `json:"workloadType,omitempty"`
}

var _ json.Marshaler = AzureVMWorkloadSAPHanaHSRProtectableItem{}

func (s AzureVMWorkloadSAPHanaHSRProtectableItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMWorkloadSAPHanaHSRProtectableItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMWorkloadSAPHanaHSRProtectableItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMWorkloadSAPHanaHSRProtectableItem: %+v", err)
	}
	decoded["protectableItemType"] = "HanaHSRContainer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMWorkloadSAPHanaHSRProtectableItem: %+v", err)
	}

	return encoded, nil
}
