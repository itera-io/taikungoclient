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

// KubernetesGetConfigMapReader is a Reader for the KubernetesGetConfigMap structure.
type KubernetesGetConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetConfigMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetConfigMapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetConfigMapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetConfigMapOK creates a KubernetesGetConfigMapOK with default headers values
func NewKubernetesGetConfigMapOK() *KubernetesGetConfigMapOK {
	return &KubernetesGetConfigMapOK{}
}

/*
KubernetesGetConfigMapOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetConfigMapOK struct {
	Payload *models.ConfigMaps
}

// IsSuccess returns true when this kubernetes get config map o k response has a 2xx status code
func (o *KubernetesGetConfigMapOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get config map o k response has a 3xx status code
func (o *KubernetesGetConfigMapOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get config map o k response has a 4xx status code
func (o *KubernetesGetConfigMapOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get config map o k response has a 5xx status code
func (o *KubernetesGetConfigMapOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get config map o k response a status code equal to that given
func (o *KubernetesGetConfigMapOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetConfigMapOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetConfigMapOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetConfigMapOK) GetPayload() *models.ConfigMaps {
	return o.Payload
}

func (o *KubernetesGetConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigMaps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetConfigMapBadRequest creates a KubernetesGetConfigMapBadRequest with default headers values
func NewKubernetesGetConfigMapBadRequest() *KubernetesGetConfigMapBadRequest {
	return &KubernetesGetConfigMapBadRequest{}
}

/*
KubernetesGetConfigMapBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetConfigMapBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes get config map bad request response has a 2xx status code
func (o *KubernetesGetConfigMapBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get config map bad request response has a 3xx status code
func (o *KubernetesGetConfigMapBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get config map bad request response has a 4xx status code
func (o *KubernetesGetConfigMapBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get config map bad request response has a 5xx status code
func (o *KubernetesGetConfigMapBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get config map bad request response a status code equal to that given
func (o *KubernetesGetConfigMapBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetConfigMapBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetConfigMapBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetConfigMapBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesGetConfigMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetConfigMapUnauthorized creates a KubernetesGetConfigMapUnauthorized with default headers values
func NewKubernetesGetConfigMapUnauthorized() *KubernetesGetConfigMapUnauthorized {
	return &KubernetesGetConfigMapUnauthorized{}
}

/*
KubernetesGetConfigMapUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetConfigMapUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get config map unauthorized response has a 2xx status code
func (o *KubernetesGetConfigMapUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get config map unauthorized response has a 3xx status code
func (o *KubernetesGetConfigMapUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get config map unauthorized response has a 4xx status code
func (o *KubernetesGetConfigMapUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get config map unauthorized response has a 5xx status code
func (o *KubernetesGetConfigMapUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get config map unauthorized response a status code equal to that given
func (o *KubernetesGetConfigMapUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetConfigMapUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetConfigMapUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetConfigMapUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetConfigMapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetConfigMapForbidden creates a KubernetesGetConfigMapForbidden with default headers values
func NewKubernetesGetConfigMapForbidden() *KubernetesGetConfigMapForbidden {
	return &KubernetesGetConfigMapForbidden{}
}

/*
KubernetesGetConfigMapForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetConfigMapForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get config map forbidden response has a 2xx status code
func (o *KubernetesGetConfigMapForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get config map forbidden response has a 3xx status code
func (o *KubernetesGetConfigMapForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get config map forbidden response has a 4xx status code
func (o *KubernetesGetConfigMapForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get config map forbidden response has a 5xx status code
func (o *KubernetesGetConfigMapForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get config map forbidden response a status code equal to that given
func (o *KubernetesGetConfigMapForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetConfigMapForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetConfigMapForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetConfigMapForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetConfigMapNotFound creates a KubernetesGetConfigMapNotFound with default headers values
func NewKubernetesGetConfigMapNotFound() *KubernetesGetConfigMapNotFound {
	return &KubernetesGetConfigMapNotFound{}
}

/*
KubernetesGetConfigMapNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetConfigMapNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get config map not found response has a 2xx status code
func (o *KubernetesGetConfigMapNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get config map not found response has a 3xx status code
func (o *KubernetesGetConfigMapNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get config map not found response has a 4xx status code
func (o *KubernetesGetConfigMapNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get config map not found response has a 5xx status code
func (o *KubernetesGetConfigMapNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get config map not found response a status code equal to that given
func (o *KubernetesGetConfigMapNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetConfigMapNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetConfigMapNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetConfigMapNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetConfigMapInternalServerError creates a KubernetesGetConfigMapInternalServerError with default headers values
func NewKubernetesGetConfigMapInternalServerError() *KubernetesGetConfigMapInternalServerError {
	return &KubernetesGetConfigMapInternalServerError{}
}

/*
KubernetesGetConfigMapInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetConfigMapInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get config map internal server error response has a 2xx status code
func (o *KubernetesGetConfigMapInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get config map internal server error response has a 3xx status code
func (o *KubernetesGetConfigMapInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get config map internal server error response has a 4xx status code
func (o *KubernetesGetConfigMapInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get config map internal server error response has a 5xx status code
func (o *KubernetesGetConfigMapInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get config map internal server error response a status code equal to that given
func (o *KubernetesGetConfigMapInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetConfigMapInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapInternalServerError ", 500)
}

func (o *KubernetesGetConfigMapInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/configmap][%d] kubernetesGetConfigMapInternalServerError ", 500)
}

func (o *KubernetesGetConfigMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
