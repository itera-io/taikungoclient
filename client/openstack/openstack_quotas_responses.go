// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpenstackQuotasReader is a Reader for the OpenstackQuotas structure.
type OpenstackQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackQuotasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackQuotasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackQuotasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackQuotasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOpenstackQuotasTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackQuotasOK creates a OpenstackQuotasOK with default headers values
func NewOpenstackQuotasOK() *OpenstackQuotasOK {
	return &OpenstackQuotasOK{}
}

/* OpenstackQuotasOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackQuotasOK struct {
	Payload *models.OpenstackQuotaList
}

func (o *OpenstackQuotasOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasOK  %+v", 200, o.Payload)
}
func (o *OpenstackQuotasOK) GetPayload() *models.OpenstackQuotaList {
	return o.Payload
}

func (o *OpenstackQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenstackQuotaList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackQuotasBadRequest creates a OpenstackQuotasBadRequest with default headers values
func NewOpenstackQuotasBadRequest() *OpenstackQuotasBadRequest {
	return &OpenstackQuotasBadRequest{}
}

/* OpenstackQuotasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackQuotasBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OpenstackQuotasBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasBadRequest  %+v", 400, o.Payload)
}
func (o *OpenstackQuotasBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpenstackQuotasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackQuotasUnauthorized creates a OpenstackQuotasUnauthorized with default headers values
func NewOpenstackQuotasUnauthorized() *OpenstackQuotasUnauthorized {
	return &OpenstackQuotasUnauthorized{}
}

/* OpenstackQuotasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackQuotasUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackQuotasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasUnauthorized  %+v", 401, o.Payload)
}
func (o *OpenstackQuotasUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackQuotasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackQuotasForbidden creates a OpenstackQuotasForbidden with default headers values
func NewOpenstackQuotasForbidden() *OpenstackQuotasForbidden {
	return &OpenstackQuotasForbidden{}
}

/* OpenstackQuotasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackQuotasForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackQuotasForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasForbidden  %+v", 403, o.Payload)
}
func (o *OpenstackQuotasForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackQuotasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackQuotasNotFound creates a OpenstackQuotasNotFound with default headers values
func NewOpenstackQuotasNotFound() *OpenstackQuotasNotFound {
	return &OpenstackQuotasNotFound{}
}

/* OpenstackQuotasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackQuotasNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackQuotasNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasNotFound  %+v", 404, o.Payload)
}
func (o *OpenstackQuotasNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackQuotasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackQuotasTooManyRequests creates a OpenstackQuotasTooManyRequests with default headers values
func NewOpenstackQuotasTooManyRequests() *OpenstackQuotasTooManyRequests {
	return &OpenstackQuotasTooManyRequests{}
}

/* OpenstackQuotasTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type OpenstackQuotasTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackQuotasTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasTooManyRequests  %+v", 429, o.Payload)
}
func (o *OpenstackQuotasTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackQuotasTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackQuotasInternalServerError creates a OpenstackQuotasInternalServerError with default headers values
func NewOpenstackQuotasInternalServerError() *OpenstackQuotasInternalServerError {
	return &OpenstackQuotasInternalServerError{}
}

/* OpenstackQuotasInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackQuotasInternalServerError struct {
}

func (o *OpenstackQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/quotas][%d] openstackQuotasInternalServerError ", 500)
}

func (o *OpenstackQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
