// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OpenstackZoneListReader is a Reader for the OpenstackZoneList structure.
type OpenstackZoneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackZoneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackZoneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackZoneListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackZoneListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackZoneListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackZoneListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackZoneListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackZoneListOK creates a OpenstackZoneListOK with default headers values
func NewOpenstackZoneListOK() *OpenstackZoneListOK {
	return &OpenstackZoneListOK{}
}

/*
OpenstackZoneListOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackZoneListOK struct {
	Payload []string
}

// IsSuccess returns true when this openstack zone list o k response has a 2xx status code
func (o *OpenstackZoneListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this openstack zone list o k response has a 3xx status code
func (o *OpenstackZoneListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack zone list o k response has a 4xx status code
func (o *OpenstackZoneListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack zone list o k response has a 5xx status code
func (o *OpenstackZoneListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack zone list o k response a status code equal to that given
func (o *OpenstackZoneListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the openstack zone list o k response
func (o *OpenstackZoneListOK) Code() int {
	return 200
}

func (o *OpenstackZoneListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListOK  %+v", 200, o.Payload)
}

func (o *OpenstackZoneListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListOK  %+v", 200, o.Payload)
}

func (o *OpenstackZoneListOK) GetPayload() []string {
	return o.Payload
}

func (o *OpenstackZoneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackZoneListBadRequest creates a OpenstackZoneListBadRequest with default headers values
func NewOpenstackZoneListBadRequest() *OpenstackZoneListBadRequest {
	return &OpenstackZoneListBadRequest{}
}

/*
OpenstackZoneListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackZoneListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this openstack zone list bad request response has a 2xx status code
func (o *OpenstackZoneListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack zone list bad request response has a 3xx status code
func (o *OpenstackZoneListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack zone list bad request response has a 4xx status code
func (o *OpenstackZoneListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack zone list bad request response has a 5xx status code
func (o *OpenstackZoneListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack zone list bad request response a status code equal to that given
func (o *OpenstackZoneListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the openstack zone list bad request response
func (o *OpenstackZoneListBadRequest) Code() int {
	return 400
}

func (o *OpenstackZoneListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackZoneListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackZoneListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *OpenstackZoneListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackZoneListUnauthorized creates a OpenstackZoneListUnauthorized with default headers values
func NewOpenstackZoneListUnauthorized() *OpenstackZoneListUnauthorized {
	return &OpenstackZoneListUnauthorized{}
}

/*
OpenstackZoneListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackZoneListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this openstack zone list unauthorized response has a 2xx status code
func (o *OpenstackZoneListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack zone list unauthorized response has a 3xx status code
func (o *OpenstackZoneListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack zone list unauthorized response has a 4xx status code
func (o *OpenstackZoneListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack zone list unauthorized response has a 5xx status code
func (o *OpenstackZoneListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack zone list unauthorized response a status code equal to that given
func (o *OpenstackZoneListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the openstack zone list unauthorized response
func (o *OpenstackZoneListUnauthorized) Code() int {
	return 401
}

func (o *OpenstackZoneListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackZoneListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackZoneListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *OpenstackZoneListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackZoneListForbidden creates a OpenstackZoneListForbidden with default headers values
func NewOpenstackZoneListForbidden() *OpenstackZoneListForbidden {
	return &OpenstackZoneListForbidden{}
}

/*
OpenstackZoneListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackZoneListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this openstack zone list forbidden response has a 2xx status code
func (o *OpenstackZoneListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack zone list forbidden response has a 3xx status code
func (o *OpenstackZoneListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack zone list forbidden response has a 4xx status code
func (o *OpenstackZoneListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack zone list forbidden response has a 5xx status code
func (o *OpenstackZoneListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack zone list forbidden response a status code equal to that given
func (o *OpenstackZoneListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the openstack zone list forbidden response
func (o *OpenstackZoneListForbidden) Code() int {
	return 403
}

func (o *OpenstackZoneListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackZoneListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackZoneListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *OpenstackZoneListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackZoneListNotFound creates a OpenstackZoneListNotFound with default headers values
func NewOpenstackZoneListNotFound() *OpenstackZoneListNotFound {
	return &OpenstackZoneListNotFound{}
}

/*
OpenstackZoneListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackZoneListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this openstack zone list not found response has a 2xx status code
func (o *OpenstackZoneListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack zone list not found response has a 3xx status code
func (o *OpenstackZoneListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack zone list not found response has a 4xx status code
func (o *OpenstackZoneListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack zone list not found response has a 5xx status code
func (o *OpenstackZoneListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack zone list not found response a status code equal to that given
func (o *OpenstackZoneListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the openstack zone list not found response
func (o *OpenstackZoneListNotFound) Code() int {
	return 404
}

func (o *OpenstackZoneListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackZoneListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackZoneListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *OpenstackZoneListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackZoneListInternalServerError creates a OpenstackZoneListInternalServerError with default headers values
func NewOpenstackZoneListInternalServerError() *OpenstackZoneListInternalServerError {
	return &OpenstackZoneListInternalServerError{}
}

/*
OpenstackZoneListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackZoneListInternalServerError struct {
}

// IsSuccess returns true when this openstack zone list internal server error response has a 2xx status code
func (o *OpenstackZoneListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack zone list internal server error response has a 3xx status code
func (o *OpenstackZoneListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack zone list internal server error response has a 4xx status code
func (o *OpenstackZoneListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack zone list internal server error response has a 5xx status code
func (o *OpenstackZoneListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this openstack zone list internal server error response a status code equal to that given
func (o *OpenstackZoneListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the openstack zone list internal server error response
func (o *OpenstackZoneListInternalServerError) Code() int {
	return 500
}

func (o *OpenstackZoneListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListInternalServerError ", 500)
}

func (o *OpenstackZoneListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/zones][%d] openstackZoneListInternalServerError ", 500)
}

func (o *OpenstackZoneListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
