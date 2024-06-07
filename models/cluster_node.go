// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterNode cluster node
//
// swagger:model ClusterNode
type ClusterNode struct {

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
	Region string `json:"region,omitempty"`

	// volume iops
	VolumeIops int64 `json:"volume_iops,omitempty"`

	// volume size
	VolumeSize int64 `json:"volume_size,omitempty"`

	// volume type
	VolumeType string `json:"volume_type,omitempty"`
}

// Validate validates this cluster node
func (m *ClusterNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster node based on context it is used
func (m *ClusterNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNode) UnmarshalBinary(b []byte) error {
	var res ClusterNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
