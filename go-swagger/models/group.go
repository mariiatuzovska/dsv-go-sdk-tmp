// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Group group
//
// swagger:model Group
type Group struct {

	// the id for this item
	ID string `json:"id,omitempty"`

	// Members
	Members []string `json:"members"`

	// MetaData
	MetaData []map[string]string `json:"metaData"`

	// Name
	Name string `json:"groupName,omitempty"`

	// Total number of members
	Total int64 `json:"total,omitempty"`
}

// Validate validates this group
func (m *Group) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group based on context it is used
func (m *Group) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Group) UnmarshalBinary(b []byte) error {
	var res Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
