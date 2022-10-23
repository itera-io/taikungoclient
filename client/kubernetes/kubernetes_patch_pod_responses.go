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

// KubernetesPatchPodReader is a Reader for the KubernetesPatchPod structure.
type KubernetesPatchPodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchPodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchPodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchPodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchPodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchPodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchPodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchPodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchPodOK creates a KubernetesPatchPodOK with default headers values
func NewKubernetesPatchPodOK() *KubernetesPatchPodOK {
	return &KubernetesPatchPodOK{}
}

/*
KubernetesPatchPodOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchPodOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch pod o k response has a 2xx status code
func (o *KubernetesPatchPodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch pod o k response has a 3xx status code
func (o *KubernetesPatchPodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pod o k response has a 4xx status code
func (o *KubernetesPatchPodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch pod o k response has a 5xx status code
func (o *KubernetesPatchPodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pod o k response a status code equal to that given
func (o *KubernetesPatchPodOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchPodOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchPodOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchPodOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchPodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPodBadRequest creates a KubernetesPatchPodBadRequest with default headers values
func NewKubernetesPatchPodBadRequest() *KubernetesPatchPodBadRequest {
	return &KubernetesPatchPodBadRequest{}
}

/*
KubernetesPatchPodBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchPodBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes patch pod bad request response has a 2xx status code
func (o *KubernetesPatchPodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pod bad request response has a 3xx status code
func (o *KubernetesPatchPodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pod bad request response has a 4xx status code
func (o *KubernetesPatchPodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pod bad request response has a 5xx status code
func (o *KubernetesPatchPodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pod bad request response a status code equal to that given
func (o *KubernetesPatchPodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchPodBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchPodBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchPodBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesPatchPodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPodUnauthorized creates a KubernetesPatchPodUnauthorized with default headers values
func NewKubernetesPatchPodUnauthorized() *KubernetesPatchPodUnauthorized {
	return &KubernetesPatchPodUnauthorized{}
}

/*
KubernetesPatchPodUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchPodUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch pod unauthorized response has a 2xx status code
func (o *KubernetesPatchPodUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pod unauthorized response has a 3xx status code
func (o *KubernetesPatchPodUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pod unauthorized response has a 4xx status code
func (o *KubernetesPatchPodUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pod unauthorized response has a 5xx status code
func (o *KubernetesPatchPodUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pod unauthorized response a status code equal to that given
func (o *KubernetesPatchPodUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchPodUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchPodUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchPodUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPodForbidden creates a KubernetesPatchPodForbidden with default headers values
func NewKubernetesPatchPodForbidden() *KubernetesPatchPodForbidden {
	return &KubernetesPatchPodForbidden{}
}

/*
KubernetesPatchPodForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchPodForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch pod forbidden response has a 2xx status code
func (o *KubernetesPatchPodForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pod forbidden response has a 3xx status code
func (o *KubernetesPatchPodForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pod forbidden response has a 4xx status code
func (o *KubernetesPatchPodForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pod forbidden response has a 5xx status code
func (o *KubernetesPatchPodForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pod forbidden response a status code equal to that given
func (o *KubernetesPatchPodForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchPodForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchPodForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchPodForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPodNotFound creates a KubernetesPatchPodNotFound with default headers values
func NewKubernetesPatchPodNotFound() *KubernetesPatchPodNotFound {
	return &KubernetesPatchPodNotFound{}
}

/*
KubernetesPatchPodNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchPodNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch pod not found response has a 2xx status code
func (o *KubernetesPatchPodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pod not found response has a 3xx status code
func (o *KubernetesPatchPodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pod not found response has a 4xx status code
func (o *KubernetesPatchPodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pod not found response has a 5xx status code
func (o *KubernetesPatchPodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pod not found response a status code equal to that given
func (o *KubernetesPatchPodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchPodNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchPodNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchPodNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPodInternalServerError creates a KubernetesPatchPodInternalServerError with default headers values
func NewKubernetesPatchPodInternalServerError() *KubernetesPatchPodInternalServerError {
	return &KubernetesPatchPodInternalServerError{}
}

/*
KubernetesPatchPodInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchPodInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch pod internal server error response has a 2xx status code
func (o *KubernetesPatchPodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pod internal server error response has a 3xx status code
func (o *KubernetesPatchPodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pod internal server error response has a 4xx status code
func (o *KubernetesPatchPodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch pod internal server error response has a 5xx status code
func (o *KubernetesPatchPodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch pod internal server error response a status code equal to that given
func (o *KubernetesPatchPodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchPodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodInternalServerError ", 500)
}

func (o *KubernetesPatchPodInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pod][%d] kubernetesPatchPodInternalServerError ", 500)
}

func (o *KubernetesPatchPodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
