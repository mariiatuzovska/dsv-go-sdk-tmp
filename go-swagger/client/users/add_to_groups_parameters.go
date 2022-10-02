// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewAddToGroupsParams creates a new AddToGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddToGroupsParams() *AddToGroupsParams {
	return &AddToGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddToGroupsParamsWithTimeout creates a new AddToGroupsParams object
// with the ability to set a timeout on a request.
func NewAddToGroupsParamsWithTimeout(timeout time.Duration) *AddToGroupsParams {
	return &AddToGroupsParams{
		timeout: timeout,
	}
}

// NewAddToGroupsParamsWithContext creates a new AddToGroupsParams object
// with the ability to set a context for a request.
func NewAddToGroupsParamsWithContext(ctx context.Context) *AddToGroupsParams {
	return &AddToGroupsParams{
		Context: ctx,
	}
}

// NewAddToGroupsParamsWithHTTPClient creates a new AddToGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddToGroupsParamsWithHTTPClient(client *http.Client) *AddToGroupsParams {
	return &AddToGroupsParams{
		HTTPClient: client,
	}
}

/* AddToGroupsParams contains all the parameters to send to the API endpoint
   for the add to groups operation.

   Typically these are written to a http.Request.
*/
type AddToGroupsParams struct {

	// Body.
	Body *models.AddToGroupsRequest

	/* Name.

	   Full user name to lookup user by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add to groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddToGroupsParams) WithDefaults() *AddToGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add to groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddToGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add to groups params
func (o *AddToGroupsParams) WithTimeout(timeout time.Duration) *AddToGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add to groups params
func (o *AddToGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add to groups params
func (o *AddToGroupsParams) WithContext(ctx context.Context) *AddToGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add to groups params
func (o *AddToGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add to groups params
func (o *AddToGroupsParams) WithHTTPClient(client *http.Client) *AddToGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add to groups params
func (o *AddToGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add to groups params
func (o *AddToGroupsParams) WithBody(body *models.AddToGroupsRequest) *AddToGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add to groups params
func (o *AddToGroupsParams) SetBody(body *models.AddToGroupsRequest) {
	o.Body = body
}

// WithName adds the name to the add to groups params
func (o *AddToGroupsParams) WithName(name string) *AddToGroupsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the add to groups params
func (o *AddToGroupsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AddToGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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