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
	"github.com/go-openapi/validate"
)

// DatabaseTable database table
//
// swagger:model DatabaseTable
type DatabaseTable struct {

	// columns
	Columns []*DatabaseColumn `json:"columns"`

	// name
	// Required: true
	Name *string `json:"name"`

	// primary key
	// Required: true
	PrimaryKey []string `json:"primary_key"`

	// replication sets
	// Required: true
	ReplicationSets []string `json:"replication_sets"`

	// schema
	// Required: true
	Schema *string `json:"schema"`

	// status
	// Required: true
	Status []*DatabaseTableStatus `json:"status"`
}

// Validate validates this database table
func (m *DatabaseTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseTable) validateColumns(formats strfmt.Registry) error {
	if swag.IsZero(m.Columns) { // not required
		return nil
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatabaseTable) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseTable) validatePrimaryKey(formats strfmt.Registry) error {

	if err := validate.Required("primary_key", "body", m.PrimaryKey); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseTable) validateReplicationSets(formats strfmt.Registry) error {

	if err := validate.Required("replication_sets", "body", m.ReplicationSets); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseTable) validateSchema(formats strfmt.Registry) error {

	if err := validate.Required("schema", "body", m.Schema); err != nil {
		return err
	}

	return nil
}

func (m *DatabaseTable) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	for i := 0; i < len(m.Status); i++ {
		if swag.IsZero(m.Status[i]) { // not required
			continue
		}

		if m.Status[i] != nil {
			if err := m.Status[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database table based on the context it is used
func (m *DatabaseTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseTable) contextValidateColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Columns); i++ {

		if m.Columns[i] != nil {

			if swag.IsZero(m.Columns[i]) { // not required
				return nil
			}

			if err := m.Columns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatabaseTable) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Status); i++ {

		if m.Status[i] != nil {

			if swag.IsZero(m.Status[i]) { // not required
				return nil
			}

			if err := m.Status[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseTable) UnmarshalBinary(b []byte) error {
	var res DatabaseTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}