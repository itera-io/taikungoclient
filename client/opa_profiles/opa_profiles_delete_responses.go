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

// OpaProfilesDeleteReader is a Reader for the OpaProfilesDelete structure.
type OpaProfilesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesDeleteOK creates a OpaProfilesDeleteOK with default headers values
func NewOpaProfilesDeleteOK() *OpaProfilesDeleteOK {
	return &OpaProfilesDeleteOK{}
}

/*
OpaProfilesDeleteOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this opa profiles delete o k response has a 2xx status code
func (o *OpaProfilesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this opa profiles delete o k response has a 3xx status code
func (o *OpaProfilesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles delete o k response has a 4xx status code
func (o *OpaProfilesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles delete o k response has a 5xx status code
func (o *OpaProfilesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles delete o k response a status code equal to that given
func (o *OpaProfilesDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the opa profiles delete o k response
func (o *OpaProfilesDeleteOK) Code() int {
	return 200
}

func (o *OpaProfilesDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpaProfilesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteBadRequest creates a OpaProfilesDeleteBadRequest with default headers values
func NewOpaProfilesDeleteBadRequest() *OpaProfilesDeleteBadRequest {
	return &OpaProfilesDeleteBadRequest{}
}

/*
OpaProfilesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesDeleteBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles delete bad request response has a 2xx status code
func (o *OpaProfilesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles delete bad request response has a 3xx status code
func (o *OpaProfilesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles delete bad request response has a 4xx status code
func (o *OpaProfilesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles delete bad request response has a 5xx status code
func (o *OpaProfilesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles delete bad request response a status code equal to that given
func (o *OpaProfilesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the opa profiles delete bad request response
func (o *OpaProfilesDeleteBadRequest) Code() int {
	return 400
}

func (o *OpaProfilesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesDeleteBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteUnauthorized creates a OpaProfilesDeleteUnauthorized with default headers values
func NewOpaProfilesDeleteUnauthorized() *OpaProfilesDeleteUnauthorized {
	return &OpaProfilesDeleteUnauthorized{}
}

/*
OpaProfilesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles delete unauthorized response has a 2xx status code
func (o *OpaProfilesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles delete unauthorized response has a 3xx status code
func (o *OpaProfilesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles delete unauthorized response has a 4xx status code
func (o *OpaProfilesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles delete unauthorized response has a 5xx status code
func (o *OpaProfilesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles delete unauthorized response a status code equal to that given
func (o *OpaProfilesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the opa profiles delete unauthorized response
func (o *OpaProfilesDeleteUnauthorized) Code() int {
	return 401
}

func (o *OpaProfilesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteForbidden creates a OpaProfilesDeleteForbidden with default headers values
func NewOpaProfilesDeleteForbidden() *OpaProfilesDeleteForbidden {
	return &OpaProfilesDeleteForbidden{}
}

/*
OpaProfilesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles delete forbidden response has a 2xx status code
func (o *OpaProfilesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles delete forbidden response has a 3xx status code
func (o *OpaProfilesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles delete forbidden response has a 4xx status code
func (o *OpaProfilesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles delete forbidden response has a 5xx status code
func (o *OpaProfilesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles delete forbidden response a status code equal to that given
func (o *OpaProfilesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the opa profiles delete forbidden response
func (o *OpaProfilesDeleteForbidden) Code() int {
	return 403
}

func (o *OpaProfilesDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteNotFound creates a OpaProfilesDeleteNotFound with default headers values
func NewOpaProfilesDeleteNotFound() *OpaProfilesDeleteNotFound {
	return &OpaProfilesDeleteNotFound{}
}

/*
OpaProfilesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles delete not found response has a 2xx status code
func (o *OpaProfilesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles delete not found response has a 3xx status code
func (o *OpaProfilesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles delete not found response has a 4xx status code
func (o *OpaProfilesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles delete not found response has a 5xx status code
func (o *OpaProfilesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles delete not found response a status code equal to that given
func (o *OpaProfilesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the opa profiles delete not found response
func (o *OpaProfilesDeleteNotFound) Code() int {
	return 404
}

func (o *OpaProfilesDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteInternalServerError creates a OpaProfilesDeleteInternalServerError with default headers values
func NewOpaProfilesDeleteInternalServerError() *OpaProfilesDeleteInternalServerError {
	return &OpaProfilesDeleteInternalServerError{}
}

/*
OpaProfilesDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesDeleteInternalServerError struct {
}

// IsSuccess returns true when this opa profiles delete internal server error response has a 2xx status code
func (o *OpaProfilesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles delete internal server error response has a 3xx status code
func (o *OpaProfilesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles delete internal server error response has a 4xx status code
func (o *OpaProfilesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles delete internal server error response has a 5xx status code
func (o *OpaProfilesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this opa profiles delete internal server error response a status code equal to that given
func (o *OpaProfilesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the opa profiles delete internal server error response
func (o *OpaProfilesDeleteInternalServerError) Code() int {
	return 500
}

func (o *OpaProfilesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteInternalServerError ", 500)
}

func (o *OpaProfilesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteInternalServerError ", 500)
}

func (o *OpaProfilesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
