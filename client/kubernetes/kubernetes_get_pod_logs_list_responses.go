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

// KubernetesGetPodLogsListReader is a Reader for the KubernetesGetPodLogsList structure.
type KubernetesGetPodLogsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetPodLogsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetPodLogsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetPodLogsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetPodLogsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetPodLogsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetPodLogsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetPodLogsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetPodLogsListOK creates a KubernetesGetPodLogsListOK with default headers values
func NewKubernetesGetPodLogsListOK() *KubernetesGetPodLogsListOK {
	return &KubernetesGetPodLogsListOK{}
}

/*
KubernetesGetPodLogsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetPodLogsListOK struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get pod logs list o k response has a 2xx status code
func (o *KubernetesGetPodLogsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get pod logs list o k response has a 3xx status code
func (o *KubernetesGetPodLogsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod logs list o k response has a 4xx status code
func (o *KubernetesGetPodLogsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get pod logs list o k response has a 5xx status code
func (o *KubernetesGetPodLogsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod logs list o k response a status code equal to that given
func (o *KubernetesGetPodLogsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetPodLogsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetPodLogsListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetPodLogsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetPodLogsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodLogsListBadRequest creates a KubernetesGetPodLogsListBadRequest with default headers values
func NewKubernetesGetPodLogsListBadRequest() *KubernetesGetPodLogsListBadRequest {
	return &KubernetesGetPodLogsListBadRequest{}
}

/*
KubernetesGetPodLogsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetPodLogsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes get pod logs list bad request response has a 2xx status code
func (o *KubernetesGetPodLogsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod logs list bad request response has a 3xx status code
func (o *KubernetesGetPodLogsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod logs list bad request response has a 4xx status code
func (o *KubernetesGetPodLogsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod logs list bad request response has a 5xx status code
func (o *KubernetesGetPodLogsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod logs list bad request response a status code equal to that given
func (o *KubernetesGetPodLogsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetPodLogsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetPodLogsListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetPodLogsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetPodLogsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodLogsListUnauthorized creates a KubernetesGetPodLogsListUnauthorized with default headers values
func NewKubernetesGetPodLogsListUnauthorized() *KubernetesGetPodLogsListUnauthorized {
	return &KubernetesGetPodLogsListUnauthorized{}
}

/*
KubernetesGetPodLogsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetPodLogsListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get pod logs list unauthorized response has a 2xx status code
func (o *KubernetesGetPodLogsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod logs list unauthorized response has a 3xx status code
func (o *KubernetesGetPodLogsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod logs list unauthorized response has a 4xx status code
func (o *KubernetesGetPodLogsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod logs list unauthorized response has a 5xx status code
func (o *KubernetesGetPodLogsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod logs list unauthorized response a status code equal to that given
func (o *KubernetesGetPodLogsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetPodLogsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetPodLogsListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetPodLogsListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetPodLogsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodLogsListForbidden creates a KubernetesGetPodLogsListForbidden with default headers values
func NewKubernetesGetPodLogsListForbidden() *KubernetesGetPodLogsListForbidden {
	return &KubernetesGetPodLogsListForbidden{}
}

/*
KubernetesGetPodLogsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetPodLogsListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get pod logs list forbidden response has a 2xx status code
func (o *KubernetesGetPodLogsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod logs list forbidden response has a 3xx status code
func (o *KubernetesGetPodLogsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod logs list forbidden response has a 4xx status code
func (o *KubernetesGetPodLogsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod logs list forbidden response has a 5xx status code
func (o *KubernetesGetPodLogsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod logs list forbidden response a status code equal to that given
func (o *KubernetesGetPodLogsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetPodLogsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetPodLogsListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetPodLogsListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetPodLogsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodLogsListNotFound creates a KubernetesGetPodLogsListNotFound with default headers values
func NewKubernetesGetPodLogsListNotFound() *KubernetesGetPodLogsListNotFound {
	return &KubernetesGetPodLogsListNotFound{}
}

/*
KubernetesGetPodLogsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetPodLogsListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes get pod logs list not found response has a 2xx status code
func (o *KubernetesGetPodLogsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod logs list not found response has a 3xx status code
func (o *KubernetesGetPodLogsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod logs list not found response has a 4xx status code
func (o *KubernetesGetPodLogsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get pod logs list not found response has a 5xx status code
func (o *KubernetesGetPodLogsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get pod logs list not found response a status code equal to that given
func (o *KubernetesGetPodLogsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetPodLogsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetPodLogsListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetPodLogsListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesGetPodLogsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetPodLogsListInternalServerError creates a KubernetesGetPodLogsListInternalServerError with default headers values
func NewKubernetesGetPodLogsListInternalServerError() *KubernetesGetPodLogsListInternalServerError {
	return &KubernetesGetPodLogsListInternalServerError{}
}

/*
KubernetesGetPodLogsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetPodLogsListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get pod logs list internal server error response has a 2xx status code
func (o *KubernetesGetPodLogsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get pod logs list internal server error response has a 3xx status code
func (o *KubernetesGetPodLogsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get pod logs list internal server error response has a 4xx status code
func (o *KubernetesGetPodLogsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get pod logs list internal server error response has a 5xx status code
func (o *KubernetesGetPodLogsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get pod logs list internal server error response a status code equal to that given
func (o *KubernetesGetPodLogsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetPodLogsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListInternalServerError ", 500)
}

func (o *KubernetesGetPodLogsListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs][%d] kubernetesGetPodLogsListInternalServerError ", 500)
}

func (o *KubernetesGetPodLogsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
