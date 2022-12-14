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

// DeleteGroupReader is a Reader for the DeleteGroup structure.
type DeleteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGroupOK creates a DeleteGroupOK with default headers values
func NewDeleteGroupOK() *DeleteGroupOK {
	return &DeleteGroupOK{}
}

/* DeleteGroupOK describes a response with status code 200, with default header values.

no error
*/
type DeleteGroupOK struct {
	Payload *models.MessageResponse
}

// IsSuccess returns true when this delete group o k response has a 2xx status code
func (o *DeleteGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group o k response has a 3xx status code
func (o *DeleteGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group o k response has a 4xx status code
func (o *DeleteGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group o k response has a 5xx status code
func (o *DeleteGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group o k response a status code equal to that given
func (o *DeleteGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteGroupOK) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteGroupOK) GetPayload() *models.MessageResponse {
	return o.Payload
}

func (o *DeleteGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupNoContent creates a DeleteGroupNoContent with default headers values
func NewDeleteGroupNoContent() *DeleteGroupNoContent {
	return &DeleteGroupNoContent{}
}

/* DeleteGroupNoContent describes a response with status code 204, with default header values.

no error
*/
type DeleteGroupNoContent struct {
}

// IsSuccess returns true when this delete group no content response has a 2xx status code
func (o *DeleteGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group no content response has a 3xx status code
func (o *DeleteGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group no content response has a 4xx status code
func (o *DeleteGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group no content response has a 5xx status code
func (o *DeleteGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group no content response a status code equal to that given
func (o *DeleteGroupNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupNoContent ", 204)
}

func (o *DeleteGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupNoContent ", 204)
}

func (o *DeleteGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupBadRequest creates a DeleteGroupBadRequest with default headers values
func NewDeleteGroupBadRequest() *DeleteGroupBadRequest {
	return &DeleteGroupBadRequest{}
}

/* DeleteGroupBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DeleteGroupBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete group bad request response has a 2xx status code
func (o *DeleteGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group bad request response has a 3xx status code
func (o *DeleteGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group bad request response has a 4xx status code
func (o *DeleteGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group bad request response has a 5xx status code
func (o *DeleteGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group bad request response a status code equal to that given
func (o *DeleteGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGroupBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupUnauthorized creates a DeleteGroupUnauthorized with default headers values
func NewDeleteGroupUnauthorized() *DeleteGroupUnauthorized {
	return &DeleteGroupUnauthorized{}
}

/* DeleteGroupUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteGroupUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete group unauthorized response has a 2xx status code
func (o *DeleteGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group unauthorized response has a 3xx status code
func (o *DeleteGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group unauthorized response has a 4xx status code
func (o *DeleteGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group unauthorized response has a 5xx status code
func (o *DeleteGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group unauthorized response a status code equal to that given
func (o *DeleteGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGroupUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGroupUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupForbidden creates a DeleteGroupForbidden with default headers values
func NewDeleteGroupForbidden() *DeleteGroupForbidden {
	return &DeleteGroupForbidden{}
}

/* DeleteGroupForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteGroupForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete group forbidden response has a 2xx status code
func (o *DeleteGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group forbidden response has a 3xx status code
func (o *DeleteGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group forbidden response has a 4xx status code
func (o *DeleteGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete group forbidden response has a 5xx status code
func (o *DeleteGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group forbidden response a status code equal to that given
func (o *DeleteGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGroupForbidden) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteGroupForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupInternalServerError creates a DeleteGroupInternalServerError with default headers values
func NewDeleteGroupInternalServerError() *DeleteGroupInternalServerError {
	return &DeleteGroupInternalServerError{}
}

/* DeleteGroupInternalServerError describes a response with status code 500, with default header values.

server error
*/
type DeleteGroupInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete group internal server error response has a 2xx status code
func (o *DeleteGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete group internal server error response has a 3xx status code
func (o *DeleteGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group internal server error response has a 4xx status code
func (o *DeleteGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group internal server error response has a 5xx status code
func (o *DeleteGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete group internal server error response a status code equal to that given
func (o *DeleteGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGroupInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /groups/{name}][%d] deleteGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGroupInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
