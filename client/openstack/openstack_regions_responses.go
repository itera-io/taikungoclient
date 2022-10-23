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

/*
OpenstackRegionsOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackRegionsOK struct {
	Payload []string
}

// IsSuccess returns true when this openstack regions o k response has a 2xx status code
func (o *OpenstackRegionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this openstack regions o k response has a 3xx status code
func (o *OpenstackRegionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack regions o k response has a 4xx status code
func (o *OpenstackRegionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack regions o k response has a 5xx status code
func (o *OpenstackRegionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack regions o k response a status code equal to that given
func (o *OpenstackRegionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpenstackRegionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsOK  %+v", 200, o.Payload)
}

func (o *OpenstackRegionsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsOK  %+v", 200, o.Payload)
}

func (o *OpenstackRegionsOK) GetPayload() []string {
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

/*
OpenstackRegionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackRegionsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this openstack regions bad request response has a 2xx status code
func (o *OpenstackRegionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack regions bad request response has a 3xx status code
func (o *OpenstackRegionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack regions bad request response has a 4xx status code
func (o *OpenstackRegionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack regions bad request response has a 5xx status code
func (o *OpenstackRegionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack regions bad request response a status code equal to that given
func (o *OpenstackRegionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpenstackRegionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackRegionsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackRegionsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *OpenstackRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackRegionsUnauthorized creates a OpenstackRegionsUnauthorized with default headers values
func NewOpenstackRegionsUnauthorized() *OpenstackRegionsUnauthorized {
	return &OpenstackRegionsUnauthorized{}
}

/*
OpenstackRegionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackRegionsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack regions unauthorized response has a 2xx status code
func (o *OpenstackRegionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack regions unauthorized response has a 3xx status code
func (o *OpenstackRegionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack regions unauthorized response has a 4xx status code
func (o *OpenstackRegionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack regions unauthorized response has a 5xx status code
func (o *OpenstackRegionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack regions unauthorized response a status code equal to that given
func (o *OpenstackRegionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpenstackRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackRegionsUnauthorized) String() string {
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

/*
OpenstackRegionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackRegionsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack regions forbidden response has a 2xx status code
func (o *OpenstackRegionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack regions forbidden response has a 3xx status code
func (o *OpenstackRegionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack regions forbidden response has a 4xx status code
func (o *OpenstackRegionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack regions forbidden response has a 5xx status code
func (o *OpenstackRegionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack regions forbidden response a status code equal to that given
func (o *OpenstackRegionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpenstackRegionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackRegionsForbidden) String() string {
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

/*
OpenstackRegionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackRegionsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack regions not found response has a 2xx status code
func (o *OpenstackRegionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack regions not found response has a 3xx status code
func (o *OpenstackRegionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack regions not found response has a 4xx status code
func (o *OpenstackRegionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack regions not found response has a 5xx status code
func (o *OpenstackRegionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack regions not found response a status code equal to that given
func (o *OpenstackRegionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpenstackRegionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackRegionsNotFound) String() string {
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

// NewOpenstackRegionsInternalServerError creates a OpenstackRegionsInternalServerError with default headers values
func NewOpenstackRegionsInternalServerError() *OpenstackRegionsInternalServerError {
	return &OpenstackRegionsInternalServerError{}
}

/*
OpenstackRegionsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackRegionsInternalServerError struct {
}

// IsSuccess returns true when this openstack regions internal server error response has a 2xx status code
func (o *OpenstackRegionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack regions internal server error response has a 3xx status code
func (o *OpenstackRegionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack regions internal server error response has a 4xx status code
func (o *OpenstackRegionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack regions internal server error response has a 5xx status code
func (o *OpenstackRegionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this openstack regions internal server error response a status code equal to that given
func (o *OpenstackRegionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpenstackRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsInternalServerError ", 500)
}

func (o *OpenstackRegionsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/regions][%d] openstackRegionsInternalServerError ", 500)
}

func (o *OpenstackRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
