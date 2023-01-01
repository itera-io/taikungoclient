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

// KubernetesDescribeNetworkPolicyReader is a Reader for the KubernetesDescribeNetworkPolicy structure.
type KubernetesDescribeNetworkPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeNetworkPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeNetworkPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeNetworkPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeNetworkPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeNetworkPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeNetworkPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeNetworkPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeNetworkPolicyOK creates a KubernetesDescribeNetworkPolicyOK with default headers values
func NewKubernetesDescribeNetworkPolicyOK() *KubernetesDescribeNetworkPolicyOK {
	return &KubernetesDescribeNetworkPolicyOK{}
}

/*
KubernetesDescribeNetworkPolicyOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeNetworkPolicyOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe network policy o k response has a 2xx status code
func (o *KubernetesDescribeNetworkPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe network policy o k response has a 3xx status code
func (o *KubernetesDescribeNetworkPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe network policy o k response has a 4xx status code
func (o *KubernetesDescribeNetworkPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe network policy o k response has a 5xx status code
func (o *KubernetesDescribeNetworkPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe network policy o k response a status code equal to that given
func (o *KubernetesDescribeNetworkPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeNetworkPolicyOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeNetworkPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNetworkPolicyBadRequest creates a KubernetesDescribeNetworkPolicyBadRequest with default headers values
func NewKubernetesDescribeNetworkPolicyBadRequest() *KubernetesDescribeNetworkPolicyBadRequest {
	return &KubernetesDescribeNetworkPolicyBadRequest{}
}

/*
KubernetesDescribeNetworkPolicyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeNetworkPolicyBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe network policy bad request response has a 2xx status code
func (o *KubernetesDescribeNetworkPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe network policy bad request response has a 3xx status code
func (o *KubernetesDescribeNetworkPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe network policy bad request response has a 4xx status code
func (o *KubernetesDescribeNetworkPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe network policy bad request response has a 5xx status code
func (o *KubernetesDescribeNetworkPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe network policy bad request response a status code equal to that given
func (o *KubernetesDescribeNetworkPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeNetworkPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeNetworkPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNetworkPolicyUnauthorized creates a KubernetesDescribeNetworkPolicyUnauthorized with default headers values
func NewKubernetesDescribeNetworkPolicyUnauthorized() *KubernetesDescribeNetworkPolicyUnauthorized {
	return &KubernetesDescribeNetworkPolicyUnauthorized{}
}

/*
KubernetesDescribeNetworkPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeNetworkPolicyUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe network policy unauthorized response has a 2xx status code
func (o *KubernetesDescribeNetworkPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe network policy unauthorized response has a 3xx status code
func (o *KubernetesDescribeNetworkPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe network policy unauthorized response has a 4xx status code
func (o *KubernetesDescribeNetworkPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe network policy unauthorized response has a 5xx status code
func (o *KubernetesDescribeNetworkPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe network policy unauthorized response a status code equal to that given
func (o *KubernetesDescribeNetworkPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeNetworkPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeNetworkPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNetworkPolicyForbidden creates a KubernetesDescribeNetworkPolicyForbidden with default headers values
func NewKubernetesDescribeNetworkPolicyForbidden() *KubernetesDescribeNetworkPolicyForbidden {
	return &KubernetesDescribeNetworkPolicyForbidden{}
}

/*
KubernetesDescribeNetworkPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeNetworkPolicyForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe network policy forbidden response has a 2xx status code
func (o *KubernetesDescribeNetworkPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe network policy forbidden response has a 3xx status code
func (o *KubernetesDescribeNetworkPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe network policy forbidden response has a 4xx status code
func (o *KubernetesDescribeNetworkPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe network policy forbidden response has a 5xx status code
func (o *KubernetesDescribeNetworkPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe network policy forbidden response a status code equal to that given
func (o *KubernetesDescribeNetworkPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeNetworkPolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeNetworkPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNetworkPolicyNotFound creates a KubernetesDescribeNetworkPolicyNotFound with default headers values
func NewKubernetesDescribeNetworkPolicyNotFound() *KubernetesDescribeNetworkPolicyNotFound {
	return &KubernetesDescribeNetworkPolicyNotFound{}
}

/*
KubernetesDescribeNetworkPolicyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeNetworkPolicyNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe network policy not found response has a 2xx status code
func (o *KubernetesDescribeNetworkPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe network policy not found response has a 3xx status code
func (o *KubernetesDescribeNetworkPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe network policy not found response has a 4xx status code
func (o *KubernetesDescribeNetworkPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe network policy not found response has a 5xx status code
func (o *KubernetesDescribeNetworkPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe network policy not found response a status code equal to that given
func (o *KubernetesDescribeNetworkPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeNetworkPolicyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeNetworkPolicyNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeNetworkPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeNetworkPolicyInternalServerError creates a KubernetesDescribeNetworkPolicyInternalServerError with default headers values
func NewKubernetesDescribeNetworkPolicyInternalServerError() *KubernetesDescribeNetworkPolicyInternalServerError {
	return &KubernetesDescribeNetworkPolicyInternalServerError{}
}

/*
KubernetesDescribeNetworkPolicyInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeNetworkPolicyInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe network policy internal server error response has a 2xx status code
func (o *KubernetesDescribeNetworkPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe network policy internal server error response has a 3xx status code
func (o *KubernetesDescribeNetworkPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe network policy internal server error response has a 4xx status code
func (o *KubernetesDescribeNetworkPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe network policy internal server error response has a 5xx status code
func (o *KubernetesDescribeNetworkPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe network policy internal server error response a status code equal to that given
func (o *KubernetesDescribeNetworkPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeNetworkPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyInternalServerError ", 500)
}

func (o *KubernetesDescribeNetworkPolicyInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/network-policy][%d] kubernetesDescribeNetworkPolicyInternalServerError ", 500)
}

func (o *KubernetesDescribeNetworkPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
