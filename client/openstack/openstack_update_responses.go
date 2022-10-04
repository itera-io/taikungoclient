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

// OpenstackUpdateReader is a Reader for the OpenstackUpdate structure.
type OpenstackUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackUpdateOK creates a OpenstackUpdateOK with default headers values
func NewOpenstackUpdateOK() *OpenstackUpdateOK {
	return &OpenstackUpdateOK{}
}

/*
OpenstackUpdateOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this openstack update o k response has a 2xx status code
func (o *OpenstackUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this openstack update o k response has a 3xx status code
func (o *OpenstackUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack update o k response has a 4xx status code
func (o *OpenstackUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack update o k response has a 5xx status code
func (o *OpenstackUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack update o k response a status code equal to that given
func (o *OpenstackUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpenstackUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateOK  %+v", 200, o.Payload)
}

func (o *OpenstackUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateOK  %+v", 200, o.Payload)
}

func (o *OpenstackUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpenstackUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackUpdateBadRequest creates a OpenstackUpdateBadRequest with default headers values
func NewOpenstackUpdateBadRequest() *OpenstackUpdateBadRequest {
	return &OpenstackUpdateBadRequest{}
}

/*
OpenstackUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this openstack update bad request response has a 2xx status code
func (o *OpenstackUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack update bad request response has a 3xx status code
func (o *OpenstackUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack update bad request response has a 4xx status code
func (o *OpenstackUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack update bad request response has a 5xx status code
func (o *OpenstackUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack update bad request response a status code equal to that given
func (o *OpenstackUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpenstackUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpenstackUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackUpdateUnauthorized creates a OpenstackUpdateUnauthorized with default headers values
func NewOpenstackUpdateUnauthorized() *OpenstackUpdateUnauthorized {
	return &OpenstackUpdateUnauthorized{}
}

/*
OpenstackUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackUpdateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this openstack update unauthorized response has a 2xx status code
func (o *OpenstackUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack update unauthorized response has a 3xx status code
func (o *OpenstackUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack update unauthorized response has a 4xx status code
func (o *OpenstackUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack update unauthorized response has a 5xx status code
func (o *OpenstackUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack update unauthorized response a status code equal to that given
func (o *OpenstackUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpenstackUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackUpdateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpenstackUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackUpdateForbidden creates a OpenstackUpdateForbidden with default headers values
func NewOpenstackUpdateForbidden() *OpenstackUpdateForbidden {
	return &OpenstackUpdateForbidden{}
}

/*
OpenstackUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackUpdateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this openstack update forbidden response has a 2xx status code
func (o *OpenstackUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack update forbidden response has a 3xx status code
func (o *OpenstackUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack update forbidden response has a 4xx status code
func (o *OpenstackUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack update forbidden response has a 5xx status code
func (o *OpenstackUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack update forbidden response a status code equal to that given
func (o *OpenstackUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpenstackUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackUpdateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpenstackUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackUpdateNotFound creates a OpenstackUpdateNotFound with default headers values
func NewOpenstackUpdateNotFound() *OpenstackUpdateNotFound {
	return &OpenstackUpdateNotFound{}
}

/*
OpenstackUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackUpdateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this openstack update not found response has a 2xx status code
func (o *OpenstackUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack update not found response has a 3xx status code
func (o *OpenstackUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack update not found response has a 4xx status code
func (o *OpenstackUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack update not found response has a 5xx status code
func (o *OpenstackUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack update not found response a status code equal to that given
func (o *OpenstackUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpenstackUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackUpdateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpenstackUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackUpdateInternalServerError creates a OpenstackUpdateInternalServerError with default headers values
func NewOpenstackUpdateInternalServerError() *OpenstackUpdateInternalServerError {
	return &OpenstackUpdateInternalServerError{}
}

/*
OpenstackUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackUpdateInternalServerError struct {
}

// IsSuccess returns true when this openstack update internal server error response has a 2xx status code
func (o *OpenstackUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack update internal server error response has a 3xx status code
func (o *OpenstackUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack update internal server error response has a 4xx status code
func (o *OpenstackUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack update internal server error response has a 5xx status code
func (o *OpenstackUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this openstack update internal server error response a status code equal to that given
func (o *OpenstackUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpenstackUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateInternalServerError ", 500)
}

func (o *OpenstackUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/update][%d] openstackUpdateInternalServerError ", 500)
}

func (o *OpenstackUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
