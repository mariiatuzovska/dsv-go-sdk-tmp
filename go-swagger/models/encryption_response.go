// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncryptionResponse EncryptionResponse contains ciphertext produced from encrypting a plaintext string with a key
//
// swagger:model EncryptionResponse
type EncryptionResponse struct {

	// ciphertext
	Ciphertext []uint8 `json:"ciphertext"`

	// Path of the key with which encryption was performed
	Path string `json:"path,omitempty"`

	// Version of the key with which encryption was performed
	Version string `json:"version,omitempty"`
}

// Validate validates this encryption response
func (m *EncryptionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this encryption response based on context it is used
func (m *EncryptionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionResponse) UnmarshalBinary(b []byte) error {
	var res EncryptionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}