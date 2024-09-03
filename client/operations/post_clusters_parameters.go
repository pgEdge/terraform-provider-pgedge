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

// NewPostClustersParams creates a new PostClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostClustersParams() *PostClustersParams {
	return &PostClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostClustersParamsWithTimeout creates a new PostClustersParams object
// with the ability to set a timeout on a request.
func NewPostClustersParamsWithTimeout(timeout time.Duration) *PostClustersParams {
	return &PostClustersParams{
		timeout: timeout,
	}
}

// NewPostClustersParamsWithContext creates a new PostClustersParams object
// with the ability to set a context for a request.
func NewPostClustersParamsWithContext(ctx context.Context) *PostClustersParams {
	return &PostClustersParams{
		Context: ctx,
	}
}

// NewPostClustersParamsWithHTTPClient creates a new PostClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostClustersParamsWithHTTPClient(client *http.Client) *PostClustersParams {
	return &PostClustersParams{
		HTTPClient: client,
	}
}

/*
PostClustersParams contains all the parameters to send to the API endpoint

	for the post clusters operation.

	Typically these are written to a http.Request.
*/
type PostClustersParams struct {

	// Authorization.
	//
	// Format: Bearer {access_token}
	Authorization string

	/* Body.

	   Cluster creation request body
	*/
	Body *models.ClusterCreationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostClustersParams) WithDefaults() *PostClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post clusters params
func (o *PostClustersParams) WithTimeout(timeout time.Duration) *PostClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clusters params
func (o *PostClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clusters params
func (o *PostClustersParams) WithContext(ctx context.Context) *PostClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clusters params
func (o *PostClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post clusters params
func (o *PostClustersParams) WithHTTPClient(client *http.Client) *PostClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post clusters params
func (o *PostClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post clusters params
func (o *PostClustersParams) WithAuthorization(authorization string) *PostClustersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post clusters params
func (o *PostClustersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the post clusters params
func (o *PostClustersParams) WithBody(body *models.ClusterCreationRequest) *PostClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post clusters params
func (o *PostClustersParams) SetBody(body *models.ClusterCreationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
