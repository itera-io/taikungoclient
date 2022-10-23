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

// KubernetesDescribePvcReader is a Reader for the KubernetesDescribePvc structure.
type KubernetesDescribePvcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesDescribePvcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesDescribePvcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesDescribePvcBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesDescribePvcUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesDescribePvcForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesDescribePvcNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesDescribePvcInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesDescribePvcOK creates a KubernetesDescribePvcOK with default headers values
func NewKubernetesDescribePvcOK() *KubernetesDescribePvcOK {
	return &KubernetesDescribePvcOK{}
}

/*
KubernetesDescribePvcOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesDescribePvcOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes describe pvc o k response has a 2xx status code
func (o *KubernetesDescribePvcOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes describe pvc o k response has a 3xx status code
func (o *KubernetesDescribePvcOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pvc o k response has a 4xx status code
func (o *KubernetesDescribePvcOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe pvc o k response has a 5xx status code
func (o *KubernetesDescribePvcOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pvc o k response a status code equal to that given
func (o *KubernetesDescribePvcOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesDescribePvcOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribePvcOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcOK  %+v", 200, o.Payload)
}

func (o *KubernetesDescribePvcOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesDescribePvcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePvcBadRequest creates a KubernetesDescribePvcBadRequest with default headers values
func NewKubernetesDescribePvcBadRequest() *KubernetesDescribePvcBadRequest {
	return &KubernetesDescribePvcBadRequest{}
}

/*
KubernetesDescribePvcBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesDescribePvcBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes describe pvc bad request response has a 2xx status code
func (o *KubernetesDescribePvcBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pvc bad request response has a 3xx status code
func (o *KubernetesDescribePvcBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pvc bad request response has a 4xx status code
func (o *KubernetesDescribePvcBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pvc bad request response has a 5xx status code
func (o *KubernetesDescribePvcBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pvc bad request response a status code equal to that given
func (o *KubernetesDescribePvcBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesDescribePvcBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribePvcBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesDescribePvcBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesDescribePvcBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePvcUnauthorized creates a KubernetesDescribePvcUnauthorized with default headers values
func NewKubernetesDescribePvcUnauthorized() *KubernetesDescribePvcUnauthorized {
	return &KubernetesDescribePvcUnauthorized{}
}

/*
KubernetesDescribePvcUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesDescribePvcUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe pvc unauthorized response has a 2xx status code
func (o *KubernetesDescribePvcUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pvc unauthorized response has a 3xx status code
func (o *KubernetesDescribePvcUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pvc unauthorized response has a 4xx status code
func (o *KubernetesDescribePvcUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pvc unauthorized response has a 5xx status code
func (o *KubernetesDescribePvcUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pvc unauthorized response a status code equal to that given
func (o *KubernetesDescribePvcUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesDescribePvcUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribePvcUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesDescribePvcUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribePvcUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePvcForbidden creates a KubernetesDescribePvcForbidden with default headers values
func NewKubernetesDescribePvcForbidden() *KubernetesDescribePvcForbidden {
	return &KubernetesDescribePvcForbidden{}
}

/*
KubernetesDescribePvcForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesDescribePvcForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe pvc forbidden response has a 2xx status code
func (o *KubernetesDescribePvcForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pvc forbidden response has a 3xx status code
func (o *KubernetesDescribePvcForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pvc forbidden response has a 4xx status code
func (o *KubernetesDescribePvcForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pvc forbidden response has a 5xx status code
func (o *KubernetesDescribePvcForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pvc forbidden response a status code equal to that given
func (o *KubernetesDescribePvcForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesDescribePvcForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribePvcForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesDescribePvcForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribePvcForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePvcNotFound creates a KubernetesDescribePvcNotFound with default headers values
func NewKubernetesDescribePvcNotFound() *KubernetesDescribePvcNotFound {
	return &KubernetesDescribePvcNotFound{}
}

/*
KubernetesDescribePvcNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesDescribePvcNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes describe pvc not found response has a 2xx status code
func (o *KubernetesDescribePvcNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pvc not found response has a 3xx status code
func (o *KubernetesDescribePvcNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pvc not found response has a 4xx status code
func (o *KubernetesDescribePvcNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes describe pvc not found response has a 5xx status code
func (o *KubernetesDescribePvcNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes describe pvc not found response a status code equal to that given
func (o *KubernetesDescribePvcNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesDescribePvcNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribePvcNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesDescribePvcNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesDescribePvcNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesDescribePvcInternalServerError creates a KubernetesDescribePvcInternalServerError with default headers values
func NewKubernetesDescribePvcInternalServerError() *KubernetesDescribePvcInternalServerError {
	return &KubernetesDescribePvcInternalServerError{}
}

/*
KubernetesDescribePvcInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesDescribePvcInternalServerError struct {
}

// IsSuccess returns true when this kubernetes describe pvc internal server error response has a 2xx status code
func (o *KubernetesDescribePvcInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes describe pvc internal server error response has a 3xx status code
func (o *KubernetesDescribePvcInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes describe pvc internal server error response has a 4xx status code
func (o *KubernetesDescribePvcInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes describe pvc internal server error response has a 5xx status code
func (o *KubernetesDescribePvcInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes describe pvc internal server error response a status code equal to that given
func (o *KubernetesDescribePvcInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesDescribePvcInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcInternalServerError ", 500)
}

func (o *KubernetesDescribePvcInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/describe/pvc][%d] kubernetesDescribePvcInternalServerError ", 500)
}

func (o *KubernetesDescribePvcInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
