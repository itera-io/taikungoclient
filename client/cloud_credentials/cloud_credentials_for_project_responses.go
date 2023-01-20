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

// CloudCredentialsForProjectReader is a Reader for the CloudCredentialsForProject structure.
type CloudCredentialsForProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsForProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsForProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsForProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsForProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsForProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsForProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsForProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsForProjectOK creates a CloudCredentialsForProjectOK with default headers values
func NewCloudCredentialsForProjectOK() *CloudCredentialsForProjectOK {
	return &CloudCredentialsForProjectOK{}
}

/*
CloudCredentialsForProjectOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsForProjectOK struct {
	Payload *models.CredentialsForProjectList
}

// IsSuccess returns true when this cloud credentials for project o k response has a 2xx status code
func (o *CloudCredentialsForProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud credentials for project o k response has a 3xx status code
func (o *CloudCredentialsForProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for project o k response has a 4xx status code
func (o *CloudCredentialsForProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials for project o k response has a 5xx status code
func (o *CloudCredentialsForProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for project o k response a status code equal to that given
func (o *CloudCredentialsForProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud credentials for project o k response
func (o *CloudCredentialsForProjectOK) Code() int {
	return 200
}

func (o *CloudCredentialsForProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsForProjectOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsForProjectOK) GetPayload() *models.CredentialsForProjectList {
	return o.Payload
}

func (o *CloudCredentialsForProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsForProjectList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForProjectBadRequest creates a CloudCredentialsForProjectBadRequest with default headers values
func NewCloudCredentialsForProjectBadRequest() *CloudCredentialsForProjectBadRequest {
	return &CloudCredentialsForProjectBadRequest{}
}

/*
CloudCredentialsForProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsForProjectBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for project bad request response has a 2xx status code
func (o *CloudCredentialsForProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for project bad request response has a 3xx status code
func (o *CloudCredentialsForProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for project bad request response has a 4xx status code
func (o *CloudCredentialsForProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for project bad request response has a 5xx status code
func (o *CloudCredentialsForProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for project bad request response a status code equal to that given
func (o *CloudCredentialsForProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cloud credentials for project bad request response
func (o *CloudCredentialsForProjectBadRequest) Code() int {
	return 400
}

func (o *CloudCredentialsForProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsForProjectBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsForProjectBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForProjectUnauthorized creates a CloudCredentialsForProjectUnauthorized with default headers values
func NewCloudCredentialsForProjectUnauthorized() *CloudCredentialsForProjectUnauthorized {
	return &CloudCredentialsForProjectUnauthorized{}
}

/*
CloudCredentialsForProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsForProjectUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for project unauthorized response has a 2xx status code
func (o *CloudCredentialsForProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for project unauthorized response has a 3xx status code
func (o *CloudCredentialsForProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for project unauthorized response has a 4xx status code
func (o *CloudCredentialsForProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for project unauthorized response has a 5xx status code
func (o *CloudCredentialsForProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for project unauthorized response a status code equal to that given
func (o *CloudCredentialsForProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cloud credentials for project unauthorized response
func (o *CloudCredentialsForProjectUnauthorized) Code() int {
	return 401
}

func (o *CloudCredentialsForProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsForProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsForProjectUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForProjectForbidden creates a CloudCredentialsForProjectForbidden with default headers values
func NewCloudCredentialsForProjectForbidden() *CloudCredentialsForProjectForbidden {
	return &CloudCredentialsForProjectForbidden{}
}

/*
CloudCredentialsForProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsForProjectForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for project forbidden response has a 2xx status code
func (o *CloudCredentialsForProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for project forbidden response has a 3xx status code
func (o *CloudCredentialsForProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for project forbidden response has a 4xx status code
func (o *CloudCredentialsForProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for project forbidden response has a 5xx status code
func (o *CloudCredentialsForProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for project forbidden response a status code equal to that given
func (o *CloudCredentialsForProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cloud credentials for project forbidden response
func (o *CloudCredentialsForProjectForbidden) Code() int {
	return 403
}

func (o *CloudCredentialsForProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsForProjectForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsForProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForProjectNotFound creates a CloudCredentialsForProjectNotFound with default headers values
func NewCloudCredentialsForProjectNotFound() *CloudCredentialsForProjectNotFound {
	return &CloudCredentialsForProjectNotFound{}
}

/*
CloudCredentialsForProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsForProjectNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this cloud credentials for project not found response has a 2xx status code
func (o *CloudCredentialsForProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for project not found response has a 3xx status code
func (o *CloudCredentialsForProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for project not found response has a 4xx status code
func (o *CloudCredentialsForProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials for project not found response has a 5xx status code
func (o *CloudCredentialsForProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials for project not found response a status code equal to that given
func (o *CloudCredentialsForProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cloud credentials for project not found response
func (o *CloudCredentialsForProjectNotFound) Code() int {
	return 404
}

func (o *CloudCredentialsForProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsForProjectNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsForProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CloudCredentialsForProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsForProjectInternalServerError creates a CloudCredentialsForProjectInternalServerError with default headers values
func NewCloudCredentialsForProjectInternalServerError() *CloudCredentialsForProjectInternalServerError {
	return &CloudCredentialsForProjectInternalServerError{}
}

/*
CloudCredentialsForProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsForProjectInternalServerError struct {
}

// IsSuccess returns true when this cloud credentials for project internal server error response has a 2xx status code
func (o *CloudCredentialsForProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials for project internal server error response has a 3xx status code
func (o *CloudCredentialsForProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials for project internal server error response has a 4xx status code
func (o *CloudCredentialsForProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials for project internal server error response has a 5xx status code
func (o *CloudCredentialsForProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud credentials for project internal server error response a status code equal to that given
func (o *CloudCredentialsForProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cloud credentials for project internal server error response
func (o *CloudCredentialsForProjectInternalServerError) Code() int {
	return 500
}

func (o *CloudCredentialsForProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectInternalServerError ", 500)
}

func (o *CloudCredentialsForProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/details][%d] cloudCredentialsForProjectInternalServerError ", 500)
}

func (o *CloudCredentialsForProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
