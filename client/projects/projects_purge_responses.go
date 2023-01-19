// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectsPurgeReader is a Reader for the ProjectsPurge structure.
type ProjectsPurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsPurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsPurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsPurgeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsPurgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsPurgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsPurgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsPurgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsPurgeOK creates a ProjectsPurgeOK with default headers values
func NewProjectsPurgeOK() *ProjectsPurgeOK {
	return &ProjectsPurgeOK{}
}

/*
ProjectsPurgeOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsPurgeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects purge o k response has a 2xx status code
func (o *ProjectsPurgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects purge o k response has a 3xx status code
func (o *ProjectsPurgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects purge o k response has a 4xx status code
func (o *ProjectsPurgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects purge o k response has a 5xx status code
func (o *ProjectsPurgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects purge o k response a status code equal to that given
func (o *ProjectsPurgeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the projects purge o k response
func (o *ProjectsPurgeOK) Code() int {
	return 200
}

func (o *ProjectsPurgeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeOK  %+v", 200, o.Payload)
}

func (o *ProjectsPurgeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeOK  %+v", 200, o.Payload)
}

func (o *ProjectsPurgeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsPurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeBadRequest creates a ProjectsPurgeBadRequest with default headers values
func NewProjectsPurgeBadRequest() *ProjectsPurgeBadRequest {
	return &ProjectsPurgeBadRequest{}
}

/*
ProjectsPurgeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsPurgeBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects purge bad request response has a 2xx status code
func (o *ProjectsPurgeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects purge bad request response has a 3xx status code
func (o *ProjectsPurgeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects purge bad request response has a 4xx status code
func (o *ProjectsPurgeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects purge bad request response has a 5xx status code
func (o *ProjectsPurgeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects purge bad request response a status code equal to that given
func (o *ProjectsPurgeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the projects purge bad request response
func (o *ProjectsPurgeBadRequest) Code() int {
	return 400
}

func (o *ProjectsPurgeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsPurgeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsPurgeBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeUnauthorized creates a ProjectsPurgeUnauthorized with default headers values
func NewProjectsPurgeUnauthorized() *ProjectsPurgeUnauthorized {
	return &ProjectsPurgeUnauthorized{}
}

/*
ProjectsPurgeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsPurgeUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects purge unauthorized response has a 2xx status code
func (o *ProjectsPurgeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects purge unauthorized response has a 3xx status code
func (o *ProjectsPurgeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects purge unauthorized response has a 4xx status code
func (o *ProjectsPurgeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects purge unauthorized response has a 5xx status code
func (o *ProjectsPurgeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects purge unauthorized response a status code equal to that given
func (o *ProjectsPurgeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the projects purge unauthorized response
func (o *ProjectsPurgeUnauthorized) Code() int {
	return 401
}

func (o *ProjectsPurgeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsPurgeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsPurgeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeForbidden creates a ProjectsPurgeForbidden with default headers values
func NewProjectsPurgeForbidden() *ProjectsPurgeForbidden {
	return &ProjectsPurgeForbidden{}
}

/*
ProjectsPurgeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsPurgeForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects purge forbidden response has a 2xx status code
func (o *ProjectsPurgeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects purge forbidden response has a 3xx status code
func (o *ProjectsPurgeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects purge forbidden response has a 4xx status code
func (o *ProjectsPurgeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects purge forbidden response has a 5xx status code
func (o *ProjectsPurgeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects purge forbidden response a status code equal to that given
func (o *ProjectsPurgeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the projects purge forbidden response
func (o *ProjectsPurgeForbidden) Code() int {
	return 403
}

func (o *ProjectsPurgeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsPurgeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsPurgeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeNotFound creates a ProjectsPurgeNotFound with default headers values
func NewProjectsPurgeNotFound() *ProjectsPurgeNotFound {
	return &ProjectsPurgeNotFound{}
}

/*
ProjectsPurgeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsPurgeNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects purge not found response has a 2xx status code
func (o *ProjectsPurgeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects purge not found response has a 3xx status code
func (o *ProjectsPurgeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects purge not found response has a 4xx status code
func (o *ProjectsPurgeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects purge not found response has a 5xx status code
func (o *ProjectsPurgeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects purge not found response a status code equal to that given
func (o *ProjectsPurgeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the projects purge not found response
func (o *ProjectsPurgeNotFound) Code() int {
	return 404
}

func (o *ProjectsPurgeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsPurgeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsPurgeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeInternalServerError creates a ProjectsPurgeInternalServerError with default headers values
func NewProjectsPurgeInternalServerError() *ProjectsPurgeInternalServerError {
	return &ProjectsPurgeInternalServerError{}
}

/*
ProjectsPurgeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsPurgeInternalServerError struct {
}

// IsSuccess returns true when this projects purge internal server error response has a 2xx status code
func (o *ProjectsPurgeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects purge internal server error response has a 3xx status code
func (o *ProjectsPurgeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects purge internal server error response has a 4xx status code
func (o *ProjectsPurgeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects purge internal server error response has a 5xx status code
func (o *ProjectsPurgeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects purge internal server error response a status code equal to that given
func (o *ProjectsPurgeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the projects purge internal server error response
func (o *ProjectsPurgeInternalServerError) Code() int {
	return 500
}

func (o *ProjectsPurgeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeInternalServerError ", 500)
}

func (o *ProjectsPurgeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purge][%d] projectsPurgeInternalServerError ", 500)
}

func (o *ProjectsPurgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
