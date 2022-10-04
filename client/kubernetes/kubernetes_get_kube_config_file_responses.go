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

// KubernetesGetKubeConfigFileReader is a Reader for the KubernetesGetKubeConfigFile structure.
type KubernetesGetKubeConfigFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetKubeConfigFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetKubeConfigFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetKubeConfigFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetKubeConfigFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetKubeConfigFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetKubeConfigFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetKubeConfigFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetKubeConfigFileOK creates a KubernetesGetKubeConfigFileOK with default headers values
func NewKubernetesGetKubeConfigFileOK() *KubernetesGetKubeConfigFileOK {
	return &KubernetesGetKubeConfigFileOK{}
}

/*
KubernetesGetKubeConfigFileOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetKubeConfigFileOK struct {
	Payload *models.KubeConfigResponse
}

// IsSuccess returns true when this kubernetes get kube config file o k response has a 2xx status code
func (o *KubernetesGetKubeConfigFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get kube config file o k response has a 3xx status code
func (o *KubernetesGetKubeConfigFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kube config file o k response has a 4xx status code
func (o *KubernetesGetKubeConfigFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get kube config file o k response has a 5xx status code
func (o *KubernetesGetKubeConfigFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kube config file o k response a status code equal to that given
func (o *KubernetesGetKubeConfigFileOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetKubeConfigFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetKubeConfigFileOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetKubeConfigFileOK) GetPayload() *models.KubeConfigResponse {
	return o.Payload
}

func (o *KubernetesGetKubeConfigFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubeConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubeConfigFileBadRequest creates a KubernetesGetKubeConfigFileBadRequest with default headers values
func NewKubernetesGetKubeConfigFileBadRequest() *KubernetesGetKubeConfigFileBadRequest {
	return &KubernetesGetKubeConfigFileBadRequest{}
}

/*
KubernetesGetKubeConfigFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetKubeConfigFileBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes get kube config file bad request response has a 2xx status code
func (o *KubernetesGetKubeConfigFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kube config file bad request response has a 3xx status code
func (o *KubernetesGetKubeConfigFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kube config file bad request response has a 4xx status code
func (o *KubernetesGetKubeConfigFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kube config file bad request response has a 5xx status code
func (o *KubernetesGetKubeConfigFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kube config file bad request response a status code equal to that given
func (o *KubernetesGetKubeConfigFileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetKubeConfigFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetKubeConfigFileBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetKubeConfigFileBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetKubeConfigFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubeConfigFileUnauthorized creates a KubernetesGetKubeConfigFileUnauthorized with default headers values
func NewKubernetesGetKubeConfigFileUnauthorized() *KubernetesGetKubeConfigFileUnauthorized {
	return &KubernetesGetKubeConfigFileUnauthorized{}
}

/*
KubernetesGetKubeConfigFileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetKubeConfigFileUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get kube config file unauthorized response has a 2xx status code
func (o *KubernetesGetKubeConfigFileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kube config file unauthorized response has a 3xx status code
func (o *KubernetesGetKubeConfigFileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kube config file unauthorized response has a 4xx status code
func (o *KubernetesGetKubeConfigFileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kube config file unauthorized response has a 5xx status code
func (o *KubernetesGetKubeConfigFileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kube config file unauthorized response a status code equal to that given
func (o *KubernetesGetKubeConfigFileUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetKubeConfigFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetKubeConfigFileUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetKubeConfigFileUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetKubeConfigFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubeConfigFileForbidden creates a KubernetesGetKubeConfigFileForbidden with default headers values
func NewKubernetesGetKubeConfigFileForbidden() *KubernetesGetKubeConfigFileForbidden {
	return &KubernetesGetKubeConfigFileForbidden{}
}

/*
KubernetesGetKubeConfigFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetKubeConfigFileForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get kube config file forbidden response has a 2xx status code
func (o *KubernetesGetKubeConfigFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kube config file forbidden response has a 3xx status code
func (o *KubernetesGetKubeConfigFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kube config file forbidden response has a 4xx status code
func (o *KubernetesGetKubeConfigFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kube config file forbidden response has a 5xx status code
func (o *KubernetesGetKubeConfigFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kube config file forbidden response a status code equal to that given
func (o *KubernetesGetKubeConfigFileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetKubeConfigFileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetKubeConfigFileForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetKubeConfigFileForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetKubeConfigFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubeConfigFileNotFound creates a KubernetesGetKubeConfigFileNotFound with default headers values
func NewKubernetesGetKubeConfigFileNotFound() *KubernetesGetKubeConfigFileNotFound {
	return &KubernetesGetKubeConfigFileNotFound{}
}

/*
KubernetesGetKubeConfigFileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetKubeConfigFileNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get kube config file not found response has a 2xx status code
func (o *KubernetesGetKubeConfigFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kube config file not found response has a 3xx status code
func (o *KubernetesGetKubeConfigFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kube config file not found response has a 4xx status code
func (o *KubernetesGetKubeConfigFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get kube config file not found response has a 5xx status code
func (o *KubernetesGetKubeConfigFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get kube config file not found response a status code equal to that given
func (o *KubernetesGetKubeConfigFileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetKubeConfigFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetKubeConfigFileNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetKubeConfigFileNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetKubeConfigFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetKubeConfigFileInternalServerError creates a KubernetesGetKubeConfigFileInternalServerError with default headers values
func NewKubernetesGetKubeConfigFileInternalServerError() *KubernetesGetKubeConfigFileInternalServerError {
	return &KubernetesGetKubeConfigFileInternalServerError{}
}

/*
KubernetesGetKubeConfigFileInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetKubeConfigFileInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get kube config file internal server error response has a 2xx status code
func (o *KubernetesGetKubeConfigFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get kube config file internal server error response has a 3xx status code
func (o *KubernetesGetKubeConfigFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get kube config file internal server error response has a 4xx status code
func (o *KubernetesGetKubeConfigFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get kube config file internal server error response has a 5xx status code
func (o *KubernetesGetKubeConfigFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get kube config file internal server error response a status code equal to that given
func (o *KubernetesGetKubeConfigFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetKubeConfigFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileInternalServerError ", 500)
}

func (o *KubernetesGetKubeConfigFileInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/kubeconfig][%d] kubernetesGetKubeConfigFileInternalServerError ", 500)
}

func (o *KubernetesGetKubeConfigFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
