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

// KubernetesDescribeSecretReader is a Reader for the KubernetesDescribeSecret structure.
type KubernetesDescribeSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribeSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribeSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribeSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribeSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribeSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribeSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribeSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribeSecretOK creates a KubernetesDescribeSecretOK with default headers values
func NewKubernetesDescribeSecretOK() *KubernetesDescribeSecretOK {
	return &KubernetesDescribeSecretOK{}
}

/*
KubernetesDescribeSecretOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribeSecretOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe secret o k response has a 2xx status code
func (o *KubernetesDescribeSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe secret o k response has a 3xx status code
func (o *KubernetesDescribeSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe secret o k response has a 4xx status code
func (o *KubernetesDescribeSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe secret o k response has a 5xx status code
func (o *KubernetesDescribeSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe secret o k response a status code equal to that given
func (o *KubernetesDescribeSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribeSecretOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeSecretOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribeSecretOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribeSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeSecretBadRequest creates a KubernetesDescribeSecretBadRequest with default headers values
func NewKubernetesDescribeSecretBadRequest() *KubernetesDescribeSecretBadRequest {
	return &KubernetesDescribeSecretBadRequest{}
}

/*
KubernetesDescribeSecretBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribeSecretBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe secret bad request response has a 2xx status code
func (o *KubernetesDescribeSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe secret bad request response has a 3xx status code
func (o *KubernetesDescribeSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe secret bad request response has a 4xx status code
func (o *KubernetesDescribeSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe secret bad request response has a 5xx status code
func (o *KubernetesDescribeSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe secret bad request response a status code equal to that given
func (o *KubernetesDescribeSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribeSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeSecretBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribeSecretBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeSecretUnauthorized creates a KubernetesDescribeSecretUnauthorized with default headers values
func NewKubernetesDescribeSecretUnauthorized() *KubernetesDescribeSecretUnauthorized {
	return &KubernetesDescribeSecretUnauthorized{}
}

/*
KubernetesDescribeSecretUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribeSecretUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe secret unauthorized response has a 2xx status code
func (o *KubernetesDescribeSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe secret unauthorized response has a 3xx status code
func (o *KubernetesDescribeSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe secret unauthorized response has a 4xx status code
func (o *KubernetesDescribeSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe secret unauthorized response has a 5xx status code
func (o *KubernetesDescribeSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe secret unauthorized response a status code equal to that given
func (o *KubernetesDescribeSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribeSecretUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeSecretUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribeSecretUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeSecretForbidden creates a KubernetesDescribeSecretForbidden with default headers values
func NewKubernetesDescribeSecretForbidden() *KubernetesDescribeSecretForbidden {
	return &KubernetesDescribeSecretForbidden{}
}

/*
KubernetesDescribeSecretForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribeSecretForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe secret forbidden response has a 2xx status code
func (o *KubernetesDescribeSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe secret forbidden response has a 3xx status code
func (o *KubernetesDescribeSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe secret forbidden response has a 4xx status code
func (o *KubernetesDescribeSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe secret forbidden response has a 5xx status code
func (o *KubernetesDescribeSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe secret forbidden response a status code equal to that given
func (o *KubernetesDescribeSecretForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribeSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeSecretForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribeSecretForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeSecretNotFound creates a KubernetesDescribeSecretNotFound with default headers values
func NewKubernetesDescribeSecretNotFound() *KubernetesDescribeSecretNotFound {
	return &KubernetesDescribeSecretNotFound{}
}

/*
KubernetesDescribeSecretNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribeSecretNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe secret not found response has a 2xx status code
func (o *KubernetesDescribeSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe secret not found response has a 3xx status code
func (o *KubernetesDescribeSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe secret not found response has a 4xx status code
func (o *KubernetesDescribeSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe secret not found response has a 5xx status code
func (o *KubernetesDescribeSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe secret not found response a status code equal to that given
func (o *KubernetesDescribeSecretNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribeSecretNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeSecretNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribeSecretNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribeSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribeSecretInternalServerError creates a KubernetesDescribeSecretInternalServerError with default headers values
func NewKubernetesDescribeSecretInternalServerError() *KubernetesDescribeSecretInternalServerError {
	return &KubernetesDescribeSecretInternalServerError{}
}

/*
KubernetesDescribeSecretInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribeSecretInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe secret internal server error response has a 2xx status code
func (o *KubernetesDescribeSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe secret internal server error response has a 3xx status code
func (o *KubernetesDescribeSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe secret internal server error response has a 4xx status code
func (o *KubernetesDescribeSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe secret internal server error response has a 5xx status code
func (o *KubernetesDescribeSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe secret internal server error response a status code equal to that given
func (o *KubernetesDescribeSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribeSecretInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretInternalServerError ", 500)
}

func (o *KubernetesDescribeSecretInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/secret][%d] kubernetesDescribeSecretInternalServerError ", 500)
}

func (o *KubernetesDescribeSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
