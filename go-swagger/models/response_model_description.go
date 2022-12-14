// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseModelDescription ResponseModelDescription contains metadata but not the sensitive secret data
//
// swagger:model ResponseModelDescription
type ResponseModelDescription struct {

	// The user defined metadata
	Attributes interface{} `json:"attributes,omitempty"`

	// Created date
	Created string `json:"created,omitempty"`

	// Who created
	CreatedBy string `json:"createdBy,omitempty"`

	// Description of secret
	Description string `json:"description,omitempty"`

	// The unique id for this item
	ID string `json:"id,omitempty"`

	// Last updated date
	LastModified string `json:"lastModified,omitempty"`

	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// The path the secret is located at
	Path string `json:"path,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this response model description
func (m *ResponseModelDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response model description based on context it is used
func (m *ResponseModelDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseModelDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseModelDescription) UnmarshalBinary(b []byte) error {
	var res ResponseModelDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
