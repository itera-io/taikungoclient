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

// ServersDeleteReader is a Reader for the ServersDelete structure.
type ServersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewServersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersDeleteOK creates a ServersDeleteOK with default headers values
func NewServersDeleteOK() *ServersDeleteOK {
	return &ServersDeleteOK{}
}

/*
ServersDeleteOK describes a response with status code 200, with default header values.

Success
*/
type ServersDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this servers delete o k response has a 2xx status code
func (o *ServersDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers delete o k response has a 3xx status code
func (o *ServersDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete o k response has a 4xx status code
func (o *ServersDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers delete o k response has a 5xx status code
func (o *ServersDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this servers delete o k response a status code equal to that given
func (o *ServersDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServersDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteOK  %+v", 200, o.Payload)
}

func (o *ServersDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteOK  %+v", 200, o.Payload)
}

func (o *ServersDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ServersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDeleteNoContent creates a ServersDeleteNoContent with default headers values
func NewServersDeleteNoContent() *ServersDeleteNoContent {
	return &ServersDeleteNoContent{}
}

/*
ServersDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type ServersDeleteNoContent struct {
}

// IsSuccess returns true when this servers delete no content response has a 2xx status code
func (o *ServersDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers delete no content response has a 3xx status code
func (o *ServersDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete no content response has a 4xx status code
func (o *ServersDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers delete no content response has a 5xx status code
func (o *ServersDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this servers delete no content response a status code equal to that given
func (o *ServersDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ServersDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteNoContent ", 204)
}

func (o *ServersDeleteNoContent) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteNoContent ", 204)
}

func (o *ServersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServersDeleteBadRequest creates a ServersDeleteBadRequest with default headers values
func NewServersDeleteBadRequest() *ServersDeleteBadRequest {
	return &ServersDeleteBadRequest{}
}

/*
ServersDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this servers delete bad request response has a 2xx status code
func (o *ServersDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers delete bad request response has a 3xx status code
func (o *ServersDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete bad request response has a 4xx status code
func (o *ServersDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers delete bad request response has a 5xx status code
func (o *ServersDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this servers delete bad request response a status code equal to that given
func (o *ServersDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServersDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ServersDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ServersDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ServersDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDeleteUnauthorized creates a ServersDeleteUnauthorized with default headers values
func NewServersDeleteUnauthorized() *ServersDeleteUnauthorized {
	return &ServersDeleteUnauthorized{}
}

/*
ServersDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this servers delete unauthorized response has a 2xx status code
func (o *ServersDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers delete unauthorized response has a 3xx status code
func (o *ServersDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete unauthorized response has a 4xx status code
func (o *ServersDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers delete unauthorized response has a 5xx status code
func (o *ServersDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this servers delete unauthorized response a status code equal to that given
func (o *ServersDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServersDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDeleteForbidden creates a ServersDeleteForbidden with default headers values
func NewServersDeleteForbidden() *ServersDeleteForbidden {
	return &ServersDeleteForbidden{}
}

/*
ServersDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this servers delete forbidden response has a 2xx status code
func (o *ServersDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers delete forbidden response has a 3xx status code
func (o *ServersDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete forbidden response has a 4xx status code
func (o *ServersDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers delete forbidden response has a 5xx status code
func (o *ServersDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this servers delete forbidden response a status code equal to that given
func (o *ServersDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ServersDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ServersDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ServersDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDeleteNotFound creates a ServersDeleteNotFound with default headers values
func NewServersDeleteNotFound() *ServersDeleteNotFound {
	return &ServersDeleteNotFound{}
}

/*
ServersDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this servers delete not found response has a 2xx status code
func (o *ServersDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers delete not found response has a 3xx status code
func (o *ServersDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete not found response has a 4xx status code
func (o *ServersDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers delete not found response has a 5xx status code
func (o *ServersDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this servers delete not found response a status code equal to that given
func (o *ServersDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ServersDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ServersDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ServersDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersDeleteInternalServerError creates a ServersDeleteInternalServerError with default headers values
func NewServersDeleteInternalServerError() *ServersDeleteInternalServerError {
	return &ServersDeleteInternalServerError{}
}

/*
ServersDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersDeleteInternalServerError struct {
}

// IsSuccess returns true when this servers delete internal server error response has a 2xx status code
func (o *ServersDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers delete internal server error response has a 3xx status code
func (o *ServersDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers delete internal server error response has a 4xx status code
func (o *ServersDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers delete internal server error response has a 5xx status code
func (o *ServersDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this servers delete internal server error response a status code equal to that given
func (o *ServersDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServersDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteInternalServerError ", 500)
}

func (o *ServersDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/delete][%d] serversDeleteInternalServerError ", 500)
}

func (o *ServersDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
