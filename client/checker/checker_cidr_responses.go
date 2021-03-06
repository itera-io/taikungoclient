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

// CheckerCidrReader is a Reader for the CheckerCidr structure.
type CheckerCidrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerCidrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerCidrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerCidrBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerCidrUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerCidrForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerCidrNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerCidrInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerCidrOK creates a CheckerCidrOK with default headers values
func NewCheckerCidrOK() *CheckerCidrOK {
	return &CheckerCidrOK{}
}

/* CheckerCidrOK describes a response with status code 200, with default header values.

Success
*/
type CheckerCidrOK struct {
	Payload models.Unit
}

func (o *CheckerCidrOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cidr][%d] checkerCidrOK  %+v", 200, o.Payload)
}
func (o *CheckerCidrOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerCidrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCidrBadRequest creates a CheckerCidrBadRequest with default headers values
func NewCheckerCidrBadRequest() *CheckerCidrBadRequest {
	return &CheckerCidrBadRequest{}
}

/* CheckerCidrBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerCidrBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerCidrBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cidr][%d] checkerCidrBadRequest  %+v", 400, o.Payload)
}
func (o *CheckerCidrBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerCidrBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCidrUnauthorized creates a CheckerCidrUnauthorized with default headers values
func NewCheckerCidrUnauthorized() *CheckerCidrUnauthorized {
	return &CheckerCidrUnauthorized{}
}

/* CheckerCidrUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerCidrUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerCidrUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cidr][%d] checkerCidrUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerCidrUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerCidrUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCidrForbidden creates a CheckerCidrForbidden with default headers values
func NewCheckerCidrForbidden() *CheckerCidrForbidden {
	return &CheckerCidrForbidden{}
}

/* CheckerCidrForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerCidrForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerCidrForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cidr][%d] checkerCidrForbidden  %+v", 403, o.Payload)
}
func (o *CheckerCidrForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerCidrForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCidrNotFound creates a CheckerCidrNotFound with default headers values
func NewCheckerCidrNotFound() *CheckerCidrNotFound {
	return &CheckerCidrNotFound{}
}

/* CheckerCidrNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerCidrNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerCidrNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cidr][%d] checkerCidrNotFound  %+v", 404, o.Payload)
}
func (o *CheckerCidrNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerCidrNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCidrInternalServerError creates a CheckerCidrInternalServerError with default headers values
func NewCheckerCidrInternalServerError() *CheckerCidrInternalServerError {
	return &CheckerCidrInternalServerError{}
}

/* CheckerCidrInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerCidrInternalServerError struct {
}

func (o *CheckerCidrInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cidr][%d] checkerCidrInternalServerError ", 500)
}

func (o *CheckerCidrInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
