// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneCreateReader is a Reader for the StandAloneCreate structure.
type StandAloneCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewStandAloneCreateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneCreateOK creates a StandAloneCreateOK with default headers values
func NewStandAloneCreateOK() *StandAloneCreateOK {
	return &StandAloneCreateOK{}
}

/* StandAloneCreateOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneCreateOK struct {
	Payload models.Unit
}

func (o *StandAloneCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateOK  %+v", 200, o.Payload)
}
func (o *StandAloneCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCreateBadRequest creates a StandAloneCreateBadRequest with default headers values
func NewStandAloneCreateBadRequest() *StandAloneCreateBadRequest {
	return &StandAloneCreateBadRequest{}
}

/* StandAloneCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCreateUnauthorized creates a StandAloneCreateUnauthorized with default headers values
func NewStandAloneCreateUnauthorized() *StandAloneCreateUnauthorized {
	return &StandAloneCreateUnauthorized{}
}

/* StandAloneCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCreateForbidden creates a StandAloneCreateForbidden with default headers values
func NewStandAloneCreateForbidden() *StandAloneCreateForbidden {
	return &StandAloneCreateForbidden{}
}

/* StandAloneCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCreateNotFound creates a StandAloneCreateNotFound with default headers values
func NewStandAloneCreateNotFound() *StandAloneCreateNotFound {
	return &StandAloneCreateNotFound{}
}

/* StandAloneCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCreateTooManyRequests creates a StandAloneCreateTooManyRequests with default headers values
func NewStandAloneCreateTooManyRequests() *StandAloneCreateTooManyRequests {
	return &StandAloneCreateTooManyRequests{}
}

/* StandAloneCreateTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type StandAloneCreateTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneCreateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateTooManyRequests  %+v", 429, o.Payload)
}
func (o *StandAloneCreateTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneCreateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCreateInternalServerError creates a StandAloneCreateInternalServerError with default headers values
func NewStandAloneCreateInternalServerError() *StandAloneCreateInternalServerError {
	return &StandAloneCreateInternalServerError{}
}

/* StandAloneCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneCreateInternalServerError struct {
}

func (o *StandAloneCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/create][%d] standAloneCreateInternalServerError ", 500)
}

func (o *StandAloneCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
