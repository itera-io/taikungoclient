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

// OpenstackNetworksReader is a Reader for the OpenstackNetworks structure.
type OpenstackNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackNetworksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackNetworksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackNetworksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackNetworksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackNetworksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackNetworksOK creates a OpenstackNetworksOK with default headers values
func NewOpenstackNetworksOK() *OpenstackNetworksOK {
	return &OpenstackNetworksOK{}
}

/*
OpenstackNetworksOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackNetworksOK struct {
	Payload []*models.CommonStringBasedDropdownDto
}

// IsSuccess returns true when this openstack networks o k response has a 2xx status code
func (o *OpenstackNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this openstack networks o k response has a 3xx status code
func (o *OpenstackNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack networks o k response has a 4xx status code
func (o *OpenstackNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack networks o k response has a 5xx status code
func (o *OpenstackNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack networks o k response a status code equal to that given
func (o *OpenstackNetworksOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpenstackNetworksOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksOK  %+v", 200, o.Payload)
}

func (o *OpenstackNetworksOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksOK  %+v", 200, o.Payload)
}

func (o *OpenstackNetworksOK) GetPayload() []*models.CommonStringBasedDropdownDto {
	return o.Payload
}

func (o *OpenstackNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackNetworksBadRequest creates a OpenstackNetworksBadRequest with default headers values
func NewOpenstackNetworksBadRequest() *OpenstackNetworksBadRequest {
	return &OpenstackNetworksBadRequest{}
}

/*
OpenstackNetworksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackNetworksBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this openstack networks bad request response has a 2xx status code
func (o *OpenstackNetworksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack networks bad request response has a 3xx status code
func (o *OpenstackNetworksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack networks bad request response has a 4xx status code
func (o *OpenstackNetworksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack networks bad request response has a 5xx status code
func (o *OpenstackNetworksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack networks bad request response a status code equal to that given
func (o *OpenstackNetworksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpenstackNetworksBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackNetworksBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackNetworksBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpenstackNetworksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackNetworksUnauthorized creates a OpenstackNetworksUnauthorized with default headers values
func NewOpenstackNetworksUnauthorized() *OpenstackNetworksUnauthorized {
	return &OpenstackNetworksUnauthorized{}
}

/*
OpenstackNetworksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackNetworksUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this openstack networks unauthorized response has a 2xx status code
func (o *OpenstackNetworksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack networks unauthorized response has a 3xx status code
func (o *OpenstackNetworksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack networks unauthorized response has a 4xx status code
func (o *OpenstackNetworksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack networks unauthorized response has a 5xx status code
func (o *OpenstackNetworksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack networks unauthorized response a status code equal to that given
func (o *OpenstackNetworksUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpenstackNetworksUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackNetworksUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackNetworksUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpenstackNetworksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackNetworksForbidden creates a OpenstackNetworksForbidden with default headers values
func NewOpenstackNetworksForbidden() *OpenstackNetworksForbidden {
	return &OpenstackNetworksForbidden{}
}

/*
OpenstackNetworksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackNetworksForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this openstack networks forbidden response has a 2xx status code
func (o *OpenstackNetworksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack networks forbidden response has a 3xx status code
func (o *OpenstackNetworksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack networks forbidden response has a 4xx status code
func (o *OpenstackNetworksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack networks forbidden response has a 5xx status code
func (o *OpenstackNetworksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack networks forbidden response a status code equal to that given
func (o *OpenstackNetworksForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpenstackNetworksForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackNetworksForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackNetworksForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpenstackNetworksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackNetworksNotFound creates a OpenstackNetworksNotFound with default headers values
func NewOpenstackNetworksNotFound() *OpenstackNetworksNotFound {
	return &OpenstackNetworksNotFound{}
}

/*
OpenstackNetworksNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackNetworksNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this openstack networks not found response has a 2xx status code
func (o *OpenstackNetworksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack networks not found response has a 3xx status code
func (o *OpenstackNetworksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack networks not found response has a 4xx status code
func (o *OpenstackNetworksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack networks not found response has a 5xx status code
func (o *OpenstackNetworksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack networks not found response a status code equal to that given
func (o *OpenstackNetworksNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpenstackNetworksNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackNetworksNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackNetworksNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpenstackNetworksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackNetworksInternalServerError creates a OpenstackNetworksInternalServerError with default headers values
func NewOpenstackNetworksInternalServerError() *OpenstackNetworksInternalServerError {
	return &OpenstackNetworksInternalServerError{}
}

/*
OpenstackNetworksInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackNetworksInternalServerError struct {
}

// IsSuccess returns true when this openstack networks internal server error response has a 2xx status code
func (o *OpenstackNetworksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack networks internal server error response has a 3xx status code
func (o *OpenstackNetworksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack networks internal server error response has a 4xx status code
func (o *OpenstackNetworksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack networks internal server error response has a 5xx status code
func (o *OpenstackNetworksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this openstack networks internal server error response a status code equal to that given
func (o *OpenstackNetworksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpenstackNetworksInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksInternalServerError ", 500)
}

func (o *OpenstackNetworksInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/networks][%d] openstackNetworksInternalServerError ", 500)
}

func (o *OpenstackNetworksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
