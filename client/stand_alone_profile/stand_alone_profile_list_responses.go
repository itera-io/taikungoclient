// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneProfileListReader is a Reader for the StandAloneProfileList structure.
type StandAloneProfileListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneProfileListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneProfileListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneProfileListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneProfileListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneProfileListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneProfileListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneProfileListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneProfileListOK creates a StandAloneProfileListOK with default headers values
func NewStandAloneProfileListOK() *StandAloneProfileListOK {
	return &StandAloneProfileListOK{}
}

/* StandAloneProfileListOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneProfileListOK struct {
	Payload *models.StandAloneProfiles
}

func (o *StandAloneProfileListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile][%d] standAloneProfileListOK  %+v", 200, o.Payload)
}
func (o *StandAloneProfileListOK) GetPayload() *models.StandAloneProfiles {
	return o.Payload
}

func (o *StandAloneProfileListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandAloneProfiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileListBadRequest creates a StandAloneProfileListBadRequest with default headers values
func NewStandAloneProfileListBadRequest() *StandAloneProfileListBadRequest {
	return &StandAloneProfileListBadRequest{}
}

/* StandAloneProfileListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneProfileListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneProfileListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile][%d] standAloneProfileListBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneProfileListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileListUnauthorized creates a StandAloneProfileListUnauthorized with default headers values
func NewStandAloneProfileListUnauthorized() *StandAloneProfileListUnauthorized {
	return &StandAloneProfileListUnauthorized{}
}

/* StandAloneProfileListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneProfileListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneProfileListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile][%d] standAloneProfileListUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneProfileListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileListForbidden creates a StandAloneProfileListForbidden with default headers values
func NewStandAloneProfileListForbidden() *StandAloneProfileListForbidden {
	return &StandAloneProfileListForbidden{}
}

/* StandAloneProfileListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneProfileListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneProfileListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile][%d] standAloneProfileListForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneProfileListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileListNotFound creates a StandAloneProfileListNotFound with default headers values
func NewStandAloneProfileListNotFound() *StandAloneProfileListNotFound {
	return &StandAloneProfileListNotFound{}
}

/* StandAloneProfileListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneProfileListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneProfileListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile][%d] standAloneProfileListNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneProfileListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileListInternalServerError creates a StandAloneProfileListInternalServerError with default headers values
func NewStandAloneProfileListInternalServerError() *StandAloneProfileListInternalServerError {
	return &StandAloneProfileListInternalServerError{}
}

/* StandAloneProfileListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneProfileListInternalServerError struct {
}

func (o *StandAloneProfileListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneProfile][%d] standAloneProfileListInternalServerError ", 500)
}

func (o *StandAloneProfileListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
