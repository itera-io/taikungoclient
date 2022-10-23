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

// KubernetesPatchPvcReader is a Reader for the KubernetesPatchPvc structure.
type KubernetesPatchPvcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchPvcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchPvcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchPvcBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchPvcUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchPvcForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchPvcNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchPvcInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchPvcOK creates a KubernetesPatchPvcOK with default headers values
func NewKubernetesPatchPvcOK() *KubernetesPatchPvcOK {
	return &KubernetesPatchPvcOK{}
}

/*
KubernetesPatchPvcOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchPvcOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch pvc o k response has a 2xx status code
func (o *KubernetesPatchPvcOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch pvc o k response has a 3xx status code
func (o *KubernetesPatchPvcOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pvc o k response has a 4xx status code
func (o *KubernetesPatchPvcOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch pvc o k response has a 5xx status code
func (o *KubernetesPatchPvcOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pvc o k response a status code equal to that given
func (o *KubernetesPatchPvcOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchPvcOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchPvcOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchPvcOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchPvcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPvcBadRequest creates a KubernetesPatchPvcBadRequest with default headers values
func NewKubernetesPatchPvcBadRequest() *KubernetesPatchPvcBadRequest {
	return &KubernetesPatchPvcBadRequest{}
}

/*
KubernetesPatchPvcBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchPvcBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this kubernetes patch pvc bad request response has a 2xx status code
func (o *KubernetesPatchPvcBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pvc bad request response has a 3xx status code
func (o *KubernetesPatchPvcBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pvc bad request response has a 4xx status code
func (o *KubernetesPatchPvcBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pvc bad request response has a 5xx status code
func (o *KubernetesPatchPvcBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pvc bad request response a status code equal to that given
func (o *KubernetesPatchPvcBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchPvcBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchPvcBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchPvcBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *KubernetesPatchPvcBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPvcUnauthorized creates a KubernetesPatchPvcUnauthorized with default headers values
func NewKubernetesPatchPvcUnauthorized() *KubernetesPatchPvcUnauthorized {
	return &KubernetesPatchPvcUnauthorized{}
}

/*
KubernetesPatchPvcUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchPvcUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes patch pvc unauthorized response has a 2xx status code
func (o *KubernetesPatchPvcUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pvc unauthorized response has a 3xx status code
func (o *KubernetesPatchPvcUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pvc unauthorized response has a 4xx status code
func (o *KubernetesPatchPvcUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pvc unauthorized response has a 5xx status code
func (o *KubernetesPatchPvcUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pvc unauthorized response a status code equal to that given
func (o *KubernetesPatchPvcUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchPvcUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchPvcUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchPvcUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesPatchPvcUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPvcForbidden creates a KubernetesPatchPvcForbidden with default headers values
func NewKubernetesPatchPvcForbidden() *KubernetesPatchPvcForbidden {
	return &KubernetesPatchPvcForbidden{}
}

/*
KubernetesPatchPvcForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchPvcForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes patch pvc forbidden response has a 2xx status code
func (o *KubernetesPatchPvcForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pvc forbidden response has a 3xx status code
func (o *KubernetesPatchPvcForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pvc forbidden response has a 4xx status code
func (o *KubernetesPatchPvcForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pvc forbidden response has a 5xx status code
func (o *KubernetesPatchPvcForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pvc forbidden response a status code equal to that given
func (o *KubernetesPatchPvcForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchPvcForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchPvcForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchPvcForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesPatchPvcForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPvcNotFound creates a KubernetesPatchPvcNotFound with default headers values
func NewKubernetesPatchPvcNotFound() *KubernetesPatchPvcNotFound {
	return &KubernetesPatchPvcNotFound{}
}

/*
KubernetesPatchPvcNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchPvcNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this kubernetes patch pvc not found response has a 2xx status code
func (o *KubernetesPatchPvcNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pvc not found response has a 3xx status code
func (o *KubernetesPatchPvcNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pvc not found response has a 4xx status code
func (o *KubernetesPatchPvcNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch pvc not found response has a 5xx status code
func (o *KubernetesPatchPvcNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch pvc not found response a status code equal to that given
func (o *KubernetesPatchPvcNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchPvcNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchPvcNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchPvcNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KubernetesPatchPvcNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPvcInternalServerError creates a KubernetesPatchPvcInternalServerError with default headers values
func NewKubernetesPatchPvcInternalServerError() *KubernetesPatchPvcInternalServerError {
	return &KubernetesPatchPvcInternalServerError{}
}

/*
KubernetesPatchPvcInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchPvcInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch pvc internal server error response has a 2xx status code
func (o *KubernetesPatchPvcInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch pvc internal server error response has a 3xx status code
func (o *KubernetesPatchPvcInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch pvc internal server error response has a 4xx status code
func (o *KubernetesPatchPvcInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch pvc internal server error response has a 5xx status code
func (o *KubernetesPatchPvcInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch pvc internal server error response a status code equal to that given
func (o *KubernetesPatchPvcInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchPvcInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcInternalServerError ", 500)
}

func (o *KubernetesPatchPvcInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pvc][%d] kubernetesPatchPvcInternalServerError ", 500)
}

func (o *KubernetesPatchPvcInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
