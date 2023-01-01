// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AwsAwsOwnersReader is a Reader for the AwsAwsOwners structure.
type AwsAwsOwnersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsAwsOwnersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsAwsOwnersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsAwsOwnersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsAwsOwnersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsAwsOwnersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsAwsOwnersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsAwsOwnersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsAwsOwnersOK creates a AwsAwsOwnersOK with default headers values
func NewAwsAwsOwnersOK() *AwsAwsOwnersOK {
	return &AwsAwsOwnersOK{}
}

/*
AwsAwsOwnersOK describes a response with status code 200, with default header values.

Success
*/
type AwsAwsOwnersOK struct {
	Payload []*models.CommonStringBasedDropdownDto
}

// IsSuccess returns true when this aws aws owners o k response has a 2xx status code
func (o *AwsAwsOwnersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws aws owners o k response has a 3xx status code
func (o *AwsAwsOwnersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws owners o k response has a 4xx status code
func (o *AwsAwsOwnersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws aws owners o k response has a 5xx status code
func (o *AwsAwsOwnersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws owners o k response a status code equal to that given
func (o *AwsAwsOwnersOK) IsCode(code int) bool {
	return code == 200
}

func (o *AwsAwsOwnersOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersOK  %+v", 200, o.Payload)
}

func (o *AwsAwsOwnersOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersOK  %+v", 200, o.Payload)
}

func (o *AwsAwsOwnersOK) GetPayload() []*models.CommonStringBasedDropdownDto {
	return o.Payload
}

func (o *AwsAwsOwnersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsOwnersBadRequest creates a AwsAwsOwnersBadRequest with default headers values
func NewAwsAwsOwnersBadRequest() *AwsAwsOwnersBadRequest {
	return &AwsAwsOwnersBadRequest{}
}

/*
AwsAwsOwnersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsAwsOwnersBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws owners bad request response has a 2xx status code
func (o *AwsAwsOwnersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws owners bad request response has a 3xx status code
func (o *AwsAwsOwnersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws owners bad request response has a 4xx status code
func (o *AwsAwsOwnersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws owners bad request response has a 5xx status code
func (o *AwsAwsOwnersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws owners bad request response a status code equal to that given
func (o *AwsAwsOwnersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AwsAwsOwnersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersBadRequest  %+v", 400, o.Payload)
}

func (o *AwsAwsOwnersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersBadRequest  %+v", 400, o.Payload)
}

func (o *AwsAwsOwnersBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsOwnersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsOwnersUnauthorized creates a AwsAwsOwnersUnauthorized with default headers values
func NewAwsAwsOwnersUnauthorized() *AwsAwsOwnersUnauthorized {
	return &AwsAwsOwnersUnauthorized{}
}

/*
AwsAwsOwnersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsAwsOwnersUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws owners unauthorized response has a 2xx status code
func (o *AwsAwsOwnersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws owners unauthorized response has a 3xx status code
func (o *AwsAwsOwnersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws owners unauthorized response has a 4xx status code
func (o *AwsAwsOwnersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws owners unauthorized response has a 5xx status code
func (o *AwsAwsOwnersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws owners unauthorized response a status code equal to that given
func (o *AwsAwsOwnersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AwsAwsOwnersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsAwsOwnersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsAwsOwnersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsOwnersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsOwnersForbidden creates a AwsAwsOwnersForbidden with default headers values
func NewAwsAwsOwnersForbidden() *AwsAwsOwnersForbidden {
	return &AwsAwsOwnersForbidden{}
}

/*
AwsAwsOwnersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsAwsOwnersForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws owners forbidden response has a 2xx status code
func (o *AwsAwsOwnersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws owners forbidden response has a 3xx status code
func (o *AwsAwsOwnersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws owners forbidden response has a 4xx status code
func (o *AwsAwsOwnersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws owners forbidden response has a 5xx status code
func (o *AwsAwsOwnersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws owners forbidden response a status code equal to that given
func (o *AwsAwsOwnersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AwsAwsOwnersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersForbidden  %+v", 403, o.Payload)
}

func (o *AwsAwsOwnersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersForbidden  %+v", 403, o.Payload)
}

func (o *AwsAwsOwnersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsOwnersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsOwnersNotFound creates a AwsAwsOwnersNotFound with default headers values
func NewAwsAwsOwnersNotFound() *AwsAwsOwnersNotFound {
	return &AwsAwsOwnersNotFound{}
}

/*
AwsAwsOwnersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsAwsOwnersNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws owners not found response has a 2xx status code
func (o *AwsAwsOwnersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws owners not found response has a 3xx status code
func (o *AwsAwsOwnersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws owners not found response has a 4xx status code
func (o *AwsAwsOwnersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws owners not found response has a 5xx status code
func (o *AwsAwsOwnersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws owners not found response a status code equal to that given
func (o *AwsAwsOwnersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AwsAwsOwnersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersNotFound  %+v", 404, o.Payload)
}

func (o *AwsAwsOwnersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersNotFound  %+v", 404, o.Payload)
}

func (o *AwsAwsOwnersNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsOwnersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsOwnersInternalServerError creates a AwsAwsOwnersInternalServerError with default headers values
func NewAwsAwsOwnersInternalServerError() *AwsAwsOwnersInternalServerError {
	return &AwsAwsOwnersInternalServerError{}
}

/*
AwsAwsOwnersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsAwsOwnersInternalServerError struct {
}

// IsSuccess returns true when this aws aws owners internal server error response has a 2xx status code
func (o *AwsAwsOwnersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws owners internal server error response has a 3xx status code
func (o *AwsAwsOwnersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws owners internal server error response has a 4xx status code
func (o *AwsAwsOwnersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws aws owners internal server error response has a 5xx status code
func (o *AwsAwsOwnersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aws aws owners internal server error response a status code equal to that given
func (o *AwsAwsOwnersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AwsAwsOwnersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersInternalServerError ", 500)
}

func (o *AwsAwsOwnersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/owners][%d] awsAwsOwnersInternalServerError ", 500)
}

func (o *AwsAwsOwnersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
