// Code generated by go-swagger; DO NOT EDIT.

package home

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

// NewSearchHomeSecretsParams creates a new SearchHomeSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchHomeSecretsParams() *SearchHomeSecretsParams {
	return &SearchHomeSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchHomeSecretsParamsWithTimeout creates a new SearchHomeSecretsParams object
// with the ability to set a timeout on a request.
func NewSearchHomeSecretsParamsWithTimeout(timeout time.Duration) *SearchHomeSecretsParams {
	return &SearchHomeSecretsParams{
		timeout: timeout,
	}
}

// NewSearchHomeSecretsParamsWithContext creates a new SearchHomeSecretsParams object
// with the ability to set a context for a request.
func NewSearchHomeSecretsParamsWithContext(ctx context.Context) *SearchHomeSecretsParams {
	return &SearchHomeSecretsParams{
		Context: ctx,
	}
}

// NewSearchHomeSecretsParamsWithHTTPClient creates a new SearchHomeSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchHomeSecretsParamsWithHTTPClient(client *http.Client) *SearchHomeSecretsParams {
	return &SearchHomeSecretsParams{
		HTTPClient: client,
	}
}

/* SearchHomeSecretsParams contains all the parameters to send to the API endpoint
   for the search home secrets operation.

   Typically these are written to a http.Request.
*/
type SearchHomeSecretsParams struct {

	/* Cursor.

	   Cursor to next batch of results
	*/
	Cursor *string

	/* Limit.

	   Limit for the number of results per page (cursor)

	   Format: int64
	*/
	Limit *int64

	/* SearchComparison.

	   Comparison type (equal, contains, begins_with) for advanced searching
	*/
	Comparison *string

	/* SearchField.

	   Secret field for advanced searching
	*/
	Field *string

	/* SearchLinks.

	   Whether to search for secrets that link to the path in the search term
	*/
	Link *bool

	/* SearchType.

	   Attribute type (string, number) to search on
	*/
	UnderlyingType *string

	/* SearchTerm.

	   Partial search term for search by path
	*/
	SearchTerm *string

	/* Sort.

	   Sort results ascending (asc) or descending (desc) order by lastModified attribute on field search. Default is desc
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search home secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchHomeSecretsParams) WithDefaults() *SearchHomeSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search home secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchHomeSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search home secrets params
func (o *SearchHomeSecretsParams) WithTimeout(timeout time.Duration) *SearchHomeSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search home secrets params
func (o *SearchHomeSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search home secrets params
func (o *SearchHomeSecretsParams) WithContext(ctx context.Context) *SearchHomeSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search home secrets params
func (o *SearchHomeSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search home secrets params
func (o *SearchHomeSecretsParams) WithHTTPClient(client *http.Client) *SearchHomeSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search home secrets params
func (o *SearchHomeSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the search home secrets params
func (o *SearchHomeSecretsParams) WithCursor(cursor *string) *SearchHomeSecretsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the search home secrets params
func (o *SearchHomeSecretsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the search home secrets params
func (o *SearchHomeSecretsParams) WithLimit(limit *int64) *SearchHomeSecretsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search home secrets params
func (o *SearchHomeSecretsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithComparison adds the searchComparison to the search home secrets params
func (o *SearchHomeSecretsParams) WithComparison(searchComparison *string) *SearchHomeSecretsParams {
	o.SetComparison(searchComparison)
	return o
}

// SetComparison adds the searchComparison to the search home secrets params
func (o *SearchHomeSecretsParams) SetComparison(searchComparison *string) {
	o.Comparison = searchComparison
}

// WithField adds the searchField to the search home secrets params
func (o *SearchHomeSecretsParams) WithField(searchField *string) *SearchHomeSecretsParams {
	o.SetField(searchField)
	return o
}

// SetField adds the searchField to the search home secrets params
func (o *SearchHomeSecretsParams) SetField(searchField *string) {
	o.Field = searchField
}

// WithLink adds the searchLinks to the search home secrets params
func (o *SearchHomeSecretsParams) WithLink(searchLinks *bool) *SearchHomeSecretsParams {
	o.SetLink(searchLinks)
	return o
}

// SetLink adds the searchLinks to the search home secrets params
func (o *SearchHomeSecretsParams) SetLink(searchLinks *bool) {
	o.Link = searchLinks
}

// WithUnderlyingType adds the searchType to the search home secrets params
func (o *SearchHomeSecretsParams) WithUnderlyingType(searchType *string) *SearchHomeSecretsParams {
	o.SetUnderlyingType(searchType)
	return o
}

// SetUnderlyingType adds the searchType to the search home secrets params
func (o *SearchHomeSecretsParams) SetUnderlyingType(searchType *string) {
	o.UnderlyingType = searchType
}

// WithSearchTerm adds the searchTerm to the search home secrets params
func (o *SearchHomeSecretsParams) WithSearchTerm(searchTerm *string) *SearchHomeSecretsParams {
	o.SetSearchTerm(searchTerm)
	return o
}

// SetSearchTerm adds the searchTerm to the search home secrets params
func (o *SearchHomeSecretsParams) SetSearchTerm(searchTerm *string) {
	o.SearchTerm = searchTerm
}

// WithSort adds the sort to the search home secrets params
func (o *SearchHomeSecretsParams) WithSort(sort *string) *SearchHomeSecretsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search home secrets params
func (o *SearchHomeSecretsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchHomeSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}
	}

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

	if o.Comparison != nil {

		// query param search.comparison
		var qrSearchComparison string

		if o.Comparison != nil {
			qrSearchComparison = *o.Comparison
		}
		qSearchComparison := qrSearchComparison
		if qSearchComparison != "" {

			if err := r.SetQueryParam("search.comparison", qSearchComparison); err != nil {
				return err
			}
		}
	}

	if o.Field != nil {

		// query param search.field
		var qrSearchField string

		if o.Field != nil {
			qrSearchField = *o.Field
		}
		qSearchField := qrSearchField
		if qSearchField != "" {

			if err := r.SetQueryParam("search.field", qSearchField); err != nil {
				return err
			}
		}
	}

	if o.Link != nil {

		// query param search.links
		var qrSearchLinks bool

		if o.Link != nil {
			qrSearchLinks = *o.Link
		}
		qSearchLinks := swag.FormatBool(qrSearchLinks)
		if qSearchLinks != "" {

			if err := r.SetQueryParam("search.links", qSearchLinks); err != nil {
				return err
			}
		}
	}

	if o.UnderlyingType != nil {

		// query param search.type
		var qrSearchType string

		if o.UnderlyingType != nil {
			qrSearchType = *o.UnderlyingType
		}
		qSearchType := qrSearchType
		if qSearchType != "" {

			if err := r.SetQueryParam("search.type", qSearchType); err != nil {
				return err
			}
		}
	}

	if o.SearchTerm != nil {

		// query param searchTerm
		var qrSearchTerm string

		if o.SearchTerm != nil {
			qrSearchTerm = *o.SearchTerm
		}
		qSearchTerm := qrSearchTerm
		if qSearchTerm != "" {

			if err := r.SetQueryParam("searchTerm", qSearchTerm); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}