// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// S3CredentialsListReader is a Reader for the S3CredentialsList structure.
type S3CredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3CredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3CredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3CredentialsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS3CredentialsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS3CredentialsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS3CredentialsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewS3CredentialsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewS3CredentialsListOK creates a S3CredentialsListOK with default headers values
func NewS3CredentialsListOK() *S3CredentialsListOK {
	return &S3CredentialsListOK{}
}

/*
S3CredentialsListOK describes a response with status code 200, with default header values.

Success
*/
type S3CredentialsListOK struct {
	Payload *models.BackupCredentials
}

// IsSuccess returns true when this s3 credentials list o k response has a 2xx status code
func (o *S3CredentialsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 credentials list o k response has a 3xx status code
func (o *S3CredentialsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials list o k response has a 4xx status code
func (o *S3CredentialsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials list o k response has a 5xx status code
func (o *S3CredentialsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials list o k response a status code equal to that given
func (o *S3CredentialsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3CredentialsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsListOK) GetPayload() *models.BackupCredentials {
	return o.Payload
}

func (o *S3CredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsListBadRequest creates a S3CredentialsListBadRequest with default headers values
func NewS3CredentialsListBadRequest() *S3CredentialsListBadRequest {
	return &S3CredentialsListBadRequest{}
}

/*
S3CredentialsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type S3CredentialsListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this s3 credentials list bad request response has a 2xx status code
func (o *S3CredentialsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials list bad request response has a 3xx status code
func (o *S3CredentialsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials list bad request response has a 4xx status code
func (o *S3CredentialsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials list bad request response has a 5xx status code
func (o *S3CredentialsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials list bad request response a status code equal to that given
func (o *S3CredentialsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *S3CredentialsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *S3CredentialsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsListUnauthorized creates a S3CredentialsListUnauthorized with default headers values
func NewS3CredentialsListUnauthorized() *S3CredentialsListUnauthorized {
	return &S3CredentialsListUnauthorized{}
}

/*
S3CredentialsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type S3CredentialsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this s3 credentials list unauthorized response has a 2xx status code
func (o *S3CredentialsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials list unauthorized response has a 3xx status code
func (o *S3CredentialsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials list unauthorized response has a 4xx status code
func (o *S3CredentialsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials list unauthorized response has a 5xx status code
func (o *S3CredentialsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials list unauthorized response a status code equal to that given
func (o *S3CredentialsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *S3CredentialsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *S3CredentialsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsListForbidden creates a S3CredentialsListForbidden with default headers values
func NewS3CredentialsListForbidden() *S3CredentialsListForbidden {
	return &S3CredentialsListForbidden{}
}

/*
S3CredentialsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type S3CredentialsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this s3 credentials list forbidden response has a 2xx status code
func (o *S3CredentialsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials list forbidden response has a 3xx status code
func (o *S3CredentialsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials list forbidden response has a 4xx status code
func (o *S3CredentialsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials list forbidden response has a 5xx status code
func (o *S3CredentialsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials list forbidden response a status code equal to that given
func (o *S3CredentialsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *S3CredentialsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *S3CredentialsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsListNotFound creates a S3CredentialsListNotFound with default headers values
func NewS3CredentialsListNotFound() *S3CredentialsListNotFound {
	return &S3CredentialsListNotFound{}
}

/*
S3CredentialsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type S3CredentialsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this s3 credentials list not found response has a 2xx status code
func (o *S3CredentialsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials list not found response has a 3xx status code
func (o *S3CredentialsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials list not found response has a 4xx status code
func (o *S3CredentialsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials list not found response has a 5xx status code
func (o *S3CredentialsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials list not found response a status code equal to that given
func (o *S3CredentialsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *S3CredentialsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *S3CredentialsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsListInternalServerError creates a S3CredentialsListInternalServerError with default headers values
func NewS3CredentialsListInternalServerError() *S3CredentialsListInternalServerError {
	return &S3CredentialsListInternalServerError{}
}

/*
S3CredentialsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type S3CredentialsListInternalServerError struct {
}

// IsSuccess returns true when this s3 credentials list internal server error response has a 2xx status code
func (o *S3CredentialsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials list internal server error response has a 3xx status code
func (o *S3CredentialsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials list internal server error response has a 4xx status code
func (o *S3CredentialsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials list internal server error response has a 5xx status code
func (o *S3CredentialsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this s3 credentials list internal server error response a status code equal to that given
func (o *S3CredentialsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *S3CredentialsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListInternalServerError ", 500)
}

func (o *S3CredentialsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/S3Credentials/list][%d] s3CredentialsListInternalServerError ", 500)
}

func (o *S3CredentialsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
