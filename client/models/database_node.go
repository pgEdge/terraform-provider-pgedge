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

// DatabaseNode database node
//
// swagger:model DatabaseNode
type DatabaseNode struct {

	// connection
	// Required: true
	Connection *Connection `json:"connection"`

	// distance measurement
	DistanceMeasurement *DistanceMeasurement `json:"distance_measurement,omitempty"`

	// extensions
	Extensions *DatabaseNodeExtensions `json:"extensions,omitempty"`

	// location
	// Required: true
	Location *Location `json:"location"`

	// name
	// Required: true
	Name *string `json:"name"`

	// region
	Region *Region `json:"region,omitempty"`
}

// Validate validates this database node
func (m *DatabaseNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistanceMeasurement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseNode) validateConnection(formats strfmt.Registry) error {

	if err := validate.Required("connection", "body", m.Connection); err != nil {
		return err
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseNode) validateDistanceMeasurement(formats strfmt.Registry) error {
	if swag.IsZero(m.DistanceMeasurement) { // not required
		return nil
	}

	if m.DistanceMeasurement != nil {
		if err := m.DistanceMeasurement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distance_measurement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distance_measurement")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseNode) validateExtensions(formats strfmt.Registry) error {
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

func (m *DatabaseNode) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseNode) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseNode) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this database node based on the context it is used
func (m *DatabaseNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDistanceMeasurement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseNode) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {

		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseNode) contextValidateDistanceMeasurement(ctx context.Context, formats strfmt.Registry) error {

	if m.DistanceMeasurement != nil {

		if swag.IsZero(m.DistanceMeasurement) { // not required
			return nil
		}

		if err := m.DistanceMeasurement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distance_measurement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distance_measurement")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseNode) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DatabaseNode) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseNode) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if swag.IsZero(m.Region) { // not required
			return nil
		}

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseNode) UnmarshalBinary(b []byte) error {
	var res DatabaseNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}