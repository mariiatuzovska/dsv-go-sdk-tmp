// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// PostConfigReader is a Reader for the PostConfig structure.
type PostConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConfigOK creates a PostConfigOK with default headers values
func NewPostConfigOK() *PostConfigOK {
	return &PostConfigOK{}
}

/* PostConfigOK describes a response with status code 200, with default header values.

no error
*/
type PostConfigOK struct {
	Payload *models.Document
}

// IsSuccess returns true when this post config o k response has a 2xx status code
func (o *PostConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post config o k response has a 3xx status code
func (o *PostConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config o k response has a 4xx status code
func (o *PostConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post config o k response has a 5xx status code
func (o *PostConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post config o k response a status code equal to that given
func (o *PostConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostConfigOK) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigOK  %+v", 200, o.Payload)
}

func (o *PostConfigOK) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigOK  %+v", 200, o.Payload)
}

func (o *PostConfigOK) GetPayload() *models.Document {
	return o.Payload
}

func (o *PostConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConfigCreated creates a PostConfigCreated with default headers values
func NewPostConfigCreated() *PostConfigCreated {
	return &PostConfigCreated{}
}

/* PostConfigCreated describes a response with status code 201, with default header values.

no error
*/
type PostConfigCreated struct {
	Payload *models.Document
}

// IsSuccess returns true when this post config created response has a 2xx status code
func (o *PostConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post config created response has a 3xx status code
func (o *PostConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config created response has a 4xx status code
func (o *PostConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post config created response has a 5xx status code
func (o *PostConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post config created response a status code equal to that given
func (o *PostConfigCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostConfigCreated) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigCreated  %+v", 201, o.Payload)
}

func (o *PostConfigCreated) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigCreated  %+v", 201, o.Payload)
}

func (o *PostConfigCreated) GetPayload() *models.Document {
	return o.Payload
}

func (o *PostConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConfigBadRequest creates a PostConfigBadRequest with default headers values
func NewPostConfigBadRequest() *PostConfigBadRequest {
	return &PostConfigBadRequest{}
}

/* PostConfigBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PostConfigBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this post config bad request response has a 2xx status code
func (o *PostConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post config bad request response has a 3xx status code
func (o *PostConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config bad request response has a 4xx status code
func (o *PostConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post config bad request response has a 5xx status code
func (o *PostConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post config bad request response a status code equal to that given
func (o *PostConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigBadRequest  %+v", 400, o.Payload)
}

func (o *PostConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigBadRequest  %+v", 400, o.Payload)
}

func (o *PostConfigBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PostConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConfigUnauthorized creates a PostConfigUnauthorized with default headers values
func NewPostConfigUnauthorized() *PostConfigUnauthorized {
	return &PostConfigUnauthorized{}
}

/* PostConfigUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type PostConfigUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this post config unauthorized response has a 2xx status code
func (o *PostConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post config unauthorized response has a 3xx status code
func (o *PostConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config unauthorized response has a 4xx status code
func (o *PostConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post config unauthorized response has a 5xx status code
func (o *PostConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post config unauthorized response a status code equal to that given
func (o *PostConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConfigUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PostConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConfigForbidden creates a PostConfigForbidden with default headers values
func NewPostConfigForbidden() *PostConfigForbidden {
	return &PostConfigForbidden{}
}

/* PostConfigForbidden describes a response with status code 403, with default header values.

forbidden
*/
type PostConfigForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this post config forbidden response has a 2xx status code
func (o *PostConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post config forbidden response has a 3xx status code
func (o *PostConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config forbidden response has a 4xx status code
func (o *PostConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post config forbidden response has a 5xx status code
func (o *PostConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post config forbidden response a status code equal to that given
func (o *PostConfigForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigForbidden  %+v", 403, o.Payload)
}

func (o *PostConfigForbidden) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigForbidden  %+v", 403, o.Payload)
}

func (o *PostConfigForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PostConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConfigNotFound creates a PostConfigNotFound with default headers values
func NewPostConfigNotFound() *PostConfigNotFound {
	return &PostConfigNotFound{}
}

/* PostConfigNotFound describes a response with status code 404, with default header values.

not found
*/
type PostConfigNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this post config not found response has a 2xx status code
func (o *PostConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post config not found response has a 3xx status code
func (o *PostConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config not found response has a 4xx status code
func (o *PostConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post config not found response has a 5xx status code
func (o *PostConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post config not found response a status code equal to that given
func (o *PostConfigNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigNotFound  %+v", 404, o.Payload)
}

func (o *PostConfigNotFound) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigNotFound  %+v", 404, o.Payload)
}

func (o *PostConfigNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PostConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConfigInternalServerError creates a PostConfigInternalServerError with default headers values
func NewPostConfigInternalServerError() *PostConfigInternalServerError {
	return &PostConfigInternalServerError{}
}

/* PostConfigInternalServerError describes a response with status code 500, with default header values.

server error
*/
type PostConfigInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this post config internal server error response has a 2xx status code
func (o *PostConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post config internal server error response has a 3xx status code
func (o *PostConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post config internal server error response has a 4xx status code
func (o *PostConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post config internal server error response has a 5xx status code
func (o *PostConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post config internal server error response a status code equal to that given
func (o *PostConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config][%d] postConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /config][%d] postConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConfigInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *PostConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
