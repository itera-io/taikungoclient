// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ServersRebootReader is a Reader for the ServersReboot structure.
type ServersRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersRebootBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersRebootOK creates a ServersRebootOK with default headers values
func NewServersRebootOK() *ServersRebootOK {
	return &ServersRebootOK{}
}

/*
ServersRebootOK describes a response with status code 200, with default header values.

Success
*/
type ServersRebootOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this servers reboot o k response has a 2xx status code
func (o *ServersRebootOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers reboot o k response has a 3xx status code
func (o *ServersRebootOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers reboot o k response has a 4xx status code
func (o *ServersRebootOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers reboot o k response has a 5xx status code
func (o *ServersRebootOK) IsServerError() bool {
	return false
}

// IsCode returns true when this servers reboot o k response a status code equal to that given
func (o *ServersRebootOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServersRebootOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootOK  %+v", 200, o.Payload)
}

func (o *ServersRebootOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootOK  %+v", 200, o.Payload)
}

func (o *ServersRebootOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ServersRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersRebootBadRequest creates a ServersRebootBadRequest with default headers values
func NewServersRebootBadRequest() *ServersRebootBadRequest {
	return &ServersRebootBadRequest{}
}

/*
ServersRebootBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersRebootBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this servers reboot bad request response has a 2xx status code
func (o *ServersRebootBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers reboot bad request response has a 3xx status code
func (o *ServersRebootBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers reboot bad request response has a 4xx status code
func (o *ServersRebootBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers reboot bad request response has a 5xx status code
func (o *ServersRebootBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this servers reboot bad request response a status code equal to that given
func (o *ServersRebootBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServersRebootBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootBadRequest  %+v", 400, o.Payload)
}

func (o *ServersRebootBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootBadRequest  %+v", 400, o.Payload)
}

func (o *ServersRebootBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ServersRebootBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersRebootUnauthorized creates a ServersRebootUnauthorized with default headers values
func NewServersRebootUnauthorized() *ServersRebootUnauthorized {
	return &ServersRebootUnauthorized{}
}

/*
ServersRebootUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersRebootUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers reboot unauthorized response has a 2xx status code
func (o *ServersRebootUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers reboot unauthorized response has a 3xx status code
func (o *ServersRebootUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers reboot unauthorized response has a 4xx status code
func (o *ServersRebootUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers reboot unauthorized response has a 5xx status code
func (o *ServersRebootUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this servers reboot unauthorized response a status code equal to that given
func (o *ServersRebootUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServersRebootUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersRebootUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersRebootUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersRebootForbidden creates a ServersRebootForbidden with default headers values
func NewServersRebootForbidden() *ServersRebootForbidden {
	return &ServersRebootForbidden{}
}

/*
ServersRebootForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersRebootForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers reboot forbidden response has a 2xx status code
func (o *ServersRebootForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers reboot forbidden response has a 3xx status code
func (o *ServersRebootForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers reboot forbidden response has a 4xx status code
func (o *ServersRebootForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers reboot forbidden response has a 5xx status code
func (o *ServersRebootForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this servers reboot forbidden response a status code equal to that given
func (o *ServersRebootForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ServersRebootForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootForbidden  %+v", 403, o.Payload)
}

func (o *ServersRebootForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootForbidden  %+v", 403, o.Payload)
}

func (o *ServersRebootForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersRebootNotFound creates a ServersRebootNotFound with default headers values
func NewServersRebootNotFound() *ServersRebootNotFound {
	return &ServersRebootNotFound{}
}

/*
ServersRebootNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersRebootNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers reboot not found response has a 2xx status code
func (o *ServersRebootNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers reboot not found response has a 3xx status code
func (o *ServersRebootNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers reboot not found response has a 4xx status code
func (o *ServersRebootNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers reboot not found response has a 5xx status code
func (o *ServersRebootNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this servers reboot not found response a status code equal to that given
func (o *ServersRebootNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ServersRebootNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootNotFound  %+v", 404, o.Payload)
}

func (o *ServersRebootNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootNotFound  %+v", 404, o.Payload)
}

func (o *ServersRebootNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersRebootInternalServerError creates a ServersRebootInternalServerError with default headers values
func NewServersRebootInternalServerError() *ServersRebootInternalServerError {
	return &ServersRebootInternalServerError{}
}

/*
ServersRebootInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersRebootInternalServerError struct {
}

// IsSuccess returns true when this servers reboot internal server error response has a 2xx status code
func (o *ServersRebootInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers reboot internal server error response has a 3xx status code
func (o *ServersRebootInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers reboot internal server error response has a 4xx status code
func (o *ServersRebootInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers reboot internal server error response has a 5xx status code
func (o *ServersRebootInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this servers reboot internal server error response a status code equal to that given
func (o *ServersRebootInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServersRebootInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootInternalServerError ", 500)
}

func (o *ServersRebootInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/reboot][%d] serversRebootInternalServerError ", 500)
}

func (o *ServersRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
