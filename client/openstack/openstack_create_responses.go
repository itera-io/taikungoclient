// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpenstackCreateReader is a Reader for the OpenstackCreate structure.
type OpenstackCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackCreateOK creates a OpenstackCreateOK with default headers values
func NewOpenstackCreateOK() *OpenstackCreateOK {
	return &OpenstackCreateOK{}
}

/* OpenstackCreateOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackCreateOK struct {
	Payload *models.APIResponse
}

func (o *OpenstackCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/create][%d] openstackCreateOK  %+v", 200, o.Payload)
}
func (o *OpenstackCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *OpenstackCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackCreateBadRequest creates a OpenstackCreateBadRequest with default headers values
func NewOpenstackCreateBadRequest() *OpenstackCreateBadRequest {
	return &OpenstackCreateBadRequest{}
}

/* OpenstackCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OpenstackCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/create][%d] openstackCreateBadRequest  %+v", 400, o.Payload)
}
func (o *OpenstackCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpenstackCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackCreateUnauthorized creates a OpenstackCreateUnauthorized with default headers values
func NewOpenstackCreateUnauthorized() *OpenstackCreateUnauthorized {
	return &OpenstackCreateUnauthorized{}
}

/* OpenstackCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/create][%d] openstackCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *OpenstackCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackCreateForbidden creates a OpenstackCreateForbidden with default headers values
func NewOpenstackCreateForbidden() *OpenstackCreateForbidden {
	return &OpenstackCreateForbidden{}
}

/* OpenstackCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/create][%d] openstackCreateForbidden  %+v", 403, o.Payload)
}
func (o *OpenstackCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackCreateNotFound creates a OpenstackCreateNotFound with default headers values
func NewOpenstackCreateNotFound() *OpenstackCreateNotFound {
	return &OpenstackCreateNotFound{}
}

/* OpenstackCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/create][%d] openstackCreateNotFound  %+v", 404, o.Payload)
}
func (o *OpenstackCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackCreateInternalServerError creates a OpenstackCreateInternalServerError with default headers values
func NewOpenstackCreateInternalServerError() *OpenstackCreateInternalServerError {
	return &OpenstackCreateInternalServerError{}
}

/* OpenstackCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackCreateInternalServerError struct {
}

func (o *OpenstackCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/create][%d] openstackCreateInternalServerError ", 500)
}

func (o *OpenstackCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
