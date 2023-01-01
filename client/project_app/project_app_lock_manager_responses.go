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

// ProjectAppLockManagerReader is a Reader for the ProjectAppLockManager structure.
type ProjectAppLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppLockManagerOK creates a ProjectAppLockManagerOK with default headers values
func NewProjectAppLockManagerOK() *ProjectAppLockManagerOK {
	return &ProjectAppLockManagerOK{}
}

/*
ProjectAppLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppLockManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project app lock manager o k response has a 2xx status code
func (o *ProjectAppLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project app lock manager o k response has a 3xx status code
func (o *ProjectAppLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app lock manager o k response has a 4xx status code
func (o *ProjectAppLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app lock manager o k response has a 5xx status code
func (o *ProjectAppLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project app lock manager o k response a status code equal to that given
func (o *ProjectAppLockManagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectAppLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerOK  %+v", 200, o.Payload)
}

func (o *ProjectAppLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerOK  %+v", 200, o.Payload)
}

func (o *ProjectAppLockManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectAppLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppLockManagerBadRequest creates a ProjectAppLockManagerBadRequest with default headers values
func NewProjectAppLockManagerBadRequest() *ProjectAppLockManagerBadRequest {
	return &ProjectAppLockManagerBadRequest{}
}

/*
ProjectAppLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppLockManagerBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app lock manager bad request response has a 2xx status code
func (o *ProjectAppLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app lock manager bad request response has a 3xx status code
func (o *ProjectAppLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app lock manager bad request response has a 4xx status code
func (o *ProjectAppLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app lock manager bad request response has a 5xx status code
func (o *ProjectAppLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project app lock manager bad request response a status code equal to that given
func (o *ProjectAppLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectAppLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppLockManagerBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppLockManagerUnauthorized creates a ProjectAppLockManagerUnauthorized with default headers values
func NewProjectAppLockManagerUnauthorized() *ProjectAppLockManagerUnauthorized {
	return &ProjectAppLockManagerUnauthorized{}
}

/*
ProjectAppLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppLockManagerUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app lock manager unauthorized response has a 2xx status code
func (o *ProjectAppLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app lock manager unauthorized response has a 3xx status code
func (o *ProjectAppLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app lock manager unauthorized response has a 4xx status code
func (o *ProjectAppLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app lock manager unauthorized response has a 5xx status code
func (o *ProjectAppLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project app lock manager unauthorized response a status code equal to that given
func (o *ProjectAppLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectAppLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppLockManagerUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppLockManagerForbidden creates a ProjectAppLockManagerForbidden with default headers values
func NewProjectAppLockManagerForbidden() *ProjectAppLockManagerForbidden {
	return &ProjectAppLockManagerForbidden{}
}

/*
ProjectAppLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppLockManagerForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app lock manager forbidden response has a 2xx status code
func (o *ProjectAppLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app lock manager forbidden response has a 3xx status code
func (o *ProjectAppLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app lock manager forbidden response has a 4xx status code
func (o *ProjectAppLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app lock manager forbidden response has a 5xx status code
func (o *ProjectAppLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project app lock manager forbidden response a status code equal to that given
func (o *ProjectAppLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectAppLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppLockManagerForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppLockManagerNotFound creates a ProjectAppLockManagerNotFound with default headers values
func NewProjectAppLockManagerNotFound() *ProjectAppLockManagerNotFound {
	return &ProjectAppLockManagerNotFound{}
}

/*
ProjectAppLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppLockManagerNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project app lock manager not found response has a 2xx status code
func (o *ProjectAppLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app lock manager not found response has a 3xx status code
func (o *ProjectAppLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app lock manager not found response has a 4xx status code
func (o *ProjectAppLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app lock manager not found response has a 5xx status code
func (o *ProjectAppLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project app lock manager not found response a status code equal to that given
func (o *ProjectAppLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectAppLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppLockManagerNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppLockManagerInternalServerError creates a ProjectAppLockManagerInternalServerError with default headers values
func NewProjectAppLockManagerInternalServerError() *ProjectAppLockManagerInternalServerError {
	return &ProjectAppLockManagerInternalServerError{}
}

/*
ProjectAppLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppLockManagerInternalServerError struct {
}

// IsSuccess returns true when this project app lock manager internal server error response has a 2xx status code
func (o *ProjectAppLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app lock manager internal server error response has a 3xx status code
func (o *ProjectAppLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app lock manager internal server error response has a 4xx status code
func (o *ProjectAppLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app lock manager internal server error response has a 5xx status code
func (o *ProjectAppLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project app lock manager internal server error response a status code equal to that given
func (o *ProjectAppLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectAppLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerInternalServerError ", 500)
}

func (o *ProjectAppLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/lockmanager][%d] projectAppLockManagerInternalServerError ", 500)
}

func (o *ProjectAppLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
