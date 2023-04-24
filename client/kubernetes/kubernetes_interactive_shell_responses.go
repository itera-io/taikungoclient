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

// KubernetesInteractiveShellReader is a Reader for the KubernetesInteractiveShell structure.
type KubernetesInteractiveShellReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesInteractiveShellReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesInteractiveShellOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesInteractiveShellBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesInteractiveShellUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesInteractiveShellForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesInteractiveShellNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesInteractiveShellInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesInteractiveShellOK creates a KubernetesInteractiveShellOK with default headers values
func NewKubernetesInteractiveShellOK() *KubernetesInteractiveShellOK {
	return &KubernetesInteractiveShellOK{}
}

/*
KubernetesInteractiveShellOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesInteractiveShellOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes interactive shell o k response has a 2xx status code
func (o *KubernetesInteractiveShellOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes interactive shell o k response has a 3xx status code
func (o *KubernetesInteractiveShellOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes interactive shell o k response has a 4xx status code
func (o *KubernetesInteractiveShellOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes interactive shell o k response has a 5xx status code
func (o *KubernetesInteractiveShellOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes interactive shell o k response a status code equal to that given
func (o *KubernetesInteractiveShellOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes interactive shell o k response
func (o *KubernetesInteractiveShellOK) Code() int {
	return 200
}

func (o *KubernetesInteractiveShellOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellOK  %+v", 200, o.Payload)
}

func (o *KubernetesInteractiveShellOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellOK  %+v", 200, o.Payload)
}

func (o *KubernetesInteractiveShellOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesInteractiveShellOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesInteractiveShellBadRequest creates a KubernetesInteractiveShellBadRequest with default headers values
func NewKubernetesInteractiveShellBadRequest() *KubernetesInteractiveShellBadRequest {
	return &KubernetesInteractiveShellBadRequest{}
}

/*
KubernetesInteractiveShellBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesInteractiveShellBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes interactive shell bad request response has a 2xx status code
func (o *KubernetesInteractiveShellBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes interactive shell bad request response has a 3xx status code
func (o *KubernetesInteractiveShellBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes interactive shell bad request response has a 4xx status code
func (o *KubernetesInteractiveShellBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes interactive shell bad request response has a 5xx status code
func (o *KubernetesInteractiveShellBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes interactive shell bad request response a status code equal to that given
func (o *KubernetesInteractiveShellBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes interactive shell bad request response
func (o *KubernetesInteractiveShellBadRequest) Code() int {
	return 400
}

func (o *KubernetesInteractiveShellBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesInteractiveShellBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesInteractiveShellBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesInteractiveShellBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesInteractiveShellUnauthorized creates a KubernetesInteractiveShellUnauthorized with default headers values
func NewKubernetesInteractiveShellUnauthorized() *KubernetesInteractiveShellUnauthorized {
	return &KubernetesInteractiveShellUnauthorized{}
}

/*
KubernetesInteractiveShellUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesInteractiveShellUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes interactive shell unauthorized response has a 2xx status code
func (o *KubernetesInteractiveShellUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes interactive shell unauthorized response has a 3xx status code
func (o *KubernetesInteractiveShellUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes interactive shell unauthorized response has a 4xx status code
func (o *KubernetesInteractiveShellUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes interactive shell unauthorized response has a 5xx status code
func (o *KubernetesInteractiveShellUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes interactive shell unauthorized response a status code equal to that given
func (o *KubernetesInteractiveShellUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes interactive shell unauthorized response
func (o *KubernetesInteractiveShellUnauthorized) Code() int {
	return 401
}

func (o *KubernetesInteractiveShellUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesInteractiveShellUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesInteractiveShellUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesInteractiveShellUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesInteractiveShellForbidden creates a KubernetesInteractiveShellForbidden with default headers values
func NewKubernetesInteractiveShellForbidden() *KubernetesInteractiveShellForbidden {
	return &KubernetesInteractiveShellForbidden{}
}

/*
KubernetesInteractiveShellForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesInteractiveShellForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes interactive shell forbidden response has a 2xx status code
func (o *KubernetesInteractiveShellForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes interactive shell forbidden response has a 3xx status code
func (o *KubernetesInteractiveShellForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes interactive shell forbidden response has a 4xx status code
func (o *KubernetesInteractiveShellForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes interactive shell forbidden response has a 5xx status code
func (o *KubernetesInteractiveShellForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes interactive shell forbidden response a status code equal to that given
func (o *KubernetesInteractiveShellForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes interactive shell forbidden response
func (o *KubernetesInteractiveShellForbidden) Code() int {
	return 403
}

func (o *KubernetesInteractiveShellForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesInteractiveShellForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesInteractiveShellForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesInteractiveShellForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesInteractiveShellNotFound creates a KubernetesInteractiveShellNotFound with default headers values
func NewKubernetesInteractiveShellNotFound() *KubernetesInteractiveShellNotFound {
	return &KubernetesInteractiveShellNotFound{}
}

/*
KubernetesInteractiveShellNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesInteractiveShellNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes interactive shell not found response has a 2xx status code
func (o *KubernetesInteractiveShellNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes interactive shell not found response has a 3xx status code
func (o *KubernetesInteractiveShellNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes interactive shell not found response has a 4xx status code
func (o *KubernetesInteractiveShellNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes interactive shell not found response has a 5xx status code
func (o *KubernetesInteractiveShellNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes interactive shell not found response a status code equal to that given
func (o *KubernetesInteractiveShellNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes interactive shell not found response
func (o *KubernetesInteractiveShellNotFound) Code() int {
	return 404
}

func (o *KubernetesInteractiveShellNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesInteractiveShellNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesInteractiveShellNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesInteractiveShellNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesInteractiveShellInternalServerError creates a KubernetesInteractiveShellInternalServerError with default headers values
func NewKubernetesInteractiveShellInternalServerError() *KubernetesInteractiveShellInternalServerError {
	return &KubernetesInteractiveShellInternalServerError{}
}

/*
KubernetesInteractiveShellInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesInteractiveShellInternalServerError struct {
}

// IsSuccess returns true when this kubernetes interactive shell internal server error response has a 2xx status code
func (o *KubernetesInteractiveShellInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes interactive shell internal server error response has a 3xx status code
func (o *KubernetesInteractiveShellInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes interactive shell internal server error response has a 4xx status code
func (o *KubernetesInteractiveShellInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes interactive shell internal server error response has a 5xx status code
func (o *KubernetesInteractiveShellInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes interactive shell internal server error response a status code equal to that given
func (o *KubernetesInteractiveShellInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes interactive shell internal server error response
func (o *KubernetesInteractiveShellInternalServerError) Code() int {
	return 500
}

func (o *KubernetesInteractiveShellInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellInternalServerError ", 500)
}

func (o *KubernetesInteractiveShellInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/interactive-shell][%d] kubernetesInteractiveShellInternalServerError ", 500)
}

func (o *KubernetesInteractiveShellInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
