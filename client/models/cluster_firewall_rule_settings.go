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

// ClusterFirewallRuleSettings cluster firewall rule settings
//
// swagger:model ClusterFirewallRuleSettings
type ClusterFirewallRuleSettings struct {

	// name
	Name string `json:"name,omitempty"`

	// port
	// Required: true
	Port *int64 `json:"port"`

	// sources
	// Required: true
	Sources []string `json:"sources"`
}

// Validate validates this cluster firewall rule settings
func (m *ClusterFirewallRuleSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterFirewallRuleSettings) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *ClusterFirewallRuleSettings) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster firewall rule settings based on context it is used
func (m *ClusterFirewallRuleSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterFirewallRuleSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterFirewallRuleSettings) UnmarshalBinary(b []byte) error {
	var res ClusterFirewallRuleSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
