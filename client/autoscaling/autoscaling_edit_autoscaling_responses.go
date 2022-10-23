// Code generated by go-swagger; DO NOT EDIT.

package autoscaling

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AutoscalingEditAutoscalingReader is a Reader for the AutoscalingEditAutoscaling structure.
type AutoscalingEditAutoscalingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoscalingEditAutoscalingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoscalingEditAutoscalingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAutoscalingEditAutoscalingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAutoscalingEditAutoscalingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAutoscalingEditAutoscalingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAutoscalingEditAutoscalingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAutoscalingEditAutoscalingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAutoscalingEditAutoscalingOK creates a AutoscalingEditAutoscalingOK with default headers values
func NewAutoscalingEditAutoscalingOK() *AutoscalingEditAutoscalingOK {
	return &AutoscalingEditAutoscalingOK{}
}

/*
AutoscalingEditAutoscalingOK describes a response with status code 200, with default header values.

Success
*/
type AutoscalingEditAutoscalingOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this autoscaling edit autoscaling o k response has a 2xx status code
func (o *AutoscalingEditAutoscalingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this autoscaling edit autoscaling o k response has a 3xx status code
func (o *AutoscalingEditAutoscalingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling edit autoscaling o k response has a 4xx status code
func (o *AutoscalingEditAutoscalingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this autoscaling edit autoscaling o k response has a 5xx status code
func (o *AutoscalingEditAutoscalingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling edit autoscaling o k response a status code equal to that given
func (o *AutoscalingEditAutoscalingOK) IsCode(code int) bool {
	return code == 200
}

func (o *AutoscalingEditAutoscalingOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingOK  %+v", 200, o.Payload)
}

func (o *AutoscalingEditAutoscalingOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingOK  %+v", 200, o.Payload)
}

func (o *AutoscalingEditAutoscalingOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AutoscalingEditAutoscalingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEditAutoscalingBadRequest creates a AutoscalingEditAutoscalingBadRequest with default headers values
func NewAutoscalingEditAutoscalingBadRequest() *AutoscalingEditAutoscalingBadRequest {
	return &AutoscalingEditAutoscalingBadRequest{}
}

/*
AutoscalingEditAutoscalingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AutoscalingEditAutoscalingBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this autoscaling edit autoscaling bad request response has a 2xx status code
func (o *AutoscalingEditAutoscalingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling edit autoscaling bad request response has a 3xx status code
func (o *AutoscalingEditAutoscalingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling edit autoscaling bad request response has a 4xx status code
func (o *AutoscalingEditAutoscalingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling edit autoscaling bad request response has a 5xx status code
func (o *AutoscalingEditAutoscalingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling edit autoscaling bad request response a status code equal to that given
func (o *AutoscalingEditAutoscalingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AutoscalingEditAutoscalingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingBadRequest  %+v", 400, o.Payload)
}

func (o *AutoscalingEditAutoscalingBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingBadRequest  %+v", 400, o.Payload)
}

func (o *AutoscalingEditAutoscalingBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AutoscalingEditAutoscalingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEditAutoscalingUnauthorized creates a AutoscalingEditAutoscalingUnauthorized with default headers values
func NewAutoscalingEditAutoscalingUnauthorized() *AutoscalingEditAutoscalingUnauthorized {
	return &AutoscalingEditAutoscalingUnauthorized{}
}

/*
AutoscalingEditAutoscalingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AutoscalingEditAutoscalingUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this autoscaling edit autoscaling unauthorized response has a 2xx status code
func (o *AutoscalingEditAutoscalingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling edit autoscaling unauthorized response has a 3xx status code
func (o *AutoscalingEditAutoscalingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling edit autoscaling unauthorized response has a 4xx status code
func (o *AutoscalingEditAutoscalingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling edit autoscaling unauthorized response has a 5xx status code
func (o *AutoscalingEditAutoscalingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling edit autoscaling unauthorized response a status code equal to that given
func (o *AutoscalingEditAutoscalingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AutoscalingEditAutoscalingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingUnauthorized  %+v", 401, o.Payload)
}

func (o *AutoscalingEditAutoscalingUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingUnauthorized  %+v", 401, o.Payload)
}

func (o *AutoscalingEditAutoscalingUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AutoscalingEditAutoscalingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEditAutoscalingForbidden creates a AutoscalingEditAutoscalingForbidden with default headers values
func NewAutoscalingEditAutoscalingForbidden() *AutoscalingEditAutoscalingForbidden {
	return &AutoscalingEditAutoscalingForbidden{}
}

/*
AutoscalingEditAutoscalingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AutoscalingEditAutoscalingForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this autoscaling edit autoscaling forbidden response has a 2xx status code
func (o *AutoscalingEditAutoscalingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling edit autoscaling forbidden response has a 3xx status code
func (o *AutoscalingEditAutoscalingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling edit autoscaling forbidden response has a 4xx status code
func (o *AutoscalingEditAutoscalingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling edit autoscaling forbidden response has a 5xx status code
func (o *AutoscalingEditAutoscalingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling edit autoscaling forbidden response a status code equal to that given
func (o *AutoscalingEditAutoscalingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AutoscalingEditAutoscalingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingForbidden  %+v", 403, o.Payload)
}

func (o *AutoscalingEditAutoscalingForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingForbidden  %+v", 403, o.Payload)
}

func (o *AutoscalingEditAutoscalingForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AutoscalingEditAutoscalingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEditAutoscalingNotFound creates a AutoscalingEditAutoscalingNotFound with default headers values
func NewAutoscalingEditAutoscalingNotFound() *AutoscalingEditAutoscalingNotFound {
	return &AutoscalingEditAutoscalingNotFound{}
}

/*
AutoscalingEditAutoscalingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AutoscalingEditAutoscalingNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this autoscaling edit autoscaling not found response has a 2xx status code
func (o *AutoscalingEditAutoscalingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling edit autoscaling not found response has a 3xx status code
func (o *AutoscalingEditAutoscalingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling edit autoscaling not found response has a 4xx status code
func (o *AutoscalingEditAutoscalingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling edit autoscaling not found response has a 5xx status code
func (o *AutoscalingEditAutoscalingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling edit autoscaling not found response a status code equal to that given
func (o *AutoscalingEditAutoscalingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AutoscalingEditAutoscalingNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingNotFound  %+v", 404, o.Payload)
}

func (o *AutoscalingEditAutoscalingNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingNotFound  %+v", 404, o.Payload)
}

func (o *AutoscalingEditAutoscalingNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AutoscalingEditAutoscalingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingEditAutoscalingInternalServerError creates a AutoscalingEditAutoscalingInternalServerError with default headers values
func NewAutoscalingEditAutoscalingInternalServerError() *AutoscalingEditAutoscalingInternalServerError {
	return &AutoscalingEditAutoscalingInternalServerError{}
}

/*
AutoscalingEditAutoscalingInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AutoscalingEditAutoscalingInternalServerError struct {
}

// IsSuccess returns true when this autoscaling edit autoscaling internal server error response has a 2xx status code
func (o *AutoscalingEditAutoscalingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling edit autoscaling internal server error response has a 3xx status code
func (o *AutoscalingEditAutoscalingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling edit autoscaling internal server error response has a 4xx status code
func (o *AutoscalingEditAutoscalingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this autoscaling edit autoscaling internal server error response has a 5xx status code
func (o *AutoscalingEditAutoscalingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this autoscaling edit autoscaling internal server error response a status code equal to that given
func (o *AutoscalingEditAutoscalingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AutoscalingEditAutoscalingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingInternalServerError ", 500)
}

func (o *AutoscalingEditAutoscalingInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/edit][%d] autoscalingEditAutoscalingInternalServerError ", 500)
}

func (o *AutoscalingEditAutoscalingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}