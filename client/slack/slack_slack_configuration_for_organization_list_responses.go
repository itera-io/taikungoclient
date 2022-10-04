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

// SlackSlackConfigurationForOrganizationListReader is a Reader for the SlackSlackConfigurationForOrganizationList structure.
type SlackSlackConfigurationForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SlackSlackConfigurationForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSlackSlackConfigurationForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSlackSlackConfigurationForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSlackSlackConfigurationForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSlackSlackConfigurationForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSlackSlackConfigurationForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSlackSlackConfigurationForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSlackSlackConfigurationForOrganizationListOK creates a SlackSlackConfigurationForOrganizationListOK with default headers values
func NewSlackSlackConfigurationForOrganizationListOK() *SlackSlackConfigurationForOrganizationListOK {
	return &SlackSlackConfigurationForOrganizationListOK{}
}

/*
SlackSlackConfigurationForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type SlackSlackConfigurationForOrganizationListOK struct {
	Payload []*models.SlackConfigurationEntity
}

// IsSuccess returns true when this slack slack configuration for organization list o k response has a 2xx status code
func (o *SlackSlackConfigurationForOrganizationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this slack slack configuration for organization list o k response has a 3xx status code
func (o *SlackSlackConfigurationForOrganizationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack slack configuration for organization list o k response has a 4xx status code
func (o *SlackSlackConfigurationForOrganizationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack slack configuration for organization list o k response has a 5xx status code
func (o *SlackSlackConfigurationForOrganizationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this slack slack configuration for organization list o k response a status code equal to that given
func (o *SlackSlackConfigurationForOrganizationListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SlackSlackConfigurationForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListOK) GetPayload() []*models.SlackConfigurationEntity {
	return o.Payload
}

func (o *SlackSlackConfigurationForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackSlackConfigurationForOrganizationListBadRequest creates a SlackSlackConfigurationForOrganizationListBadRequest with default headers values
func NewSlackSlackConfigurationForOrganizationListBadRequest() *SlackSlackConfigurationForOrganizationListBadRequest {
	return &SlackSlackConfigurationForOrganizationListBadRequest{}
}

/*
SlackSlackConfigurationForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SlackSlackConfigurationForOrganizationListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this slack slack configuration for organization list bad request response has a 2xx status code
func (o *SlackSlackConfigurationForOrganizationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack slack configuration for organization list bad request response has a 3xx status code
func (o *SlackSlackConfigurationForOrganizationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack slack configuration for organization list bad request response has a 4xx status code
func (o *SlackSlackConfigurationForOrganizationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack slack configuration for organization list bad request response has a 5xx status code
func (o *SlackSlackConfigurationForOrganizationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this slack slack configuration for organization list bad request response a status code equal to that given
func (o *SlackSlackConfigurationForOrganizationListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SlackSlackConfigurationForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SlackSlackConfigurationForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackSlackConfigurationForOrganizationListUnauthorized creates a SlackSlackConfigurationForOrganizationListUnauthorized with default headers values
func NewSlackSlackConfigurationForOrganizationListUnauthorized() *SlackSlackConfigurationForOrganizationListUnauthorized {
	return &SlackSlackConfigurationForOrganizationListUnauthorized{}
}

/*
SlackSlackConfigurationForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SlackSlackConfigurationForOrganizationListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this slack slack configuration for organization list unauthorized response has a 2xx status code
func (o *SlackSlackConfigurationForOrganizationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack slack configuration for organization list unauthorized response has a 3xx status code
func (o *SlackSlackConfigurationForOrganizationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack slack configuration for organization list unauthorized response has a 4xx status code
func (o *SlackSlackConfigurationForOrganizationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack slack configuration for organization list unauthorized response has a 5xx status code
func (o *SlackSlackConfigurationForOrganizationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this slack slack configuration for organization list unauthorized response a status code equal to that given
func (o *SlackSlackConfigurationForOrganizationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SlackSlackConfigurationForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackSlackConfigurationForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackSlackConfigurationForOrganizationListForbidden creates a SlackSlackConfigurationForOrganizationListForbidden with default headers values
func NewSlackSlackConfigurationForOrganizationListForbidden() *SlackSlackConfigurationForOrganizationListForbidden {
	return &SlackSlackConfigurationForOrganizationListForbidden{}
}

/*
SlackSlackConfigurationForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SlackSlackConfigurationForOrganizationListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this slack slack configuration for organization list forbidden response has a 2xx status code
func (o *SlackSlackConfigurationForOrganizationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack slack configuration for organization list forbidden response has a 3xx status code
func (o *SlackSlackConfigurationForOrganizationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack slack configuration for organization list forbidden response has a 4xx status code
func (o *SlackSlackConfigurationForOrganizationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack slack configuration for organization list forbidden response has a 5xx status code
func (o *SlackSlackConfigurationForOrganizationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this slack slack configuration for organization list forbidden response a status code equal to that given
func (o *SlackSlackConfigurationForOrganizationListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SlackSlackConfigurationForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackSlackConfigurationForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackSlackConfigurationForOrganizationListNotFound creates a SlackSlackConfigurationForOrganizationListNotFound with default headers values
func NewSlackSlackConfigurationForOrganizationListNotFound() *SlackSlackConfigurationForOrganizationListNotFound {
	return &SlackSlackConfigurationForOrganizationListNotFound{}
}

/*
SlackSlackConfigurationForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SlackSlackConfigurationForOrganizationListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this slack slack configuration for organization list not found response has a 2xx status code
func (o *SlackSlackConfigurationForOrganizationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack slack configuration for organization list not found response has a 3xx status code
func (o *SlackSlackConfigurationForOrganizationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack slack configuration for organization list not found response has a 4xx status code
func (o *SlackSlackConfigurationForOrganizationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this slack slack configuration for organization list not found response has a 5xx status code
func (o *SlackSlackConfigurationForOrganizationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this slack slack configuration for organization list not found response a status code equal to that given
func (o *SlackSlackConfigurationForOrganizationListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SlackSlackConfigurationForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *SlackSlackConfigurationForOrganizationListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SlackSlackConfigurationForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSlackSlackConfigurationForOrganizationListInternalServerError creates a SlackSlackConfigurationForOrganizationListInternalServerError with default headers values
func NewSlackSlackConfigurationForOrganizationListInternalServerError() *SlackSlackConfigurationForOrganizationListInternalServerError {
	return &SlackSlackConfigurationForOrganizationListInternalServerError{}
}

/*
SlackSlackConfigurationForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SlackSlackConfigurationForOrganizationListInternalServerError struct {
}

// IsSuccess returns true when this slack slack configuration for organization list internal server error response has a 2xx status code
func (o *SlackSlackConfigurationForOrganizationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this slack slack configuration for organization list internal server error response has a 3xx status code
func (o *SlackSlackConfigurationForOrganizationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this slack slack configuration for organization list internal server error response has a 4xx status code
func (o *SlackSlackConfigurationForOrganizationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this slack slack configuration for organization list internal server error response has a 5xx status code
func (o *SlackSlackConfigurationForOrganizationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this slack slack configuration for organization list internal server error response a status code equal to that given
func (o *SlackSlackConfigurationForOrganizationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SlackSlackConfigurationForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListInternalServerError ", 500)
}

func (o *SlackSlackConfigurationForOrganizationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Slack/list][%d] slackSlackConfigurationForOrganizationListInternalServerError ", 500)
}

func (o *SlackSlackConfigurationForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
