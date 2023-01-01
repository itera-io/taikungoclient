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

// KubernetesGetNamespacesListReader is a Reader for the KubernetesGetNamespacesList structure.
type KubernetesGetNamespacesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetNamespacesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetNamespacesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetNamespacesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetNamespacesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetNamespacesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetNamespacesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetNamespacesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetNamespacesListOK creates a KubernetesGetNamespacesListOK with default headers values
func NewKubernetesGetNamespacesListOK() *KubernetesGetNamespacesListOK {
	return &KubernetesGetNamespacesListOK{}
}

/*
KubernetesGetNamespacesListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetNamespacesListOK struct {
	Payload []string
}

// IsSuccess returns true when this kubernetes get namespaces list o k response has a 2xx status code
func (o *KubernetesGetNamespacesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get namespaces list o k response has a 3xx status code
func (o *KubernetesGetNamespacesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get namespaces list o k response has a 4xx status code
func (o *KubernetesGetNamespacesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get namespaces list o k response has a 5xx status code
func (o *KubernetesGetNamespacesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get namespaces list o k response a status code equal to that given
func (o *KubernetesGetNamespacesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetNamespacesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetNamespacesListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetNamespacesListOK) GetPayload() []string {
	return o.Payload
}

func (o *KubernetesGetNamespacesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNamespacesListBadRequest creates a KubernetesGetNamespacesListBadRequest with default headers values
func NewKubernetesGetNamespacesListBadRequest() *KubernetesGetNamespacesListBadRequest {
	return &KubernetesGetNamespacesListBadRequest{}
}

/*
KubernetesGetNamespacesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetNamespacesListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get namespaces list bad request response has a 2xx status code
func (o *KubernetesGetNamespacesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get namespaces list bad request response has a 3xx status code
func (o *KubernetesGetNamespacesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get namespaces list bad request response has a 4xx status code
func (o *KubernetesGetNamespacesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get namespaces list bad request response has a 5xx status code
func (o *KubernetesGetNamespacesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get namespaces list bad request response a status code equal to that given
func (o *KubernetesGetNamespacesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetNamespacesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetNamespacesListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetNamespacesListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetNamespacesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNamespacesListUnauthorized creates a KubernetesGetNamespacesListUnauthorized with default headers values
func NewKubernetesGetNamespacesListUnauthorized() *KubernetesGetNamespacesListUnauthorized {
	return &KubernetesGetNamespacesListUnauthorized{}
}

/*
KubernetesGetNamespacesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetNamespacesListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get namespaces list unauthorized response has a 2xx status code
func (o *KubernetesGetNamespacesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get namespaces list unauthorized response has a 3xx status code
func (o *KubernetesGetNamespacesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get namespaces list unauthorized response has a 4xx status code
func (o *KubernetesGetNamespacesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get namespaces list unauthorized response has a 5xx status code
func (o *KubernetesGetNamespacesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get namespaces list unauthorized response a status code equal to that given
func (o *KubernetesGetNamespacesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetNamespacesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetNamespacesListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetNamespacesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetNamespacesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNamespacesListForbidden creates a KubernetesGetNamespacesListForbidden with default headers values
func NewKubernetesGetNamespacesListForbidden() *KubernetesGetNamespacesListForbidden {
	return &KubernetesGetNamespacesListForbidden{}
}

/*
KubernetesGetNamespacesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetNamespacesListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get namespaces list forbidden response has a 2xx status code
func (o *KubernetesGetNamespacesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get namespaces list forbidden response has a 3xx status code
func (o *KubernetesGetNamespacesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get namespaces list forbidden response has a 4xx status code
func (o *KubernetesGetNamespacesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get namespaces list forbidden response has a 5xx status code
func (o *KubernetesGetNamespacesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get namespaces list forbidden response a status code equal to that given
func (o *KubernetesGetNamespacesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetNamespacesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetNamespacesListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetNamespacesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetNamespacesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNamespacesListNotFound creates a KubernetesGetNamespacesListNotFound with default headers values
func NewKubernetesGetNamespacesListNotFound() *KubernetesGetNamespacesListNotFound {
	return &KubernetesGetNamespacesListNotFound{}
}

/*
KubernetesGetNamespacesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetNamespacesListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get namespaces list not found response has a 2xx status code
func (o *KubernetesGetNamespacesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get namespaces list not found response has a 3xx status code
func (o *KubernetesGetNamespacesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get namespaces list not found response has a 4xx status code
func (o *KubernetesGetNamespacesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get namespaces list not found response has a 5xx status code
func (o *KubernetesGetNamespacesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get namespaces list not found response a status code equal to that given
func (o *KubernetesGetNamespacesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetNamespacesListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetNamespacesListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetNamespacesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetNamespacesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetNamespacesListInternalServerError creates a KubernetesGetNamespacesListInternalServerError with default headers values
func NewKubernetesGetNamespacesListInternalServerError() *KubernetesGetNamespacesListInternalServerError {
	return &KubernetesGetNamespacesListInternalServerError{}
}

/*
KubernetesGetNamespacesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetNamespacesListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get namespaces list internal server error response has a 2xx status code
func (o *KubernetesGetNamespacesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get namespaces list internal server error response has a 3xx status code
func (o *KubernetesGetNamespacesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get namespaces list internal server error response has a 4xx status code
func (o *KubernetesGetNamespacesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get namespaces list internal server error response has a 5xx status code
func (o *KubernetesGetNamespacesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get namespaces list internal server error response a status code equal to that given
func (o *KubernetesGetNamespacesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetNamespacesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListInternalServerError ", 500)
}

func (o *KubernetesGetNamespacesListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/namespaces][%d] kubernetesGetNamespacesListInternalServerError ", 500)
}

func (o *KubernetesGetNamespacesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
