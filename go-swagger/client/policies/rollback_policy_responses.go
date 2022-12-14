// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// RollbackPolicyReader is a Reader for the RollbackPolicy structure.
type RollbackPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RollbackPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRollbackPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRollbackPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRollbackPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRollbackPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRollbackPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRollbackPolicyOK creates a RollbackPolicyOK with default headers values
func NewRollbackPolicyOK() *RollbackPolicyOK {
	return &RollbackPolicyOK{}
}

/* RollbackPolicyOK describes a response with status code 200, with default header values.

no error
*/
type RollbackPolicyOK struct {
	Payload *models.Policy
}

// IsSuccess returns true when this rollback policy o k response has a 2xx status code
func (o *RollbackPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rollback policy o k response has a 3xx status code
func (o *RollbackPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rollback policy o k response has a 4xx status code
func (o *RollbackPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rollback policy o k response has a 5xx status code
func (o *RollbackPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rollback policy o k response a status code equal to that given
func (o *RollbackPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *RollbackPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyOK  %+v", 200, o.Payload)
}

func (o *RollbackPolicyOK) String() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyOK  %+v", 200, o.Payload)
}

func (o *RollbackPolicyOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *RollbackPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackPolicyBadRequest creates a RollbackPolicyBadRequest with default headers values
func NewRollbackPolicyBadRequest() *RollbackPolicyBadRequest {
	return &RollbackPolicyBadRequest{}
}

/* RollbackPolicyBadRequest describes a response with status code 400, with default header values.

bad request
*/
type RollbackPolicyBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this rollback policy bad request response has a 2xx status code
func (o *RollbackPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rollback policy bad request response has a 3xx status code
func (o *RollbackPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rollback policy bad request response has a 4xx status code
func (o *RollbackPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rollback policy bad request response has a 5xx status code
func (o *RollbackPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rollback policy bad request response a status code equal to that given
func (o *RollbackPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RollbackPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *RollbackPolicyBadRequest) String() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *RollbackPolicyBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *RollbackPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackPolicyUnauthorized creates a RollbackPolicyUnauthorized with default headers values
func NewRollbackPolicyUnauthorized() *RollbackPolicyUnauthorized {
	return &RollbackPolicyUnauthorized{}
}

/* RollbackPolicyUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type RollbackPolicyUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this rollback policy unauthorized response has a 2xx status code
func (o *RollbackPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rollback policy unauthorized response has a 3xx status code
func (o *RollbackPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rollback policy unauthorized response has a 4xx status code
func (o *RollbackPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this rollback policy unauthorized response has a 5xx status code
func (o *RollbackPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this rollback policy unauthorized response a status code equal to that given
func (o *RollbackPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RollbackPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *RollbackPolicyUnauthorized) String() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *RollbackPolicyUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *RollbackPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackPolicyForbidden creates a RollbackPolicyForbidden with default headers values
func NewRollbackPolicyForbidden() *RollbackPolicyForbidden {
	return &RollbackPolicyForbidden{}
}

/* RollbackPolicyForbidden describes a response with status code 403, with default header values.

forbidden
*/
type RollbackPolicyForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this rollback policy forbidden response has a 2xx status code
func (o *RollbackPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rollback policy forbidden response has a 3xx status code
func (o *RollbackPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rollback policy forbidden response has a 4xx status code
func (o *RollbackPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this rollback policy forbidden response has a 5xx status code
func (o *RollbackPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this rollback policy forbidden response a status code equal to that given
func (o *RollbackPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RollbackPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyForbidden  %+v", 403, o.Payload)
}

func (o *RollbackPolicyForbidden) String() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyForbidden  %+v", 403, o.Payload)
}

func (o *RollbackPolicyForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *RollbackPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackPolicyInternalServerError creates a RollbackPolicyInternalServerError with default headers values
func NewRollbackPolicyInternalServerError() *RollbackPolicyInternalServerError {
	return &RollbackPolicyInternalServerError{}
}

/* RollbackPolicyInternalServerError describes a response with status code 500, with default header values.

server error
*/
type RollbackPolicyInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this rollback policy internal server error response has a 2xx status code
func (o *RollbackPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rollback policy internal server error response has a 3xx status code
func (o *RollbackPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rollback policy internal server error response has a 4xx status code
func (o *RollbackPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rollback policy internal server error response has a 5xx status code
func (o *RollbackPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rollback policy internal server error response a status code equal to that given
func (o *RollbackPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RollbackPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *RollbackPolicyInternalServerError) String() string {
	return fmt.Sprintf("[PUT /config/policies/{path}/rollback/{version}][%d] rollbackPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *RollbackPolicyInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *RollbackPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
