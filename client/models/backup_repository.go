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

// BackupRepository backup repository
//
// swagger:model BackupRepository
type BackupRepository struct {

	// azure account
	AzureAccount string `json:"azure_account,omitempty"`

	// azure container
	AzureContainer string `json:"azure_container,omitempty"`

	// azure endpoint
	AzureEndpoint string `json:"azure_endpoint,omitempty"`

	// backup store id
	// Format: uuid
	BackupStoreID strfmt.UUID `json:"backup_store_id,omitempty"`

	// base path
	BasePath string `json:"base_path,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// database id
	// Format: uuid
	DatabaseID strfmt.UUID `json:"database_id,omitempty"`

	// gcs bucket
	GcsBucket string `json:"gcs_bucket,omitempty"`

	// gcs endpoint
	GcsEndpoint string `json:"gcs_endpoint,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// retention full
	RetentionFull int64 `json:"retention_full,omitempty"`

	// retention full type
	RetentionFullType string `json:"retention_full_type,omitempty"`

	// s3 bucket
	S3Bucket string `json:"s3_bucket,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3_endpoint,omitempty"`

	// s3 region
	S3Region string `json:"s3_region,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this backup repository
func (m *BackupRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupStoreID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRepository) validateBackupStoreID(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupStoreID) { // not required
		return nil
	}

	if err := validate.FormatOf("backup_store_id", "body", "uuid", m.BackupStoreID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupRepository) validateDatabaseID(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseID) { // not required
		return nil
	}

	if err := validate.FormatOf("database_id", "body", "uuid", m.DatabaseID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupRepository) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup repository based on context it is used
func (m *BackupRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRepository) UnmarshalBinary(b []byte) error {
	var res BackupRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}