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

// KubernetesDescribeDeploymentReader is a Reader for the KubernetesDescribeDeployment structure.
type KubernetesDescribeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeDeploymentOK creates a KubernetesDescribeDeploymentOK with default headers values
func NewKubernetesDescribeDeploymentOK() *KubernetesDescribeDeploymentOK {
	return &KubernetesDescribeDeploymentOK{}
}

/*
KubernetesDescribeDeploymentOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeDeploymentOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe deployment o k response has a 2xx status code
func (o *KubernetesDescribeDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe deployment o k response has a 3xx status code
func (o *KubernetesDescribeDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe deployment o k response has a 4xx status code
func (o *KubernetesDescribeDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe deployment o k response has a 5xx status code
func (o *KubernetesDescribeDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe deployment o k response a status code equal to that given
func (o *KubernetesDescribeDeploymentOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeDeploymentOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeDeploymentOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentBadRequest creates a KubernetesDescribeDeploymentBadRequest with default headers values
func NewKubernetesDescribeDeploymentBadRequest() *KubernetesDescribeDeploymentBadRequest {
	return &KubernetesDescribeDeploymentBadRequest{}
}

/*
KubernetesDescribeDeploymentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeDeploymentBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe deployment bad request response has a 2xx status code
func (o *KubernetesDescribeDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe deployment bad request response has a 3xx status code
func (o *KubernetesDescribeDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe deployment bad request response has a 4xx status code
func (o *KubernetesDescribeDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe deployment bad request response has a 5xx status code
func (o *KubernetesDescribeDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe deployment bad request response a status code equal to that given
func (o *KubernetesDescribeDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeDeploymentBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeDeploymentBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentUnauthorized creates a KubernetesDescribeDeploymentUnauthorized with default headers values
func NewKubernetesDescribeDeploymentUnauthorized() *KubernetesDescribeDeploymentUnauthorized {
	return &KubernetesDescribeDeploymentUnauthorized{}
}

/*
KubernetesDescribeDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeDeploymentUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe deployment unauthorized response has a 2xx status code
func (o *KubernetesDescribeDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe deployment unauthorized response has a 3xx status code
func (o *KubernetesDescribeDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe deployment unauthorized response has a 4xx status code
func (o *KubernetesDescribeDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe deployment unauthorized response has a 5xx status code
func (o *KubernetesDescribeDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe deployment unauthorized response a status code equal to that given
func (o *KubernetesDescribeDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeDeploymentUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentForbidden creates a KubernetesDescribeDeploymentForbidden with default headers values
func NewKubernetesDescribeDeploymentForbidden() *KubernetesDescribeDeploymentForbidden {
	return &KubernetesDescribeDeploymentForbidden{}
}

/*
KubernetesDescribeDeploymentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeDeploymentForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe deployment forbidden response has a 2xx status code
func (o *KubernetesDescribeDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe deployment forbidden response has a 3xx status code
func (o *KubernetesDescribeDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe deployment forbidden response has a 4xx status code
func (o *KubernetesDescribeDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe deployment forbidden response has a 5xx status code
func (o *KubernetesDescribeDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe deployment forbidden response a status code equal to that given
func (o *KubernetesDescribeDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeDeploymentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeDeploymentForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentNotFound creates a KubernetesDescribeDeploymentNotFound with default headers values
func NewKubernetesDescribeDeploymentNotFound() *KubernetesDescribeDeploymentNotFound {
	return &KubernetesDescribeDeploymentNotFound{}
}

/*
KubernetesDescribeDeploymentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeDeploymentNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes describe deployment not found response has a 2xx status code
func (o *KubernetesDescribeDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe deployment not found response has a 3xx status code
func (o *KubernetesDescribeDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe deployment not found response has a 4xx status code
func (o *KubernetesDescribeDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe deployment not found response has a 5xx status code
func (o *KubernetesDescribeDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe deployment not found response a status code equal to that given
func (o *KubernetesDescribeDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeDeploymentNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeDeploymentNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesDescribeDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeDeploymentInternalServerError creates a KubernetesDescribeDeploymentInternalServerError with default headers values
func NewKubernetesDescribeDeploymentInternalServerError() *KubernetesDescribeDeploymentInternalServerError {
	return &KubernetesDescribeDeploymentInternalServerError{}
}

/*
KubernetesDescribeDeploymentInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeDeploymentInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe deployment internal server error response has a 2xx status code
func (o *KubernetesDescribeDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe deployment internal server error response has a 3xx status code
func (o *KubernetesDescribeDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe deployment internal server error response has a 4xx status code
func (o *KubernetesDescribeDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe deployment internal server error response has a 5xx status code
func (o *KubernetesDescribeDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe deployment internal server error response a status code equal to that given
func (o *KubernetesDescribeDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentInternalServerError ", 500)
}

func (o *KubernetesDescribeDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/deployment][%d] kubernetesDescribeDeploymentInternalServerError ", 500)
}

func (o *KubernetesDescribeDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
