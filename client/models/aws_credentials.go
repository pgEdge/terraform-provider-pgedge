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

// AwsCredentials aws credentials
//
// swagger:model AwsCredentials
type AwsCredentials struct {

	// role arn
	// Required: true
	RoleArn *string `json:"role_arn"`
}

// Validate validates this aws credentials
func (m *AwsCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleArn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCredentials) validateRoleArn(formats strfmt.Registry) error {

	if err := validate.Required("role_arn", "body", m.RoleArn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aws credentials based on context it is used
func (m *AwsCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCredentials) UnmarshalBinary(b []byte) error {
	var res AwsCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}