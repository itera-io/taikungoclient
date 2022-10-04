// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// RepositoryCreateReader is a Reader for the RepositoryCreate structure.
type RepositoryCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepositoryCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepositoryCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepositoryCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRepositoryCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRepositoryCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepositoryCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRepositoryCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRepositoryCreateOK creates a RepositoryCreateOK with default headers values
func NewRepositoryCreateOK() *RepositoryCreateOK {
	return &RepositoryCreateOK{}
}

/*
RepositoryCreateOK describes a response with status code 200, with default header values.

Success
*/
type RepositoryCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this repository create o k response has a 2xx status code
func (o *RepositoryCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repository create o k response has a 3xx status code
func (o *RepositoryCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository create o k response has a 4xx status code
func (o *RepositoryCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repository create o k response has a 5xx status code
func (o *RepositoryCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repository create o k response a status code equal to that given
func (o *RepositoryCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *RepositoryCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateOK  %+v", 200, o.Payload)
}

func (o *RepositoryCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateOK  %+v", 200, o.Payload)
}

func (o *RepositoryCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *RepositoryCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryCreateBadRequest creates a RepositoryCreateBadRequest with default headers values
func NewRepositoryCreateBadRequest() *RepositoryCreateBadRequest {
	return &RepositoryCreateBadRequest{}
}

/*
RepositoryCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RepositoryCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this repository create bad request response has a 2xx status code
func (o *RepositoryCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository create bad request response has a 3xx status code
func (o *RepositoryCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository create bad request response has a 4xx status code
func (o *RepositoryCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository create bad request response has a 5xx status code
func (o *RepositoryCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repository create bad request response a status code equal to that given
func (o *RepositoryCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RepositoryCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateBadRequest  %+v", 400, o.Payload)
}

func (o *RepositoryCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateBadRequest  %+v", 400, o.Payload)
}

func (o *RepositoryCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *RepositoryCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryCreateUnauthorized creates a RepositoryCreateUnauthorized with default headers values
func NewRepositoryCreateUnauthorized() *RepositoryCreateUnauthorized {
	return &RepositoryCreateUnauthorized{}
}

/*
RepositoryCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RepositoryCreateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this repository create unauthorized response has a 2xx status code
func (o *RepositoryCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository create unauthorized response has a 3xx status code
func (o *RepositoryCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository create unauthorized response has a 4xx status code
func (o *RepositoryCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository create unauthorized response has a 5xx status code
func (o *RepositoryCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this repository create unauthorized response a status code equal to that given
func (o *RepositoryCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RepositoryCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *RepositoryCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *RepositoryCreateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *RepositoryCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryCreateForbidden creates a RepositoryCreateForbidden with default headers values
func NewRepositoryCreateForbidden() *RepositoryCreateForbidden {
	return &RepositoryCreateForbidden{}
}

/*
RepositoryCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RepositoryCreateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this repository create forbidden response has a 2xx status code
func (o *RepositoryCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository create forbidden response has a 3xx status code
func (o *RepositoryCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository create forbidden response has a 4xx status code
func (o *RepositoryCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository create forbidden response has a 5xx status code
func (o *RepositoryCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repository create forbidden response a status code equal to that given
func (o *RepositoryCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RepositoryCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateForbidden  %+v", 403, o.Payload)
}

func (o *RepositoryCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateForbidden  %+v", 403, o.Payload)
}

func (o *RepositoryCreateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *RepositoryCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryCreateNotFound creates a RepositoryCreateNotFound with default headers values
func NewRepositoryCreateNotFound() *RepositoryCreateNotFound {
	return &RepositoryCreateNotFound{}
}

/*
RepositoryCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RepositoryCreateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this repository create not found response has a 2xx status code
func (o *RepositoryCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository create not found response has a 3xx status code
func (o *RepositoryCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository create not found response has a 4xx status code
func (o *RepositoryCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository create not found response has a 5xx status code
func (o *RepositoryCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repository create not found response a status code equal to that given
func (o *RepositoryCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RepositoryCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateNotFound  %+v", 404, o.Payload)
}

func (o *RepositoryCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateNotFound  %+v", 404, o.Payload)
}

func (o *RepositoryCreateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *RepositoryCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryCreateInternalServerError creates a RepositoryCreateInternalServerError with default headers values
func NewRepositoryCreateInternalServerError() *RepositoryCreateInternalServerError {
	return &RepositoryCreateInternalServerError{}
}

/*
RepositoryCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type RepositoryCreateInternalServerError struct {
}

// IsSuccess returns true when this repository create internal server error response has a 2xx status code
func (o *RepositoryCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository create internal server error response has a 3xx status code
func (o *RepositoryCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository create internal server error response has a 4xx status code
func (o *RepositoryCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this repository create internal server error response has a 5xx status code
func (o *RepositoryCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this repository create internal server error response a status code equal to that given
func (o *RepositoryCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RepositoryCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateInternalServerError ", 500)
}

func (o *RepositoryCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/bind][%d] repositoryCreateInternalServerError ", 500)
}

func (o *RepositoryCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
