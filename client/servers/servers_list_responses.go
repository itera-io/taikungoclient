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

// ServersListReader is a Reader for the ServersList structure.
type ServersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersListOK creates a ServersListOK with default headers values
func NewServersListOK() *ServersListOK {
	return &ServersListOK{}
}

/*
ServersListOK describes a response with status code 200, with default header values.

Success
*/
type ServersListOK struct {
	Payload *models.ServersList
}

// IsSuccess returns true when this servers list o k response has a 2xx status code
func (o *ServersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers list o k response has a 3xx status code
func (o *ServersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers list o k response has a 4xx status code
func (o *ServersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers list o k response has a 5xx status code
func (o *ServersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this servers list o k response a status code equal to that given
func (o *ServersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListOK  %+v", 200, o.Payload)
}

func (o *ServersListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListOK  %+v", 200, o.Payload)
}

func (o *ServersListOK) GetPayload() *models.ServersList {
	return o.Payload
}

func (o *ServersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServersList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersListBadRequest creates a ServersListBadRequest with default headers values
func NewServersListBadRequest() *ServersListBadRequest {
	return &ServersListBadRequest{}
}

/*
ServersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this servers list bad request response has a 2xx status code
func (o *ServersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers list bad request response has a 3xx status code
func (o *ServersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers list bad request response has a 4xx status code
func (o *ServersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers list bad request response has a 5xx status code
func (o *ServersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this servers list bad request response a status code equal to that given
func (o *ServersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListBadRequest  %+v", 400, o.Payload)
}

func (o *ServersListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListBadRequest  %+v", 400, o.Payload)
}

func (o *ServersListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ServersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersListUnauthorized creates a ServersListUnauthorized with default headers values
func NewServersListUnauthorized() *ServersListUnauthorized {
	return &ServersListUnauthorized{}
}

/*
ServersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this servers list unauthorized response has a 2xx status code
func (o *ServersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers list unauthorized response has a 3xx status code
func (o *ServersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers list unauthorized response has a 4xx status code
func (o *ServersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers list unauthorized response has a 5xx status code
func (o *ServersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this servers list unauthorized response a status code equal to that given
func (o *ServersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersListForbidden creates a ServersListForbidden with default headers values
func NewServersListForbidden() *ServersListForbidden {
	return &ServersListForbidden{}
}

/*
ServersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this servers list forbidden response has a 2xx status code
func (o *ServersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers list forbidden response has a 3xx status code
func (o *ServersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers list forbidden response has a 4xx status code
func (o *ServersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers list forbidden response has a 5xx status code
func (o *ServersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this servers list forbidden response a status code equal to that given
func (o *ServersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ServersListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListForbidden  %+v", 403, o.Payload)
}

func (o *ServersListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListForbidden  %+v", 403, o.Payload)
}

func (o *ServersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersListNotFound creates a ServersListNotFound with default headers values
func NewServersListNotFound() *ServersListNotFound {
	return &ServersListNotFound{}
}

/*
ServersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this servers list not found response has a 2xx status code
func (o *ServersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers list not found response has a 3xx status code
func (o *ServersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers list not found response has a 4xx status code
func (o *ServersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers list not found response has a 5xx status code
func (o *ServersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this servers list not found response a status code equal to that given
func (o *ServersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ServersListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListNotFound  %+v", 404, o.Payload)
}

func (o *ServersListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListNotFound  %+v", 404, o.Payload)
}

func (o *ServersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ServersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersListInternalServerError creates a ServersListInternalServerError with default headers values
func NewServersListInternalServerError() *ServersListInternalServerError {
	return &ServersListInternalServerError{}
}

/*
ServersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersListInternalServerError struct {
}

// IsSuccess returns true when this servers list internal server error response has a 2xx status code
func (o *ServersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers list internal server error response has a 3xx status code
func (o *ServersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers list internal server error response has a 4xx status code
func (o *ServersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers list internal server error response has a 5xx status code
func (o *ServersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this servers list internal server error response a status code equal to that given
func (o *ServersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListInternalServerError ", 500)
}

func (o *ServersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Servers][%d] serversListInternalServerError ", 500)
}

func (o *ServersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
