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

// KubernetesDescribeCronJobReader is a Reader for the KubernetesDescribeCronJob structure.
type KubernetesDescribeCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeCronJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeCronJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeCronJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeCronJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeCronJobOK creates a KubernetesDescribeCronJobOK with default headers values
func NewKubernetesDescribeCronJobOK() *KubernetesDescribeCronJobOK {
	return &KubernetesDescribeCronJobOK{}
}

/*
KubernetesDescribeCronJobOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeCronJobOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe cron job o k response has a 2xx status code
func (o *KubernetesDescribeCronJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe cron job o k response has a 3xx status code
func (o *KubernetesDescribeCronJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe cron job o k response has a 4xx status code
func (o *KubernetesDescribeCronJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe cron job o k response has a 5xx status code
func (o *KubernetesDescribeCronJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe cron job o k response a status code equal to that given
func (o *KubernetesDescribeCronJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeCronJobOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeCronJobOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeCronJobOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeCronJobBadRequest creates a KubernetesDescribeCronJobBadRequest with default headers values
func NewKubernetesDescribeCronJobBadRequest() *KubernetesDescribeCronJobBadRequest {
	return &KubernetesDescribeCronJobBadRequest{}
}

/*
KubernetesDescribeCronJobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeCronJobBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes describe cron job bad request response has a 2xx status code
func (o *KubernetesDescribeCronJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe cron job bad request response has a 3xx status code
func (o *KubernetesDescribeCronJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe cron job bad request response has a 4xx status code
func (o *KubernetesDescribeCronJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe cron job bad request response has a 5xx status code
func (o *KubernetesDescribeCronJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe cron job bad request response a status code equal to that given
func (o *KubernetesDescribeCronJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeCronJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeCronJobBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeCronJobBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeCronJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeCronJobUnauthorized creates a KubernetesDescribeCronJobUnauthorized with default headers values
func NewKubernetesDescribeCronJobUnauthorized() *KubernetesDescribeCronJobUnauthorized {
	return &KubernetesDescribeCronJobUnauthorized{}
}

/*
KubernetesDescribeCronJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeCronJobUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe cron job unauthorized response has a 2xx status code
func (o *KubernetesDescribeCronJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe cron job unauthorized response has a 3xx status code
func (o *KubernetesDescribeCronJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe cron job unauthorized response has a 4xx status code
func (o *KubernetesDescribeCronJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe cron job unauthorized response has a 5xx status code
func (o *KubernetesDescribeCronJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe cron job unauthorized response a status code equal to that given
func (o *KubernetesDescribeCronJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeCronJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeCronJobUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeCronJobForbidden creates a KubernetesDescribeCronJobForbidden with default headers values
func NewKubernetesDescribeCronJobForbidden() *KubernetesDescribeCronJobForbidden {
	return &KubernetesDescribeCronJobForbidden{}
}

/*
KubernetesDescribeCronJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeCronJobForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe cron job forbidden response has a 2xx status code
func (o *KubernetesDescribeCronJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe cron job forbidden response has a 3xx status code
func (o *KubernetesDescribeCronJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe cron job forbidden response has a 4xx status code
func (o *KubernetesDescribeCronJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe cron job forbidden response has a 5xx status code
func (o *KubernetesDescribeCronJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe cron job forbidden response a status code equal to that given
func (o *KubernetesDescribeCronJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeCronJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeCronJobForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeCronJobForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeCronJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeCronJobNotFound creates a KubernetesDescribeCronJobNotFound with default headers values
func NewKubernetesDescribeCronJobNotFound() *KubernetesDescribeCronJobNotFound {
	return &KubernetesDescribeCronJobNotFound{}
}

/*
KubernetesDescribeCronJobNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeCronJobNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe cron job not found response has a 2xx status code
func (o *KubernetesDescribeCronJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe cron job not found response has a 3xx status code
func (o *KubernetesDescribeCronJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe cron job not found response has a 4xx status code
func (o *KubernetesDescribeCronJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe cron job not found response has a 5xx status code
func (o *KubernetesDescribeCronJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe cron job not found response a status code equal to that given
func (o *KubernetesDescribeCronJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeCronJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeCronJobNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeCronJobNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeCronJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeCronJobInternalServerError creates a KubernetesDescribeCronJobInternalServerError with default headers values
func NewKubernetesDescribeCronJobInternalServerError() *KubernetesDescribeCronJobInternalServerError {
	return &KubernetesDescribeCronJobInternalServerError{}
}

/*
KubernetesDescribeCronJobInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeCronJobInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe cron job internal server error response has a 2xx status code
func (o *KubernetesDescribeCronJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe cron job internal server error response has a 3xx status code
func (o *KubernetesDescribeCronJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe cron job internal server error response has a 4xx status code
func (o *KubernetesDescribeCronJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe cron job internal server error response has a 5xx status code
func (o *KubernetesDescribeCronJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe cron job internal server error response a status code equal to that given
func (o *KubernetesDescribeCronJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeCronJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobInternalServerError ", 500)
}

func (o *KubernetesDescribeCronJobInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/cronjob][%d] kubernetesDescribeCronJobInternalServerError ", 500)
}

func (o *KubernetesDescribeCronJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
