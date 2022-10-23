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

// KubernetesPatchStsReader is a Reader for the KubernetesPatchSts structure.
type KubernetesPatchStsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchStsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchStsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchStsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchStsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchStsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchStsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchStsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchStsOK creates a KubernetesPatchStsOK with default headers values
func NewKubernetesPatchStsOK() *KubernetesPatchStsOK {
	return &KubernetesPatchStsOK{}
}

/*
KubernetesPatchStsOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchStsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch sts o k response has a 2xx status code
func (o *KubernetesPatchStsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch sts o k response has a 3xx status code
func (o *KubernetesPatchStsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch sts o k response has a 4xx status code
func (o *KubernetesPatchStsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch sts o k response has a 5xx status code
func (o *KubernetesPatchStsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch sts o k response a status code equal to that given
func (o *KubernetesPatchStsOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchStsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchStsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchStsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchStsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchStsBadRequest creates a KubernetesPatchStsBadRequest with default headers values
func NewKubernetesPatchStsBadRequest() *KubernetesPatchStsBadRequest {
	return &KubernetesPatchStsBadRequest{}
}

/*
KubernetesPatchStsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchStsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes patch sts bad request response has a 2xx status code
func (o *KubernetesPatchStsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch sts bad request response has a 3xx status code
func (o *KubernetesPatchStsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch sts bad request response has a 4xx status code
func (o *KubernetesPatchStsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch sts bad request response has a 5xx status code
func (o *KubernetesPatchStsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch sts bad request response a status code equal to that given
func (o *KubernetesPatchStsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchStsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchStsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchStsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesPatchStsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchStsUnauthorized creates a KubernetesPatchStsUnauthorized with default headers values
func NewKubernetesPatchStsUnauthorized() *KubernetesPatchStsUnauthorized {
	return &KubernetesPatchStsUnauthorized{}
}

/*
KubernetesPatchStsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchStsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch sts unauthorized response has a 2xx status code
func (o *KubernetesPatchStsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch sts unauthorized response has a 3xx status code
func (o *KubernetesPatchStsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch sts unauthorized response has a 4xx status code
func (o *KubernetesPatchStsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch sts unauthorized response has a 5xx status code
func (o *KubernetesPatchStsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch sts unauthorized response a status code equal to that given
func (o *KubernetesPatchStsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchStsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchStsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchStsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchStsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchStsForbidden creates a KubernetesPatchStsForbidden with default headers values
func NewKubernetesPatchStsForbidden() *KubernetesPatchStsForbidden {
	return &KubernetesPatchStsForbidden{}
}

/*
KubernetesPatchStsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchStsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch sts forbidden response has a 2xx status code
func (o *KubernetesPatchStsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch sts forbidden response has a 3xx status code
func (o *KubernetesPatchStsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch sts forbidden response has a 4xx status code
func (o *KubernetesPatchStsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch sts forbidden response has a 5xx status code
func (o *KubernetesPatchStsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch sts forbidden response a status code equal to that given
func (o *KubernetesPatchStsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchStsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchStsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchStsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchStsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchStsNotFound creates a KubernetesPatchStsNotFound with default headers values
func NewKubernetesPatchStsNotFound() *KubernetesPatchStsNotFound {
	return &KubernetesPatchStsNotFound{}
}

/*
KubernetesPatchStsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchStsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch sts not found response has a 2xx status code
func (o *KubernetesPatchStsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch sts not found response has a 3xx status code
func (o *KubernetesPatchStsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch sts not found response has a 4xx status code
func (o *KubernetesPatchStsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch sts not found response has a 5xx status code
func (o *KubernetesPatchStsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch sts not found response a status code equal to that given
func (o *KubernetesPatchStsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchStsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchStsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchStsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchStsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchStsInternalServerError creates a KubernetesPatchStsInternalServerError with default headers values
func NewKubernetesPatchStsInternalServerError() *KubernetesPatchStsInternalServerError {
	return &KubernetesPatchStsInternalServerError{}
}

/*
KubernetesPatchStsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchStsInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch sts internal server error response has a 2xx status code
func (o *KubernetesPatchStsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch sts internal server error response has a 3xx status code
func (o *KubernetesPatchStsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch sts internal server error response has a 4xx status code
func (o *KubernetesPatchStsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch sts internal server error response has a 5xx status code
func (o *KubernetesPatchStsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch sts internal server error response a status code equal to that given
func (o *KubernetesPatchStsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchStsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsInternalServerError ", 500)
}

func (o *KubernetesPatchStsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/sts][%d] kubernetesPatchStsInternalServerError ", 500)
}

func (o *KubernetesPatchStsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
