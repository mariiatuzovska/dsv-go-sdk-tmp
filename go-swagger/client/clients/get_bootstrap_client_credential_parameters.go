// Code generated by go-swagger; DO NOT EDIT.

package clients

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

// NewGetBootstrapClientCredentialParams creates a new GetBootstrapClientCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBootstrapClientCredentialParams() *GetBootstrapClientCredentialParams {
	return &GetBootstrapClientCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBootstrapClientCredentialParamsWithTimeout creates a new GetBootstrapClientCredentialParams object
// with the ability to set a timeout on a request.
func NewGetBootstrapClientCredentialParamsWithTimeout(timeout time.Duration) *GetBootstrapClientCredentialParams {
	return &GetBootstrapClientCredentialParams{
		timeout: timeout,
	}
}

// NewGetBootstrapClientCredentialParamsWithContext creates a new GetBootstrapClientCredentialParams object
// with the ability to set a context for a request.
func NewGetBootstrapClientCredentialParamsWithContext(ctx context.Context) *GetBootstrapClientCredentialParams {
	return &GetBootstrapClientCredentialParams{
		Context: ctx,
	}
}

// NewGetBootstrapClientCredentialParamsWithHTTPClient creates a new GetBootstrapClientCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBootstrapClientCredentialParamsWithHTTPClient(client *http.Client) *GetBootstrapClientCredentialParams {
	return &GetBootstrapClientCredentialParams{
		HTTPClient: client,
	}
}

/* GetBootstrapClientCredentialParams contains all the parameters to send to the API endpoint
   for the get bootstrap client credential operation.

   Typically these are written to a http.Request.
*/
type GetBootstrapClientCredentialParams struct {

	/* ClientID.

	   ClientId property of the client credentials
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bootstrap client credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBootstrapClientCredentialParams) WithDefaults() *GetBootstrapClientCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bootstrap client credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBootstrapClientCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) WithTimeout(timeout time.Duration) *GetBootstrapClientCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) WithContext(ctx context.Context) *GetBootstrapClientCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) WithHTTPClient(client *http.Client) *GetBootstrapClientCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the clientID to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) WithID(clientID string) *GetBootstrapClientCredentialParams {
	o.SetID(clientID)
	return o
}

// SetID adds the clientId to the get bootstrap client credential params
func (o *GetBootstrapClientCredentialParams) SetID(clientID string) {
	o.ID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBootstrapClientCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
