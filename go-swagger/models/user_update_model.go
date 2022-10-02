// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserUpdateModel Editable properties of a user
//
// swagger:model UserUpdateModel
type UserUpdateModel struct {

	// The display name of the user
	DisplayName string `json:"displayName,omitempty"`

	// User's password (not required if using 3rd party auth)
	Password string `json:"password,omitempty"`
}

// Validate validates this user update model
func (m *UserUpdateModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user update model based on context it is used
func (m *UserUpdateModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdateModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdateModel) UnmarshalBinary(b []byte) error {
	var res UserUpdateModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}