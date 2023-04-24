// Code generated by go-swagger; DO NOT EDIT.

package autoscaling

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AutoscalingEnableAutoscalingReader is a Reader for the AutoscalingEnableAutoscaling structure.
type AutoscalingEnableAutoscalingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoscalingEnableAutoscalingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoscalingEnableAutoscalingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAutoscalingEnableAutoscalingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAutoscalingEnableAutoscalingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAutoscalingEnableAutoscalingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAutoscalingEnableAutoscalingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAutoscalingEnableAutoscalingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAutoscalingEnableAutoscalingOK creates a AutoscalingEnableAutoscalingOK with default headers values
func NewAutoscalingEnableAutoscalingOK() *AutoscalingEnableAutoscalingOK {
	return &AutoscalingEnableAutoscalingOK{}
}

/*
AutoscalingEnableAutoscalingOK describes a response with status code 200, with default header values.

Success
*/
type AutoscalingEnableAutoscalingOK struct {
}

// IsSuccess returns true when this autoscaling enable autoscaling o k response has a 2xx status code
func (o *AutoscalingEnableAutoscalingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this autoscaling enable autoscaling o k response has a 3xx status code
func (o *AutoscalingEnableAutoscalingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling enable autoscaling o k response has a 4xx status code
func (o *AutoscalingEnableAutoscalingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this autoscaling enable autoscaling o k response has a 5xx status code
func (o *AutoscalingEnableAutoscalingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling enable autoscaling o k response a status code equal to that given
func (o *AutoscalingEnableAutoscalingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the autoscaling enable autoscaling o k response
func (o *AutoscalingEnableAutoscalingOK) Code() int {
	return 200
}

func (o *AutoscalingEnableAutoscalingOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingOK ", 200)
}

func (o *AutoscalingEnableAutoscalingOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingOK ", 200)
}

func (o *AutoscalingEnableAutoscalingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAutoscalingEnableAutoscalingBadRequest creates a AutoscalingEnableAutoscalingBadRequest with default headers values
func NewAutoscalingEnableAutoscalingBadRequest() *AutoscalingEnableAutoscalingBadRequest {
	return &AutoscalingEnableAutoscalingBadRequest{}
}

/*
AutoscalingEnableAutoscalingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AutoscalingEnableAutoscalingBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this autoscaling enable autoscaling bad request response has a 2xx status code
func (o *AutoscalingEnableAutoscalingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling enable autoscaling bad request response has a 3xx status code
func (o *AutoscalingEnableAutoscalingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling enable autoscaling bad request response has a 4xx status code
func (o *AutoscalingEnableAutoscalingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling enable autoscaling bad request response has a 5xx status code
func (o *AutoscalingEnableAutoscalingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling enable autoscaling bad request response a status code equal to that given
func (o *AutoscalingEnableAutoscalingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the autoscaling enable autoscaling bad request response
func (o *AutoscalingEnableAutoscalingBadRequest) Code() int {
	return 400
}

func (o *AutoscalingEnableAutoscalingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingBadRequest  %+v", 400, o.Payload)
}

func (o *AutoscalingEnableAutoscalingBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingBadRequest  %+v", 400, o.Payload)
}

func (o *AutoscalingEnableAutoscalingBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AutoscalingEnableAutoscalingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEnableAutoscalingUnauthorized creates a AutoscalingEnableAutoscalingUnauthorized with default headers values
func NewAutoscalingEnableAutoscalingUnauthorized() *AutoscalingEnableAutoscalingUnauthorized {
	return &AutoscalingEnableAutoscalingUnauthorized{}
}

/*
AutoscalingEnableAutoscalingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AutoscalingEnableAutoscalingUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this autoscaling enable autoscaling unauthorized response has a 2xx status code
func (o *AutoscalingEnableAutoscalingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling enable autoscaling unauthorized response has a 3xx status code
func (o *AutoscalingEnableAutoscalingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling enable autoscaling unauthorized response has a 4xx status code
func (o *AutoscalingEnableAutoscalingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling enable autoscaling unauthorized response has a 5xx status code
func (o *AutoscalingEnableAutoscalingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling enable autoscaling unauthorized response a status code equal to that given
func (o *AutoscalingEnableAutoscalingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the autoscaling enable autoscaling unauthorized response
func (o *AutoscalingEnableAutoscalingUnauthorized) Code() int {
	return 401
}

func (o *AutoscalingEnableAutoscalingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingUnauthorized  %+v", 401, o.Payload)
}

func (o *AutoscalingEnableAutoscalingUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingUnauthorized  %+v", 401, o.Payload)
}

func (o *AutoscalingEnableAutoscalingUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AutoscalingEnableAutoscalingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEnableAutoscalingForbidden creates a AutoscalingEnableAutoscalingForbidden with default headers values
func NewAutoscalingEnableAutoscalingForbidden() *AutoscalingEnableAutoscalingForbidden {
	return &AutoscalingEnableAutoscalingForbidden{}
}

/*
AutoscalingEnableAutoscalingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AutoscalingEnableAutoscalingForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this autoscaling enable autoscaling forbidden response has a 2xx status code
func (o *AutoscalingEnableAutoscalingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling enable autoscaling forbidden response has a 3xx status code
func (o *AutoscalingEnableAutoscalingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling enable autoscaling forbidden response has a 4xx status code
func (o *AutoscalingEnableAutoscalingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling enable autoscaling forbidden response has a 5xx status code
func (o *AutoscalingEnableAutoscalingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling enable autoscaling forbidden response a status code equal to that given
func (o *AutoscalingEnableAutoscalingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the autoscaling enable autoscaling forbidden response
func (o *AutoscalingEnableAutoscalingForbidden) Code() int {
	return 403
}

func (o *AutoscalingEnableAutoscalingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingForbidden  %+v", 403, o.Payload)
}

func (o *AutoscalingEnableAutoscalingForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingForbidden  %+v", 403, o.Payload)
}

func (o *AutoscalingEnableAutoscalingForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AutoscalingEnableAutoscalingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEnableAutoscalingNotFound creates a AutoscalingEnableAutoscalingNotFound with default headers values
func NewAutoscalingEnableAutoscalingNotFound() *AutoscalingEnableAutoscalingNotFound {
	return &AutoscalingEnableAutoscalingNotFound{}
}

/*
AutoscalingEnableAutoscalingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AutoscalingEnableAutoscalingNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this autoscaling enable autoscaling not found response has a 2xx status code
func (o *AutoscalingEnableAutoscalingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling enable autoscaling not found response has a 3xx status code
func (o *AutoscalingEnableAutoscalingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling enable autoscaling not found response has a 4xx status code
func (o *AutoscalingEnableAutoscalingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling enable autoscaling not found response has a 5xx status code
func (o *AutoscalingEnableAutoscalingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling enable autoscaling not found response a status code equal to that given
func (o *AutoscalingEnableAutoscalingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the autoscaling enable autoscaling not found response
func (o *AutoscalingEnableAutoscalingNotFound) Code() int {
	return 404
}

func (o *AutoscalingEnableAutoscalingNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingNotFound  %+v", 404, o.Payload)
}

func (o *AutoscalingEnableAutoscalingNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingNotFound  %+v", 404, o.Payload)
}

func (o *AutoscalingEnableAutoscalingNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AutoscalingEnableAutoscalingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEnableAutoscalingInternalServerError creates a AutoscalingEnableAutoscalingInternalServerError with default headers values
func NewAutoscalingEnableAutoscalingInternalServerError() *AutoscalingEnableAutoscalingInternalServerError {
	return &AutoscalingEnableAutoscalingInternalServerError{}
}

/*
AutoscalingEnableAutoscalingInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AutoscalingEnableAutoscalingInternalServerError struct {
}

// IsSuccess returns true when this autoscaling enable autoscaling internal server error response has a 2xx status code
func (o *AutoscalingEnableAutoscalingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling enable autoscaling internal server error response has a 3xx status code
func (o *AutoscalingEnableAutoscalingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling enable autoscaling internal server error response has a 4xx status code
func (o *AutoscalingEnableAutoscalingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this autoscaling enable autoscaling internal server error response has a 5xx status code
func (o *AutoscalingEnableAutoscalingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this autoscaling enable autoscaling internal server error response a status code equal to that given
func (o *AutoscalingEnableAutoscalingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the autoscaling enable autoscaling internal server error response
func (o *AutoscalingEnableAutoscalingInternalServerError) Code() int {
	return 500
}

func (o *AutoscalingEnableAutoscalingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingInternalServerError ", 500)
}

func (o *AutoscalingEnableAutoscalingInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/enable][%d] autoscalingEnableAutoscalingInternalServerError ", 500)
}

func (o *AutoscalingEnableAutoscalingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
