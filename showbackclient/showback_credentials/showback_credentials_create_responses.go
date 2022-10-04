// Code generated by go-swagger; DO NOT EDIT.

package showback_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackCredentialsCreateReader is a Reader for the ShowbackCredentialsCreate structure.
type ShowbackCredentialsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackCredentialsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackCredentialsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackCredentialsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackCredentialsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackCredentialsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackCredentialsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackCredentialsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackCredentialsCreateOK creates a ShowbackCredentialsCreateOK with default headers values
func NewShowbackCredentialsCreateOK() *ShowbackCredentialsCreateOK {
	return &ShowbackCredentialsCreateOK{}
}

/*
ShowbackCredentialsCreateOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackCredentialsCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this showback credentials create o k response has a 2xx status code
func (o *ShowbackCredentialsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback credentials create o k response has a 3xx status code
func (o *ShowbackCredentialsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials create o k response has a 4xx status code
func (o *ShowbackCredentialsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback credentials create o k response has a 5xx status code
func (o *ShowbackCredentialsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials create o k response a status code equal to that given
func (o *ShowbackCredentialsCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackCredentialsCreateOK) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateOK  %+v", 200, o.Payload)
}

func (o *ShowbackCredentialsCreateOK) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateOK  %+v", 200, o.Payload)
}

func (o *ShowbackCredentialsCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *ShowbackCredentialsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateBadRequest creates a ShowbackCredentialsCreateBadRequest with default headers values
func NewShowbackCredentialsCreateBadRequest() *ShowbackCredentialsCreateBadRequest {
	return &ShowbackCredentialsCreateBadRequest{}
}

/*
ShowbackCredentialsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackCredentialsCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback credentials create bad request response has a 2xx status code
func (o *ShowbackCredentialsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials create bad request response has a 3xx status code
func (o *ShowbackCredentialsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials create bad request response has a 4xx status code
func (o *ShowbackCredentialsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials create bad request response has a 5xx status code
func (o *ShowbackCredentialsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials create bad request response a status code equal to that given
func (o *ShowbackCredentialsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackCredentialsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackCredentialsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackCredentialsCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateUnauthorized creates a ShowbackCredentialsCreateUnauthorized with default headers values
func NewShowbackCredentialsCreateUnauthorized() *ShowbackCredentialsCreateUnauthorized {
	return &ShowbackCredentialsCreateUnauthorized{}
}

/*
ShowbackCredentialsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackCredentialsCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials create unauthorized response has a 2xx status code
func (o *ShowbackCredentialsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials create unauthorized response has a 3xx status code
func (o *ShowbackCredentialsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials create unauthorized response has a 4xx status code
func (o *ShowbackCredentialsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials create unauthorized response has a 5xx status code
func (o *ShowbackCredentialsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials create unauthorized response a status code equal to that given
func (o *ShowbackCredentialsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackCredentialsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackCredentialsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackCredentialsCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateForbidden creates a ShowbackCredentialsCreateForbidden with default headers values
func NewShowbackCredentialsCreateForbidden() *ShowbackCredentialsCreateForbidden {
	return &ShowbackCredentialsCreateForbidden{}
}

/*
ShowbackCredentialsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackCredentialsCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials create forbidden response has a 2xx status code
func (o *ShowbackCredentialsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials create forbidden response has a 3xx status code
func (o *ShowbackCredentialsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials create forbidden response has a 4xx status code
func (o *ShowbackCredentialsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials create forbidden response has a 5xx status code
func (o *ShowbackCredentialsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials create forbidden response a status code equal to that given
func (o *ShowbackCredentialsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackCredentialsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackCredentialsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackCredentialsCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateNotFound creates a ShowbackCredentialsCreateNotFound with default headers values
func NewShowbackCredentialsCreateNotFound() *ShowbackCredentialsCreateNotFound {
	return &ShowbackCredentialsCreateNotFound{}
}

/*
ShowbackCredentialsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackCredentialsCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials create not found response has a 2xx status code
func (o *ShowbackCredentialsCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials create not found response has a 3xx status code
func (o *ShowbackCredentialsCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials create not found response has a 4xx status code
func (o *ShowbackCredentialsCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials create not found response has a 5xx status code
func (o *ShowbackCredentialsCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials create not found response a status code equal to that given
func (o *ShowbackCredentialsCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackCredentialsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackCredentialsCreateNotFound) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackCredentialsCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateInternalServerError creates a ShowbackCredentialsCreateInternalServerError with default headers values
func NewShowbackCredentialsCreateInternalServerError() *ShowbackCredentialsCreateInternalServerError {
	return &ShowbackCredentialsCreateInternalServerError{}
}

/*
ShowbackCredentialsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackCredentialsCreateInternalServerError struct {
}

// IsSuccess returns true when this showback credentials create internal server error response has a 2xx status code
func (o *ShowbackCredentialsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials create internal server error response has a 3xx status code
func (o *ShowbackCredentialsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials create internal server error response has a 4xx status code
func (o *ShowbackCredentialsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback credentials create internal server error response has a 5xx status code
func (o *ShowbackCredentialsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback credentials create internal server error response a status code equal to that given
func (o *ShowbackCredentialsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackCredentialsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateInternalServerError ", 500)
}

func (o *ShowbackCredentialsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateInternalServerError ", 500)
}

func (o *ShowbackCredentialsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
