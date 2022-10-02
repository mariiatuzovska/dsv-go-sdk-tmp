// Code generated by go-swagger; DO NOT EDIT.

package home

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

// NewDeleteHomeSecretParams creates a new DeleteHomeSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteHomeSecretParams() *DeleteHomeSecretParams {
	return &DeleteHomeSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHomeSecretParamsWithTimeout creates a new DeleteHomeSecretParams object
// with the ability to set a timeout on a request.
func NewDeleteHomeSecretParamsWithTimeout(timeout time.Duration) *DeleteHomeSecretParams {
	return &DeleteHomeSecretParams{
		timeout: timeout,
	}
}

// NewDeleteHomeSecretParamsWithContext creates a new DeleteHomeSecretParams object
// with the ability to set a context for a request.
func NewDeleteHomeSecretParamsWithContext(ctx context.Context) *DeleteHomeSecretParams {
	return &DeleteHomeSecretParams{
		Context: ctx,
	}
}

// NewDeleteHomeSecretParamsWithHTTPClient creates a new DeleteHomeSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteHomeSecretParamsWithHTTPClient(client *http.Client) *DeleteHomeSecretParams {
	return &DeleteHomeSecretParams{
		HTTPClient: client,
	}
}

/* DeleteHomeSecretParams contains all the parameters to send to the API endpoint
   for the delete home secret operation.

   Typically these are written to a http.Request.
*/
type DeleteHomeSecretParams struct {

	/* ID.

	   Unique uuid identifying a secret
	*/
	ID *string

	/* Path.

	   The full secret path i.e. servers/prod/webserver-01
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete home secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHomeSecretParams) WithDefaults() *DeleteHomeSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete home secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHomeSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete home secret params
func (o *DeleteHomeSecretParams) WithTimeout(timeout time.Duration) *DeleteHomeSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete home secret params
func (o *DeleteHomeSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete home secret params
func (o *DeleteHomeSecretParams) WithContext(ctx context.Context) *DeleteHomeSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete home secret params
func (o *DeleteHomeSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete home secret params
func (o *DeleteHomeSecretParams) WithHTTPClient(client *http.Client) *DeleteHomeSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete home secret params
func (o *DeleteHomeSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete home secret params
func (o *DeleteHomeSecretParams) WithID(id *string) *DeleteHomeSecretParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete home secret params
func (o *DeleteHomeSecretParams) SetID(id *string) {
	o.ID = id
}

// WithPath adds the path to the delete home secret params
func (o *DeleteHomeSecretParams) WithPath(path string) *DeleteHomeSecretParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the delete home secret params
func (o *DeleteHomeSecretParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHomeSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
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