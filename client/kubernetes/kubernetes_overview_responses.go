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

// KubernetesOverviewReader is a Reader for the KubernetesOverview structure.
type KubernetesOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesOverviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesOverviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesOverviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesOverviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubernetesOverviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesOverviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesOverviewOK creates a KubernetesOverviewOK with default headers values
func NewKubernetesOverviewOK() *KubernetesOverviewOK {
	return &KubernetesOverviewOK{}
}

/* KubernetesOverviewOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesOverviewOK struct {
	Payload interface{}
}

func (o *KubernetesOverviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewOK  %+v", 200, o.Payload)
}
func (o *KubernetesOverviewOK) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewBadRequest creates a KubernetesOverviewBadRequest with default headers values
func NewKubernetesOverviewBadRequest() *KubernetesOverviewBadRequest {
	return &KubernetesOverviewBadRequest{}
}

/* KubernetesOverviewBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesOverviewBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesOverviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesOverviewBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesOverviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewUnauthorized creates a KubernetesOverviewUnauthorized with default headers values
func NewKubernetesOverviewUnauthorized() *KubernetesOverviewUnauthorized {
	return &KubernetesOverviewUnauthorized{}
}

/* KubernetesOverviewUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesOverviewUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesOverviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesOverviewUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesOverviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewForbidden creates a KubernetesOverviewForbidden with default headers values
func NewKubernetesOverviewForbidden() *KubernetesOverviewForbidden {
	return &KubernetesOverviewForbidden{}
}

/* KubernetesOverviewForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesOverviewForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesOverviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesOverviewForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesOverviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewNotFound creates a KubernetesOverviewNotFound with default headers values
func NewKubernetesOverviewNotFound() *KubernetesOverviewNotFound {
	return &KubernetesOverviewNotFound{}
}

/* KubernetesOverviewNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesOverviewNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesOverviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesOverviewNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesOverviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewTooManyRequests creates a KubernetesOverviewTooManyRequests with default headers values
func NewKubernetesOverviewTooManyRequests() *KubernetesOverviewTooManyRequests {
	return &KubernetesOverviewTooManyRequests{}
}

/* KubernetesOverviewTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type KubernetesOverviewTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesOverviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewTooManyRequests  %+v", 429, o.Payload)
}
func (o *KubernetesOverviewTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesOverviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewInternalServerError creates a KubernetesOverviewInternalServerError with default headers values
func NewKubernetesOverviewInternalServerError() *KubernetesOverviewInternalServerError {
	return &KubernetesOverviewInternalServerError{}
}

/* KubernetesOverviewInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesOverviewInternalServerError struct {
}

func (o *KubernetesOverviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewInternalServerError ", 500)
}

func (o *KubernetesOverviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
