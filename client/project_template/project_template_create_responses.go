// Code generated by go-swagger; DO NOT EDIT.

package project_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectTemplateCreateReader is a Reader for the ProjectTemplateCreate structure.
type ProjectTemplateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectTemplateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectTemplateCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectTemplateCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectTemplateCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectTemplateCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectTemplateCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectTemplateCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectTemplateCreateOK creates a ProjectTemplateCreateOK with default headers values
func NewProjectTemplateCreateOK() *ProjectTemplateCreateOK {
	return &ProjectTemplateCreateOK{}
}

/*
ProjectTemplateCreateOK describes a response with status code 200, with default header values.

Success
*/
type ProjectTemplateCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project template create o k response has a 2xx status code
func (o *ProjectTemplateCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project template create o k response has a 3xx status code
func (o *ProjectTemplateCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template create o k response has a 4xx status code
func (o *ProjectTemplateCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project template create o k response has a 5xx status code
func (o *ProjectTemplateCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project template create o k response a status code equal to that given
func (o *ProjectTemplateCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project template create o k response
func (o *ProjectTemplateCreateOK) Code() int {
	return 200
}

func (o *ProjectTemplateCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectTemplateCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectTemplateCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectTemplateCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateCreateBadRequest creates a ProjectTemplateCreateBadRequest with default headers values
func NewProjectTemplateCreateBadRequest() *ProjectTemplateCreateBadRequest {
	return &ProjectTemplateCreateBadRequest{}
}

/*
ProjectTemplateCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectTemplateCreateBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template create bad request response has a 2xx status code
func (o *ProjectTemplateCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template create bad request response has a 3xx status code
func (o *ProjectTemplateCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template create bad request response has a 4xx status code
func (o *ProjectTemplateCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template create bad request response has a 5xx status code
func (o *ProjectTemplateCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project template create bad request response a status code equal to that given
func (o *ProjectTemplateCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project template create bad request response
func (o *ProjectTemplateCreateBadRequest) Code() int {
	return 400
}

func (o *ProjectTemplateCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectTemplateCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectTemplateCreateBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateCreateUnauthorized creates a ProjectTemplateCreateUnauthorized with default headers values
func NewProjectTemplateCreateUnauthorized() *ProjectTemplateCreateUnauthorized {
	return &ProjectTemplateCreateUnauthorized{}
}

/*
ProjectTemplateCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectTemplateCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template create unauthorized response has a 2xx status code
func (o *ProjectTemplateCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template create unauthorized response has a 3xx status code
func (o *ProjectTemplateCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template create unauthorized response has a 4xx status code
func (o *ProjectTemplateCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template create unauthorized response has a 5xx status code
func (o *ProjectTemplateCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project template create unauthorized response a status code equal to that given
func (o *ProjectTemplateCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the project template create unauthorized response
func (o *ProjectTemplateCreateUnauthorized) Code() int {
	return 401
}

func (o *ProjectTemplateCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectTemplateCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectTemplateCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateCreateForbidden creates a ProjectTemplateCreateForbidden with default headers values
func NewProjectTemplateCreateForbidden() *ProjectTemplateCreateForbidden {
	return &ProjectTemplateCreateForbidden{}
}

/*
ProjectTemplateCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectTemplateCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template create forbidden response has a 2xx status code
func (o *ProjectTemplateCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template create forbidden response has a 3xx status code
func (o *ProjectTemplateCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template create forbidden response has a 4xx status code
func (o *ProjectTemplateCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template create forbidden response has a 5xx status code
func (o *ProjectTemplateCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project template create forbidden response a status code equal to that given
func (o *ProjectTemplateCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the project template create forbidden response
func (o *ProjectTemplateCreateForbidden) Code() int {
	return 403
}

func (o *ProjectTemplateCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateForbidden  %+v", 403, o.Payload)
}

func (o *ProjectTemplateCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateForbidden  %+v", 403, o.Payload)
}

func (o *ProjectTemplateCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateCreateNotFound creates a ProjectTemplateCreateNotFound with default headers values
func NewProjectTemplateCreateNotFound() *ProjectTemplateCreateNotFound {
	return &ProjectTemplateCreateNotFound{}
}

/*
ProjectTemplateCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectTemplateCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template create not found response has a 2xx status code
func (o *ProjectTemplateCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template create not found response has a 3xx status code
func (o *ProjectTemplateCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template create not found response has a 4xx status code
func (o *ProjectTemplateCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template create not found response has a 5xx status code
func (o *ProjectTemplateCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project template create not found response a status code equal to that given
func (o *ProjectTemplateCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the project template create not found response
func (o *ProjectTemplateCreateNotFound) Code() int {
	return 404
}

func (o *ProjectTemplateCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectTemplateCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectTemplateCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateCreateInternalServerError creates a ProjectTemplateCreateInternalServerError with default headers values
func NewProjectTemplateCreateInternalServerError() *ProjectTemplateCreateInternalServerError {
	return &ProjectTemplateCreateInternalServerError{}
}

/*
ProjectTemplateCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectTemplateCreateInternalServerError struct {
}

// IsSuccess returns true when this project template create internal server error response has a 2xx status code
func (o *ProjectTemplateCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template create internal server error response has a 3xx status code
func (o *ProjectTemplateCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template create internal server error response has a 4xx status code
func (o *ProjectTemplateCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project template create internal server error response has a 5xx status code
func (o *ProjectTemplateCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project template create internal server error response a status code equal to that given
func (o *ProjectTemplateCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the project template create internal server error response
func (o *ProjectTemplateCreateInternalServerError) Code() int {
	return 500
}

func (o *ProjectTemplateCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateInternalServerError ", 500)
}

func (o *ProjectTemplateCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectTemplate/create-project][%d] projectTemplateCreateInternalServerError ", 500)
}

func (o *ProjectTemplateCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
