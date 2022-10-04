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

// SlackCreateReader is a Reader for the SlackCreate structure.
type SlackCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SlackCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSlackCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSlackCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSlackCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSlackCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSlackCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSlackCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSlackCreateOK creates a SlackCreateOK with default headers values
func NewSlackCreateOK() *SlackCreateOK {
	return &SlackCreateOK{}
}

/*
SlackCreateOK describes a response with status code 200, with default header values.

Success
*/
type SlackCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this slack create o k response has a 2xx status code
func (o *SlackCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this slack create o k response has a 3xx status code
func (o *SlackCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack create o k response has a 4xx status code
func (o *SlackCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack create o k response has a 5xx status code
func (o *SlackCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this slack create o k response a status code equal to that given
func (o *SlackCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SlackCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateOK  %+v", 200, o.Payload)
}

func (o *SlackCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateOK  %+v", 200, o.Payload)
}

func (o *SlackCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *SlackCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackCreateBadRequest creates a SlackCreateBadRequest with default headers values
func NewSlackCreateBadRequest() *SlackCreateBadRequest {
	return &SlackCreateBadRequest{}
}

/*
SlackCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SlackCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this slack create bad request response has a 2xx status code
func (o *SlackCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack create bad request response has a 3xx status code
func (o *SlackCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack create bad request response has a 4xx status code
func (o *SlackCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack create bad request response has a 5xx status code
func (o *SlackCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this slack create bad request response a status code equal to that given
func (o *SlackCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SlackCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SlackCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SlackCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SlackCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackCreateUnauthorized creates a SlackCreateUnauthorized with default headers values
func NewSlackCreateUnauthorized() *SlackCreateUnauthorized {
	return &SlackCreateUnauthorized{}
}

/*
SlackCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SlackCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this slack create unauthorized response has a 2xx status code
func (o *SlackCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack create unauthorized response has a 3xx status code
func (o *SlackCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack create unauthorized response has a 4xx status code
func (o *SlackCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack create unauthorized response has a 5xx status code
func (o *SlackCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this slack create unauthorized response a status code equal to that given
func (o *SlackCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SlackCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackCreateForbidden creates a SlackCreateForbidden with default headers values
func NewSlackCreateForbidden() *SlackCreateForbidden {
	return &SlackCreateForbidden{}
}

/*
SlackCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SlackCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this slack create forbidden response has a 2xx status code
func (o *SlackCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack create forbidden response has a 3xx status code
func (o *SlackCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack create forbidden response has a 4xx status code
func (o *SlackCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack create forbidden response has a 5xx status code
func (o *SlackCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this slack create forbidden response a status code equal to that given
func (o *SlackCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SlackCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateForbidden  %+v", 403, o.Payload)
}

func (o *SlackCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateForbidden  %+v", 403, o.Payload)
}

func (o *SlackCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackCreateNotFound creates a SlackCreateNotFound with default headers values
func NewSlackCreateNotFound() *SlackCreateNotFound {
	return &SlackCreateNotFound{}
}

/*
SlackCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SlackCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this slack create not found response has a 2xx status code
func (o *SlackCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack create not found response has a 3xx status code
func (o *SlackCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack create not found response has a 4xx status code
func (o *SlackCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack create not found response has a 5xx status code
func (o *SlackCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this slack create not found response a status code equal to that given
func (o *SlackCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SlackCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateNotFound  %+v", 404, o.Payload)
}

func (o *SlackCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateNotFound  %+v", 404, o.Payload)
}

func (o *SlackCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackCreateInternalServerError creates a SlackCreateInternalServerError with default headers values
func NewSlackCreateInternalServerError() *SlackCreateInternalServerError {
	return &SlackCreateInternalServerError{}
}

/*
SlackCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SlackCreateInternalServerError struct {
}

// IsSuccess returns true when this slack create internal server error response has a 2xx status code
func (o *SlackCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack create internal server error response has a 3xx status code
func (o *SlackCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack create internal server error response has a 4xx status code
func (o *SlackCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack create internal server error response has a 5xx status code
func (o *SlackCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this slack create internal server error response a status code equal to that given
func (o *SlackCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SlackCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateInternalServerError ", 500)
}

func (o *SlackCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/create][%d] slackCreateInternalServerError ", 500)
}

func (o *SlackCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
