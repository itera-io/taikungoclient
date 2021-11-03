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

// KubernetesUpdateKubernetesAlertReader is a Reader for the KubernetesUpdateKubernetesAlert structure.
type KubernetesUpdateKubernetesAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesUpdateKubernetesAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesUpdateKubernetesAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesUpdateKubernetesAlertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesUpdateKubernetesAlertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesUpdateKubernetesAlertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesUpdateKubernetesAlertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubernetesUpdateKubernetesAlertTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesUpdateKubernetesAlertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesUpdateKubernetesAlertOK creates a KubernetesUpdateKubernetesAlertOK with default headers values
func NewKubernetesUpdateKubernetesAlertOK() *KubernetesUpdateKubernetesAlertOK {
	return &KubernetesUpdateKubernetesAlertOK{}
}

/* KubernetesUpdateKubernetesAlertOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesUpdateKubernetesAlertOK struct {
	Payload models.Unit
}

func (o *KubernetesUpdateKubernetesAlertOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertOK  %+v", 200, o.Payload)
}
func (o *KubernetesUpdateKubernetesAlertOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesUpdateKubernetesAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesUpdateKubernetesAlertBadRequest creates a KubernetesUpdateKubernetesAlertBadRequest with default headers values
func NewKubernetesUpdateKubernetesAlertBadRequest() *KubernetesUpdateKubernetesAlertBadRequest {
	return &KubernetesUpdateKubernetesAlertBadRequest{}
}

/* KubernetesUpdateKubernetesAlertBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesUpdateKubernetesAlertBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesUpdateKubernetesAlertBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesUpdateKubernetesAlertBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesUpdateKubernetesAlertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesUpdateKubernetesAlertUnauthorized creates a KubernetesUpdateKubernetesAlertUnauthorized with default headers values
func NewKubernetesUpdateKubernetesAlertUnauthorized() *KubernetesUpdateKubernetesAlertUnauthorized {
	return &KubernetesUpdateKubernetesAlertUnauthorized{}
}

/* KubernetesUpdateKubernetesAlertUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesUpdateKubernetesAlertUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesUpdateKubernetesAlertUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesUpdateKubernetesAlertUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesUpdateKubernetesAlertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesUpdateKubernetesAlertForbidden creates a KubernetesUpdateKubernetesAlertForbidden with default headers values
func NewKubernetesUpdateKubernetesAlertForbidden() *KubernetesUpdateKubernetesAlertForbidden {
	return &KubernetesUpdateKubernetesAlertForbidden{}
}

/* KubernetesUpdateKubernetesAlertForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesUpdateKubernetesAlertForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesUpdateKubernetesAlertForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesUpdateKubernetesAlertForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesUpdateKubernetesAlertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesUpdateKubernetesAlertNotFound creates a KubernetesUpdateKubernetesAlertNotFound with default headers values
func NewKubernetesUpdateKubernetesAlertNotFound() *KubernetesUpdateKubernetesAlertNotFound {
	return &KubernetesUpdateKubernetesAlertNotFound{}
}

/* KubernetesUpdateKubernetesAlertNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesUpdateKubernetesAlertNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesUpdateKubernetesAlertNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesUpdateKubernetesAlertNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesUpdateKubernetesAlertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesUpdateKubernetesAlertTooManyRequests creates a KubernetesUpdateKubernetesAlertTooManyRequests with default headers values
func NewKubernetesUpdateKubernetesAlertTooManyRequests() *KubernetesUpdateKubernetesAlertTooManyRequests {
	return &KubernetesUpdateKubernetesAlertTooManyRequests{}
}

/* KubernetesUpdateKubernetesAlertTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type KubernetesUpdateKubernetesAlertTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesUpdateKubernetesAlertTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertTooManyRequests  %+v", 429, o.Payload)
}
func (o *KubernetesUpdateKubernetesAlertTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesUpdateKubernetesAlertTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesUpdateKubernetesAlertInternalServerError creates a KubernetesUpdateKubernetesAlertInternalServerError with default headers values
func NewKubernetesUpdateKubernetesAlertInternalServerError() *KubernetesUpdateKubernetesAlertInternalServerError {
	return &KubernetesUpdateKubernetesAlertInternalServerError{}
}

/* KubernetesUpdateKubernetesAlertInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesUpdateKubernetesAlertInternalServerError struct {
}

func (o *KubernetesUpdateKubernetesAlertInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Kubernetes/updatealert/{alertId}][%d] kubernetesUpdateKubernetesAlertInternalServerError ", 500)
}

func (o *KubernetesUpdateKubernetesAlertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
