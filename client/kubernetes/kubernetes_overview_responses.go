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

// KubernetesOverviewReader is a Reader for the KubernetesOverview structure.
type KubernetesOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesOverviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesOverviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesOverviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesOverviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesOverviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesOverviewOK creates a KubernetesOverviewOK with default headers values
func NewKubernetesOverviewOK() *KubernetesOverviewOK {
	return &KubernetesOverviewOK{}
}

/*
KubernetesOverviewOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesOverviewOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes overview o k response has a 2xx status code
func (o *KubernetesOverviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes overview o k response has a 3xx status code
func (o *KubernetesOverviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes overview o k response has a 4xx status code
func (o *KubernetesOverviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes overview o k response has a 5xx status code
func (o *KubernetesOverviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes overview o k response a status code equal to that given
func (o *KubernetesOverviewOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesOverviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewOK  %+v", 200, o.Payload)
}

func (o *KubernetesOverviewOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewOK  %+v", 200, o.Payload)
}

func (o *KubernetesOverviewOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewBadRequest creates a KubernetesOverviewBadRequest with default headers values
func NewKubernetesOverviewBadRequest() *KubernetesOverviewBadRequest {
	return &KubernetesOverviewBadRequest{}
}

/*
KubernetesOverviewBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesOverviewBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes overview bad request response has a 2xx status code
func (o *KubernetesOverviewBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes overview bad request response has a 3xx status code
func (o *KubernetesOverviewBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes overview bad request response has a 4xx status code
func (o *KubernetesOverviewBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes overview bad request response has a 5xx status code
func (o *KubernetesOverviewBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes overview bad request response a status code equal to that given
func (o *KubernetesOverviewBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesOverviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesOverviewBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesOverviewBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesOverviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewUnauthorized creates a KubernetesOverviewUnauthorized with default headers values
func NewKubernetesOverviewUnauthorized() *KubernetesOverviewUnauthorized {
	return &KubernetesOverviewUnauthorized{}
}

/*
KubernetesOverviewUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesOverviewUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes overview unauthorized response has a 2xx status code
func (o *KubernetesOverviewUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes overview unauthorized response has a 3xx status code
func (o *KubernetesOverviewUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes overview unauthorized response has a 4xx status code
func (o *KubernetesOverviewUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes overview unauthorized response has a 5xx status code
func (o *KubernetesOverviewUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes overview unauthorized response a status code equal to that given
func (o *KubernetesOverviewUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesOverviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesOverviewUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesOverviewUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesOverviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewForbidden creates a KubernetesOverviewForbidden with default headers values
func NewKubernetesOverviewForbidden() *KubernetesOverviewForbidden {
	return &KubernetesOverviewForbidden{}
}

/*
KubernetesOverviewForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesOverviewForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes overview forbidden response has a 2xx status code
func (o *KubernetesOverviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes overview forbidden response has a 3xx status code
func (o *KubernetesOverviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes overview forbidden response has a 4xx status code
func (o *KubernetesOverviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes overview forbidden response has a 5xx status code
func (o *KubernetesOverviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes overview forbidden response a status code equal to that given
func (o *KubernetesOverviewForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesOverviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesOverviewForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesOverviewForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesOverviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewNotFound creates a KubernetesOverviewNotFound with default headers values
func NewKubernetesOverviewNotFound() *KubernetesOverviewNotFound {
	return &KubernetesOverviewNotFound{}
}

/*
KubernetesOverviewNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesOverviewNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this kubernetes overview not found response has a 2xx status code
func (o *KubernetesOverviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes overview not found response has a 3xx status code
func (o *KubernetesOverviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes overview not found response has a 4xx status code
func (o *KubernetesOverviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes overview not found response has a 5xx status code
func (o *KubernetesOverviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes overview not found response a status code equal to that given
func (o *KubernetesOverviewNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesOverviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesOverviewNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesOverviewNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *KubernetesOverviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesOverviewInternalServerError creates a KubernetesOverviewInternalServerError with default headers values
func NewKubernetesOverviewInternalServerError() *KubernetesOverviewInternalServerError {
	return &KubernetesOverviewInternalServerError{}
}

/*
KubernetesOverviewInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesOverviewInternalServerError struct {
}

// IsSuccess returns true when this kubernetes overview internal server error response has a 2xx status code
func (o *KubernetesOverviewInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes overview internal server error response has a 3xx status code
func (o *KubernetesOverviewInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes overview internal server error response has a 4xx status code
func (o *KubernetesOverviewInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes overview internal server error response has a 5xx status code
func (o *KubernetesOverviewInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes overview internal server error response a status code equal to that given
func (o *KubernetesOverviewInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesOverviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewInternalServerError ", 500)
}

func (o *KubernetesOverviewInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/overview][%d] kubernetesOverviewInternalServerError ", 500)
}

func (o *KubernetesOverviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
