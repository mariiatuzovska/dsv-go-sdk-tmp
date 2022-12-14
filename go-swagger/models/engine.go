// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Engine Engine is the stored record
//
// swagger:model Engine
type Engine struct {

	// Created date
	Created string `json:"created,omitempty"`

	// Who created the item
	CreatedBy string `json:"createdBy,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// the id for this item
	ID string `json:"id,omitempty"`

	// last heartbeat
	LastHeartbeat string `json:"lastHeartbeat,omitempty"`

	// Last updated date
	LastModified string `json:"lastModified,omitempty"`

	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pool name
	PoolName string `json:"poolName,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// Current version
	Version string `json:"version,omitempty"`
}

// Validate validates this engine
func (m *Engine) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this engine based on context it is used
func (m *Engine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Engine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Engine) UnmarshalBinary(b []byte) error {
	var res Engine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
