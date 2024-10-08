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

// NewGetCloudAccountsIDParams creates a new GetCloudAccountsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCloudAccountsIDParams() *GetCloudAccountsIDParams {
	return &GetCloudAccountsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudAccountsIDParamsWithTimeout creates a new GetCloudAccountsIDParams object
// with the ability to set a timeout on a request.
func NewGetCloudAccountsIDParamsWithTimeout(timeout time.Duration) *GetCloudAccountsIDParams {
	return &GetCloudAccountsIDParams{
		timeout: timeout,
	}
}

// NewGetCloudAccountsIDParamsWithContext creates a new GetCloudAccountsIDParams object
// with the ability to set a context for a request.
func NewGetCloudAccountsIDParamsWithContext(ctx context.Context) *GetCloudAccountsIDParams {
	return &GetCloudAccountsIDParams{
		Context: ctx,
	}
}

// NewGetCloudAccountsIDParamsWithHTTPClient creates a new GetCloudAccountsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCloudAccountsIDParamsWithHTTPClient(client *http.Client) *GetCloudAccountsIDParams {
	return &GetCloudAccountsIDParams{
		HTTPClient: client,
	}
}

/*
GetCloudAccountsIDParams contains all the parameters to send to the API endpoint

	for the get cloud accounts ID operation.

	Typically these are written to a http.Request.
*/
type GetCloudAccountsIDParams struct {

	// Authorization.
	//
	// Format: Bearer {access_token}
	Authorization string

	/* ID.

	   ID of the cloud account to retrieve.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cloud accounts ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudAccountsIDParams) WithDefaults() *GetCloudAccountsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cloud accounts ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudAccountsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) WithTimeout(timeout time.Duration) *GetCloudAccountsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) WithContext(ctx context.Context) *GetCloudAccountsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) WithHTTPClient(client *http.Client) *GetCloudAccountsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) WithAuthorization(authorization string) *GetCloudAccountsIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) WithID(id strfmt.UUID) *GetCloudAccountsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cloud accounts ID params
func (o *GetCloudAccountsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudAccountsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
