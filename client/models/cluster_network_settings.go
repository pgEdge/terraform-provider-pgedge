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

// ClusterNetworkSettings cluster network settings
//
// swagger:model ClusterNetworkSettings
type ClusterNetworkSettings struct {

	// cidr
	Cidr string `json:"cidr,omitempty"`

	// external
	External bool `json:"external,omitempty"`

	// external id
	ExternalID string `json:"external_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// private subnets
	PrivateSubnets []string `json:"private_subnets"`

	// public subnets
	PublicSubnets []string `json:"public_subnets"`

	// region
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this cluster network settings
func (m *ClusterNetworkSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNetworkSettings) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster network settings based on context it is used
func (m *ClusterNetworkSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNetworkSettings) UnmarshalBinary(b []byte) error {
	var res ClusterNetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}