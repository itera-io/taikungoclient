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

// RepositoryDeleteReader is a Reader for the RepositoryDelete structure.
type RepositoryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepositoryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepositoryDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepositoryDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRepositoryDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRepositoryDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepositoryDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRepositoryDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRepositoryDeleteOK creates a RepositoryDeleteOK with default headers values
func NewRepositoryDeleteOK() *RepositoryDeleteOK {
	return &RepositoryDeleteOK{}
}

/*
RepositoryDeleteOK describes a response with status code 200, with default header values.

Success
*/
type RepositoryDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this repository delete o k response has a 2xx status code
func (o *RepositoryDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repository delete o k response has a 3xx status code
func (o *RepositoryDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository delete o k response has a 4xx status code
func (o *RepositoryDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repository delete o k response has a 5xx status code
func (o *RepositoryDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repository delete o k response a status code equal to that given
func (o *RepositoryDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *RepositoryDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteOK  %+v", 200, o.Payload)
}

func (o *RepositoryDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteOK  %+v", 200, o.Payload)
}

func (o *RepositoryDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *RepositoryDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryDeleteBadRequest creates a RepositoryDeleteBadRequest with default headers values
func NewRepositoryDeleteBadRequest() *RepositoryDeleteBadRequest {
	return &RepositoryDeleteBadRequest{}
}

/*
RepositoryDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RepositoryDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this repository delete bad request response has a 2xx status code
func (o *RepositoryDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository delete bad request response has a 3xx status code
func (o *RepositoryDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository delete bad request response has a 4xx status code
func (o *RepositoryDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository delete bad request response has a 5xx status code
func (o *RepositoryDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repository delete bad request response a status code equal to that given
func (o *RepositoryDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RepositoryDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *RepositoryDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *RepositoryDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *RepositoryDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryDeleteUnauthorized creates a RepositoryDeleteUnauthorized with default headers values
func NewRepositoryDeleteUnauthorized() *RepositoryDeleteUnauthorized {
	return &RepositoryDeleteUnauthorized{}
}

/*
RepositoryDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RepositoryDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this repository delete unauthorized response has a 2xx status code
func (o *RepositoryDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository delete unauthorized response has a 3xx status code
func (o *RepositoryDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository delete unauthorized response has a 4xx status code
func (o *RepositoryDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository delete unauthorized response has a 5xx status code
func (o *RepositoryDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this repository delete unauthorized response a status code equal to that given
func (o *RepositoryDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RepositoryDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *RepositoryDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *RepositoryDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RepositoryDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryDeleteForbidden creates a RepositoryDeleteForbidden with default headers values
func NewRepositoryDeleteForbidden() *RepositoryDeleteForbidden {
	return &RepositoryDeleteForbidden{}
}

/*
RepositoryDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RepositoryDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this repository delete forbidden response has a 2xx status code
func (o *RepositoryDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository delete forbidden response has a 3xx status code
func (o *RepositoryDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository delete forbidden response has a 4xx status code
func (o *RepositoryDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository delete forbidden response has a 5xx status code
func (o *RepositoryDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repository delete forbidden response a status code equal to that given
func (o *RepositoryDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RepositoryDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteForbidden  %+v", 403, o.Payload)
}

func (o *RepositoryDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteForbidden  %+v", 403, o.Payload)
}

func (o *RepositoryDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RepositoryDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryDeleteNotFound creates a RepositoryDeleteNotFound with default headers values
func NewRepositoryDeleteNotFound() *RepositoryDeleteNotFound {
	return &RepositoryDeleteNotFound{}
}

/*
RepositoryDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RepositoryDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this repository delete not found response has a 2xx status code
func (o *RepositoryDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository delete not found response has a 3xx status code
func (o *RepositoryDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository delete not found response has a 4xx status code
func (o *RepositoryDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repository delete not found response has a 5xx status code
func (o *RepositoryDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repository delete not found response a status code equal to that given
func (o *RepositoryDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RepositoryDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteNotFound  %+v", 404, o.Payload)
}

func (o *RepositoryDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteNotFound  %+v", 404, o.Payload)
}

func (o *RepositoryDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RepositoryDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepositoryDeleteInternalServerError creates a RepositoryDeleteInternalServerError with default headers values
func NewRepositoryDeleteInternalServerError() *RepositoryDeleteInternalServerError {
	return &RepositoryDeleteInternalServerError{}
}

/*
RepositoryDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type RepositoryDeleteInternalServerError struct {
}

// IsSuccess returns true when this repository delete internal server error response has a 2xx status code
func (o *RepositoryDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repository delete internal server error response has a 3xx status code
func (o *RepositoryDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repository delete internal server error response has a 4xx status code
func (o *RepositoryDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this repository delete internal server error response has a 5xx status code
func (o *RepositoryDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this repository delete internal server error response a status code equal to that given
func (o *RepositoryDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RepositoryDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteInternalServerError ", 500)
}

func (o *RepositoryDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Repository/delete][%d] repositoryDeleteInternalServerError ", 500)
}

func (o *RepositoryDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
