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

// ClusterCreationResponse cluster creation response
//
// swagger:model ClusterCreationResponse
type ClusterCreationResponse struct {

	// aws
	Aws *ClusterCreationResponseAws `json:"aws,omitempty"`

	// cloud account id
	CloudAccountID string `json:"cloud_account_id,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// database
	Database *ClusterCreationResponseDatabase `json:"database,omitempty"`

	// firewall
	Firewall *ClusterCreationResponseFirewall `json:"firewall,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// node groups
	NodeGroups *ClusterCreationResponseNodeGroups `json:"node_groups,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster creation response
func (m *ClusterCreationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationResponse) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreationResponse) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreationResponse) validateDatabase(formats strfmt.Registry) error {
	if swag.IsZero(m.Database) { // not required
		return nil
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreationResponse) validateFirewall(formats strfmt.Registry) error {
	if swag.IsZero(m.Firewall) { // not required
		return nil
	}

	if m.Firewall != nil {
		if err := m.Firewall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firewall")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreationResponse) validateNodeGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeGroups) { // not required
		return nil
	}

	if m.NodeGroups != nil {
		if err := m.NodeGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_groups")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster creation response based on the context it is used
func (m *ClusterCreationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirewall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationResponse) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {

		if swag.IsZero(m.Aws) { // not required
			return nil
		}

		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreationResponse) contextValidateDatabase(ctx context.Context, formats strfmt.Registry) error {

	if m.Database != nil {

		if swag.IsZero(m.Database) { // not required
			return nil
		}

		if err := m.Database.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreationResponse) contextValidateFirewall(ctx context.Context, formats strfmt.Registry) error {

	if m.Firewall != nil {

		if swag.IsZero(m.Firewall) { // not required
			return nil
		}

		if err := m.Firewall.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firewall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firewall")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCreationResponse) contextValidateNodeGroups(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeGroups != nil {

		if swag.IsZero(m.NodeGroups) { // not required
			return nil
		}

		if err := m.NodeGroups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_groups")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationResponse) UnmarshalBinary(b []byte) error {
	var res ClusterCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterCreationResponseAws cluster creation response aws
//
// swagger:model ClusterCreationResponseAws
type ClusterCreationResponseAws struct {

	// role arn
	RoleArn string `json:"role_arn,omitempty"`
}

// Validate validates this cluster creation response aws
func (m *ClusterCreationResponseAws) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster creation response aws based on context it is used
func (m *ClusterCreationResponseAws) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationResponseAws) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationResponseAws) UnmarshalBinary(b []byte) error {
	var res ClusterCreationResponseAws
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterCreationResponseDatabase cluster creation response database
//
// swagger:model ClusterCreationResponseDatabase
type ClusterCreationResponseDatabase struct {

	// name
	Name string `json:"name,omitempty"`

	// pg version
	PgVersion string `json:"pg_version,omitempty"`

	// scripts
	Scripts interface{} `json:"scripts,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this cluster creation response database
func (m *ClusterCreationResponseDatabase) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster creation response database based on context it is used
func (m *ClusterCreationResponseDatabase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationResponseDatabase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationResponseDatabase) UnmarshalBinary(b []byte) error {
	var res ClusterCreationResponseDatabase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterCreationResponseFirewall cluster creation response firewall
//
// swagger:model ClusterCreationResponseFirewall
type ClusterCreationResponseFirewall struct {

	// rules
	Rules []*ClusterCreationResponseFirewallRulesItems0 `json:"rules"`
}

// Validate validates this cluster creation response firewall
func (m *ClusterCreationResponseFirewall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationResponseFirewall) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewall" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("firewall" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster creation response firewall based on the context it is used
func (m *ClusterCreationResponseFirewall) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationResponseFirewall) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {

			if swag.IsZero(m.Rules[i]) { // not required
				return nil
			}

			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewall" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("firewall" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationResponseFirewall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationResponseFirewall) UnmarshalBinary(b []byte) error {
	var res ClusterCreationResponseFirewall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterCreationResponseFirewallRulesItems0 cluster creation response firewall rules items0
//
// swagger:model ClusterCreationResponseFirewallRulesItems0
type ClusterCreationResponseFirewallRulesItems0 struct {

	// port
	Port int64 `json:"port,omitempty"`

	// sources
	Sources []string `json:"sources"`

	// type
	Type string `json:"type_,omitempty"`
}

// Validate validates this cluster creation response firewall rules items0
func (m *ClusterCreationResponseFirewallRulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster creation response firewall rules items0 based on context it is used
func (m *ClusterCreationResponseFirewallRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationResponseFirewallRulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationResponseFirewallRulesItems0) UnmarshalBinary(b []byte) error {
	var res ClusterCreationResponseFirewallRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterCreationResponseNodeGroups cluster creation response node groups
//
// swagger:model ClusterCreationResponseNodeGroups
type ClusterCreationResponseNodeGroups struct {

	// aws
	Aws []*NodeGroup `json:"aws"`

	// azure
	Azure []*NodeGroup `json:"azure"`

	// google
	Google []*NodeGroup `json:"google"`
}

// Validate validates this cluster creation response node groups
func (m *ClusterCreationResponseNodeGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationResponseNodeGroups) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	for i := 0; i < len(m.Aws); i++ {
		if swag.IsZero(m.Aws[i]) { // not required
			continue
		}

		if m.Aws[i] != nil {
			if err := m.Aws[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_groups" + "." + "aws" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_groups" + "." + "aws" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreationResponseNodeGroups) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	for i := 0; i < len(m.Azure); i++ {
		if swag.IsZero(m.Azure[i]) { // not required
			continue
		}

		if m.Azure[i] != nil {
			if err := m.Azure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_groups" + "." + "azure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_groups" + "." + "azure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreationResponseNodeGroups) validateGoogle(formats strfmt.Registry) error {
	if swag.IsZero(m.Google) { // not required
		return nil
	}

	for i := 0; i < len(m.Google); i++ {
		if swag.IsZero(m.Google[i]) { // not required
			continue
		}

		if m.Google[i] != nil {
			if err := m.Google[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_groups" + "." + "google" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_groups" + "." + "google" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster creation response node groups based on the context it is used
func (m *ClusterCreationResponseNodeGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationResponseNodeGroups) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aws); i++ {

		if m.Aws[i] != nil {

			if swag.IsZero(m.Aws[i]) { // not required
				return nil
			}

			if err := m.Aws[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_groups" + "." + "aws" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_groups" + "." + "aws" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreationResponseNodeGroups) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Azure); i++ {

		if m.Azure[i] != nil {

			if swag.IsZero(m.Azure[i]) { // not required
				return nil
			}

			if err := m.Azure[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_groups" + "." + "azure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_groups" + "." + "azure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterCreationResponseNodeGroups) contextValidateGoogle(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Google); i++ {

		if m.Google[i] != nil {

			if swag.IsZero(m.Google[i]) { // not required
				return nil
			}

			if err := m.Google[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_groups" + "." + "google" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_groups" + "." + "google" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationResponseNodeGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationResponseNodeGroups) UnmarshalBinary(b []byte) error {
	var res ClusterCreationResponseNodeGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
