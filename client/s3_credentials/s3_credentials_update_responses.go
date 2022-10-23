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

// S3CredentialsUpdateReader is a Reader for the S3CredentialsUpdate structure.
type S3CredentialsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3CredentialsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3CredentialsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3CredentialsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS3CredentialsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS3CredentialsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS3CredentialsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewS3CredentialsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewS3CredentialsUpdateOK creates a S3CredentialsUpdateOK with default headers values
func NewS3CredentialsUpdateOK() *S3CredentialsUpdateOK {
	return &S3CredentialsUpdateOK{}
}

/*
S3CredentialsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type S3CredentialsUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this s3 credentials update o k response has a 2xx status code
func (o *S3CredentialsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 credentials update o k response has a 3xx status code
func (o *S3CredentialsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials update o k response has a 4xx status code
func (o *S3CredentialsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials update o k response has a 5xx status code
func (o *S3CredentialsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials update o k response a status code equal to that given
func (o *S3CredentialsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3CredentialsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *S3CredentialsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsUpdateBadRequest creates a S3CredentialsUpdateBadRequest with default headers values
func NewS3CredentialsUpdateBadRequest() *S3CredentialsUpdateBadRequest {
	return &S3CredentialsUpdateBadRequest{}
}

/*
S3CredentialsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type S3CredentialsUpdateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this s3 credentials update bad request response has a 2xx status code
func (o *S3CredentialsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials update bad request response has a 3xx status code
func (o *S3CredentialsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials update bad request response has a 4xx status code
func (o *S3CredentialsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials update bad request response has a 5xx status code
func (o *S3CredentialsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials update bad request response a status code equal to that given
func (o *S3CredentialsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *S3CredentialsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsUpdateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *S3CredentialsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsUpdateUnauthorized creates a S3CredentialsUpdateUnauthorized with default headers values
func NewS3CredentialsUpdateUnauthorized() *S3CredentialsUpdateUnauthorized {
	return &S3CredentialsUpdateUnauthorized{}
}

/*
S3CredentialsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type S3CredentialsUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this s3 credentials update unauthorized response has a 2xx status code
func (o *S3CredentialsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials update unauthorized response has a 3xx status code
func (o *S3CredentialsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials update unauthorized response has a 4xx status code
func (o *S3CredentialsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials update unauthorized response has a 5xx status code
func (o *S3CredentialsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials update unauthorized response a status code equal to that given
func (o *S3CredentialsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *S3CredentialsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *S3CredentialsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsUpdateForbidden creates a S3CredentialsUpdateForbidden with default headers values
func NewS3CredentialsUpdateForbidden() *S3CredentialsUpdateForbidden {
	return &S3CredentialsUpdateForbidden{}
}

/*
S3CredentialsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type S3CredentialsUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this s3 credentials update forbidden response has a 2xx status code
func (o *S3CredentialsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials update forbidden response has a 3xx status code
func (o *S3CredentialsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials update forbidden response has a 4xx status code
func (o *S3CredentialsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials update forbidden response has a 5xx status code
func (o *S3CredentialsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials update forbidden response a status code equal to that given
func (o *S3CredentialsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *S3CredentialsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *S3CredentialsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsUpdateNotFound creates a S3CredentialsUpdateNotFound with default headers values
func NewS3CredentialsUpdateNotFound() *S3CredentialsUpdateNotFound {
	return &S3CredentialsUpdateNotFound{}
}

/*
S3CredentialsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type S3CredentialsUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this s3 credentials update not found response has a 2xx status code
func (o *S3CredentialsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials update not found response has a 3xx status code
func (o *S3CredentialsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials update not found response has a 4xx status code
func (o *S3CredentialsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials update not found response has a 5xx status code
func (o *S3CredentialsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials update not found response a status code equal to that given
func (o *S3CredentialsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *S3CredentialsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *S3CredentialsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsUpdateInternalServerError creates a S3CredentialsUpdateInternalServerError with default headers values
func NewS3CredentialsUpdateInternalServerError() *S3CredentialsUpdateInternalServerError {
	return &S3CredentialsUpdateInternalServerError{}
}

/*
S3CredentialsUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type S3CredentialsUpdateInternalServerError struct {
}

// IsSuccess returns true when this s3 credentials update internal server error response has a 2xx status code
func (o *S3CredentialsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials update internal server error response has a 3xx status code
func (o *S3CredentialsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials update internal server error response has a 4xx status code
func (o *S3CredentialsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials update internal server error response has a 5xx status code
func (o *S3CredentialsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this s3 credentials update internal server error response a status code equal to that given
func (o *S3CredentialsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *S3CredentialsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateInternalServerError ", 500)
}

func (o *S3CredentialsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/S3Credentials][%d] s3CredentialsUpdateInternalServerError ", 500)
}

func (o *S3CredentialsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
