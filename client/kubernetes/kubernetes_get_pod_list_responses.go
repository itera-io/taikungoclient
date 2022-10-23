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

// KubernetesGetPodListReader is a Reader for the KubernetesGetPodList structure.
type KubernetesGetPodListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetPodListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetPodListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetPodListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetPodListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetPodListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetPodListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetPodListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetPodListOK creates a KubernetesGetPodListOK with default headers values
func NewKubernetesGetPodListOK() *KubernetesGetPodListOK {
	return &KubernetesGetPodListOK{}
}

/*
KubernetesGetPodListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetPodListOK struct {
	Payload *models.Pods
}

// IsSuccess returns true when this kubernetes get pod list o k response has a 2xx status code
func (o *KubernetesGetPodListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get pod list o k response has a 3xx status code
func (o *KubernetesGetPodListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod list o k response has a 4xx status code
func (o *KubernetesGetPodListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get pod list o k response has a 5xx status code
func (o *KubernetesGetPodListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod list o k response a status code equal to that given
func (o *KubernetesGetPodListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetPodListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetPodListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetPodListOK) GetPayload() *models.Pods {
	return o.Payload
}

func (o *KubernetesGetPodListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pods)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodListBadRequest creates a KubernetesGetPodListBadRequest with default headers values
func NewKubernetesGetPodListBadRequest() *KubernetesGetPodListBadRequest {
	return &KubernetesGetPodListBadRequest{}
}

/*
KubernetesGetPodListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetPodListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes get pod list bad request response has a 2xx status code
func (o *KubernetesGetPodListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod list bad request response has a 3xx status code
func (o *KubernetesGetPodListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod list bad request response has a 4xx status code
func (o *KubernetesGetPodListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod list bad request response has a 5xx status code
func (o *KubernetesGetPodListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod list bad request response a status code equal to that given
func (o *KubernetesGetPodListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetPodListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetPodListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetPodListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesGetPodListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodListUnauthorized creates a KubernetesGetPodListUnauthorized with default headers values
func NewKubernetesGetPodListUnauthorized() *KubernetesGetPodListUnauthorized {
	return &KubernetesGetPodListUnauthorized{}
}

/*
KubernetesGetPodListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetPodListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pod list unauthorized response has a 2xx status code
func (o *KubernetesGetPodListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod list unauthorized response has a 3xx status code
func (o *KubernetesGetPodListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod list unauthorized response has a 4xx status code
func (o *KubernetesGetPodListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod list unauthorized response has a 5xx status code
func (o *KubernetesGetPodListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod list unauthorized response a status code equal to that given
func (o *KubernetesGetPodListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetPodListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetPodListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetPodListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPodListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodListForbidden creates a KubernetesGetPodListForbidden with default headers values
func NewKubernetesGetPodListForbidden() *KubernetesGetPodListForbidden {
	return &KubernetesGetPodListForbidden{}
}

/*
KubernetesGetPodListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetPodListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pod list forbidden response has a 2xx status code
func (o *KubernetesGetPodListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod list forbidden response has a 3xx status code
func (o *KubernetesGetPodListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod list forbidden response has a 4xx status code
func (o *KubernetesGetPodListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod list forbidden response has a 5xx status code
func (o *KubernetesGetPodListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod list forbidden response a status code equal to that given
func (o *KubernetesGetPodListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetPodListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetPodListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetPodListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPodListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodListNotFound creates a KubernetesGetPodListNotFound with default headers values
func NewKubernetesGetPodListNotFound() *KubernetesGetPodListNotFound {
	return &KubernetesGetPodListNotFound{}
}

/*
KubernetesGetPodListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetPodListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get pod list not found response has a 2xx status code
func (o *KubernetesGetPodListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod list not found response has a 3xx status code
func (o *KubernetesGetPodListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod list not found response has a 4xx status code
func (o *KubernetesGetPodListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod list not found response has a 5xx status code
func (o *KubernetesGetPodListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod list not found response a status code equal to that given
func (o *KubernetesGetPodListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetPodListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetPodListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetPodListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPodListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodListInternalServerError creates a KubernetesGetPodListInternalServerError with default headers values
func NewKubernetesGetPodListInternalServerError() *KubernetesGetPodListInternalServerError {
	return &KubernetesGetPodListInternalServerError{}
}

/*
KubernetesGetPodListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetPodListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get pod list internal server error response has a 2xx status code
func (o *KubernetesGetPodListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod list internal server error response has a 3xx status code
func (o *KubernetesGetPodListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod list internal server error response has a 4xx status code
func (o *KubernetesGetPodListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get pod list internal server error response has a 5xx status code
func (o *KubernetesGetPodListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get pod list internal server error response a status code equal to that given
func (o *KubernetesGetPodListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetPodListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListInternalServerError ", 500)
}

func (o *KubernetesGetPodListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/pod][%d] kubernetesGetPodListInternalServerError ", 500)
}

func (o *KubernetesGetPodListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
