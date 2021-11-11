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

// KubernetesGetJobsListReader is a Reader for the KubernetesGetJobsList structure.
type KubernetesGetJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetJobsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetJobsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetJobsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetJobsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetJobsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetJobsListOK creates a KubernetesGetJobsListOK with default headers values
func NewKubernetesGetJobsListOK() *KubernetesGetJobsListOK {
	return &KubernetesGetJobsListOK{}
}

/* KubernetesGetJobsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetJobsListOK struct {
	Payload *models.KubernetesJobList
}

func (o *KubernetesGetJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/jobs][%d] kubernetesGetJobsListOK  %+v", 200, o.Payload)
}
func (o *KubernetesGetJobsListOK) GetPayload() *models.KubernetesJobList {
	return o.Payload
}

func (o *KubernetesGetJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesJobList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetJobsListBadRequest creates a KubernetesGetJobsListBadRequest with default headers values
func NewKubernetesGetJobsListBadRequest() *KubernetesGetJobsListBadRequest {
	return &KubernetesGetJobsListBadRequest{}
}

/* KubernetesGetJobsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetJobsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesGetJobsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/jobs][%d] kubernetesGetJobsListBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesGetJobsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetJobsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetJobsListUnauthorized creates a KubernetesGetJobsListUnauthorized with default headers values
func NewKubernetesGetJobsListUnauthorized() *KubernetesGetJobsListUnauthorized {
	return &KubernetesGetJobsListUnauthorized{}
}

/* KubernetesGetJobsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetJobsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetJobsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/jobs][%d] kubernetesGetJobsListUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesGetJobsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetJobsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetJobsListForbidden creates a KubernetesGetJobsListForbidden with default headers values
func NewKubernetesGetJobsListForbidden() *KubernetesGetJobsListForbidden {
	return &KubernetesGetJobsListForbidden{}
}

/* KubernetesGetJobsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetJobsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetJobsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/jobs][%d] kubernetesGetJobsListForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesGetJobsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetJobsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetJobsListNotFound creates a KubernetesGetJobsListNotFound with default headers values
func NewKubernetesGetJobsListNotFound() *KubernetesGetJobsListNotFound {
	return &KubernetesGetJobsListNotFound{}
}

/* KubernetesGetJobsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetJobsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetJobsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/jobs][%d] kubernetesGetJobsListNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesGetJobsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetJobsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetJobsListInternalServerError creates a KubernetesGetJobsListInternalServerError with default headers values
func NewKubernetesGetJobsListInternalServerError() *KubernetesGetJobsListInternalServerError {
	return &KubernetesGetJobsListInternalServerError{}
}

/* KubernetesGetJobsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetJobsListInternalServerError struct {
}

func (o *KubernetesGetJobsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/jobs][%d] kubernetesGetJobsListInternalServerError ", 500)
}

func (o *KubernetesGetJobsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
