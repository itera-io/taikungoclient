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

// CloudCredentialsLockManagerReader is a Reader for the CloudCredentialsLockManager structure.
type CloudCredentialsLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsLockManagerOK creates a CloudCredentialsLockManagerOK with default headers values
func NewCloudCredentialsLockManagerOK() *CloudCredentialsLockManagerOK {
	return &CloudCredentialsLockManagerOK{}
}

/*
CloudCredentialsLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsLockManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cloud credentials lock manager o k response has a 2xx status code
func (o *CloudCredentialsLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud credentials lock manager o k response has a 3xx status code
func (o *CloudCredentialsLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials lock manager o k response has a 4xx status code
func (o *CloudCredentialsLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials lock manager o k response has a 5xx status code
func (o *CloudCredentialsLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials lock manager o k response a status code equal to that given
func (o *CloudCredentialsLockManagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudCredentialsLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsLockManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CloudCredentialsLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsLockManagerBadRequest creates a CloudCredentialsLockManagerBadRequest with default headers values
func NewCloudCredentialsLockManagerBadRequest() *CloudCredentialsLockManagerBadRequest {
	return &CloudCredentialsLockManagerBadRequest{}
}

/*
CloudCredentialsLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsLockManagerBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this cloud credentials lock manager bad request response has a 2xx status code
func (o *CloudCredentialsLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials lock manager bad request response has a 3xx status code
func (o *CloudCredentialsLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials lock manager bad request response has a 4xx status code
func (o *CloudCredentialsLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials lock manager bad request response has a 5xx status code
func (o *CloudCredentialsLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials lock manager bad request response a status code equal to that given
func (o *CloudCredentialsLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CloudCredentialsLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsLockManagerBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CloudCredentialsLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsLockManagerUnauthorized creates a CloudCredentialsLockManagerUnauthorized with default headers values
func NewCloudCredentialsLockManagerUnauthorized() *CloudCredentialsLockManagerUnauthorized {
	return &CloudCredentialsLockManagerUnauthorized{}
}

/*
CloudCredentialsLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsLockManagerUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cloud credentials lock manager unauthorized response has a 2xx status code
func (o *CloudCredentialsLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials lock manager unauthorized response has a 3xx status code
func (o *CloudCredentialsLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials lock manager unauthorized response has a 4xx status code
func (o *CloudCredentialsLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials lock manager unauthorized response has a 5xx status code
func (o *CloudCredentialsLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials lock manager unauthorized response a status code equal to that given
func (o *CloudCredentialsLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloudCredentialsLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsLockManagerUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CloudCredentialsLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsLockManagerForbidden creates a CloudCredentialsLockManagerForbidden with default headers values
func NewCloudCredentialsLockManagerForbidden() *CloudCredentialsLockManagerForbidden {
	return &CloudCredentialsLockManagerForbidden{}
}

/*
CloudCredentialsLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsLockManagerForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cloud credentials lock manager forbidden response has a 2xx status code
func (o *CloudCredentialsLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials lock manager forbidden response has a 3xx status code
func (o *CloudCredentialsLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials lock manager forbidden response has a 4xx status code
func (o *CloudCredentialsLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials lock manager forbidden response has a 5xx status code
func (o *CloudCredentialsLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials lock manager forbidden response a status code equal to that given
func (o *CloudCredentialsLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloudCredentialsLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsLockManagerForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CloudCredentialsLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsLockManagerNotFound creates a CloudCredentialsLockManagerNotFound with default headers values
func NewCloudCredentialsLockManagerNotFound() *CloudCredentialsLockManagerNotFound {
	return &CloudCredentialsLockManagerNotFound{}
}

/*
CloudCredentialsLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsLockManagerNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cloud credentials lock manager not found response has a 2xx status code
func (o *CloudCredentialsLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials lock manager not found response has a 3xx status code
func (o *CloudCredentialsLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials lock manager not found response has a 4xx status code
func (o *CloudCredentialsLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials lock manager not found response has a 5xx status code
func (o *CloudCredentialsLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials lock manager not found response a status code equal to that given
func (o *CloudCredentialsLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloudCredentialsLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsLockManagerNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CloudCredentialsLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsLockManagerInternalServerError creates a CloudCredentialsLockManagerInternalServerError with default headers values
func NewCloudCredentialsLockManagerInternalServerError() *CloudCredentialsLockManagerInternalServerError {
	return &CloudCredentialsLockManagerInternalServerError{}
}

/*
CloudCredentialsLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsLockManagerInternalServerError struct {
}

// IsSuccess returns true when this cloud credentials lock manager internal server error response has a 2xx status code
func (o *CloudCredentialsLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials lock manager internal server error response has a 3xx status code
func (o *CloudCredentialsLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials lock manager internal server error response has a 4xx status code
func (o *CloudCredentialsLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials lock manager internal server error response has a 5xx status code
func (o *CloudCredentialsLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud credentials lock manager internal server error response a status code equal to that given
func (o *CloudCredentialsLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudCredentialsLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerInternalServerError ", 500)
}

func (o *CloudCredentialsLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CloudCredentials/lockmanager][%d] cloudCredentialsLockManagerInternalServerError ", 500)
}

func (o *CloudCredentialsLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
