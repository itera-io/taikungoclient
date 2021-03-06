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

// CheckerYamlReader is a Reader for the CheckerYaml structure.
type CheckerYamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerYamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerYamlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerYamlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerYamlUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerYamlForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerYamlNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerYamlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerYamlOK creates a CheckerYamlOK with default headers values
func NewCheckerYamlOK() *CheckerYamlOK {
	return &CheckerYamlOK{}
}

/* CheckerYamlOK describes a response with status code 200, with default header values.

Success
*/
type CheckerYamlOK struct {
	Payload models.Unit
}

func (o *CheckerYamlOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/yaml][%d] checkerYamlOK  %+v", 200, o.Payload)
}
func (o *CheckerYamlOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerYamlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerYamlBadRequest creates a CheckerYamlBadRequest with default headers values
func NewCheckerYamlBadRequest() *CheckerYamlBadRequest {
	return &CheckerYamlBadRequest{}
}

/* CheckerYamlBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerYamlBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CheckerYamlBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/yaml][%d] checkerYamlBadRequest  %+v", 400, o.Payload)
}
func (o *CheckerYamlBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerYamlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerYamlUnauthorized creates a CheckerYamlUnauthorized with default headers values
func NewCheckerYamlUnauthorized() *CheckerYamlUnauthorized {
	return &CheckerYamlUnauthorized{}
}

/* CheckerYamlUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerYamlUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CheckerYamlUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/yaml][%d] checkerYamlUnauthorized  %+v", 401, o.Payload)
}
func (o *CheckerYamlUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerYamlUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerYamlForbidden creates a CheckerYamlForbidden with default headers values
func NewCheckerYamlForbidden() *CheckerYamlForbidden {
	return &CheckerYamlForbidden{}
}

/* CheckerYamlForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerYamlForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CheckerYamlForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/yaml][%d] checkerYamlForbidden  %+v", 403, o.Payload)
}
func (o *CheckerYamlForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerYamlForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerYamlNotFound creates a CheckerYamlNotFound with default headers values
func NewCheckerYamlNotFound() *CheckerYamlNotFound {
	return &CheckerYamlNotFound{}
}

/* CheckerYamlNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerYamlNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CheckerYamlNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/yaml][%d] checkerYamlNotFound  %+v", 404, o.Payload)
}
func (o *CheckerYamlNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerYamlNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerYamlInternalServerError creates a CheckerYamlInternalServerError with default headers values
func NewCheckerYamlInternalServerError() *CheckerYamlInternalServerError {
	return &CheckerYamlInternalServerError{}
}

/* CheckerYamlInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerYamlInternalServerError struct {
}

func (o *CheckerYamlInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/yaml][%d] checkerYamlInternalServerError ", 500)
}

func (o *CheckerYamlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
