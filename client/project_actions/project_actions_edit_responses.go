// Code generated by go-swagger; DO NOT EDIT.

package project_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectActionsEditReader is a Reader for the ProjectActionsEdit structure.
type ProjectActionsEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectActionsEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectActionsEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectActionsEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectActionsEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectActionsEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectActionsEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectActionsEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectActionsEditOK creates a ProjectActionsEditOK with default headers values
func NewProjectActionsEditOK() *ProjectActionsEditOK {
	return &ProjectActionsEditOK{}
}

/*
ProjectActionsEditOK describes a response with status code 200, with default header values.

Success
*/
type ProjectActionsEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project actions edit o k response has a 2xx status code
func (o *ProjectActionsEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project actions edit o k response has a 3xx status code
func (o *ProjectActionsEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions edit o k response has a 4xx status code
func (o *ProjectActionsEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project actions edit o k response has a 5xx status code
func (o *ProjectActionsEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions edit o k response a status code equal to that given
func (o *ProjectActionsEditOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectActionsEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditOK  %+v", 200, o.Payload)
}

func (o *ProjectActionsEditOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditOK  %+v", 200, o.Payload)
}

func (o *ProjectActionsEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectActionsEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsEditBadRequest creates a ProjectActionsEditBadRequest with default headers values
func NewProjectActionsEditBadRequest() *ProjectActionsEditBadRequest {
	return &ProjectActionsEditBadRequest{}
}

/*
ProjectActionsEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectActionsEditBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project actions edit bad request response has a 2xx status code
func (o *ProjectActionsEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions edit bad request response has a 3xx status code
func (o *ProjectActionsEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions edit bad request response has a 4xx status code
func (o *ProjectActionsEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions edit bad request response has a 5xx status code
func (o *ProjectActionsEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions edit bad request response a status code equal to that given
func (o *ProjectActionsEditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectActionsEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectActionsEditBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectActionsEditBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectActionsEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsEditUnauthorized creates a ProjectActionsEditUnauthorized with default headers values
func NewProjectActionsEditUnauthorized() *ProjectActionsEditUnauthorized {
	return &ProjectActionsEditUnauthorized{}
}

/*
ProjectActionsEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectActionsEditUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project actions edit unauthorized response has a 2xx status code
func (o *ProjectActionsEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions edit unauthorized response has a 3xx status code
func (o *ProjectActionsEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions edit unauthorized response has a 4xx status code
func (o *ProjectActionsEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions edit unauthorized response has a 5xx status code
func (o *ProjectActionsEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions edit unauthorized response a status code equal to that given
func (o *ProjectActionsEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectActionsEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectActionsEditUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectActionsEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectActionsEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsEditForbidden creates a ProjectActionsEditForbidden with default headers values
func NewProjectActionsEditForbidden() *ProjectActionsEditForbidden {
	return &ProjectActionsEditForbidden{}
}

/*
ProjectActionsEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectActionsEditForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project actions edit forbidden response has a 2xx status code
func (o *ProjectActionsEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions edit forbidden response has a 3xx status code
func (o *ProjectActionsEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions edit forbidden response has a 4xx status code
func (o *ProjectActionsEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions edit forbidden response has a 5xx status code
func (o *ProjectActionsEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions edit forbidden response a status code equal to that given
func (o *ProjectActionsEditForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectActionsEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditForbidden  %+v", 403, o.Payload)
}

func (o *ProjectActionsEditForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditForbidden  %+v", 403, o.Payload)
}

func (o *ProjectActionsEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectActionsEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsEditNotFound creates a ProjectActionsEditNotFound with default headers values
func NewProjectActionsEditNotFound() *ProjectActionsEditNotFound {
	return &ProjectActionsEditNotFound{}
}

/*
ProjectActionsEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectActionsEditNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project actions edit not found response has a 2xx status code
func (o *ProjectActionsEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions edit not found response has a 3xx status code
func (o *ProjectActionsEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions edit not found response has a 4xx status code
func (o *ProjectActionsEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project actions edit not found response has a 5xx status code
func (o *ProjectActionsEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project actions edit not found response a status code equal to that given
func (o *ProjectActionsEditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectActionsEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditNotFound  %+v", 404, o.Payload)
}

func (o *ProjectActionsEditNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditNotFound  %+v", 404, o.Payload)
}

func (o *ProjectActionsEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectActionsEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectActionsEditInternalServerError creates a ProjectActionsEditInternalServerError with default headers values
func NewProjectActionsEditInternalServerError() *ProjectActionsEditInternalServerError {
	return &ProjectActionsEditInternalServerError{}
}

/*
ProjectActionsEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectActionsEditInternalServerError struct {
}

// IsSuccess returns true when this project actions edit internal server error response has a 2xx status code
func (o *ProjectActionsEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project actions edit internal server error response has a 3xx status code
func (o *ProjectActionsEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project actions edit internal server error response has a 4xx status code
func (o *ProjectActionsEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project actions edit internal server error response has a 5xx status code
func (o *ProjectActionsEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project actions edit internal server error response a status code equal to that given
func (o *ProjectActionsEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectActionsEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditInternalServerError ", 500)
}

func (o *ProjectActionsEditInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectActions/update/{projectId}][%d] projectActionsEditInternalServerError ", 500)
}

func (o *ProjectActionsEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
