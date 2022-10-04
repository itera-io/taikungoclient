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

// S3CredentialsDeleteReader is a Reader for the S3CredentialsDelete structure.
type S3CredentialsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3CredentialsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3CredentialsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewS3CredentialsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3CredentialsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS3CredentialsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS3CredentialsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS3CredentialsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewS3CredentialsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewS3CredentialsDeleteOK creates a S3CredentialsDeleteOK with default headers values
func NewS3CredentialsDeleteOK() *S3CredentialsDeleteOK {
	return &S3CredentialsDeleteOK{}
}

/*
S3CredentialsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type S3CredentialsDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this s3 credentials delete o k response has a 2xx status code
func (o *S3CredentialsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 credentials delete o k response has a 3xx status code
func (o *S3CredentialsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete o k response has a 4xx status code
func (o *S3CredentialsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials delete o k response has a 5xx status code
func (o *S3CredentialsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials delete o k response a status code equal to that given
func (o *S3CredentialsDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3CredentialsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *S3CredentialsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsDeleteNoContent creates a S3CredentialsDeleteNoContent with default headers values
func NewS3CredentialsDeleteNoContent() *S3CredentialsDeleteNoContent {
	return &S3CredentialsDeleteNoContent{}
}

/*
S3CredentialsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type S3CredentialsDeleteNoContent struct {
}

// IsSuccess returns true when this s3 credentials delete no content response has a 2xx status code
func (o *S3CredentialsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 credentials delete no content response has a 3xx status code
func (o *S3CredentialsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete no content response has a 4xx status code
func (o *S3CredentialsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials delete no content response has a 5xx status code
func (o *S3CredentialsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials delete no content response a status code equal to that given
func (o *S3CredentialsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *S3CredentialsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteNoContent ", 204)
}

func (o *S3CredentialsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteNoContent ", 204)
}

func (o *S3CredentialsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewS3CredentialsDeleteBadRequest creates a S3CredentialsDeleteBadRequest with default headers values
func NewS3CredentialsDeleteBadRequest() *S3CredentialsDeleteBadRequest {
	return &S3CredentialsDeleteBadRequest{}
}

/*
S3CredentialsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type S3CredentialsDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this s3 credentials delete bad request response has a 2xx status code
func (o *S3CredentialsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials delete bad request response has a 3xx status code
func (o *S3CredentialsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete bad request response has a 4xx status code
func (o *S3CredentialsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials delete bad request response has a 5xx status code
func (o *S3CredentialsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials delete bad request response a status code equal to that given
func (o *S3CredentialsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *S3CredentialsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *S3CredentialsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsDeleteUnauthorized creates a S3CredentialsDeleteUnauthorized with default headers values
func NewS3CredentialsDeleteUnauthorized() *S3CredentialsDeleteUnauthorized {
	return &S3CredentialsDeleteUnauthorized{}
}

/*
S3CredentialsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type S3CredentialsDeleteUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this s3 credentials delete unauthorized response has a 2xx status code
func (o *S3CredentialsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials delete unauthorized response has a 3xx status code
func (o *S3CredentialsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete unauthorized response has a 4xx status code
func (o *S3CredentialsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials delete unauthorized response has a 5xx status code
func (o *S3CredentialsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials delete unauthorized response a status code equal to that given
func (o *S3CredentialsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *S3CredentialsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsDeleteUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *S3CredentialsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsDeleteForbidden creates a S3CredentialsDeleteForbidden with default headers values
func NewS3CredentialsDeleteForbidden() *S3CredentialsDeleteForbidden {
	return &S3CredentialsDeleteForbidden{}
}

/*
S3CredentialsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type S3CredentialsDeleteForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this s3 credentials delete forbidden response has a 2xx status code
func (o *S3CredentialsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials delete forbidden response has a 3xx status code
func (o *S3CredentialsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete forbidden response has a 4xx status code
func (o *S3CredentialsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials delete forbidden response has a 5xx status code
func (o *S3CredentialsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials delete forbidden response a status code equal to that given
func (o *S3CredentialsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *S3CredentialsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsDeleteForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *S3CredentialsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsDeleteNotFound creates a S3CredentialsDeleteNotFound with default headers values
func NewS3CredentialsDeleteNotFound() *S3CredentialsDeleteNotFound {
	return &S3CredentialsDeleteNotFound{}
}

/*
S3CredentialsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type S3CredentialsDeleteNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this s3 credentials delete not found response has a 2xx status code
func (o *S3CredentialsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials delete not found response has a 3xx status code
func (o *S3CredentialsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete not found response has a 4xx status code
func (o *S3CredentialsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials delete not found response has a 5xx status code
func (o *S3CredentialsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials delete not found response a status code equal to that given
func (o *S3CredentialsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *S3CredentialsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsDeleteNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *S3CredentialsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsDeleteInternalServerError creates a S3CredentialsDeleteInternalServerError with default headers values
func NewS3CredentialsDeleteInternalServerError() *S3CredentialsDeleteInternalServerError {
	return &S3CredentialsDeleteInternalServerError{}
}

/*
S3CredentialsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type S3CredentialsDeleteInternalServerError struct {
}

// IsSuccess returns true when this s3 credentials delete internal server error response has a 2xx status code
func (o *S3CredentialsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials delete internal server error response has a 3xx status code
func (o *S3CredentialsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials delete internal server error response has a 4xx status code
func (o *S3CredentialsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials delete internal server error response has a 5xx status code
func (o *S3CredentialsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this s3 credentials delete internal server error response a status code equal to that given
func (o *S3CredentialsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *S3CredentialsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteInternalServerError ", 500)
}

func (o *S3CredentialsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/S3Credentials/{id}][%d] s3CredentialsDeleteInternalServerError ", 500)
}

func (o *S3CredentialsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
