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

// KubernetesPatchConfigMapReader is a Reader for the KubernetesPatchConfigMap structure.
type KubernetesPatchConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchConfigMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchConfigMapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchConfigMapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchConfigMapOK creates a KubernetesPatchConfigMapOK with default headers values
func NewKubernetesPatchConfigMapOK() *KubernetesPatchConfigMapOK {
	return &KubernetesPatchConfigMapOK{}
}

/*
KubernetesPatchConfigMapOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchConfigMapOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch config map o k response has a 2xx status code
func (o *KubernetesPatchConfigMapOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch config map o k response has a 3xx status code
func (o *KubernetesPatchConfigMapOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch config map o k response has a 4xx status code
func (o *KubernetesPatchConfigMapOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch config map o k response has a 5xx status code
func (o *KubernetesPatchConfigMapOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch config map o k response a status code equal to that given
func (o *KubernetesPatchConfigMapOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchConfigMapOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchConfigMapOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchConfigMapOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchConfigMapBadRequest creates a KubernetesPatchConfigMapBadRequest with default headers values
func NewKubernetesPatchConfigMapBadRequest() *KubernetesPatchConfigMapBadRequest {
	return &KubernetesPatchConfigMapBadRequest{}
}

/*
KubernetesPatchConfigMapBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchConfigMapBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch config map bad request response has a 2xx status code
func (o *KubernetesPatchConfigMapBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch config map bad request response has a 3xx status code
func (o *KubernetesPatchConfigMapBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch config map bad request response has a 4xx status code
func (o *KubernetesPatchConfigMapBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch config map bad request response has a 5xx status code
func (o *KubernetesPatchConfigMapBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch config map bad request response a status code equal to that given
func (o *KubernetesPatchConfigMapBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchConfigMapBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchConfigMapBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchConfigMapBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchConfigMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchConfigMapUnauthorized creates a KubernetesPatchConfigMapUnauthorized with default headers values
func NewKubernetesPatchConfigMapUnauthorized() *KubernetesPatchConfigMapUnauthorized {
	return &KubernetesPatchConfigMapUnauthorized{}
}

/*
KubernetesPatchConfigMapUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchConfigMapUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch config map unauthorized response has a 2xx status code
func (o *KubernetesPatchConfigMapUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch config map unauthorized response has a 3xx status code
func (o *KubernetesPatchConfigMapUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch config map unauthorized response has a 4xx status code
func (o *KubernetesPatchConfigMapUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch config map unauthorized response has a 5xx status code
func (o *KubernetesPatchConfigMapUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch config map unauthorized response a status code equal to that given
func (o *KubernetesPatchConfigMapUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchConfigMapUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchConfigMapUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchConfigMapUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchConfigMapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchConfigMapForbidden creates a KubernetesPatchConfigMapForbidden with default headers values
func NewKubernetesPatchConfigMapForbidden() *KubernetesPatchConfigMapForbidden {
	return &KubernetesPatchConfigMapForbidden{}
}

/*
KubernetesPatchConfigMapForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchConfigMapForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch config map forbidden response has a 2xx status code
func (o *KubernetesPatchConfigMapForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch config map forbidden response has a 3xx status code
func (o *KubernetesPatchConfigMapForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch config map forbidden response has a 4xx status code
func (o *KubernetesPatchConfigMapForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch config map forbidden response has a 5xx status code
func (o *KubernetesPatchConfigMapForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch config map forbidden response a status code equal to that given
func (o *KubernetesPatchConfigMapForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchConfigMapForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchConfigMapForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchConfigMapForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchConfigMapNotFound creates a KubernetesPatchConfigMapNotFound with default headers values
func NewKubernetesPatchConfigMapNotFound() *KubernetesPatchConfigMapNotFound {
	return &KubernetesPatchConfigMapNotFound{}
}

/*
KubernetesPatchConfigMapNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchConfigMapNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch config map not found response has a 2xx status code
func (o *KubernetesPatchConfigMapNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch config map not found response has a 3xx status code
func (o *KubernetesPatchConfigMapNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch config map not found response has a 4xx status code
func (o *KubernetesPatchConfigMapNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch config map not found response has a 5xx status code
func (o *KubernetesPatchConfigMapNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch config map not found response a status code equal to that given
func (o *KubernetesPatchConfigMapNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchConfigMapNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchConfigMapNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchConfigMapNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchConfigMapInternalServerError creates a KubernetesPatchConfigMapInternalServerError with default headers values
func NewKubernetesPatchConfigMapInternalServerError() *KubernetesPatchConfigMapInternalServerError {
	return &KubernetesPatchConfigMapInternalServerError{}
}

/*
KubernetesPatchConfigMapInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchConfigMapInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch config map internal server error response has a 2xx status code
func (o *KubernetesPatchConfigMapInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch config map internal server error response has a 3xx status code
func (o *KubernetesPatchConfigMapInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch config map internal server error response has a 4xx status code
func (o *KubernetesPatchConfigMapInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch config map internal server error response has a 5xx status code
func (o *KubernetesPatchConfigMapInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch config map internal server error response a status code equal to that given
func (o *KubernetesPatchConfigMapInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchConfigMapInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapInternalServerError ", 500)
}

func (o *KubernetesPatchConfigMapInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/configmap][%d] kubernetesPatchConfigMapInternalServerError ", 500)
}

func (o *KubernetesPatchConfigMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
