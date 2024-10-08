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

// Region region
//
// swagger:model Region
type Region struct {

	// active
	Active bool `json:"active,omitempty"`

	// availability zones
	// Required: true
	AvailabilityZones []string `json:"availability_zones"`

	// cloud
	// Required: true
	Cloud *string `json:"cloud"`

	// code
	// Required: true
	Code *string `json:"code"`

	// name
	// Required: true
	Name *string `json:"name"`

	// parent
	Parent string `json:"parent,omitempty"`
}

// Validate validates this region
func (m *Region) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Region) validateAvailabilityZones(formats strfmt.Registry) error {

	if err := validate.Required("availability_zones", "body", m.AvailabilityZones); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateCloud(formats strfmt.Registry) error {

	if err := validate.Required("cloud", "body", m.Cloud); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Region) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this region based on context it is used
func (m *Region) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Region) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Region) UnmarshalBinary(b []byte) error {
	var res Region
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
