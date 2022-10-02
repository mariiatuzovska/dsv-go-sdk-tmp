// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// DeleteMemberReader is a Reader for the DeleteMember structure.
type DeleteMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMemberNoContent creates a DeleteMemberNoContent with default headers values
func NewDeleteMemberNoContent() *DeleteMemberNoContent {
	return &DeleteMemberNoContent{}
}

/* DeleteMemberNoContent describes a response with status code 204, with default header values.

no error
*/
type DeleteMemberNoContent struct {
}

// IsSuccess returns true when this delete member no content response has a 2xx status code
func (o *DeleteMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete member no content response has a 3xx status code
func (o *DeleteMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete member no content response has a 4xx status code
func (o *DeleteMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete member no content response has a 5xx status code
func (o *DeleteMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete member no content response a status code equal to that given
func (o *DeleteMemberNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberNoContent ", 204)
}

func (o *DeleteMemberNoContent) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberNoContent ", 204)
}

func (o *DeleteMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMemberBadRequest creates a DeleteMemberBadRequest with default headers values
func NewDeleteMemberBadRequest() *DeleteMemberBadRequest {
	return &DeleteMemberBadRequest{}
}

/* DeleteMemberBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DeleteMemberBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete member bad request response has a 2xx status code
func (o *DeleteMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete member bad request response has a 3xx status code
func (o *DeleteMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete member bad request response has a 4xx status code
func (o *DeleteMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete member bad request response has a 5xx status code
func (o *DeleteMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete member bad request response a status code equal to that given
func (o *DeleteMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMemberBadRequest) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMemberBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMemberUnauthorized creates a DeleteMemberUnauthorized with default headers values
func NewDeleteMemberUnauthorized() *DeleteMemberUnauthorized {
	return &DeleteMemberUnauthorized{}
}

/* DeleteMemberUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteMemberUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete member unauthorized response has a 2xx status code
func (o *DeleteMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete member unauthorized response has a 3xx status code
func (o *DeleteMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete member unauthorized response has a 4xx status code
func (o *DeleteMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete member unauthorized response has a 5xx status code
func (o *DeleteMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete member unauthorized response a status code equal to that given
func (o *DeleteMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMemberUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMemberUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMemberForbidden creates a DeleteMemberForbidden with default headers values
func NewDeleteMemberForbidden() *DeleteMemberForbidden {
	return &DeleteMemberForbidden{}
}

/* DeleteMemberForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteMemberForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete member forbidden response has a 2xx status code
func (o *DeleteMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete member forbidden response has a 3xx status code
func (o *DeleteMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete member forbidden response has a 4xx status code
func (o *DeleteMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete member forbidden response has a 5xx status code
func (o *DeleteMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete member forbidden response a status code equal to that given
func (o *DeleteMemberForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMemberForbidden) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMemberForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMemberInternalServerError creates a DeleteMemberInternalServerError with default headers values
func NewDeleteMemberInternalServerError() *DeleteMemberInternalServerError {
	return &DeleteMemberInternalServerError{}
}

/* DeleteMemberInternalServerError describes a response with status code 500, with default header values.

server error
*/
type DeleteMemberInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete member internal server error response has a 2xx status code
func (o *DeleteMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete member internal server error response has a 3xx status code
func (o *DeleteMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete member internal server error response has a 4xx status code
func (o *DeleteMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete member internal server error response has a 5xx status code
func (o *DeleteMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete member internal server error response a status code equal to that given
func (o *DeleteMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMemberInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}/members][%d] deleteMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMemberInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}