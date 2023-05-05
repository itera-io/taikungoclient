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

// ShowbackCredentialsLockManagerReader is a Reader for the ShowbackCredentialsLockManager structure.
type ShowbackCredentialsLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackCredentialsLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackCredentialsLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackCredentialsLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackCredentialsLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackCredentialsLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackCredentialsLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackCredentialsLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackCredentialsLockManagerOK creates a ShowbackCredentialsLockManagerOK with default headers values
func NewShowbackCredentialsLockManagerOK() *ShowbackCredentialsLockManagerOK {
	return &ShowbackCredentialsLockManagerOK{}
}

/*
ShowbackCredentialsLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackCredentialsLockManagerOK struct {
}

// IsSuccess returns true when this showback credentials lock manager o k response has a 2xx status code
func (o *ShowbackCredentialsLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback credentials lock manager o k response has a 3xx status code
func (o *ShowbackCredentialsLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials lock manager o k response has a 4xx status code
func (o *ShowbackCredentialsLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback credentials lock manager o k response has a 5xx status code
func (o *ShowbackCredentialsLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials lock manager o k response a status code equal to that given
func (o *ShowbackCredentialsLockManagerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the showback credentials lock manager o k response
func (o *ShowbackCredentialsLockManagerOK) Code() int {
	return 200
}

func (o *ShowbackCredentialsLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerOK ", 200)
}

func (o *ShowbackCredentialsLockManagerOK) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerOK ", 200)
}

func (o *ShowbackCredentialsLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowbackCredentialsLockManagerBadRequest creates a ShowbackCredentialsLockManagerBadRequest with default headers values
func NewShowbackCredentialsLockManagerBadRequest() *ShowbackCredentialsLockManagerBadRequest {
	return &ShowbackCredentialsLockManagerBadRequest{}
}

/*
ShowbackCredentialsLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackCredentialsLockManagerBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback credentials lock manager bad request response has a 2xx status code
func (o *ShowbackCredentialsLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials lock manager bad request response has a 3xx status code
func (o *ShowbackCredentialsLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials lock manager bad request response has a 4xx status code
func (o *ShowbackCredentialsLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials lock manager bad request response has a 5xx status code
func (o *ShowbackCredentialsLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials lock manager bad request response a status code equal to that given
func (o *ShowbackCredentialsLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the showback credentials lock manager bad request response
func (o *ShowbackCredentialsLockManagerBadRequest) Code() int {
	return 400
}

func (o *ShowbackCredentialsLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackCredentialsLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackCredentialsLockManagerBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsLockManagerUnauthorized creates a ShowbackCredentialsLockManagerUnauthorized with default headers values
func NewShowbackCredentialsLockManagerUnauthorized() *ShowbackCredentialsLockManagerUnauthorized {
	return &ShowbackCredentialsLockManagerUnauthorized{}
}

/*
ShowbackCredentialsLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackCredentialsLockManagerUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials lock manager unauthorized response has a 2xx status code
func (o *ShowbackCredentialsLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials lock manager unauthorized response has a 3xx status code
func (o *ShowbackCredentialsLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials lock manager unauthorized response has a 4xx status code
func (o *ShowbackCredentialsLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials lock manager unauthorized response has a 5xx status code
func (o *ShowbackCredentialsLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials lock manager unauthorized response a status code equal to that given
func (o *ShowbackCredentialsLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the showback credentials lock manager unauthorized response
func (o *ShowbackCredentialsLockManagerUnauthorized) Code() int {
	return 401
}

func (o *ShowbackCredentialsLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackCredentialsLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackCredentialsLockManagerUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsLockManagerForbidden creates a ShowbackCredentialsLockManagerForbidden with default headers values
func NewShowbackCredentialsLockManagerForbidden() *ShowbackCredentialsLockManagerForbidden {
	return &ShowbackCredentialsLockManagerForbidden{}
}

/*
ShowbackCredentialsLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackCredentialsLockManagerForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials lock manager forbidden response has a 2xx status code
func (o *ShowbackCredentialsLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials lock manager forbidden response has a 3xx status code
func (o *ShowbackCredentialsLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials lock manager forbidden response has a 4xx status code
func (o *ShowbackCredentialsLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials lock manager forbidden response has a 5xx status code
func (o *ShowbackCredentialsLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials lock manager forbidden response a status code equal to that given
func (o *ShowbackCredentialsLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the showback credentials lock manager forbidden response
func (o *ShowbackCredentialsLockManagerForbidden) Code() int {
	return 403
}

func (o *ShowbackCredentialsLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackCredentialsLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackCredentialsLockManagerForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsLockManagerNotFound creates a ShowbackCredentialsLockManagerNotFound with default headers values
func NewShowbackCredentialsLockManagerNotFound() *ShowbackCredentialsLockManagerNotFound {
	return &ShowbackCredentialsLockManagerNotFound{}
}

/*
ShowbackCredentialsLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackCredentialsLockManagerNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback credentials lock manager not found response has a 2xx status code
func (o *ShowbackCredentialsLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials lock manager not found response has a 3xx status code
func (o *ShowbackCredentialsLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials lock manager not found response has a 4xx status code
func (o *ShowbackCredentialsLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback credentials lock manager not found response has a 5xx status code
func (o *ShowbackCredentialsLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback credentials lock manager not found response a status code equal to that given
func (o *ShowbackCredentialsLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the showback credentials lock manager not found response
func (o *ShowbackCredentialsLockManagerNotFound) Code() int {
	return 404
}

func (o *ShowbackCredentialsLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackCredentialsLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackCredentialsLockManagerNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackCredentialsLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsLockManagerInternalServerError creates a ShowbackCredentialsLockManagerInternalServerError with default headers values
func NewShowbackCredentialsLockManagerInternalServerError() *ShowbackCredentialsLockManagerInternalServerError {
	return &ShowbackCredentialsLockManagerInternalServerError{}
}

/*
ShowbackCredentialsLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackCredentialsLockManagerInternalServerError struct {
}

// IsSuccess returns true when this showback credentials lock manager internal server error response has a 2xx status code
func (o *ShowbackCredentialsLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback credentials lock manager internal server error response has a 3xx status code
func (o *ShowbackCredentialsLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback credentials lock manager internal server error response has a 4xx status code
func (o *ShowbackCredentialsLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback credentials lock manager internal server error response has a 5xx status code
func (o *ShowbackCredentialsLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback credentials lock manager internal server error response a status code equal to that given
func (o *ShowbackCredentialsLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the showback credentials lock manager internal server error response
func (o *ShowbackCredentialsLockManagerInternalServerError) Code() int {
	return 500
}

func (o *ShowbackCredentialsLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerInternalServerError ", 500)
}

func (o *ShowbackCredentialsLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/lockmanager][%d] showbackCredentialsLockManagerInternalServerError ", 500)
}

func (o *ShowbackCredentialsLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
