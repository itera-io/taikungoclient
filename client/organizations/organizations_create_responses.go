// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OrganizationsCreateReader is a Reader for the OrganizationsCreate structure.
type OrganizationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsCreateOK creates a OrganizationsCreateOK with default headers values
func NewOrganizationsCreateOK() *OrganizationsCreateOK {
	return &OrganizationsCreateOK{}
}

/*
OrganizationsCreateOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this organizations create o k response has a 2xx status code
func (o *OrganizationsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations create o k response has a 3xx status code
func (o *OrganizationsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create o k response has a 4xx status code
func (o *OrganizationsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations create o k response has a 5xx status code
func (o *OrganizationsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create o k response a status code equal to that given
func (o *OrganizationsCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateOK  %+v", 200, o.Payload)
}

func (o *OrganizationsCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateOK  %+v", 200, o.Payload)
}

func (o *OrganizationsCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *OrganizationsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateBadRequest creates a OrganizationsCreateBadRequest with default headers values
func NewOrganizationsCreateBadRequest() *OrganizationsCreateBadRequest {
	return &OrganizationsCreateBadRequest{}
}

/*
OrganizationsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this organizations create bad request response has a 2xx status code
func (o *OrganizationsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create bad request response has a 3xx status code
func (o *OrganizationsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create bad request response has a 4xx status code
func (o *OrganizationsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create bad request response has a 5xx status code
func (o *OrganizationsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create bad request response a status code equal to that given
func (o *OrganizationsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OrganizationsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateUnauthorized creates a OrganizationsCreateUnauthorized with default headers values
func NewOrganizationsCreateUnauthorized() *OrganizationsCreateUnauthorized {
	return &OrganizationsCreateUnauthorized{}
}

/*
OrganizationsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations create unauthorized response has a 2xx status code
func (o *OrganizationsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create unauthorized response has a 3xx status code
func (o *OrganizationsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create unauthorized response has a 4xx status code
func (o *OrganizationsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create unauthorized response has a 5xx status code
func (o *OrganizationsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create unauthorized response a status code equal to that given
func (o *OrganizationsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OrganizationsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateForbidden creates a OrganizationsCreateForbidden with default headers values
func NewOrganizationsCreateForbidden() *OrganizationsCreateForbidden {
	return &OrganizationsCreateForbidden{}
}

/*
OrganizationsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations create forbidden response has a 2xx status code
func (o *OrganizationsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create forbidden response has a 3xx status code
func (o *OrganizationsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create forbidden response has a 4xx status code
func (o *OrganizationsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create forbidden response has a 5xx status code
func (o *OrganizationsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create forbidden response a status code equal to that given
func (o *OrganizationsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateNotFound creates a OrganizationsCreateNotFound with default headers values
func NewOrganizationsCreateNotFound() *OrganizationsCreateNotFound {
	return &OrganizationsCreateNotFound{}
}

/*
OrganizationsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations create not found response has a 2xx status code
func (o *OrganizationsCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create not found response has a 3xx status code
func (o *OrganizationsCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create not found response has a 4xx status code
func (o *OrganizationsCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations create not found response has a 5xx status code
func (o *OrganizationsCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations create not found response a status code equal to that given
func (o *OrganizationsCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OrganizationsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsCreateInternalServerError creates a OrganizationsCreateInternalServerError with default headers values
func NewOrganizationsCreateInternalServerError() *OrganizationsCreateInternalServerError {
	return &OrganizationsCreateInternalServerError{}
}

/*
OrganizationsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsCreateInternalServerError struct {
}

// IsSuccess returns true when this organizations create internal server error response has a 2xx status code
func (o *OrganizationsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations create internal server error response has a 3xx status code
func (o *OrganizationsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations create internal server error response has a 4xx status code
func (o *OrganizationsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations create internal server error response has a 5xx status code
func (o *OrganizationsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations create internal server error response a status code equal to that given
func (o *OrganizationsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OrganizationsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateInternalServerError ", 500)
}

func (o *OrganizationsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations][%d] organizationsCreateInternalServerError ", 500)
}

func (o *OrganizationsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
