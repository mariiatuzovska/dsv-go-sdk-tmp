// Code generated by go-swagger; DO NOT EDIT.

package s_i_e_m

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

// NewSiemUpdateParams creates a new SiemUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSiemUpdateParams() *SiemUpdateParams {
	return &SiemUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSiemUpdateParamsWithTimeout creates a new SiemUpdateParams object
// with the ability to set a timeout on a request.
func NewSiemUpdateParamsWithTimeout(timeout time.Duration) *SiemUpdateParams {
	return &SiemUpdateParams{
		timeout: timeout,
	}
}

// NewSiemUpdateParamsWithContext creates a new SiemUpdateParams object
// with the ability to set a context for a request.
func NewSiemUpdateParamsWithContext(ctx context.Context) *SiemUpdateParams {
	return &SiemUpdateParams{
		Context: ctx,
	}
}

// NewSiemUpdateParamsWithHTTPClient creates a new SiemUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSiemUpdateParamsWithHTTPClient(client *http.Client) *SiemUpdateParams {
	return &SiemUpdateParams{
		HTTPClient: client,
	}
}

/* SiemUpdateParams contains all the parameters to send to the API endpoint
   for the siem update operation.

   Typically these are written to a http.Request.
*/
type SiemUpdateParams struct {

	// Body.
	Body *models.SiemCreateUpdateRequestModel

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the siem update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SiemUpdateParams) WithDefaults() *SiemUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the siem update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SiemUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the siem update params
func (o *SiemUpdateParams) WithTimeout(timeout time.Duration) *SiemUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the siem update params
func (o *SiemUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the siem update params
func (o *SiemUpdateParams) WithContext(ctx context.Context) *SiemUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the siem update params
func (o *SiemUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the siem update params
func (o *SiemUpdateParams) WithHTTPClient(client *http.Client) *SiemUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the siem update params
func (o *SiemUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the siem update params
func (o *SiemUpdateParams) WithBody(body *models.SiemCreateUpdateRequestModel) *SiemUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the siem update params
func (o *SiemUpdateParams) SetBody(body *models.SiemCreateUpdateRequestModel) {
	o.Body = body
}

// WithName adds the name to the siem update params
func (o *SiemUpdateParams) WithName(name string) *SiemUpdateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the siem update params
func (o *SiemUpdateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SiemUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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