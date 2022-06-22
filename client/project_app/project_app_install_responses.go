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

// ProjectAppInstallReader is a Reader for the ProjectAppInstall structure.
type ProjectAppInstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppInstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppInstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppInstallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppInstallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppInstallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppInstallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppInstallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppInstallOK creates a ProjectAppInstallOK with default headers values
func NewProjectAppInstallOK() *ProjectAppInstallOK {
	return &ProjectAppInstallOK{}
}

/* ProjectAppInstallOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppInstallOK struct {
	Payload *models.APIResponse
}

func (o *ProjectAppInstallOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/install][%d] projectAppInstallOK  %+v", 200, o.Payload)
}
func (o *ProjectAppInstallOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *ProjectAppInstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppInstallBadRequest creates a ProjectAppInstallBadRequest with default headers values
func NewProjectAppInstallBadRequest() *ProjectAppInstallBadRequest {
	return &ProjectAppInstallBadRequest{}
}

/* ProjectAppInstallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppInstallBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectAppInstallBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/install][%d] projectAppInstallBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectAppInstallBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectAppInstallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppInstallUnauthorized creates a ProjectAppInstallUnauthorized with default headers values
func NewProjectAppInstallUnauthorized() *ProjectAppInstallUnauthorized {
	return &ProjectAppInstallUnauthorized{}
}

/* ProjectAppInstallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppInstallUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppInstallUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/install][%d] projectAppInstallUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectAppInstallUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppInstallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppInstallForbidden creates a ProjectAppInstallForbidden with default headers values
func NewProjectAppInstallForbidden() *ProjectAppInstallForbidden {
	return &ProjectAppInstallForbidden{}
}

/* ProjectAppInstallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppInstallForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppInstallForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/install][%d] projectAppInstallForbidden  %+v", 403, o.Payload)
}
func (o *ProjectAppInstallForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppInstallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppInstallNotFound creates a ProjectAppInstallNotFound with default headers values
func NewProjectAppInstallNotFound() *ProjectAppInstallNotFound {
	return &ProjectAppInstallNotFound{}
}

/* ProjectAppInstallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppInstallNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppInstallNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/install][%d] projectAppInstallNotFound  %+v", 404, o.Payload)
}
func (o *ProjectAppInstallNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppInstallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppInstallInternalServerError creates a ProjectAppInstallInternalServerError with default headers values
func NewProjectAppInstallInternalServerError() *ProjectAppInstallInternalServerError {
	return &ProjectAppInstallInternalServerError{}
}

/* ProjectAppInstallInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppInstallInternalServerError struct {
}

func (o *ProjectAppInstallInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectApp/install][%d] projectAppInstallInternalServerError ", 500)
}

func (o *ProjectAppInstallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}