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

// NewSearchSecretsParams creates a new SearchSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchSecretsParams() *SearchSecretsParams {
	return &SearchSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchSecretsParamsWithTimeout creates a new SearchSecretsParams object
// with the ability to set a timeout on a request.
func NewSearchSecretsParamsWithTimeout(timeout time.Duration) *SearchSecretsParams {
	return &SearchSecretsParams{
		timeout: timeout,
	}
}

// NewSearchSecretsParamsWithContext creates a new SearchSecretsParams object
// with the ability to set a context for a request.
func NewSearchSecretsParamsWithContext(ctx context.Context) *SearchSecretsParams {
	return &SearchSecretsParams{
		Context: ctx,
	}
}

// NewSearchSecretsParamsWithHTTPClient creates a new SearchSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchSecretsParamsWithHTTPClient(client *http.Client) *SearchSecretsParams {
	return &SearchSecretsParams{
		HTTPClient: client,
	}
}

/* SearchSecretsParams contains all the parameters to send to the API endpoint
   for the search secrets operation.

   Typically these are written to a http.Request.
*/
type SearchSecretsParams struct {

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

// WithDefaults hydrates default values in the search secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSecretsParams) WithDefaults() *SearchSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search secrets params
func (o *SearchSecretsParams) WithTimeout(timeout time.Duration) *SearchSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search secrets params
func (o *SearchSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search secrets params
func (o *SearchSecretsParams) WithContext(ctx context.Context) *SearchSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search secrets params
func (o *SearchSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search secrets params
func (o *SearchSecretsParams) WithHTTPClient(client *http.Client) *SearchSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search secrets params
func (o *SearchSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the search secrets params
func (o *SearchSecretsParams) WithCursor(cursor *string) *SearchSecretsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the search secrets params
func (o *SearchSecretsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the search secrets params
func (o *SearchSecretsParams) WithLimit(limit *int64) *SearchSecretsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search secrets params
func (o *SearchSecretsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithComparison adds the searchComparison to the search secrets params
func (o *SearchSecretsParams) WithComparison(searchComparison *string) *SearchSecretsParams {
	o.SetComparison(searchComparison)
	return o
}

// SetComparison adds the searchComparison to the search secrets params
func (o *SearchSecretsParams) SetComparison(searchComparison *string) {
	o.Comparison = searchComparison
}

// WithField adds the searchField to the search secrets params
func (o *SearchSecretsParams) WithField(searchField *string) *SearchSecretsParams {
	o.SetField(searchField)
	return o
}

// SetField adds the searchField to the search secrets params
func (o *SearchSecretsParams) SetField(searchField *string) {
	o.Field = searchField
}

// WithLink adds the searchLinks to the search secrets params
func (o *SearchSecretsParams) WithLink(searchLinks *bool) *SearchSecretsParams {
	o.SetLink(searchLinks)
	return o
}

// SetLink adds the searchLinks to the search secrets params
func (o *SearchSecretsParams) SetLink(searchLinks *bool) {
	o.Link = searchLinks
}

// WithUnderlyingType adds the searchType to the search secrets params
func (o *SearchSecretsParams) WithUnderlyingType(searchType *string) *SearchSecretsParams {
	o.SetUnderlyingType(searchType)
	return o
}

// SetUnderlyingType adds the searchType to the search secrets params
func (o *SearchSecretsParams) SetUnderlyingType(searchType *string) {
	o.UnderlyingType = searchType
}

// WithSearchTerm adds the searchTerm to the search secrets params
func (o *SearchSecretsParams) WithSearchTerm(searchTerm *string) *SearchSecretsParams {
	o.SetSearchTerm(searchTerm)
	return o
}

// SetSearchTerm adds the searchTerm to the search secrets params
func (o *SearchSecretsParams) SetSearchTerm(searchTerm *string) {
	o.SearchTerm = searchTerm
}

// WithSort adds the sort to the search secrets params
func (o *SearchSecretsParams) WithSort(sort *string) *SearchSecretsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search secrets params
func (o *SearchSecretsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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