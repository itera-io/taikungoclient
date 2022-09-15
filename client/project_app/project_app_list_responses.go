// Code generated by go-swagger; DO NOT EDIT.

package project_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectAppListReader is a Reader for the ProjectAppList structure.
type ProjectAppListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppListOK creates a ProjectAppListOK with default headers values
func NewProjectAppListOK() *ProjectAppListOK {
	return &ProjectAppListOK{}
}

/*
ProjectAppListOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppListOK struct {
	Payload *models.ProjectAppList
}

// IsSuccess returns true when this project app list o k response has a 2xx status code
func (o *ProjectAppListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project app list o k response has a 3xx status code
func (o *ProjectAppListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app list o k response has a 4xx status code
func (o *ProjectAppListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app list o k response has a 5xx status code
func (o *ProjectAppListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project app list o k response a status code equal to that given
func (o *ProjectAppListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectAppListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListOK  %+v", 200, o.Payload)
}

func (o *ProjectAppListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListOK  %+v", 200, o.Payload)
}

func (o *ProjectAppListOK) GetPayload() *models.ProjectAppList {
	return o.Payload
}

func (o *ProjectAppListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectAppList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppListBadRequest creates a ProjectAppListBadRequest with default headers values
func NewProjectAppListBadRequest() *ProjectAppListBadRequest {
	return &ProjectAppListBadRequest{}
}

/*
ProjectAppListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this project app list bad request response has a 2xx status code
func (o *ProjectAppListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app list bad request response has a 3xx status code
func (o *ProjectAppListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app list bad request response has a 4xx status code
func (o *ProjectAppListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app list bad request response has a 5xx status code
func (o *ProjectAppListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project app list bad request response a status code equal to that given
func (o *ProjectAppListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectAppListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectAppListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppListUnauthorized creates a ProjectAppListUnauthorized with default headers values
func NewProjectAppListUnauthorized() *ProjectAppListUnauthorized {
	return &ProjectAppListUnauthorized{}
}

/*
ProjectAppListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app list unauthorized response has a 2xx status code
func (o *ProjectAppListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app list unauthorized response has a 3xx status code
func (o *ProjectAppListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app list unauthorized response has a 4xx status code
func (o *ProjectAppListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app list unauthorized response has a 5xx status code
func (o *ProjectAppListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project app list unauthorized response a status code equal to that given
func (o *ProjectAppListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectAppListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppListForbidden creates a ProjectAppListForbidden with default headers values
func NewProjectAppListForbidden() *ProjectAppListForbidden {
	return &ProjectAppListForbidden{}
}

/*
ProjectAppListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app list forbidden response has a 2xx status code
func (o *ProjectAppListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app list forbidden response has a 3xx status code
func (o *ProjectAppListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app list forbidden response has a 4xx status code
func (o *ProjectAppListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app list forbidden response has a 5xx status code
func (o *ProjectAppListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project app list forbidden response a status code equal to that given
func (o *ProjectAppListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectAppListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppListNotFound creates a ProjectAppListNotFound with default headers values
func NewProjectAppListNotFound() *ProjectAppListNotFound {
	return &ProjectAppListNotFound{}
}

/*
ProjectAppListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app list not found response has a 2xx status code
func (o *ProjectAppListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app list not found response has a 3xx status code
func (o *ProjectAppListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app list not found response has a 4xx status code
func (o *ProjectAppListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app list not found response has a 5xx status code
func (o *ProjectAppListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project app list not found response a status code equal to that given
func (o *ProjectAppListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectAppListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppListInternalServerError creates a ProjectAppListInternalServerError with default headers values
func NewProjectAppListInternalServerError() *ProjectAppListInternalServerError {
	return &ProjectAppListInternalServerError{}
}

/*
ProjectAppListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppListInternalServerError struct {
}

// IsSuccess returns true when this project app list internal server error response has a 2xx status code
func (o *ProjectAppListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app list internal server error response has a 3xx status code
func (o *ProjectAppListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app list internal server error response has a 4xx status code
func (o *ProjectAppListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app list internal server error response has a 5xx status code
func (o *ProjectAppListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project app list internal server error response a status code equal to that given
func (o *ProjectAppListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectAppListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListInternalServerError ", 500)
}

func (o *ProjectAppListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectApp/list][%d] projectAppListInternalServerError ", 500)
}

func (o *ProjectAppListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
