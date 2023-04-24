// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ServersUpdateReader is a Reader for the ServersUpdate structure.
type ServersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersUpdateOK creates a ServersUpdateOK with default headers values
func NewServersUpdateOK() *ServersUpdateOK {
	return &ServersUpdateOK{}
}

/*
ServersUpdateOK describes a response with status code 200, with default header values.

Success
*/
type ServersUpdateOK struct {
}

// IsSuccess returns true when this servers update o k response has a 2xx status code
func (o *ServersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers update o k response has a 3xx status code
func (o *ServersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update o k response has a 4xx status code
func (o *ServersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers update o k response has a 5xx status code
func (o *ServersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update o k response a status code equal to that given
func (o *ServersUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the servers update o k response
func (o *ServersUpdateOK) Code() int {
	return 200
}

func (o *ServersUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateOK ", 200)
}

func (o *ServersUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateOK ", 200)
}

func (o *ServersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewServersUpdateBadRequest creates a ServersUpdateBadRequest with default headers values
func NewServersUpdateBadRequest() *ServersUpdateBadRequest {
	return &ServersUpdateBadRequest{}
}

/*
ServersUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this servers update bad request response has a 2xx status code
func (o *ServersUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update bad request response has a 3xx status code
func (o *ServersUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update bad request response has a 4xx status code
func (o *ServersUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update bad request response has a 5xx status code
func (o *ServersUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update bad request response a status code equal to that given
func (o *ServersUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the servers update bad request response
func (o *ServersUpdateBadRequest) Code() int {
	return 400
}

func (o *ServersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ServersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateUnauthorized creates a ServersUpdateUnauthorized with default headers values
func NewServersUpdateUnauthorized() *ServersUpdateUnauthorized {
	return &ServersUpdateUnauthorized{}
}

/*
ServersUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersUpdateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this servers update unauthorized response has a 2xx status code
func (o *ServersUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update unauthorized response has a 3xx status code
func (o *ServersUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update unauthorized response has a 4xx status code
func (o *ServersUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update unauthorized response has a 5xx status code
func (o *ServersUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update unauthorized response a status code equal to that given
func (o *ServersUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the servers update unauthorized response
func (o *ServersUpdateUnauthorized) Code() int {
	return 401
}

func (o *ServersUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersUpdateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ServersUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateForbidden creates a ServersUpdateForbidden with default headers values
func NewServersUpdateForbidden() *ServersUpdateForbidden {
	return &ServersUpdateForbidden{}
}

/*
ServersUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersUpdateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this servers update forbidden response has a 2xx status code
func (o *ServersUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update forbidden response has a 3xx status code
func (o *ServersUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update forbidden response has a 4xx status code
func (o *ServersUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update forbidden response has a 5xx status code
func (o *ServersUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update forbidden response a status code equal to that given
func (o *ServersUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the servers update forbidden response
func (o *ServersUpdateForbidden) Code() int {
	return 403
}

func (o *ServersUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ServersUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ServersUpdateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ServersUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateNotFound creates a ServersUpdateNotFound with default headers values
func NewServersUpdateNotFound() *ServersUpdateNotFound {
	return &ServersUpdateNotFound{}
}

/*
ServersUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersUpdateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this servers update not found response has a 2xx status code
func (o *ServersUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update not found response has a 3xx status code
func (o *ServersUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update not found response has a 4xx status code
func (o *ServersUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers update not found response has a 5xx status code
func (o *ServersUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this servers update not found response a status code equal to that given
func (o *ServersUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the servers update not found response
func (o *ServersUpdateNotFound) Code() int {
	return 404
}

func (o *ServersUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ServersUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ServersUpdateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ServersUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersUpdateInternalServerError creates a ServersUpdateInternalServerError with default headers values
func NewServersUpdateInternalServerError() *ServersUpdateInternalServerError {
	return &ServersUpdateInternalServerError{}
}

/*
ServersUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersUpdateInternalServerError struct {
}

// IsSuccess returns true when this servers update internal server error response has a 2xx status code
func (o *ServersUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers update internal server error response has a 3xx status code
func (o *ServersUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers update internal server error response has a 4xx status code
func (o *ServersUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers update internal server error response has a 5xx status code
func (o *ServersUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this servers update internal server error response a status code equal to that given
func (o *ServersUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the servers update internal server error response
func (o *ServersUpdateInternalServerError) Code() int {
	return 500
}

func (o *ServersUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateInternalServerError ", 500)
}

func (o *ServersUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers/update][%d] serversUpdateInternalServerError ", 500)
}

func (o *ServersUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
