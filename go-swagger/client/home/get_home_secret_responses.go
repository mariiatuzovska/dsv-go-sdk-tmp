// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// GetHomeSecretReader is a Reader for the GetHomeSecret structure.
type GetHomeSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHomeSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHomeSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHomeSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetHomeSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHomeSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHomeSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHomeSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHomeSecretOK creates a GetHomeSecretOK with default headers values
func NewGetHomeSecretOK() *GetHomeSecretOK {
	return &GetHomeSecretOK{}
}

/* GetHomeSecretOK describes a response with status code 200, with default header values.

no error
*/
type GetHomeSecretOK struct {
	Payload *models.ResponseModelFull
}

// IsSuccess returns true when this get home secret o k response has a 2xx status code
func (o *GetHomeSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get home secret o k response has a 3xx status code
func (o *GetHomeSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home secret o k response has a 4xx status code
func (o *GetHomeSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get home secret o k response has a 5xx status code
func (o *GetHomeSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get home secret o k response a status code equal to that given
func (o *GetHomeSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHomeSecretOK) Error() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretOK  %+v", 200, o.Payload)
}

func (o *GetHomeSecretOK) String() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretOK  %+v", 200, o.Payload)
}

func (o *GetHomeSecretOK) GetPayload() *models.ResponseModelFull {
	return o.Payload
}

func (o *GetHomeSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseModelFull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeSecretBadRequest creates a GetHomeSecretBadRequest with default headers values
func NewGetHomeSecretBadRequest() *GetHomeSecretBadRequest {
	return &GetHomeSecretBadRequest{}
}

/* GetHomeSecretBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetHomeSecretBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get home secret bad request response has a 2xx status code
func (o *GetHomeSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home secret bad request response has a 3xx status code
func (o *GetHomeSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home secret bad request response has a 4xx status code
func (o *GetHomeSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get home secret bad request response has a 5xx status code
func (o *GetHomeSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get home secret bad request response a status code equal to that given
func (o *GetHomeSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetHomeSecretBadRequest) Error() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretBadRequest  %+v", 400, o.Payload)
}

func (o *GetHomeSecretBadRequest) String() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretBadRequest  %+v", 400, o.Payload)
}

func (o *GetHomeSecretBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetHomeSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeSecretUnauthorized creates a GetHomeSecretUnauthorized with default headers values
func NewGetHomeSecretUnauthorized() *GetHomeSecretUnauthorized {
	return &GetHomeSecretUnauthorized{}
}

/* GetHomeSecretUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetHomeSecretUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get home secret unauthorized response has a 2xx status code
func (o *GetHomeSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home secret unauthorized response has a 3xx status code
func (o *GetHomeSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home secret unauthorized response has a 4xx status code
func (o *GetHomeSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get home secret unauthorized response has a 5xx status code
func (o *GetHomeSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get home secret unauthorized response a status code equal to that given
func (o *GetHomeSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetHomeSecretUnauthorized) Error() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHomeSecretUnauthorized) String() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHomeSecretUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetHomeSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeSecretForbidden creates a GetHomeSecretForbidden with default headers values
func NewGetHomeSecretForbidden() *GetHomeSecretForbidden {
	return &GetHomeSecretForbidden{}
}

/* GetHomeSecretForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetHomeSecretForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get home secret forbidden response has a 2xx status code
func (o *GetHomeSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home secret forbidden response has a 3xx status code
func (o *GetHomeSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home secret forbidden response has a 4xx status code
func (o *GetHomeSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get home secret forbidden response has a 5xx status code
func (o *GetHomeSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get home secret forbidden response a status code equal to that given
func (o *GetHomeSecretForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetHomeSecretForbidden) Error() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretForbidden  %+v", 403, o.Payload)
}

func (o *GetHomeSecretForbidden) String() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretForbidden  %+v", 403, o.Payload)
}

func (o *GetHomeSecretForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetHomeSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeSecretNotFound creates a GetHomeSecretNotFound with default headers values
func NewGetHomeSecretNotFound() *GetHomeSecretNotFound {
	return &GetHomeSecretNotFound{}
}

/* GetHomeSecretNotFound describes a response with status code 404, with default header values.

not found
*/
type GetHomeSecretNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get home secret not found response has a 2xx status code
func (o *GetHomeSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home secret not found response has a 3xx status code
func (o *GetHomeSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home secret not found response has a 4xx status code
func (o *GetHomeSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get home secret not found response has a 5xx status code
func (o *GetHomeSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get home secret not found response a status code equal to that given
func (o *GetHomeSecretNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetHomeSecretNotFound) Error() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretNotFound  %+v", 404, o.Payload)
}

func (o *GetHomeSecretNotFound) String() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretNotFound  %+v", 404, o.Payload)
}

func (o *GetHomeSecretNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetHomeSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeSecretInternalServerError creates a GetHomeSecretInternalServerError with default headers values
func NewGetHomeSecretInternalServerError() *GetHomeSecretInternalServerError {
	return &GetHomeSecretInternalServerError{}
}

/* GetHomeSecretInternalServerError describes a response with status code 500, with default header values.

server error
*/
type GetHomeSecretInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get home secret internal server error response has a 2xx status code
func (o *GetHomeSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home secret internal server error response has a 3xx status code
func (o *GetHomeSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home secret internal server error response has a 4xx status code
func (o *GetHomeSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get home secret internal server error response has a 5xx status code
func (o *GetHomeSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get home secret internal server error response a status code equal to that given
func (o *GetHomeSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHomeSecretInternalServerError) Error() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHomeSecretInternalServerError) String() string {
	return fmt.Sprintf("[GET /home/{principalName}/{path}][%d] getHomeSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHomeSecretInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetHomeSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
