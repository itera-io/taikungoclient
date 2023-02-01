// Code generated by go-swagger; DO NOT EDIT.

package project_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectGroupsCreateReader is a Reader for the ProjectGroupsCreate structure.
type ProjectGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGroupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectGroupsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectGroupsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectGroupsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectGroupsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectGroupsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectGroupsCreateOK creates a ProjectGroupsCreateOK with default headers values
func NewProjectGroupsCreateOK() *ProjectGroupsCreateOK {
	return &ProjectGroupsCreateOK{}
}

/*
ProjectGroupsCreateOK describes a response with status code 200, with default header values.

Success
*/
type ProjectGroupsCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this project groups create o k response has a 2xx status code
func (o *ProjectGroupsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project groups create o k response has a 3xx status code
func (o *ProjectGroupsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups create o k response has a 4xx status code
func (o *ProjectGroupsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups create o k response has a 5xx status code
func (o *ProjectGroupsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups create o k response a status code equal to that given
func (o *ProjectGroupsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the project groups create o k response
func (o *ProjectGroupsCreateOK) Code() int {
	return 200
}

func (o *ProjectGroupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectGroupsCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *ProjectGroupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsCreateBadRequest creates a ProjectGroupsCreateBadRequest with default headers values
func NewProjectGroupsCreateBadRequest() *ProjectGroupsCreateBadRequest {
	return &ProjectGroupsCreateBadRequest{}
}

/*
ProjectGroupsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectGroupsCreateBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups create bad request response has a 2xx status code
func (o *ProjectGroupsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups create bad request response has a 3xx status code
func (o *ProjectGroupsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups create bad request response has a 4xx status code
func (o *ProjectGroupsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups create bad request response has a 5xx status code
func (o *ProjectGroupsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups create bad request response a status code equal to that given
func (o *ProjectGroupsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the project groups create bad request response
func (o *ProjectGroupsCreateBadRequest) Code() int {
	return 400
}

func (o *ProjectGroupsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectGroupsCreateBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsCreateUnauthorized creates a ProjectGroupsCreateUnauthorized with default headers values
func NewProjectGroupsCreateUnauthorized() *ProjectGroupsCreateUnauthorized {
	return &ProjectGroupsCreateUnauthorized{}
}

/*
ProjectGroupsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectGroupsCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups create unauthorized response has a 2xx status code
func (o *ProjectGroupsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups create unauthorized response has a 3xx status code
func (o *ProjectGroupsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups create unauthorized response has a 4xx status code
func (o *ProjectGroupsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups create unauthorized response has a 5xx status code
func (o *ProjectGroupsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups create unauthorized response a status code equal to that given
func (o *ProjectGroupsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the project groups create unauthorized response
func (o *ProjectGroupsCreateUnauthorized) Code() int {
	return 401
}

func (o *ProjectGroupsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectGroupsCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsCreateForbidden creates a ProjectGroupsCreateForbidden with default headers values
func NewProjectGroupsCreateForbidden() *ProjectGroupsCreateForbidden {
	return &ProjectGroupsCreateForbidden{}
}

/*
ProjectGroupsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectGroupsCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups create forbidden response has a 2xx status code
func (o *ProjectGroupsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups create forbidden response has a 3xx status code
func (o *ProjectGroupsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups create forbidden response has a 4xx status code
func (o *ProjectGroupsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups create forbidden response has a 5xx status code
func (o *ProjectGroupsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups create forbidden response a status code equal to that given
func (o *ProjectGroupsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the project groups create forbidden response
func (o *ProjectGroupsCreateForbidden) Code() int {
	return 403
}

func (o *ProjectGroupsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ProjectGroupsCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsCreateNotFound creates a ProjectGroupsCreateNotFound with default headers values
func NewProjectGroupsCreateNotFound() *ProjectGroupsCreateNotFound {
	return &ProjectGroupsCreateNotFound{}
}

/*
ProjectGroupsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectGroupsCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project groups create not found response has a 2xx status code
func (o *ProjectGroupsCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups create not found response has a 3xx status code
func (o *ProjectGroupsCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups create not found response has a 4xx status code
func (o *ProjectGroupsCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project groups create not found response has a 5xx status code
func (o *ProjectGroupsCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project groups create not found response a status code equal to that given
func (o *ProjectGroupsCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the project groups create not found response
func (o *ProjectGroupsCreateNotFound) Code() int {
	return 404
}

func (o *ProjectGroupsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateNotFound  %+v", 404, o.Payload)
}

func (o *ProjectGroupsCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsCreateInternalServerError creates a ProjectGroupsCreateInternalServerError with default headers values
func NewProjectGroupsCreateInternalServerError() *ProjectGroupsCreateInternalServerError {
	return &ProjectGroupsCreateInternalServerError{}
}

/*
ProjectGroupsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectGroupsCreateInternalServerError struct {
}

// IsSuccess returns true when this project groups create internal server error response has a 2xx status code
func (o *ProjectGroupsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project groups create internal server error response has a 3xx status code
func (o *ProjectGroupsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project groups create internal server error response has a 4xx status code
func (o *ProjectGroupsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project groups create internal server error response has a 5xx status code
func (o *ProjectGroupsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project groups create internal server error response a status code equal to that given
func (o *ProjectGroupsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the project groups create internal server error response
func (o *ProjectGroupsCreateInternalServerError) Code() int {
	return 500
}

func (o *ProjectGroupsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateInternalServerError ", 500)
}

func (o *ProjectGroupsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectGroups/create][%d] projectGroupsCreateInternalServerError ", 500)
}

func (o *ProjectGroupsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
