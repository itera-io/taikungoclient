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

// ProjectAppUninstallReader is a Reader for the ProjectAppUninstall structure.
type ProjectAppUninstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectAppUninstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectAppUninstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectAppUninstallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectAppUninstallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectAppUninstallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectAppUninstallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectAppUninstallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectAppUninstallOK creates a ProjectAppUninstallOK with default headers values
func NewProjectAppUninstallOK() *ProjectAppUninstallOK {
	return &ProjectAppUninstallOK{}
}

/* ProjectAppUninstallOK describes a response with status code 200, with default header values.

Success
*/
type ProjectAppUninstallOK struct {
	Payload models.Unit
}

func (o *ProjectAppUninstallOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallOK  %+v", 200, o.Payload)
}
func (o *ProjectAppUninstallOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectAppUninstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallBadRequest creates a ProjectAppUninstallBadRequest with default headers values
func NewProjectAppUninstallBadRequest() *ProjectAppUninstallBadRequest {
	return &ProjectAppUninstallBadRequest{}
}

/* ProjectAppUninstallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectAppUninstallBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectAppUninstallBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectAppUninstallBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectAppUninstallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallUnauthorized creates a ProjectAppUninstallUnauthorized with default headers values
func NewProjectAppUninstallUnauthorized() *ProjectAppUninstallUnauthorized {
	return &ProjectAppUninstallUnauthorized{}
}

/* ProjectAppUninstallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectAppUninstallUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppUninstallUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectAppUninstallUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppUninstallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallForbidden creates a ProjectAppUninstallForbidden with default headers values
func NewProjectAppUninstallForbidden() *ProjectAppUninstallForbidden {
	return &ProjectAppUninstallForbidden{}
}

/* ProjectAppUninstallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectAppUninstallForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppUninstallForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallForbidden  %+v", 403, o.Payload)
}
func (o *ProjectAppUninstallForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppUninstallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallNotFound creates a ProjectAppUninstallNotFound with default headers values
func NewProjectAppUninstallNotFound() *ProjectAppUninstallNotFound {
	return &ProjectAppUninstallNotFound{}
}

/* ProjectAppUninstallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectAppUninstallNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectAppUninstallNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallNotFound  %+v", 404, o.Payload)
}
func (o *ProjectAppUninstallNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectAppUninstallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectAppUninstallInternalServerError creates a ProjectAppUninstallInternalServerError with default headers values
func NewProjectAppUninstallInternalServerError() *ProjectAppUninstallInternalServerError {
	return &ProjectAppUninstallInternalServerError{}
}

/* ProjectAppUninstallInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectAppUninstallInternalServerError struct {
}

func (o *ProjectAppUninstallInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/ProjectApp/uninstall/{projectAppId}][%d] projectAppUninstallInternalServerError ", 500)
}

func (o *ProjectAppUninstallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
