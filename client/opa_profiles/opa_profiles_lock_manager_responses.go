// Code generated by go-swagger; DO NOT EDIT.

package opa_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpaProfilesLockManagerReader is a Reader for the OpaProfilesLockManager structure.
type OpaProfilesLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesLockManagerOK creates a OpaProfilesLockManagerOK with default headers values
func NewOpaProfilesLockManagerOK() *OpaProfilesLockManagerOK {
	return &OpaProfilesLockManagerOK{}
}

/*
OpaProfilesLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesLockManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this opa profiles lock manager o k response has a 2xx status code
func (o *OpaProfilesLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this opa profiles lock manager o k response has a 3xx status code
func (o *OpaProfilesLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles lock manager o k response has a 4xx status code
func (o *OpaProfilesLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles lock manager o k response has a 5xx status code
func (o *OpaProfilesLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles lock manager o k response a status code equal to that given
func (o *OpaProfilesLockManagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpaProfilesLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesLockManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpaProfilesLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesLockManagerBadRequest creates a OpaProfilesLockManagerBadRequest with default headers values
func NewOpaProfilesLockManagerBadRequest() *OpaProfilesLockManagerBadRequest {
	return &OpaProfilesLockManagerBadRequest{}
}

/*
OpaProfilesLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesLockManagerBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this opa profiles lock manager bad request response has a 2xx status code
func (o *OpaProfilesLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles lock manager bad request response has a 3xx status code
func (o *OpaProfilesLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles lock manager bad request response has a 4xx status code
func (o *OpaProfilesLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles lock manager bad request response has a 5xx status code
func (o *OpaProfilesLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles lock manager bad request response a status code equal to that given
func (o *OpaProfilesLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpaProfilesLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesLockManagerBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpaProfilesLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesLockManagerUnauthorized creates a OpaProfilesLockManagerUnauthorized with default headers values
func NewOpaProfilesLockManagerUnauthorized() *OpaProfilesLockManagerUnauthorized {
	return &OpaProfilesLockManagerUnauthorized{}
}

/*
OpaProfilesLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesLockManagerUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this opa profiles lock manager unauthorized response has a 2xx status code
func (o *OpaProfilesLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles lock manager unauthorized response has a 3xx status code
func (o *OpaProfilesLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles lock manager unauthorized response has a 4xx status code
func (o *OpaProfilesLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles lock manager unauthorized response has a 5xx status code
func (o *OpaProfilesLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles lock manager unauthorized response a status code equal to that given
func (o *OpaProfilesLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpaProfilesLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesLockManagerUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpaProfilesLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesLockManagerForbidden creates a OpaProfilesLockManagerForbidden with default headers values
func NewOpaProfilesLockManagerForbidden() *OpaProfilesLockManagerForbidden {
	return &OpaProfilesLockManagerForbidden{}
}

/*
OpaProfilesLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesLockManagerForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this opa profiles lock manager forbidden response has a 2xx status code
func (o *OpaProfilesLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles lock manager forbidden response has a 3xx status code
func (o *OpaProfilesLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles lock manager forbidden response has a 4xx status code
func (o *OpaProfilesLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles lock manager forbidden response has a 5xx status code
func (o *OpaProfilesLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles lock manager forbidden response a status code equal to that given
func (o *OpaProfilesLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpaProfilesLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesLockManagerForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpaProfilesLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesLockManagerNotFound creates a OpaProfilesLockManagerNotFound with default headers values
func NewOpaProfilesLockManagerNotFound() *OpaProfilesLockManagerNotFound {
	return &OpaProfilesLockManagerNotFound{}
}

/*
OpaProfilesLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesLockManagerNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this opa profiles lock manager not found response has a 2xx status code
func (o *OpaProfilesLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles lock manager not found response has a 3xx status code
func (o *OpaProfilesLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles lock manager not found response has a 4xx status code
func (o *OpaProfilesLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles lock manager not found response has a 5xx status code
func (o *OpaProfilesLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles lock manager not found response a status code equal to that given
func (o *OpaProfilesLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpaProfilesLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesLockManagerNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpaProfilesLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesLockManagerInternalServerError creates a OpaProfilesLockManagerInternalServerError with default headers values
func NewOpaProfilesLockManagerInternalServerError() *OpaProfilesLockManagerInternalServerError {
	return &OpaProfilesLockManagerInternalServerError{}
}

/*
OpaProfilesLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesLockManagerInternalServerError struct {
}

// IsSuccess returns true when this opa profiles lock manager internal server error response has a 2xx status code
func (o *OpaProfilesLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles lock manager internal server error response has a 3xx status code
func (o *OpaProfilesLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles lock manager internal server error response has a 4xx status code
func (o *OpaProfilesLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles lock manager internal server error response has a 5xx status code
func (o *OpaProfilesLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this opa profiles lock manager internal server error response a status code equal to that given
func (o *OpaProfilesLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpaProfilesLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerInternalServerError ", 500)
}

func (o *OpaProfilesLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/lockmanager][%d] opaProfilesLockManagerInternalServerError ", 500)
}

func (o *OpaProfilesLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
