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

// AwsCreateReader is a Reader for the AwsCreate structure.
type AwsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsCreateOK creates a AwsCreateOK with default headers values
func NewAwsCreateOK() *AwsCreateOK {
	return &AwsCreateOK{}
}

/*
AwsCreateOK describes a response with status code 200, with default header values.

Success
*/
type AwsCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this aws create o k response has a 2xx status code
func (o *AwsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws create o k response has a 3xx status code
func (o *AwsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws create o k response has a 4xx status code
func (o *AwsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws create o k response has a 5xx status code
func (o *AwsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aws create o k response a status code equal to that given
func (o *AwsCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AwsCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateOK  %+v", 200, o.Payload)
}

func (o *AwsCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateOK  %+v", 200, o.Payload)
}

func (o *AwsCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *AwsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsCreateBadRequest creates a AwsCreateBadRequest with default headers values
func NewAwsCreateBadRequest() *AwsCreateBadRequest {
	return &AwsCreateBadRequest{}
}

/*
AwsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsCreateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this aws create bad request response has a 2xx status code
func (o *AwsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws create bad request response has a 3xx status code
func (o *AwsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws create bad request response has a 4xx status code
func (o *AwsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws create bad request response has a 5xx status code
func (o *AwsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aws create bad request response a status code equal to that given
func (o *AwsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AwsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AwsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AwsCreateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AwsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsCreateUnauthorized creates a AwsCreateUnauthorized with default headers values
func NewAwsCreateUnauthorized() *AwsCreateUnauthorized {
	return &AwsCreateUnauthorized{}
}

/*
AwsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsCreateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this aws create unauthorized response has a 2xx status code
func (o *AwsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws create unauthorized response has a 3xx status code
func (o *AwsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws create unauthorized response has a 4xx status code
func (o *AwsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws create unauthorized response has a 5xx status code
func (o *AwsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this aws create unauthorized response a status code equal to that given
func (o *AwsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AwsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AwsCreateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AwsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsCreateForbidden creates a AwsCreateForbidden with default headers values
func NewAwsCreateForbidden() *AwsCreateForbidden {
	return &AwsCreateForbidden{}
}

/*
AwsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsCreateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this aws create forbidden response has a 2xx status code
func (o *AwsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws create forbidden response has a 3xx status code
func (o *AwsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws create forbidden response has a 4xx status code
func (o *AwsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws create forbidden response has a 5xx status code
func (o *AwsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aws create forbidden response a status code equal to that given
func (o *AwsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AwsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateForbidden  %+v", 403, o.Payload)
}

func (o *AwsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateForbidden  %+v", 403, o.Payload)
}

func (o *AwsCreateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AwsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsCreateNotFound creates a AwsCreateNotFound with default headers values
func NewAwsCreateNotFound() *AwsCreateNotFound {
	return &AwsCreateNotFound{}
}

/*
AwsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsCreateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this aws create not found response has a 2xx status code
func (o *AwsCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws create not found response has a 3xx status code
func (o *AwsCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws create not found response has a 4xx status code
func (o *AwsCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this aws create not found response has a 5xx status code
func (o *AwsCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this aws create not found response a status code equal to that given
func (o *AwsCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AwsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateNotFound  %+v", 404, o.Payload)
}

func (o *AwsCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateNotFound  %+v", 404, o.Payload)
}

func (o *AwsCreateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AwsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsCreateInternalServerError creates a AwsCreateInternalServerError with default headers values
func NewAwsCreateInternalServerError() *AwsCreateInternalServerError {
	return &AwsCreateInternalServerError{}
}

/*
AwsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsCreateInternalServerError struct {
}

// IsSuccess returns true when this aws create internal server error response has a 2xx status code
func (o *AwsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aws create internal server error response has a 3xx status code
func (o *AwsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws create internal server error response has a 4xx status code
func (o *AwsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws create internal server error response has a 5xx status code
func (o *AwsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aws create internal server error response a status code equal to that given
func (o *AwsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AwsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateInternalServerError ", 500)
}

func (o *AwsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/create][%d] awsCreateInternalServerError ", 500)
}

func (o *AwsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
