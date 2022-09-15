// Code generated by go-swagger; DO NOT EDIT.

package slack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SlackUpdateReader is a Reader for the SlackUpdate structure.
type SlackUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SlackUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSlackUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSlackUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSlackUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSlackUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSlackUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSlackUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSlackUpdateOK creates a SlackUpdateOK with default headers values
func NewSlackUpdateOK() *SlackUpdateOK {
	return &SlackUpdateOK{}
}

/*
SlackUpdateOK describes a response with status code 200, with default header values.

Success
*/
type SlackUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this slack update o k response has a 2xx status code
func (o *SlackUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this slack update o k response has a 3xx status code
func (o *SlackUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack update o k response has a 4xx status code
func (o *SlackUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack update o k response has a 5xx status code
func (o *SlackUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this slack update o k response a status code equal to that given
func (o *SlackUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SlackUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateOK  %+v", 200, o.Payload)
}

func (o *SlackUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateOK  %+v", 200, o.Payload)
}

func (o *SlackUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *SlackUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackUpdateBadRequest creates a SlackUpdateBadRequest with default headers values
func NewSlackUpdateBadRequest() *SlackUpdateBadRequest {
	return &SlackUpdateBadRequest{}
}

/*
SlackUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SlackUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this slack update bad request response has a 2xx status code
func (o *SlackUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack update bad request response has a 3xx status code
func (o *SlackUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack update bad request response has a 4xx status code
func (o *SlackUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack update bad request response has a 5xx status code
func (o *SlackUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this slack update bad request response a status code equal to that given
func (o *SlackUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SlackUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SlackUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SlackUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SlackUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackUpdateUnauthorized creates a SlackUpdateUnauthorized with default headers values
func NewSlackUpdateUnauthorized() *SlackUpdateUnauthorized {
	return &SlackUpdateUnauthorized{}
}

/*
SlackUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SlackUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this slack update unauthorized response has a 2xx status code
func (o *SlackUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack update unauthorized response has a 3xx status code
func (o *SlackUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack update unauthorized response has a 4xx status code
func (o *SlackUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack update unauthorized response has a 5xx status code
func (o *SlackUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this slack update unauthorized response a status code equal to that given
func (o *SlackUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SlackUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SlackUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackUpdateForbidden creates a SlackUpdateForbidden with default headers values
func NewSlackUpdateForbidden() *SlackUpdateForbidden {
	return &SlackUpdateForbidden{}
}

/*
SlackUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SlackUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this slack update forbidden response has a 2xx status code
func (o *SlackUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack update forbidden response has a 3xx status code
func (o *SlackUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack update forbidden response has a 4xx status code
func (o *SlackUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack update forbidden response has a 5xx status code
func (o *SlackUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this slack update forbidden response a status code equal to that given
func (o *SlackUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SlackUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SlackUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SlackUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SlackUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackUpdateNotFound creates a SlackUpdateNotFound with default headers values
func NewSlackUpdateNotFound() *SlackUpdateNotFound {
	return &SlackUpdateNotFound{}
}

/*
SlackUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SlackUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this slack update not found response has a 2xx status code
func (o *SlackUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack update not found response has a 3xx status code
func (o *SlackUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack update not found response has a 4xx status code
func (o *SlackUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack update not found response has a 5xx status code
func (o *SlackUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this slack update not found response a status code equal to that given
func (o *SlackUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SlackUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SlackUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SlackUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SlackUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackUpdateInternalServerError creates a SlackUpdateInternalServerError with default headers values
func NewSlackUpdateInternalServerError() *SlackUpdateInternalServerError {
	return &SlackUpdateInternalServerError{}
}

/*
SlackUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SlackUpdateInternalServerError struct {
}

// IsSuccess returns true when this slack update internal server error response has a 2xx status code
func (o *SlackUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack update internal server error response has a 3xx status code
func (o *SlackUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack update internal server error response has a 4xx status code
func (o *SlackUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack update internal server error response has a 5xx status code
func (o *SlackUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this slack update internal server error response a status code equal to that given
func (o *SlackUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SlackUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateInternalServerError ", 500)
}

func (o *SlackUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Slack/update/{id}][%d] slackUpdateInternalServerError ", 500)
}

func (o *SlackUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
