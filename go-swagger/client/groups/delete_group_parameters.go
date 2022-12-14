// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewDeleteGroupParams creates a new DeleteGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGroupParams() *DeleteGroupParams {
	return &DeleteGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupParamsWithTimeout creates a new DeleteGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteGroupParamsWithTimeout(timeout time.Duration) *DeleteGroupParams {
	return &DeleteGroupParams{
		timeout: timeout,
	}
}

// NewDeleteGroupParamsWithContext creates a new DeleteGroupParams object
// with the ability to set a context for a request.
func NewDeleteGroupParamsWithContext(ctx context.Context) *DeleteGroupParams {
	return &DeleteGroupParams{
		Context: ctx,
	}
}

// NewDeleteGroupParamsWithHTTPClient creates a new DeleteGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGroupParamsWithHTTPClient(client *http.Client) *DeleteGroupParams {
	return &DeleteGroupParams{
		HTTPClient: client,
	}
}

/* DeleteGroupParams contains all the parameters to send to the API endpoint
   for the delete group operation.

   Typically these are written to a http.Request.
*/
type DeleteGroupParams struct {

	/* Force.

	   Delete immediately
	*/
	Force *bool

	/* Name.

	   Group name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupParams) WithDefaults() *DeleteGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete group params
func (o *DeleteGroupParams) WithTimeout(timeout time.Duration) *DeleteGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete group params
func (o *DeleteGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete group params
func (o *DeleteGroupParams) WithContext(ctx context.Context) *DeleteGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete group params
func (o *DeleteGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete group params
func (o *DeleteGroupParams) WithHTTPClient(client *http.Client) *DeleteGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete group params
func (o *DeleteGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the delete group params
func (o *DeleteGroupParams) WithForce(force *bool) *DeleteGroupParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete group params
func (o *DeleteGroupParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the delete group params
func (o *DeleteGroupParams) WithName(name string) *DeleteGroupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete group params
func (o *DeleteGroupParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
