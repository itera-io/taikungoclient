// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CloudCredentialsForCliReader is a Reader for the CloudCredentialsForCli structure.
type CloudCredentialsForCliReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsForCliReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsForCliOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsForCliBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsForCliUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsForCliForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsForCliNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsForCliInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsForCliOK creates a CloudCredentialsForCliOK with default headers values
func NewCloudCredentialsForCliOK() *CloudCredentialsForCliOK {
	return &CloudCredentialsForCliOK{}
}

/*
CloudCredentialsForCliOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsForCliOK struct {
	Payload *models.CredentialsForCli
}

// IsSuccess returns true when this cloud credentials for cli o k response has a 2xx status code
func (o *CloudCredentialsForCliOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud credentials for cli o k response has a 3xx status code
func (o *CloudCredentialsForCliOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for cli o k response has a 4xx status code
func (o *CloudCredentialsForCliOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials for cli o k response has a 5xx status code
func (o *CloudCredentialsForCliOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for cli o k response a status code equal to that given
func (o *CloudCredentialsForCliOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudCredentialsForCliOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsForCliOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsForCliOK) GetPayload() *models.CredentialsForCli {
	return o.Payload
}

func (o *CloudCredentialsForCliOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsForCli)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForCliBadRequest creates a CloudCredentialsForCliBadRequest with default headers values
func NewCloudCredentialsForCliBadRequest() *CloudCredentialsForCliBadRequest {
	return &CloudCredentialsForCliBadRequest{}
}

/*
CloudCredentialsForCliBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsForCliBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this cloud credentials for cli bad request response has a 2xx status code
func (o *CloudCredentialsForCliBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for cli bad request response has a 3xx status code
func (o *CloudCredentialsForCliBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for cli bad request response has a 4xx status code
func (o *CloudCredentialsForCliBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for cli bad request response has a 5xx status code
func (o *CloudCredentialsForCliBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for cli bad request response a status code equal to that given
func (o *CloudCredentialsForCliBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CloudCredentialsForCliBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsForCliBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsForCliBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsForCliBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForCliUnauthorized creates a CloudCredentialsForCliUnauthorized with default headers values
func NewCloudCredentialsForCliUnauthorized() *CloudCredentialsForCliUnauthorized {
	return &CloudCredentialsForCliUnauthorized{}
}

/*
CloudCredentialsForCliUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsForCliUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for cli unauthorized response has a 2xx status code
func (o *CloudCredentialsForCliUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for cli unauthorized response has a 3xx status code
func (o *CloudCredentialsForCliUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for cli unauthorized response has a 4xx status code
func (o *CloudCredentialsForCliUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for cli unauthorized response has a 5xx status code
func (o *CloudCredentialsForCliUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for cli unauthorized response a status code equal to that given
func (o *CloudCredentialsForCliUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloudCredentialsForCliUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsForCliUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsForCliUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForCliUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForCliForbidden creates a CloudCredentialsForCliForbidden with default headers values
func NewCloudCredentialsForCliForbidden() *CloudCredentialsForCliForbidden {
	return &CloudCredentialsForCliForbidden{}
}

/*
CloudCredentialsForCliForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsForCliForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for cli forbidden response has a 2xx status code
func (o *CloudCredentialsForCliForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for cli forbidden response has a 3xx status code
func (o *CloudCredentialsForCliForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for cli forbidden response has a 4xx status code
func (o *CloudCredentialsForCliForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for cli forbidden response has a 5xx status code
func (o *CloudCredentialsForCliForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for cli forbidden response a status code equal to that given
func (o *CloudCredentialsForCliForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloudCredentialsForCliForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsForCliForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsForCliForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForCliForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForCliNotFound creates a CloudCredentialsForCliNotFound with default headers values
func NewCloudCredentialsForCliNotFound() *CloudCredentialsForCliNotFound {
	return &CloudCredentialsForCliNotFound{}
}

/*
CloudCredentialsForCliNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsForCliNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for cli not found response has a 2xx status code
func (o *CloudCredentialsForCliNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for cli not found response has a 3xx status code
func (o *CloudCredentialsForCliNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for cli not found response has a 4xx status code
func (o *CloudCredentialsForCliNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for cli not found response has a 5xx status code
func (o *CloudCredentialsForCliNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for cli not found response a status code equal to that given
func (o *CloudCredentialsForCliNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloudCredentialsForCliNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsForCliNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsForCliNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForCliNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForCliInternalServerError creates a CloudCredentialsForCliInternalServerError with default headers values
func NewCloudCredentialsForCliInternalServerError() *CloudCredentialsForCliInternalServerError {
	return &CloudCredentialsForCliInternalServerError{}
}

/*
CloudCredentialsForCliInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsForCliInternalServerError struct {
}

// IsSuccess returns true when this cloud credentials for cli internal server error response has a 2xx status code
func (o *CloudCredentialsForCliInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for cli internal server error response has a 3xx status code
func (o *CloudCredentialsForCliInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for cli internal server error response has a 4xx status code
func (o *CloudCredentialsForCliInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials for cli internal server error response has a 5xx status code
func (o *CloudCredentialsForCliInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud credentials for cli internal server error response a status code equal to that given
func (o *CloudCredentialsForCliInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudCredentialsForCliInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliInternalServerError ", 500)
}

func (o *CloudCredentialsForCliInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/cli][%d] cloudCredentialsForCliInternalServerError ", 500)
}

func (o *CloudCredentialsForCliInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
