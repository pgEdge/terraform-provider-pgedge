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

// DatabaseNodeExtensions database node extensions
//
// swagger:model DatabaseNodeExtensions
type DatabaseNodeExtensions struct {

	// errors
	Errors map[string]string `json:"errors,omitempty"`

	// installed
	// Required: true
	Installed []string `json:"installed"`
}

// Validate validates this database node extensions
func (m *DatabaseNodeExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstalled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseNodeExtensions) validateInstalled(formats strfmt.Registry) error {

	if err := validate.Required("installed", "body", m.Installed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this database node extensions based on context it is used
func (m *DatabaseNodeExtensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseNodeExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseNodeExtensions) UnmarshalBinary(b []byte) error {
	var res DatabaseNodeExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}