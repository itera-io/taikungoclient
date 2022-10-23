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

// ProjectAppAutoSyncReader is a Reader for the ProjectAppAutoSync structure.
type ProjectAppAutoSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppAutoSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppAutoSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppAutoSyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppAutoSyncUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppAutoSyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppAutoSyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppAutoSyncInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppAutoSyncOK creates a ProjectAppAutoSyncOK with default headers values
func NewProjectAppAutoSyncOK() *ProjectAppAutoSyncOK {
	return &ProjectAppAutoSyncOK{}
}

/*
ProjectAppAutoSyncOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppAutoSyncOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project app auto sync o k response has a 2xx status code
func (o *ProjectAppAutoSyncOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project app auto sync o k response has a 3xx status code
func (o *ProjectAppAutoSyncOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app auto sync o k response has a 4xx status code
func (o *ProjectAppAutoSyncOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app auto sync o k response has a 5xx status code
func (o *ProjectAppAutoSyncOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project app auto sync o k response a status code equal to that given
func (o *ProjectAppAutoSyncOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectAppAutoSyncOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncOK  %+v", 200, o.Payload)
}

func (o *ProjectAppAutoSyncOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncOK  %+v", 200, o.Payload)
}

func (o *ProjectAppAutoSyncOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectAppAutoSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppAutoSyncBadRequest creates a ProjectAppAutoSyncBadRequest with default headers values
func NewProjectAppAutoSyncBadRequest() *ProjectAppAutoSyncBadRequest {
	return &ProjectAppAutoSyncBadRequest{}
}

/*
ProjectAppAutoSyncBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppAutoSyncBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this project app auto sync bad request response has a 2xx status code
func (o *ProjectAppAutoSyncBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app auto sync bad request response has a 3xx status code
func (o *ProjectAppAutoSyncBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app auto sync bad request response has a 4xx status code
func (o *ProjectAppAutoSyncBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app auto sync bad request response has a 5xx status code
func (o *ProjectAppAutoSyncBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project app auto sync bad request response a status code equal to that given
func (o *ProjectAppAutoSyncBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectAppAutoSyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppAutoSyncBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectAppAutoSyncBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ProjectAppAutoSyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppAutoSyncUnauthorized creates a ProjectAppAutoSyncUnauthorized with default headers values
func NewProjectAppAutoSyncUnauthorized() *ProjectAppAutoSyncUnauthorized {
	return &ProjectAppAutoSyncUnauthorized{}
}

/*
ProjectAppAutoSyncUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppAutoSyncUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this project app auto sync unauthorized response has a 2xx status code
func (o *ProjectAppAutoSyncUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app auto sync unauthorized response has a 3xx status code
func (o *ProjectAppAutoSyncUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app auto sync unauthorized response has a 4xx status code
func (o *ProjectAppAutoSyncUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app auto sync unauthorized response has a 5xx status code
func (o *ProjectAppAutoSyncUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project app auto sync unauthorized response a status code equal to that given
func (o *ProjectAppAutoSyncUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectAppAutoSyncUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppAutoSyncUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectAppAutoSyncUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectAppAutoSyncUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppAutoSyncForbidden creates a ProjectAppAutoSyncForbidden with default headers values
func NewProjectAppAutoSyncForbidden() *ProjectAppAutoSyncForbidden {
	return &ProjectAppAutoSyncForbidden{}
}

/*
ProjectAppAutoSyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppAutoSyncForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this project app auto sync forbidden response has a 2xx status code
func (o *ProjectAppAutoSyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app auto sync forbidden response has a 3xx status code
func (o *ProjectAppAutoSyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app auto sync forbidden response has a 4xx status code
func (o *ProjectAppAutoSyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app auto sync forbidden response has a 5xx status code
func (o *ProjectAppAutoSyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project app auto sync forbidden response a status code equal to that given
func (o *ProjectAppAutoSyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectAppAutoSyncForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppAutoSyncForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncForbidden  %+v", 403, o.Payload)
}

func (o *ProjectAppAutoSyncForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectAppAutoSyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppAutoSyncNotFound creates a ProjectAppAutoSyncNotFound with default headers values
func NewProjectAppAutoSyncNotFound() *ProjectAppAutoSyncNotFound {
	return &ProjectAppAutoSyncNotFound{}
}

/*
ProjectAppAutoSyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppAutoSyncNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this project app auto sync not found response has a 2xx status code
func (o *ProjectAppAutoSyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app auto sync not found response has a 3xx status code
func (o *ProjectAppAutoSyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app auto sync not found response has a 4xx status code
func (o *ProjectAppAutoSyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project app auto sync not found response has a 5xx status code
func (o *ProjectAppAutoSyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project app auto sync not found response a status code equal to that given
func (o *ProjectAppAutoSyncNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectAppAutoSyncNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppAutoSyncNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncNotFound  %+v", 404, o.Payload)
}

func (o *ProjectAppAutoSyncNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectAppAutoSyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppAutoSyncInternalServerError creates a ProjectAppAutoSyncInternalServerError with default headers values
func NewProjectAppAutoSyncInternalServerError() *ProjectAppAutoSyncInternalServerError {
	return &ProjectAppAutoSyncInternalServerError{}
}

/*
ProjectAppAutoSyncInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppAutoSyncInternalServerError struct {
}

// IsSuccess returns true when this project app auto sync internal server error response has a 2xx status code
func (o *ProjectAppAutoSyncInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project app auto sync internal server error response has a 3xx status code
func (o *ProjectAppAutoSyncInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project app auto sync internal server error response has a 4xx status code
func (o *ProjectAppAutoSyncInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project app auto sync internal server error response has a 5xx status code
func (o *ProjectAppAutoSyncInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project app auto sync internal server error response a status code equal to that given
func (o *ProjectAppAutoSyncInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectAppAutoSyncInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncInternalServerError ", 500)
}

func (o *ProjectAppAutoSyncInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/autosync][%d] projectAppAutoSyncInternalServerError ", 500)
}

func (o *ProjectAppAutoSyncInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
