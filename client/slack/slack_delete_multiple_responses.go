// Code generated by go-swagger; DO NOT EDIT.

package slack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SlackDeleteMultipleReader is a Reader for the SlackDeleteMultiple structure.
type SlackDeleteMultipleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SlackDeleteMultipleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSlackDeleteMultipleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSlackDeleteMultipleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSlackDeleteMultipleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSlackDeleteMultipleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSlackDeleteMultipleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSlackDeleteMultipleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSlackDeleteMultipleOK creates a SlackDeleteMultipleOK with default headers values
func NewSlackDeleteMultipleOK() *SlackDeleteMultipleOK {
	return &SlackDeleteMultipleOK{}
}

/*
SlackDeleteMultipleOK describes a response with status code 200, with default header values.

Success
*/
type SlackDeleteMultipleOK struct {
}

// IsSuccess returns true when this slack delete multiple o k response has a 2xx status code
func (o *SlackDeleteMultipleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this slack delete multiple o k response has a 3xx status code
func (o *SlackDeleteMultipleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack delete multiple o k response has a 4xx status code
func (o *SlackDeleteMultipleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack delete multiple o k response has a 5xx status code
func (o *SlackDeleteMultipleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this slack delete multiple o k response a status code equal to that given
func (o *SlackDeleteMultipleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the slack delete multiple o k response
func (o *SlackDeleteMultipleOK) Code() int {
	return 200
}

func (o *SlackDeleteMultipleOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleOK ", 200)
}

func (o *SlackDeleteMultipleOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleOK ", 200)
}

func (o *SlackDeleteMultipleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSlackDeleteMultipleBadRequest creates a SlackDeleteMultipleBadRequest with default headers values
func NewSlackDeleteMultipleBadRequest() *SlackDeleteMultipleBadRequest {
	return &SlackDeleteMultipleBadRequest{}
}

/*
SlackDeleteMultipleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SlackDeleteMultipleBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this slack delete multiple bad request response has a 2xx status code
func (o *SlackDeleteMultipleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack delete multiple bad request response has a 3xx status code
func (o *SlackDeleteMultipleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack delete multiple bad request response has a 4xx status code
func (o *SlackDeleteMultipleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack delete multiple bad request response has a 5xx status code
func (o *SlackDeleteMultipleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this slack delete multiple bad request response a status code equal to that given
func (o *SlackDeleteMultipleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the slack delete multiple bad request response
func (o *SlackDeleteMultipleBadRequest) Code() int {
	return 400
}

func (o *SlackDeleteMultipleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleBadRequest  %+v", 400, o.Payload)
}

func (o *SlackDeleteMultipleBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleBadRequest  %+v", 400, o.Payload)
}

func (o *SlackDeleteMultipleBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackDeleteMultipleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackDeleteMultipleUnauthorized creates a SlackDeleteMultipleUnauthorized with default headers values
func NewSlackDeleteMultipleUnauthorized() *SlackDeleteMultipleUnauthorized {
	return &SlackDeleteMultipleUnauthorized{}
}

/*
SlackDeleteMultipleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SlackDeleteMultipleUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this slack delete multiple unauthorized response has a 2xx status code
func (o *SlackDeleteMultipleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack delete multiple unauthorized response has a 3xx status code
func (o *SlackDeleteMultipleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack delete multiple unauthorized response has a 4xx status code
func (o *SlackDeleteMultipleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack delete multiple unauthorized response has a 5xx status code
func (o *SlackDeleteMultipleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this slack delete multiple unauthorized response a status code equal to that given
func (o *SlackDeleteMultipleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the slack delete multiple unauthorized response
func (o *SlackDeleteMultipleUnauthorized) Code() int {
	return 401
}

func (o *SlackDeleteMultipleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackDeleteMultipleUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackDeleteMultipleUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackDeleteMultipleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackDeleteMultipleForbidden creates a SlackDeleteMultipleForbidden with default headers values
func NewSlackDeleteMultipleForbidden() *SlackDeleteMultipleForbidden {
	return &SlackDeleteMultipleForbidden{}
}

/*
SlackDeleteMultipleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SlackDeleteMultipleForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this slack delete multiple forbidden response has a 2xx status code
func (o *SlackDeleteMultipleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack delete multiple forbidden response has a 3xx status code
func (o *SlackDeleteMultipleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack delete multiple forbidden response has a 4xx status code
func (o *SlackDeleteMultipleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack delete multiple forbidden response has a 5xx status code
func (o *SlackDeleteMultipleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this slack delete multiple forbidden response a status code equal to that given
func (o *SlackDeleteMultipleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the slack delete multiple forbidden response
func (o *SlackDeleteMultipleForbidden) Code() int {
	return 403
}

func (o *SlackDeleteMultipleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleForbidden  %+v", 403, o.Payload)
}

func (o *SlackDeleteMultipleForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleForbidden  %+v", 403, o.Payload)
}

func (o *SlackDeleteMultipleForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackDeleteMultipleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackDeleteMultipleNotFound creates a SlackDeleteMultipleNotFound with default headers values
func NewSlackDeleteMultipleNotFound() *SlackDeleteMultipleNotFound {
	return &SlackDeleteMultipleNotFound{}
}

/*
SlackDeleteMultipleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SlackDeleteMultipleNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this slack delete multiple not found response has a 2xx status code
func (o *SlackDeleteMultipleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack delete multiple not found response has a 3xx status code
func (o *SlackDeleteMultipleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack delete multiple not found response has a 4xx status code
func (o *SlackDeleteMultipleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack delete multiple not found response has a 5xx status code
func (o *SlackDeleteMultipleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this slack delete multiple not found response a status code equal to that given
func (o *SlackDeleteMultipleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the slack delete multiple not found response
func (o *SlackDeleteMultipleNotFound) Code() int {
	return 404
}

func (o *SlackDeleteMultipleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleNotFound  %+v", 404, o.Payload)
}

func (o *SlackDeleteMultipleNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleNotFound  %+v", 404, o.Payload)
}

func (o *SlackDeleteMultipleNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackDeleteMultipleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackDeleteMultipleInternalServerError creates a SlackDeleteMultipleInternalServerError with default headers values
func NewSlackDeleteMultipleInternalServerError() *SlackDeleteMultipleInternalServerError {
	return &SlackDeleteMultipleInternalServerError{}
}

/*
SlackDeleteMultipleInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SlackDeleteMultipleInternalServerError struct {
}

// IsSuccess returns true when this slack delete multiple internal server error response has a 2xx status code
func (o *SlackDeleteMultipleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack delete multiple internal server error response has a 3xx status code
func (o *SlackDeleteMultipleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack delete multiple internal server error response has a 4xx status code
func (o *SlackDeleteMultipleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack delete multiple internal server error response has a 5xx status code
func (o *SlackDeleteMultipleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this slack delete multiple internal server error response a status code equal to that given
func (o *SlackDeleteMultipleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the slack delete multiple internal server error response
func (o *SlackDeleteMultipleInternalServerError) Code() int {
	return 500
}

func (o *SlackDeleteMultipleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleInternalServerError ", 500)
}

func (o *SlackDeleteMultipleInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/delete-multiple][%d] slackDeleteMultipleInternalServerError ", 500)
}

func (o *SlackDeleteMultipleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
