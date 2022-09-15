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

// KubernetesDescribeServiceReader is a Reader for the KubernetesDescribeService structure.
type KubernetesDescribeServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeServiceOK creates a KubernetesDescribeServiceOK with default headers values
func NewKubernetesDescribeServiceOK() *KubernetesDescribeServiceOK {
	return &KubernetesDescribeServiceOK{}
}

/*
KubernetesDescribeServiceOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeServiceOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe service o k response has a 2xx status code
func (o *KubernetesDescribeServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe service o k response has a 3xx status code
func (o *KubernetesDescribeServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe service o k response has a 4xx status code
func (o *KubernetesDescribeServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe service o k response has a 5xx status code
func (o *KubernetesDescribeServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe service o k response a status code equal to that given
func (o *KubernetesDescribeServiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeServiceOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeServiceOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeServiceOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeServiceBadRequest creates a KubernetesDescribeServiceBadRequest with default headers values
func NewKubernetesDescribeServiceBadRequest() *KubernetesDescribeServiceBadRequest {
	return &KubernetesDescribeServiceBadRequest{}
}

/*
KubernetesDescribeServiceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeServiceBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes describe service bad request response has a 2xx status code
func (o *KubernetesDescribeServiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe service bad request response has a 3xx status code
func (o *KubernetesDescribeServiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe service bad request response has a 4xx status code
func (o *KubernetesDescribeServiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe service bad request response has a 5xx status code
func (o *KubernetesDescribeServiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe service bad request response a status code equal to that given
func (o *KubernetesDescribeServiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeServiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeServiceBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeServiceBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeServiceUnauthorized creates a KubernetesDescribeServiceUnauthorized with default headers values
func NewKubernetesDescribeServiceUnauthorized() *KubernetesDescribeServiceUnauthorized {
	return &KubernetesDescribeServiceUnauthorized{}
}

/*
KubernetesDescribeServiceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeServiceUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe service unauthorized response has a 2xx status code
func (o *KubernetesDescribeServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe service unauthorized response has a 3xx status code
func (o *KubernetesDescribeServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe service unauthorized response has a 4xx status code
func (o *KubernetesDescribeServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe service unauthorized response has a 5xx status code
func (o *KubernetesDescribeServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe service unauthorized response a status code equal to that given
func (o *KubernetesDescribeServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeServiceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeServiceUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeServiceForbidden creates a KubernetesDescribeServiceForbidden with default headers values
func NewKubernetesDescribeServiceForbidden() *KubernetesDescribeServiceForbidden {
	return &KubernetesDescribeServiceForbidden{}
}

/*
KubernetesDescribeServiceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeServiceForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe service forbidden response has a 2xx status code
func (o *KubernetesDescribeServiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe service forbidden response has a 3xx status code
func (o *KubernetesDescribeServiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe service forbidden response has a 4xx status code
func (o *KubernetesDescribeServiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe service forbidden response has a 5xx status code
func (o *KubernetesDescribeServiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe service forbidden response a status code equal to that given
func (o *KubernetesDescribeServiceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeServiceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeServiceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeServiceForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeServiceNotFound creates a KubernetesDescribeServiceNotFound with default headers values
func NewKubernetesDescribeServiceNotFound() *KubernetesDescribeServiceNotFound {
	return &KubernetesDescribeServiceNotFound{}
}

/*
KubernetesDescribeServiceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeServiceNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe service not found response has a 2xx status code
func (o *KubernetesDescribeServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe service not found response has a 3xx status code
func (o *KubernetesDescribeServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe service not found response has a 4xx status code
func (o *KubernetesDescribeServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe service not found response has a 5xx status code
func (o *KubernetesDescribeServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe service not found response a status code equal to that given
func (o *KubernetesDescribeServiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeServiceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeServiceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeServiceNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeServiceInternalServerError creates a KubernetesDescribeServiceInternalServerError with default headers values
func NewKubernetesDescribeServiceInternalServerError() *KubernetesDescribeServiceInternalServerError {
	return &KubernetesDescribeServiceInternalServerError{}
}

/*
KubernetesDescribeServiceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeServiceInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe service internal server error response has a 2xx status code
func (o *KubernetesDescribeServiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe service internal server error response has a 3xx status code
func (o *KubernetesDescribeServiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe service internal server error response has a 4xx status code
func (o *KubernetesDescribeServiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe service internal server error response has a 5xx status code
func (o *KubernetesDescribeServiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe service internal server error response a status code equal to that given
func (o *KubernetesDescribeServiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeServiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceInternalServerError ", 500)
}

func (o *KubernetesDescribeServiceInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/service][%d] kubernetesDescribeServiceInternalServerError ", 500)
}

func (o *KubernetesDescribeServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
