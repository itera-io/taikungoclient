// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// KubernetesPatchDeploymentReader is a Reader for the KubernetesPatchDeployment structure.
type KubernetesPatchDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchDeploymentOK creates a KubernetesPatchDeploymentOK with default headers values
func NewKubernetesPatchDeploymentOK() *KubernetesPatchDeploymentOK {
	return &KubernetesPatchDeploymentOK{}
}

/*
KubernetesPatchDeploymentOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchDeploymentOK struct {
}

// IsSuccess returns true when this kubernetes patch deployment o k response has a 2xx status code
func (o *KubernetesPatchDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch deployment o k response has a 3xx status code
func (o *KubernetesPatchDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch deployment o k response has a 4xx status code
func (o *KubernetesPatchDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch deployment o k response has a 5xx status code
func (o *KubernetesPatchDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch deployment o k response a status code equal to that given
func (o *KubernetesPatchDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes patch deployment o k response
func (o *KubernetesPatchDeploymentOK) Code() int {
	return 200
}

func (o *KubernetesPatchDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentOK ", 200)
}

func (o *KubernetesPatchDeploymentOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentOK ", 200)
}

func (o *KubernetesPatchDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKubernetesPatchDeploymentBadRequest creates a KubernetesPatchDeploymentBadRequest with default headers values
func NewKubernetesPatchDeploymentBadRequest() *KubernetesPatchDeploymentBadRequest {
	return &KubernetesPatchDeploymentBadRequest{}
}

/*
KubernetesPatchDeploymentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchDeploymentBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch deployment bad request response has a 2xx status code
func (o *KubernetesPatchDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch deployment bad request response has a 3xx status code
func (o *KubernetesPatchDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch deployment bad request response has a 4xx status code
func (o *KubernetesPatchDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch deployment bad request response has a 5xx status code
func (o *KubernetesPatchDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch deployment bad request response a status code equal to that given
func (o *KubernetesPatchDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes patch deployment bad request response
func (o *KubernetesPatchDeploymentBadRequest) Code() int {
	return 400
}

func (o *KubernetesPatchDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchDeploymentBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchDeploymentBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchDeploymentUnauthorized creates a KubernetesPatchDeploymentUnauthorized with default headers values
func NewKubernetesPatchDeploymentUnauthorized() *KubernetesPatchDeploymentUnauthorized {
	return &KubernetesPatchDeploymentUnauthorized{}
}

/*
KubernetesPatchDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchDeploymentUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch deployment unauthorized response has a 2xx status code
func (o *KubernetesPatchDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch deployment unauthorized response has a 3xx status code
func (o *KubernetesPatchDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch deployment unauthorized response has a 4xx status code
func (o *KubernetesPatchDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch deployment unauthorized response has a 5xx status code
func (o *KubernetesPatchDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch deployment unauthorized response a status code equal to that given
func (o *KubernetesPatchDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes patch deployment unauthorized response
func (o *KubernetesPatchDeploymentUnauthorized) Code() int {
	return 401
}

func (o *KubernetesPatchDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchDeploymentUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchDeploymentForbidden creates a KubernetesPatchDeploymentForbidden with default headers values
func NewKubernetesPatchDeploymentForbidden() *KubernetesPatchDeploymentForbidden {
	return &KubernetesPatchDeploymentForbidden{}
}

/*
KubernetesPatchDeploymentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchDeploymentForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch deployment forbidden response has a 2xx status code
func (o *KubernetesPatchDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch deployment forbidden response has a 3xx status code
func (o *KubernetesPatchDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch deployment forbidden response has a 4xx status code
func (o *KubernetesPatchDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch deployment forbidden response has a 5xx status code
func (o *KubernetesPatchDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch deployment forbidden response a status code equal to that given
func (o *KubernetesPatchDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes patch deployment forbidden response
func (o *KubernetesPatchDeploymentForbidden) Code() int {
	return 403
}

func (o *KubernetesPatchDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchDeploymentForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchDeploymentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchDeploymentNotFound creates a KubernetesPatchDeploymentNotFound with default headers values
func NewKubernetesPatchDeploymentNotFound() *KubernetesPatchDeploymentNotFound {
	return &KubernetesPatchDeploymentNotFound{}
}

/*
KubernetesPatchDeploymentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchDeploymentNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes patch deployment not found response has a 2xx status code
func (o *KubernetesPatchDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch deployment not found response has a 3xx status code
func (o *KubernetesPatchDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch deployment not found response has a 4xx status code
func (o *KubernetesPatchDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch deployment not found response has a 5xx status code
func (o *KubernetesPatchDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch deployment not found response a status code equal to that given
func (o *KubernetesPatchDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes patch deployment not found response
func (o *KubernetesPatchDeploymentNotFound) Code() int {
	return 404
}

func (o *KubernetesPatchDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchDeploymentNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchDeploymentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesPatchDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchDeploymentInternalServerError creates a KubernetesPatchDeploymentInternalServerError with default headers values
func NewKubernetesPatchDeploymentInternalServerError() *KubernetesPatchDeploymentInternalServerError {
	return &KubernetesPatchDeploymentInternalServerError{}
}

/*
KubernetesPatchDeploymentInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchDeploymentInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch deployment internal server error response has a 2xx status code
func (o *KubernetesPatchDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch deployment internal server error response has a 3xx status code
func (o *KubernetesPatchDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch deployment internal server error response has a 4xx status code
func (o *KubernetesPatchDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch deployment internal server error response has a 5xx status code
func (o *KubernetesPatchDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch deployment internal server error response a status code equal to that given
func (o *KubernetesPatchDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes patch deployment internal server error response
func (o *KubernetesPatchDeploymentInternalServerError) Code() int {
	return 500
}

func (o *KubernetesPatchDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentInternalServerError ", 500)
}

func (o *KubernetesPatchDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/deployment][%d] kubernetesPatchDeploymentInternalServerError ", 500)
}

func (o *KubernetesPatchDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
