// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_auto

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

// NewRotationRequestParams creates a new RotationRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRotationRequestParams() *RotationRequestParams {
	return &RotationRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRotationRequestParamsWithTimeout creates a new RotationRequestParams object
// with the ability to set a timeout on a request.
func NewRotationRequestParamsWithTimeout(timeout time.Duration) *RotationRequestParams {
	return &RotationRequestParams{
		timeout: timeout,
	}
}

// NewRotationRequestParamsWithContext creates a new RotationRequestParams object
// with the ability to set a context for a request.
func NewRotationRequestParamsWithContext(ctx context.Context) *RotationRequestParams {
	return &RotationRequestParams{
		Context: ctx,
	}
}

// NewRotationRequestParamsWithHTTPClient creates a new RotationRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRotationRequestParamsWithHTTPClient(client *http.Client) *RotationRequestParams {
	return &RotationRequestParams{
		HTTPClient: client,
	}
}

/* RotationRequestParams contains all the parameters to send to the API endpoint
   for the rotation request operation.

   Typically these are written to a http.Request.
*/
type RotationRequestParams struct {

	/* Ciphertext.

	   A value to be rotated (re-encrypted)
	*/
	Ciphertext []uint8

	/* EndingVersion.

	   The ending version of the key with which to re-encrypt data
	*/
	EndingVersion *string

	/* Path.

	   The path to data key
	*/
	Path string

	/* StartingVersion.

	   The starting version of the key with which to re-encrypt data
	*/
	StartingVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rotation request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RotationRequestParams) WithDefaults() *RotationRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rotation request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RotationRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rotation request params
func (o *RotationRequestParams) WithTimeout(timeout time.Duration) *RotationRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rotation request params
func (o *RotationRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rotation request params
func (o *RotationRequestParams) WithContext(ctx context.Context) *RotationRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rotation request params
func (o *RotationRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rotation request params
func (o *RotationRequestParams) WithHTTPClient(client *http.Client) *RotationRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rotation request params
func (o *RotationRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCiphertext adds the ciphertext to the rotation request params
func (o *RotationRequestParams) WithCiphertext(ciphertext []uint8) *RotationRequestParams {
	o.SetCiphertext(ciphertext)
	return o
}

// SetCiphertext adds the ciphertext to the rotation request params
func (o *RotationRequestParams) SetCiphertext(ciphertext []uint8) {
	o.Ciphertext = ciphertext
}

// WithEndingVersion adds the endingVersion to the rotation request params
func (o *RotationRequestParams) WithEndingVersion(endingVersion *string) *RotationRequestParams {
	o.SetEndingVersion(endingVersion)
	return o
}

// SetEndingVersion adds the endingVersion to the rotation request params
func (o *RotationRequestParams) SetEndingVersion(endingVersion *string) {
	o.EndingVersion = endingVersion
}

// WithPath adds the path to the rotation request params
func (o *RotationRequestParams) WithPath(path string) *RotationRequestParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the rotation request params
func (o *RotationRequestParams) SetPath(path string) {
	o.Path = path
}

// WithStartingVersion adds the startingVersion to the rotation request params
func (o *RotationRequestParams) WithStartingVersion(startingVersion string) *RotationRequestParams {
	o.SetStartingVersion(startingVersion)
	return o
}

// SetStartingVersion adds the startingVersion to the rotation request params
func (o *RotationRequestParams) SetStartingVersion(startingVersion string) {
	o.StartingVersion = startingVersion
}

// WriteToRequest writes these params to a swagger request
func (o *RotationRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ciphertext != nil {

		// binding items for ciphertext
		joinedCiphertext := o.bindParamCiphertext(reg)

		// query array param ciphertext
		if err := r.SetQueryParam("ciphertext", joinedCiphertext...); err != nil {
			return err
		}
	}

	if o.EndingVersion != nil {

		// query param endingVersion
		var qrEndingVersion string

		if o.EndingVersion != nil {
			qrEndingVersion = *o.EndingVersion
		}
		qEndingVersion := qrEndingVersion
		if qEndingVersion != "" {

			if err := r.SetQueryParam("endingVersion", qEndingVersion); err != nil {
				return err
			}
		}
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	// query param startingVersion
	qrStartingVersion := o.StartingVersion
	qStartingVersion := qrStartingVersion
	if qStartingVersion != "" {

		if err := r.SetQueryParam("startingVersion", qStartingVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRotationRequest binds the parameter ciphertext
func (o *RotationRequestParams) bindParamCiphertext(formats strfmt.Registry) []string {
	ciphertextIR := o.Ciphertext

	var ciphertextIC []string
	for _, ciphertextIIR := range ciphertextIR { // explode []uint8

		ciphertextIIV := swag.FormatUint8(ciphertextIIR) // uint8 as string
		ciphertextIC = append(ciphertextIC, ciphertextIIV)
	}

	// items.CollectionFormat: ""
	ciphertextIS := swag.JoinByFormat(ciphertextIC, "")

	return ciphertextIS
}
