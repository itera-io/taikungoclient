// Code generated by go-swagger; DO NOT EDIT.

package project_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectGroupsUpdateReader is a Reader for the ProjectGroupsUpdate structure.
type ProjectGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectGroupsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectGroupsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectGroupsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectGroupsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectGroupsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectGroupsUpdateOK creates a ProjectGroupsUpdateOK with default headers values
func NewProjectGroupsUpdateOK() *ProjectGroupsUpdateOK {
	return &ProjectGroupsUpdateOK{}
}

/* ProjectGroupsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type ProjectGroupsUpdateOK struct {
	Payload models.Unit
}

func (o *ProjectGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/ProjectGroups/update][%d] projectGroupsUpdateOK  %+v", 200, o.Payload)
}
func (o *ProjectGroupsUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsUpdateBadRequest creates a ProjectGroupsUpdateBadRequest with default headers values
func NewProjectGroupsUpdateBadRequest() *ProjectGroupsUpdateBadRequest {
	return &ProjectGroupsUpdateBadRequest{}
}

/* ProjectGroupsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectGroupsUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectGroupsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/ProjectGroups/update][%d] projectGroupsUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectGroupsUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsUpdateUnauthorized creates a ProjectGroupsUpdateUnauthorized with default headers values
func NewProjectGroupsUpdateUnauthorized() *ProjectGroupsUpdateUnauthorized {
	return &ProjectGroupsUpdateUnauthorized{}
}

/* ProjectGroupsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectGroupsUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectGroupsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/ProjectGroups/update][%d] projectGroupsUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectGroupsUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsUpdateForbidden creates a ProjectGroupsUpdateForbidden with default headers values
func NewProjectGroupsUpdateForbidden() *ProjectGroupsUpdateForbidden {
	return &ProjectGroupsUpdateForbidden{}
}

/* ProjectGroupsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectGroupsUpdateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectGroupsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/ProjectGroups/update][%d] projectGroupsUpdateForbidden  %+v", 403, o.Payload)
}
func (o *ProjectGroupsUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsUpdateNotFound creates a ProjectGroupsUpdateNotFound with default headers values
func NewProjectGroupsUpdateNotFound() *ProjectGroupsUpdateNotFound {
	return &ProjectGroupsUpdateNotFound{}
}

/* ProjectGroupsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectGroupsUpdateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectGroupsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/ProjectGroups/update][%d] projectGroupsUpdateNotFound  %+v", 404, o.Payload)
}
func (o *ProjectGroupsUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectGroupsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGroupsUpdateInternalServerError creates a ProjectGroupsUpdateInternalServerError with default headers values
func NewProjectGroupsUpdateInternalServerError() *ProjectGroupsUpdateInternalServerError {
	return &ProjectGroupsUpdateInternalServerError{}
}

/* ProjectGroupsUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectGroupsUpdateInternalServerError struct {
}

func (o *ProjectGroupsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/ProjectGroups/update][%d] projectGroupsUpdateInternalServerError ", 500)
}

func (o *ProjectGroupsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
