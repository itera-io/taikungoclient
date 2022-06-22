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

// KubernetesPatchSecretReader is a Reader for the KubernetesPatchSecret structure.
type KubernetesPatchSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesPatchSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesPatchSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesPatchSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesPatchSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesPatchSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesPatchSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesPatchSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesPatchSecretOK creates a KubernetesPatchSecretOK with default headers values
func NewKubernetesPatchSecretOK() *KubernetesPatchSecretOK {
	return &KubernetesPatchSecretOK{}
}

/* KubernetesPatchSecretOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesPatchSecretOK struct {
	Payload models.Unit
}

func (o *KubernetesPatchSecretOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/secret][%d] kubernetesPatchSecretOK  %+v", 200, o.Payload)
}
func (o *KubernetesPatchSecretOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KubernetesPatchSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchSecretBadRequest creates a KubernetesPatchSecretBadRequest with default headers values
func NewKubernetesPatchSecretBadRequest() *KubernetesPatchSecretBadRequest {
	return &KubernetesPatchSecretBadRequest{}
}

/* KubernetesPatchSecretBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesPatchSecretBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesPatchSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/secret][%d] kubernetesPatchSecretBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesPatchSecretBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchSecretUnauthorized creates a KubernetesPatchSecretUnauthorized with default headers values
func NewKubernetesPatchSecretUnauthorized() *KubernetesPatchSecretUnauthorized {
	return &KubernetesPatchSecretUnauthorized{}
}

/* KubernetesPatchSecretUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesPatchSecretUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesPatchSecretUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/secret][%d] kubernetesPatchSecretUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesPatchSecretUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchSecretForbidden creates a KubernetesPatchSecretForbidden with default headers values
func NewKubernetesPatchSecretForbidden() *KubernetesPatchSecretForbidden {
	return &KubernetesPatchSecretForbidden{}
}

/* KubernetesPatchSecretForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesPatchSecretForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesPatchSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/secret][%d] kubernetesPatchSecretForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesPatchSecretForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchSecretNotFound creates a KubernetesPatchSecretNotFound with default headers values
func NewKubernetesPatchSecretNotFound() *KubernetesPatchSecretNotFound {
	return &KubernetesPatchSecretNotFound{}
}

/* KubernetesPatchSecretNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesPatchSecretNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesPatchSecretNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/secret][%d] kubernetesPatchSecretNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesPatchSecretNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesPatchSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesPatchSecretInternalServerError creates a KubernetesPatchSecretInternalServerError with default headers values
func NewKubernetesPatchSecretInternalServerError() *KubernetesPatchSecretInternalServerError {
	return &KubernetesPatchSecretInternalServerError{}
}

/* KubernetesPatchSecretInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesPatchSecretInternalServerError struct {
}

func (o *KubernetesPatchSecretInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Kubernetes/patch/secret][%d] kubernetesPatchSecretInternalServerError ", 500)
}

func (o *KubernetesPatchSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}