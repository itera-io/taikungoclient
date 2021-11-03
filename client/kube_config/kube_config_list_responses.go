// Code generated by go-swagger; DO NOT EDIT.

package kube_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubeConfigListReader is a Reader for the KubeConfigList structure.
type KubeConfigListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubeConfigListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubeConfigListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubeConfigListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubeConfigListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubeConfigListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubeConfigListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewKubeConfigListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubeConfigListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubeConfigListOK creates a KubeConfigListOK with default headers values
func NewKubeConfigListOK() *KubeConfigListOK {
	return &KubeConfigListOK{}
}

/* KubeConfigListOK describes a response with status code 200, with default header values.

Success
*/
type KubeConfigListOK struct {
	Payload *models.KubeConfigForUserList
}

func (o *KubeConfigListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListOK  %+v", 200, o.Payload)
}
func (o *KubeConfigListOK) GetPayload() *models.KubeConfigForUserList {
	return o.Payload
}

func (o *KubeConfigListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubeConfigForUserList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigListBadRequest creates a KubeConfigListBadRequest with default headers values
func NewKubeConfigListBadRequest() *KubeConfigListBadRequest {
	return &KubeConfigListBadRequest{}
}

/* KubeConfigListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubeConfigListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubeConfigListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListBadRequest  %+v", 400, o.Payload)
}
func (o *KubeConfigListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubeConfigListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigListUnauthorized creates a KubeConfigListUnauthorized with default headers values
func NewKubeConfigListUnauthorized() *KubeConfigListUnauthorized {
	return &KubeConfigListUnauthorized{}
}

/* KubeConfigListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubeConfigListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListUnauthorized  %+v", 401, o.Payload)
}
func (o *KubeConfigListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigListForbidden creates a KubeConfigListForbidden with default headers values
func NewKubeConfigListForbidden() *KubeConfigListForbidden {
	return &KubeConfigListForbidden{}
}

/* KubeConfigListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubeConfigListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListForbidden  %+v", 403, o.Payload)
}
func (o *KubeConfigListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigListNotFound creates a KubeConfigListNotFound with default headers values
func NewKubeConfigListNotFound() *KubeConfigListNotFound {
	return &KubeConfigListNotFound{}
}

/* KubeConfigListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubeConfigListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListNotFound  %+v", 404, o.Payload)
}
func (o *KubeConfigListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigListTooManyRequests creates a KubeConfigListTooManyRequests with default headers values
func NewKubeConfigListTooManyRequests() *KubeConfigListTooManyRequests {
	return &KubeConfigListTooManyRequests{}
}

/* KubeConfigListTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type KubeConfigListTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *KubeConfigListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListTooManyRequests  %+v", 429, o.Payload)
}
func (o *KubeConfigListTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubeConfigListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubeConfigListInternalServerError creates a KubeConfigListInternalServerError with default headers values
func NewKubeConfigListInternalServerError() *KubeConfigListInternalServerError {
	return &KubeConfigListInternalServerError{}
}

/* KubeConfigListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubeConfigListInternalServerError struct {
}

func (o *KubeConfigListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubeConfig][%d] kubeConfigListInternalServerError ", 500)
}

func (o *KubeConfigListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
