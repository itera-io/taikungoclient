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

// CheckerOpenstackReader is a Reader for the CheckerOpenstack structure.
type CheckerOpenstackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerOpenstackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerOpenstackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerOpenstackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerOpenstackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerOpenstackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerOpenstackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerOpenstackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerOpenstackOK creates a CheckerOpenstackOK with default headers values
func NewCheckerOpenstackOK() *CheckerOpenstackOK {
	return &CheckerOpenstackOK{}
}

/* CheckerOpenstackOK describes a response with status code 200, with default header values.

Success
*/
type CheckerOpenstackOK struct {
	Payload models.Unit
}

func (o *CheckerOpenstackOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack][%d] checkerOpenstackOK  %+v", 200, o.Payload)
}
func (o *CheckerOpenstackOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerOpenstackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackBadRequest creates a CheckerOpenstackBadRequest with default headers values
func NewCheckerOpenstackBadRequest() *CheckerOpenstackBadRequest {
	return &CheckerOpenstackBadRequest{}
}

/* CheckerOpenstackBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerOpenstackBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerOpenstackBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack][%d] checkerOpenstackBadRequest  %+v", 400, o.Payload)
}
func (o *CheckerOpenstackBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackUnauthorized creates a CheckerOpenstackUnauthorized with default headers values
func NewCheckerOpenstackUnauthorized() *CheckerOpenstackUnauthorized {
	return &CheckerOpenstackUnauthorized{}
}

/* CheckerOpenstackUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerOpenstackUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOpenstackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack][%d] checkerOpenstackUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerOpenstackUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackForbidden creates a CheckerOpenstackForbidden with default headers values
func NewCheckerOpenstackForbidden() *CheckerOpenstackForbidden {
	return &CheckerOpenstackForbidden{}
}

/* CheckerOpenstackForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerOpenstackForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOpenstackForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack][%d] checkerOpenstackForbidden  %+v", 403, o.Payload)
}
func (o *CheckerOpenstackForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackNotFound creates a CheckerOpenstackNotFound with default headers values
func NewCheckerOpenstackNotFound() *CheckerOpenstackNotFound {
	return &CheckerOpenstackNotFound{}
}

/* CheckerOpenstackNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerOpenstackNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOpenstackNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack][%d] checkerOpenstackNotFound  %+v", 404, o.Payload)
}
func (o *CheckerOpenstackNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackInternalServerError creates a CheckerOpenstackInternalServerError with default headers values
func NewCheckerOpenstackInternalServerError() *CheckerOpenstackInternalServerError {
	return &CheckerOpenstackInternalServerError{}
}

/* CheckerOpenstackInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerOpenstackInternalServerError struct {
}

func (o *CheckerOpenstackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/openstack][%d] checkerOpenstackInternalServerError ", 500)
}

func (o *CheckerOpenstackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}