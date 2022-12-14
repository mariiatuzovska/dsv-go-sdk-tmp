// Code generated by go-swagger; DO NOT EDIT.

package engines

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

// NewPingEngineParams creates a new PingEngineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPingEngineParams() *PingEngineParams {
	return &PingEngineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPingEngineParamsWithTimeout creates a new PingEngineParams object
// with the ability to set a timeout on a request.
func NewPingEngineParamsWithTimeout(timeout time.Duration) *PingEngineParams {
	return &PingEngineParams{
		timeout: timeout,
	}
}

// NewPingEngineParamsWithContext creates a new PingEngineParams object
// with the ability to set a context for a request.
func NewPingEngineParamsWithContext(ctx context.Context) *PingEngineParams {
	return &PingEngineParams{
		Context: ctx,
	}
}

// NewPingEngineParamsWithHTTPClient creates a new PingEngineParams object
// with the ability to set a custom HTTPClient for a request.
func NewPingEngineParamsWithHTTPClient(client *http.Client) *PingEngineParams {
	return &PingEngineParams{
		HTTPClient: client,
	}
}

/* PingEngineParams contains all the parameters to send to the API endpoint
   for the ping engine operation.

   Typically these are written to a http.Request.
*/
type PingEngineParams struct {

	/* Name.

	   Name of engine
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ping engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingEngineParams) WithDefaults() *PingEngineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ping engine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PingEngineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ping engine params
func (o *PingEngineParams) WithTimeout(timeout time.Duration) *PingEngineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ping engine params
func (o *PingEngineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ping engine params
func (o *PingEngineParams) WithContext(ctx context.Context) *PingEngineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ping engine params
func (o *PingEngineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ping engine params
func (o *PingEngineParams) WithHTTPClient(client *http.Client) *PingEngineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ping engine params
func (o *PingEngineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the ping engine params
func (o *PingEngineParams) WithName(name string) *PingEngineParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ping engine params
func (o *PingEngineParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PingEngineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
