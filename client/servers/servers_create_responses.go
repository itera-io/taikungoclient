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

// ServersCreateReader is a Reader for the ServersCreate structure.
type ServersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServersCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServersCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServersCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServersCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServersCreateOK creates a ServersCreateOK with default headers values
func NewServersCreateOK() *ServersCreateOK {
	return &ServersCreateOK{}
}

/*
ServersCreateOK describes a response with status code 200, with default header values.

Success
*/
type ServersCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this servers create o k response has a 2xx status code
func (o *ServersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this servers create o k response has a 3xx status code
func (o *ServersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers create o k response has a 4xx status code
func (o *ServersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers create o k response has a 5xx status code
func (o *ServersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this servers create o k response a status code equal to that given
func (o *ServersCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServersCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateOK  %+v", 200, o.Payload)
}

func (o *ServersCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateOK  %+v", 200, o.Payload)
}

func (o *ServersCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *ServersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersCreateBadRequest creates a ServersCreateBadRequest with default headers values
func NewServersCreateBadRequest() *ServersCreateBadRequest {
	return &ServersCreateBadRequest{}
}

/*
ServersCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServersCreateBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers create bad request response has a 2xx status code
func (o *ServersCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers create bad request response has a 3xx status code
func (o *ServersCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers create bad request response has a 4xx status code
func (o *ServersCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers create bad request response has a 5xx status code
func (o *ServersCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this servers create bad request response a status code equal to that given
func (o *ServersCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ServersCreateBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersCreateUnauthorized creates a ServersCreateUnauthorized with default headers values
func NewServersCreateUnauthorized() *ServersCreateUnauthorized {
	return &ServersCreateUnauthorized{}
}

/*
ServersCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServersCreateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers create unauthorized response has a 2xx status code
func (o *ServersCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers create unauthorized response has a 3xx status code
func (o *ServersCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers create unauthorized response has a 4xx status code
func (o *ServersCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers create unauthorized response has a 5xx status code
func (o *ServersCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this servers create unauthorized response a status code equal to that given
func (o *ServersCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServersCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ServersCreateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersCreateForbidden creates a ServersCreateForbidden with default headers values
func NewServersCreateForbidden() *ServersCreateForbidden {
	return &ServersCreateForbidden{}
}

/*
ServersCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServersCreateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers create forbidden response has a 2xx status code
func (o *ServersCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers create forbidden response has a 3xx status code
func (o *ServersCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers create forbidden response has a 4xx status code
func (o *ServersCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers create forbidden response has a 5xx status code
func (o *ServersCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this servers create forbidden response a status code equal to that given
func (o *ServersCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ServersCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateForbidden  %+v", 403, o.Payload)
}

func (o *ServersCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateForbidden  %+v", 403, o.Payload)
}

func (o *ServersCreateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersCreateNotFound creates a ServersCreateNotFound with default headers values
func NewServersCreateNotFound() *ServersCreateNotFound {
	return &ServersCreateNotFound{}
}

/*
ServersCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServersCreateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this servers create not found response has a 2xx status code
func (o *ServersCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers create not found response has a 3xx status code
func (o *ServersCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers create not found response has a 4xx status code
func (o *ServersCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this servers create not found response has a 5xx status code
func (o *ServersCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this servers create not found response a status code equal to that given
func (o *ServersCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ServersCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateNotFound  %+v", 404, o.Payload)
}

func (o *ServersCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateNotFound  %+v", 404, o.Payload)
}

func (o *ServersCreateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ServersCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServersCreateInternalServerError creates a ServersCreateInternalServerError with default headers values
func NewServersCreateInternalServerError() *ServersCreateInternalServerError {
	return &ServersCreateInternalServerError{}
}

/*
ServersCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ServersCreateInternalServerError struct {
}

// IsSuccess returns true when this servers create internal server error response has a 2xx status code
func (o *ServersCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this servers create internal server error response has a 3xx status code
func (o *ServersCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this servers create internal server error response has a 4xx status code
func (o *ServersCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this servers create internal server error response has a 5xx status code
func (o *ServersCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this servers create internal server error response a status code equal to that given
func (o *ServersCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServersCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateInternalServerError ", 500)
}

func (o *ServersCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Servers][%d] serversCreateInternalServerError ", 500)
}

func (o *ServersCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
