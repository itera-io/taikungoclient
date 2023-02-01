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

// KubernetesDescribeDaemonSetReader is a Reader for the KubernetesDescribeDaemonSet structure.
type KubernetesDescribeDaemonSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeDaemonSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeDaemonSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeDaemonSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeDaemonSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeDaemonSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeDaemonSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeDaemonSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeDaemonSetOK creates a KubernetesDescribeDaemonSetOK with default headers values
func NewKubernetesDescribeDaemonSetOK() *KubernetesDescribeDaemonSetOK {
	return &KubernetesDescribeDaemonSetOK{}
}

/*
KubernetesDescribeDaemonSetOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeDaemonSetOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe daemon set o k response has a 2xx status code
func (o *KubernetesDescribeDaemonSetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe daemon set o k response has a 3xx status code
func (o *KubernetesDescribeDaemonSetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe daemon set o k response has a 4xx status code
func (o *KubernetesDescribeDaemonSetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe daemon set o k response has a 5xx status code
func (o *KubernetesDescribeDaemonSetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe daemon set o k response a status code equal to that given
func (o *KubernetesDescribeDaemonSetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes describe daemon set o k response
func (o *KubernetesDescribeDaemonSetOK) Code() int {
	return 200
}

func (o *KubernetesDescribeDaemonSetOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeDaemonSetOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeDaemonSetOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeDaemonSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDaemonSetBadRequest creates a KubernetesDescribeDaemonSetBadRequest with default headers values
func NewKubernetesDescribeDaemonSetBadRequest() *KubernetesDescribeDaemonSetBadRequest {
	return &KubernetesDescribeDaemonSetBadRequest{}
}

/*
KubernetesDescribeDaemonSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeDaemonSetBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe daemon set bad request response has a 2xx status code
func (o *KubernetesDescribeDaemonSetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe daemon set bad request response has a 3xx status code
func (o *KubernetesDescribeDaemonSetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe daemon set bad request response has a 4xx status code
func (o *KubernetesDescribeDaemonSetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe daemon set bad request response has a 5xx status code
func (o *KubernetesDescribeDaemonSetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe daemon set bad request response a status code equal to that given
func (o *KubernetesDescribeDaemonSetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes describe daemon set bad request response
func (o *KubernetesDescribeDaemonSetBadRequest) Code() int {
	return 400
}

func (o *KubernetesDescribeDaemonSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeDaemonSetBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeDaemonSetBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDaemonSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDaemonSetUnauthorized creates a KubernetesDescribeDaemonSetUnauthorized with default headers values
func NewKubernetesDescribeDaemonSetUnauthorized() *KubernetesDescribeDaemonSetUnauthorized {
	return &KubernetesDescribeDaemonSetUnauthorized{}
}

/*
KubernetesDescribeDaemonSetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeDaemonSetUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe daemon set unauthorized response has a 2xx status code
func (o *KubernetesDescribeDaemonSetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe daemon set unauthorized response has a 3xx status code
func (o *KubernetesDescribeDaemonSetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe daemon set unauthorized response has a 4xx status code
func (o *KubernetesDescribeDaemonSetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe daemon set unauthorized response has a 5xx status code
func (o *KubernetesDescribeDaemonSetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe daemon set unauthorized response a status code equal to that given
func (o *KubernetesDescribeDaemonSetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes describe daemon set unauthorized response
func (o *KubernetesDescribeDaemonSetUnauthorized) Code() int {
	return 401
}

func (o *KubernetesDescribeDaemonSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeDaemonSetUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeDaemonSetUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDaemonSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDaemonSetForbidden creates a KubernetesDescribeDaemonSetForbidden with default headers values
func NewKubernetesDescribeDaemonSetForbidden() *KubernetesDescribeDaemonSetForbidden {
	return &KubernetesDescribeDaemonSetForbidden{}
}

/*
KubernetesDescribeDaemonSetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeDaemonSetForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe daemon set forbidden response has a 2xx status code
func (o *KubernetesDescribeDaemonSetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe daemon set forbidden response has a 3xx status code
func (o *KubernetesDescribeDaemonSetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe daemon set forbidden response has a 4xx status code
func (o *KubernetesDescribeDaemonSetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe daemon set forbidden response has a 5xx status code
func (o *KubernetesDescribeDaemonSetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe daemon set forbidden response a status code equal to that given
func (o *KubernetesDescribeDaemonSetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes describe daemon set forbidden response
func (o *KubernetesDescribeDaemonSetForbidden) Code() int {
	return 403
}

func (o *KubernetesDescribeDaemonSetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeDaemonSetForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeDaemonSetForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDaemonSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDaemonSetNotFound creates a KubernetesDescribeDaemonSetNotFound with default headers values
func NewKubernetesDescribeDaemonSetNotFound() *KubernetesDescribeDaemonSetNotFound {
	return &KubernetesDescribeDaemonSetNotFound{}
}

/*
KubernetesDescribeDaemonSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeDaemonSetNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe daemon set not found response has a 2xx status code
func (o *KubernetesDescribeDaemonSetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe daemon set not found response has a 3xx status code
func (o *KubernetesDescribeDaemonSetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe daemon set not found response has a 4xx status code
func (o *KubernetesDescribeDaemonSetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe daemon set not found response has a 5xx status code
func (o *KubernetesDescribeDaemonSetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe daemon set not found response a status code equal to that given
func (o *KubernetesDescribeDaemonSetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes describe daemon set not found response
func (o *KubernetesDescribeDaemonSetNotFound) Code() int {
	return 404
}

func (o *KubernetesDescribeDaemonSetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeDaemonSetNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeDaemonSetNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeDaemonSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDaemonSetInternalServerError creates a KubernetesDescribeDaemonSetInternalServerError with default headers values
func NewKubernetesDescribeDaemonSetInternalServerError() *KubernetesDescribeDaemonSetInternalServerError {
	return &KubernetesDescribeDaemonSetInternalServerError{}
}

/*
KubernetesDescribeDaemonSetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeDaemonSetInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe daemon set internal server error response has a 2xx status code
func (o *KubernetesDescribeDaemonSetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe daemon set internal server error response has a 3xx status code
func (o *KubernetesDescribeDaemonSetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe daemon set internal server error response has a 4xx status code
func (o *KubernetesDescribeDaemonSetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe daemon set internal server error response has a 5xx status code
func (o *KubernetesDescribeDaemonSetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe daemon set internal server error response a status code equal to that given
func (o *KubernetesDescribeDaemonSetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes describe daemon set internal server error response
func (o *KubernetesDescribeDaemonSetInternalServerError) Code() int {
	return 500
}

func (o *KubernetesDescribeDaemonSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetInternalServerError ", 500)
}

func (o *KubernetesDescribeDaemonSetInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/daemonset][%d] kubernetesDescribeDaemonSetInternalServerError ", 500)
}

func (o *KubernetesDescribeDaemonSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
