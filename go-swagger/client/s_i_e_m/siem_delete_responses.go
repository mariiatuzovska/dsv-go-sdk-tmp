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

// SiemDeleteReader is a Reader for the SiemDelete structure.
type SiemDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SiemDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSiemDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSiemDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSiemDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSiemDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSiemDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSiemDeleteOK creates a SiemDeleteOK with default headers values
func NewSiemDeleteOK() *SiemDeleteOK {
	return &SiemDeleteOK{}
}

/* SiemDeleteOK describes a response with status code 200, with default header values.

no error
*/
type SiemDeleteOK struct {
}

// IsSuccess returns true when this siem delete o k response has a 2xx status code
func (o *SiemDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this siem delete o k response has a 3xx status code
func (o *SiemDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem delete o k response has a 4xx status code
func (o *SiemDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this siem delete o k response has a 5xx status code
func (o *SiemDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this siem delete o k response a status code equal to that given
func (o *SiemDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *SiemDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteOK ", 200)
}

func (o *SiemDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteOK ", 200)
}

func (o *SiemDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSiemDeleteBadRequest creates a SiemDeleteBadRequest with default headers values
func NewSiemDeleteBadRequest() *SiemDeleteBadRequest {
	return &SiemDeleteBadRequest{}
}

/* SiemDeleteBadRequest describes a response with status code 400, with default header values.

bad request
*/
type SiemDeleteBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem delete bad request response has a 2xx status code
func (o *SiemDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem delete bad request response has a 3xx status code
func (o *SiemDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem delete bad request response has a 4xx status code
func (o *SiemDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this siem delete bad request response has a 5xx status code
func (o *SiemDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this siem delete bad request response a status code equal to that given
func (o *SiemDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SiemDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SiemDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SiemDeleteBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemDeleteUnauthorized creates a SiemDeleteUnauthorized with default headers values
func NewSiemDeleteUnauthorized() *SiemDeleteUnauthorized {
	return &SiemDeleteUnauthorized{}
}

/* SiemDeleteUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type SiemDeleteUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem delete unauthorized response has a 2xx status code
func (o *SiemDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem delete unauthorized response has a 3xx status code
func (o *SiemDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem delete unauthorized response has a 4xx status code
func (o *SiemDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this siem delete unauthorized response has a 5xx status code
func (o *SiemDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this siem delete unauthorized response a status code equal to that given
func (o *SiemDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SiemDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SiemDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *SiemDeleteUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemDeleteForbidden creates a SiemDeleteForbidden with default headers values
func NewSiemDeleteForbidden() *SiemDeleteForbidden {
	return &SiemDeleteForbidden{}
}

/* SiemDeleteForbidden describes a response with status code 403, with default header values.

forbidden
*/
type SiemDeleteForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem delete forbidden response has a 2xx status code
func (o *SiemDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem delete forbidden response has a 3xx status code
func (o *SiemDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem delete forbidden response has a 4xx status code
func (o *SiemDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this siem delete forbidden response has a 5xx status code
func (o *SiemDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this siem delete forbidden response a status code equal to that given
func (o *SiemDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SiemDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SiemDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SiemDeleteForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSiemDeleteInternalServerError creates a SiemDeleteInternalServerError with default headers values
func NewSiemDeleteInternalServerError() *SiemDeleteInternalServerError {
	return &SiemDeleteInternalServerError{}
}

/* SiemDeleteInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SiemDeleteInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this siem delete internal server error response has a 2xx status code
func (o *SiemDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this siem delete internal server error response has a 3xx status code
func (o *SiemDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this siem delete internal server error response has a 4xx status code
func (o *SiemDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this siem delete internal server error response has a 5xx status code
func (o *SiemDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this siem delete internal server error response a status code equal to that given
func (o *SiemDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SiemDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SiemDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /config/siem/{name}][%d] siemDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SiemDeleteInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *SiemDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}