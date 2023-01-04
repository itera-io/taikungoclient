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

/*
KubernetesDescribeConfigMapOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeConfigMapOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe config map o k response has a 2xx status code
func (o *KubernetesDescribeConfigMapOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe config map o k response has a 3xx status code
func (o *KubernetesDescribeConfigMapOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe config map o k response has a 4xx status code
func (o *KubernetesDescribeConfigMapOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe config map o k response has a 5xx status code
func (o *KubernetesDescribeConfigMapOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe config map o k response a status code equal to that given
func (o *KubernetesDescribeConfigMapOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeConfigMapOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeConfigMapOK) String() string {
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

/*
KubernetesDescribeConfigMapBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeConfigMapBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe config map bad request response has a 2xx status code
func (o *KubernetesDescribeConfigMapBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe config map bad request response has a 3xx status code
func (o *KubernetesDescribeConfigMapBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe config map bad request response has a 4xx status code
func (o *KubernetesDescribeConfigMapBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe config map bad request response has a 5xx status code
func (o *KubernetesDescribeConfigMapBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe config map bad request response a status code equal to that given
func (o *KubernetesDescribeConfigMapBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeConfigMapBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeConfigMapBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeConfigMapBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeConfigMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeConfigMapUnauthorized creates a KubernetesDescribeConfigMapUnauthorized with default headers values
func NewKubernetesDescribeConfigMapUnauthorized() *KubernetesDescribeConfigMapUnauthorized {
	return &KubernetesDescribeConfigMapUnauthorized{}
}

/*
KubernetesDescribeConfigMapUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeConfigMapUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe config map unauthorized response has a 2xx status code
func (o *KubernetesDescribeConfigMapUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe config map unauthorized response has a 3xx status code
func (o *KubernetesDescribeConfigMapUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe config map unauthorized response has a 4xx status code
func (o *KubernetesDescribeConfigMapUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe config map unauthorized response has a 5xx status code
func (o *KubernetesDescribeConfigMapUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe config map unauthorized response a status code equal to that given
func (o *KubernetesDescribeConfigMapUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeConfigMapUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeConfigMapUnauthorized) String() string {
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

/*
KubernetesDescribeConfigMapForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeConfigMapForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe config map forbidden response has a 2xx status code
func (o *KubernetesDescribeConfigMapForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe config map forbidden response has a 3xx status code
func (o *KubernetesDescribeConfigMapForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe config map forbidden response has a 4xx status code
func (o *KubernetesDescribeConfigMapForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe config map forbidden response has a 5xx status code
func (o *KubernetesDescribeConfigMapForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe config map forbidden response a status code equal to that given
func (o *KubernetesDescribeConfigMapForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeConfigMapForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeConfigMapForbidden) String() string {
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

/*
KubernetesDescribeConfigMapNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeConfigMapNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe config map not found response has a 2xx status code
func (o *KubernetesDescribeConfigMapNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe config map not found response has a 3xx status code
func (o *KubernetesDescribeConfigMapNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe config map not found response has a 4xx status code
func (o *KubernetesDescribeConfigMapNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe config map not found response has a 5xx status code
func (o *KubernetesDescribeConfigMapNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe config map not found response a status code equal to that given
func (o *KubernetesDescribeConfigMapNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeConfigMapNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeConfigMapNotFound) String() string {
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

/*
KubernetesDescribeConfigMapInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeConfigMapInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe config map internal server error response has a 2xx status code
func (o *KubernetesDescribeConfigMapInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe config map internal server error response has a 3xx status code
func (o *KubernetesDescribeConfigMapInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe config map internal server error response has a 4xx status code
func (o *KubernetesDescribeConfigMapInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe config map internal server error response has a 5xx status code
func (o *KubernetesDescribeConfigMapInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe config map internal server error response a status code equal to that given
func (o *KubernetesDescribeConfigMapInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeConfigMapInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapInternalServerError ", 500)
}

func (o *KubernetesDescribeConfigMapInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/configmap][%d] kubernetesDescribeConfigMapInternalServerError ", 500)
}

func (o *KubernetesDescribeConfigMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
