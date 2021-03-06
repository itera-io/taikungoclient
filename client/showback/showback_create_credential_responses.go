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

// ShowbackCreateCredentialReader is a Reader for the ShowbackCreateCredential structure.
type ShowbackCreateCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackCreateCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackCreateCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackCreateCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackCreateCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackCreateCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackCreateCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackCreateCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackCreateCredentialOK creates a ShowbackCreateCredentialOK with default headers values
func NewShowbackCreateCredentialOK() *ShowbackCreateCredentialOK {
	return &ShowbackCreateCredentialOK{}
}

/* ShowbackCreateCredentialOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackCreateCredentialOK struct {
	Payload *models.APIResponse
}

func (o *ShowbackCreateCredentialOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/create][%d] showbackCreateCredentialOK  %+v", 200, o.Payload)
}
func (o *ShowbackCreateCredentialOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *ShowbackCreateCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCreateCredentialBadRequest creates a ShowbackCreateCredentialBadRequest with default headers values
func NewShowbackCreateCredentialBadRequest() *ShowbackCreateCredentialBadRequest {
	return &ShowbackCreateCredentialBadRequest{}
}

/* ShowbackCreateCredentialBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackCreateCredentialBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackCreateCredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/create][%d] showbackCreateCredentialBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackCreateCredentialBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackCreateCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCreateCredentialUnauthorized creates a ShowbackCreateCredentialUnauthorized with default headers values
func NewShowbackCreateCredentialUnauthorized() *ShowbackCreateCredentialUnauthorized {
	return &ShowbackCreateCredentialUnauthorized{}
}

/* ShowbackCreateCredentialUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackCreateCredentialUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCreateCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/create][%d] showbackCreateCredentialUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackCreateCredentialUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCreateCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCreateCredentialForbidden creates a ShowbackCreateCredentialForbidden with default headers values
func NewShowbackCreateCredentialForbidden() *ShowbackCreateCredentialForbidden {
	return &ShowbackCreateCredentialForbidden{}
}

/* ShowbackCreateCredentialForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackCreateCredentialForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCreateCredentialForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/create][%d] showbackCreateCredentialForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackCreateCredentialForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCreateCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCreateCredentialNotFound creates a ShowbackCreateCredentialNotFound with default headers values
func NewShowbackCreateCredentialNotFound() *ShowbackCreateCredentialNotFound {
	return &ShowbackCreateCredentialNotFound{}
}

/* ShowbackCreateCredentialNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackCreateCredentialNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCreateCredentialNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/create][%d] showbackCreateCredentialNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackCreateCredentialNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCreateCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCreateCredentialInternalServerError creates a ShowbackCreateCredentialInternalServerError with default headers values
func NewShowbackCreateCredentialInternalServerError() *ShowbackCreateCredentialInternalServerError {
	return &ShowbackCreateCredentialInternalServerError{}
}

/* ShowbackCreateCredentialInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackCreateCredentialInternalServerError struct {
}

func (o *ShowbackCreateCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/credential/create][%d] showbackCreateCredentialInternalServerError ", 500)
}

func (o *ShowbackCreateCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
