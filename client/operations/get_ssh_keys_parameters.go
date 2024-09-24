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

// NewGetSSHKeysParams creates a new GetSSHKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSSHKeysParams() *GetSSHKeysParams {
	return &GetSSHKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSSHKeysParamsWithTimeout creates a new GetSSHKeysParams object
// with the ability to set a timeout on a request.
func NewGetSSHKeysParamsWithTimeout(timeout time.Duration) *GetSSHKeysParams {
	return &GetSSHKeysParams{
		timeout: timeout,
	}
}

// NewGetSSHKeysParamsWithContext creates a new GetSSHKeysParams object
// with the ability to set a context for a request.
func NewGetSSHKeysParamsWithContext(ctx context.Context) *GetSSHKeysParams {
	return &GetSSHKeysParams{
		Context: ctx,
	}
}

// NewGetSSHKeysParamsWithHTTPClient creates a new GetSSHKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSSHKeysParamsWithHTTPClient(client *http.Client) *GetSSHKeysParams {
	return &GetSSHKeysParams{
		HTTPClient: client,
	}
}

/*
GetSSHKeysParams contains all the parameters to send to the API endpoint

	for the get SSH keys operation.

	Typically these are written to a http.Request.
*/
type GetSSHKeysParams struct {

	// Authorization.
	//
	// Format: Bearer {access_token}
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get SSH keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSSHKeysParams) WithDefaults() *GetSSHKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get SSH keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSSHKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get SSH keys params
func (o *GetSSHKeysParams) WithTimeout(timeout time.Duration) *GetSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get SSH keys params
func (o *GetSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get SSH keys params
func (o *GetSSHKeysParams) WithContext(ctx context.Context) *GetSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get SSH keys params
func (o *GetSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get SSH keys params
func (o *GetSSHKeysParams) WithHTTPClient(client *http.Client) *GetSSHKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get SSH keys params
func (o *GetSSHKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get SSH keys params
func (o *GetSSHKeysParams) WithAuthorization(authorization string) *GetSSHKeysParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get SSH keys params
func (o *GetSSHKeysParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *GetSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}