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

// ProjectTemplateListReader is a Reader for the ProjectTemplateList structure.
type ProjectTemplateListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectTemplateListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectTemplateListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectTemplateListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectTemplateListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectTemplateListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectTemplateListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectTemplateListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectTemplateListOK creates a ProjectTemplateListOK with default headers values
func NewProjectTemplateListOK() *ProjectTemplateListOK {
	return &ProjectTemplateListOK{}
}

/*
ProjectTemplateListOK describes a response with status code 200, with default header values.

Success
*/
type ProjectTemplateListOK struct {
	Payload *models.ProjectTemplateList
}

// IsSuccess returns true when this project template list o k response has a 2xx status code
func (o *ProjectTemplateListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project template list o k response has a 3xx status code
func (o *ProjectTemplateListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template list o k response has a 4xx status code
func (o *ProjectTemplateListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project template list o k response has a 5xx status code
func (o *ProjectTemplateListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project template list o k response a status code equal to that given
func (o *ProjectTemplateListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectTemplateListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListOK  %+v", 200, o.Payload)
}

func (o *ProjectTemplateListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListOK  %+v", 200, o.Payload)
}

func (o *ProjectTemplateListOK) GetPayload() *models.ProjectTemplateList {
	return o.Payload
}

func (o *ProjectTemplateListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectTemplateList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateListBadRequest creates a ProjectTemplateListBadRequest with default headers values
func NewProjectTemplateListBadRequest() *ProjectTemplateListBadRequest {
	return &ProjectTemplateListBadRequest{}
}

/*
ProjectTemplateListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectTemplateListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template list bad request response has a 2xx status code
func (o *ProjectTemplateListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template list bad request response has a 3xx status code
func (o *ProjectTemplateListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template list bad request response has a 4xx status code
func (o *ProjectTemplateListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template list bad request response has a 5xx status code
func (o *ProjectTemplateListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project template list bad request response a status code equal to that given
func (o *ProjectTemplateListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectTemplateListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectTemplateListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectTemplateListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateListUnauthorized creates a ProjectTemplateListUnauthorized with default headers values
func NewProjectTemplateListUnauthorized() *ProjectTemplateListUnauthorized {
	return &ProjectTemplateListUnauthorized{}
}

/*
ProjectTemplateListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectTemplateListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template list unauthorized response has a 2xx status code
func (o *ProjectTemplateListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template list unauthorized response has a 3xx status code
func (o *ProjectTemplateListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template list unauthorized response has a 4xx status code
func (o *ProjectTemplateListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template list unauthorized response has a 5xx status code
func (o *ProjectTemplateListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project template list unauthorized response a status code equal to that given
func (o *ProjectTemplateListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectTemplateListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectTemplateListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectTemplateListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateListForbidden creates a ProjectTemplateListForbidden with default headers values
func NewProjectTemplateListForbidden() *ProjectTemplateListForbidden {
	return &ProjectTemplateListForbidden{}
}

/*
ProjectTemplateListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectTemplateListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template list forbidden response has a 2xx status code
func (o *ProjectTemplateListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template list forbidden response has a 3xx status code
func (o *ProjectTemplateListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template list forbidden response has a 4xx status code
func (o *ProjectTemplateListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template list forbidden response has a 5xx status code
func (o *ProjectTemplateListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project template list forbidden response a status code equal to that given
func (o *ProjectTemplateListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectTemplateListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectTemplateListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectTemplateListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateListNotFound creates a ProjectTemplateListNotFound with default headers values
func NewProjectTemplateListNotFound() *ProjectTemplateListNotFound {
	return &ProjectTemplateListNotFound{}
}

/*
ProjectTemplateListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectTemplateListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project template list not found response has a 2xx status code
func (o *ProjectTemplateListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template list not found response has a 3xx status code
func (o *ProjectTemplateListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template list not found response has a 4xx status code
func (o *ProjectTemplateListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project template list not found response has a 5xx status code
func (o *ProjectTemplateListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project template list not found response a status code equal to that given
func (o *ProjectTemplateListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectTemplateListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectTemplateListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectTemplateListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectTemplateListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectTemplateListInternalServerError creates a ProjectTemplateListInternalServerError with default headers values
func NewProjectTemplateListInternalServerError() *ProjectTemplateListInternalServerError {
	return &ProjectTemplateListInternalServerError{}
}

/*
ProjectTemplateListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectTemplateListInternalServerError struct {
}

// IsSuccess returns true when this project template list internal server error response has a 2xx status code
func (o *ProjectTemplateListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project template list internal server error response has a 3xx status code
func (o *ProjectTemplateListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project template list internal server error response has a 4xx status code
func (o *ProjectTemplateListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project template list internal server error response has a 5xx status code
func (o *ProjectTemplateListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project template list internal server error response a status code equal to that given
func (o *ProjectTemplateListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectTemplateListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListInternalServerError ", 500)
}

func (o *ProjectTemplateListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/ProjectTemplate][%d] projectTemplateListInternalServerError ", 500)
}

func (o *ProjectTemplateListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
