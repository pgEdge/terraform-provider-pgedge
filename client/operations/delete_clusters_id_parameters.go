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

// NewDeleteClustersIDParams creates a new DeleteClustersIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClustersIDParams() *DeleteClustersIDParams {
	return &DeleteClustersIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClustersIDParamsWithTimeout creates a new DeleteClustersIDParams object
// with the ability to set a timeout on a request.
func NewDeleteClustersIDParamsWithTimeout(timeout time.Duration) *DeleteClustersIDParams {
	return &DeleteClustersIDParams{
		timeout: timeout,
	}
}

// NewDeleteClustersIDParamsWithContext creates a new DeleteClustersIDParams object
// with the ability to set a context for a request.
func NewDeleteClustersIDParamsWithContext(ctx context.Context) *DeleteClustersIDParams {
	return &DeleteClustersIDParams{
		Context: ctx,
	}
}

// NewDeleteClustersIDParamsWithHTTPClient creates a new DeleteClustersIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClustersIDParamsWithHTTPClient(client *http.Client) *DeleteClustersIDParams {
	return &DeleteClustersIDParams{
		HTTPClient: client,
	}
}

/*
DeleteClustersIDParams contains all the parameters to send to the API endpoint

	for the delete clusters ID operation.

	Typically these are written to a http.Request.
*/
type DeleteClustersIDParams struct {

	// Authorization.
	//
	// Format: Bearer {access_token}
	Authorization string

	// ID.
	//
	// Format: uuid
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete clusters ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClustersIDParams) WithDefaults() *DeleteClustersIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete clusters ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClustersIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete clusters ID params
func (o *DeleteClustersIDParams) WithTimeout(timeout time.Duration) *DeleteClustersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete clusters ID params
func (o *DeleteClustersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete clusters ID params
func (o *DeleteClustersIDParams) WithContext(ctx context.Context) *DeleteClustersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete clusters ID params
func (o *DeleteClustersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete clusters ID params
func (o *DeleteClustersIDParams) WithHTTPClient(client *http.Client) *DeleteClustersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete clusters ID params
func (o *DeleteClustersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete clusters ID params
func (o *DeleteClustersIDParams) WithAuthorization(authorization string) *DeleteClustersIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete clusters ID params
func (o *DeleteClustersIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the delete clusters ID params
func (o *DeleteClustersIDParams) WithID(id strfmt.UUID) *DeleteClustersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete clusters ID params
func (o *DeleteClustersIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClustersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
