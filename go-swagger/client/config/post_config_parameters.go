// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewPostConfigParams creates a new PostConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostConfigParams() *PostConfigParams {
	return &PostConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostConfigParamsWithTimeout creates a new PostConfigParams object
// with the ability to set a timeout on a request.
func NewPostConfigParamsWithTimeout(timeout time.Duration) *PostConfigParams {
	return &PostConfigParams{
		timeout: timeout,
	}
}

// NewPostConfigParamsWithContext creates a new PostConfigParams object
// with the ability to set a context for a request.
func NewPostConfigParamsWithContext(ctx context.Context) *PostConfigParams {
	return &PostConfigParams{
		Context: ctx,
	}
}

// NewPostConfigParamsWithHTTPClient creates a new PostConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostConfigParamsWithHTTPClient(client *http.Client) *PostConfigParams {
	return &PostConfigParams{
		HTTPClient: client,
	}
}

/* PostConfigParams contains all the parameters to send to the API endpoint
   for the post config operation.

   Typically these are written to a http.Request.
*/
type PostConfigParams struct {

	// Body.
	Body *models.PostConfigModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConfigParams) WithDefaults() *PostConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post config params
func (o *PostConfigParams) WithTimeout(timeout time.Duration) *PostConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post config params
func (o *PostConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post config params
func (o *PostConfigParams) WithContext(ctx context.Context) *PostConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post config params
func (o *PostConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post config params
func (o *PostConfigParams) WithHTTPClient(client *http.Client) *PostConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post config params
func (o *PostConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post config params
func (o *PostConfigParams) WithBody(body *models.PostConfigModel) *PostConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post config params
func (o *PostConfigParams) SetBody(body *models.PostConfigModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
