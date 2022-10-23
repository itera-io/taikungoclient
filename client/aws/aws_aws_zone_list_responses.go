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

// AwsAwsZoneListReader is a Reader for the AwsAwsZoneList structure.
type AwsAwsZoneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsAwsZoneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsAwsZoneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsAwsZoneListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsAwsZoneListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsAwsZoneListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsAwsZoneListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsAwsZoneListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsAwsZoneListOK creates a AwsAwsZoneListOK with default headers values
func NewAwsAwsZoneListOK() *AwsAwsZoneListOK {
	return &AwsAwsZoneListOK{}
}

/*
AwsAwsZoneListOK describes a response with status code 200, with default header values.

Success
*/
type AwsAwsZoneListOK struct {
	Payload *models.AzResult
}

// IsSuccess returns true when this aws aws zone list o k response has a 2xx status code
func (o *AwsAwsZoneListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws aws zone list o k response has a 3xx status code
func (o *AwsAwsZoneListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws zone list o k response has a 4xx status code
func (o *AwsAwsZoneListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws aws zone list o k response has a 5xx status code
func (o *AwsAwsZoneListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws zone list o k response a status code equal to that given
func (o *AwsAwsZoneListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AwsAwsZoneListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListOK  %+v", 200, o.Payload)
}

func (o *AwsAwsZoneListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListOK  %+v", 200, o.Payload)
}

func (o *AwsAwsZoneListOK) GetPayload() *models.AzResult {
	return o.Payload
}

func (o *AwsAwsZoneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListBadRequest creates a AwsAwsZoneListBadRequest with default headers values
func NewAwsAwsZoneListBadRequest() *AwsAwsZoneListBadRequest {
	return &AwsAwsZoneListBadRequest{}
}

/*
AwsAwsZoneListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsAwsZoneListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this aws aws zone list bad request response has a 2xx status code
func (o *AwsAwsZoneListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws zone list bad request response has a 3xx status code
func (o *AwsAwsZoneListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws zone list bad request response has a 4xx status code
func (o *AwsAwsZoneListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws zone list bad request response has a 5xx status code
func (o *AwsAwsZoneListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws zone list bad request response a status code equal to that given
func (o *AwsAwsZoneListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AwsAwsZoneListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListBadRequest  %+v", 400, o.Payload)
}

func (o *AwsAwsZoneListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListBadRequest  %+v", 400, o.Payload)
}

func (o *AwsAwsZoneListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AwsAwsZoneListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListUnauthorized creates a AwsAwsZoneListUnauthorized with default headers values
func NewAwsAwsZoneListUnauthorized() *AwsAwsZoneListUnauthorized {
	return &AwsAwsZoneListUnauthorized{}
}

/*
AwsAwsZoneListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsAwsZoneListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws zone list unauthorized response has a 2xx status code
func (o *AwsAwsZoneListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws zone list unauthorized response has a 3xx status code
func (o *AwsAwsZoneListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws zone list unauthorized response has a 4xx status code
func (o *AwsAwsZoneListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws zone list unauthorized response has a 5xx status code
func (o *AwsAwsZoneListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws zone list unauthorized response a status code equal to that given
func (o *AwsAwsZoneListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AwsAwsZoneListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsAwsZoneListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsAwsZoneListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListForbidden creates a AwsAwsZoneListForbidden with default headers values
func NewAwsAwsZoneListForbidden() *AwsAwsZoneListForbidden {
	return &AwsAwsZoneListForbidden{}
}

/*
AwsAwsZoneListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsAwsZoneListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws zone list forbidden response has a 2xx status code
func (o *AwsAwsZoneListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws zone list forbidden response has a 3xx status code
func (o *AwsAwsZoneListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws zone list forbidden response has a 4xx status code
func (o *AwsAwsZoneListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws zone list forbidden response has a 5xx status code
func (o *AwsAwsZoneListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws zone list forbidden response a status code equal to that given
func (o *AwsAwsZoneListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AwsAwsZoneListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListForbidden  %+v", 403, o.Payload)
}

func (o *AwsAwsZoneListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListForbidden  %+v", 403, o.Payload)
}

func (o *AwsAwsZoneListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListNotFound creates a AwsAwsZoneListNotFound with default headers values
func NewAwsAwsZoneListNotFound() *AwsAwsZoneListNotFound {
	return &AwsAwsZoneListNotFound{}
}

/*
AwsAwsZoneListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsAwsZoneListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws aws zone list not found response has a 2xx status code
func (o *AwsAwsZoneListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws zone list not found response has a 3xx status code
func (o *AwsAwsZoneListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws zone list not found response has a 4xx status code
func (o *AwsAwsZoneListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws aws zone list not found response has a 5xx status code
func (o *AwsAwsZoneListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this aws aws zone list not found response a status code equal to that given
func (o *AwsAwsZoneListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AwsAwsZoneListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListNotFound  %+v", 404, o.Payload)
}

func (o *AwsAwsZoneListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListNotFound  %+v", 404, o.Payload)
}

func (o *AwsAwsZoneListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListInternalServerError creates a AwsAwsZoneListInternalServerError with default headers values
func NewAwsAwsZoneListInternalServerError() *AwsAwsZoneListInternalServerError {
	return &AwsAwsZoneListInternalServerError{}
}

/*
AwsAwsZoneListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsAwsZoneListInternalServerError struct {
}

// IsSuccess returns true when this aws aws zone list internal server error response has a 2xx status code
func (o *AwsAwsZoneListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws aws zone list internal server error response has a 3xx status code
func (o *AwsAwsZoneListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws aws zone list internal server error response has a 4xx status code
func (o *AwsAwsZoneListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws aws zone list internal server error response has a 5xx status code
func (o *AwsAwsZoneListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aws aws zone list internal server error response a status code equal to that given
func (o *AwsAwsZoneListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AwsAwsZoneListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListInternalServerError ", 500)
}

func (o *AwsAwsZoneListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListInternalServerError ", 500)
}

func (o *AwsAwsZoneListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
