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

/*
StandAloneListForPollerOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneListForPollerOK struct {
	Payload *models.StandaloneVmsListForPoller
}

// IsSuccess returns true when this stand alone list for poller o k response has a 2xx status code
func (o *StandAloneListForPollerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone list for poller o k response has a 3xx status code
func (o *StandAloneListForPollerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list for poller o k response has a 4xx status code
func (o *StandAloneListForPollerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone list for poller o k response has a 5xx status code
func (o *StandAloneListForPollerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list for poller o k response a status code equal to that given
func (o *StandAloneListForPollerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone list for poller o k response
func (o *StandAloneListForPollerOK) Code() int {
	return 200
}

func (o *StandAloneListForPollerOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerOK  %+v", 200, o.Payload)
}

func (o *StandAloneListForPollerOK) String() string {
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

/*
StandAloneListForPollerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneListForPollerBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list for poller bad request response has a 2xx status code
func (o *StandAloneListForPollerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list for poller bad request response has a 3xx status code
func (o *StandAloneListForPollerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list for poller bad request response has a 4xx status code
func (o *StandAloneListForPollerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list for poller bad request response has a 5xx status code
func (o *StandAloneListForPollerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list for poller bad request response a status code equal to that given
func (o *StandAloneListForPollerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone list for poller bad request response
func (o *StandAloneListForPollerBadRequest) Code() int {
	return 400
}

func (o *StandAloneListForPollerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneListForPollerBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneListForPollerBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneListForPollerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

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

/*
StandAloneListForPollerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneListForPollerUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list for poller unauthorized response has a 2xx status code
func (o *StandAloneListForPollerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list for poller unauthorized response has a 3xx status code
func (o *StandAloneListForPollerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list for poller unauthorized response has a 4xx status code
func (o *StandAloneListForPollerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list for poller unauthorized response has a 5xx status code
func (o *StandAloneListForPollerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list for poller unauthorized response a status code equal to that given
func (o *StandAloneListForPollerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone list for poller unauthorized response
func (o *StandAloneListForPollerUnauthorized) Code() int {
	return 401
}

func (o *StandAloneListForPollerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneListForPollerUnauthorized) String() string {
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

/*
StandAloneListForPollerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneListForPollerForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list for poller forbidden response has a 2xx status code
func (o *StandAloneListForPollerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list for poller forbidden response has a 3xx status code
func (o *StandAloneListForPollerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list for poller forbidden response has a 4xx status code
func (o *StandAloneListForPollerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list for poller forbidden response has a 5xx status code
func (o *StandAloneListForPollerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list for poller forbidden response a status code equal to that given
func (o *StandAloneListForPollerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone list for poller forbidden response
func (o *StandAloneListForPollerForbidden) Code() int {
	return 403
}

func (o *StandAloneListForPollerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneListForPollerForbidden) String() string {
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

/*
StandAloneListForPollerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneListForPollerNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone list for poller not found response has a 2xx status code
func (o *StandAloneListForPollerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list for poller not found response has a 3xx status code
func (o *StandAloneListForPollerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list for poller not found response has a 4xx status code
func (o *StandAloneListForPollerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone list for poller not found response has a 5xx status code
func (o *StandAloneListForPollerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone list for poller not found response a status code equal to that given
func (o *StandAloneListForPollerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone list for poller not found response
func (o *StandAloneListForPollerNotFound) Code() int {
	return 404
}

func (o *StandAloneListForPollerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneListForPollerNotFound) String() string {
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

/*
StandAloneListForPollerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneListForPollerInternalServerError struct {
}

// IsSuccess returns true when this stand alone list for poller internal server error response has a 2xx status code
func (o *StandAloneListForPollerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone list for poller internal server error response has a 3xx status code
func (o *StandAloneListForPollerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone list for poller internal server error response has a 4xx status code
func (o *StandAloneListForPollerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone list for poller internal server error response has a 5xx status code
func (o *StandAloneListForPollerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone list for poller internal server error response a status code equal to that given
func (o *StandAloneListForPollerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone list for poller internal server error response
func (o *StandAloneListForPollerInternalServerError) Code() int {
	return 500
}

func (o *StandAloneListForPollerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerInternalServerError ", 500)
}

func (o *StandAloneListForPollerInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/forpoller][%d] standAloneListForPollerInternalServerError ", 500)
}

func (o *StandAloneListForPollerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
