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

// KubernetesGetCrdListReader is a Reader for the KubernetesGetCrdList structure.
type KubernetesGetCrdListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetCrdListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetCrdListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetCrdListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetCrdListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetCrdListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetCrdListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetCrdListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetCrdListOK creates a KubernetesGetCrdListOK with default headers values
func NewKubernetesGetCrdListOK() *KubernetesGetCrdListOK {
	return &KubernetesGetCrdListOK{}
}

/*
KubernetesGetCrdListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetCrdListOK struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get crd list o k response has a 2xx status code
func (o *KubernetesGetCrdListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get crd list o k response has a 3xx status code
func (o *KubernetesGetCrdListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get crd list o k response has a 4xx status code
func (o *KubernetesGetCrdListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get crd list o k response has a 5xx status code
func (o *KubernetesGetCrdListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get crd list o k response a status code equal to that given
func (o *KubernetesGetCrdListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetCrdListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetCrdListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetCrdListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetCrdListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCrdListBadRequest creates a KubernetesGetCrdListBadRequest with default headers values
func NewKubernetesGetCrdListBadRequest() *KubernetesGetCrdListBadRequest {
	return &KubernetesGetCrdListBadRequest{}
}

/*
KubernetesGetCrdListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetCrdListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes get crd list bad request response has a 2xx status code
func (o *KubernetesGetCrdListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get crd list bad request response has a 3xx status code
func (o *KubernetesGetCrdListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get crd list bad request response has a 4xx status code
func (o *KubernetesGetCrdListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get crd list bad request response has a 5xx status code
func (o *KubernetesGetCrdListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get crd list bad request response a status code equal to that given
func (o *KubernetesGetCrdListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetCrdListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetCrdListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetCrdListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesGetCrdListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCrdListUnauthorized creates a KubernetesGetCrdListUnauthorized with default headers values
func NewKubernetesGetCrdListUnauthorized() *KubernetesGetCrdListUnauthorized {
	return &KubernetesGetCrdListUnauthorized{}
}

/*
KubernetesGetCrdListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetCrdListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get crd list unauthorized response has a 2xx status code
func (o *KubernetesGetCrdListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get crd list unauthorized response has a 3xx status code
func (o *KubernetesGetCrdListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get crd list unauthorized response has a 4xx status code
func (o *KubernetesGetCrdListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get crd list unauthorized response has a 5xx status code
func (o *KubernetesGetCrdListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get crd list unauthorized response a status code equal to that given
func (o *KubernetesGetCrdListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetCrdListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetCrdListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetCrdListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCrdListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCrdListForbidden creates a KubernetesGetCrdListForbidden with default headers values
func NewKubernetesGetCrdListForbidden() *KubernetesGetCrdListForbidden {
	return &KubernetesGetCrdListForbidden{}
}

/*
KubernetesGetCrdListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetCrdListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get crd list forbidden response has a 2xx status code
func (o *KubernetesGetCrdListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get crd list forbidden response has a 3xx status code
func (o *KubernetesGetCrdListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get crd list forbidden response has a 4xx status code
func (o *KubernetesGetCrdListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get crd list forbidden response has a 5xx status code
func (o *KubernetesGetCrdListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get crd list forbidden response a status code equal to that given
func (o *KubernetesGetCrdListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetCrdListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetCrdListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetCrdListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCrdListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCrdListNotFound creates a KubernetesGetCrdListNotFound with default headers values
func NewKubernetesGetCrdListNotFound() *KubernetesGetCrdListNotFound {
	return &KubernetesGetCrdListNotFound{}
}

/*
KubernetesGetCrdListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetCrdListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get crd list not found response has a 2xx status code
func (o *KubernetesGetCrdListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get crd list not found response has a 3xx status code
func (o *KubernetesGetCrdListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get crd list not found response has a 4xx status code
func (o *KubernetesGetCrdListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get crd list not found response has a 5xx status code
func (o *KubernetesGetCrdListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get crd list not found response a status code equal to that given
func (o *KubernetesGetCrdListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetCrdListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetCrdListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetCrdListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetCrdListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetCrdListInternalServerError creates a KubernetesGetCrdListInternalServerError with default headers values
func NewKubernetesGetCrdListInternalServerError() *KubernetesGetCrdListInternalServerError {
	return &KubernetesGetCrdListInternalServerError{}
}

/*
KubernetesGetCrdListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetCrdListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get crd list internal server error response has a 2xx status code
func (o *KubernetesGetCrdListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get crd list internal server error response has a 3xx status code
func (o *KubernetesGetCrdListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get crd list internal server error response has a 4xx status code
func (o *KubernetesGetCrdListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get crd list internal server error response has a 5xx status code
func (o *KubernetesGetCrdListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get crd list internal server error response a status code equal to that given
func (o *KubernetesGetCrdListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetCrdListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListInternalServerError ", 500)
}

func (o *KubernetesGetCrdListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/crd][%d] kubernetesGetCrdListInternalServerError ", 500)
}

func (o *KubernetesGetCrdListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
