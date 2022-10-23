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

// KubernetesGetStreamPodLogsListReader is a Reader for the KubernetesGetStreamPodLogsList structure.
type KubernetesGetStreamPodLogsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetStreamPodLogsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetStreamPodLogsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetStreamPodLogsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetStreamPodLogsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetStreamPodLogsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetStreamPodLogsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetStreamPodLogsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetStreamPodLogsListOK creates a KubernetesGetStreamPodLogsListOK with default headers values
func NewKubernetesGetStreamPodLogsListOK() *KubernetesGetStreamPodLogsListOK {
	return &KubernetesGetStreamPodLogsListOK{}
}

/*
KubernetesGetStreamPodLogsListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetStreamPodLogsListOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes get stream pod logs list o k response has a 2xx status code
func (o *KubernetesGetStreamPodLogsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get stream pod logs list o k response has a 3xx status code
func (o *KubernetesGetStreamPodLogsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get stream pod logs list o k response has a 4xx status code
func (o *KubernetesGetStreamPodLogsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get stream pod logs list o k response has a 5xx status code
func (o *KubernetesGetStreamPodLogsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get stream pod logs list o k response a status code equal to that given
func (o *KubernetesGetStreamPodLogsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetStreamPodLogsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesGetStreamPodLogsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetStreamPodLogsListBadRequest creates a KubernetesGetStreamPodLogsListBadRequest with default headers values
func NewKubernetesGetStreamPodLogsListBadRequest() *KubernetesGetStreamPodLogsListBadRequest {
	return &KubernetesGetStreamPodLogsListBadRequest{}
}

/*
KubernetesGetStreamPodLogsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetStreamPodLogsListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes get stream pod logs list bad request response has a 2xx status code
func (o *KubernetesGetStreamPodLogsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get stream pod logs list bad request response has a 3xx status code
func (o *KubernetesGetStreamPodLogsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get stream pod logs list bad request response has a 4xx status code
func (o *KubernetesGetStreamPodLogsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get stream pod logs list bad request response has a 5xx status code
func (o *KubernetesGetStreamPodLogsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get stream pod logs list bad request response a status code equal to that given
func (o *KubernetesGetStreamPodLogsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetStreamPodLogsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesGetStreamPodLogsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetStreamPodLogsListUnauthorized creates a KubernetesGetStreamPodLogsListUnauthorized with default headers values
func NewKubernetesGetStreamPodLogsListUnauthorized() *KubernetesGetStreamPodLogsListUnauthorized {
	return &KubernetesGetStreamPodLogsListUnauthorized{}
}

/*
KubernetesGetStreamPodLogsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetStreamPodLogsListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get stream pod logs list unauthorized response has a 2xx status code
func (o *KubernetesGetStreamPodLogsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get stream pod logs list unauthorized response has a 3xx status code
func (o *KubernetesGetStreamPodLogsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get stream pod logs list unauthorized response has a 4xx status code
func (o *KubernetesGetStreamPodLogsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get stream pod logs list unauthorized response has a 5xx status code
func (o *KubernetesGetStreamPodLogsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get stream pod logs list unauthorized response a status code equal to that given
func (o *KubernetesGetStreamPodLogsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetStreamPodLogsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetStreamPodLogsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetStreamPodLogsListForbidden creates a KubernetesGetStreamPodLogsListForbidden with default headers values
func NewKubernetesGetStreamPodLogsListForbidden() *KubernetesGetStreamPodLogsListForbidden {
	return &KubernetesGetStreamPodLogsListForbidden{}
}

/*
KubernetesGetStreamPodLogsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetStreamPodLogsListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get stream pod logs list forbidden response has a 2xx status code
func (o *KubernetesGetStreamPodLogsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get stream pod logs list forbidden response has a 3xx status code
func (o *KubernetesGetStreamPodLogsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get stream pod logs list forbidden response has a 4xx status code
func (o *KubernetesGetStreamPodLogsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get stream pod logs list forbidden response has a 5xx status code
func (o *KubernetesGetStreamPodLogsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get stream pod logs list forbidden response a status code equal to that given
func (o *KubernetesGetStreamPodLogsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetStreamPodLogsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetStreamPodLogsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetStreamPodLogsListNotFound creates a KubernetesGetStreamPodLogsListNotFound with default headers values
func NewKubernetesGetStreamPodLogsListNotFound() *KubernetesGetStreamPodLogsListNotFound {
	return &KubernetesGetStreamPodLogsListNotFound{}
}

/*
KubernetesGetStreamPodLogsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetStreamPodLogsListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes get stream pod logs list not found response has a 2xx status code
func (o *KubernetesGetStreamPodLogsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get stream pod logs list not found response has a 3xx status code
func (o *KubernetesGetStreamPodLogsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get stream pod logs list not found response has a 4xx status code
func (o *KubernetesGetStreamPodLogsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get stream pod logs list not found response has a 5xx status code
func (o *KubernetesGetStreamPodLogsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get stream pod logs list not found response a status code equal to that given
func (o *KubernetesGetStreamPodLogsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetStreamPodLogsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetStreamPodLogsListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesGetStreamPodLogsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetStreamPodLogsListInternalServerError creates a KubernetesGetStreamPodLogsListInternalServerError with default headers values
func NewKubernetesGetStreamPodLogsListInternalServerError() *KubernetesGetStreamPodLogsListInternalServerError {
	return &KubernetesGetStreamPodLogsListInternalServerError{}
}

/*
KubernetesGetStreamPodLogsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetStreamPodLogsListInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get stream pod logs list internal server error response has a 2xx status code
func (o *KubernetesGetStreamPodLogsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get stream pod logs list internal server error response has a 3xx status code
func (o *KubernetesGetStreamPodLogsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get stream pod logs list internal server error response has a 4xx status code
func (o *KubernetesGetStreamPodLogsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get stream pod logs list internal server error response has a 5xx status code
func (o *KubernetesGetStreamPodLogsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get stream pod logs list internal server error response a status code equal to that given
func (o *KubernetesGetStreamPodLogsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetStreamPodLogsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListInternalServerError ", 500)
}

func (o *KubernetesGetStreamPodLogsListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/podlogs-stream][%d] kubernetesGetStreamPodLogsListInternalServerError ", 500)
}

func (o *KubernetesGetStreamPodLogsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
