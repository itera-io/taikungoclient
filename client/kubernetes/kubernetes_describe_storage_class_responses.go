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

// KubernetesDescribeStorageClassReader is a Reader for the KubernetesDescribeStorageClass structure.
type KubernetesDescribeStorageClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeStorageClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeStorageClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeStorageClassBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeStorageClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeStorageClassForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeStorageClassNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeStorageClassInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeStorageClassOK creates a KubernetesDescribeStorageClassOK with default headers values
func NewKubernetesDescribeStorageClassOK() *KubernetesDescribeStorageClassOK {
	return &KubernetesDescribeStorageClassOK{}
}

/* KubernetesDescribeStorageClassOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeStorageClassOK struct {
	Payload string
}

func (o *KubernetesDescribeStorageClassOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassOK  %+v", 200, o.Payload)
}
func (o *KubernetesDescribeStorageClassOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeStorageClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStorageClassBadRequest creates a KubernetesDescribeStorageClassBadRequest with default headers values
func NewKubernetesDescribeStorageClassBadRequest() *KubernetesDescribeStorageClassBadRequest {
	return &KubernetesDescribeStorageClassBadRequest{}
}

/* KubernetesDescribeStorageClassBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeStorageClassBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesDescribeStorageClassBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesDescribeStorageClassBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStorageClassBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStorageClassUnauthorized creates a KubernetesDescribeStorageClassUnauthorized with default headers values
func NewKubernetesDescribeStorageClassUnauthorized() *KubernetesDescribeStorageClassUnauthorized {
	return &KubernetesDescribeStorageClassUnauthorized{}
}

/* KubernetesDescribeStorageClassUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeStorageClassUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeStorageClassUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesDescribeStorageClassUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStorageClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStorageClassForbidden creates a KubernetesDescribeStorageClassForbidden with default headers values
func NewKubernetesDescribeStorageClassForbidden() *KubernetesDescribeStorageClassForbidden {
	return &KubernetesDescribeStorageClassForbidden{}
}

/* KubernetesDescribeStorageClassForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeStorageClassForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeStorageClassForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesDescribeStorageClassForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStorageClassForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStorageClassNotFound creates a KubernetesDescribeStorageClassNotFound with default headers values
func NewKubernetesDescribeStorageClassNotFound() *KubernetesDescribeStorageClassNotFound {
	return &KubernetesDescribeStorageClassNotFound{}
}

/* KubernetesDescribeStorageClassNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeStorageClassNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesDescribeStorageClassNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesDescribeStorageClassNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStorageClassNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStorageClassInternalServerError creates a KubernetesDescribeStorageClassInternalServerError with default headers values
func NewKubernetesDescribeStorageClassInternalServerError() *KubernetesDescribeStorageClassInternalServerError {
	return &KubernetesDescribeStorageClassInternalServerError{}
}

/* KubernetesDescribeStorageClassInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeStorageClassInternalServerError struct {
}

func (o *KubernetesDescribeStorageClassInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassInternalServerError ", 500)
}

func (o *KubernetesDescribeStorageClassInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
