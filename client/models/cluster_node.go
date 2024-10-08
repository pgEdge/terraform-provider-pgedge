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

// ClusterNode cluster node
//
// swagger:model ClusterNode
type ClusterNode struct {

	// availability zone
	// Required: true
	AvailabilityZone *string `json:"availability_zone"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// distance measurement
	DistanceMeasurement *DistanceMeasurement `json:"distance_measurement,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// image id
	// Required: true
	ImageID *string `json:"image_id"`

	// instance id
	// Required: true
	InstanceID *string `json:"instance_id"`

	// instance type
	// Required: true
	InstanceType *string `json:"instance_type"`

	// ip address
	// Required: true
	IPAddress *string `json:"ip_address"`

	// is active
	// Required: true
	IsActive *bool `json:"is_active"`

	// key name
	// Required: true
	KeyName *string `json:"key_name"`

	// location
	// Required: true
	Location *Location `json:"location"`

	// name
	// Required: true
	Name *string `json:"name"`

	// pg version
	// Required: true
	PgVersion *string `json:"pg_version"`

	// public ip address
	PublicIPAddress string `json:"public_ip_address,omitempty"`

	// region
	// Required: true
	Region *string `json:"region"`

	// region detail
	// Required: true
	RegionDetail *Region `json:"region_detail"`

	// volume iops
	// Required: true
	VolumeIops *int64 `json:"volume_iops"`

	// volume size
	// Required: true
	VolumeSize *int64 `json:"volume_size"`

	// volume type
	// Required: true
	VolumeType *string `json:"volume_type"`
}

// Validate validates this cluster node
func (m *ClusterNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistanceMeasurement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePgVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNode) validateAvailabilityZone(formats strfmt.Registry) error {

	if err := validate.Required("availability_zone", "body", m.AvailabilityZone); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateDistanceMeasurement(formats strfmt.Registry) error {
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

func (m *ClusterNode) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("image_id", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instance_id", "body", m.InstanceID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instance_type", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ip_address", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("is_active", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateKeyName(formats strfmt.Registry) error {

	if err := validate.Required("key_name", "body", m.KeyName); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateLocation(formats strfmt.Registry) error {

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

func (m *ClusterNode) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validatePgVersion(formats strfmt.Registry) error {

	if err := validate.Required("pg_version", "body", m.PgVersion); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateRegionDetail(formats strfmt.Registry) error {

	if err := validate.Required("region_detail", "body", m.RegionDetail); err != nil {
		return err
	}

	if m.RegionDetail != nil {
		if err := m.RegionDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region_detail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region_detail")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNode) validateVolumeIops(formats strfmt.Registry) error {

	if err := validate.Required("volume_iops", "body", m.VolumeIops); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateVolumeSize(formats strfmt.Registry) error {

	if err := validate.Required("volume_size", "body", m.VolumeSize); err != nil {
		return err
	}

	return nil
}

func (m *ClusterNode) validateVolumeType(formats strfmt.Registry) error {

	if err := validate.Required("volume_type", "body", m.VolumeType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster node based on the context it is used
func (m *ClusterNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDistanceMeasurement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegionDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNode) contextValidateDistanceMeasurement(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterNode) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterNode) contextValidateRegionDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.RegionDetail != nil {

		if err := m.RegionDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region_detail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region_detail")
			}
			return err
		}
	}

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
