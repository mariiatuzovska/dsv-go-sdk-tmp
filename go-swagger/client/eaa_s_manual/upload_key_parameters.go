// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_manual

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

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// NewUploadKeyParams creates a new UploadKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadKeyParams() *UploadKeyParams {
	return &UploadKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadKeyParamsWithTimeout creates a new UploadKeyParams object
// with the ability to set a timeout on a request.
func NewUploadKeyParamsWithTimeout(timeout time.Duration) *UploadKeyParams {
	return &UploadKeyParams{
		timeout: timeout,
	}
}

// NewUploadKeyParamsWithContext creates a new UploadKeyParams object
// with the ability to set a context for a request.
func NewUploadKeyParamsWithContext(ctx context.Context) *UploadKeyParams {
	return &UploadKeyParams{
		Context: ctx,
	}
}

// NewUploadKeyParamsWithHTTPClient creates a new UploadKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadKeyParamsWithHTTPClient(client *http.Client) *UploadKeyParams {
	return &UploadKeyParams{
		HTTPClient: client,
	}
}

/* UploadKeyParams contains all the parameters to send to the API endpoint
   for the upload key operation.

   Typically these are written to a http.Request.
*/
type UploadKeyParams struct {

	// Body.
	Body *models.ManualKeyData

	/* Path.

	   The full key path, for example, mykeys/key01
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadKeyParams) WithDefaults() *UploadKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload key params
func (o *UploadKeyParams) WithTimeout(timeout time.Duration) *UploadKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload key params
func (o *UploadKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload key params
func (o *UploadKeyParams) WithContext(ctx context.Context) *UploadKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload key params
func (o *UploadKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload key params
func (o *UploadKeyParams) WithHTTPClient(client *http.Client) *UploadKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload key params
func (o *UploadKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upload key params
func (o *UploadKeyParams) WithBody(body *models.ManualKeyData) *UploadKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload key params
func (o *UploadKeyParams) SetBody(body *models.ManualKeyData) {
	o.Body = body
}

// WithPath adds the path to the upload key params
func (o *UploadKeyParams) WithPath(path string) *UploadKeyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the upload key params
func (o *UploadKeyParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *UploadKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}