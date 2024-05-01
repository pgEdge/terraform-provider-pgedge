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

// DatabaseCreationRequest database creation request
//
// swagger:model DatabaseCreationRequest
type DatabaseCreationRequest struct {

	// backups
	Backups *DatabaseCreationRequestBackups `json:"backups,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// config version
	ConfigVersion string `json:"config_version,omitempty"`

	// extensions
	Extensions *DatabaseCreationRequestExtensions `json:"extensions,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// options
	Options []string `json:"options"`
}

// Validate validates this database creation request
func (m *DatabaseCreationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseCreationRequest) validateBackups(formats strfmt.Registry) error {
	if swag.IsZero(m.Backups) { // not required
		return nil
	}

	if m.Backups != nil {
		if err := m.Backups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backups")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseCreationRequest) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	if m.Extensions != nil {
		if err := m.Extensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this database creation request based on the context it is used
func (m *DatabaseCreationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseCreationRequest) contextValidateBackups(ctx context.Context, formats strfmt.Registry) error {

	if m.Backups != nil {

		if swag.IsZero(m.Backups) { // not required
			return nil
		}

		if err := m.Backups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backups")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseCreationRequest) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Extensions != nil {

		if swag.IsZero(m.Extensions) { // not required
			return nil
		}

		if err := m.Extensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseCreationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseCreationRequest) UnmarshalBinary(b []byte) error {
	var res DatabaseCreationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseCreationRequestBackups database creation request backups
//
// swagger:model DatabaseCreationRequestBackups
type DatabaseCreationRequestBackups struct {

	// config
	Config []*BackupConfig `json:"config"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this database creation request backups
func (m *DatabaseCreationRequestBackups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseCreationRequestBackups) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	for i := 0; i < len(m.Config); i++ {
		if swag.IsZero(m.Config[i]) { // not required
			continue
		}

		if m.Config[i] != nil {
			if err := m.Config[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database creation request backups based on the context it is used
func (m *DatabaseCreationRequestBackups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseCreationRequestBackups) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Config); i++ {

		if m.Config[i] != nil {

			if swag.IsZero(m.Config[i]) { // not required
				return nil
			}

			if err := m.Config[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseCreationRequestBackups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseCreationRequestBackups) UnmarshalBinary(b []byte) error {
	var res DatabaseCreationRequestBackups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseCreationRequestExtensions database creation request extensions
//
// swagger:model DatabaseCreationRequestExtensions
type DatabaseCreationRequestExtensions struct {

	// auto manage
	AutoManage bool `json:"auto_manage,omitempty"`

	// available
	Available []string `json:"available"`

	// requested
	Requested []string `json:"requested"`
}

// Validate validates this database creation request extensions
func (m *DatabaseCreationRequestExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database creation request extensions based on context it is used
func (m *DatabaseCreationRequestExtensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseCreationRequestExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseCreationRequestExtensions) UnmarshalBinary(b []byte) error {
	var res DatabaseCreationRequestExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
