// Code generated by go-swagger; DO NOT EDIT.

package kube_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubeConfigDeleteReader is a Reader for the KubeConfigDelete structure.
type KubeConfigDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubeConfigDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubeConfigDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubeConfigDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubeConfigDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubeConfigDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubeConfigDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubeConfigDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubeConfigDeleteOK creates a KubeConfigDeleteOK with default headers values
func NewKubeConfigDeleteOK() *KubeConfigDeleteOK {
	return &KubeConfigDeleteOK{}
}

/*
KubeConfigDeleteOK describes a response with status code 200, with default header values.

Success
*/
type KubeConfigDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kube config delete o k response has a 2xx status code
func (o *KubeConfigDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kube config delete o k response has a 3xx status code
func (o *KubeConfigDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete o k response has a 4xx status code
func (o *KubeConfigDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kube config delete o k response has a 5xx status code
func (o *KubeConfigDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete o k response a status code equal to that given
func (o *KubeConfigDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubeConfigDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteOK  %+v", 200, o.Payload)
}

func (o *KubeConfigDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteOK  %+v", 200, o.Payload)
}

func (o *KubeConfigDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubeConfigDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteBadRequest creates a KubeConfigDeleteBadRequest with default headers values
func NewKubeConfigDeleteBadRequest() *KubeConfigDeleteBadRequest {
	return &KubeConfigDeleteBadRequest{}
}

/*
KubeConfigDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubeConfigDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this kube config delete bad request response has a 2xx status code
func (o *KubeConfigDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete bad request response has a 3xx status code
func (o *KubeConfigDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete bad request response has a 4xx status code
func (o *KubeConfigDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete bad request response has a 5xx status code
func (o *KubeConfigDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete bad request response a status code equal to that given
func (o *KubeConfigDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubeConfigDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KubeConfigDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KubeConfigDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *KubeConfigDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteUnauthorized creates a KubeConfigDeleteUnauthorized with default headers values
func NewKubeConfigDeleteUnauthorized() *KubeConfigDeleteUnauthorized {
	return &KubeConfigDeleteUnauthorized{}
}

/*
KubeConfigDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubeConfigDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kube config delete unauthorized response has a 2xx status code
func (o *KubeConfigDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete unauthorized response has a 3xx status code
func (o *KubeConfigDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete unauthorized response has a 4xx status code
func (o *KubeConfigDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete unauthorized response has a 5xx status code
func (o *KubeConfigDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete unauthorized response a status code equal to that given
func (o *KubeConfigDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubeConfigDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KubeConfigDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KubeConfigDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteForbidden creates a KubeConfigDeleteForbidden with default headers values
func NewKubeConfigDeleteForbidden() *KubeConfigDeleteForbidden {
	return &KubeConfigDeleteForbidden{}
}

/*
KubeConfigDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubeConfigDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kube config delete forbidden response has a 2xx status code
func (o *KubeConfigDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete forbidden response has a 3xx status code
func (o *KubeConfigDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete forbidden response has a 4xx status code
func (o *KubeConfigDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete forbidden response has a 5xx status code
func (o *KubeConfigDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete forbidden response a status code equal to that given
func (o *KubeConfigDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubeConfigDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KubeConfigDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KubeConfigDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteNotFound creates a KubeConfigDeleteNotFound with default headers values
func NewKubeConfigDeleteNotFound() *KubeConfigDeleteNotFound {
	return &KubeConfigDeleteNotFound{}
}

/*
KubeConfigDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubeConfigDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kube config delete not found response has a 2xx status code
func (o *KubeConfigDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete not found response has a 3xx status code
func (o *KubeConfigDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete not found response has a 4xx status code
func (o *KubeConfigDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete not found response has a 5xx status code
func (o *KubeConfigDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete not found response a status code equal to that given
func (o *KubeConfigDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubeConfigDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KubeConfigDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KubeConfigDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteInternalServerError creates a KubeConfigDeleteInternalServerError with default headers values
func NewKubeConfigDeleteInternalServerError() *KubeConfigDeleteInternalServerError {
	return &KubeConfigDeleteInternalServerError{}
}

/*
KubeConfigDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubeConfigDeleteInternalServerError struct {
}

// IsSuccess returns true when this kube config delete internal server error response has a 2xx status code
func (o *KubeConfigDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete internal server error response has a 3xx status code
func (o *KubeConfigDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete internal server error response has a 4xx status code
func (o *KubeConfigDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kube config delete internal server error response has a 5xx status code
func (o *KubeConfigDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kube config delete internal server error response a status code equal to that given
func (o *KubeConfigDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubeConfigDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteInternalServerError ", 500)
}

func (o *KubeConfigDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete][%d] kubeConfigDeleteInternalServerError ", 500)
}

func (o *KubeConfigDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
