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

// KubernetesPatchJobReader is a Reader for the KubernetesPatchJob structure.
type KubernetesPatchJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchJobOK creates a KubernetesPatchJobOK with default headers values
func NewKubernetesPatchJobOK() *KubernetesPatchJobOK {
	return &KubernetesPatchJobOK{}
}

/*
KubernetesPatchJobOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchJobOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch job o k response has a 2xx status code
func (o *KubernetesPatchJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch job o k response has a 3xx status code
func (o *KubernetesPatchJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch job o k response has a 4xx status code
func (o *KubernetesPatchJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch job o k response has a 5xx status code
func (o *KubernetesPatchJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch job o k response a status code equal to that given
func (o *KubernetesPatchJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchJobOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchJobOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchJobOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchJobBadRequest creates a KubernetesPatchJobBadRequest with default headers values
func NewKubernetesPatchJobBadRequest() *KubernetesPatchJobBadRequest {
	return &KubernetesPatchJobBadRequest{}
}

/*
KubernetesPatchJobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchJobBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch job bad request response has a 2xx status code
func (o *KubernetesPatchJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch job bad request response has a 3xx status code
func (o *KubernetesPatchJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch job bad request response has a 4xx status code
func (o *KubernetesPatchJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch job bad request response has a 5xx status code
func (o *KubernetesPatchJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch job bad request response a status code equal to that given
func (o *KubernetesPatchJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchJobBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchJobBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchJobUnauthorized creates a KubernetesPatchJobUnauthorized with default headers values
func NewKubernetesPatchJobUnauthorized() *KubernetesPatchJobUnauthorized {
	return &KubernetesPatchJobUnauthorized{}
}

/*
KubernetesPatchJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchJobUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch job unauthorized response has a 2xx status code
func (o *KubernetesPatchJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch job unauthorized response has a 3xx status code
func (o *KubernetesPatchJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch job unauthorized response has a 4xx status code
func (o *KubernetesPatchJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch job unauthorized response has a 5xx status code
func (o *KubernetesPatchJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch job unauthorized response a status code equal to that given
func (o *KubernetesPatchJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchJobUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchJobForbidden creates a KubernetesPatchJobForbidden with default headers values
func NewKubernetesPatchJobForbidden() *KubernetesPatchJobForbidden {
	return &KubernetesPatchJobForbidden{}
}

/*
KubernetesPatchJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchJobForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch job forbidden response has a 2xx status code
func (o *KubernetesPatchJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch job forbidden response has a 3xx status code
func (o *KubernetesPatchJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch job forbidden response has a 4xx status code
func (o *KubernetesPatchJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch job forbidden response has a 5xx status code
func (o *KubernetesPatchJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch job forbidden response a status code equal to that given
func (o *KubernetesPatchJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchJobForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchJobForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchJobForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchJobNotFound creates a KubernetesPatchJobNotFound with default headers values
func NewKubernetesPatchJobNotFound() *KubernetesPatchJobNotFound {
	return &KubernetesPatchJobNotFound{}
}

/*
KubernetesPatchJobNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchJobNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch job not found response has a 2xx status code
func (o *KubernetesPatchJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch job not found response has a 3xx status code
func (o *KubernetesPatchJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch job not found response has a 4xx status code
func (o *KubernetesPatchJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch job not found response has a 5xx status code
func (o *KubernetesPatchJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch job not found response a status code equal to that given
func (o *KubernetesPatchJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchJobNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchJobNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchJobInternalServerError creates a KubernetesPatchJobInternalServerError with default headers values
func NewKubernetesPatchJobInternalServerError() *KubernetesPatchJobInternalServerError {
	return &KubernetesPatchJobInternalServerError{}
}

/*
KubernetesPatchJobInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchJobInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch job internal server error response has a 2xx status code
func (o *KubernetesPatchJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch job internal server error response has a 3xx status code
func (o *KubernetesPatchJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch job internal server error response has a 4xx status code
func (o *KubernetesPatchJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch job internal server error response has a 5xx status code
func (o *KubernetesPatchJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch job internal server error response a status code equal to that given
func (o *KubernetesPatchJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobInternalServerError ", 500)
}

func (o *KubernetesPatchJobInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/job][%d] kubernetesPatchJobInternalServerError ", 500)
}

func (o *KubernetesPatchJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
