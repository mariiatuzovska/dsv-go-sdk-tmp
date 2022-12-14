// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Document Document is the per-tenant configuration store
//
// swagger:model Document
type Document struct {

	// created
	Created string `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// permission document
	PermissionDocument []*DefaultPolicy `json:"permissionDocument"`

	// refresh token TTL hours
	RefreshTokenTTLHours int64 `json:"refreshTokenTTLHours,omitempty"`

	// tenant name
	TenantName string `json:"tenantName,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// settings
	Settings *Settings `json:"settings,omitempty"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissionDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) validatePermissionDocument(formats strfmt.Registry) error {
	if swag.IsZero(m.PermissionDocument) { // not required
		return nil
	}

	for i := 0; i < len(m.PermissionDocument); i++ {
		if swag.IsZero(m.PermissionDocument[i]) { // not required
			continue
		}

		if m.PermissionDocument[i] != nil {
			if err := m.PermissionDocument[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissionDocument" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissionDocument" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this document based on the context it is used
func (m *Document) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissionDocument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) contextValidatePermissionDocument(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PermissionDocument); i++ {

		if m.PermissionDocument[i] != nil {
			if err := m.PermissionDocument[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissionDocument" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissionDocument" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Document) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {
		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
