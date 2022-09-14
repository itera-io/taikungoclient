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

/*
KubernetesGetCronJobsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetCronJobsListOK struct {
	Payload *models.KubernetesCronJobsList
}

// IsSuccess returns true when this kubernetes get cron jobs list o k response has a 2xx status code
func (o *KubernetesGetCronJobsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get cron jobs list o k response has a 3xx status code
func (o *KubernetesGetCronJobsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get cron jobs list o k response has a 4xx status code
func (o *KubernetesGetCronJobsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get cron jobs list o k response has a 5xx status code
func (o *KubernetesGetCronJobsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get cron jobs list o k response a status code equal to that given
func (o *KubernetesGetCronJobsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetCronJobsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetCronJobsListOK) String() string {
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

/*
KubernetesGetCronJobsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetCronJobsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes get cron jobs list bad request response has a 2xx status code
func (o *KubernetesGetCronJobsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get cron jobs list bad request response has a 3xx status code
func (o *KubernetesGetCronJobsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get cron jobs list bad request response has a 4xx status code
func (o *KubernetesGetCronJobsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get cron jobs list bad request response has a 5xx status code
func (o *KubernetesGetCronJobsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get cron jobs list bad request response a status code equal to that given
func (o *KubernetesGetCronJobsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetCronJobsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetCronJobsListBadRequest) String() string {
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

/*
KubernetesGetCronJobsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetCronJobsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get cron jobs list unauthorized response has a 2xx status code
func (o *KubernetesGetCronJobsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get cron jobs list unauthorized response has a 3xx status code
func (o *KubernetesGetCronJobsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get cron jobs list unauthorized response has a 4xx status code
func (o *KubernetesGetCronJobsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get cron jobs list unauthorized response has a 5xx status code
func (o *KubernetesGetCronJobsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get cron jobs list unauthorized response a status code equal to that given
func (o *KubernetesGetCronJobsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetCronJobsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetCronJobsListUnauthorized) String() string {
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

/*
KubernetesGetCronJobsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetCronJobsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get cron jobs list forbidden response has a 2xx status code
func (o *KubernetesGetCronJobsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get cron jobs list forbidden response has a 3xx status code
func (o *KubernetesGetCronJobsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get cron jobs list forbidden response has a 4xx status code
func (o *KubernetesGetCronJobsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get cron jobs list forbidden response has a 5xx status code
func (o *KubernetesGetCronJobsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get cron jobs list forbidden response a status code equal to that given
func (o *KubernetesGetCronJobsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetCronJobsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetCronJobsListForbidden) String() string {
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

/*
KubernetesGetCronJobsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetCronJobsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get cron jobs list not found response has a 2xx status code
func (o *KubernetesGetCronJobsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get cron jobs list not found response has a 3xx status code
func (o *KubernetesGetCronJobsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get cron jobs list not found response has a 4xx status code
func (o *KubernetesGetCronJobsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get cron jobs list not found response has a 5xx status code
func (o *KubernetesGetCronJobsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get cron jobs list not found response a status code equal to that given
func (o *KubernetesGetCronJobsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetCronJobsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetCronJobsListNotFound) String() string {
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

/*
KubernetesGetCronJobsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetCronJobsListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get cron jobs list internal server error response has a 2xx status code
func (o *KubernetesGetCronJobsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get cron jobs list internal server error response has a 3xx status code
func (o *KubernetesGetCronJobsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get cron jobs list internal server error response has a 4xx status code
func (o *KubernetesGetCronJobsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get cron jobs list internal server error response has a 5xx status code
func (o *KubernetesGetCronJobsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get cron jobs list internal server error response a status code equal to that given
func (o *KubernetesGetCronJobsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetCronJobsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListInternalServerError ", 500)
}

func (o *KubernetesGetCronJobsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/cronjobs][%d] kubernetesGetCronJobsListInternalServerError ", 500)
}

func (o *KubernetesGetCronJobsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
