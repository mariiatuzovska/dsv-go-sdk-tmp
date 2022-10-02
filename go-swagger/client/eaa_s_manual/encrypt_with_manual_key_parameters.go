// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_manual

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

// NewEncryptWithManualKeyParams creates a new EncryptWithManualKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEncryptWithManualKeyParams() *EncryptWithManualKeyParams {
	return &EncryptWithManualKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEncryptWithManualKeyParamsWithTimeout creates a new EncryptWithManualKeyParams object
// with the ability to set a timeout on a request.
func NewEncryptWithManualKeyParamsWithTimeout(timeout time.Duration) *EncryptWithManualKeyParams {
	return &EncryptWithManualKeyParams{
		timeout: timeout,
	}
}

// NewEncryptWithManualKeyParamsWithContext creates a new EncryptWithManualKeyParams object
// with the ability to set a context for a request.
func NewEncryptWithManualKeyParamsWithContext(ctx context.Context) *EncryptWithManualKeyParams {
	return &EncryptWithManualKeyParams{
		Context: ctx,
	}
}

// NewEncryptWithManualKeyParamsWithHTTPClient creates a new EncryptWithManualKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewEncryptWithManualKeyParamsWithHTTPClient(client *http.Client) *EncryptWithManualKeyParams {
	return &EncryptWithManualKeyParams{
		HTTPClient: client,
	}
}

/* EncryptWithManualKeyParams contains all the parameters to send to the API endpoint
   for the encrypt with manual key operation.

   Typically these are written to a http.Request.
*/
type EncryptWithManualKeyParams struct {

	/* Path.

	   The path to data key
	*/
	Path string

	/* Plaintext.

	   A value to be encrypted
	*/
	Plaintext string

	/* Version.

	   The version of the key with which to encrypt data
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the encrypt with manual key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EncryptWithManualKeyParams) WithDefaults() *EncryptWithManualKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the encrypt with manual key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EncryptWithManualKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) WithTimeout(timeout time.Duration) *EncryptWithManualKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) WithContext(ctx context.Context) *EncryptWithManualKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) WithHTTPClient(client *http.Client) *EncryptWithManualKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPath adds the path to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) WithPath(path string) *EncryptWithManualKeyParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) SetPath(path string) {
	o.Path = path
}

// WithPlaintext adds the plaintext to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) WithPlaintext(plaintext string) *EncryptWithManualKeyParams {
	o.SetPlaintext(plaintext)
	return o
}

// SetPlaintext adds the plaintext to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) SetPlaintext(plaintext string) {
	o.Plaintext = plaintext
}

// WithVersion adds the version to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) WithVersion(version *string) *EncryptWithManualKeyParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the encrypt with manual key params
func (o *EncryptWithManualKeyParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *EncryptWithManualKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	// query param plaintext
	qrPlaintext := o.Plaintext
	qPlaintext := qrPlaintext
	if qPlaintext != "" {

		if err := r.SetQueryParam("plaintext", qPlaintext); err != nil {
			return err
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}