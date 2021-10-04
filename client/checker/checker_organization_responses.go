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

// CheckerOrganizationReader is a Reader for the CheckerOrganization structure.
type CheckerOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerOrganizationOK creates a CheckerOrganizationOK with default headers values
func NewCheckerOrganizationOK() *CheckerOrganizationOK {
	return &CheckerOrganizationOK{}
}

/* CheckerOrganizationOK describes a response with status code 200, with default header values.

Success
*/
type CheckerOrganizationOK struct {
	Payload models.Unit
}

func (o *CheckerOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationOK  %+v", 200, o.Payload)
}
func (o *CheckerOrganizationOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationBadRequest creates a CheckerOrganizationBadRequest with default headers values
func NewCheckerOrganizationBadRequest() *CheckerOrganizationBadRequest {
	return &CheckerOrganizationBadRequest{}
}

/* CheckerOrganizationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerOrganizationBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationBadRequest  %+v", 400, o.Payload)
}
func (o *CheckerOrganizationBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationUnauthorized creates a CheckerOrganizationUnauthorized with default headers values
func NewCheckerOrganizationUnauthorized() *CheckerOrganizationUnauthorized {
	return &CheckerOrganizationUnauthorized{}
}

/* CheckerOrganizationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerOrganizationUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerOrganizationUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationForbidden creates a CheckerOrganizationForbidden with default headers values
func NewCheckerOrganizationForbidden() *CheckerOrganizationForbidden {
	return &CheckerOrganizationForbidden{}
}

/* CheckerOrganizationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerOrganizationForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationForbidden  %+v", 403, o.Payload)
}
func (o *CheckerOrganizationForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationNotFound creates a CheckerOrganizationNotFound with default headers values
func NewCheckerOrganizationNotFound() *CheckerOrganizationNotFound {
	return &CheckerOrganizationNotFound{}
}

/* CheckerOrganizationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerOrganizationNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationNotFound  %+v", 404, o.Payload)
}
func (o *CheckerOrganizationNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationInternalServerError creates a CheckerOrganizationInternalServerError with default headers values
func NewCheckerOrganizationInternalServerError() *CheckerOrganizationInternalServerError {
	return &CheckerOrganizationInternalServerError{}
}

/* CheckerOrganizationInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerOrganizationInternalServerError struct {
}

func (o *CheckerOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationInternalServerError ", 500)
}

func (o *CheckerOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
