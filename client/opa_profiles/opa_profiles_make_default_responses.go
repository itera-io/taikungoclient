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

// OpaProfilesMakeDefaultReader is a Reader for the OpaProfilesMakeDefault structure.
type OpaProfilesMakeDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesMakeDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesMakeDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesMakeDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesMakeDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesMakeDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesMakeDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesMakeDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesMakeDefaultOK creates a OpaProfilesMakeDefaultOK with default headers values
func NewOpaProfilesMakeDefaultOK() *OpaProfilesMakeDefaultOK {
	return &OpaProfilesMakeDefaultOK{}
}

/*
OpaProfilesMakeDefaultOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesMakeDefaultOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this opa profiles make default o k response has a 2xx status code
func (o *OpaProfilesMakeDefaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this opa profiles make default o k response has a 3xx status code
func (o *OpaProfilesMakeDefaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles make default o k response has a 4xx status code
func (o *OpaProfilesMakeDefaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles make default o k response has a 5xx status code
func (o *OpaProfilesMakeDefaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles make default o k response a status code equal to that given
func (o *OpaProfilesMakeDefaultOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpaProfilesMakeDefaultOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesMakeDefaultOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesMakeDefaultOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpaProfilesMakeDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesMakeDefaultBadRequest creates a OpaProfilesMakeDefaultBadRequest with default headers values
func NewOpaProfilesMakeDefaultBadRequest() *OpaProfilesMakeDefaultBadRequest {
	return &OpaProfilesMakeDefaultBadRequest{}
}

/*
OpaProfilesMakeDefaultBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesMakeDefaultBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this opa profiles make default bad request response has a 2xx status code
func (o *OpaProfilesMakeDefaultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles make default bad request response has a 3xx status code
func (o *OpaProfilesMakeDefaultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles make default bad request response has a 4xx status code
func (o *OpaProfilesMakeDefaultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles make default bad request response has a 5xx status code
func (o *OpaProfilesMakeDefaultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles make default bad request response a status code equal to that given
func (o *OpaProfilesMakeDefaultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpaProfilesMakeDefaultBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesMakeDefaultBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesMakeDefaultBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *OpaProfilesMakeDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesMakeDefaultUnauthorized creates a OpaProfilesMakeDefaultUnauthorized with default headers values
func NewOpaProfilesMakeDefaultUnauthorized() *OpaProfilesMakeDefaultUnauthorized {
	return &OpaProfilesMakeDefaultUnauthorized{}
}

/*
OpaProfilesMakeDefaultUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesMakeDefaultUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles make default unauthorized response has a 2xx status code
func (o *OpaProfilesMakeDefaultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles make default unauthorized response has a 3xx status code
func (o *OpaProfilesMakeDefaultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles make default unauthorized response has a 4xx status code
func (o *OpaProfilesMakeDefaultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles make default unauthorized response has a 5xx status code
func (o *OpaProfilesMakeDefaultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles make default unauthorized response a status code equal to that given
func (o *OpaProfilesMakeDefaultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpaProfilesMakeDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesMakeDefaultUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesMakeDefaultUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesMakeDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesMakeDefaultForbidden creates a OpaProfilesMakeDefaultForbidden with default headers values
func NewOpaProfilesMakeDefaultForbidden() *OpaProfilesMakeDefaultForbidden {
	return &OpaProfilesMakeDefaultForbidden{}
}

/*
OpaProfilesMakeDefaultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesMakeDefaultForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles make default forbidden response has a 2xx status code
func (o *OpaProfilesMakeDefaultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles make default forbidden response has a 3xx status code
func (o *OpaProfilesMakeDefaultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles make default forbidden response has a 4xx status code
func (o *OpaProfilesMakeDefaultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles make default forbidden response has a 5xx status code
func (o *OpaProfilesMakeDefaultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles make default forbidden response a status code equal to that given
func (o *OpaProfilesMakeDefaultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpaProfilesMakeDefaultForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesMakeDefaultForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesMakeDefaultForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesMakeDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesMakeDefaultNotFound creates a OpaProfilesMakeDefaultNotFound with default headers values
func NewOpaProfilesMakeDefaultNotFound() *OpaProfilesMakeDefaultNotFound {
	return &OpaProfilesMakeDefaultNotFound{}
}

/*
OpaProfilesMakeDefaultNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesMakeDefaultNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this opa profiles make default not found response has a 2xx status code
func (o *OpaProfilesMakeDefaultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles make default not found response has a 3xx status code
func (o *OpaProfilesMakeDefaultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles make default not found response has a 4xx status code
func (o *OpaProfilesMakeDefaultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles make default not found response has a 5xx status code
func (o *OpaProfilesMakeDefaultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles make default not found response a status code equal to that given
func (o *OpaProfilesMakeDefaultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpaProfilesMakeDefaultNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesMakeDefaultNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesMakeDefaultNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesMakeDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesMakeDefaultInternalServerError creates a OpaProfilesMakeDefaultInternalServerError with default headers values
func NewOpaProfilesMakeDefaultInternalServerError() *OpaProfilesMakeDefaultInternalServerError {
	return &OpaProfilesMakeDefaultInternalServerError{}
}

/*
OpaProfilesMakeDefaultInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesMakeDefaultInternalServerError struct {
}

// IsSuccess returns true when this opa profiles make default internal server error response has a 2xx status code
func (o *OpaProfilesMakeDefaultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles make default internal server error response has a 3xx status code
func (o *OpaProfilesMakeDefaultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles make default internal server error response has a 4xx status code
func (o *OpaProfilesMakeDefaultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles make default internal server error response has a 5xx status code
func (o *OpaProfilesMakeDefaultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this opa profiles make default internal server error response a status code equal to that given
func (o *OpaProfilesMakeDefaultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpaProfilesMakeDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultInternalServerError ", 500)
}

func (o *OpaProfilesMakeDefaultInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/make-default][%d] opaProfilesMakeDefaultInternalServerError ", 500)
}

func (o *OpaProfilesMakeDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
