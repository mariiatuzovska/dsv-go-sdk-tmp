// Code generated by go-swagger; DO NOT EDIT.

package audit

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

// NewDownloadAuditParams creates a new DownloadAuditParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadAuditParams() *DownloadAuditParams {
	return &DownloadAuditParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadAuditParamsWithTimeout creates a new DownloadAuditParams object
// with the ability to set a timeout on a request.
func NewDownloadAuditParamsWithTimeout(timeout time.Duration) *DownloadAuditParams {
	return &DownloadAuditParams{
		timeout: timeout,
	}
}

// NewDownloadAuditParamsWithContext creates a new DownloadAuditParams object
// with the ability to set a context for a request.
func NewDownloadAuditParamsWithContext(ctx context.Context) *DownloadAuditParams {
	return &DownloadAuditParams{
		Context: ctx,
	}
}

// NewDownloadAuditParamsWithHTTPClient creates a new DownloadAuditParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadAuditParamsWithHTTPClient(client *http.Client) *DownloadAuditParams {
	return &DownloadAuditParams{
		HTTPClient: client,
	}
}

/* DownloadAuditParams contains all the parameters to send to the API endpoint
   for the download audit operation.

   Typically these are written to a http.Request.
*/
type DownloadAuditParams struct {

	/* EndDate.

	   The end date to find audits to
	*/
	EndDate string

	/* StartDate.

	   The start date to find audits from
	*/
	StartDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download audit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadAuditParams) WithDefaults() *DownloadAuditParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download audit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadAuditParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download audit params
func (o *DownloadAuditParams) WithTimeout(timeout time.Duration) *DownloadAuditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download audit params
func (o *DownloadAuditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download audit params
func (o *DownloadAuditParams) WithContext(ctx context.Context) *DownloadAuditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download audit params
func (o *DownloadAuditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download audit params
func (o *DownloadAuditParams) WithHTTPClient(client *http.Client) *DownloadAuditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download audit params
func (o *DownloadAuditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the download audit params
func (o *DownloadAuditParams) WithEndDate(endDate string) *DownloadAuditParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the download audit params
func (o *DownloadAuditParams) SetEndDate(endDate string) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the download audit params
func (o *DownloadAuditParams) WithStartDate(startDate string) *DownloadAuditParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the download audit params
func (o *DownloadAuditParams) SetStartDate(startDate string) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadAuditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endDate
	qrEndDate := o.EndDate
	qEndDate := qrEndDate
	if qEndDate != "" {

		if err := r.SetQueryParam("endDate", qEndDate); err != nil {
			return err
		}
	}

	// query param startDate
	qrStartDate := o.StartDate
	qStartDate := qrStartDate
	if qStartDate != "" {

		if err := r.SetQueryParam("startDate", qStartDate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
