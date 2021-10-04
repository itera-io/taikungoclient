// Code generated by go-swagger; DO NOT EDIT.

package kube_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubeConfigCreateReader is a Reader for the KubeConfigCreate structure.
type KubeConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubeConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubeConfigCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubeConfigCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubeConfigCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubeConfigCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubeConfigCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubeConfigCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubeConfigCreateOK creates a KubeConfigCreateOK with default headers values
func NewKubeConfigCreateOK() *KubeConfigCreateOK {
	return &KubeConfigCreateOK{}
}

/* KubeConfigCreateOK describes a response with status code 200, with default header values.

Success
*/
type KubeConfigCreateOK struct {
	Payload models.Unit
}

func (o *KubeConfigCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig][%d] kubeConfigCreateOK  %+v", 200, o.Payload)
}
func (o *KubeConfigCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubeConfigCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigCreateBadRequest creates a KubeConfigCreateBadRequest with default headers values
func NewKubeConfigCreateBadRequest() *KubeConfigCreateBadRequest {
	return &KubeConfigCreateBadRequest{}
}

/* KubeConfigCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubeConfigCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubeConfigCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig][%d] kubeConfigCreateBadRequest  %+v", 400, o.Payload)
}
func (o *KubeConfigCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubeConfigCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigCreateUnauthorized creates a KubeConfigCreateUnauthorized with default headers values
func NewKubeConfigCreateUnauthorized() *KubeConfigCreateUnauthorized {
	return &KubeConfigCreateUnauthorized{}
}

/* KubeConfigCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubeConfigCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig][%d] kubeConfigCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *KubeConfigCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigCreateForbidden creates a KubeConfigCreateForbidden with default headers values
func NewKubeConfigCreateForbidden() *KubeConfigCreateForbidden {
	return &KubeConfigCreateForbidden{}
}

/* KubeConfigCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubeConfigCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig][%d] kubeConfigCreateForbidden  %+v", 403, o.Payload)
}
func (o *KubeConfigCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigCreateNotFound creates a KubeConfigCreateNotFound with default headers values
func NewKubeConfigCreateNotFound() *KubeConfigCreateNotFound {
	return &KubeConfigCreateNotFound{}
}

/* KubeConfigCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubeConfigCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig][%d] kubeConfigCreateNotFound  %+v", 404, o.Payload)
}
func (o *KubeConfigCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigCreateInternalServerError creates a KubeConfigCreateInternalServerError with default headers values
func NewKubeConfigCreateInternalServerError() *KubeConfigCreateInternalServerError {
	return &KubeConfigCreateInternalServerError{}
}

/* KubeConfigCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubeConfigCreateInternalServerError struct {
}

func (o *KubeConfigCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig][%d] kubeConfigCreateInternalServerError ", 500)
}

func (o *KubeConfigCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
