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

// AwsListReader is a Reader for the AwsList structure.
type AwsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsListOK creates a AwsListOK with default headers values
func NewAwsListOK() *AwsListOK {
	return &AwsListOK{}
}

/*
AwsListOK describes a response with status code 200, with default header values.

Success
*/
type AwsListOK struct {
	Payload *models.AwsCredentialList
}

// IsSuccess returns true when this aws list o k response has a 2xx status code
func (o *AwsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws list o k response has a 3xx status code
func (o *AwsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws list o k response has a 4xx status code
func (o *AwsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws list o k response has a 5xx status code
func (o *AwsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aws list o k response a status code equal to that given
func (o *AwsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AwsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListOK  %+v", 200, o.Payload)
}

func (o *AwsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListOK  %+v", 200, o.Payload)
}

func (o *AwsListOK) GetPayload() *models.AwsCredentialList {
	return o.Payload
}

func (o *AwsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsCredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsListBadRequest creates a AwsListBadRequest with default headers values
func NewAwsListBadRequest() *AwsListBadRequest {
	return &AwsListBadRequest{}
}

/*
AwsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws list bad request response has a 2xx status code
func (o *AwsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws list bad request response has a 3xx status code
func (o *AwsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws list bad request response has a 4xx status code
func (o *AwsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws list bad request response has a 5xx status code
func (o *AwsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aws list bad request response a status code equal to that given
func (o *AwsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AwsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListBadRequest  %+v", 400, o.Payload)
}

func (o *AwsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListBadRequest  %+v", 400, o.Payload)
}

func (o *AwsListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsListUnauthorized creates a AwsListUnauthorized with default headers values
func NewAwsListUnauthorized() *AwsListUnauthorized {
	return &AwsListUnauthorized{}
}

/*
AwsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws list unauthorized response has a 2xx status code
func (o *AwsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws list unauthorized response has a 3xx status code
func (o *AwsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws list unauthorized response has a 4xx status code
func (o *AwsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws list unauthorized response has a 5xx status code
func (o *AwsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this aws list unauthorized response a status code equal to that given
func (o *AwsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AwsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsListForbidden creates a AwsListForbidden with default headers values
func NewAwsListForbidden() *AwsListForbidden {
	return &AwsListForbidden{}
}

/*
AwsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws list forbidden response has a 2xx status code
func (o *AwsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws list forbidden response has a 3xx status code
func (o *AwsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws list forbidden response has a 4xx status code
func (o *AwsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws list forbidden response has a 5xx status code
func (o *AwsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aws list forbidden response a status code equal to that given
func (o *AwsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AwsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListForbidden  %+v", 403, o.Payload)
}

func (o *AwsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListForbidden  %+v", 403, o.Payload)
}

func (o *AwsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsListNotFound creates a AwsListNotFound with default headers values
func NewAwsListNotFound() *AwsListNotFound {
	return &AwsListNotFound{}
}

/*
AwsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this aws list not found response has a 2xx status code
func (o *AwsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws list not found response has a 3xx status code
func (o *AwsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws list not found response has a 4xx status code
func (o *AwsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws list not found response has a 5xx status code
func (o *AwsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this aws list not found response a status code equal to that given
func (o *AwsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AwsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListNotFound  %+v", 404, o.Payload)
}

func (o *AwsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListNotFound  %+v", 404, o.Payload)
}

func (o *AwsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsListInternalServerError creates a AwsListInternalServerError with default headers values
func NewAwsListInternalServerError() *AwsListInternalServerError {
	return &AwsListInternalServerError{}
}

/*
AwsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsListInternalServerError struct {
}

// IsSuccess returns true when this aws list internal server error response has a 2xx status code
func (o *AwsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws list internal server error response has a 3xx status code
func (o *AwsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws list internal server error response has a 4xx status code
func (o *AwsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws list internal server error response has a 5xx status code
func (o *AwsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aws list internal server error response a status code equal to that given
func (o *AwsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AwsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListInternalServerError ", 500)
}

func (o *AwsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Aws/list][%d] awsListInternalServerError ", 500)
}

func (o *AwsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
