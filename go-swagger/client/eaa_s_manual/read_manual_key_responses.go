// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_manual

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// ReadManualKeyReader is a Reader for the ReadManualKey structure.
type ReadManualKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadManualKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadManualKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadManualKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadManualKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadManualKeyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadManualKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadManualKeyOK creates a ReadManualKeyOK with default headers values
func NewReadManualKeyOK() *ReadManualKeyOK {
	return &ReadManualKeyOK{}
}

/* ReadManualKeyOK describes a response with status code 200, with default header values.

no error
*/
type ReadManualKeyOK struct {
	Payload *models.ResponseModelFull
}

// IsSuccess returns true when this read manual key o k response has a 2xx status code
func (o *ReadManualKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read manual key o k response has a 3xx status code
func (o *ReadManualKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read manual key o k response has a 4xx status code
func (o *ReadManualKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read manual key o k response has a 5xx status code
func (o *ReadManualKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read manual key o k response a status code equal to that given
func (o *ReadManualKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *ReadManualKeyOK) Error() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyOK  %+v", 200, o.Payload)
}

func (o *ReadManualKeyOK) String() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyOK  %+v", 200, o.Payload)
}

func (o *ReadManualKeyOK) GetPayload() *models.ResponseModelFull {
	return o.Payload
}

func (o *ReadManualKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseModelFull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadManualKeyBadRequest creates a ReadManualKeyBadRequest with default headers values
func NewReadManualKeyBadRequest() *ReadManualKeyBadRequest {
	return &ReadManualKeyBadRequest{}
}

/* ReadManualKeyBadRequest describes a response with status code 400, with default header values.

bad request
*/
type ReadManualKeyBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this read manual key bad request response has a 2xx status code
func (o *ReadManualKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read manual key bad request response has a 3xx status code
func (o *ReadManualKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read manual key bad request response has a 4xx status code
func (o *ReadManualKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this read manual key bad request response has a 5xx status code
func (o *ReadManualKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this read manual key bad request response a status code equal to that given
func (o *ReadManualKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ReadManualKeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyBadRequest  %+v", 400, o.Payload)
}

func (o *ReadManualKeyBadRequest) String() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyBadRequest  %+v", 400, o.Payload)
}

func (o *ReadManualKeyBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *ReadManualKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadManualKeyUnauthorized creates a ReadManualKeyUnauthorized with default headers values
func NewReadManualKeyUnauthorized() *ReadManualKeyUnauthorized {
	return &ReadManualKeyUnauthorized{}
}

/* ReadManualKeyUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ReadManualKeyUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this read manual key unauthorized response has a 2xx status code
func (o *ReadManualKeyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read manual key unauthorized response has a 3xx status code
func (o *ReadManualKeyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read manual key unauthorized response has a 4xx status code
func (o *ReadManualKeyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this read manual key unauthorized response has a 5xx status code
func (o *ReadManualKeyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this read manual key unauthorized response a status code equal to that given
func (o *ReadManualKeyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ReadManualKeyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadManualKeyUnauthorized) String() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadManualKeyUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *ReadManualKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadManualKeyForbidden creates a ReadManualKeyForbidden with default headers values
func NewReadManualKeyForbidden() *ReadManualKeyForbidden {
	return &ReadManualKeyForbidden{}
}

/* ReadManualKeyForbidden describes a response with status code 403, with default header values.

forbidden
*/
type ReadManualKeyForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this read manual key forbidden response has a 2xx status code
func (o *ReadManualKeyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read manual key forbidden response has a 3xx status code
func (o *ReadManualKeyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read manual key forbidden response has a 4xx status code
func (o *ReadManualKeyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read manual key forbidden response has a 5xx status code
func (o *ReadManualKeyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read manual key forbidden response a status code equal to that given
func (o *ReadManualKeyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ReadManualKeyForbidden) Error() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyForbidden  %+v", 403, o.Payload)
}

func (o *ReadManualKeyForbidden) String() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyForbidden  %+v", 403, o.Payload)
}

func (o *ReadManualKeyForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *ReadManualKeyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadManualKeyInternalServerError creates a ReadManualKeyInternalServerError with default headers values
func NewReadManualKeyInternalServerError() *ReadManualKeyInternalServerError {
	return &ReadManualKeyInternalServerError{}
}

/* ReadManualKeyInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ReadManualKeyInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this read manual key internal server error response has a 2xx status code
func (o *ReadManualKeyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read manual key internal server error response has a 3xx status code
func (o *ReadManualKeyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read manual key internal server error response has a 4xx status code
func (o *ReadManualKeyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read manual key internal server error response has a 5xx status code
func (o *ReadManualKeyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read manual key internal server error response a status code equal to that given
func (o *ReadManualKeyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ReadManualKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadManualKeyInternalServerError) String() string {
	return fmt.Sprintf("[GET /crypto/manual/key/{path}][%d] readManualKeyInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadManualKeyInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *ReadManualKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}