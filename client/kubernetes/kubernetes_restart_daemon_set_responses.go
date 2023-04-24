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

// KubernetesRestartDaemonSetReader is a Reader for the KubernetesRestartDaemonSet structure.
type KubernetesRestartDaemonSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesRestartDaemonSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesRestartDaemonSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesRestartDaemonSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesRestartDaemonSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesRestartDaemonSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesRestartDaemonSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesRestartDaemonSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesRestartDaemonSetOK creates a KubernetesRestartDaemonSetOK with default headers values
func NewKubernetesRestartDaemonSetOK() *KubernetesRestartDaemonSetOK {
	return &KubernetesRestartDaemonSetOK{}
}

/*
KubernetesRestartDaemonSetOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesRestartDaemonSetOK struct {
}

// IsSuccess returns true when this kubernetes restart daemon set o k response has a 2xx status code
func (o *KubernetesRestartDaemonSetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes restart daemon set o k response has a 3xx status code
func (o *KubernetesRestartDaemonSetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart daemon set o k response has a 4xx status code
func (o *KubernetesRestartDaemonSetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes restart daemon set o k response has a 5xx status code
func (o *KubernetesRestartDaemonSetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart daemon set o k response a status code equal to that given
func (o *KubernetesRestartDaemonSetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the kubernetes restart daemon set o k response
func (o *KubernetesRestartDaemonSetOK) Code() int {
	return 200
}

func (o *KubernetesRestartDaemonSetOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetOK ", 200)
}

func (o *KubernetesRestartDaemonSetOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetOK ", 200)
}

func (o *KubernetesRestartDaemonSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewKubernetesRestartDaemonSetBadRequest creates a KubernetesRestartDaemonSetBadRequest with default headers values
func NewKubernetesRestartDaemonSetBadRequest() *KubernetesRestartDaemonSetBadRequest {
	return &KubernetesRestartDaemonSetBadRequest{}
}

/*
KubernetesRestartDaemonSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesRestartDaemonSetBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes restart daemon set bad request response has a 2xx status code
func (o *KubernetesRestartDaemonSetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart daemon set bad request response has a 3xx status code
func (o *KubernetesRestartDaemonSetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart daemon set bad request response has a 4xx status code
func (o *KubernetesRestartDaemonSetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart daemon set bad request response has a 5xx status code
func (o *KubernetesRestartDaemonSetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart daemon set bad request response a status code equal to that given
func (o *KubernetesRestartDaemonSetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the kubernetes restart daemon set bad request response
func (o *KubernetesRestartDaemonSetBadRequest) Code() int {
	return 400
}

func (o *KubernetesRestartDaemonSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesRestartDaemonSetBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesRestartDaemonSetBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetUnauthorized creates a KubernetesRestartDaemonSetUnauthorized with default headers values
func NewKubernetesRestartDaemonSetUnauthorized() *KubernetesRestartDaemonSetUnauthorized {
	return &KubernetesRestartDaemonSetUnauthorized{}
}

/*
KubernetesRestartDaemonSetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesRestartDaemonSetUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes restart daemon set unauthorized response has a 2xx status code
func (o *KubernetesRestartDaemonSetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart daemon set unauthorized response has a 3xx status code
func (o *KubernetesRestartDaemonSetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart daemon set unauthorized response has a 4xx status code
func (o *KubernetesRestartDaemonSetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart daemon set unauthorized response has a 5xx status code
func (o *KubernetesRestartDaemonSetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart daemon set unauthorized response a status code equal to that given
func (o *KubernetesRestartDaemonSetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the kubernetes restart daemon set unauthorized response
func (o *KubernetesRestartDaemonSetUnauthorized) Code() int {
	return 401
}

func (o *KubernetesRestartDaemonSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesRestartDaemonSetUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesRestartDaemonSetUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetForbidden creates a KubernetesRestartDaemonSetForbidden with default headers values
func NewKubernetesRestartDaemonSetForbidden() *KubernetesRestartDaemonSetForbidden {
	return &KubernetesRestartDaemonSetForbidden{}
}

/*
KubernetesRestartDaemonSetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesRestartDaemonSetForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes restart daemon set forbidden response has a 2xx status code
func (o *KubernetesRestartDaemonSetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart daemon set forbidden response has a 3xx status code
func (o *KubernetesRestartDaemonSetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart daemon set forbidden response has a 4xx status code
func (o *KubernetesRestartDaemonSetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart daemon set forbidden response has a 5xx status code
func (o *KubernetesRestartDaemonSetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart daemon set forbidden response a status code equal to that given
func (o *KubernetesRestartDaemonSetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the kubernetes restart daemon set forbidden response
func (o *KubernetesRestartDaemonSetForbidden) Code() int {
	return 403
}

func (o *KubernetesRestartDaemonSetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesRestartDaemonSetForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesRestartDaemonSetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetNotFound creates a KubernetesRestartDaemonSetNotFound with default headers values
func NewKubernetesRestartDaemonSetNotFound() *KubernetesRestartDaemonSetNotFound {
	return &KubernetesRestartDaemonSetNotFound{}
}

/*
KubernetesRestartDaemonSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesRestartDaemonSetNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes restart daemon set not found response has a 2xx status code
func (o *KubernetesRestartDaemonSetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart daemon set not found response has a 3xx status code
func (o *KubernetesRestartDaemonSetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart daemon set not found response has a 4xx status code
func (o *KubernetesRestartDaemonSetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes restart daemon set not found response has a 5xx status code
func (o *KubernetesRestartDaemonSetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes restart daemon set not found response a status code equal to that given
func (o *KubernetesRestartDaemonSetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the kubernetes restart daemon set not found response
func (o *KubernetesRestartDaemonSetNotFound) Code() int {
	return 404
}

func (o *KubernetesRestartDaemonSetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesRestartDaemonSetNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesRestartDaemonSetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesRestartDaemonSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesRestartDaemonSetInternalServerError creates a KubernetesRestartDaemonSetInternalServerError with default headers values
func NewKubernetesRestartDaemonSetInternalServerError() *KubernetesRestartDaemonSetInternalServerError {
	return &KubernetesRestartDaemonSetInternalServerError{}
}

/*
KubernetesRestartDaemonSetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesRestartDaemonSetInternalServerError struct {
}

// IsSuccess returns true when this kubernetes restart daemon set internal server error response has a 2xx status code
func (o *KubernetesRestartDaemonSetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes restart daemon set internal server error response has a 3xx status code
func (o *KubernetesRestartDaemonSetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes restart daemon set internal server error response has a 4xx status code
func (o *KubernetesRestartDaemonSetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes restart daemon set internal server error response has a 5xx status code
func (o *KubernetesRestartDaemonSetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes restart daemon set internal server error response a status code equal to that given
func (o *KubernetesRestartDaemonSetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the kubernetes restart daemon set internal server error response
func (o *KubernetesRestartDaemonSetInternalServerError) Code() int {
	return 500
}

func (o *KubernetesRestartDaemonSetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetInternalServerError ", 500)
}

func (o *KubernetesRestartDaemonSetInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/restart/daemonset][%d] kubernetesRestartDaemonSetInternalServerError ", 500)
}

func (o *KubernetesRestartDaemonSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
