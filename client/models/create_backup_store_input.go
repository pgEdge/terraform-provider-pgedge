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

// CreateBackupStoreInput create backup store input
//
// swagger:model CreateBackupStoreInput
type CreateBackupStoreInput struct {

	// The ID of the cloud account to use for the backup store.
	// Required: true
	CloudAccountID *string `json:"cloud_account_id"`

	// The name of the backup store.
	// Required: true
	Name *string `json:"name"`

	// The region to use for the backup store.
	Region string `json:"region,omitempty"`
}

// Validate validates this create backup store input
func (m *CreateBackupStoreInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudAccountID(formats); err != nil {
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

func (m *CreateBackupStoreInput) validateCloudAccountID(formats strfmt.Registry) error {

	if err := validate.Required("cloud_account_id", "body", m.CloudAccountID); err != nil {
		return err
	}

	return nil
}

func (m *CreateBackupStoreInput) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create backup store input based on context it is used
func (m *CreateBackupStoreInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBackupStoreInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBackupStoreInput) UnmarshalBinary(b []byte) error {
	var res CreateBackupStoreInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}