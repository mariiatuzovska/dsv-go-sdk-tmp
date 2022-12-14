// Code generated by go-swagger; DO NOT EDIT.

package s_i_e_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// SiemCreateReader is a Reader for the SiemCreate structure.
type SiemCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SiemCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSiemCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSiemCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSiemCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSiemCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSiemCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSiemCreateCreated creates a SiemCreateCreated with default headers values
func NewSiemCreateCreated() *SiemCreateCreated {
	return &SiemCreateCreated{}
}

/* SiemCreateCreated describes a response with status code 201, with default header values.

no error
*/
type SiemCreateCreated struct {
	Payload *models.SiemResponseModel
}

// IsSuccess returns true when this siem create created response has a 2xx status code
func (o *SiemCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this siem create created response has a 3xx status code
func (o *SiemCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem create created response has a 4xx status code
func (o *SiemCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this siem create created response has a 5xx status code
func (o *SiemCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this siem create created response a status code equal to that given
func (o *SiemCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SiemCreateCreated) Error() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateCreated  %+v", 201, o.Payload)
}

func (o *SiemCreateCreated) String() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateCreated  %+v", 201, o.Payload)
}

func (o *SiemCreateCreated) GetPayload() *models.SiemResponseModel {
	return o.Payload
}

func (o *SiemCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiemResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemCreateBadRequest creates a SiemCreateBadRequest with default headers values
func NewSiemCreateBadRequest() *SiemCreateBadRequest {
	return &SiemCreateBadRequest{}
}

/* SiemCreateBadRequest describes a response with status code 400, with default header values.

bad request
*/
type SiemCreateBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem create bad request response has a 2xx status code
func (o *SiemCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem create bad request response has a 3xx status code
func (o *SiemCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem create bad request response has a 4xx status code
func (o *SiemCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this siem create bad request response has a 5xx status code
func (o *SiemCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this siem create bad request response a status code equal to that given
func (o *SiemCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SiemCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SiemCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SiemCreateBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemCreateUnauthorized creates a SiemCreateUnauthorized with default headers values
func NewSiemCreateUnauthorized() *SiemCreateUnauthorized {
	return &SiemCreateUnauthorized{}
}

/* SiemCreateUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type SiemCreateUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem create unauthorized response has a 2xx status code
func (o *SiemCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem create unauthorized response has a 3xx status code
func (o *SiemCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem create unauthorized response has a 4xx status code
func (o *SiemCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this siem create unauthorized response has a 5xx status code
func (o *SiemCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this siem create unauthorized response a status code equal to that given
func (o *SiemCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SiemCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SiemCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SiemCreateUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemCreateForbidden creates a SiemCreateForbidden with default headers values
func NewSiemCreateForbidden() *SiemCreateForbidden {
	return &SiemCreateForbidden{}
}

/* SiemCreateForbidden describes a response with status code 403, with default header values.

forbidden
*/
type SiemCreateForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem create forbidden response has a 2xx status code
func (o *SiemCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem create forbidden response has a 3xx status code
func (o *SiemCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem create forbidden response has a 4xx status code
func (o *SiemCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this siem create forbidden response has a 5xx status code
func (o *SiemCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this siem create forbidden response a status code equal to that given
func (o *SiemCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SiemCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateForbidden  %+v", 403, o.Payload)
}

func (o *SiemCreateForbidden) String() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateForbidden  %+v", 403, o.Payload)
}

func (o *SiemCreateForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemCreateInternalServerError creates a SiemCreateInternalServerError with default headers values
func NewSiemCreateInternalServerError() *SiemCreateInternalServerError {
	return &SiemCreateInternalServerError{}
}

/* SiemCreateInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SiemCreateInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem create internal server error response has a 2xx status code
func (o *SiemCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem create internal server error response has a 3xx status code
func (o *SiemCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem create internal server error response has a 4xx status code
func (o *SiemCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this siem create internal server error response has a 5xx status code
func (o *SiemCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this siem create internal server error response a status code equal to that given
func (o *SiemCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SiemCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *SiemCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /config/siem][%d] siemCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *SiemCreateInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
