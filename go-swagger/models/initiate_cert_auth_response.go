// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InitiateCertAuthResponse InitiateCertAuthResponse contains challenge to decrypt and challenge id
//
// swagger:model InitiateCertAuthResponse
type InitiateCertAuthResponse struct {

	// Encrypted and base64 encoded challenge
	Encrypted string `json:"encrypted,omitempty"`

	// Challenge id
	ID string `json:"cert_challenge_id,omitempty"`
}

// Validate validates this initiate cert auth response
func (m *InitiateCertAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this initiate cert auth response based on context it is used
func (m *InitiateCertAuthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitiateCertAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitiateCertAuthResponse) UnmarshalBinary(b []byte) error {
	var res InitiateCertAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
