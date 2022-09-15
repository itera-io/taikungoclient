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

// KubernetesPatchNodeReader is a Reader for the KubernetesPatchNode structure.
type KubernetesPatchNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchNodeOK creates a KubernetesPatchNodeOK with default headers values
func NewKubernetesPatchNodeOK() *KubernetesPatchNodeOK {
	return &KubernetesPatchNodeOK{}
}

/*
KubernetesPatchNodeOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchNodeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this kubernetes patch node o k response has a 2xx status code
func (o *KubernetesPatchNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this kubernetes patch node o k response has a 3xx status code
func (o *KubernetesPatchNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch node o k response has a 4xx status code
func (o *KubernetesPatchNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch node o k response has a 5xx status code
func (o *KubernetesPatchNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch node o k response a status code equal to that given
func (o *KubernetesPatchNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *KubernetesPatchNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchNodeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeOK  %+v", 200, o.Payload)
}

func (o *KubernetesPatchNodeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchNodeBadRequest creates a KubernetesPatchNodeBadRequest with default headers values
func NewKubernetesPatchNodeBadRequest() *KubernetesPatchNodeBadRequest {
	return &KubernetesPatchNodeBadRequest{}
}

/*
KubernetesPatchNodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchNodeBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this kubernetes patch node bad request response has a 2xx status code
func (o *KubernetesPatchNodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch node bad request response has a 3xx status code
func (o *KubernetesPatchNodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch node bad request response has a 4xx status code
func (o *KubernetesPatchNodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch node bad request response has a 5xx status code
func (o *KubernetesPatchNodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch node bad request response a status code equal to that given
func (o *KubernetesPatchNodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KubernetesPatchNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchNodeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeBadRequest  %+v", 400, o.Payload)
}

func (o *KubernetesPatchNodeBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchNodeUnauthorized creates a KubernetesPatchNodeUnauthorized with default headers values
func NewKubernetesPatchNodeUnauthorized() *KubernetesPatchNodeUnauthorized {
	return &KubernetesPatchNodeUnauthorized{}
}

/*
KubernetesPatchNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchNodeUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch node unauthorized response has a 2xx status code
func (o *KubernetesPatchNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch node unauthorized response has a 3xx status code
func (o *KubernetesPatchNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch node unauthorized response has a 4xx status code
func (o *KubernetesPatchNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch node unauthorized response has a 5xx status code
func (o *KubernetesPatchNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch node unauthorized response a status code equal to that given
func (o *KubernetesPatchNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KubernetesPatchNodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchNodeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *KubernetesPatchNodeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchNodeForbidden creates a KubernetesPatchNodeForbidden with default headers values
func NewKubernetesPatchNodeForbidden() *KubernetesPatchNodeForbidden {
	return &KubernetesPatchNodeForbidden{}
}

/*
KubernetesPatchNodeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchNodeForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch node forbidden response has a 2xx status code
func (o *KubernetesPatchNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch node forbidden response has a 3xx status code
func (o *KubernetesPatchNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch node forbidden response has a 4xx status code
func (o *KubernetesPatchNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch node forbidden response has a 5xx status code
func (o *KubernetesPatchNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch node forbidden response a status code equal to that given
func (o *KubernetesPatchNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KubernetesPatchNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchNodeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeForbidden  %+v", 403, o.Payload)
}

func (o *KubernetesPatchNodeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchNodeNotFound creates a KubernetesPatchNodeNotFound with default headers values
func NewKubernetesPatchNodeNotFound() *KubernetesPatchNodeNotFound {
	return &KubernetesPatchNodeNotFound{}
}

/*
KubernetesPatchNodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchNodeNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this kubernetes patch node not found response has a 2xx status code
func (o *KubernetesPatchNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch node not found response has a 3xx status code
func (o *KubernetesPatchNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch node not found response has a 4xx status code
func (o *KubernetesPatchNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this kubernetes patch node not found response has a 5xx status code
func (o *KubernetesPatchNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this kubernetes patch node not found response a status code equal to that given
func (o *KubernetesPatchNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KubernetesPatchNodeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchNodeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeNotFound  %+v", 404, o.Payload)
}

func (o *KubernetesPatchNodeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchNodeInternalServerError creates a KubernetesPatchNodeInternalServerError with default headers values
func NewKubernetesPatchNodeInternalServerError() *KubernetesPatchNodeInternalServerError {
	return &KubernetesPatchNodeInternalServerError{}
}

/*
KubernetesPatchNodeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchNodeInternalServerError struct {
}

// IsSuccess returns true when this kubernetes patch node internal server error response has a 2xx status code
func (o *KubernetesPatchNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this kubernetes patch node internal server error response has a 3xx status code
func (o *KubernetesPatchNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this kubernetes patch node internal server error response has a 4xx status code
func (o *KubernetesPatchNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this kubernetes patch node internal server error response has a 5xx status code
func (o *KubernetesPatchNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this kubernetes patch node internal server error response a status code equal to that given
func (o *KubernetesPatchNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KubernetesPatchNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeInternalServerError ", 500)
}

func (o *KubernetesPatchNodeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/node][%d] kubernetesPatchNodeInternalServerError ", 500)
}

func (o *KubernetesPatchNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
