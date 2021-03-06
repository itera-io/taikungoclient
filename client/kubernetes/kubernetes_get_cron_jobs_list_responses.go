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

// KubernetesGetCronJobsListReader is a Reader for the KubernetesGetCronJobsList structure.
type KubernetesGetCronJobsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetCronJobsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetCronJobsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetCronJobsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetCronJobsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetCronJobsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetCronJobsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetCronJobsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetCronJobsListOK creates a KubernetesGetCronJobsListOK with default headers values
func NewKubernetesGetCronJobsListOK() *KubernetesGetCronJobsListOK {
	return &KubernetesGetCronJobsListOK{}
}

/* KubernetesGetCronJobsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetCronJobsListOK struct {
	Payload *models.KubernetesCronJobsList
}

func (o *KubernetesGetCronJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListOK  %+v", 200, o.Payload)
}
func (o *KubernetesGetCronJobsListOK) GetPayload() *models.KubernetesCronJobsList {
	return o.Payload
}

func (o *KubernetesGetCronJobsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesCronJobsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCronJobsListBadRequest creates a KubernetesGetCronJobsListBadRequest with default headers values
func NewKubernetesGetCronJobsListBadRequest() *KubernetesGetCronJobsListBadRequest {
	return &KubernetesGetCronJobsListBadRequest{}
}

/* KubernetesGetCronJobsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetCronJobsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesGetCronJobsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesGetCronJobsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCronJobsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCronJobsListUnauthorized creates a KubernetesGetCronJobsListUnauthorized with default headers values
func NewKubernetesGetCronJobsListUnauthorized() *KubernetesGetCronJobsListUnauthorized {
	return &KubernetesGetCronJobsListUnauthorized{}
}

/* KubernetesGetCronJobsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetCronJobsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetCronJobsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesGetCronJobsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCronJobsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCronJobsListForbidden creates a KubernetesGetCronJobsListForbidden with default headers values
func NewKubernetesGetCronJobsListForbidden() *KubernetesGetCronJobsListForbidden {
	return &KubernetesGetCronJobsListForbidden{}
}

/* KubernetesGetCronJobsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetCronJobsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetCronJobsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesGetCronJobsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCronJobsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCronJobsListNotFound creates a KubernetesGetCronJobsListNotFound with default headers values
func NewKubernetesGetCronJobsListNotFound() *KubernetesGetCronJobsListNotFound {
	return &KubernetesGetCronJobsListNotFound{}
}

/* KubernetesGetCronJobsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetCronJobsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesGetCronJobsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesGetCronJobsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCronJobsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCronJobsListInternalServerError creates a KubernetesGetCronJobsListInternalServerError with default headers values
func NewKubernetesGetCronJobsListInternalServerError() *KubernetesGetCronJobsListInternalServerError {
	return &KubernetesGetCronJobsListInternalServerError{}
}

/* KubernetesGetCronJobsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetCronJobsListInternalServerError struct {
}

func (o *KubernetesGetCronJobsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListInternalServerError ", 500)
}

func (o *KubernetesGetCronJobsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
