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

// KubernetesPatchPdbReader is a Reader for the KubernetesPatchPdb structure.
type KubernetesPatchPdbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchPdbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchPdbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchPdbBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchPdbUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchPdbForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchPdbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchPdbInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchPdbOK creates a KubernetesPatchPdbOK with default headers values
func NewKubernetesPatchPdbOK() *KubernetesPatchPdbOK {
	return &KubernetesPatchPdbOK{}
}

/* KubernetesPatchPdbOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchPdbOK struct {
	Payload models.Unit
}

func (o *KubernetesPatchPdbOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pdb][%d] kubernetesPatchPdbOK  %+v", 200, o.Payload)
}
func (o *KubernetesPatchPdbOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchPdbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPdbBadRequest creates a KubernetesPatchPdbBadRequest with default headers values
func NewKubernetesPatchPdbBadRequest() *KubernetesPatchPdbBadRequest {
	return &KubernetesPatchPdbBadRequest{}
}

/* KubernetesPatchPdbBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchPdbBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesPatchPdbBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pdb][%d] kubernetesPatchPdbBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesPatchPdbBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPdbBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPdbUnauthorized creates a KubernetesPatchPdbUnauthorized with default headers values
func NewKubernetesPatchPdbUnauthorized() *KubernetesPatchPdbUnauthorized {
	return &KubernetesPatchPdbUnauthorized{}
}

/* KubernetesPatchPdbUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchPdbUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesPatchPdbUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pdb][%d] kubernetesPatchPdbUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesPatchPdbUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPdbUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPdbForbidden creates a KubernetesPatchPdbForbidden with default headers values
func NewKubernetesPatchPdbForbidden() *KubernetesPatchPdbForbidden {
	return &KubernetesPatchPdbForbidden{}
}

/* KubernetesPatchPdbForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchPdbForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesPatchPdbForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pdb][%d] kubernetesPatchPdbForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesPatchPdbForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPdbForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPdbNotFound creates a KubernetesPatchPdbNotFound with default headers values
func NewKubernetesPatchPdbNotFound() *KubernetesPatchPdbNotFound {
	return &KubernetesPatchPdbNotFound{}
}

/* KubernetesPatchPdbNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchPdbNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesPatchPdbNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pdb][%d] kubernetesPatchPdbNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesPatchPdbNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchPdbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchPdbInternalServerError creates a KubernetesPatchPdbInternalServerError with default headers values
func NewKubernetesPatchPdbInternalServerError() *KubernetesPatchPdbInternalServerError {
	return &KubernetesPatchPdbInternalServerError{}
}

/* KubernetesPatchPdbInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchPdbInternalServerError struct {
}

func (o *KubernetesPatchPdbInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/pdb][%d] kubernetesPatchPdbInternalServerError ", 500)
}

func (o *KubernetesPatchPdbInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
