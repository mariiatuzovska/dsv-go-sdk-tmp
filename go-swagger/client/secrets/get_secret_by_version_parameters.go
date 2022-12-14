// Code generated by go-swagger; DO NOT EDIT.

package secrets

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
	"github.com/go-openapi/swag"
)

// NewGetSecretByVersionParams creates a new GetSecretByVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSecretByVersionParams() *GetSecretByVersionParams {
	return &GetSecretByVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecretByVersionParamsWithTimeout creates a new GetSecretByVersionParams object
// with the ability to set a timeout on a request.
func NewGetSecretByVersionParamsWithTimeout(timeout time.Duration) *GetSecretByVersionParams {
	return &GetSecretByVersionParams{
		timeout: timeout,
	}
}

// NewGetSecretByVersionParamsWithContext creates a new GetSecretByVersionParams object
// with the ability to set a context for a request.
func NewGetSecretByVersionParamsWithContext(ctx context.Context) *GetSecretByVersionParams {
	return &GetSecretByVersionParams{
		Context: ctx,
	}
}

// NewGetSecretByVersionParamsWithHTTPClient creates a new GetSecretByVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSecretByVersionParamsWithHTTPClient(client *http.Client) *GetSecretByVersionParams {
	return &GetSecretByVersionParams{
		HTTPClient: client,
	}
}

/* GetSecretByVersionParams contains all the parameters to send to the API endpoint
   for the get secret by version operation.

   Typically these are written to a http.Request.
*/
type GetSecretByVersionParams struct {

	/* Path.

	   The full secret path i.e. servers/prod/webserver-01
	*/
	Path string

	/* Version.

	   Versions to return

	   Format: int64
	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get secret by version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecretByVersionParams) WithDefaults() *GetSecretByVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get secret by version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecretByVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get secret by version params
func (o *GetSecretByVersionParams) WithTimeout(timeout time.Duration) *GetSecretByVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get secret by version params
func (o *GetSecretByVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get secret by version params
func (o *GetSecretByVersionParams) WithContext(ctx context.Context) *GetSecretByVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get secret by version params
func (o *GetSecretByVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get secret by version params
func (o *GetSecretByVersionParams) WithHTTPClient(client *http.Client) *GetSecretByVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get secret by version params
func (o *GetSecretByVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the get secret by version params
func (o *GetSecretByVersionParams) WithPath(path string) *GetSecretByVersionParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get secret by version params
func (o *GetSecretByVersionParams) SetPath(path string) {
	o.Path = path
}

// WithVersion adds the version to the get secret by version params
func (o *GetSecretByVersionParams) WithVersion(version int64) *GetSecretByVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get secret by version params
func (o *GetSecretByVersionParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecretByVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", swag.FormatInt64(o.Version)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
