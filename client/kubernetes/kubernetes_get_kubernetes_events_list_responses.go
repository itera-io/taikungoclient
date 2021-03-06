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

// KubernetesGetKubernetesEventsListReader is a Reader for the KubernetesGetKubernetesEventsList structure.
type KubernetesGetKubernetesEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetKubernetesEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetKubernetesEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetKubernetesEventsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetKubernetesEventsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetKubernetesEventsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetKubernetesEventsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetKubernetesEventsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetKubernetesEventsListOK creates a KubernetesGetKubernetesEventsListOK with default headers values
func NewKubernetesGetKubernetesEventsListOK() *KubernetesGetKubernetesEventsListOK {
	return &KubernetesGetKubernetesEventsListOK{}
}

/* KubernetesGetKubernetesEventsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetKubernetesEventsListOK struct {
	Payload *models.KubernetesEventsList
}

func (o *KubernetesGetKubernetesEventsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/events][%d] kubernetesGetKubernetesEventsListOK  %+v", 200, o.Payload)
}
func (o *KubernetesGetKubernetesEventsListOK) GetPayload() *models.KubernetesEventsList {
	return o.Payload
}

func (o *KubernetesGetKubernetesEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesEventsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesEventsListBadRequest creates a KubernetesGetKubernetesEventsListBadRequest with default headers values
func NewKubernetesGetKubernetesEventsListBadRequest() *KubernetesGetKubernetesEventsListBadRequest {
	return &KubernetesGetKubernetesEventsListBadRequest{}
}

/* KubernetesGetKubernetesEventsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetKubernetesEventsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesGetKubernetesEventsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/events][%d] kubernetesGetKubernetesEventsListBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesGetKubernetesEventsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetKubernetesEventsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesEventsListUnauthorized creates a KubernetesGetKubernetesEventsListUnauthorized with default headers values
func NewKubernetesGetKubernetesEventsListUnauthorized() *KubernetesGetKubernetesEventsListUnauthorized {
	return &KubernetesGetKubernetesEventsListUnauthorized{}
}

/* KubernetesGetKubernetesEventsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetKubernetesEventsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetKubernetesEventsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/events][%d] kubernetesGetKubernetesEventsListUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesGetKubernetesEventsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetKubernetesEventsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesEventsListForbidden creates a KubernetesGetKubernetesEventsListForbidden with default headers values
func NewKubernetesGetKubernetesEventsListForbidden() *KubernetesGetKubernetesEventsListForbidden {
	return &KubernetesGetKubernetesEventsListForbidden{}
}

/* KubernetesGetKubernetesEventsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetKubernetesEventsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetKubernetesEventsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/events][%d] kubernetesGetKubernetesEventsListForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesGetKubernetesEventsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetKubernetesEventsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesEventsListNotFound creates a KubernetesGetKubernetesEventsListNotFound with default headers values
func NewKubernetesGetKubernetesEventsListNotFound() *KubernetesGetKubernetesEventsListNotFound {
	return &KubernetesGetKubernetesEventsListNotFound{}
}

/* KubernetesGetKubernetesEventsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetKubernetesEventsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetKubernetesEventsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/events][%d] kubernetesGetKubernetesEventsListNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesGetKubernetesEventsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetKubernetesEventsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubernetesEventsListInternalServerError creates a KubernetesGetKubernetesEventsListInternalServerError with default headers values
func NewKubernetesGetKubernetesEventsListInternalServerError() *KubernetesGetKubernetesEventsListInternalServerError {
	return &KubernetesGetKubernetesEventsListInternalServerError{}
}

/* KubernetesGetKubernetesEventsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetKubernetesEventsListInternalServerError struct {
}

func (o *KubernetesGetKubernetesEventsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/events][%d] kubernetesGetKubernetesEventsListInternalServerError ", 500)
}

func (o *KubernetesGetKubernetesEventsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
