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

// KubernetesRestartDaemonSetReader is a Reader for the KubernetesRestartDaemonSet structure.
type KubernetesRestartDaemonSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesRestartDaemonSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesRestartDaemonSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesRestartDaemonSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesRestartDaemonSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesRestartDaemonSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesRestartDaemonSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubernetesRestartDaemonSetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesRestartDaemonSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesRestartDaemonSetOK creates a KubernetesRestartDaemonSetOK with default headers values
func NewKubernetesRestartDaemonSetOK() *KubernetesRestartDaemonSetOK {
	return &KubernetesRestartDaemonSetOK{}
}

/* KubernetesRestartDaemonSetOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesRestartDaemonSetOK struct {
	Payload models.Unit
}

func (o *KubernetesRestartDaemonSetOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetOK  %+v", 200, o.Payload)
}
func (o *KubernetesRestartDaemonSetOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetBadRequest creates a KubernetesRestartDaemonSetBadRequest with default headers values
func NewKubernetesRestartDaemonSetBadRequest() *KubernetesRestartDaemonSetBadRequest {
	return &KubernetesRestartDaemonSetBadRequest{}
}

/* KubernetesRestartDaemonSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesRestartDaemonSetBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesRestartDaemonSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesRestartDaemonSetBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetUnauthorized creates a KubernetesRestartDaemonSetUnauthorized with default headers values
func NewKubernetesRestartDaemonSetUnauthorized() *KubernetesRestartDaemonSetUnauthorized {
	return &KubernetesRestartDaemonSetUnauthorized{}
}

/* KubernetesRestartDaemonSetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesRestartDaemonSetUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesRestartDaemonSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesRestartDaemonSetUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetForbidden creates a KubernetesRestartDaemonSetForbidden with default headers values
func NewKubernetesRestartDaemonSetForbidden() *KubernetesRestartDaemonSetForbidden {
	return &KubernetesRestartDaemonSetForbidden{}
}

/* KubernetesRestartDaemonSetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesRestartDaemonSetForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesRestartDaemonSetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesRestartDaemonSetForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetNotFound creates a KubernetesRestartDaemonSetNotFound with default headers values
func NewKubernetesRestartDaemonSetNotFound() *KubernetesRestartDaemonSetNotFound {
	return &KubernetesRestartDaemonSetNotFound{}
}

/* KubernetesRestartDaemonSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesRestartDaemonSetNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesRestartDaemonSetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesRestartDaemonSetNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetTooManyRequests creates a KubernetesRestartDaemonSetTooManyRequests with default headers values
func NewKubernetesRestartDaemonSetTooManyRequests() *KubernetesRestartDaemonSetTooManyRequests {
	return &KubernetesRestartDaemonSetTooManyRequests{}
}

/* KubernetesRestartDaemonSetTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type KubernetesRestartDaemonSetTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesRestartDaemonSetTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetTooManyRequests  %+v", 429, o.Payload)
}
func (o *KubernetesRestartDaemonSetTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetInternalServerError creates a KubernetesRestartDaemonSetInternalServerError with default headers values
func NewKubernetesRestartDaemonSetInternalServerError() *KubernetesRestartDaemonSetInternalServerError {
	return &KubernetesRestartDaemonSetInternalServerError{}
}

/* KubernetesRestartDaemonSetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesRestartDaemonSetInternalServerError struct {
}

func (o *KubernetesRestartDaemonSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetInternalServerError ", 500)
}

func (o *KubernetesRestartDaemonSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
