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

// KubernetesDescribeDeploymentReader is a Reader for the KubernetesDescribeDeployment structure.
type KubernetesDescribeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeDeploymentOK creates a KubernetesDescribeDeploymentOK with default headers values
func NewKubernetesDescribeDeploymentOK() *KubernetesDescribeDeploymentOK {
	return &KubernetesDescribeDeploymentOK{}
}

/* KubernetesDescribeDeploymentOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeDeploymentOK struct {
	Payload string
}

func (o *KubernetesDescribeDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentOK  %+v", 200, o.Payload)
}
func (o *KubernetesDescribeDeploymentOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentBadRequest creates a KubernetesDescribeDeploymentBadRequest with default headers values
func NewKubernetesDescribeDeploymentBadRequest() *KubernetesDescribeDeploymentBadRequest {
	return &KubernetesDescribeDeploymentBadRequest{}
}

/* KubernetesDescribeDeploymentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeDeploymentBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesDescribeDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesDescribeDeploymentBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentUnauthorized creates a KubernetesDescribeDeploymentUnauthorized with default headers values
func NewKubernetesDescribeDeploymentUnauthorized() *KubernetesDescribeDeploymentUnauthorized {
	return &KubernetesDescribeDeploymentUnauthorized{}
}

/* KubernetesDescribeDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeDeploymentUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesDescribeDeploymentUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentForbidden creates a KubernetesDescribeDeploymentForbidden with default headers values
func NewKubernetesDescribeDeploymentForbidden() *KubernetesDescribeDeploymentForbidden {
	return &KubernetesDescribeDeploymentForbidden{}
}

/* KubernetesDescribeDeploymentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeDeploymentForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesDescribeDeploymentForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentNotFound creates a KubernetesDescribeDeploymentNotFound with default headers values
func NewKubernetesDescribeDeploymentNotFound() *KubernetesDescribeDeploymentNotFound {
	return &KubernetesDescribeDeploymentNotFound{}
}

/* KubernetesDescribeDeploymentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeDeploymentNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesDescribeDeploymentNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentInternalServerError creates a KubernetesDescribeDeploymentInternalServerError with default headers values
func NewKubernetesDescribeDeploymentInternalServerError() *KubernetesDescribeDeploymentInternalServerError {
	return &KubernetesDescribeDeploymentInternalServerError{}
}

/* KubernetesDescribeDeploymentInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeDeploymentInternalServerError struct {
}

func (o *KubernetesDescribeDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentInternalServerError ", 500)
}

func (o *KubernetesDescribeDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
