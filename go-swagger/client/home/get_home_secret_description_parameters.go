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

// NewGetHomeSecretDescriptionParams creates a new GetHomeSecretDescriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHomeSecretDescriptionParams() *GetHomeSecretDescriptionParams {
	return &GetHomeSecretDescriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHomeSecretDescriptionParamsWithTimeout creates a new GetHomeSecretDescriptionParams object
// with the ability to set a timeout on a request.
func NewGetHomeSecretDescriptionParamsWithTimeout(timeout time.Duration) *GetHomeSecretDescriptionParams {
	return &GetHomeSecretDescriptionParams{
		timeout: timeout,
	}
}

// NewGetHomeSecretDescriptionParamsWithContext creates a new GetHomeSecretDescriptionParams object
// with the ability to set a context for a request.
func NewGetHomeSecretDescriptionParamsWithContext(ctx context.Context) *GetHomeSecretDescriptionParams {
	return &GetHomeSecretDescriptionParams{
		Context: ctx,
	}
}

// NewGetHomeSecretDescriptionParamsWithHTTPClient creates a new GetHomeSecretDescriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHomeSecretDescriptionParamsWithHTTPClient(client *http.Client) *GetHomeSecretDescriptionParams {
	return &GetHomeSecretDescriptionParams{
		HTTPClient: client,
	}
}

/* GetHomeSecretDescriptionParams contains all the parameters to send to the API endpoint
   for the get home secret description operation.

   Typically these are written to a http.Request.
*/
type GetHomeSecretDescriptionParams struct {

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

// WithDefaults hydrates default values in the get home secret description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHomeSecretDescriptionParams) WithDefaults() *GetHomeSecretDescriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get home secret description params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHomeSecretDescriptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get home secret description params
func (o *GetHomeSecretDescriptionParams) WithTimeout(timeout time.Duration) *GetHomeSecretDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get home secret description params
func (o *GetHomeSecretDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get home secret description params
func (o *GetHomeSecretDescriptionParams) WithContext(ctx context.Context) *GetHomeSecretDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get home secret description params
func (o *GetHomeSecretDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get home secret description params
func (o *GetHomeSecretDescriptionParams) WithHTTPClient(client *http.Client) *GetHomeSecretDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get home secret description params
func (o *GetHomeSecretDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get home secret description params
func (o *GetHomeSecretDescriptionParams) WithID(id *string) *GetHomeSecretDescriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get home secret description params
func (o *GetHomeSecretDescriptionParams) SetID(id *string) {
	o.ID = id
}

// WithPath adds the path to the get home secret description params
func (o *GetHomeSecretDescriptionParams) WithPath(path string) *GetHomeSecretDescriptionParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get home secret description params
func (o *GetHomeSecretDescriptionParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *GetHomeSecretDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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