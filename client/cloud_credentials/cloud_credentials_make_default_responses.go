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

// CloudCredentialsMakeDefaultReader is a Reader for the CloudCredentialsMakeDefault structure.
type CloudCredentialsMakeDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsMakeDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsMakeDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsMakeDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsMakeDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsMakeDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsMakeDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsMakeDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsMakeDefaultOK creates a CloudCredentialsMakeDefaultOK with default headers values
func NewCloudCredentialsMakeDefaultOK() *CloudCredentialsMakeDefaultOK {
	return &CloudCredentialsMakeDefaultOK{}
}

/*
CloudCredentialsMakeDefaultOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsMakeDefaultOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cloud credentials make default o k response has a 2xx status code
func (o *CloudCredentialsMakeDefaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud credentials make default o k response has a 3xx status code
func (o *CloudCredentialsMakeDefaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials make default o k response has a 4xx status code
func (o *CloudCredentialsMakeDefaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials make default o k response has a 5xx status code
func (o *CloudCredentialsMakeDefaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials make default o k response a status code equal to that given
func (o *CloudCredentialsMakeDefaultOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudCredentialsMakeDefaultOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsMakeDefaultOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsMakeDefaultOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CloudCredentialsMakeDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsMakeDefaultBadRequest creates a CloudCredentialsMakeDefaultBadRequest with default headers values
func NewCloudCredentialsMakeDefaultBadRequest() *CloudCredentialsMakeDefaultBadRequest {
	return &CloudCredentialsMakeDefaultBadRequest{}
}

/*
CloudCredentialsMakeDefaultBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsMakeDefaultBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this cloud credentials make default bad request response has a 2xx status code
func (o *CloudCredentialsMakeDefaultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials make default bad request response has a 3xx status code
func (o *CloudCredentialsMakeDefaultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials make default bad request response has a 4xx status code
func (o *CloudCredentialsMakeDefaultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials make default bad request response has a 5xx status code
func (o *CloudCredentialsMakeDefaultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials make default bad request response a status code equal to that given
func (o *CloudCredentialsMakeDefaultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CloudCredentialsMakeDefaultBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsMakeDefaultBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsMakeDefaultBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CloudCredentialsMakeDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsMakeDefaultUnauthorized creates a CloudCredentialsMakeDefaultUnauthorized with default headers values
func NewCloudCredentialsMakeDefaultUnauthorized() *CloudCredentialsMakeDefaultUnauthorized {
	return &CloudCredentialsMakeDefaultUnauthorized{}
}

/*
CloudCredentialsMakeDefaultUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsMakeDefaultUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cloud credentials make default unauthorized response has a 2xx status code
func (o *CloudCredentialsMakeDefaultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials make default unauthorized response has a 3xx status code
func (o *CloudCredentialsMakeDefaultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials make default unauthorized response has a 4xx status code
func (o *CloudCredentialsMakeDefaultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials make default unauthorized response has a 5xx status code
func (o *CloudCredentialsMakeDefaultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials make default unauthorized response a status code equal to that given
func (o *CloudCredentialsMakeDefaultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloudCredentialsMakeDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsMakeDefaultUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsMakeDefaultUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsMakeDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsMakeDefaultForbidden creates a CloudCredentialsMakeDefaultForbidden with default headers values
func NewCloudCredentialsMakeDefaultForbidden() *CloudCredentialsMakeDefaultForbidden {
	return &CloudCredentialsMakeDefaultForbidden{}
}

/*
CloudCredentialsMakeDefaultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsMakeDefaultForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cloud credentials make default forbidden response has a 2xx status code
func (o *CloudCredentialsMakeDefaultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials make default forbidden response has a 3xx status code
func (o *CloudCredentialsMakeDefaultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials make default forbidden response has a 4xx status code
func (o *CloudCredentialsMakeDefaultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials make default forbidden response has a 5xx status code
func (o *CloudCredentialsMakeDefaultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials make default forbidden response a status code equal to that given
func (o *CloudCredentialsMakeDefaultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloudCredentialsMakeDefaultForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsMakeDefaultForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsMakeDefaultForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsMakeDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsMakeDefaultNotFound creates a CloudCredentialsMakeDefaultNotFound with default headers values
func NewCloudCredentialsMakeDefaultNotFound() *CloudCredentialsMakeDefaultNotFound {
	return &CloudCredentialsMakeDefaultNotFound{}
}

/*
CloudCredentialsMakeDefaultNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsMakeDefaultNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cloud credentials make default not found response has a 2xx status code
func (o *CloudCredentialsMakeDefaultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials make default not found response has a 3xx status code
func (o *CloudCredentialsMakeDefaultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials make default not found response has a 4xx status code
func (o *CloudCredentialsMakeDefaultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials make default not found response has a 5xx status code
func (o *CloudCredentialsMakeDefaultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials make default not found response a status code equal to that given
func (o *CloudCredentialsMakeDefaultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloudCredentialsMakeDefaultNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsMakeDefaultNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsMakeDefaultNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsMakeDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsMakeDefaultInternalServerError creates a CloudCredentialsMakeDefaultInternalServerError with default headers values
func NewCloudCredentialsMakeDefaultInternalServerError() *CloudCredentialsMakeDefaultInternalServerError {
	return &CloudCredentialsMakeDefaultInternalServerError{}
}

/*
CloudCredentialsMakeDefaultInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsMakeDefaultInternalServerError struct {
}

// IsSuccess returns true when this cloud credentials make default internal server error response has a 2xx status code
func (o *CloudCredentialsMakeDefaultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials make default internal server error response has a 3xx status code
func (o *CloudCredentialsMakeDefaultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials make default internal server error response has a 4xx status code
func (o *CloudCredentialsMakeDefaultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials make default internal server error response has a 5xx status code
func (o *CloudCredentialsMakeDefaultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud credentials make default internal server error response a status code equal to that given
func (o *CloudCredentialsMakeDefaultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudCredentialsMakeDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultInternalServerError ", 500)
}

func (o *CloudCredentialsMakeDefaultInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/makedefault][%d] cloudCredentialsMakeDefaultInternalServerError ", 500)
}

func (o *CloudCredentialsMakeDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
