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

// KubernetesDescribeStsReader is a Reader for the KubernetesDescribeSts structure.
type KubernetesDescribeStsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeStsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeStsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeStsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeStsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeStsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeStsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeStsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeStsOK creates a KubernetesDescribeStsOK with default headers values
func NewKubernetesDescribeStsOK() *KubernetesDescribeStsOK {
	return &KubernetesDescribeStsOK{}
}

/*
KubernetesDescribeStsOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeStsOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe sts o k response has a 2xx status code
func (o *KubernetesDescribeStsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe sts o k response has a 3xx status code
func (o *KubernetesDescribeStsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe sts o k response has a 4xx status code
func (o *KubernetesDescribeStsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe sts o k response has a 5xx status code
func (o *KubernetesDescribeStsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe sts o k response a status code equal to that given
func (o *KubernetesDescribeStsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes describe sts o k response
func (o *KubernetesDescribeStsOK) Code() int {
	return 200
}

func (o *KubernetesDescribeStsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeStsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeStsOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeStsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStsBadRequest creates a KubernetesDescribeStsBadRequest with default headers values
func NewKubernetesDescribeStsBadRequest() *KubernetesDescribeStsBadRequest {
	return &KubernetesDescribeStsBadRequest{}
}

/*
KubernetesDescribeStsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeStsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe sts bad request response has a 2xx status code
func (o *KubernetesDescribeStsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe sts bad request response has a 3xx status code
func (o *KubernetesDescribeStsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe sts bad request response has a 4xx status code
func (o *KubernetesDescribeStsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe sts bad request response has a 5xx status code
func (o *KubernetesDescribeStsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe sts bad request response a status code equal to that given
func (o *KubernetesDescribeStsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes describe sts bad request response
func (o *KubernetesDescribeStsBadRequest) Code() int {
	return 400
}

func (o *KubernetesDescribeStsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeStsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeStsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStsUnauthorized creates a KubernetesDescribeStsUnauthorized with default headers values
func NewKubernetesDescribeStsUnauthorized() *KubernetesDescribeStsUnauthorized {
	return &KubernetesDescribeStsUnauthorized{}
}

/*
KubernetesDescribeStsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeStsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe sts unauthorized response has a 2xx status code
func (o *KubernetesDescribeStsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe sts unauthorized response has a 3xx status code
func (o *KubernetesDescribeStsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe sts unauthorized response has a 4xx status code
func (o *KubernetesDescribeStsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe sts unauthorized response has a 5xx status code
func (o *KubernetesDescribeStsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe sts unauthorized response a status code equal to that given
func (o *KubernetesDescribeStsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes describe sts unauthorized response
func (o *KubernetesDescribeStsUnauthorized) Code() int {
	return 401
}

func (o *KubernetesDescribeStsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeStsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeStsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStsForbidden creates a KubernetesDescribeStsForbidden with default headers values
func NewKubernetesDescribeStsForbidden() *KubernetesDescribeStsForbidden {
	return &KubernetesDescribeStsForbidden{}
}

/*
KubernetesDescribeStsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeStsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe sts forbidden response has a 2xx status code
func (o *KubernetesDescribeStsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe sts forbidden response has a 3xx status code
func (o *KubernetesDescribeStsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe sts forbidden response has a 4xx status code
func (o *KubernetesDescribeStsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe sts forbidden response has a 5xx status code
func (o *KubernetesDescribeStsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe sts forbidden response a status code equal to that given
func (o *KubernetesDescribeStsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes describe sts forbidden response
func (o *KubernetesDescribeStsForbidden) Code() int {
	return 403
}

func (o *KubernetesDescribeStsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeStsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeStsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStsNotFound creates a KubernetesDescribeStsNotFound with default headers values
func NewKubernetesDescribeStsNotFound() *KubernetesDescribeStsNotFound {
	return &KubernetesDescribeStsNotFound{}
}

/*
KubernetesDescribeStsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeStsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe sts not found response has a 2xx status code
func (o *KubernetesDescribeStsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe sts not found response has a 3xx status code
func (o *KubernetesDescribeStsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe sts not found response has a 4xx status code
func (o *KubernetesDescribeStsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe sts not found response has a 5xx status code
func (o *KubernetesDescribeStsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe sts not found response a status code equal to that given
func (o *KubernetesDescribeStsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes describe sts not found response
func (o *KubernetesDescribeStsNotFound) Code() int {
	return 404
}

func (o *KubernetesDescribeStsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeStsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeStsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeStsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeStsInternalServerError creates a KubernetesDescribeStsInternalServerError with default headers values
func NewKubernetesDescribeStsInternalServerError() *KubernetesDescribeStsInternalServerError {
	return &KubernetesDescribeStsInternalServerError{}
}

/*
KubernetesDescribeStsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeStsInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe sts internal server error response has a 2xx status code
func (o *KubernetesDescribeStsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe sts internal server error response has a 3xx status code
func (o *KubernetesDescribeStsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe sts internal server error response has a 4xx status code
func (o *KubernetesDescribeStsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe sts internal server error response has a 5xx status code
func (o *KubernetesDescribeStsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe sts internal server error response a status code equal to that given
func (o *KubernetesDescribeStsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes describe sts internal server error response
func (o *KubernetesDescribeStsInternalServerError) Code() int {
	return 500
}

func (o *KubernetesDescribeStsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsInternalServerError ", 500)
}

func (o *KubernetesDescribeStsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/sts][%d] kubernetesDescribeStsInternalServerError ", 500)
}

func (o *KubernetesDescribeStsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
