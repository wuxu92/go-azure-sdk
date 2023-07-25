package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainOperationPredicate struct {
	Id       *string
	Kind     *string
	Location *string
	Name     *string
	Type     *string
}

func (p DomainOperationPredicate) Matches(input Domain) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type DomainOwnershipIdentifierOperationPredicate struct {
	Id   *string
	Kind *string
	Name *string
	Type *string
}

func (p DomainOwnershipIdentifierOperationPredicate) Matches(input DomainOwnershipIdentifier) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Kind != nil && (input.Kind == nil && *p.Kind != *input.Kind) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type NameIdentifierOperationPredicate struct {
	Name *string
}

func (p NameIdentifierOperationPredicate) Matches(input NameIdentifier) bool {

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
