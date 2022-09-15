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

// KubernetesGetDashboardReader is a Reader for the KubernetesGetDashboard structure.
type KubernetesGetDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesGetDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesGetDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesGetDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesGetDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesGetDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesGetDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesGetDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesGetDashboardOK creates a KubernetesGetDashboardOK with default headers values
func NewKubernetesGetDashboardOK() *KubernetesGetDashboardOK {
	return &KubernetesGetDashboardOK{}
}

/*
KubernetesGetDashboardOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesGetDashboardOK struct {
	Payload string
}

// IsSuccess returns true when this kubernetes get dashboard o k response has a 2xx status code
func (o *KubernetesGetDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes get dashboard o k response has a 3xx status code
func (o *KubernetesGetDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get dashboard o k response has a 4xx status code
func (o *KubernetesGetDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get dashboard o k response has a 5xx status code
func (o *KubernetesGetDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get dashboard o k response a status code equal to that given
func (o *KubernetesGetDashboardOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesGetDashboardOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetDashboardOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardOK  %+v", 200, o.Payload)
}

func (o *KubernetesGetDashboardOK) GetPayload() string {
	return o.Payload
}

func (o *KubernetesGetDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDashboardBadRequest creates a KubernetesGetDashboardBadRequest with default headers values
func NewKubernetesGetDashboardBadRequest() *KubernetesGetDashboardBadRequest {
	return &KubernetesGetDashboardBadRequest{}
}

/*
KubernetesGetDashboardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesGetDashboardBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes get dashboard bad request response has a 2xx status code
func (o *KubernetesGetDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get dashboard bad request response has a 3xx status code
func (o *KubernetesGetDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get dashboard bad request response has a 4xx status code
func (o *KubernetesGetDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get dashboard bad request response has a 5xx status code
func (o *KubernetesGetDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get dashboard bad request response a status code equal to that given
func (o *KubernetesGetDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesGetDashboardBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetDashboardBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesGetDashboardBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDashboardUnauthorized creates a KubernetesGetDashboardUnauthorized with default headers values
func NewKubernetesGetDashboardUnauthorized() *KubernetesGetDashboardUnauthorized {
	return &KubernetesGetDashboardUnauthorized{}
}

/*
KubernetesGetDashboardUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesGetDashboardUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get dashboard unauthorized response has a 2xx status code
func (o *KubernetesGetDashboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get dashboard unauthorized response has a 3xx status code
func (o *KubernetesGetDashboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get dashboard unauthorized response has a 4xx status code
func (o *KubernetesGetDashboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get dashboard unauthorized response has a 5xx status code
func (o *KubernetesGetDashboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get dashboard unauthorized response a status code equal to that given
func (o *KubernetesGetDashboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesGetDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetDashboardUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesGetDashboardUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDashboardForbidden creates a KubernetesGetDashboardForbidden with default headers values
func NewKubernetesGetDashboardForbidden() *KubernetesGetDashboardForbidden {
	return &KubernetesGetDashboardForbidden{}
}

/*
KubernetesGetDashboardForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesGetDashboardForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get dashboard forbidden response has a 2xx status code
func (o *KubernetesGetDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get dashboard forbidden response has a 3xx status code
func (o *KubernetesGetDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get dashboard forbidden response has a 4xx status code
func (o *KubernetesGetDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get dashboard forbidden response has a 5xx status code
func (o *KubernetesGetDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get dashboard forbidden response a status code equal to that given
func (o *KubernetesGetDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesGetDashboardForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetDashboardForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesGetDashboardForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDashboardNotFound creates a KubernetesGetDashboardNotFound with default headers values
func NewKubernetesGetDashboardNotFound() *KubernetesGetDashboardNotFound {
	return &KubernetesGetDashboardNotFound{}
}

/*
KubernetesGetDashboardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesGetDashboardNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes get dashboard not found response has a 2xx status code
func (o *KubernetesGetDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get dashboard not found response has a 3xx status code
func (o *KubernetesGetDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get dashboard not found response has a 4xx status code
func (o *KubernetesGetDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes get dashboard not found response has a 5xx status code
func (o *KubernetesGetDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes get dashboard not found response a status code equal to that given
func (o *KubernetesGetDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesGetDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetDashboardNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesGetDashboardNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesGetDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesGetDashboardInternalServerError creates a KubernetesGetDashboardInternalServerError with default headers values
func NewKubernetesGetDashboardInternalServerError() *KubernetesGetDashboardInternalServerError {
	return &KubernetesGetDashboardInternalServerError{}
}

/*
KubernetesGetDashboardInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesGetDashboardInternalServerError struct {
}

// IsSuccess returns true when this kubernetes get dashboard internal server error response has a 2xx status code
func (o *KubernetesGetDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes get dashboard internal server error response has a 3xx status code
func (o *KubernetesGetDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes get dashboard internal server error response has a 4xx status code
func (o *KubernetesGetDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes get dashboard internal server error response has a 5xx status code
func (o *KubernetesGetDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes get dashboard internal server error response a status code equal to that given
func (o *KubernetesGetDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesGetDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardInternalServerError ", 500)
}

func (o *KubernetesGetDashboardInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Kubernetes/{projectId}/dashboard][%d] kubernetesGetDashboardInternalServerError ", 500)
}

func (o *KubernetesGetDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
