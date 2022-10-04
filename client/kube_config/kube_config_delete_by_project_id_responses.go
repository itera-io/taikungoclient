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

// KubeConfigDeleteByProjectIDReader is a Reader for the KubeConfigDeleteByProjectID structure.
type KubeConfigDeleteByProjectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubeConfigDeleteByProjectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubeConfigDeleteByProjectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubeConfigDeleteByProjectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubeConfigDeleteByProjectIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubeConfigDeleteByProjectIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubeConfigDeleteByProjectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubeConfigDeleteByProjectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubeConfigDeleteByProjectIDOK creates a KubeConfigDeleteByProjectIDOK with default headers values
func NewKubeConfigDeleteByProjectIDOK() *KubeConfigDeleteByProjectIDOK {
	return &KubeConfigDeleteByProjectIDOK{}
}

/*
KubeConfigDeleteByProjectIDOK describes a response with status code 200, with default header values.

Success
*/
type KubeConfigDeleteByProjectIDOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kube config delete by project Id o k response has a 2xx status code
func (o *KubeConfigDeleteByProjectIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kube config delete by project Id o k response has a 3xx status code
func (o *KubeConfigDeleteByProjectIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete by project Id o k response has a 4xx status code
func (o *KubeConfigDeleteByProjectIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kube config delete by project Id o k response has a 5xx status code
func (o *KubeConfigDeleteByProjectIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete by project Id o k response a status code equal to that given
func (o *KubeConfigDeleteByProjectIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubeConfigDeleteByProjectIDOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdOK  %+v", 200, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdOK  %+v", 200, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubeConfigDeleteByProjectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteByProjectIDBadRequest creates a KubeConfigDeleteByProjectIDBadRequest with default headers values
func NewKubeConfigDeleteByProjectIDBadRequest() *KubeConfigDeleteByProjectIDBadRequest {
	return &KubeConfigDeleteByProjectIDBadRequest{}
}

/*
KubeConfigDeleteByProjectIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubeConfigDeleteByProjectIDBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kube config delete by project Id bad request response has a 2xx status code
func (o *KubeConfigDeleteByProjectIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete by project Id bad request response has a 3xx status code
func (o *KubeConfigDeleteByProjectIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete by project Id bad request response has a 4xx status code
func (o *KubeConfigDeleteByProjectIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete by project Id bad request response has a 5xx status code
func (o *KubeConfigDeleteByProjectIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete by project Id bad request response a status code equal to that given
func (o *KubeConfigDeleteByProjectIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubeConfigDeleteByProjectIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdBadRequest  %+v", 400, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdBadRequest  %+v", 400, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubeConfigDeleteByProjectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteByProjectIDUnauthorized creates a KubeConfigDeleteByProjectIDUnauthorized with default headers values
func NewKubeConfigDeleteByProjectIDUnauthorized() *KubeConfigDeleteByProjectIDUnauthorized {
	return &KubeConfigDeleteByProjectIDUnauthorized{}
}

/*
KubeConfigDeleteByProjectIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubeConfigDeleteByProjectIDUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kube config delete by project Id unauthorized response has a 2xx status code
func (o *KubeConfigDeleteByProjectIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete by project Id unauthorized response has a 3xx status code
func (o *KubeConfigDeleteByProjectIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete by project Id unauthorized response has a 4xx status code
func (o *KubeConfigDeleteByProjectIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete by project Id unauthorized response has a 5xx status code
func (o *KubeConfigDeleteByProjectIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete by project Id unauthorized response a status code equal to that given
func (o *KubeConfigDeleteByProjectIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubeConfigDeleteByProjectIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdUnauthorized  %+v", 401, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdUnauthorized  %+v", 401, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubeConfigDeleteByProjectIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteByProjectIDForbidden creates a KubeConfigDeleteByProjectIDForbidden with default headers values
func NewKubeConfigDeleteByProjectIDForbidden() *KubeConfigDeleteByProjectIDForbidden {
	return &KubeConfigDeleteByProjectIDForbidden{}
}

/*
KubeConfigDeleteByProjectIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubeConfigDeleteByProjectIDForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kube config delete by project Id forbidden response has a 2xx status code
func (o *KubeConfigDeleteByProjectIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete by project Id forbidden response has a 3xx status code
func (o *KubeConfigDeleteByProjectIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete by project Id forbidden response has a 4xx status code
func (o *KubeConfigDeleteByProjectIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete by project Id forbidden response has a 5xx status code
func (o *KubeConfigDeleteByProjectIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete by project Id forbidden response a status code equal to that given
func (o *KubeConfigDeleteByProjectIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubeConfigDeleteByProjectIDForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdForbidden  %+v", 403, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdForbidden  %+v", 403, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubeConfigDeleteByProjectIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteByProjectIDNotFound creates a KubeConfigDeleteByProjectIDNotFound with default headers values
func NewKubeConfigDeleteByProjectIDNotFound() *KubeConfigDeleteByProjectIDNotFound {
	return &KubeConfigDeleteByProjectIDNotFound{}
}

/*
KubeConfigDeleteByProjectIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubeConfigDeleteByProjectIDNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kube config delete by project Id not found response has a 2xx status code
func (o *KubeConfigDeleteByProjectIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete by project Id not found response has a 3xx status code
func (o *KubeConfigDeleteByProjectIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete by project Id not found response has a 4xx status code
func (o *KubeConfigDeleteByProjectIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kube config delete by project Id not found response has a 5xx status code
func (o *KubeConfigDeleteByProjectIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kube config delete by project Id not found response a status code equal to that given
func (o *KubeConfigDeleteByProjectIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubeConfigDeleteByProjectIDNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdNotFound  %+v", 404, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdNotFound  %+v", 404, o.Payload)
}

func (o *KubeConfigDeleteByProjectIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubeConfigDeleteByProjectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigDeleteByProjectIDInternalServerError creates a KubeConfigDeleteByProjectIDInternalServerError with default headers values
func NewKubeConfigDeleteByProjectIDInternalServerError() *KubeConfigDeleteByProjectIDInternalServerError {
	return &KubeConfigDeleteByProjectIDInternalServerError{}
}

/*
KubeConfigDeleteByProjectIDInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubeConfigDeleteByProjectIDInternalServerError struct {
}

// IsSuccess returns true when this kube config delete by project Id internal server error response has a 2xx status code
func (o *KubeConfigDeleteByProjectIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kube config delete by project Id internal server error response has a 3xx status code
func (o *KubeConfigDeleteByProjectIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kube config delete by project Id internal server error response has a 4xx status code
func (o *KubeConfigDeleteByProjectIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kube config delete by project Id internal server error response has a 5xx status code
func (o *KubeConfigDeleteByProjectIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kube config delete by project Id internal server error response a status code equal to that given
func (o *KubeConfigDeleteByProjectIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubeConfigDeleteByProjectIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdInternalServerError ", 500)
}

func (o *KubeConfigDeleteByProjectIDInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/KubeConfig/delete-by-project-id][%d] kubeConfigDeleteByProjectIdInternalServerError ", 500)
}

func (o *KubeConfigDeleteByProjectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
