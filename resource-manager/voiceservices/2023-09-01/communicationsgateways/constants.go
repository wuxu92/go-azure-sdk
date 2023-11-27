package communicationsgateways

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiBridgeActivationState string

const (
	ApiBridgeActivationStateDisabled ApiBridgeActivationState = "disabled"
	ApiBridgeActivationStateEnabled  ApiBridgeActivationState = "enabled"
)

func PossibleValuesForApiBridgeActivationState() []string {
	return []string{
		string(ApiBridgeActivationStateDisabled),
		string(ApiBridgeActivationStateEnabled),
	}
}

func (s *ApiBridgeActivationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApiBridgeActivationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApiBridgeActivationState(input string) (*ApiBridgeActivationState, error) {
	vals := map[string]ApiBridgeActivationState{
		"disabled": ApiBridgeActivationStateDisabled,
		"enabled":  ApiBridgeActivationStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiBridgeActivationState(input)
	return &out, nil
}

type AutoGeneratedDomainNameLabelScope string

const (
	AutoGeneratedDomainNameLabelScopeNoReuse            AutoGeneratedDomainNameLabelScope = "NoReuse"
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse  AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	AutoGeneratedDomainNameLabelScopeTenantReuse        AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

func PossibleValuesForAutoGeneratedDomainNameLabelScope() []string {
	return []string{
		string(AutoGeneratedDomainNameLabelScopeNoReuse),
		string(AutoGeneratedDomainNameLabelScopeResourceGroupReuse),
		string(AutoGeneratedDomainNameLabelScopeSubscriptionReuse),
		string(AutoGeneratedDomainNameLabelScopeTenantReuse),
	}
}

func (s *AutoGeneratedDomainNameLabelScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoGeneratedDomainNameLabelScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoGeneratedDomainNameLabelScope(input string) (*AutoGeneratedDomainNameLabelScope, error) {
	vals := map[string]AutoGeneratedDomainNameLabelScope{
		"noreuse":            AutoGeneratedDomainNameLabelScopeNoReuse,
		"resourcegroupreuse": AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		"subscriptionreuse":  AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		"tenantreuse":        AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoGeneratedDomainNameLabelScope(input)
	return &out, nil
}

type CommunicationsPlatform string

const (
	CommunicationsPlatformOperatorConnect    CommunicationsPlatform = "OperatorConnect"
	CommunicationsPlatformTeamsDirectRouting CommunicationsPlatform = "TeamsDirectRouting"
	CommunicationsPlatformTeamsPhoneMobile   CommunicationsPlatform = "TeamsPhoneMobile"
)

func PossibleValuesForCommunicationsPlatform() []string {
	return []string{
		string(CommunicationsPlatformOperatorConnect),
		string(CommunicationsPlatformTeamsDirectRouting),
		string(CommunicationsPlatformTeamsPhoneMobile),
	}
}

func (s *CommunicationsPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommunicationsPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommunicationsPlatform(input string) (*CommunicationsPlatform, error) {
	vals := map[string]CommunicationsPlatform{
		"operatorconnect":    CommunicationsPlatformOperatorConnect,
		"teamsdirectrouting": CommunicationsPlatformTeamsDirectRouting,
		"teamsphonemobile":   CommunicationsPlatformTeamsPhoneMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommunicationsPlatform(input)
	return &out, nil
}

type Connectivity string

const (
	ConnectivityPublicAddress Connectivity = "PublicAddress"
)

func PossibleValuesForConnectivity() []string {
	return []string{
		string(ConnectivityPublicAddress),
	}
}

func (s *Connectivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectivity(input string) (*Connectivity, error) {
	vals := map[string]Connectivity{
		"publicaddress": ConnectivityPublicAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Connectivity(input)
	return &out, nil
}

type E911Type string

const (
	E911TypeDirectToEsrp E911Type = "DirectToEsrp"
	E911TypeStandard     E911Type = "Standard"
)

func PossibleValuesForE911Type() []string {
	return []string{
		string(E911TypeDirectToEsrp),
		string(E911TypeStandard),
	}
}

func (s *E911Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseE911Type(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseE911Type(input string) (*E911Type, error) {
	vals := map[string]E911Type{
		"directtoesrp": E911TypeDirectToEsrp,
		"standard":     E911TypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := E911Type(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func (s *SkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}

type Status string

const (
	StatusChangePending Status = "ChangePending"
	StatusComplete      Status = "Complete"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusChangePending),
		string(StatusComplete),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"changepending": StatusChangePending,
		"complete":      StatusComplete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}

type TeamsCodecs string

const (
	TeamsCodecsGSevenTwoTwo    TeamsCodecs = "G722"
	TeamsCodecsGSevenTwoTwoTwo TeamsCodecs = "G722_2"
	TeamsCodecsPCMA            TeamsCodecs = "PCMA"
	TeamsCodecsPCMU            TeamsCodecs = "PCMU"
	TeamsCodecsSILKEight       TeamsCodecs = "SILK_8"
	TeamsCodecsSILKOneSix      TeamsCodecs = "SILK_16"
)

func PossibleValuesForTeamsCodecs() []string {
	return []string{
		string(TeamsCodecsGSevenTwoTwo),
		string(TeamsCodecsGSevenTwoTwoTwo),
		string(TeamsCodecsPCMA),
		string(TeamsCodecsPCMU),
		string(TeamsCodecsSILKEight),
		string(TeamsCodecsSILKOneSix),
	}
}

func (s *TeamsCodecs) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsCodecs(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsCodecs(input string) (*TeamsCodecs, error) {
	vals := map[string]TeamsCodecs{
		"g722":    TeamsCodecsGSevenTwoTwo,
		"g722_2":  TeamsCodecsGSevenTwoTwoTwo,
		"pcma":    TeamsCodecsPCMA,
		"pcmu":    TeamsCodecsPCMU,
		"silk_8":  TeamsCodecsSILKEight,
		"silk_16": TeamsCodecsSILKOneSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsCodecs(input)
	return &out, nil
}
