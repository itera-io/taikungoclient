// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AccessProfilesLockManagerReader is a Reader for the AccessProfilesLockManager structure.
type AccessProfilesLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessProfilesLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessProfilesLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessProfilesLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccessProfilesLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccessProfilesLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessProfilesLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessProfilesLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccessProfilesLockManagerOK creates a AccessProfilesLockManagerOK with default headers values
func NewAccessProfilesLockManagerOK() *AccessProfilesLockManagerOK {
	return &AccessProfilesLockManagerOK{}
}

/*
AccessProfilesLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type AccessProfilesLockManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this access profiles lock manager o k response has a 2xx status code
func (o *AccessProfilesLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access profiles lock manager o k response has a 3xx status code
func (o *AccessProfilesLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles lock manager o k response has a 4xx status code
func (o *AccessProfilesLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles lock manager o k response has a 5xx status code
func (o *AccessProfilesLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles lock manager o k response a status code equal to that given
func (o *AccessProfilesLockManagerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the access profiles lock manager o k response
func (o *AccessProfilesLockManagerOK) Code() int {
	return 200
}

func (o *AccessProfilesLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesLockManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AccessProfilesLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesLockManagerBadRequest creates a AccessProfilesLockManagerBadRequest with default headers values
func NewAccessProfilesLockManagerBadRequest() *AccessProfilesLockManagerBadRequest {
	return &AccessProfilesLockManagerBadRequest{}
}

/*
AccessProfilesLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AccessProfilesLockManagerBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this access profiles lock manager bad request response has a 2xx status code
func (o *AccessProfilesLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles lock manager bad request response has a 3xx status code
func (o *AccessProfilesLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles lock manager bad request response has a 4xx status code
func (o *AccessProfilesLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles lock manager bad request response has a 5xx status code
func (o *AccessProfilesLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles lock manager bad request response a status code equal to that given
func (o *AccessProfilesLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the access profiles lock manager bad request response
func (o *AccessProfilesLockManagerBadRequest) Code() int {
	return 400
}

func (o *AccessProfilesLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesLockManagerBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessProfilesLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesLockManagerUnauthorized creates a AccessProfilesLockManagerUnauthorized with default headers values
func NewAccessProfilesLockManagerUnauthorized() *AccessProfilesLockManagerUnauthorized {
	return &AccessProfilesLockManagerUnauthorized{}
}

/*
AccessProfilesLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AccessProfilesLockManagerUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this access profiles lock manager unauthorized response has a 2xx status code
func (o *AccessProfilesLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles lock manager unauthorized response has a 3xx status code
func (o *AccessProfilesLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles lock manager unauthorized response has a 4xx status code
func (o *AccessProfilesLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles lock manager unauthorized response has a 5xx status code
func (o *AccessProfilesLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles lock manager unauthorized response a status code equal to that given
func (o *AccessProfilesLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the access profiles lock manager unauthorized response
func (o *AccessProfilesLockManagerUnauthorized) Code() int {
	return 401
}

func (o *AccessProfilesLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesLockManagerUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessProfilesLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesLockManagerForbidden creates a AccessProfilesLockManagerForbidden with default headers values
func NewAccessProfilesLockManagerForbidden() *AccessProfilesLockManagerForbidden {
	return &AccessProfilesLockManagerForbidden{}
}

/*
AccessProfilesLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccessProfilesLockManagerForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this access profiles lock manager forbidden response has a 2xx status code
func (o *AccessProfilesLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles lock manager forbidden response has a 3xx status code
func (o *AccessProfilesLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles lock manager forbidden response has a 4xx status code
func (o *AccessProfilesLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles lock manager forbidden response has a 5xx status code
func (o *AccessProfilesLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles lock manager forbidden response a status code equal to that given
func (o *AccessProfilesLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the access profiles lock manager forbidden response
func (o *AccessProfilesLockManagerForbidden) Code() int {
	return 403
}

func (o *AccessProfilesLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesLockManagerForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessProfilesLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesLockManagerNotFound creates a AccessProfilesLockManagerNotFound with default headers values
func NewAccessProfilesLockManagerNotFound() *AccessProfilesLockManagerNotFound {
	return &AccessProfilesLockManagerNotFound{}
}

/*
AccessProfilesLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AccessProfilesLockManagerNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this access profiles lock manager not found response has a 2xx status code
func (o *AccessProfilesLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles lock manager not found response has a 3xx status code
func (o *AccessProfilesLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles lock manager not found response has a 4xx status code
func (o *AccessProfilesLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles lock manager not found response has a 5xx status code
func (o *AccessProfilesLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles lock manager not found response a status code equal to that given
func (o *AccessProfilesLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the access profiles lock manager not found response
func (o *AccessProfilesLockManagerNotFound) Code() int {
	return 404
}

func (o *AccessProfilesLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesLockManagerNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessProfilesLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesLockManagerInternalServerError creates a AccessProfilesLockManagerInternalServerError with default headers values
func NewAccessProfilesLockManagerInternalServerError() *AccessProfilesLockManagerInternalServerError {
	return &AccessProfilesLockManagerInternalServerError{}
}

/*
AccessProfilesLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AccessProfilesLockManagerInternalServerError struct {
}

// IsSuccess returns true when this access profiles lock manager internal server error response has a 2xx status code
func (o *AccessProfilesLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles lock manager internal server error response has a 3xx status code
func (o *AccessProfilesLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles lock manager internal server error response has a 4xx status code
func (o *AccessProfilesLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles lock manager internal server error response has a 5xx status code
func (o *AccessProfilesLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access profiles lock manager internal server error response a status code equal to that given
func (o *AccessProfilesLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the access profiles lock manager internal server error response
func (o *AccessProfilesLockManagerInternalServerError) Code() int {
	return 500
}

func (o *AccessProfilesLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerInternalServerError ", 500)
}

func (o *AccessProfilesLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/lockmanager][%d] accessProfilesLockManagerInternalServerError ", 500)
}

func (o *AccessProfilesLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
