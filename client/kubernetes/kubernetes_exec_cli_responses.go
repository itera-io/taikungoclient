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

// KubernetesExecCliReader is a Reader for the KubernetesExecCli structure.
type KubernetesExecCliReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesExecCliReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesExecCliOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesExecCliBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesExecCliUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesExecCliForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesExecCliNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesExecCliInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesExecCliOK creates a KubernetesExecCliOK with default headers values
func NewKubernetesExecCliOK() *KubernetesExecCliOK {
	return &KubernetesExecCliOK{}
}

/*
KubernetesExecCliOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesExecCliOK struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes exec cli o k response has a 2xx status code
func (o *KubernetesExecCliOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes exec cli o k response has a 3xx status code
func (o *KubernetesExecCliOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes exec cli o k response has a 4xx status code
func (o *KubernetesExecCliOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes exec cli o k response has a 5xx status code
func (o *KubernetesExecCliOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes exec cli o k response a status code equal to that given
func (o *KubernetesExecCliOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesExecCliOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliOK  %+v", 200, o.Payload)
}

func (o *KubernetesExecCliOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliOK  %+v", 200, o.Payload)
}

func (o *KubernetesExecCliOK) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesExecCliOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesExecCliBadRequest creates a KubernetesExecCliBadRequest with default headers values
func NewKubernetesExecCliBadRequest() *KubernetesExecCliBadRequest {
	return &KubernetesExecCliBadRequest{}
}

/*
KubernetesExecCliBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesExecCliBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes exec cli bad request response has a 2xx status code
func (o *KubernetesExecCliBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes exec cli bad request response has a 3xx status code
func (o *KubernetesExecCliBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes exec cli bad request response has a 4xx status code
func (o *KubernetesExecCliBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes exec cli bad request response has a 5xx status code
func (o *KubernetesExecCliBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes exec cli bad request response a status code equal to that given
func (o *KubernetesExecCliBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesExecCliBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesExecCliBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesExecCliBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesExecCliBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesExecCliUnauthorized creates a KubernetesExecCliUnauthorized with default headers values
func NewKubernetesExecCliUnauthorized() *KubernetesExecCliUnauthorized {
	return &KubernetesExecCliUnauthorized{}
}

/*
KubernetesExecCliUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesExecCliUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes exec cli unauthorized response has a 2xx status code
func (o *KubernetesExecCliUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes exec cli unauthorized response has a 3xx status code
func (o *KubernetesExecCliUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes exec cli unauthorized response has a 4xx status code
func (o *KubernetesExecCliUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes exec cli unauthorized response has a 5xx status code
func (o *KubernetesExecCliUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes exec cli unauthorized response a status code equal to that given
func (o *KubernetesExecCliUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesExecCliUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesExecCliUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesExecCliUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesExecCliUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesExecCliForbidden creates a KubernetesExecCliForbidden with default headers values
func NewKubernetesExecCliForbidden() *KubernetesExecCliForbidden {
	return &KubernetesExecCliForbidden{}
}

/*
KubernetesExecCliForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesExecCliForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes exec cli forbidden response has a 2xx status code
func (o *KubernetesExecCliForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes exec cli forbidden response has a 3xx status code
func (o *KubernetesExecCliForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes exec cli forbidden response has a 4xx status code
func (o *KubernetesExecCliForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes exec cli forbidden response has a 5xx status code
func (o *KubernetesExecCliForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes exec cli forbidden response a status code equal to that given
func (o *KubernetesExecCliForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesExecCliForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesExecCliForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesExecCliForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesExecCliForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesExecCliNotFound creates a KubernetesExecCliNotFound with default headers values
func NewKubernetesExecCliNotFound() *KubernetesExecCliNotFound {
	return &KubernetesExecCliNotFound{}
}

/*
KubernetesExecCliNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesExecCliNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes exec cli not found response has a 2xx status code
func (o *KubernetesExecCliNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes exec cli not found response has a 3xx status code
func (o *KubernetesExecCliNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes exec cli not found response has a 4xx status code
func (o *KubernetesExecCliNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes exec cli not found response has a 5xx status code
func (o *KubernetesExecCliNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes exec cli not found response a status code equal to that given
func (o *KubernetesExecCliNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesExecCliNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesExecCliNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesExecCliNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesExecCliNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesExecCliInternalServerError creates a KubernetesExecCliInternalServerError with default headers values
func NewKubernetesExecCliInternalServerError() *KubernetesExecCliInternalServerError {
	return &KubernetesExecCliInternalServerError{}
}

/*
KubernetesExecCliInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesExecCliInternalServerError struct {
}

// IsSuccess returns true when this kubernetes exec cli internal server error response has a 2xx status code
func (o *KubernetesExecCliInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes exec cli internal server error response has a 3xx status code
func (o *KubernetesExecCliInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes exec cli internal server error response has a 4xx status code
func (o *KubernetesExecCliInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes exec cli internal server error response has a 5xx status code
func (o *KubernetesExecCliInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes exec cli internal server error response a status code equal to that given
func (o *KubernetesExecCliInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesExecCliInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliInternalServerError ", 500)
}

func (o *KubernetesExecCliInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/cli][%d] kubernetesExecCliInternalServerError ", 500)
}

func (o *KubernetesExecCliInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
