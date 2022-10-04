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

// KubernetesPatchCrdReader is a Reader for the KubernetesPatchCrd structure.
type KubernetesPatchCrdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchCrdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchCrdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchCrdBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchCrdUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchCrdForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchCrdNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchCrdInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchCrdOK creates a KubernetesPatchCrdOK with default headers values
func NewKubernetesPatchCrdOK() *KubernetesPatchCrdOK {
	return &KubernetesPatchCrdOK{}
}

/*
KubernetesPatchCrdOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchCrdOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch crd o k response has a 2xx status code
func (o *KubernetesPatchCrdOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch crd o k response has a 3xx status code
func (o *KubernetesPatchCrdOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd o k response has a 4xx status code
func (o *KubernetesPatchCrdOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch crd o k response has a 5xx status code
func (o *KubernetesPatchCrdOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd o k response a status code equal to that given
func (o *KubernetesPatchCrdOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchCrdOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchCrdOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchCrdOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchCrdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdBadRequest creates a KubernetesPatchCrdBadRequest with default headers values
func NewKubernetesPatchCrdBadRequest() *KubernetesPatchCrdBadRequest {
	return &KubernetesPatchCrdBadRequest{}
}

/*
KubernetesPatchCrdBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchCrdBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes patch crd bad request response has a 2xx status code
func (o *KubernetesPatchCrdBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd bad request response has a 3xx status code
func (o *KubernetesPatchCrdBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd bad request response has a 4xx status code
func (o *KubernetesPatchCrdBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd bad request response has a 5xx status code
func (o *KubernetesPatchCrdBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd bad request response a status code equal to that given
func (o *KubernetesPatchCrdBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchCrdBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchCrdBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchCrdBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchCrdBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdUnauthorized creates a KubernetesPatchCrdUnauthorized with default headers values
func NewKubernetesPatchCrdUnauthorized() *KubernetesPatchCrdUnauthorized {
	return &KubernetesPatchCrdUnauthorized{}
}

/*
KubernetesPatchCrdUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchCrdUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch crd unauthorized response has a 2xx status code
func (o *KubernetesPatchCrdUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd unauthorized response has a 3xx status code
func (o *KubernetesPatchCrdUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd unauthorized response has a 4xx status code
func (o *KubernetesPatchCrdUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd unauthorized response has a 5xx status code
func (o *KubernetesPatchCrdUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd unauthorized response a status code equal to that given
func (o *KubernetesPatchCrdUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchCrdUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchCrdUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchCrdUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchCrdUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdForbidden creates a KubernetesPatchCrdForbidden with default headers values
func NewKubernetesPatchCrdForbidden() *KubernetesPatchCrdForbidden {
	return &KubernetesPatchCrdForbidden{}
}

/*
KubernetesPatchCrdForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchCrdForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch crd forbidden response has a 2xx status code
func (o *KubernetesPatchCrdForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd forbidden response has a 3xx status code
func (o *KubernetesPatchCrdForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd forbidden response has a 4xx status code
func (o *KubernetesPatchCrdForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd forbidden response has a 5xx status code
func (o *KubernetesPatchCrdForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd forbidden response a status code equal to that given
func (o *KubernetesPatchCrdForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchCrdForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchCrdForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchCrdForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchCrdForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdNotFound creates a KubernetesPatchCrdNotFound with default headers values
func NewKubernetesPatchCrdNotFound() *KubernetesPatchCrdNotFound {
	return &KubernetesPatchCrdNotFound{}
}

/*
KubernetesPatchCrdNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchCrdNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch crd not found response has a 2xx status code
func (o *KubernetesPatchCrdNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd not found response has a 3xx status code
func (o *KubernetesPatchCrdNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd not found response has a 4xx status code
func (o *KubernetesPatchCrdNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch crd not found response has a 5xx status code
func (o *KubernetesPatchCrdNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch crd not found response a status code equal to that given
func (o *KubernetesPatchCrdNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchCrdNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchCrdNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchCrdNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchCrdNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchCrdInternalServerError creates a KubernetesPatchCrdInternalServerError with default headers values
func NewKubernetesPatchCrdInternalServerError() *KubernetesPatchCrdInternalServerError {
	return &KubernetesPatchCrdInternalServerError{}
}

/*
KubernetesPatchCrdInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchCrdInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch crd internal server error response has a 2xx status code
func (o *KubernetesPatchCrdInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch crd internal server error response has a 3xx status code
func (o *KubernetesPatchCrdInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch crd internal server error response has a 4xx status code
func (o *KubernetesPatchCrdInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch crd internal server error response has a 5xx status code
func (o *KubernetesPatchCrdInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch crd internal server error response a status code equal to that given
func (o *KubernetesPatchCrdInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchCrdInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdInternalServerError ", 500)
}

func (o *KubernetesPatchCrdInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/crd][%d] kubernetesPatchCrdInternalServerError ", 500)
}

func (o *KubernetesPatchCrdInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
