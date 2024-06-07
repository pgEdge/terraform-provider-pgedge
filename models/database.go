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

// Database database
//
// swagger:model Database
type Database struct {

	// backups
	Backups *DatabaseBackups `json:"backups,omitempty"`

	// cluster id
	// Format: uuid
	ClusterID strfmt.UUID `json:"cluster_id,omitempty"`

	// components
	Components []*DatabaseComponentsItems0 `json:"components"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// extensions
	Extensions *DatabaseExtensions `json:"extensions,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nodes
	Nodes []*Node `json:"nodes"`

	// options
	Options []string `json:"options"`

	// pg version
	PgVersion string `json:"pg_version,omitempty"`

	// roles
	Roles []*DatabaseRolesItems0 `json:"roles"`

	// status
	Status string `json:"status,omitempty"`

	// storage used
	StorageUsed int64 `json:"storage_used,omitempty"`

	// tables
	Tables []*DatabaseTablesItems0 `json:"tables"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this database
func (m *Database) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Database) validateBackups(formats strfmt.Registry) error {
	if swag.IsZero(m.Backups) { // not required
		return nil
	}

	if m.Backups != nil {
		if err := m.Backups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backups")
			}
			return err
		}
	}

	return nil
}

func (m *Database) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Database) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Database) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	if m.Extensions != nil {
		if err := m.Extensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensions")
			}
			return err
		}
	}

	return nil
}

func (m *Database) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Database) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Database) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Database) validateTables(formats strfmt.Registry) error {
	if swag.IsZero(m.Tables) { // not required
		return nil
	}

	for i := 0; i < len(m.Tables); i++ {
		if swag.IsZero(m.Tables[i]) { // not required
			continue
		}

		if m.Tables[i] != nil {
			if err := m.Tables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database based on the context it is used
func (m *Database) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Database) contextValidateBackups(ctx context.Context, formats strfmt.Registry) error {

	if m.Backups != nil {

		if swag.IsZero(m.Backups) { // not required
			return nil
		}

		if err := m.Backups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backups")
			}
			return err
		}
	}

	return nil
}

func (m *Database) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Components); i++ {

		if m.Components[i] != nil {

			if swag.IsZero(m.Components[i]) { // not required
				return nil
			}

			if err := m.Components[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Database) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Extensions != nil {

		if swag.IsZero(m.Extensions) { // not required
			return nil
		}

		if err := m.Extensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensions")
			}
			return err
		}
	}

	return nil
}

func (m *Database) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {

			if swag.IsZero(m.Nodes[i]) { // not required
				return nil
			}

			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Database) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {

			if swag.IsZero(m.Roles[i]) { // not required
				return nil
			}

			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Database) contextValidateTables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tables); i++ {

		if m.Tables[i] != nil {

			if swag.IsZero(m.Tables[i]) { // not required
				return nil
			}

			if err := m.Tables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Database) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Database) UnmarshalBinary(b []byte) error {
	var res Database
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseBackups database backups
//
// swagger:model DatabaseBackups
type DatabaseBackups struct {

	// config
	Config []*DatabaseBackupsConfigItems0 `json:"config"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this database backups
func (m *DatabaseBackups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseBackups) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	for i := 0; i < len(m.Config); i++ {
		if swag.IsZero(m.Config[i]) { // not required
			continue
		}

		if m.Config[i] != nil {
			if err := m.Config[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database backups based on the context it is used
func (m *DatabaseBackups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseBackups) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Config); i++ {

		if m.Config[i] != nil {

			if swag.IsZero(m.Config[i]) { // not required
				return nil
			}

			if err := m.Config[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backups" + "." + "config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseBackups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseBackups) UnmarshalBinary(b []byte) error {
	var res DatabaseBackups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseBackupsConfigItems0 database backups config items0
//
// swagger:model DatabaseBackupsConfigItems0
type DatabaseBackupsConfigItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// node name
	NodeName string `json:"node_name,omitempty"`

	// repositories
	Repositories []*DatabaseBackupsConfigItems0RepositoriesItems0 `json:"repositories"`

	// schedules
	Schedules []*DatabaseBackupsConfigItems0SchedulesItems0 `json:"schedules"`
}

// Validate validates this database backups config items0
func (m *DatabaseBackupsConfigItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepositories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseBackupsConfigItems0) validateRepositories(formats strfmt.Registry) error {
	if swag.IsZero(m.Repositories) { // not required
		return nil
	}

	for i := 0; i < len(m.Repositories); i++ {
		if swag.IsZero(m.Repositories[i]) { // not required
			continue
		}

		if m.Repositories[i] != nil {
			if err := m.Repositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatabaseBackupsConfigItems0) validateSchedules(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedules) { // not required
		return nil
	}

	for i := 0; i < len(m.Schedules); i++ {
		if swag.IsZero(m.Schedules[i]) { // not required
			continue
		}

		if m.Schedules[i] != nil {
			if err := m.Schedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database backups config items0 based on the context it is used
func (m *DatabaseBackupsConfigItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepositories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseBackupsConfigItems0) contextValidateRepositories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Repositories); i++ {

		if m.Repositories[i] != nil {

			if swag.IsZero(m.Repositories[i]) { // not required
				return nil
			}

			if err := m.Repositories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatabaseBackupsConfigItems0) contextValidateSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Schedules); i++ {

		if m.Schedules[i] != nil {

			if swag.IsZero(m.Schedules[i]) { // not required
				return nil
			}

			if err := m.Schedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseBackupsConfigItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseBackupsConfigItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseBackupsConfigItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseBackupsConfigItems0RepositoriesItems0 database backups config items0 repositories items0
//
// swagger:model DatabaseBackupsConfigItems0RepositoriesItems0
type DatabaseBackupsConfigItems0RepositoriesItems0 struct {

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

// Validate validates this database backups config items0 repositories items0
func (m *DatabaseBackupsConfigItems0RepositoriesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database backups config items0 repositories items0 based on context it is used
func (m *DatabaseBackupsConfigItems0RepositoriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseBackupsConfigItems0RepositoriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseBackupsConfigItems0RepositoriesItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseBackupsConfigItems0RepositoriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseBackupsConfigItems0SchedulesItems0 database backups config items0 schedules items0
//
// swagger:model DatabaseBackupsConfigItems0SchedulesItems0
type DatabaseBackupsConfigItems0SchedulesItems0 struct {

	// cron expression
	CronExpression string `json:"cron_expression,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this database backups config items0 schedules items0
func (m *DatabaseBackupsConfigItems0SchedulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database backups config items0 schedules items0 based on context it is used
func (m *DatabaseBackupsConfigItems0SchedulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseBackupsConfigItems0SchedulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseBackupsConfigItems0SchedulesItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseBackupsConfigItems0SchedulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseComponentsItems0 database components items0
//
// swagger:model DatabaseComponentsItems0
type DatabaseComponentsItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// release date
	ReleaseDate string `json:"release_date,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this database components items0
func (m *DatabaseComponentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database components items0 based on context it is used
func (m *DatabaseComponentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseComponentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseComponentsItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseComponentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseExtensions database extensions
//
// swagger:model DatabaseExtensions
type DatabaseExtensions struct {

	// auto manage
	AutoManage bool `json:"auto_manage,omitempty"`

	// available
	Available []string `json:"available"`

	// requested
	Requested []string `json:"requested"`
}

// Validate validates this database extensions
func (m *DatabaseExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database extensions based on context it is used
func (m *DatabaseExtensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseExtensions) UnmarshalBinary(b []byte) error {
	var res DatabaseExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseRolesItems0 database roles items0
//
// swagger:model DatabaseRolesItems0
type DatabaseRolesItems0 struct {

	// bypass rls
	BypassRls bool `json:"bypass_rls,omitempty"`

	// connection limit
	ConnectionLimit int64 `json:"connection_limit,omitempty"`

	// create db
	CreateDb bool `json:"create_db,omitempty"`

	// create role
	CreateRole bool `json:"create_role,omitempty"`

	// inherit
	Inherit bool `json:"inherit,omitempty"`

	// login
	Login bool `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// replication
	Replication bool `json:"replication,omitempty"`

	// superuser
	Superuser bool `json:"superuser,omitempty"`
}

// Validate validates this database roles items0
func (m *DatabaseRolesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database roles items0 based on context it is used
func (m *DatabaseRolesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseRolesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseRolesItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseRolesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseTablesItems0 database tables items0
//
// swagger:model DatabaseTablesItems0
type DatabaseTablesItems0 struct {

	// columns
	Columns []*DatabaseTablesItems0ColumnsItems0 `json:"columns"`

	// name
	Name string `json:"name,omitempty"`

	// primary key
	PrimaryKey []string `json:"primary_key"`

	// replication sets
	ReplicationSets []string `json:"replication_sets"`

	// schema
	Schema string `json:"schema,omitempty"`

	// status
	Status []*DatabaseTablesItems0StatusItems0 `json:"status"`
}

// Validate validates this database tables items0
func (m *DatabaseTablesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseTablesItems0) validateColumns(formats strfmt.Registry) error {
	if swag.IsZero(m.Columns) { // not required
		return nil
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatabaseTablesItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	for i := 0; i < len(m.Status); i++ {
		if swag.IsZero(m.Status[i]) { // not required
			continue
		}

		if m.Status[i] != nil {
			if err := m.Status[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this database tables items0 based on the context it is used
func (m *DatabaseTablesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseTablesItems0) contextValidateColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Columns); i++ {

		if m.Columns[i] != nil {

			if swag.IsZero(m.Columns[i]) { // not required
				return nil
			}

			if err := m.Columns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatabaseTablesItems0) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Status); i++ {

		if m.Status[i] != nil {

			if swag.IsZero(m.Status[i]) { // not required
				return nil
			}

			if err := m.Status[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseTablesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseTablesItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseTablesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseTablesItems0ColumnsItems0 database tables items0 columns items0
//
// swagger:model DatabaseTablesItems0ColumnsItems0
type DatabaseTablesItems0ColumnsItems0 struct {

	// data type
	DataType string `json:"data_type,omitempty"`

	// default
	Default string `json:"default,omitempty"`

	// is nullable
	IsNullable bool `json:"is_nullable,omitempty"`

	// is primary key
	IsPrimaryKey bool `json:"is_primary_key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ordinal position
	OrdinalPosition int64 `json:"ordinal_position,omitempty"`
}

// Validate validates this database tables items0 columns items0
func (m *DatabaseTablesItems0ColumnsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database tables items0 columns items0 based on context it is used
func (m *DatabaseTablesItems0ColumnsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseTablesItems0ColumnsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseTablesItems0ColumnsItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseTablesItems0ColumnsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DatabaseTablesItems0StatusItems0 database tables items0 status items0
//
// swagger:model DatabaseTablesItems0StatusItems0
type DatabaseTablesItems0StatusItems0 struct {

	// aligned
	Aligned bool `json:"aligned,omitempty"`

	// node name
	NodeName string `json:"node_name,omitempty"`

	// present
	Present bool `json:"present,omitempty"`

	// replicating
	Replicating bool `json:"replicating,omitempty"`
}

// Validate validates this database tables items0 status items0
func (m *DatabaseTablesItems0StatusItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database tables items0 status items0 based on context it is used
func (m *DatabaseTablesItems0StatusItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseTablesItems0StatusItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseTablesItems0StatusItems0) UnmarshalBinary(b []byte) error {
	var res DatabaseTablesItems0StatusItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
