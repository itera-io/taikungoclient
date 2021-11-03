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

// OpenstackRegionsReader is a Reader for the OpenstackRegions structure.
type OpenstackRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackRegionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewOpenstackRegionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackRegionsOK creates a OpenstackRegionsOK with default headers values
func NewOpenstackRegionsOK() *OpenstackRegionsOK {
	return &OpenstackRegionsOK{}
}

/* OpenstackRegionsOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackRegionsOK struct {
	Payload interface{}
}

func (o *OpenstackRegionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsOK  %+v", 200, o.Payload)
}
func (o *OpenstackRegionsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *OpenstackRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsBadRequest creates a OpenstackRegionsBadRequest with default headers values
func NewOpenstackRegionsBadRequest() *OpenstackRegionsBadRequest {
	return &OpenstackRegionsBadRequest{}
}

/* OpenstackRegionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackRegionsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OpenstackRegionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsBadRequest  %+v", 400, o.Payload)
}
func (o *OpenstackRegionsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpenstackRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsUnauthorized creates a OpenstackRegionsUnauthorized with default headers values
func NewOpenstackRegionsUnauthorized() *OpenstackRegionsUnauthorized {
	return &OpenstackRegionsUnauthorized{}
}

/* OpenstackRegionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackRegionsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsUnauthorized  %+v", 401, o.Payload)
}
func (o *OpenstackRegionsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsForbidden creates a OpenstackRegionsForbidden with default headers values
func NewOpenstackRegionsForbidden() *OpenstackRegionsForbidden {
	return &OpenstackRegionsForbidden{}
}

/* OpenstackRegionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackRegionsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackRegionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsForbidden  %+v", 403, o.Payload)
}
func (o *OpenstackRegionsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsNotFound creates a OpenstackRegionsNotFound with default headers values
func NewOpenstackRegionsNotFound() *OpenstackRegionsNotFound {
	return &OpenstackRegionsNotFound{}
}

/* OpenstackRegionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackRegionsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackRegionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsNotFound  %+v", 404, o.Payload)
}
func (o *OpenstackRegionsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackRegionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsTooManyRequests creates a OpenstackRegionsTooManyRequests with default headers values
func NewOpenstackRegionsTooManyRequests() *OpenstackRegionsTooManyRequests {
	return &OpenstackRegionsTooManyRequests{}
}

/* OpenstackRegionsTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type OpenstackRegionsTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *OpenstackRegionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsTooManyRequests  %+v", 429, o.Payload)
}
func (o *OpenstackRegionsTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackRegionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsInternalServerError creates a OpenstackRegionsInternalServerError with default headers values
func NewOpenstackRegionsInternalServerError() *OpenstackRegionsInternalServerError {
	return &OpenstackRegionsInternalServerError{}
}

/* OpenstackRegionsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackRegionsInternalServerError struct {
}

func (o *OpenstackRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsInternalServerError ", 500)
}

func (o *OpenstackRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
