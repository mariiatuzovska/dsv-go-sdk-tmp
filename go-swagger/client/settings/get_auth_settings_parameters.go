// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewGetAuthSettingsParams creates a new GetAuthSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthSettingsParams() *GetAuthSettingsParams {
	return &GetAuthSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthSettingsParamsWithTimeout creates a new GetAuthSettingsParams object
// with the ability to set a timeout on a request.
func NewGetAuthSettingsParamsWithTimeout(timeout time.Duration) *GetAuthSettingsParams {
	return &GetAuthSettingsParams{
		timeout: timeout,
	}
}

// NewGetAuthSettingsParamsWithContext creates a new GetAuthSettingsParams object
// with the ability to set a context for a request.
func NewGetAuthSettingsParamsWithContext(ctx context.Context) *GetAuthSettingsParams {
	return &GetAuthSettingsParams{
		Context: ctx,
	}
}

// NewGetAuthSettingsParamsWithHTTPClient creates a new GetAuthSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthSettingsParamsWithHTTPClient(client *http.Client) *GetAuthSettingsParams {
	return &GetAuthSettingsParams{
		HTTPClient: client,
	}
}

/* GetAuthSettingsParams contains all the parameters to send to the API endpoint
   for the get auth settings operation.

   Typically these are written to a http.Request.
*/
type GetAuthSettingsParams struct {

	/* Name.

	   Full name to lookup authentication settings by.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthSettingsParams) WithDefaults() *GetAuthSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get auth settings params
func (o *GetAuthSettingsParams) WithTimeout(timeout time.Duration) *GetAuthSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth settings params
func (o *GetAuthSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth settings params
func (o *GetAuthSettingsParams) WithContext(ctx context.Context) *GetAuthSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth settings params
func (o *GetAuthSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth settings params
func (o *GetAuthSettingsParams) WithHTTPClient(client *http.Client) *GetAuthSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth settings params
func (o *GetAuthSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get auth settings params
func (o *GetAuthSettingsParams) WithName(name string) *GetAuthSettingsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get auth settings params
func (o *GetAuthSettingsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
