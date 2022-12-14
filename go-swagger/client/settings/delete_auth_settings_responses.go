// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/models"
)

// DeleteAuthSettingsReader is a Reader for the DeleteAuthSettings structure.
type DeleteAuthSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAuthSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAuthSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAuthSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAuthSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAuthSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAuthSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAuthSettingsOK creates a DeleteAuthSettingsOK with default headers values
func NewDeleteAuthSettingsOK() *DeleteAuthSettingsOK {
	return &DeleteAuthSettingsOK{}
}

/* DeleteAuthSettingsOK describes a response with status code 200, with default header values.

no error
*/
type DeleteAuthSettingsOK struct {
	Payload *models.MessageResponse
}

// IsSuccess returns true when this delete auth settings o k response has a 2xx status code
func (o *DeleteAuthSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete auth settings o k response has a 3xx status code
func (o *DeleteAuthSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth settings o k response has a 4xx status code
func (o *DeleteAuthSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete auth settings o k response has a 5xx status code
func (o *DeleteAuthSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth settings o k response a status code equal to that given
func (o *DeleteAuthSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAuthSettingsOK) Error() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsOK  %+v", 200, o.Payload)
}

func (o *DeleteAuthSettingsOK) String() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsOK  %+v", 200, o.Payload)
}

func (o *DeleteAuthSettingsOK) GetPayload() *models.MessageResponse {
	return o.Payload
}

func (o *DeleteAuthSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthSettingsBadRequest creates a DeleteAuthSettingsBadRequest with default headers values
func NewDeleteAuthSettingsBadRequest() *DeleteAuthSettingsBadRequest {
	return &DeleteAuthSettingsBadRequest{}
}

/* DeleteAuthSettingsBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DeleteAuthSettingsBadRequest struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete auth settings bad request response has a 2xx status code
func (o *DeleteAuthSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete auth settings bad request response has a 3xx status code
func (o *DeleteAuthSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth settings bad request response has a 4xx status code
func (o *DeleteAuthSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete auth settings bad request response has a 5xx status code
func (o *DeleteAuthSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth settings bad request response a status code equal to that given
func (o *DeleteAuthSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteAuthSettingsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAuthSettingsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAuthSettingsBadRequest) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteAuthSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthSettingsUnauthorized creates a DeleteAuthSettingsUnauthorized with default headers values
func NewDeleteAuthSettingsUnauthorized() *DeleteAuthSettingsUnauthorized {
	return &DeleteAuthSettingsUnauthorized{}
}

/* DeleteAuthSettingsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteAuthSettingsUnauthorized struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete auth settings unauthorized response has a 2xx status code
func (o *DeleteAuthSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete auth settings unauthorized response has a 3xx status code
func (o *DeleteAuthSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth settings unauthorized response has a 4xx status code
func (o *DeleteAuthSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete auth settings unauthorized response has a 5xx status code
func (o *DeleteAuthSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth settings unauthorized response a status code equal to that given
func (o *DeleteAuthSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteAuthSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAuthSettingsUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAuthSettingsUnauthorized) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteAuthSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthSettingsForbidden creates a DeleteAuthSettingsForbidden with default headers values
func NewDeleteAuthSettingsForbidden() *DeleteAuthSettingsForbidden {
	return &DeleteAuthSettingsForbidden{}
}

/* DeleteAuthSettingsForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteAuthSettingsForbidden struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete auth settings forbidden response has a 2xx status code
func (o *DeleteAuthSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete auth settings forbidden response has a 3xx status code
func (o *DeleteAuthSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth settings forbidden response has a 4xx status code
func (o *DeleteAuthSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete auth settings forbidden response has a 5xx status code
func (o *DeleteAuthSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth settings forbidden response a status code equal to that given
func (o *DeleteAuthSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteAuthSettingsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAuthSettingsForbidden) String() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAuthSettingsForbidden) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteAuthSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthSettingsNotFound creates a DeleteAuthSettingsNotFound with default headers values
func NewDeleteAuthSettingsNotFound() *DeleteAuthSettingsNotFound {
	return &DeleteAuthSettingsNotFound{}
}

/* DeleteAuthSettingsNotFound describes a response with status code 404, with default header values.

not found
*/
type DeleteAuthSettingsNotFound struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete auth settings not found response has a 2xx status code
func (o *DeleteAuthSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete auth settings not found response has a 3xx status code
func (o *DeleteAuthSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth settings not found response has a 4xx status code
func (o *DeleteAuthSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete auth settings not found response has a 5xx status code
func (o *DeleteAuthSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth settings not found response a status code equal to that given
func (o *DeleteAuthSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteAuthSettingsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAuthSettingsNotFound) String() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAuthSettingsNotFound) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteAuthSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthSettingsInternalServerError creates a DeleteAuthSettingsInternalServerError with default headers values
func NewDeleteAuthSettingsInternalServerError() *DeleteAuthSettingsInternalServerError {
	return &DeleteAuthSettingsInternalServerError{}
}

/* DeleteAuthSettingsInternalServerError describes a response with status code 500, with default header values.

server error
*/
type DeleteAuthSettingsInternalServerError struct {
	Payload *models.HTTPError
}

// IsSuccess returns true when this delete auth settings internal server error response has a 2xx status code
func (o *DeleteAuthSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete auth settings internal server error response has a 3xx status code
func (o *DeleteAuthSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth settings internal server error response has a 4xx status code
func (o *DeleteAuthSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete auth settings internal server error response has a 5xx status code
func (o *DeleteAuthSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete auth settings internal server error response a status code equal to that given
func (o *DeleteAuthSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteAuthSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAuthSettingsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /config/auth/{name}][%d] deleteAuthSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAuthSettingsInternalServerError) GetPayload() *models.HTTPError {
	return o.Payload
}

func (o *DeleteAuthSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
