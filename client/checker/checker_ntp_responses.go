// Code generated by go-swagger; DO NOT EDIT.

package checker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CheckerNtpReader is a Reader for the CheckerNtp structure.
type CheckerNtpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerNtpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerNtpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerNtpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerNtpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerNtpForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerNtpNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCheckerNtpTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerNtpInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerNtpOK creates a CheckerNtpOK with default headers values
func NewCheckerNtpOK() *CheckerNtpOK {
	return &CheckerNtpOK{}
}

/* CheckerNtpOK describes a response with status code 200, with default header values.

Success
*/
type CheckerNtpOK struct {
	Payload models.Unit
}

func (o *CheckerNtpOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpOK  %+v", 200, o.Payload)
}
func (o *CheckerNtpOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerNtpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNtpBadRequest creates a CheckerNtpBadRequest with default headers values
func NewCheckerNtpBadRequest() *CheckerNtpBadRequest {
	return &CheckerNtpBadRequest{}
}

/* CheckerNtpBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerNtpBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerNtpBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpBadRequest  %+v", 400, o.Payload)
}
func (o *CheckerNtpBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerNtpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNtpUnauthorized creates a CheckerNtpUnauthorized with default headers values
func NewCheckerNtpUnauthorized() *CheckerNtpUnauthorized {
	return &CheckerNtpUnauthorized{}
}

/* CheckerNtpUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerNtpUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerNtpUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerNtpUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerNtpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNtpForbidden creates a CheckerNtpForbidden with default headers values
func NewCheckerNtpForbidden() *CheckerNtpForbidden {
	return &CheckerNtpForbidden{}
}

/* CheckerNtpForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerNtpForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerNtpForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpForbidden  %+v", 403, o.Payload)
}
func (o *CheckerNtpForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerNtpForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNtpNotFound creates a CheckerNtpNotFound with default headers values
func NewCheckerNtpNotFound() *CheckerNtpNotFound {
	return &CheckerNtpNotFound{}
}

/* CheckerNtpNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerNtpNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerNtpNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpNotFound  %+v", 404, o.Payload)
}
func (o *CheckerNtpNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerNtpNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNtpTooManyRequests creates a CheckerNtpTooManyRequests with default headers values
func NewCheckerNtpTooManyRequests() *CheckerNtpTooManyRequests {
	return &CheckerNtpTooManyRequests{}
}

/* CheckerNtpTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type CheckerNtpTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *CheckerNtpTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpTooManyRequests  %+v", 429, o.Payload)
}
func (o *CheckerNtpTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerNtpTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerNtpInternalServerError creates a CheckerNtpInternalServerError with default headers values
func NewCheckerNtpInternalServerError() *CheckerNtpInternalServerError {
	return &CheckerNtpInternalServerError{}
}

/* CheckerNtpInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerNtpInternalServerError struct {
}

func (o *CheckerNtpInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/ntp][%d] checkerNtpInternalServerError ", 500)
}

func (o *CheckerNtpInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
