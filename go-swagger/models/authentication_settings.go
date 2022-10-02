// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthenticationSettings AuthenticationSettings is the 3rd party authentication providers.
//
// swagger:model AuthenticationSettings
type AuthenticationSettings struct {

	// AuthenticationSettings provider system type.
	AuthType string `json:"type,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// properties
	Properties *ProviderProperties `json:"properties,omitempty"`
}

// Validate validates this authentication settings
func (m *AuthenticationSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationSettings) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this authentication settings based on the context it is used
func (m *AuthenticationSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationSettings) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	if m.Properties != nil {
		if err := m.Properties.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationSettings) UnmarshalBinary(b []byte) error {
	var res AuthenticationSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}