// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoleCreate role create
//
// swagger:model RoleCreate
type RoleCreate struct {

	// Role description
	Description string `json:"description,omitempty"`

	// External identifier, such as an AWS arn for 3rd party authentication
	ExternalID string `json:"externalId,omitempty"`

	// Name of role
	// Required: true
	Name *string `json:"name"`

	// Provider name defined in the authentication settings section of configuration
	Provider string `json:"provider,omitempty"`
}

// Validate validates this role create
func (m *RoleCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role create based on context it is used
func (m *RoleCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleCreate) UnmarshalBinary(b []byte) error {
	var res RoleCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}