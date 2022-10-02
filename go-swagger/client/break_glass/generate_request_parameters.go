// Code generated by go-swagger; DO NOT EDIT.

package break_glass

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

// NewGenerateRequestParams creates a new GenerateRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateRequestParams() *GenerateRequestParams {
	return &GenerateRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateRequestParamsWithTimeout creates a new GenerateRequestParams object
// with the ability to set a timeout on a request.
func NewGenerateRequestParamsWithTimeout(timeout time.Duration) *GenerateRequestParams {
	return &GenerateRequestParams{
		timeout: timeout,
	}
}

// NewGenerateRequestParamsWithContext creates a new GenerateRequestParams object
// with the ability to set a context for a request.
func NewGenerateRequestParamsWithContext(ctx context.Context) *GenerateRequestParams {
	return &GenerateRequestParams{
		Context: ctx,
	}
}

// NewGenerateRequestParamsWithHTTPClient creates a new GenerateRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateRequestParamsWithHTTPClient(client *http.Client) *GenerateRequestParams {
	return &GenerateRequestParams{
		HTTPClient: client,
	}
}

/* GenerateRequestParams contains all the parameters to send to the API endpoint
   for the generate request operation.

   Typically these are written to a http.Request.
*/
type GenerateRequestParams struct {

	// MinNumberOfShares.
	//
	// Format: int64
	MinNumberOfShares int64

	// NewAdmins.
	NewAdmins []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateRequestParams) WithDefaults() *GenerateRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate request params
func (o *GenerateRequestParams) WithTimeout(timeout time.Duration) *GenerateRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate request params
func (o *GenerateRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate request params
func (o *GenerateRequestParams) WithContext(ctx context.Context) *GenerateRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate request params
func (o *GenerateRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate request params
func (o *GenerateRequestParams) WithHTTPClient(client *http.Client) *GenerateRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate request params
func (o *GenerateRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMinNumberOfShares adds the minNumberOfShares to the generate request params
func (o *GenerateRequestParams) WithMinNumberOfShares(minNumberOfShares int64) *GenerateRequestParams {
	o.SetMinNumberOfShares(minNumberOfShares)
	return o
}

// SetMinNumberOfShares adds the minNumberOfShares to the generate request params
func (o *GenerateRequestParams) SetMinNumberOfShares(minNumberOfShares int64) {
	o.MinNumberOfShares = minNumberOfShares
}

// WithNewAdmins adds the newAdmins to the generate request params
func (o *GenerateRequestParams) WithNewAdmins(newAdmins []string) *GenerateRequestParams {
	o.SetNewAdmins(newAdmins)
	return o
}

// SetNewAdmins adds the newAdmins to the generate request params
func (o *GenerateRequestParams) SetNewAdmins(newAdmins []string) {
	o.NewAdmins = newAdmins
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param minNumberOfShares
	qrMinNumberOfShares := o.MinNumberOfShares
	qMinNumberOfShares := swag.FormatInt64(qrMinNumberOfShares)
	if qMinNumberOfShares != "" {

		if err := r.SetQueryParam("minNumberOfShares", qMinNumberOfShares); err != nil {
			return err
		}
	}

	if o.NewAdmins != nil {

		// binding items for newAdmins
		joinedNewAdmins := o.bindParamNewAdmins(reg)

		// query array param newAdmins
		if err := r.SetQueryParam("newAdmins", joinedNewAdmins...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGenerateRequest binds the parameter newAdmins
func (o *GenerateRequestParams) bindParamNewAdmins(formats strfmt.Registry) []string {
	newAdminsIR := o.NewAdmins

	var newAdminsIC []string
	for _, newAdminsIIR := range newAdminsIR { // explode []string

		newAdminsIIV := newAdminsIIR // string as string
		newAdminsIC = append(newAdminsIC, newAdminsIIV)
	}

	// items.CollectionFormat: ""
	newAdminsIS := swag.JoinByFormat(newAdminsIC, "")

	return newAdminsIS
}