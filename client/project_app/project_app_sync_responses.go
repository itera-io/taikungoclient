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

// ProjectAppSyncReader is a Reader for the ProjectAppSync structure.
type ProjectAppSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppSyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppSyncUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppSyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppSyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppSyncInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppSyncOK creates a ProjectAppSyncOK with default headers values
func NewProjectAppSyncOK() *ProjectAppSyncOK {
	return &ProjectAppSyncOK{}
}

/*
ProjectAppSyncOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppSyncOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project app sync o k response has a 2xx status code
func (o *ProjectAppSyncOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project app sync o k response has a 3xx status code
func (o *ProjectAppSyncOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app sync o k response has a 4xx status code
func (o *ProjectAppSyncOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app sync o k response has a 5xx status code
func (o *ProjectAppSyncOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project app sync o k response a status code equal to that given
func (o *ProjectAppSyncOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectAppSyncOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncOK  %+v", 200, o.Payload)
}

func (o *ProjectAppSyncOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncOK  %+v", 200, o.Payload)
}

func (o *ProjectAppSyncOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectAppSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppSyncBadRequest creates a ProjectAppSyncBadRequest with default headers values
func NewProjectAppSyncBadRequest() *ProjectAppSyncBadRequest {
	return &ProjectAppSyncBadRequest{}
}

/*
ProjectAppSyncBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppSyncBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this project app sync bad request response has a 2xx status code
func (o *ProjectAppSyncBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app sync bad request response has a 3xx status code
func (o *ProjectAppSyncBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app sync bad request response has a 4xx status code
func (o *ProjectAppSyncBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app sync bad request response has a 5xx status code
func (o *ProjectAppSyncBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project app sync bad request response a status code equal to that given
func (o *ProjectAppSyncBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectAppSyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppSyncBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppSyncBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ProjectAppSyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppSyncUnauthorized creates a ProjectAppSyncUnauthorized with default headers values
func NewProjectAppSyncUnauthorized() *ProjectAppSyncUnauthorized {
	return &ProjectAppSyncUnauthorized{}
}

/*
ProjectAppSyncUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppSyncUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app sync unauthorized response has a 2xx status code
func (o *ProjectAppSyncUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app sync unauthorized response has a 3xx status code
func (o *ProjectAppSyncUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app sync unauthorized response has a 4xx status code
func (o *ProjectAppSyncUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app sync unauthorized response has a 5xx status code
func (o *ProjectAppSyncUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project app sync unauthorized response a status code equal to that given
func (o *ProjectAppSyncUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectAppSyncUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppSyncUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppSyncUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppSyncUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppSyncForbidden creates a ProjectAppSyncForbidden with default headers values
func NewProjectAppSyncForbidden() *ProjectAppSyncForbidden {
	return &ProjectAppSyncForbidden{}
}

/*
ProjectAppSyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppSyncForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app sync forbidden response has a 2xx status code
func (o *ProjectAppSyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app sync forbidden response has a 3xx status code
func (o *ProjectAppSyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app sync forbidden response has a 4xx status code
func (o *ProjectAppSyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app sync forbidden response has a 5xx status code
func (o *ProjectAppSyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project app sync forbidden response a status code equal to that given
func (o *ProjectAppSyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectAppSyncForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppSyncForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppSyncForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppSyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppSyncNotFound creates a ProjectAppSyncNotFound with default headers values
func NewProjectAppSyncNotFound() *ProjectAppSyncNotFound {
	return &ProjectAppSyncNotFound{}
}

/*
ProjectAppSyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppSyncNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app sync not found response has a 2xx status code
func (o *ProjectAppSyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app sync not found response has a 3xx status code
func (o *ProjectAppSyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app sync not found response has a 4xx status code
func (o *ProjectAppSyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app sync not found response has a 5xx status code
func (o *ProjectAppSyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project app sync not found response a status code equal to that given
func (o *ProjectAppSyncNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectAppSyncNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppSyncNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppSyncNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppSyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppSyncInternalServerError creates a ProjectAppSyncInternalServerError with default headers values
func NewProjectAppSyncInternalServerError() *ProjectAppSyncInternalServerError {
	return &ProjectAppSyncInternalServerError{}
}

/*
ProjectAppSyncInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppSyncInternalServerError struct {
}

// IsSuccess returns true when this project app sync internal server error response has a 2xx status code
func (o *ProjectAppSyncInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app sync internal server error response has a 3xx status code
func (o *ProjectAppSyncInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app sync internal server error response has a 4xx status code
func (o *ProjectAppSyncInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app sync internal server error response has a 5xx status code
func (o *ProjectAppSyncInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project app sync internal server error response a status code equal to that given
func (o *ProjectAppSyncInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectAppSyncInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncInternalServerError ", 500)
}

func (o *ProjectAppSyncInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/sync][%d] projectAppSyncInternalServerError ", 500)
}

func (o *ProjectAppSyncInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
