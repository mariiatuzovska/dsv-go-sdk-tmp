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

// ClientCreate client create
//
// swagger:model ClientCreate
type ClientCreate struct {

	// Client description
	Description string `json:"description,omitempty"`

	// Name of role to assign to client
	// Required: true
	Role *string `json:"role"`

	// TTL expiration in seconds
	TTL int64 `json:"ttl,omitempty"`

	// Url
	URL bool `json:"url,omitempty"`

	// Url TTL
	URLTTL int64 `json:"urlTTL,omitempty"`

	// Uses the number of times the client credential can be read. if set to 0, it can be used infinitely. default is 0.
	UsesLimit int64 `json:"usesLimit,omitempty"`
}

// Validate validates this client create
func (m *ClientCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientCreate) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client create based on context it is used
func (m *ClientCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientCreate) UnmarshalBinary(b []byte) error {
	var res ClientCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
