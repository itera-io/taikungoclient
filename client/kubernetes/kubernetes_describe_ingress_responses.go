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

// KubernetesDescribeIngressReader is a Reader for the KubernetesDescribeIngress structure.
type KubernetesDescribeIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeIngressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeIngressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeIngressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeIngressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeIngressOK creates a KubernetesDescribeIngressOK with default headers values
func NewKubernetesDescribeIngressOK() *KubernetesDescribeIngressOK {
	return &KubernetesDescribeIngressOK{}
}

/*
KubernetesDescribeIngressOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeIngressOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe ingress o k response has a 2xx status code
func (o *KubernetesDescribeIngressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe ingress o k response has a 3xx status code
func (o *KubernetesDescribeIngressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe ingress o k response has a 4xx status code
func (o *KubernetesDescribeIngressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe ingress o k response has a 5xx status code
func (o *KubernetesDescribeIngressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe ingress o k response a status code equal to that given
func (o *KubernetesDescribeIngressOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeIngressOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeIngressOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeIngressOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeIngressBadRequest creates a KubernetesDescribeIngressBadRequest with default headers values
func NewKubernetesDescribeIngressBadRequest() *KubernetesDescribeIngressBadRequest {
	return &KubernetesDescribeIngressBadRequest{}
}

/*
KubernetesDescribeIngressBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeIngressBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes describe ingress bad request response has a 2xx status code
func (o *KubernetesDescribeIngressBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe ingress bad request response has a 3xx status code
func (o *KubernetesDescribeIngressBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe ingress bad request response has a 4xx status code
func (o *KubernetesDescribeIngressBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe ingress bad request response has a 5xx status code
func (o *KubernetesDescribeIngressBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe ingress bad request response a status code equal to that given
func (o *KubernetesDescribeIngressBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeIngressBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeIngressBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeIngressBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeIngressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeIngressUnauthorized creates a KubernetesDescribeIngressUnauthorized with default headers values
func NewKubernetesDescribeIngressUnauthorized() *KubernetesDescribeIngressUnauthorized {
	return &KubernetesDescribeIngressUnauthorized{}
}

/*
KubernetesDescribeIngressUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeIngressUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe ingress unauthorized response has a 2xx status code
func (o *KubernetesDescribeIngressUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe ingress unauthorized response has a 3xx status code
func (o *KubernetesDescribeIngressUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe ingress unauthorized response has a 4xx status code
func (o *KubernetesDescribeIngressUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe ingress unauthorized response has a 5xx status code
func (o *KubernetesDescribeIngressUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe ingress unauthorized response a status code equal to that given
func (o *KubernetesDescribeIngressUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeIngressUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeIngressUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeIngressUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeIngressForbidden creates a KubernetesDescribeIngressForbidden with default headers values
func NewKubernetesDescribeIngressForbidden() *KubernetesDescribeIngressForbidden {
	return &KubernetesDescribeIngressForbidden{}
}

/*
KubernetesDescribeIngressForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeIngressForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe ingress forbidden response has a 2xx status code
func (o *KubernetesDescribeIngressForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe ingress forbidden response has a 3xx status code
func (o *KubernetesDescribeIngressForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe ingress forbidden response has a 4xx status code
func (o *KubernetesDescribeIngressForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe ingress forbidden response has a 5xx status code
func (o *KubernetesDescribeIngressForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe ingress forbidden response a status code equal to that given
func (o *KubernetesDescribeIngressForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeIngressForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeIngressForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeIngressForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeIngressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeIngressNotFound creates a KubernetesDescribeIngressNotFound with default headers values
func NewKubernetesDescribeIngressNotFound() *KubernetesDescribeIngressNotFound {
	return &KubernetesDescribeIngressNotFound{}
}

/*
KubernetesDescribeIngressNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeIngressNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes describe ingress not found response has a 2xx status code
func (o *KubernetesDescribeIngressNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe ingress not found response has a 3xx status code
func (o *KubernetesDescribeIngressNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe ingress not found response has a 4xx status code
func (o *KubernetesDescribeIngressNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe ingress not found response has a 5xx status code
func (o *KubernetesDescribeIngressNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe ingress not found response a status code equal to that given
func (o *KubernetesDescribeIngressNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeIngressNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeIngressNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeIngressNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesDescribeIngressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeIngressInternalServerError creates a KubernetesDescribeIngressInternalServerError with default headers values
func NewKubernetesDescribeIngressInternalServerError() *KubernetesDescribeIngressInternalServerError {
	return &KubernetesDescribeIngressInternalServerError{}
}

/*
KubernetesDescribeIngressInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeIngressInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe ingress internal server error response has a 2xx status code
func (o *KubernetesDescribeIngressInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe ingress internal server error response has a 3xx status code
func (o *KubernetesDescribeIngressInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe ingress internal server error response has a 4xx status code
func (o *KubernetesDescribeIngressInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe ingress internal server error response has a 5xx status code
func (o *KubernetesDescribeIngressInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe ingress internal server error response a status code equal to that given
func (o *KubernetesDescribeIngressInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeIngressInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressInternalServerError ", 500)
}

func (o *KubernetesDescribeIngressInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/ingress][%d] kubernetesDescribeIngressInternalServerError ", 500)
}

func (o *KubernetesDescribeIngressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
