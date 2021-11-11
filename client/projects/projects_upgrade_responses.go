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

// ProjectsUpgradeReader is a Reader for the ProjectsUpgrade structure.
type ProjectsUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsUpgradeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsUpgradeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsUpgradeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsUpgradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsUpgradeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsUpgradeOK creates a ProjectsUpgradeOK with default headers values
func NewProjectsUpgradeOK() *ProjectsUpgradeOK {
	return &ProjectsUpgradeOK{}
}

/* ProjectsUpgradeOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsUpgradeOK struct {
	Payload models.Unit
}

func (o *ProjectsUpgradeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeOK  %+v", 200, o.Payload)
}
func (o *ProjectsUpgradeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeBadRequest creates a ProjectsUpgradeBadRequest with default headers values
func NewProjectsUpgradeBadRequest() *ProjectsUpgradeBadRequest {
	return &ProjectsUpgradeBadRequest{}
}

/* ProjectsUpgradeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsUpgradeBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectsUpgradeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectsUpgradeBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsUpgradeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeUnauthorized creates a ProjectsUpgradeUnauthorized with default headers values
func NewProjectsUpgradeUnauthorized() *ProjectsUpgradeUnauthorized {
	return &ProjectsUpgradeUnauthorized{}
}

/* ProjectsUpgradeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsUpgradeUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsUpgradeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectsUpgradeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsUpgradeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeForbidden creates a ProjectsUpgradeForbidden with default headers values
func NewProjectsUpgradeForbidden() *ProjectsUpgradeForbidden {
	return &ProjectsUpgradeForbidden{}
}

/* ProjectsUpgradeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsUpgradeForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsUpgradeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeForbidden  %+v", 403, o.Payload)
}
func (o *ProjectsUpgradeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsUpgradeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeNotFound creates a ProjectsUpgradeNotFound with default headers values
func NewProjectsUpgradeNotFound() *ProjectsUpgradeNotFound {
	return &ProjectsUpgradeNotFound{}
}

/* ProjectsUpgradeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsUpgradeNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsUpgradeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeNotFound  %+v", 404, o.Payload)
}
func (o *ProjectsUpgradeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsUpgradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsUpgradeInternalServerError creates a ProjectsUpgradeInternalServerError with default headers values
func NewProjectsUpgradeInternalServerError() *ProjectsUpgradeInternalServerError {
	return &ProjectsUpgradeInternalServerError{}
}

/* ProjectsUpgradeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsUpgradeInternalServerError struct {
}

func (o *ProjectsUpgradeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/upgrade/{projectId}][%d] projectsUpgradeInternalServerError ", 500)
}

func (o *ProjectsUpgradeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
