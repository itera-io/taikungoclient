// Code generated by go-swagger; DO NOT EDIT.

package project_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectQuotasEditReader is a Reader for the ProjectQuotasEdit structure.
type ProjectQuotasEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectQuotasEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectQuotasEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectQuotasEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectQuotasEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectQuotasEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectQuotasEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectQuotasEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectQuotasEditOK creates a ProjectQuotasEditOK with default headers values
func NewProjectQuotasEditOK() *ProjectQuotasEditOK {
	return &ProjectQuotasEditOK{}
}

/* ProjectQuotasEditOK describes a response with status code 200, with default header values.

Success
*/
type ProjectQuotasEditOK struct {
	Payload models.Unit
}

func (o *ProjectQuotasEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectQuotas/update/{quotaId}][%d] projectQuotasEditOK  %+v", 200, o.Payload)
}
func (o *ProjectQuotasEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectQuotasEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasEditBadRequest creates a ProjectQuotasEditBadRequest with default headers values
func NewProjectQuotasEditBadRequest() *ProjectQuotasEditBadRequest {
	return &ProjectQuotasEditBadRequest{}
}

/* ProjectQuotasEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectQuotasEditBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectQuotasEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectQuotas/update/{quotaId}][%d] projectQuotasEditBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectQuotasEditBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasEditUnauthorized creates a ProjectQuotasEditUnauthorized with default headers values
func NewProjectQuotasEditUnauthorized() *ProjectQuotasEditUnauthorized {
	return &ProjectQuotasEditUnauthorized{}
}

/* ProjectQuotasEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectQuotasEditUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectQuotasEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectQuotas/update/{quotaId}][%d] projectQuotasEditUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectQuotasEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasEditForbidden creates a ProjectQuotasEditForbidden with default headers values
func NewProjectQuotasEditForbidden() *ProjectQuotasEditForbidden {
	return &ProjectQuotasEditForbidden{}
}

/* ProjectQuotasEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectQuotasEditForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectQuotasEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectQuotas/update/{quotaId}][%d] projectQuotasEditForbidden  %+v", 403, o.Payload)
}
func (o *ProjectQuotasEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasEditNotFound creates a ProjectQuotasEditNotFound with default headers values
func NewProjectQuotasEditNotFound() *ProjectQuotasEditNotFound {
	return &ProjectQuotasEditNotFound{}
}

/* ProjectQuotasEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectQuotasEditNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectQuotasEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectQuotas/update/{quotaId}][%d] projectQuotasEditNotFound  %+v", 404, o.Payload)
}
func (o *ProjectQuotasEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectQuotasEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectQuotasEditInternalServerError creates a ProjectQuotasEditInternalServerError with default headers values
func NewProjectQuotasEditInternalServerError() *ProjectQuotasEditInternalServerError {
	return &ProjectQuotasEditInternalServerError{}
}

/* ProjectQuotasEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectQuotasEditInternalServerError struct {
}

func (o *ProjectQuotasEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectQuotas/update/{quotaId}][%d] projectQuotasEditInternalServerError ", 500)
}

func (o *ProjectQuotasEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
