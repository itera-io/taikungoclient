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

// KubernetesDescribeConfigMapReader is a Reader for the KubernetesDescribeConfigMap structure.
type KubernetesDescribeConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeConfigMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeConfigMapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeConfigMapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeConfigMapOK creates a KubernetesDescribeConfigMapOK with default headers values
func NewKubernetesDescribeConfigMapOK() *KubernetesDescribeConfigMapOK {
	return &KubernetesDescribeConfigMapOK{}
}

/* KubernetesDescribeConfigMapOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeConfigMapOK struct {
	Payload string
}

func (o *KubernetesDescribeConfigMapOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapOK  %+v", 200, o.Payload)
}
func (o *KubernetesDescribeConfigMapOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeConfigMapBadRequest creates a KubernetesDescribeConfigMapBadRequest with default headers values
func NewKubernetesDescribeConfigMapBadRequest() *KubernetesDescribeConfigMapBadRequest {
	return &KubernetesDescribeConfigMapBadRequest{}
}

/* KubernetesDescribeConfigMapBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeConfigMapBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesDescribeConfigMapBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesDescribeConfigMapBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeConfigMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeConfigMapUnauthorized creates a KubernetesDescribeConfigMapUnauthorized with default headers values
func NewKubernetesDescribeConfigMapUnauthorized() *KubernetesDescribeConfigMapUnauthorized {
	return &KubernetesDescribeConfigMapUnauthorized{}
}

/* KubernetesDescribeConfigMapUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeConfigMapUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeConfigMapUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesDescribeConfigMapUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeConfigMapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeConfigMapForbidden creates a KubernetesDescribeConfigMapForbidden with default headers values
func NewKubernetesDescribeConfigMapForbidden() *KubernetesDescribeConfigMapForbidden {
	return &KubernetesDescribeConfigMapForbidden{}
}

/* KubernetesDescribeConfigMapForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeConfigMapForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeConfigMapForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesDescribeConfigMapForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeConfigMapNotFound creates a KubernetesDescribeConfigMapNotFound with default headers values
func NewKubernetesDescribeConfigMapNotFound() *KubernetesDescribeConfigMapNotFound {
	return &KubernetesDescribeConfigMapNotFound{}
}

/* KubernetesDescribeConfigMapNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeConfigMapNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeConfigMapNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesDescribeConfigMapNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeConfigMapInternalServerError creates a KubernetesDescribeConfigMapInternalServerError with default headers values
func NewKubernetesDescribeConfigMapInternalServerError() *KubernetesDescribeConfigMapInternalServerError {
	return &KubernetesDescribeConfigMapInternalServerError{}
}

/* KubernetesDescribeConfigMapInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeConfigMapInternalServerError struct {
}

func (o *KubernetesDescribeConfigMapInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapInternalServerError ", 500)
}

func (o *KubernetesDescribeConfigMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
