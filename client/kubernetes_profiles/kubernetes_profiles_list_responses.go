// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KubernetesProfilesListReader is a Reader for the KubernetesProfilesList structure.
type KubernetesProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKubernetesProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKubernetesProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKubernetesProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKubernetesProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKubernetesProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKubernetesProfilesListOK creates a KubernetesProfilesListOK with default headers values
func NewKubernetesProfilesListOK() *KubernetesProfilesListOK {
	return &KubernetesProfilesListOK{}
}

/* KubernetesProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type KubernetesProfilesListOK struct {
	Payload *models.KubernetesProfilesList
}

func (o *KubernetesProfilesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListOK  %+v", 200, o.Payload)
}
func (o *KubernetesProfilesListOK) GetPayload() *models.KubernetesProfilesList {
	return o.Payload
}

func (o *KubernetesProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesProfilesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListBadRequest creates a KubernetesProfilesListBadRequest with default headers values
func NewKubernetesProfilesListBadRequest() *KubernetesProfilesListBadRequest {
	return &KubernetesProfilesListBadRequest{}
}

/* KubernetesProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KubernetesProfilesListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KubernetesProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListBadRequest  %+v", 400, o.Payload)
}
func (o *KubernetesProfilesListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListUnauthorized creates a KubernetesProfilesListUnauthorized with default headers values
func NewKubernetesProfilesListUnauthorized() *KubernetesProfilesListUnauthorized {
	return &KubernetesProfilesListUnauthorized{}
}

/* KubernetesProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KubernetesProfilesListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListUnauthorized  %+v", 401, o.Payload)
}
func (o *KubernetesProfilesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListForbidden creates a KubernetesProfilesListForbidden with default headers values
func NewKubernetesProfilesListForbidden() *KubernetesProfilesListForbidden {
	return &KubernetesProfilesListForbidden{}
}

/* KubernetesProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KubernetesProfilesListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesProfilesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListForbidden  %+v", 403, o.Payload)
}
func (o *KubernetesProfilesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListNotFound creates a KubernetesProfilesListNotFound with default headers values
func NewKubernetesProfilesListNotFound() *KubernetesProfilesListNotFound {
	return &KubernetesProfilesListNotFound{}
}

/* KubernetesProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KubernetesProfilesListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KubernetesProfilesListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListNotFound  %+v", 404, o.Payload)
}
func (o *KubernetesProfilesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KubernetesProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesProfilesListInternalServerError creates a KubernetesProfilesListInternalServerError with default headers values
func NewKubernetesProfilesListInternalServerError() *KubernetesProfilesListInternalServerError {
	return &KubernetesProfilesListInternalServerError{}
}

/* KubernetesProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KubernetesProfilesListInternalServerError struct {
}

func (o *KubernetesProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/KubernetesProfiles/list][%d] kubernetesProfilesListInternalServerError ", 500)
}

func (o *KubernetesProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
