// Code generated by go-swagger; DO NOT EDIT.

package project_infracosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectInfracostsDeleteReader is a Reader for the ProjectInfracostsDelete structure.
type ProjectInfracostsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectInfracostsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectInfracostsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectInfracostsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectInfracostsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectInfracostsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectInfracostsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectInfracostsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectInfracostsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectInfracostsDeleteOK creates a ProjectInfracostsDeleteOK with default headers values
func NewProjectInfracostsDeleteOK() *ProjectInfracostsDeleteOK {
	return &ProjectInfracostsDeleteOK{}
}

/*
ProjectInfracostsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type ProjectInfracostsDeleteOK struct {
}

// IsSuccess returns true when this project infracosts delete o k response has a 2xx status code
func (o *ProjectInfracostsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project infracosts delete o k response has a 3xx status code
func (o *ProjectInfracostsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete o k response has a 4xx status code
func (o *ProjectInfracostsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project infracosts delete o k response has a 5xx status code
func (o *ProjectInfracostsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts delete o k response a status code equal to that given
func (o *ProjectInfracostsDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectInfracostsDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteOK ", 200)
}

func (o *ProjectInfracostsDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteOK ", 200)
}

func (o *ProjectInfracostsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectInfracostsDeleteNoContent creates a ProjectInfracostsDeleteNoContent with default headers values
func NewProjectInfracostsDeleteNoContent() *ProjectInfracostsDeleteNoContent {
	return &ProjectInfracostsDeleteNoContent{}
}

/*
ProjectInfracostsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type ProjectInfracostsDeleteNoContent struct {
}

// IsSuccess returns true when this project infracosts delete no content response has a 2xx status code
func (o *ProjectInfracostsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project infracosts delete no content response has a 3xx status code
func (o *ProjectInfracostsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete no content response has a 4xx status code
func (o *ProjectInfracostsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this project infracosts delete no content response has a 5xx status code
func (o *ProjectInfracostsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts delete no content response a status code equal to that given
func (o *ProjectInfracostsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ProjectInfracostsDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteNoContent ", 204)
}

func (o *ProjectInfracostsDeleteNoContent) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteNoContent ", 204)
}

func (o *ProjectInfracostsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectInfracostsDeleteBadRequest creates a ProjectInfracostsDeleteBadRequest with default headers values
func NewProjectInfracostsDeleteBadRequest() *ProjectInfracostsDeleteBadRequest {
	return &ProjectInfracostsDeleteBadRequest{}
}

/*
ProjectInfracostsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectInfracostsDeleteBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts delete bad request response has a 2xx status code
func (o *ProjectInfracostsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts delete bad request response has a 3xx status code
func (o *ProjectInfracostsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete bad request response has a 4xx status code
func (o *ProjectInfracostsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts delete bad request response has a 5xx status code
func (o *ProjectInfracostsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts delete bad request response a status code equal to that given
func (o *ProjectInfracostsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectInfracostsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectInfracostsDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectInfracostsDeleteBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsDeleteUnauthorized creates a ProjectInfracostsDeleteUnauthorized with default headers values
func NewProjectInfracostsDeleteUnauthorized() *ProjectInfracostsDeleteUnauthorized {
	return &ProjectInfracostsDeleteUnauthorized{}
}

/*
ProjectInfracostsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectInfracostsDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts delete unauthorized response has a 2xx status code
func (o *ProjectInfracostsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts delete unauthorized response has a 3xx status code
func (o *ProjectInfracostsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete unauthorized response has a 4xx status code
func (o *ProjectInfracostsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts delete unauthorized response has a 5xx status code
func (o *ProjectInfracostsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts delete unauthorized response a status code equal to that given
func (o *ProjectInfracostsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectInfracostsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectInfracostsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectInfracostsDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsDeleteForbidden creates a ProjectInfracostsDeleteForbidden with default headers values
func NewProjectInfracostsDeleteForbidden() *ProjectInfracostsDeleteForbidden {
	return &ProjectInfracostsDeleteForbidden{}
}

/*
ProjectInfracostsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectInfracostsDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts delete forbidden response has a 2xx status code
func (o *ProjectInfracostsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts delete forbidden response has a 3xx status code
func (o *ProjectInfracostsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete forbidden response has a 4xx status code
func (o *ProjectInfracostsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts delete forbidden response has a 5xx status code
func (o *ProjectInfracostsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts delete forbidden response a status code equal to that given
func (o *ProjectInfracostsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectInfracostsDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectInfracostsDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ProjectInfracostsDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsDeleteNotFound creates a ProjectInfracostsDeleteNotFound with default headers values
func NewProjectInfracostsDeleteNotFound() *ProjectInfracostsDeleteNotFound {
	return &ProjectInfracostsDeleteNotFound{}
}

/*
ProjectInfracostsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectInfracostsDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts delete not found response has a 2xx status code
func (o *ProjectInfracostsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts delete not found response has a 3xx status code
func (o *ProjectInfracostsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete not found response has a 4xx status code
func (o *ProjectInfracostsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts delete not found response has a 5xx status code
func (o *ProjectInfracostsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts delete not found response a status code equal to that given
func (o *ProjectInfracostsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectInfracostsDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectInfracostsDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ProjectInfracostsDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsDeleteInternalServerError creates a ProjectInfracostsDeleteInternalServerError with default headers values
func NewProjectInfracostsDeleteInternalServerError() *ProjectInfracostsDeleteInternalServerError {
	return &ProjectInfracostsDeleteInternalServerError{}
}

/*
ProjectInfracostsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectInfracostsDeleteInternalServerError struct {
}

// IsSuccess returns true when this project infracosts delete internal server error response has a 2xx status code
func (o *ProjectInfracostsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts delete internal server error response has a 3xx status code
func (o *ProjectInfracostsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts delete internal server error response has a 4xx status code
func (o *ProjectInfracostsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project infracosts delete internal server error response has a 5xx status code
func (o *ProjectInfracostsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project infracosts delete internal server error response a status code equal to that given
func (o *ProjectInfracostsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectInfracostsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteInternalServerError ", 500)
}

func (o *ProjectInfracostsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/delete][%d] projectInfracostsDeleteInternalServerError ", 500)
}

func (o *ProjectInfracostsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
