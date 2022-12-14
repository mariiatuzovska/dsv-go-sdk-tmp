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

// NewListSecretPathsParams creates a new ListSecretPathsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSecretPathsParams() *ListSecretPathsParams {
	return &ListSecretPathsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSecretPathsParamsWithTimeout creates a new ListSecretPathsParams object
// with the ability to set a timeout on a request.
func NewListSecretPathsParamsWithTimeout(timeout time.Duration) *ListSecretPathsParams {
	return &ListSecretPathsParams{
		timeout: timeout,
	}
}

// NewListSecretPathsParamsWithContext creates a new ListSecretPathsParams object
// with the ability to set a context for a request.
func NewListSecretPathsParamsWithContext(ctx context.Context) *ListSecretPathsParams {
	return &ListSecretPathsParams{
		Context: ctx,
	}
}

// NewListSecretPathsParamsWithHTTPClient creates a new ListSecretPathsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSecretPathsParamsWithHTTPClient(client *http.Client) *ListSecretPathsParams {
	return &ListSecretPathsParams{
		HTTPClient: client,
	}
}

/* ListSecretPathsParams contains all the parameters to send to the API endpoint
   for the list secret paths operation.

   Typically these are written to a http.Request.
*/
type ListSecretPathsParams struct {

	/* Limit.

	   The maximum number of path matches to return

	   Format: int64
	*/
	Limit *int64

	/* Path.

	   The secret path to match on
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list secret paths params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSecretPathsParams) WithDefaults() *ListSecretPathsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list secret paths params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSecretPathsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list secret paths params
func (o *ListSecretPathsParams) WithTimeout(timeout time.Duration) *ListSecretPathsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list secret paths params
func (o *ListSecretPathsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list secret paths params
func (o *ListSecretPathsParams) WithContext(ctx context.Context) *ListSecretPathsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list secret paths params
func (o *ListSecretPathsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list secret paths params
func (o *ListSecretPathsParams) WithHTTPClient(client *http.Client) *ListSecretPathsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list secret paths params
func (o *ListSecretPathsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list secret paths params
func (o *ListSecretPathsParams) WithLimit(limit *int64) *ListSecretPathsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list secret paths params
func (o *ListSecretPathsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPath adds the path to the list secret paths params
func (o *ListSecretPathsParams) WithPath(path string) *ListSecretPathsParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the list secret paths params
func (o *ListSecretPathsParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *ListSecretPathsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
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
