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

// CheckerOpenstackTaikunLbImageReader is a Reader for the CheckerOpenstackTaikunLbImage structure.
type CheckerOpenstackTaikunLbImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerOpenstackTaikunLbImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerOpenstackTaikunLbImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerOpenstackTaikunLbImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerOpenstackTaikunLbImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerOpenstackTaikunLbImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerOpenstackTaikunLbImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerOpenstackTaikunLbImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerOpenstackTaikunLbImageOK creates a CheckerOpenstackTaikunLbImageOK with default headers values
func NewCheckerOpenstackTaikunLbImageOK() *CheckerOpenstackTaikunLbImageOK {
	return &CheckerOpenstackTaikunLbImageOK{}
}

/* CheckerOpenstackTaikunLbImageOK describes a response with status code 200, with default header values.

Success
*/
type CheckerOpenstackTaikunLbImageOK struct {
	Payload models.Unit
}

func (o *CheckerOpenstackTaikunLbImageOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageOK  %+v", 200, o.Payload)
}
func (o *CheckerOpenstackTaikunLbImageOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageBadRequest creates a CheckerOpenstackTaikunLbImageBadRequest with default headers values
func NewCheckerOpenstackTaikunLbImageBadRequest() *CheckerOpenstackTaikunLbImageBadRequest {
	return &CheckerOpenstackTaikunLbImageBadRequest{}
}

/* CheckerOpenstackTaikunLbImageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerOpenstackTaikunLbImageBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerOpenstackTaikunLbImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageBadRequest  %+v", 400, o.Payload)
}
func (o *CheckerOpenstackTaikunLbImageBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageUnauthorized creates a CheckerOpenstackTaikunLbImageUnauthorized with default headers values
func NewCheckerOpenstackTaikunLbImageUnauthorized() *CheckerOpenstackTaikunLbImageUnauthorized {
	return &CheckerOpenstackTaikunLbImageUnauthorized{}
}

/* CheckerOpenstackTaikunLbImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerOpenstackTaikunLbImageUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOpenstackTaikunLbImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerOpenstackTaikunLbImageUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageForbidden creates a CheckerOpenstackTaikunLbImageForbidden with default headers values
func NewCheckerOpenstackTaikunLbImageForbidden() *CheckerOpenstackTaikunLbImageForbidden {
	return &CheckerOpenstackTaikunLbImageForbidden{}
}

/* CheckerOpenstackTaikunLbImageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerOpenstackTaikunLbImageForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOpenstackTaikunLbImageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageForbidden  %+v", 403, o.Payload)
}
func (o *CheckerOpenstackTaikunLbImageForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageNotFound creates a CheckerOpenstackTaikunLbImageNotFound with default headers values
func NewCheckerOpenstackTaikunLbImageNotFound() *CheckerOpenstackTaikunLbImageNotFound {
	return &CheckerOpenstackTaikunLbImageNotFound{}
}

/* CheckerOpenstackTaikunLbImageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerOpenstackTaikunLbImageNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerOpenstackTaikunLbImageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageNotFound  %+v", 404, o.Payload)
}
func (o *CheckerOpenstackTaikunLbImageNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerOpenstackTaikunLbImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOpenstackTaikunLbImageInternalServerError creates a CheckerOpenstackTaikunLbImageInternalServerError with default headers values
func NewCheckerOpenstackTaikunLbImageInternalServerError() *CheckerOpenstackTaikunLbImageInternalServerError {
	return &CheckerOpenstackTaikunLbImageInternalServerError{}
}

/* CheckerOpenstackTaikunLbImageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerOpenstackTaikunLbImageInternalServerError struct {
}

func (o *CheckerOpenstackTaikunLbImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/taikun-lb-image/{id}][%d] checkerOpenstackTaikunLbImageInternalServerError ", 500)
}

func (o *CheckerOpenstackTaikunLbImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
