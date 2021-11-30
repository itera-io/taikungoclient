// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminUpdateUserKubeConfigReader is a Reader for the AdminUpdateUserKubeConfig structure.
type AdminUpdateUserKubeConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserKubeConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserKubeConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserKubeConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserKubeConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserKubeConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUpdateUserKubeConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserKubeConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateUserKubeConfigOK creates a AdminUpdateUserKubeConfigOK with default headers values
func NewAdminUpdateUserKubeConfigOK() *AdminUpdateUserKubeConfigOK {
	return &AdminUpdateUserKubeConfigOK{}
}

/* AdminUpdateUserKubeConfigOK describes a response with status code 200, with default header values.

Success
*/
type AdminUpdateUserKubeConfigOK struct {
	Payload models.Unit
}

func (o *AdminUpdateUserKubeConfigOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/userkube][%d] adminUpdateUserKubeConfigOK  %+v", 200, o.Payload)
}
func (o *AdminUpdateUserKubeConfigOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminUpdateUserKubeConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserKubeConfigBadRequest creates a AdminUpdateUserKubeConfigBadRequest with default headers values
func NewAdminUpdateUserKubeConfigBadRequest() *AdminUpdateUserKubeConfigBadRequest {
	return &AdminUpdateUserKubeConfigBadRequest{}
}

/* AdminUpdateUserKubeConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUpdateUserKubeConfigBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AdminUpdateUserKubeConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/userkube][%d] adminUpdateUserKubeConfigBadRequest  %+v", 400, o.Payload)
}
func (o *AdminUpdateUserKubeConfigBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserKubeConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserKubeConfigUnauthorized creates a AdminUpdateUserKubeConfigUnauthorized with default headers values
func NewAdminUpdateUserKubeConfigUnauthorized() *AdminUpdateUserKubeConfigUnauthorized {
	return &AdminUpdateUserKubeConfigUnauthorized{}
}

/* AdminUpdateUserKubeConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUpdateUserKubeConfigUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserKubeConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/userkube][%d] adminUpdateUserKubeConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminUpdateUserKubeConfigUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserKubeConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserKubeConfigForbidden creates a AdminUpdateUserKubeConfigForbidden with default headers values
func NewAdminUpdateUserKubeConfigForbidden() *AdminUpdateUserKubeConfigForbidden {
	return &AdminUpdateUserKubeConfigForbidden{}
}

/* AdminUpdateUserKubeConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUpdateUserKubeConfigForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserKubeConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/userkube][%d] adminUpdateUserKubeConfigForbidden  %+v", 403, o.Payload)
}
func (o *AdminUpdateUserKubeConfigForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserKubeConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserKubeConfigNotFound creates a AdminUpdateUserKubeConfigNotFound with default headers values
func NewAdminUpdateUserKubeConfigNotFound() *AdminUpdateUserKubeConfigNotFound {
	return &AdminUpdateUserKubeConfigNotFound{}
}

/* AdminUpdateUserKubeConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUpdateUserKubeConfigNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserKubeConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/userkube][%d] adminUpdateUserKubeConfigNotFound  %+v", 404, o.Payload)
}
func (o *AdminUpdateUserKubeConfigNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserKubeConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserKubeConfigInternalServerError creates a AdminUpdateUserKubeConfigInternalServerError with default headers values
func NewAdminUpdateUserKubeConfigInternalServerError() *AdminUpdateUserKubeConfigInternalServerError {
	return &AdminUpdateUserKubeConfigInternalServerError{}
}

/* AdminUpdateUserKubeConfigInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUpdateUserKubeConfigInternalServerError struct {
}

func (o *AdminUpdateUserKubeConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/userkube][%d] adminUpdateUserKubeConfigInternalServerError ", 500)
}

func (o *AdminUpdateUserKubeConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
