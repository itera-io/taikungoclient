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

// StandAloneListForPollerReader is a Reader for the StandAloneListForPoller structure.
type StandAloneListForPollerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneListForPollerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneListForPollerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneListForPollerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneListForPollerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneListForPollerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneListForPollerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneListForPollerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneListForPollerOK creates a StandAloneListForPollerOK with default headers values
func NewStandAloneListForPollerOK() *StandAloneListForPollerOK {
	return &StandAloneListForPollerOK{}
}

/* StandAloneListForPollerOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneListForPollerOK struct {
	Payload *models.StandaloneVmsListForPoller
}

func (o *StandAloneListForPollerOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerOK  %+v", 200, o.Payload)
}
func (o *StandAloneListForPollerOK) GetPayload() *models.StandaloneVmsListForPoller {
	return o.Payload
}

func (o *StandAloneListForPollerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandaloneVmsListForPoller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListForPollerBadRequest creates a StandAloneListForPollerBadRequest with default headers values
func NewStandAloneListForPollerBadRequest() *StandAloneListForPollerBadRequest {
	return &StandAloneListForPollerBadRequest{}
}

/* StandAloneListForPollerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneListForPollerBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneListForPollerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneListForPollerBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneListForPollerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListForPollerUnauthorized creates a StandAloneListForPollerUnauthorized with default headers values
func NewStandAloneListForPollerUnauthorized() *StandAloneListForPollerUnauthorized {
	return &StandAloneListForPollerUnauthorized{}
}

/* StandAloneListForPollerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneListForPollerUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneListForPollerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneListForPollerUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListForPollerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListForPollerForbidden creates a StandAloneListForPollerForbidden with default headers values
func NewStandAloneListForPollerForbidden() *StandAloneListForPollerForbidden {
	return &StandAloneListForPollerForbidden{}
}

/* StandAloneListForPollerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneListForPollerForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneListForPollerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneListForPollerForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListForPollerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListForPollerNotFound creates a StandAloneListForPollerNotFound with default headers values
func NewStandAloneListForPollerNotFound() *StandAloneListForPollerNotFound {
	return &StandAloneListForPollerNotFound{}
}

/* StandAloneListForPollerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneListForPollerNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneListForPollerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneListForPollerNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListForPollerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneListForPollerInternalServerError creates a StandAloneListForPollerInternalServerError with default headers values
func NewStandAloneListForPollerInternalServerError() *StandAloneListForPollerInternalServerError {
	return &StandAloneListForPollerInternalServerError{}
}

/* StandAloneListForPollerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneListForPollerInternalServerError struct {
}

func (o *StandAloneListForPollerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerInternalServerError ", 500)
}

func (o *StandAloneListForPollerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
