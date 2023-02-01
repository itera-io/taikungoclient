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

/*
KubernetesDescribeStorageClassOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeStorageClassOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe storage class o k response has a 2xx status code
func (o *KubernetesDescribeStorageClassOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe storage class o k response has a 3xx status code
func (o *KubernetesDescribeStorageClassOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe storage class o k response has a 4xx status code
func (o *KubernetesDescribeStorageClassOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe storage class o k response has a 5xx status code
func (o *KubernetesDescribeStorageClassOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe storage class o k response a status code equal to that given
func (o *KubernetesDescribeStorageClassOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes describe storage class o k response
func (o *KubernetesDescribeStorageClassOK) Code() int {
	return 200
}

func (o *KubernetesDescribeStorageClassOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeStorageClassOK) String() string {
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

/*
KubernetesDescribeStorageClassBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeStorageClassBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe storage class bad request response has a 2xx status code
func (o *KubernetesDescribeStorageClassBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe storage class bad request response has a 3xx status code
func (o *KubernetesDescribeStorageClassBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe storage class bad request response has a 4xx status code
func (o *KubernetesDescribeStorageClassBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe storage class bad request response has a 5xx status code
func (o *KubernetesDescribeStorageClassBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe storage class bad request response a status code equal to that given
func (o *KubernetesDescribeStorageClassBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes describe storage class bad request response
func (o *KubernetesDescribeStorageClassBadRequest) Code() int {
	return 400
}

func (o *KubernetesDescribeStorageClassBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeStorageClassBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeStorageClassBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStorageClassBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

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

/*
KubernetesDescribeStorageClassUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeStorageClassUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe storage class unauthorized response has a 2xx status code
func (o *KubernetesDescribeStorageClassUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe storage class unauthorized response has a 3xx status code
func (o *KubernetesDescribeStorageClassUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe storage class unauthorized response has a 4xx status code
func (o *KubernetesDescribeStorageClassUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe storage class unauthorized response has a 5xx status code
func (o *KubernetesDescribeStorageClassUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe storage class unauthorized response a status code equal to that given
func (o *KubernetesDescribeStorageClassUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes describe storage class unauthorized response
func (o *KubernetesDescribeStorageClassUnauthorized) Code() int {
	return 401
}

func (o *KubernetesDescribeStorageClassUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeStorageClassUnauthorized) String() string {
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

/*
KubernetesDescribeStorageClassForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeStorageClassForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe storage class forbidden response has a 2xx status code
func (o *KubernetesDescribeStorageClassForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe storage class forbidden response has a 3xx status code
func (o *KubernetesDescribeStorageClassForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe storage class forbidden response has a 4xx status code
func (o *KubernetesDescribeStorageClassForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe storage class forbidden response has a 5xx status code
func (o *KubernetesDescribeStorageClassForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe storage class forbidden response a status code equal to that given
func (o *KubernetesDescribeStorageClassForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes describe storage class forbidden response
func (o *KubernetesDescribeStorageClassForbidden) Code() int {
	return 403
}

func (o *KubernetesDescribeStorageClassForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeStorageClassForbidden) String() string {
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

/*
KubernetesDescribeStorageClassNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeStorageClassNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe storage class not found response has a 2xx status code
func (o *KubernetesDescribeStorageClassNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe storage class not found response has a 3xx status code
func (o *KubernetesDescribeStorageClassNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe storage class not found response has a 4xx status code
func (o *KubernetesDescribeStorageClassNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe storage class not found response has a 5xx status code
func (o *KubernetesDescribeStorageClassNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe storage class not found response a status code equal to that given
func (o *KubernetesDescribeStorageClassNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes describe storage class not found response
func (o *KubernetesDescribeStorageClassNotFound) Code() int {
	return 404
}

func (o *KubernetesDescribeStorageClassNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeStorageClassNotFound) String() string {
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

/*
KubernetesDescribeStorageClassInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeStorageClassInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe storage class internal server error response has a 2xx status code
func (o *KubernetesDescribeStorageClassInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe storage class internal server error response has a 3xx status code
func (o *KubernetesDescribeStorageClassInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe storage class internal server error response has a 4xx status code
func (o *KubernetesDescribeStorageClassInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe storage class internal server error response has a 5xx status code
func (o *KubernetesDescribeStorageClassInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe storage class internal server error response a status code equal to that given
func (o *KubernetesDescribeStorageClassInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes describe storage class internal server error response
func (o *KubernetesDescribeStorageClassInternalServerError) Code() int {
	return 500
}

func (o *KubernetesDescribeStorageClassInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassInternalServerError ", 500)
}

func (o *KubernetesDescribeStorageClassInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/storageclass][%d] kubernetesDescribeStorageClassInternalServerError ", 500)
}

func (o *KubernetesDescribeStorageClassInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
