// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MemberRequest MemberRequest to add users to or delete from a group
//
// swagger:model MemberRequest
type MemberRequest struct {

	// Names
	Names []string `json:"memberNames"`
}

// Validate validates this member request
func (m *MemberRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this member request based on context it is used
func (m *MemberRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MemberRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberRequest) UnmarshalBinary(b []byte) error {
	var res MemberRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}