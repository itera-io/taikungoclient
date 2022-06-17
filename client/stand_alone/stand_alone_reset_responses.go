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

// StandAloneResetReader is a Reader for the StandAloneReset structure.
type StandAloneResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneResetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneResetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneResetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneResetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneResetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneResetOK creates a StandAloneResetOK with default headers values
func NewStandAloneResetOK() *StandAloneResetOK {
	return &StandAloneResetOK{}
}

/* StandAloneResetOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneResetOK struct {
	Payload models.Unit
}

func (o *StandAloneResetOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/reset][%d] standAloneResetOK  %+v", 200, o.Payload)
}
func (o *StandAloneResetOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneResetBadRequest creates a StandAloneResetBadRequest with default headers values
func NewStandAloneResetBadRequest() *StandAloneResetBadRequest {
	return &StandAloneResetBadRequest{}
}

/* StandAloneResetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneResetBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneResetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/reset][%d] standAloneResetBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneResetBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneResetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneResetUnauthorized creates a StandAloneResetUnauthorized with default headers values
func NewStandAloneResetUnauthorized() *StandAloneResetUnauthorized {
	return &StandAloneResetUnauthorized{}
}

/* StandAloneResetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneResetUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneResetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/reset][%d] standAloneResetUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneResetUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneResetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneResetForbidden creates a StandAloneResetForbidden with default headers values
func NewStandAloneResetForbidden() *StandAloneResetForbidden {
	return &StandAloneResetForbidden{}
}

/* StandAloneResetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneResetForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneResetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/reset][%d] standAloneResetForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneResetForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneResetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneResetNotFound creates a StandAloneResetNotFound with default headers values
func NewStandAloneResetNotFound() *StandAloneResetNotFound {
	return &StandAloneResetNotFound{}
}

/* StandAloneResetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneResetNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneResetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/reset][%d] standAloneResetNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneResetNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneResetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneResetInternalServerError creates a StandAloneResetInternalServerError with default headers values
func NewStandAloneResetInternalServerError() *StandAloneResetInternalServerError {
	return &StandAloneResetInternalServerError{}
}

/* StandAloneResetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneResetInternalServerError struct {
}

func (o *StandAloneResetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/reset][%d] standAloneResetInternalServerError ", 500)
}

func (o *StandAloneResetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
