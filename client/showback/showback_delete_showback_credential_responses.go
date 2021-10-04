// Code generated by go-swagger; DO NOT EDIT.

package showback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackDeleteShowbackCredentialReader is a Reader for the ShowbackDeleteShowbackCredential structure.
type ShowbackDeleteShowbackCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackDeleteShowbackCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackDeleteShowbackCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackDeleteShowbackCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackDeleteShowbackCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackDeleteShowbackCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackDeleteShowbackCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackDeleteShowbackCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackDeleteShowbackCredentialOK creates a ShowbackDeleteShowbackCredentialOK with default headers values
func NewShowbackDeleteShowbackCredentialOK() *ShowbackDeleteShowbackCredentialOK {
	return &ShowbackDeleteShowbackCredentialOK{}
}

/* ShowbackDeleteShowbackCredentialOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackDeleteShowbackCredentialOK struct {
	Payload models.Unit
}

func (o *ShowbackDeleteShowbackCredentialOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/delete][%d] showbackDeleteShowbackCredentialOK  %+v", 200, o.Payload)
}
func (o *ShowbackDeleteShowbackCredentialOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ShowbackDeleteShowbackCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackDeleteShowbackCredentialBadRequest creates a ShowbackDeleteShowbackCredentialBadRequest with default headers values
func NewShowbackDeleteShowbackCredentialBadRequest() *ShowbackDeleteShowbackCredentialBadRequest {
	return &ShowbackDeleteShowbackCredentialBadRequest{}
}

/* ShowbackDeleteShowbackCredentialBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackDeleteShowbackCredentialBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackDeleteShowbackCredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/delete][%d] showbackDeleteShowbackCredentialBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackDeleteShowbackCredentialBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackDeleteShowbackCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackDeleteShowbackCredentialUnauthorized creates a ShowbackDeleteShowbackCredentialUnauthorized with default headers values
func NewShowbackDeleteShowbackCredentialUnauthorized() *ShowbackDeleteShowbackCredentialUnauthorized {
	return &ShowbackDeleteShowbackCredentialUnauthorized{}
}

/* ShowbackDeleteShowbackCredentialUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackDeleteShowbackCredentialUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackDeleteShowbackCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/delete][%d] showbackDeleteShowbackCredentialUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackDeleteShowbackCredentialUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackDeleteShowbackCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackDeleteShowbackCredentialForbidden creates a ShowbackDeleteShowbackCredentialForbidden with default headers values
func NewShowbackDeleteShowbackCredentialForbidden() *ShowbackDeleteShowbackCredentialForbidden {
	return &ShowbackDeleteShowbackCredentialForbidden{}
}

/* ShowbackDeleteShowbackCredentialForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackDeleteShowbackCredentialForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackDeleteShowbackCredentialForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/delete][%d] showbackDeleteShowbackCredentialForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackDeleteShowbackCredentialForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackDeleteShowbackCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackDeleteShowbackCredentialNotFound creates a ShowbackDeleteShowbackCredentialNotFound with default headers values
func NewShowbackDeleteShowbackCredentialNotFound() *ShowbackDeleteShowbackCredentialNotFound {
	return &ShowbackDeleteShowbackCredentialNotFound{}
}

/* ShowbackDeleteShowbackCredentialNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackDeleteShowbackCredentialNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackDeleteShowbackCredentialNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/delete][%d] showbackDeleteShowbackCredentialNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackDeleteShowbackCredentialNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackDeleteShowbackCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackDeleteShowbackCredentialInternalServerError creates a ShowbackDeleteShowbackCredentialInternalServerError with default headers values
func NewShowbackDeleteShowbackCredentialInternalServerError() *ShowbackDeleteShowbackCredentialInternalServerError {
	return &ShowbackDeleteShowbackCredentialInternalServerError{}
}

/* ShowbackDeleteShowbackCredentialInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackDeleteShowbackCredentialInternalServerError struct {
}

func (o *ShowbackDeleteShowbackCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/delete][%d] showbackDeleteShowbackCredentialInternalServerError ", 500)
}

func (o *ShowbackDeleteShowbackCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
