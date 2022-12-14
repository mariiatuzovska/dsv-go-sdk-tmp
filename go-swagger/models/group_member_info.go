// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupMemberInfo group member info
//
// swagger:model GroupMemberInfo
type GroupMemberInfo struct {

	// Created at
	Created string `json:"created,omitempty"`

	// Who created
	CreatedBy string `json:"createdBy,omitempty"`

	// GroupName
	GroupName string `json:"groupName,omitempty"`
}

// Validate validates this group member info
func (m *GroupMemberInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group member info based on context it is used
func (m *GroupMemberInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupMemberInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupMemberInfo) UnmarshalBinary(b []byte) error {
	var res GroupMemberInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
