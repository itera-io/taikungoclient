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

// ProjectInfracostsEditReader is a Reader for the ProjectInfracostsEdit structure.
type ProjectInfracostsEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectInfracostsEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectInfracostsEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectInfracostsEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectInfracostsEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectInfracostsEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectInfracostsEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectInfracostsEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectInfracostsEditOK creates a ProjectInfracostsEditOK with default headers values
func NewProjectInfracostsEditOK() *ProjectInfracostsEditOK {
	return &ProjectInfracostsEditOK{}
}

/*
ProjectInfracostsEditOK describes a response with status code 200, with default header values.

Success
*/
type ProjectInfracostsEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this project infracosts edit o k response has a 2xx status code
func (o *ProjectInfracostsEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project infracosts edit o k response has a 3xx status code
func (o *ProjectInfracostsEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts edit o k response has a 4xx status code
func (o *ProjectInfracostsEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project infracosts edit o k response has a 5xx status code
func (o *ProjectInfracostsEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts edit o k response a status code equal to that given
func (o *ProjectInfracostsEditOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectInfracostsEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditOK  %+v", 200, o.Payload)
}

func (o *ProjectInfracostsEditOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditOK  %+v", 200, o.Payload)
}

func (o *ProjectInfracostsEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectInfracostsEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsEditBadRequest creates a ProjectInfracostsEditBadRequest with default headers values
func NewProjectInfracostsEditBadRequest() *ProjectInfracostsEditBadRequest {
	return &ProjectInfracostsEditBadRequest{}
}

/*
ProjectInfracostsEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectInfracostsEditBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this project infracosts edit bad request response has a 2xx status code
func (o *ProjectInfracostsEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts edit bad request response has a 3xx status code
func (o *ProjectInfracostsEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts edit bad request response has a 4xx status code
func (o *ProjectInfracostsEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts edit bad request response has a 5xx status code
func (o *ProjectInfracostsEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts edit bad request response a status code equal to that given
func (o *ProjectInfracostsEditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectInfracostsEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectInfracostsEditBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectInfracostsEditBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ProjectInfracostsEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsEditUnauthorized creates a ProjectInfracostsEditUnauthorized with default headers values
func NewProjectInfracostsEditUnauthorized() *ProjectInfracostsEditUnauthorized {
	return &ProjectInfracostsEditUnauthorized{}
}

/*
ProjectInfracostsEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectInfracostsEditUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts edit unauthorized response has a 2xx status code
func (o *ProjectInfracostsEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts edit unauthorized response has a 3xx status code
func (o *ProjectInfracostsEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts edit unauthorized response has a 4xx status code
func (o *ProjectInfracostsEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts edit unauthorized response has a 5xx status code
func (o *ProjectInfracostsEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts edit unauthorized response a status code equal to that given
func (o *ProjectInfracostsEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectInfracostsEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectInfracostsEditUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectInfracostsEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsEditForbidden creates a ProjectInfracostsEditForbidden with default headers values
func NewProjectInfracostsEditForbidden() *ProjectInfracostsEditForbidden {
	return &ProjectInfracostsEditForbidden{}
}

/*
ProjectInfracostsEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectInfracostsEditForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts edit forbidden response has a 2xx status code
func (o *ProjectInfracostsEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts edit forbidden response has a 3xx status code
func (o *ProjectInfracostsEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts edit forbidden response has a 4xx status code
func (o *ProjectInfracostsEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts edit forbidden response has a 5xx status code
func (o *ProjectInfracostsEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts edit forbidden response a status code equal to that given
func (o *ProjectInfracostsEditForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectInfracostsEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditForbidden  %+v", 403, o.Payload)
}

func (o *ProjectInfracostsEditForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditForbidden  %+v", 403, o.Payload)
}

func (o *ProjectInfracostsEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsEditNotFound creates a ProjectInfracostsEditNotFound with default headers values
func NewProjectInfracostsEditNotFound() *ProjectInfracostsEditNotFound {
	return &ProjectInfracostsEditNotFound{}
}

/*
ProjectInfracostsEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectInfracostsEditNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this project infracosts edit not found response has a 2xx status code
func (o *ProjectInfracostsEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts edit not found response has a 3xx status code
func (o *ProjectInfracostsEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts edit not found response has a 4xx status code
func (o *ProjectInfracostsEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this project infracosts edit not found response has a 5xx status code
func (o *ProjectInfracostsEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this project infracosts edit not found response a status code equal to that given
func (o *ProjectInfracostsEditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectInfracostsEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditNotFound  %+v", 404, o.Payload)
}

func (o *ProjectInfracostsEditNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditNotFound  %+v", 404, o.Payload)
}

func (o *ProjectInfracostsEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectInfracostsEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectInfracostsEditInternalServerError creates a ProjectInfracostsEditInternalServerError with default headers values
func NewProjectInfracostsEditInternalServerError() *ProjectInfracostsEditInternalServerError {
	return &ProjectInfracostsEditInternalServerError{}
}

/*
ProjectInfracostsEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectInfracostsEditInternalServerError struct {
}

// IsSuccess returns true when this project infracosts edit internal server error response has a 2xx status code
func (o *ProjectInfracostsEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this project infracosts edit internal server error response has a 3xx status code
func (o *ProjectInfracostsEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project infracosts edit internal server error response has a 4xx status code
func (o *ProjectInfracostsEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this project infracosts edit internal server error response has a 5xx status code
func (o *ProjectInfracostsEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this project infracosts edit internal server error response a status code equal to that given
func (o *ProjectInfracostsEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectInfracostsEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditInternalServerError ", 500)
}

func (o *ProjectInfracostsEditInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectInfracosts/upsert/{projectId}][%d] projectInfracostsEditInternalServerError ", 500)
}

func (o *ProjectInfracostsEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
