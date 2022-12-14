// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// GetSecretByVersionReader is a Reader for the GetSecretByVersion structure.
type GetSecretByVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecretByVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecretByVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSecretByVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSecretByVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSecretByVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSecretByVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSecretByVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSecretByVersionOK creates a GetSecretByVersionOK with default headers values
func NewGetSecretByVersionOK() *GetSecretByVersionOK {
	return &GetSecretByVersionOK{}
}

/* GetSecretByVersionOK describes a response with status code 200, with default header values.

no error
*/
type GetSecretByVersionOK struct {
	Payload *models.SecretVersionResponse
}

// IsSuccess returns true when this get secret by version o k response has a 2xx status code
func (o *GetSecretByVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get secret by version o k response has a 3xx status code
func (o *GetSecretByVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get secret by version o k response has a 4xx status code
func (o *GetSecretByVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get secret by version o k response has a 5xx status code
func (o *GetSecretByVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get secret by version o k response a status code equal to that given
func (o *GetSecretByVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSecretByVersionOK) Error() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionOK  %+v", 200, o.Payload)
}

func (o *GetSecretByVersionOK) String() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionOK  %+v", 200, o.Payload)
}

func (o *GetSecretByVersionOK) GetPayload() *models.SecretVersionResponse {
	return o.Payload
}

func (o *GetSecretByVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretByVersionBadRequest creates a GetSecretByVersionBadRequest with default headers values
func NewGetSecretByVersionBadRequest() *GetSecretByVersionBadRequest {
	return &GetSecretByVersionBadRequest{}
}

/* GetSecretByVersionBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetSecretByVersionBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get secret by version bad request response has a 2xx status code
func (o *GetSecretByVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get secret by version bad request response has a 3xx status code
func (o *GetSecretByVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get secret by version bad request response has a 4xx status code
func (o *GetSecretByVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get secret by version bad request response has a 5xx status code
func (o *GetSecretByVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get secret by version bad request response a status code equal to that given
func (o *GetSecretByVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSecretByVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetSecretByVersionBadRequest) String() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetSecretByVersionBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetSecretByVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretByVersionUnauthorized creates a GetSecretByVersionUnauthorized with default headers values
func NewGetSecretByVersionUnauthorized() *GetSecretByVersionUnauthorized {
	return &GetSecretByVersionUnauthorized{}
}

/* GetSecretByVersionUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetSecretByVersionUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get secret by version unauthorized response has a 2xx status code
func (o *GetSecretByVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get secret by version unauthorized response has a 3xx status code
func (o *GetSecretByVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get secret by version unauthorized response has a 4xx status code
func (o *GetSecretByVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get secret by version unauthorized response has a 5xx status code
func (o *GetSecretByVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get secret by version unauthorized response a status code equal to that given
func (o *GetSecretByVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSecretByVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSecretByVersionUnauthorized) String() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSecretByVersionUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetSecretByVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretByVersionForbidden creates a GetSecretByVersionForbidden with default headers values
func NewGetSecretByVersionForbidden() *GetSecretByVersionForbidden {
	return &GetSecretByVersionForbidden{}
}

/* GetSecretByVersionForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetSecretByVersionForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get secret by version forbidden response has a 2xx status code
func (o *GetSecretByVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get secret by version forbidden response has a 3xx status code
func (o *GetSecretByVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get secret by version forbidden response has a 4xx status code
func (o *GetSecretByVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get secret by version forbidden response has a 5xx status code
func (o *GetSecretByVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get secret by version forbidden response a status code equal to that given
func (o *GetSecretByVersionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSecretByVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetSecretByVersionForbidden) String() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetSecretByVersionForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetSecretByVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretByVersionNotFound creates a GetSecretByVersionNotFound with default headers values
func NewGetSecretByVersionNotFound() *GetSecretByVersionNotFound {
	return &GetSecretByVersionNotFound{}
}

/* GetSecretByVersionNotFound describes a response with status code 404, with default header values.

not found
*/
type GetSecretByVersionNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get secret by version not found response has a 2xx status code
func (o *GetSecretByVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get secret by version not found response has a 3xx status code
func (o *GetSecretByVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get secret by version not found response has a 4xx status code
func (o *GetSecretByVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get secret by version not found response has a 5xx status code
func (o *GetSecretByVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get secret by version not found response a status code equal to that given
func (o *GetSecretByVersionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSecretByVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetSecretByVersionNotFound) String() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetSecretByVersionNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetSecretByVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecretByVersionInternalServerError creates a GetSecretByVersionInternalServerError with default headers values
func NewGetSecretByVersionInternalServerError() *GetSecretByVersionInternalServerError {
	return &GetSecretByVersionInternalServerError{}
}

/* GetSecretByVersionInternalServerError describes a response with status code 500, with default header values.

server error
*/
type GetSecretByVersionInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get secret by version internal server error response has a 2xx status code
func (o *GetSecretByVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get secret by version internal server error response has a 3xx status code
func (o *GetSecretByVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get secret by version internal server error response has a 4xx status code
func (o *GetSecretByVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get secret by version internal server error response has a 5xx status code
func (o *GetSecretByVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get secret by version internal server error response a status code equal to that given
func (o *GetSecretByVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSecretByVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecretByVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /secrets/{path}/version/{version}][%d] getSecretByVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecretByVersionInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetSecretByVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
