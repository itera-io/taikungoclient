// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubernetesCreateKubernetesEventReader is a Reader for the KubernetesCreateKubernetesEvent structure.
type KubernetesCreateKubernetesEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesCreateKubernetesEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesCreateKubernetesEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesCreateKubernetesEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesCreateKubernetesEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesCreateKubernetesEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesCreateKubernetesEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesCreateKubernetesEventInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesCreateKubernetesEventOK creates a KubernetesCreateKubernetesEventOK with default headers values
func NewKubernetesCreateKubernetesEventOK() *KubernetesCreateKubernetesEventOK {
	return &KubernetesCreateKubernetesEventOK{}
}

/* KubernetesCreateKubernetesEventOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesCreateKubernetesEventOK struct {
	Payload models.Unit
}

func (o *KubernetesCreateKubernetesEventOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/event/{projectId}][%d] kubernetesCreateKubernetesEventOK  %+v", 200, o.Payload)
}
func (o *KubernetesCreateKubernetesEventOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesCreateKubernetesEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesCreateKubernetesEventBadRequest creates a KubernetesCreateKubernetesEventBadRequest with default headers values
func NewKubernetesCreateKubernetesEventBadRequest() *KubernetesCreateKubernetesEventBadRequest {
	return &KubernetesCreateKubernetesEventBadRequest{}
}

/* KubernetesCreateKubernetesEventBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesCreateKubernetesEventBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesCreateKubernetesEventBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/event/{projectId}][%d] kubernetesCreateKubernetesEventBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesCreateKubernetesEventBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesCreateKubernetesEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesCreateKubernetesEventUnauthorized creates a KubernetesCreateKubernetesEventUnauthorized with default headers values
func NewKubernetesCreateKubernetesEventUnauthorized() *KubernetesCreateKubernetesEventUnauthorized {
	return &KubernetesCreateKubernetesEventUnauthorized{}
}

/* KubernetesCreateKubernetesEventUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesCreateKubernetesEventUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesCreateKubernetesEventUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/event/{projectId}][%d] kubernetesCreateKubernetesEventUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesCreateKubernetesEventUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesCreateKubernetesEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesCreateKubernetesEventForbidden creates a KubernetesCreateKubernetesEventForbidden with default headers values
func NewKubernetesCreateKubernetesEventForbidden() *KubernetesCreateKubernetesEventForbidden {
	return &KubernetesCreateKubernetesEventForbidden{}
}

/* KubernetesCreateKubernetesEventForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesCreateKubernetesEventForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesCreateKubernetesEventForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/event/{projectId}][%d] kubernetesCreateKubernetesEventForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesCreateKubernetesEventForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesCreateKubernetesEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesCreateKubernetesEventNotFound creates a KubernetesCreateKubernetesEventNotFound with default headers values
func NewKubernetesCreateKubernetesEventNotFound() *KubernetesCreateKubernetesEventNotFound {
	return &KubernetesCreateKubernetesEventNotFound{}
}

/* KubernetesCreateKubernetesEventNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesCreateKubernetesEventNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesCreateKubernetesEventNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/event/{projectId}][%d] kubernetesCreateKubernetesEventNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesCreateKubernetesEventNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesCreateKubernetesEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesCreateKubernetesEventInternalServerError creates a KubernetesCreateKubernetesEventInternalServerError with default headers values
func NewKubernetesCreateKubernetesEventInternalServerError() *KubernetesCreateKubernetesEventInternalServerError {
	return &KubernetesCreateKubernetesEventInternalServerError{}
}

/* KubernetesCreateKubernetesEventInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesCreateKubernetesEventInternalServerError struct {
}

func (o *KubernetesCreateKubernetesEventInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/event/{projectId}][%d] kubernetesCreateKubernetesEventInternalServerError ", 500)
}

func (o *KubernetesCreateKubernetesEventInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
