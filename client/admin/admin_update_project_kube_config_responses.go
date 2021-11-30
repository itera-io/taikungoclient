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

// AdminUpdateProjectKubeConfigReader is a Reader for the AdminUpdateProjectKubeConfig structure.
type AdminUpdateProjectKubeConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateProjectKubeConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateProjectKubeConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateProjectKubeConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateProjectKubeConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateProjectKubeConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUpdateProjectKubeConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateProjectKubeConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateProjectKubeConfigOK creates a AdminUpdateProjectKubeConfigOK with default headers values
func NewAdminUpdateProjectKubeConfigOK() *AdminUpdateProjectKubeConfigOK {
	return &AdminUpdateProjectKubeConfigOK{}
}

/* AdminUpdateProjectKubeConfigOK describes a response with status code 200, with default header values.

Success
*/
type AdminUpdateProjectKubeConfigOK struct {
	Payload models.Unit
}

func (o *AdminUpdateProjectKubeConfigOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigOK  %+v", 200, o.Payload)
}
func (o *AdminUpdateProjectKubeConfigOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigBadRequest creates a AdminUpdateProjectKubeConfigBadRequest with default headers values
func NewAdminUpdateProjectKubeConfigBadRequest() *AdminUpdateProjectKubeConfigBadRequest {
	return &AdminUpdateProjectKubeConfigBadRequest{}
}

/* AdminUpdateProjectKubeConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUpdateProjectKubeConfigBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AdminUpdateProjectKubeConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigBadRequest  %+v", 400, o.Payload)
}
func (o *AdminUpdateProjectKubeConfigBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigUnauthorized creates a AdminUpdateProjectKubeConfigUnauthorized with default headers values
func NewAdminUpdateProjectKubeConfigUnauthorized() *AdminUpdateProjectKubeConfigUnauthorized {
	return &AdminUpdateProjectKubeConfigUnauthorized{}
}

/* AdminUpdateProjectKubeConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUpdateProjectKubeConfigUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateProjectKubeConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminUpdateProjectKubeConfigUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigForbidden creates a AdminUpdateProjectKubeConfigForbidden with default headers values
func NewAdminUpdateProjectKubeConfigForbidden() *AdminUpdateProjectKubeConfigForbidden {
	return &AdminUpdateProjectKubeConfigForbidden{}
}

/* AdminUpdateProjectKubeConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUpdateProjectKubeConfigForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateProjectKubeConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigForbidden  %+v", 403, o.Payload)
}
func (o *AdminUpdateProjectKubeConfigForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigNotFound creates a AdminUpdateProjectKubeConfigNotFound with default headers values
func NewAdminUpdateProjectKubeConfigNotFound() *AdminUpdateProjectKubeConfigNotFound {
	return &AdminUpdateProjectKubeConfigNotFound{}
}

/* AdminUpdateProjectKubeConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUpdateProjectKubeConfigNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateProjectKubeConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigNotFound  %+v", 404, o.Payload)
}
func (o *AdminUpdateProjectKubeConfigNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateProjectKubeConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateProjectKubeConfigInternalServerError creates a AdminUpdateProjectKubeConfigInternalServerError with default headers values
func NewAdminUpdateProjectKubeConfigInternalServerError() *AdminUpdateProjectKubeConfigInternalServerError {
	return &AdminUpdateProjectKubeConfigInternalServerError{}
}

/* AdminUpdateProjectKubeConfigInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUpdateProjectKubeConfigInternalServerError struct {
}

func (o *AdminUpdateProjectKubeConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/projects/update/kubeconfig][%d] adminUpdateProjectKubeConfigInternalServerError ", 500)
}

func (o *AdminUpdateProjectKubeConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
