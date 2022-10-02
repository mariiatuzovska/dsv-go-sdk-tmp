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

// UpdateKeyRequest update key request
//
// swagger:model UpdateKeyRequest
type UpdateKeyRequest struct {

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// Base64 encoded nonce to be used with key. If not provided, DSV generates it for the user.
	Nonce string `json:"nonce,omitempty"`

	// Base64 encoded private key
	// Required: true
	PrivateKey *string `json:"privateKey"`
}

// Validate validates this update key request
func (m *UpdateKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateKeyRequest) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update key request based on context it is used
func (m *UpdateKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateKeyRequest) UnmarshalBinary(b []byte) error {
	var res UpdateKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}