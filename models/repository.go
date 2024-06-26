// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Repository repository
//
// swagger:model Repository
type Repository struct {

	// base path
	BasePath string `json:"base_path,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// retention full
	RetentionFull int64 `json:"retention_full,omitempty"`

	// retention full type
	RetentionFullType string `json:"retention_full_type,omitempty"`

	// s3 bucket
	S3Bucket string `json:"s3_bucket,omitempty"`

	// s3 region
	S3Region string `json:"s3_region,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this repository
func (m *Repository) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this repository based on context it is used
func (m *Repository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Repository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Repository) UnmarshalBinary(b []byte) error {
	var res Repository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
