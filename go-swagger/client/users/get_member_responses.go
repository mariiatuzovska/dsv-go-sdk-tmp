// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// GetMemberReader is a Reader for the GetMember structure.
type GetMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMemberOK creates a GetMemberOK with default headers values
func NewGetMemberOK() *GetMemberOK {
	return &GetMemberOK{}
}

/* GetMemberOK describes a response with status code 200, with default header values.

no error
*/
type GetMemberOK struct {
	Payload *models.MemberResponse
}

// IsSuccess returns true when this get member o k response has a 2xx status code
func (o *GetMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get member o k response has a 3xx status code
func (o *GetMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get member o k response has a 4xx status code
func (o *GetMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get member o k response has a 5xx status code
func (o *GetMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get member o k response a status code equal to that given
func (o *GetMemberOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMemberOK) Error() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberOK  %+v", 200, o.Payload)
}

func (o *GetMemberOK) String() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberOK  %+v", 200, o.Payload)
}

func (o *GetMemberOK) GetPayload() *models.MemberResponse {
	return o.Payload
}

func (o *GetMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemberBadRequest creates a GetMemberBadRequest with default headers values
func NewGetMemberBadRequest() *GetMemberBadRequest {
	return &GetMemberBadRequest{}
}

/* GetMemberBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetMemberBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get member bad request response has a 2xx status code
func (o *GetMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get member bad request response has a 3xx status code
func (o *GetMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get member bad request response has a 4xx status code
func (o *GetMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get member bad request response has a 5xx status code
func (o *GetMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get member bad request response a status code equal to that given
func (o *GetMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetMemberBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberBadRequest  %+v", 400, o.Payload)
}

func (o *GetMemberBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberBadRequest  %+v", 400, o.Payload)
}

func (o *GetMemberBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemberUnauthorized creates a GetMemberUnauthorized with default headers values
func NewGetMemberUnauthorized() *GetMemberUnauthorized {
	return &GetMemberUnauthorized{}
}

/* GetMemberUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type GetMemberUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get member unauthorized response has a 2xx status code
func (o *GetMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get member unauthorized response has a 3xx status code
func (o *GetMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get member unauthorized response has a 4xx status code
func (o *GetMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get member unauthorized response has a 5xx status code
func (o *GetMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get member unauthorized response a status code equal to that given
func (o *GetMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetMemberUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMemberUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMemberUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemberForbidden creates a GetMemberForbidden with default headers values
func NewGetMemberForbidden() *GetMemberForbidden {
	return &GetMemberForbidden{}
}

/* GetMemberForbidden describes a response with status code 403, with default header values.

forbidden
*/
type GetMemberForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get member forbidden response has a 2xx status code
func (o *GetMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get member forbidden response has a 3xx status code
func (o *GetMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get member forbidden response has a 4xx status code
func (o *GetMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get member forbidden response has a 5xx status code
func (o *GetMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get member forbidden response a status code equal to that given
func (o *GetMemberForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetMemberForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetMemberForbidden) String() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetMemberForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemberNotFound creates a GetMemberNotFound with default headers values
func NewGetMemberNotFound() *GetMemberNotFound {
	return &GetMemberNotFound{}
}

/* GetMemberNotFound describes a response with status code 404, with default header values.

not found
*/
type GetMemberNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get member not found response has a 2xx status code
func (o *GetMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get member not found response has a 3xx status code
func (o *GetMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get member not found response has a 4xx status code
func (o *GetMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get member not found response has a 5xx status code
func (o *GetMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get member not found response a status code equal to that given
func (o *GetMemberNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetMemberNotFound) String() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetMemberNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMemberInternalServerError creates a GetMemberInternalServerError with default headers values
func NewGetMemberInternalServerError() *GetMemberInternalServerError {
	return &GetMemberInternalServerError{}
}

/* GetMemberInternalServerError describes a response with status code 500, with default header values.

server error
*/
type GetMemberInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this get member internal server error response has a 2xx status code
func (o *GetMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get member internal server error response has a 3xx status code
func (o *GetMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get member internal server error response has a 4xx status code
func (o *GetMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get member internal server error response has a 5xx status code
func (o *GetMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get member internal server error response a status code equal to that given
func (o *GetMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetMemberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMemberInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{name}/groups][%d] getMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMemberInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *GetMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}