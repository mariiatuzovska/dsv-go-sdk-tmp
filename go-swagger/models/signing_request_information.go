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

// SigningRequestInformation signing request information
//
// swagger:model SigningRequestInformation
type SigningRequestInformation struct {

	// Boolean indicating whether to return a root certificate
	Chain bool `json:"chain,omitempty"`

	// common name
	// Required: true
	CommonName *string `json:"commonName"`

	// country
	Country string `json:"country,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// email address
	EmailAddress string `json:"emailAddress,omitempty"`

	// locality
	Locality string `json:"locality,omitempty"`

	// organization
	Organization string `json:"organization,omitempty"`

	// organizational unit
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`

	// The name of the secret containing the root CA certificate
	// Required: true
	RootCAPath *string `json:"rootCAPath"`

	// state
	State string `json:"state,omitempty"`

	// The name of the secret in which to store the generated certificate and private key
	StorePath string `json:"storePath,omitempty"`

	// TTL for a generated certificate (in hours, cannot exceed the maximum TTL specified in root CA secret)
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this signing request information
func (m *SigningRequestInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommonName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootCAPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigningRequestInformation) validateCommonName(formats strfmt.Registry) error {

	if err := validate.Required("commonName", "body", m.CommonName); err != nil {
		return err
	}

	return nil
}

func (m *SigningRequestInformation) validateRootCAPath(formats strfmt.Registry) error {

	if err := validate.Required("rootCAPath", "body", m.RootCAPath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signing request information based on context it is used
func (m *SigningRequestInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SigningRequestInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigningRequestInformation) UnmarshalBinary(b []byte) error {
	var res SigningRequestInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
