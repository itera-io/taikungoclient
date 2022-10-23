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

// KubernetesGetSecretReader is a Reader for the KubernetesGetSecret structure.
type KubernetesGetSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetSecretOK creates a KubernetesGetSecretOK with default headers values
func NewKubernetesGetSecretOK() *KubernetesGetSecretOK {
	return &KubernetesGetSecretOK{}
}

/*
KubernetesGetSecretOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetSecretOK struct {
	Payload *models.Secrets
}

// IsSuccess returns true when this kubernetes get secret o k response has a 2xx status code
func (o *KubernetesGetSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get secret o k response has a 3xx status code
func (o *KubernetesGetSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get secret o k response has a 4xx status code
func (o *KubernetesGetSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get secret o k response has a 5xx status code
func (o *KubernetesGetSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get secret o k response a status code equal to that given
func (o *KubernetesGetSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetSecretOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetSecretOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetSecretOK) GetPayload() *models.Secrets {
	return o.Payload
}

func (o *KubernetesGetSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSecretBadRequest creates a KubernetesGetSecretBadRequest with default headers values
func NewKubernetesGetSecretBadRequest() *KubernetesGetSecretBadRequest {
	return &KubernetesGetSecretBadRequest{}
}

/*
KubernetesGetSecretBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetSecretBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes get secret bad request response has a 2xx status code
func (o *KubernetesGetSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get secret bad request response has a 3xx status code
func (o *KubernetesGetSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get secret bad request response has a 4xx status code
func (o *KubernetesGetSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get secret bad request response has a 5xx status code
func (o *KubernetesGetSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get secret bad request response a status code equal to that given
func (o *KubernetesGetSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetSecretBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetSecretBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetSecretBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesGetSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSecretUnauthorized creates a KubernetesGetSecretUnauthorized with default headers values
func NewKubernetesGetSecretUnauthorized() *KubernetesGetSecretUnauthorized {
	return &KubernetesGetSecretUnauthorized{}
}

/*
KubernetesGetSecretUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetSecretUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get secret unauthorized response has a 2xx status code
func (o *KubernetesGetSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get secret unauthorized response has a 3xx status code
func (o *KubernetesGetSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get secret unauthorized response has a 4xx status code
func (o *KubernetesGetSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get secret unauthorized response has a 5xx status code
func (o *KubernetesGetSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get secret unauthorized response a status code equal to that given
func (o *KubernetesGetSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetSecretUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetSecretUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetSecretUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSecretForbidden creates a KubernetesGetSecretForbidden with default headers values
func NewKubernetesGetSecretForbidden() *KubernetesGetSecretForbidden {
	return &KubernetesGetSecretForbidden{}
}

/*
KubernetesGetSecretForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetSecretForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get secret forbidden response has a 2xx status code
func (o *KubernetesGetSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get secret forbidden response has a 3xx status code
func (o *KubernetesGetSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get secret forbidden response has a 4xx status code
func (o *KubernetesGetSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get secret forbidden response has a 5xx status code
func (o *KubernetesGetSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get secret forbidden response a status code equal to that given
func (o *KubernetesGetSecretForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetSecretForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetSecretForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetSecretForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSecretNotFound creates a KubernetesGetSecretNotFound with default headers values
func NewKubernetesGetSecretNotFound() *KubernetesGetSecretNotFound {
	return &KubernetesGetSecretNotFound{}
}

/*
KubernetesGetSecretNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetSecretNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get secret not found response has a 2xx status code
func (o *KubernetesGetSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get secret not found response has a 3xx status code
func (o *KubernetesGetSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get secret not found response has a 4xx status code
func (o *KubernetesGetSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get secret not found response has a 5xx status code
func (o *KubernetesGetSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get secret not found response a status code equal to that given
func (o *KubernetesGetSecretNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetSecretNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetSecretNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetSecretNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetSecretInternalServerError creates a KubernetesGetSecretInternalServerError with default headers values
func NewKubernetesGetSecretInternalServerError() *KubernetesGetSecretInternalServerError {
	return &KubernetesGetSecretInternalServerError{}
}

/*
KubernetesGetSecretInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetSecretInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get secret internal server error response has a 2xx status code
func (o *KubernetesGetSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get secret internal server error response has a 3xx status code
func (o *KubernetesGetSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get secret internal server error response has a 4xx status code
func (o *KubernetesGetSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get secret internal server error response has a 5xx status code
func (o *KubernetesGetSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get secret internal server error response a status code equal to that given
func (o *KubernetesGetSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetSecretInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretInternalServerError ", 500)
}

func (o *KubernetesGetSecretInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/secret][%d] kubernetesGetSecretInternalServerError ", 500)
}

func (o *KubernetesGetSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
