// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
)

// NewPatchDatabasesIDParams creates a new PatchDatabasesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDatabasesIDParams() *PatchDatabasesIDParams {
	return &PatchDatabasesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDatabasesIDParamsWithTimeout creates a new PatchDatabasesIDParams object
// with the ability to set a timeout on a request.
func NewPatchDatabasesIDParamsWithTimeout(timeout time.Duration) *PatchDatabasesIDParams {
	return &PatchDatabasesIDParams{
		timeout: timeout,
	}
}

// NewPatchDatabasesIDParamsWithContext creates a new PatchDatabasesIDParams object
// with the ability to set a context for a request.
func NewPatchDatabasesIDParamsWithContext(ctx context.Context) *PatchDatabasesIDParams {
	return &PatchDatabasesIDParams{
		Context: ctx,
	}
}

// NewPatchDatabasesIDParamsWithHTTPClient creates a new PatchDatabasesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDatabasesIDParamsWithHTTPClient(client *http.Client) *PatchDatabasesIDParams {
	return &PatchDatabasesIDParams{
		HTTPClient: client,
	}
}

/*
PatchDatabasesIDParams contains all the parameters to send to the API endpoint

	for the patch databases ID operation.

	Typically these are written to a http.Request.
*/
type PatchDatabasesIDParams struct {

	// Authorization.
	//
	// Format: Bearer {access_token}
	Authorization string

	/* Body.

	   The database parameters to update.
	*/
	Body *models.UpdateDatabaseInput

	/* ID.

	   ID of the database to update.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch databases ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDatabasesIDParams) WithDefaults() *PatchDatabasesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch databases ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDatabasesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch databases ID params
func (o *PatchDatabasesIDParams) WithTimeout(timeout time.Duration) *PatchDatabasesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch databases ID params
func (o *PatchDatabasesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch databases ID params
func (o *PatchDatabasesIDParams) WithContext(ctx context.Context) *PatchDatabasesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch databases ID params
func (o *PatchDatabasesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch databases ID params
func (o *PatchDatabasesIDParams) WithHTTPClient(client *http.Client) *PatchDatabasesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch databases ID params
func (o *PatchDatabasesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch databases ID params
func (o *PatchDatabasesIDParams) WithAuthorization(authorization string) *PatchDatabasesIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch databases ID params
func (o *PatchDatabasesIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the patch databases ID params
func (o *PatchDatabasesIDParams) WithBody(body *models.UpdateDatabaseInput) *PatchDatabasesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch databases ID params
func (o *PatchDatabasesIDParams) SetBody(body *models.UpdateDatabaseInput) {
	o.Body = body
}

// WithID adds the id to the patch databases ID params
func (o *PatchDatabasesIDParams) WithID(id strfmt.UUID) *PatchDatabasesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch databases ID params
func (o *PatchDatabasesIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDatabasesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}