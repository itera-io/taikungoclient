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

// KubernetesGetDeploymentListReader is a Reader for the KubernetesGetDeploymentList structure.
type KubernetesGetDeploymentListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetDeploymentListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetDeploymentListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetDeploymentListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetDeploymentListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetDeploymentListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetDeploymentListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetDeploymentListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetDeploymentListOK creates a KubernetesGetDeploymentListOK with default headers values
func NewKubernetesGetDeploymentListOK() *KubernetesGetDeploymentListOK {
	return &KubernetesGetDeploymentListOK{}
}

/*
KubernetesGetDeploymentListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetDeploymentListOK struct {
	Payload *models.Deployments
}

// IsSuccess returns true when this kubernetes get deployment list o k response has a 2xx status code
func (o *KubernetesGetDeploymentListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get deployment list o k response has a 3xx status code
func (o *KubernetesGetDeploymentListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get deployment list o k response has a 4xx status code
func (o *KubernetesGetDeploymentListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get deployment list o k response has a 5xx status code
func (o *KubernetesGetDeploymentListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get deployment list o k response a status code equal to that given
func (o *KubernetesGetDeploymentListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetDeploymentListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetDeploymentListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetDeploymentListOK) GetPayload() *models.Deployments {
	return o.Payload
}

func (o *KubernetesGetDeploymentListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDeploymentListBadRequest creates a KubernetesGetDeploymentListBadRequest with default headers values
func NewKubernetesGetDeploymentListBadRequest() *KubernetesGetDeploymentListBadRequest {
	return &KubernetesGetDeploymentListBadRequest{}
}

/*
KubernetesGetDeploymentListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetDeploymentListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes get deployment list bad request response has a 2xx status code
func (o *KubernetesGetDeploymentListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get deployment list bad request response has a 3xx status code
func (o *KubernetesGetDeploymentListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get deployment list bad request response has a 4xx status code
func (o *KubernetesGetDeploymentListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get deployment list bad request response has a 5xx status code
func (o *KubernetesGetDeploymentListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get deployment list bad request response a status code equal to that given
func (o *KubernetesGetDeploymentListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetDeploymentListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetDeploymentListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetDeploymentListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesGetDeploymentListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDeploymentListUnauthorized creates a KubernetesGetDeploymentListUnauthorized with default headers values
func NewKubernetesGetDeploymentListUnauthorized() *KubernetesGetDeploymentListUnauthorized {
	return &KubernetesGetDeploymentListUnauthorized{}
}

/*
KubernetesGetDeploymentListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetDeploymentListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get deployment list unauthorized response has a 2xx status code
func (o *KubernetesGetDeploymentListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get deployment list unauthorized response has a 3xx status code
func (o *KubernetesGetDeploymentListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get deployment list unauthorized response has a 4xx status code
func (o *KubernetesGetDeploymentListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get deployment list unauthorized response has a 5xx status code
func (o *KubernetesGetDeploymentListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get deployment list unauthorized response a status code equal to that given
func (o *KubernetesGetDeploymentListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetDeploymentListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetDeploymentListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetDeploymentListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDeploymentListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDeploymentListForbidden creates a KubernetesGetDeploymentListForbidden with default headers values
func NewKubernetesGetDeploymentListForbidden() *KubernetesGetDeploymentListForbidden {
	return &KubernetesGetDeploymentListForbidden{}
}

/*
KubernetesGetDeploymentListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetDeploymentListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get deployment list forbidden response has a 2xx status code
func (o *KubernetesGetDeploymentListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get deployment list forbidden response has a 3xx status code
func (o *KubernetesGetDeploymentListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get deployment list forbidden response has a 4xx status code
func (o *KubernetesGetDeploymentListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get deployment list forbidden response has a 5xx status code
func (o *KubernetesGetDeploymentListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get deployment list forbidden response a status code equal to that given
func (o *KubernetesGetDeploymentListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetDeploymentListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetDeploymentListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetDeploymentListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDeploymentListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDeploymentListNotFound creates a KubernetesGetDeploymentListNotFound with default headers values
func NewKubernetesGetDeploymentListNotFound() *KubernetesGetDeploymentListNotFound {
	return &KubernetesGetDeploymentListNotFound{}
}

/*
KubernetesGetDeploymentListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetDeploymentListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get deployment list not found response has a 2xx status code
func (o *KubernetesGetDeploymentListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get deployment list not found response has a 3xx status code
func (o *KubernetesGetDeploymentListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get deployment list not found response has a 4xx status code
func (o *KubernetesGetDeploymentListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get deployment list not found response has a 5xx status code
func (o *KubernetesGetDeploymentListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get deployment list not found response a status code equal to that given
func (o *KubernetesGetDeploymentListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetDeploymentListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetDeploymentListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetDeploymentListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDeploymentListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDeploymentListInternalServerError creates a KubernetesGetDeploymentListInternalServerError with default headers values
func NewKubernetesGetDeploymentListInternalServerError() *KubernetesGetDeploymentListInternalServerError {
	return &KubernetesGetDeploymentListInternalServerError{}
}

/*
KubernetesGetDeploymentListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetDeploymentListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get deployment list internal server error response has a 2xx status code
func (o *KubernetesGetDeploymentListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get deployment list internal server error response has a 3xx status code
func (o *KubernetesGetDeploymentListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get deployment list internal server error response has a 4xx status code
func (o *KubernetesGetDeploymentListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get deployment list internal server error response has a 5xx status code
func (o *KubernetesGetDeploymentListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get deployment list internal server error response a status code equal to that given
func (o *KubernetesGetDeploymentListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetDeploymentListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListInternalServerError ", 500)
}

func (o *KubernetesGetDeploymentListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/deployment][%d] kubernetesGetDeploymentListInternalServerError ", 500)
}

func (o *KubernetesGetDeploymentListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
