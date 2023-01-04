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

// KubernetesDescribePodReader is a Reader for the KubernetesDescribePod structure.
type KubernetesDescribePodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribePodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribePodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribePodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribePodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribePodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribePodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribePodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribePodOK creates a KubernetesDescribePodOK with default headers values
func NewKubernetesDescribePodOK() *KubernetesDescribePodOK {
	return &KubernetesDescribePodOK{}
}

/*
KubernetesDescribePodOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribePodOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe pod o k response has a 2xx status code
func (o *KubernetesDescribePodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe pod o k response has a 3xx status code
func (o *KubernetesDescribePodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pod o k response has a 4xx status code
func (o *KubernetesDescribePodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe pod o k response has a 5xx status code
func (o *KubernetesDescribePodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pod o k response a status code equal to that given
func (o *KubernetesDescribePodOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribePodOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribePodOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribePodOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribePodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePodBadRequest creates a KubernetesDescribePodBadRequest with default headers values
func NewKubernetesDescribePodBadRequest() *KubernetesDescribePodBadRequest {
	return &KubernetesDescribePodBadRequest{}
}

/*
KubernetesDescribePodBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribePodBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe pod bad request response has a 2xx status code
func (o *KubernetesDescribePodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pod bad request response has a 3xx status code
func (o *KubernetesDescribePodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pod bad request response has a 4xx status code
func (o *KubernetesDescribePodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pod bad request response has a 5xx status code
func (o *KubernetesDescribePodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pod bad request response a status code equal to that given
func (o *KubernetesDescribePodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribePodBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribePodBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribePodBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribePodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePodUnauthorized creates a KubernetesDescribePodUnauthorized with default headers values
func NewKubernetesDescribePodUnauthorized() *KubernetesDescribePodUnauthorized {
	return &KubernetesDescribePodUnauthorized{}
}

/*
KubernetesDescribePodUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribePodUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe pod unauthorized response has a 2xx status code
func (o *KubernetesDescribePodUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pod unauthorized response has a 3xx status code
func (o *KubernetesDescribePodUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pod unauthorized response has a 4xx status code
func (o *KubernetesDescribePodUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pod unauthorized response has a 5xx status code
func (o *KubernetesDescribePodUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pod unauthorized response a status code equal to that given
func (o *KubernetesDescribePodUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribePodUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribePodUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribePodUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribePodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePodForbidden creates a KubernetesDescribePodForbidden with default headers values
func NewKubernetesDescribePodForbidden() *KubernetesDescribePodForbidden {
	return &KubernetesDescribePodForbidden{}
}

/*
KubernetesDescribePodForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribePodForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe pod forbidden response has a 2xx status code
func (o *KubernetesDescribePodForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pod forbidden response has a 3xx status code
func (o *KubernetesDescribePodForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pod forbidden response has a 4xx status code
func (o *KubernetesDescribePodForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pod forbidden response has a 5xx status code
func (o *KubernetesDescribePodForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pod forbidden response a status code equal to that given
func (o *KubernetesDescribePodForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribePodForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribePodForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribePodForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribePodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePodNotFound creates a KubernetesDescribePodNotFound with default headers values
func NewKubernetesDescribePodNotFound() *KubernetesDescribePodNotFound {
	return &KubernetesDescribePodNotFound{}
}

/*
KubernetesDescribePodNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribePodNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe pod not found response has a 2xx status code
func (o *KubernetesDescribePodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pod not found response has a 3xx status code
func (o *KubernetesDescribePodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pod not found response has a 4xx status code
func (o *KubernetesDescribePodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pod not found response has a 5xx status code
func (o *KubernetesDescribePodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pod not found response a status code equal to that given
func (o *KubernetesDescribePodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribePodNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribePodNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribePodNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribePodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePodInternalServerError creates a KubernetesDescribePodInternalServerError with default headers values
func NewKubernetesDescribePodInternalServerError() *KubernetesDescribePodInternalServerError {
	return &KubernetesDescribePodInternalServerError{}
}

/*
KubernetesDescribePodInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribePodInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe pod internal server error response has a 2xx status code
func (o *KubernetesDescribePodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pod internal server error response has a 3xx status code
func (o *KubernetesDescribePodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pod internal server error response has a 4xx status code
func (o *KubernetesDescribePodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe pod internal server error response has a 5xx status code
func (o *KubernetesDescribePodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe pod internal server error response a status code equal to that given
func (o *KubernetesDescribePodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribePodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodInternalServerError ", 500)
}

func (o *KubernetesDescribePodInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pod][%d] kubernetesDescribePodInternalServerError ", 500)
}

func (o *KubernetesDescribePodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
