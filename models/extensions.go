// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Extensions extensions
//
// swagger:model Extensions
type Extensions struct {

	// auto manage
	AutoManage bool `json:"auto_manage,omitempty"`

	// available
	Available []string `json:"available"`

	// requested
	Requested []string `json:"requested"`
}

// Validate validates this extensions
func (m *Extensions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this extensions based on context it is used
func (m *Extensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Extensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Extensions) UnmarshalBinary(b []byte) error {
	var res Extensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
