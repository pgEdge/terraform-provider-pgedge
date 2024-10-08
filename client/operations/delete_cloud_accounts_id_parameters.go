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
)

// NewDeleteCloudAccountsIDParams creates a new DeleteCloudAccountsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCloudAccountsIDParams() *DeleteCloudAccountsIDParams {
	return &DeleteCloudAccountsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudAccountsIDParamsWithTimeout creates a new DeleteCloudAccountsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteCloudAccountsIDParamsWithTimeout(timeout time.Duration) *DeleteCloudAccountsIDParams {
	return &DeleteCloudAccountsIDParams{
		timeout: timeout,
	}
}

// NewDeleteCloudAccountsIDParamsWithContext creates a new DeleteCloudAccountsIDParams object
// with the ability to set a context for a request.
func NewDeleteCloudAccountsIDParamsWithContext(ctx context.Context) *DeleteCloudAccountsIDParams {
	return &DeleteCloudAccountsIDParams{
		Context: ctx,
	}
}

// NewDeleteCloudAccountsIDParamsWithHTTPClient creates a new DeleteCloudAccountsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCloudAccountsIDParamsWithHTTPClient(client *http.Client) *DeleteCloudAccountsIDParams {
	return &DeleteCloudAccountsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteCloudAccountsIDParams contains all the parameters to send to the API endpoint

	for the delete cloud accounts ID operation.

	Typically these are written to a http.Request.
*/
type DeleteCloudAccountsIDParams struct {

	// Authorization.
	//
	// Format: Bearer {access_token}
	Authorization string

	/* ID.

	   ID of the cloud account to offboard.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cloud accounts ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCloudAccountsIDParams) WithDefaults() *DeleteCloudAccountsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cloud accounts ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCloudAccountsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) WithTimeout(timeout time.Duration) *DeleteCloudAccountsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) WithContext(ctx context.Context) *DeleteCloudAccountsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) WithHTTPClient(client *http.Client) *DeleteCloudAccountsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) WithAuthorization(authorization string) *DeleteCloudAccountsIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) WithID(id strfmt.UUID) *DeleteCloudAccountsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete cloud accounts ID params
func (o *DeleteCloudAccountsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudAccountsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
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
