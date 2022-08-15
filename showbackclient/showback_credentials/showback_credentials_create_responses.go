// Code generated by go-swagger; DO NOT EDIT.

package showback_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackCredentialsCreateReader is a Reader for the ShowbackCredentialsCreate structure.
type ShowbackCredentialsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackCredentialsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackCredentialsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackCredentialsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackCredentialsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackCredentialsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackCredentialsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackCredentialsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackCredentialsCreateOK creates a ShowbackCredentialsCreateOK with default headers values
func NewShowbackCredentialsCreateOK() *ShowbackCredentialsCreateOK {
	return &ShowbackCredentialsCreateOK{}
}

/* ShowbackCredentialsCreateOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackCredentialsCreateOK struct {
	Payload *models.APIResponse
}

func (o *ShowbackCredentialsCreateOK) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateOK  %+v", 200, o.Payload)
}
func (o *ShowbackCredentialsCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *ShowbackCredentialsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateBadRequest creates a ShowbackCredentialsCreateBadRequest with default headers values
func NewShowbackCredentialsCreateBadRequest() *ShowbackCredentialsCreateBadRequest {
	return &ShowbackCredentialsCreateBadRequest{}
}

/* ShowbackCredentialsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackCredentialsCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackCredentialsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackCredentialsCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateUnauthorized creates a ShowbackCredentialsCreateUnauthorized with default headers values
func NewShowbackCredentialsCreateUnauthorized() *ShowbackCredentialsCreateUnauthorized {
	return &ShowbackCredentialsCreateUnauthorized{}
}

/* ShowbackCredentialsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackCredentialsCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackCredentialsCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateForbidden creates a ShowbackCredentialsCreateForbidden with default headers values
func NewShowbackCredentialsCreateForbidden() *ShowbackCredentialsCreateForbidden {
	return &ShowbackCredentialsCreateForbidden{}
}

/* ShowbackCredentialsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackCredentialsCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackCredentialsCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateNotFound creates a ShowbackCredentialsCreateNotFound with default headers values
func NewShowbackCredentialsCreateNotFound() *ShowbackCredentialsCreateNotFound {
	return &ShowbackCredentialsCreateNotFound{}
}

/* ShowbackCredentialsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackCredentialsCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackCredentialsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackCredentialsCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackCredentialsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackCredentialsCreateInternalServerError creates a ShowbackCredentialsCreateInternalServerError with default headers values
func NewShowbackCredentialsCreateInternalServerError() *ShowbackCredentialsCreateInternalServerError {
	return &ShowbackCredentialsCreateInternalServerError{}
}

/* ShowbackCredentialsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackCredentialsCreateInternalServerError struct {
}

func (o *ShowbackCredentialsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /showback/v{v}/ShowbackCredentials/create][%d] showbackCredentialsCreateInternalServerError ", 500)
}

func (o *ShowbackCredentialsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
