// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectsPurgeWholeProjectReader is a Reader for the ProjectsPurgeWholeProject structure.
type ProjectsPurgeWholeProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsPurgeWholeProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsPurgeWholeProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsPurgeWholeProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsPurgeWholeProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsPurgeWholeProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsPurgeWholeProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsPurgeWholeProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsPurgeWholeProjectOK creates a ProjectsPurgeWholeProjectOK with default headers values
func NewProjectsPurgeWholeProjectOK() *ProjectsPurgeWholeProjectOK {
	return &ProjectsPurgeWholeProjectOK{}
}

/* ProjectsPurgeWholeProjectOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsPurgeWholeProjectOK struct {
	Payload models.Unit
}

func (o *ProjectsPurgeWholeProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purgewholeproject][%d] projectsPurgeWholeProjectOK  %+v", 200, o.Payload)
}
func (o *ProjectsPurgeWholeProjectOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsPurgeWholeProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeWholeProjectBadRequest creates a ProjectsPurgeWholeProjectBadRequest with default headers values
func NewProjectsPurgeWholeProjectBadRequest() *ProjectsPurgeWholeProjectBadRequest {
	return &ProjectsPurgeWholeProjectBadRequest{}
}

/* ProjectsPurgeWholeProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsPurgeWholeProjectBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectsPurgeWholeProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purgewholeproject][%d] projectsPurgeWholeProjectBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectsPurgeWholeProjectBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeWholeProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeWholeProjectUnauthorized creates a ProjectsPurgeWholeProjectUnauthorized with default headers values
func NewProjectsPurgeWholeProjectUnauthorized() *ProjectsPurgeWholeProjectUnauthorized {
	return &ProjectsPurgeWholeProjectUnauthorized{}
}

/* ProjectsPurgeWholeProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsPurgeWholeProjectUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsPurgeWholeProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purgewholeproject][%d] projectsPurgeWholeProjectUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectsPurgeWholeProjectUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeWholeProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeWholeProjectForbidden creates a ProjectsPurgeWholeProjectForbidden with default headers values
func NewProjectsPurgeWholeProjectForbidden() *ProjectsPurgeWholeProjectForbidden {
	return &ProjectsPurgeWholeProjectForbidden{}
}

/* ProjectsPurgeWholeProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsPurgeWholeProjectForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsPurgeWholeProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purgewholeproject][%d] projectsPurgeWholeProjectForbidden  %+v", 403, o.Payload)
}
func (o *ProjectsPurgeWholeProjectForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeWholeProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeWholeProjectNotFound creates a ProjectsPurgeWholeProjectNotFound with default headers values
func NewProjectsPurgeWholeProjectNotFound() *ProjectsPurgeWholeProjectNotFound {
	return &ProjectsPurgeWholeProjectNotFound{}
}

/* ProjectsPurgeWholeProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsPurgeWholeProjectNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsPurgeWholeProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purgewholeproject][%d] projectsPurgeWholeProjectNotFound  %+v", 404, o.Payload)
}
func (o *ProjectsPurgeWholeProjectNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPurgeWholeProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPurgeWholeProjectInternalServerError creates a ProjectsPurgeWholeProjectInternalServerError with default headers values
func NewProjectsPurgeWholeProjectInternalServerError() *ProjectsPurgeWholeProjectInternalServerError {
	return &ProjectsPurgeWholeProjectInternalServerError{}
}

/* ProjectsPurgeWholeProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsPurgeWholeProjectInternalServerError struct {
}

func (o *ProjectsPurgeWholeProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/purgewholeproject][%d] projectsPurgeWholeProjectInternalServerError ", 500)
}

func (o *ProjectsPurgeWholeProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
