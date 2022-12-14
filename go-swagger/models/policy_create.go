// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyCreate PolicyCreate struct
//
// swagger:model PolicyCreate
type PolicyCreate struct {

	// path
	// Required: true
	// Max Length: 500
	// Min Length: 3
	Path *string `json:"path"`

	// policy
	// Required: true
	// Max Length: 2000
	// Min Length: 10
	Policy *string `json:"policy"`

	// serialization
	// Enum: [json yaml yml]
	Serialization *string `json:"serialization,omitempty"`
}

// Validate validates this policy create
func (m *PolicyCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyCreate) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	if err := validate.MinLength("path", "body", *m.Path, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("path", "body", *m.Path, 500); err != nil {
		return err
	}

	return nil
}

func (m *PolicyCreate) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	if err := validate.MinLength("policy", "body", *m.Policy, 10); err != nil {
		return err
	}

	if err := validate.MaxLength("policy", "body", *m.Policy, 2000); err != nil {
		return err
	}

	return nil
}

var policyCreateTypeSerializationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["json","yaml","yml"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyCreateTypeSerializationPropEnum = append(policyCreateTypeSerializationPropEnum, v)
	}
}

const (

	// PolicyCreateSerializationJSON captures enum value "json"
	PolicyCreateSerializationJSON string = "json"

	// PolicyCreateSerializationYaml captures enum value "yaml"
	PolicyCreateSerializationYaml string = "yaml"

	// PolicyCreateSerializationYml captures enum value "yml"
	PolicyCreateSerializationYml string = "yml"
)

// prop value enum
func (m *PolicyCreate) validateSerializationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyCreateTypeSerializationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyCreate) validateSerialization(formats strfmt.Registry) error {
	if swag.IsZero(m.Serialization) { // not required
		return nil
	}

	// value enum
	if err := m.validateSerializationEnum("serialization", "body", *m.Serialization); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this policy create based on context it is used
func (m *PolicyCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyCreate) UnmarshalBinary(b []byte) error {
	var res PolicyCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
