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

// ClusterNodeSettings cluster node settings
//
// swagger:model ClusterNodeSettings
type ClusterNodeSettings struct {

	// availability zone
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// instance type
	InstanceType string `json:"instance_type,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// options
	Options []string `json:"options"`

	// region
	// Required: true
	Region *string `json:"region"`

	// volume iops
	VolumeIops int64 `json:"volume_iops,omitempty"`

	// volume size
	VolumeSize int64 `json:"volume_size,omitempty"`

	// volume type
	VolumeType string `json:"volume_type,omitempty"`
}

// Validate validates this cluster node settings
func (m *ClusterNodeSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNodeSettings) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster node settings based on context it is used
func (m *ClusterNodeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNodeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNodeSettings) UnmarshalBinary(b []byte) error {
	var res ClusterNodeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}