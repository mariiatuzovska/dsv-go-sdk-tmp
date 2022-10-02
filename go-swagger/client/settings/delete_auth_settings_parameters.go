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
	"github.com/go-openapi/swag"
)

// NewDeleteAuthSettingsParams creates a new DeleteAuthSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAuthSettingsParams() *DeleteAuthSettingsParams {
	return &DeleteAuthSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuthSettingsParamsWithTimeout creates a new DeleteAuthSettingsParams object
// with the ability to set a timeout on a request.
func NewDeleteAuthSettingsParamsWithTimeout(timeout time.Duration) *DeleteAuthSettingsParams {
	return &DeleteAuthSettingsParams{
		timeout: timeout,
	}
}

// NewDeleteAuthSettingsParamsWithContext creates a new DeleteAuthSettingsParams object
// with the ability to set a context for a request.
func NewDeleteAuthSettingsParamsWithContext(ctx context.Context) *DeleteAuthSettingsParams {
	return &DeleteAuthSettingsParams{
		Context: ctx,
	}
}

// NewDeleteAuthSettingsParamsWithHTTPClient creates a new DeleteAuthSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAuthSettingsParamsWithHTTPClient(client *http.Client) *DeleteAuthSettingsParams {
	return &DeleteAuthSettingsParams{
		HTTPClient: client,
	}
}

/* DeleteAuthSettingsParams contains all the parameters to send to the API endpoint
   for the delete auth settings operation.

   Typically these are written to a http.Request.
*/
type DeleteAuthSettingsParams struct {

	/* Force.

	   Delete immediately
	*/
	Force *bool

	/* Name.

	   Full name to lookup authentication settings by.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAuthSettingsParams) WithDefaults() *DeleteAuthSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete auth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAuthSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete auth settings params
func (o *DeleteAuthSettingsParams) WithTimeout(timeout time.Duration) *DeleteAuthSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete auth settings params
func (o *DeleteAuthSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete auth settings params
func (o *DeleteAuthSettingsParams) WithContext(ctx context.Context) *DeleteAuthSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete auth settings params
func (o *DeleteAuthSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete auth settings params
func (o *DeleteAuthSettingsParams) WithHTTPClient(client *http.Client) *DeleteAuthSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete auth settings params
func (o *DeleteAuthSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the delete auth settings params
func (o *DeleteAuthSettingsParams) WithForce(force *bool) *DeleteAuthSettingsParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete auth settings params
func (o *DeleteAuthSettingsParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the delete auth settings params
func (o *DeleteAuthSettingsParams) WithName(name string) *DeleteAuthSettingsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete auth settings params
func (o *DeleteAuthSettingsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuthSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}