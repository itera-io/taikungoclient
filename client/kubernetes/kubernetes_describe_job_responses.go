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

// KubernetesDescribeJobReader is a Reader for the KubernetesDescribeJob structure.
type KubernetesDescribeJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeJobOK creates a KubernetesDescribeJobOK with default headers values
func NewKubernetesDescribeJobOK() *KubernetesDescribeJobOK {
	return &KubernetesDescribeJobOK{}
}

/*
KubernetesDescribeJobOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeJobOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe job o k response has a 2xx status code
func (o *KubernetesDescribeJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe job o k response has a 3xx status code
func (o *KubernetesDescribeJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe job o k response has a 4xx status code
func (o *KubernetesDescribeJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe job o k response has a 5xx status code
func (o *KubernetesDescribeJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe job o k response a status code equal to that given
func (o *KubernetesDescribeJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeJobOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeJobOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeJobOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeJobBadRequest creates a KubernetesDescribeJobBadRequest with default headers values
func NewKubernetesDescribeJobBadRequest() *KubernetesDescribeJobBadRequest {
	return &KubernetesDescribeJobBadRequest{}
}

/*
KubernetesDescribeJobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeJobBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes describe job bad request response has a 2xx status code
func (o *KubernetesDescribeJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe job bad request response has a 3xx status code
func (o *KubernetesDescribeJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe job bad request response has a 4xx status code
func (o *KubernetesDescribeJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe job bad request response has a 5xx status code
func (o *KubernetesDescribeJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe job bad request response a status code equal to that given
func (o *KubernetesDescribeJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeJobBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeJobBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeJobUnauthorized creates a KubernetesDescribeJobUnauthorized with default headers values
func NewKubernetesDescribeJobUnauthorized() *KubernetesDescribeJobUnauthorized {
	return &KubernetesDescribeJobUnauthorized{}
}

/*
KubernetesDescribeJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeJobUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe job unauthorized response has a 2xx status code
func (o *KubernetesDescribeJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe job unauthorized response has a 3xx status code
func (o *KubernetesDescribeJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe job unauthorized response has a 4xx status code
func (o *KubernetesDescribeJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe job unauthorized response has a 5xx status code
func (o *KubernetesDescribeJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe job unauthorized response a status code equal to that given
func (o *KubernetesDescribeJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeJobUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeJobForbidden creates a KubernetesDescribeJobForbidden with default headers values
func NewKubernetesDescribeJobForbidden() *KubernetesDescribeJobForbidden {
	return &KubernetesDescribeJobForbidden{}
}

/*
KubernetesDescribeJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeJobForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe job forbidden response has a 2xx status code
func (o *KubernetesDescribeJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe job forbidden response has a 3xx status code
func (o *KubernetesDescribeJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe job forbidden response has a 4xx status code
func (o *KubernetesDescribeJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe job forbidden response has a 5xx status code
func (o *KubernetesDescribeJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe job forbidden response a status code equal to that given
func (o *KubernetesDescribeJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeJobForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeJobForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeJobNotFound creates a KubernetesDescribeJobNotFound with default headers values
func NewKubernetesDescribeJobNotFound() *KubernetesDescribeJobNotFound {
	return &KubernetesDescribeJobNotFound{}
}

/*
KubernetesDescribeJobNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeJobNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe job not found response has a 2xx status code
func (o *KubernetesDescribeJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe job not found response has a 3xx status code
func (o *KubernetesDescribeJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe job not found response has a 4xx status code
func (o *KubernetesDescribeJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe job not found response has a 5xx status code
func (o *KubernetesDescribeJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe job not found response a status code equal to that given
func (o *KubernetesDescribeJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeJobNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeJobNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeJobInternalServerError creates a KubernetesDescribeJobInternalServerError with default headers values
func NewKubernetesDescribeJobInternalServerError() *KubernetesDescribeJobInternalServerError {
	return &KubernetesDescribeJobInternalServerError{}
}

/*
KubernetesDescribeJobInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeJobInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe job internal server error response has a 2xx status code
func (o *KubernetesDescribeJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe job internal server error response has a 3xx status code
func (o *KubernetesDescribeJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe job internal server error response has a 4xx status code
func (o *KubernetesDescribeJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe job internal server error response has a 5xx status code
func (o *KubernetesDescribeJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe job internal server error response a status code equal to that given
func (o *KubernetesDescribeJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobInternalServerError ", 500)
}

func (o *KubernetesDescribeJobInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/job][%d] kubernetesDescribeJobInternalServerError ", 500)
}

func (o *KubernetesDescribeJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
