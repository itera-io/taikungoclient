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

// OpenstackVolumeTypesReader is a Reader for the OpenstackVolumeTypes structure.
type OpenstackVolumeTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackVolumeTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackVolumeTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackVolumeTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackVolumeTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackVolumeTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackVolumeTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackVolumeTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackVolumeTypesOK creates a OpenstackVolumeTypesOK with default headers values
func NewOpenstackVolumeTypesOK() *OpenstackVolumeTypesOK {
	return &OpenstackVolumeTypesOK{}
}

/*
OpenstackVolumeTypesOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackVolumeTypesOK struct {
	Payload []string
}

// IsSuccess returns true when this openstack volume types o k response has a 2xx status code
func (o *OpenstackVolumeTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this openstack volume types o k response has a 3xx status code
func (o *OpenstackVolumeTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack volume types o k response has a 4xx status code
func (o *OpenstackVolumeTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack volume types o k response has a 5xx status code
func (o *OpenstackVolumeTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack volume types o k response a status code equal to that given
func (o *OpenstackVolumeTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpenstackVolumeTypesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesOK  %+v", 200, o.Payload)
}

func (o *OpenstackVolumeTypesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesOK  %+v", 200, o.Payload)
}

func (o *OpenstackVolumeTypesOK) GetPayload() []string {
	return o.Payload
}

func (o *OpenstackVolumeTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackVolumeTypesBadRequest creates a OpenstackVolumeTypesBadRequest with default headers values
func NewOpenstackVolumeTypesBadRequest() *OpenstackVolumeTypesBadRequest {
	return &OpenstackVolumeTypesBadRequest{}
}

/*
OpenstackVolumeTypesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackVolumeTypesBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack volume types bad request response has a 2xx status code
func (o *OpenstackVolumeTypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack volume types bad request response has a 3xx status code
func (o *OpenstackVolumeTypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack volume types bad request response has a 4xx status code
func (o *OpenstackVolumeTypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack volume types bad request response has a 5xx status code
func (o *OpenstackVolumeTypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack volume types bad request response a status code equal to that given
func (o *OpenstackVolumeTypesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpenstackVolumeTypesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackVolumeTypesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackVolumeTypesBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackVolumeTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackVolumeTypesUnauthorized creates a OpenstackVolumeTypesUnauthorized with default headers values
func NewOpenstackVolumeTypesUnauthorized() *OpenstackVolumeTypesUnauthorized {
	return &OpenstackVolumeTypesUnauthorized{}
}

/*
OpenstackVolumeTypesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackVolumeTypesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack volume types unauthorized response has a 2xx status code
func (o *OpenstackVolumeTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack volume types unauthorized response has a 3xx status code
func (o *OpenstackVolumeTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack volume types unauthorized response has a 4xx status code
func (o *OpenstackVolumeTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack volume types unauthorized response has a 5xx status code
func (o *OpenstackVolumeTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack volume types unauthorized response a status code equal to that given
func (o *OpenstackVolumeTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpenstackVolumeTypesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackVolumeTypesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackVolumeTypesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackVolumeTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackVolumeTypesForbidden creates a OpenstackVolumeTypesForbidden with default headers values
func NewOpenstackVolumeTypesForbidden() *OpenstackVolumeTypesForbidden {
	return &OpenstackVolumeTypesForbidden{}
}

/*
OpenstackVolumeTypesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackVolumeTypesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack volume types forbidden response has a 2xx status code
func (o *OpenstackVolumeTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack volume types forbidden response has a 3xx status code
func (o *OpenstackVolumeTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack volume types forbidden response has a 4xx status code
func (o *OpenstackVolumeTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack volume types forbidden response has a 5xx status code
func (o *OpenstackVolumeTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack volume types forbidden response a status code equal to that given
func (o *OpenstackVolumeTypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpenstackVolumeTypesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackVolumeTypesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackVolumeTypesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackVolumeTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackVolumeTypesNotFound creates a OpenstackVolumeTypesNotFound with default headers values
func NewOpenstackVolumeTypesNotFound() *OpenstackVolumeTypesNotFound {
	return &OpenstackVolumeTypesNotFound{}
}

/*
OpenstackVolumeTypesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackVolumeTypesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack volume types not found response has a 2xx status code
func (o *OpenstackVolumeTypesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack volume types not found response has a 3xx status code
func (o *OpenstackVolumeTypesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack volume types not found response has a 4xx status code
func (o *OpenstackVolumeTypesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack volume types not found response has a 5xx status code
func (o *OpenstackVolumeTypesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack volume types not found response a status code equal to that given
func (o *OpenstackVolumeTypesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpenstackVolumeTypesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackVolumeTypesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackVolumeTypesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackVolumeTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackVolumeTypesInternalServerError creates a OpenstackVolumeTypesInternalServerError with default headers values
func NewOpenstackVolumeTypesInternalServerError() *OpenstackVolumeTypesInternalServerError {
	return &OpenstackVolumeTypesInternalServerError{}
}

/*
OpenstackVolumeTypesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackVolumeTypesInternalServerError struct {
}

// IsSuccess returns true when this openstack volume types internal server error response has a 2xx status code
func (o *OpenstackVolumeTypesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack volume types internal server error response has a 3xx status code
func (o *OpenstackVolumeTypesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack volume types internal server error response has a 4xx status code
func (o *OpenstackVolumeTypesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack volume types internal server error response has a 5xx status code
func (o *OpenstackVolumeTypesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this openstack volume types internal server error response a status code equal to that given
func (o *OpenstackVolumeTypesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpenstackVolumeTypesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesInternalServerError ", 500)
}

func (o *OpenstackVolumeTypesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/volumes][%d] openstackVolumeTypesInternalServerError ", 500)
}

func (o *OpenstackVolumeTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
