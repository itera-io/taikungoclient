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

// KubernetesGetIngressesReader is a Reader for the KubernetesGetIngresses structure.
type KubernetesGetIngressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetIngressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetIngressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetIngressesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetIngressesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetIngressesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetIngressesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubernetesGetIngressesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetIngressesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetIngressesOK creates a KubernetesGetIngressesOK with default headers values
func NewKubernetesGetIngressesOK() *KubernetesGetIngressesOK {
	return &KubernetesGetIngressesOK{}
}

/* KubernetesGetIngressesOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetIngressesOK struct {
	Payload *models.Ingresses
}

func (o *KubernetesGetIngressesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesOK  %+v", 200, o.Payload)
}
func (o *KubernetesGetIngressesOK) GetPayload() *models.Ingresses {
	return o.Payload
}

func (o *KubernetesGetIngressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ingresses)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetIngressesBadRequest creates a KubernetesGetIngressesBadRequest with default headers values
func NewKubernetesGetIngressesBadRequest() *KubernetesGetIngressesBadRequest {
	return &KubernetesGetIngressesBadRequest{}
}

/* KubernetesGetIngressesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetIngressesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesGetIngressesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesGetIngressesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetIngressesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetIngressesUnauthorized creates a KubernetesGetIngressesUnauthorized with default headers values
func NewKubernetesGetIngressesUnauthorized() *KubernetesGetIngressesUnauthorized {
	return &KubernetesGetIngressesUnauthorized{}
}

/* KubernetesGetIngressesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetIngressesUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetIngressesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesGetIngressesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetIngressesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetIngressesForbidden creates a KubernetesGetIngressesForbidden with default headers values
func NewKubernetesGetIngressesForbidden() *KubernetesGetIngressesForbidden {
	return &KubernetesGetIngressesForbidden{}
}

/* KubernetesGetIngressesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetIngressesForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetIngressesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesGetIngressesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetIngressesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetIngressesNotFound creates a KubernetesGetIngressesNotFound with default headers values
func NewKubernetesGetIngressesNotFound() *KubernetesGetIngressesNotFound {
	return &KubernetesGetIngressesNotFound{}
}

/* KubernetesGetIngressesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetIngressesNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetIngressesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesGetIngressesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetIngressesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetIngressesTooManyRequests creates a KubernetesGetIngressesTooManyRequests with default headers values
func NewKubernetesGetIngressesTooManyRequests() *KubernetesGetIngressesTooManyRequests {
	return &KubernetesGetIngressesTooManyRequests{}
}

/* KubernetesGetIngressesTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type KubernetesGetIngressesTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetIngressesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesTooManyRequests  %+v", 429, o.Payload)
}
func (o *KubernetesGetIngressesTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetIngressesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetIngressesInternalServerError creates a KubernetesGetIngressesInternalServerError with default headers values
func NewKubernetesGetIngressesInternalServerError() *KubernetesGetIngressesInternalServerError {
	return &KubernetesGetIngressesInternalServerError{}
}

/* KubernetesGetIngressesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetIngressesInternalServerError struct {
}

func (o *KubernetesGetIngressesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/ingress][%d] kubernetesGetIngressesInternalServerError ", 500)
}

func (o *KubernetesGetIngressesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
