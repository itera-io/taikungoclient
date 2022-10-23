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

// AutoscalingDisableAutoscalingReader is a Reader for the AutoscalingDisableAutoscaling structure.
type AutoscalingDisableAutoscalingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoscalingDisableAutoscalingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoscalingDisableAutoscalingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAutoscalingDisableAutoscalingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAutoscalingDisableAutoscalingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAutoscalingDisableAutoscalingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAutoscalingDisableAutoscalingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAutoscalingDisableAutoscalingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAutoscalingDisableAutoscalingOK creates a AutoscalingDisableAutoscalingOK with default headers values
func NewAutoscalingDisableAutoscalingOK() *AutoscalingDisableAutoscalingOK {
	return &AutoscalingDisableAutoscalingOK{}
}

/*
AutoscalingDisableAutoscalingOK describes a response with status code 200, with default header values.

Success
*/
type AutoscalingDisableAutoscalingOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this autoscaling disable autoscaling o k response has a 2xx status code
func (o *AutoscalingDisableAutoscalingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this autoscaling disable autoscaling o k response has a 3xx status code
func (o *AutoscalingDisableAutoscalingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling disable autoscaling o k response has a 4xx status code
func (o *AutoscalingDisableAutoscalingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this autoscaling disable autoscaling o k response has a 5xx status code
func (o *AutoscalingDisableAutoscalingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling disable autoscaling o k response a status code equal to that given
func (o *AutoscalingDisableAutoscalingOK) IsCode(code int) bool {
	return code == 200
}

func (o *AutoscalingDisableAutoscalingOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingOK  %+v", 200, o.Payload)
}

func (o *AutoscalingDisableAutoscalingOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingOK  %+v", 200, o.Payload)
}

func (o *AutoscalingDisableAutoscalingOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AutoscalingDisableAutoscalingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingDisableAutoscalingBadRequest creates a AutoscalingDisableAutoscalingBadRequest with default headers values
func NewAutoscalingDisableAutoscalingBadRequest() *AutoscalingDisableAutoscalingBadRequest {
	return &AutoscalingDisableAutoscalingBadRequest{}
}

/*
AutoscalingDisableAutoscalingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AutoscalingDisableAutoscalingBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this autoscaling disable autoscaling bad request response has a 2xx status code
func (o *AutoscalingDisableAutoscalingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling disable autoscaling bad request response has a 3xx status code
func (o *AutoscalingDisableAutoscalingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling disable autoscaling bad request response has a 4xx status code
func (o *AutoscalingDisableAutoscalingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling disable autoscaling bad request response has a 5xx status code
func (o *AutoscalingDisableAutoscalingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling disable autoscaling bad request response a status code equal to that given
func (o *AutoscalingDisableAutoscalingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AutoscalingDisableAutoscalingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingBadRequest  %+v", 400, o.Payload)
}

func (o *AutoscalingDisableAutoscalingBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingBadRequest  %+v", 400, o.Payload)
}

func (o *AutoscalingDisableAutoscalingBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AutoscalingDisableAutoscalingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingDisableAutoscalingUnauthorized creates a AutoscalingDisableAutoscalingUnauthorized with default headers values
func NewAutoscalingDisableAutoscalingUnauthorized() *AutoscalingDisableAutoscalingUnauthorized {
	return &AutoscalingDisableAutoscalingUnauthorized{}
}

/*
AutoscalingDisableAutoscalingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AutoscalingDisableAutoscalingUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this autoscaling disable autoscaling unauthorized response has a 2xx status code
func (o *AutoscalingDisableAutoscalingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling disable autoscaling unauthorized response has a 3xx status code
func (o *AutoscalingDisableAutoscalingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling disable autoscaling unauthorized response has a 4xx status code
func (o *AutoscalingDisableAutoscalingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling disable autoscaling unauthorized response has a 5xx status code
func (o *AutoscalingDisableAutoscalingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling disable autoscaling unauthorized response a status code equal to that given
func (o *AutoscalingDisableAutoscalingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AutoscalingDisableAutoscalingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingUnauthorized  %+v", 401, o.Payload)
}

func (o *AutoscalingDisableAutoscalingUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingUnauthorized  %+v", 401, o.Payload)
}

func (o *AutoscalingDisableAutoscalingUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AutoscalingDisableAutoscalingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingDisableAutoscalingForbidden creates a AutoscalingDisableAutoscalingForbidden with default headers values
func NewAutoscalingDisableAutoscalingForbidden() *AutoscalingDisableAutoscalingForbidden {
	return &AutoscalingDisableAutoscalingForbidden{}
}

/*
AutoscalingDisableAutoscalingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AutoscalingDisableAutoscalingForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this autoscaling disable autoscaling forbidden response has a 2xx status code
func (o *AutoscalingDisableAutoscalingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling disable autoscaling forbidden response has a 3xx status code
func (o *AutoscalingDisableAutoscalingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling disable autoscaling forbidden response has a 4xx status code
func (o *AutoscalingDisableAutoscalingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling disable autoscaling forbidden response has a 5xx status code
func (o *AutoscalingDisableAutoscalingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling disable autoscaling forbidden response a status code equal to that given
func (o *AutoscalingDisableAutoscalingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AutoscalingDisableAutoscalingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingForbidden  %+v", 403, o.Payload)
}

func (o *AutoscalingDisableAutoscalingForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingForbidden  %+v", 403, o.Payload)
}

func (o *AutoscalingDisableAutoscalingForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AutoscalingDisableAutoscalingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingDisableAutoscalingNotFound creates a AutoscalingDisableAutoscalingNotFound with default headers values
func NewAutoscalingDisableAutoscalingNotFound() *AutoscalingDisableAutoscalingNotFound {
	return &AutoscalingDisableAutoscalingNotFound{}
}

/*
AutoscalingDisableAutoscalingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AutoscalingDisableAutoscalingNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this autoscaling disable autoscaling not found response has a 2xx status code
func (o *AutoscalingDisableAutoscalingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling disable autoscaling not found response has a 3xx status code
func (o *AutoscalingDisableAutoscalingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling disable autoscaling not found response has a 4xx status code
func (o *AutoscalingDisableAutoscalingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this autoscaling disable autoscaling not found response has a 5xx status code
func (o *AutoscalingDisableAutoscalingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this autoscaling disable autoscaling not found response a status code equal to that given
func (o *AutoscalingDisableAutoscalingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AutoscalingDisableAutoscalingNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingNotFound  %+v", 404, o.Payload)
}

func (o *AutoscalingDisableAutoscalingNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingNotFound  %+v", 404, o.Payload)
}

func (o *AutoscalingDisableAutoscalingNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AutoscalingDisableAutoscalingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoscalingDisableAutoscalingInternalServerError creates a AutoscalingDisableAutoscalingInternalServerError with default headers values
func NewAutoscalingDisableAutoscalingInternalServerError() *AutoscalingDisableAutoscalingInternalServerError {
	return &AutoscalingDisableAutoscalingInternalServerError{}
}

/*
AutoscalingDisableAutoscalingInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AutoscalingDisableAutoscalingInternalServerError struct {
}

// IsSuccess returns true when this autoscaling disable autoscaling internal server error response has a 2xx status code
func (o *AutoscalingDisableAutoscalingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this autoscaling disable autoscaling internal server error response has a 3xx status code
func (o *AutoscalingDisableAutoscalingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this autoscaling disable autoscaling internal server error response has a 4xx status code
func (o *AutoscalingDisableAutoscalingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this autoscaling disable autoscaling internal server error response has a 5xx status code
func (o *AutoscalingDisableAutoscalingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this autoscaling disable autoscaling internal server error response a status code equal to that given
func (o *AutoscalingDisableAutoscalingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AutoscalingDisableAutoscalingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingInternalServerError ", 500)
}

func (o *AutoscalingDisableAutoscalingInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Autoscaling/disable][%d] autoscalingDisableAutoscalingInternalServerError ", 500)
}

func (o *AutoscalingDisableAutoscalingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
