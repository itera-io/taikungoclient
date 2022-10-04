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

// SlackVerifySlackCredentialsReader is a Reader for the SlackVerifySlackCredentials structure.
type SlackVerifySlackCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SlackVerifySlackCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSlackVerifySlackCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSlackVerifySlackCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSlackVerifySlackCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSlackVerifySlackCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSlackVerifySlackCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSlackVerifySlackCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSlackVerifySlackCredentialsOK creates a SlackVerifySlackCredentialsOK with default headers values
func NewSlackVerifySlackCredentialsOK() *SlackVerifySlackCredentialsOK {
	return &SlackVerifySlackCredentialsOK{}
}

/*
SlackVerifySlackCredentialsOK describes a response with status code 200, with default header values.

Success
*/
type SlackVerifySlackCredentialsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this slack verify slack credentials o k response has a 2xx status code
func (o *SlackVerifySlackCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this slack verify slack credentials o k response has a 3xx status code
func (o *SlackVerifySlackCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack verify slack credentials o k response has a 4xx status code
func (o *SlackVerifySlackCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack verify slack credentials o k response has a 5xx status code
func (o *SlackVerifySlackCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this slack verify slack credentials o k response a status code equal to that given
func (o *SlackVerifySlackCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SlackVerifySlackCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsOK  %+v", 200, o.Payload)
}

func (o *SlackVerifySlackCredentialsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsOK  %+v", 200, o.Payload)
}

func (o *SlackVerifySlackCredentialsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *SlackVerifySlackCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackVerifySlackCredentialsBadRequest creates a SlackVerifySlackCredentialsBadRequest with default headers values
func NewSlackVerifySlackCredentialsBadRequest() *SlackVerifySlackCredentialsBadRequest {
	return &SlackVerifySlackCredentialsBadRequest{}
}

/*
SlackVerifySlackCredentialsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SlackVerifySlackCredentialsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this slack verify slack credentials bad request response has a 2xx status code
func (o *SlackVerifySlackCredentialsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack verify slack credentials bad request response has a 3xx status code
func (o *SlackVerifySlackCredentialsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack verify slack credentials bad request response has a 4xx status code
func (o *SlackVerifySlackCredentialsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack verify slack credentials bad request response has a 5xx status code
func (o *SlackVerifySlackCredentialsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this slack verify slack credentials bad request response a status code equal to that given
func (o *SlackVerifySlackCredentialsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SlackVerifySlackCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *SlackVerifySlackCredentialsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *SlackVerifySlackCredentialsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SlackVerifySlackCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackVerifySlackCredentialsUnauthorized creates a SlackVerifySlackCredentialsUnauthorized with default headers values
func NewSlackVerifySlackCredentialsUnauthorized() *SlackVerifySlackCredentialsUnauthorized {
	return &SlackVerifySlackCredentialsUnauthorized{}
}

/*
SlackVerifySlackCredentialsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SlackVerifySlackCredentialsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this slack verify slack credentials unauthorized response has a 2xx status code
func (o *SlackVerifySlackCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack verify slack credentials unauthorized response has a 3xx status code
func (o *SlackVerifySlackCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack verify slack credentials unauthorized response has a 4xx status code
func (o *SlackVerifySlackCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack verify slack credentials unauthorized response has a 5xx status code
func (o *SlackVerifySlackCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this slack verify slack credentials unauthorized response a status code equal to that given
func (o *SlackVerifySlackCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SlackVerifySlackCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackVerifySlackCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackVerifySlackCredentialsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackVerifySlackCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackVerifySlackCredentialsForbidden creates a SlackVerifySlackCredentialsForbidden with default headers values
func NewSlackVerifySlackCredentialsForbidden() *SlackVerifySlackCredentialsForbidden {
	return &SlackVerifySlackCredentialsForbidden{}
}

/*
SlackVerifySlackCredentialsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SlackVerifySlackCredentialsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this slack verify slack credentials forbidden response has a 2xx status code
func (o *SlackVerifySlackCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack verify slack credentials forbidden response has a 3xx status code
func (o *SlackVerifySlackCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack verify slack credentials forbidden response has a 4xx status code
func (o *SlackVerifySlackCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack verify slack credentials forbidden response has a 5xx status code
func (o *SlackVerifySlackCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this slack verify slack credentials forbidden response a status code equal to that given
func (o *SlackVerifySlackCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SlackVerifySlackCredentialsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *SlackVerifySlackCredentialsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *SlackVerifySlackCredentialsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackVerifySlackCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackVerifySlackCredentialsNotFound creates a SlackVerifySlackCredentialsNotFound with default headers values
func NewSlackVerifySlackCredentialsNotFound() *SlackVerifySlackCredentialsNotFound {
	return &SlackVerifySlackCredentialsNotFound{}
}

/*
SlackVerifySlackCredentialsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SlackVerifySlackCredentialsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this slack verify slack credentials not found response has a 2xx status code
func (o *SlackVerifySlackCredentialsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack verify slack credentials not found response has a 3xx status code
func (o *SlackVerifySlackCredentialsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack verify slack credentials not found response has a 4xx status code
func (o *SlackVerifySlackCredentialsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack verify slack credentials not found response has a 5xx status code
func (o *SlackVerifySlackCredentialsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this slack verify slack credentials not found response a status code equal to that given
func (o *SlackVerifySlackCredentialsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SlackVerifySlackCredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *SlackVerifySlackCredentialsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *SlackVerifySlackCredentialsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackVerifySlackCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackVerifySlackCredentialsInternalServerError creates a SlackVerifySlackCredentialsInternalServerError with default headers values
func NewSlackVerifySlackCredentialsInternalServerError() *SlackVerifySlackCredentialsInternalServerError {
	return &SlackVerifySlackCredentialsInternalServerError{}
}

/*
SlackVerifySlackCredentialsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SlackVerifySlackCredentialsInternalServerError struct {
}

// IsSuccess returns true when this slack verify slack credentials internal server error response has a 2xx status code
func (o *SlackVerifySlackCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack verify slack credentials internal server error response has a 3xx status code
func (o *SlackVerifySlackCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack verify slack credentials internal server error response has a 4xx status code
func (o *SlackVerifySlackCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack verify slack credentials internal server error response has a 5xx status code
func (o *SlackVerifySlackCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this slack verify slack credentials internal server error response a status code equal to that given
func (o *SlackVerifySlackCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SlackVerifySlackCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsInternalServerError ", 500)
}

func (o *SlackVerifySlackCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Slack/verify][%d] slackVerifySlackCredentialsInternalServerError ", 500)
}

func (o *SlackVerifySlackCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
