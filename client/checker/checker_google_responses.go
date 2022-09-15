// Code generated by go-swagger; DO NOT EDIT.

package checker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CheckerGoogleReader is a Reader for the CheckerGoogle structure.
type CheckerGoogleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerGoogleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerGoogleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerGoogleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerGoogleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerGoogleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerGoogleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerGoogleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerGoogleOK creates a CheckerGoogleOK with default headers values
func NewCheckerGoogleOK() *CheckerGoogleOK {
	return &CheckerGoogleOK{}
}

/*
CheckerGoogleOK describes a response with status code 200, with default header values.

Success
*/
type CheckerGoogleOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this checker google o k response has a 2xx status code
func (o *CheckerGoogleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker google o k response has a 3xx status code
func (o *CheckerGoogleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker google o k response has a 4xx status code
func (o *CheckerGoogleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker google o k response has a 5xx status code
func (o *CheckerGoogleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker google o k response a status code equal to that given
func (o *CheckerGoogleOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerGoogleOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleOK  %+v", 200, o.Payload)
}

func (o *CheckerGoogleOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleOK  %+v", 200, o.Payload)
}

func (o *CheckerGoogleOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerGoogleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerGoogleBadRequest creates a CheckerGoogleBadRequest with default headers values
func NewCheckerGoogleBadRequest() *CheckerGoogleBadRequest {
	return &CheckerGoogleBadRequest{}
}

/*
CheckerGoogleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerGoogleBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this checker google bad request response has a 2xx status code
func (o *CheckerGoogleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker google bad request response has a 3xx status code
func (o *CheckerGoogleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker google bad request response has a 4xx status code
func (o *CheckerGoogleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker google bad request response has a 5xx status code
func (o *CheckerGoogleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker google bad request response a status code equal to that given
func (o *CheckerGoogleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerGoogleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerGoogleBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerGoogleBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerGoogleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerGoogleUnauthorized creates a CheckerGoogleUnauthorized with default headers values
func NewCheckerGoogleUnauthorized() *CheckerGoogleUnauthorized {
	return &CheckerGoogleUnauthorized{}
}

/*
CheckerGoogleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerGoogleUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this checker google unauthorized response has a 2xx status code
func (o *CheckerGoogleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker google unauthorized response has a 3xx status code
func (o *CheckerGoogleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker google unauthorized response has a 4xx status code
func (o *CheckerGoogleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker google unauthorized response has a 5xx status code
func (o *CheckerGoogleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker google unauthorized response a status code equal to that given
func (o *CheckerGoogleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerGoogleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerGoogleUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerGoogleUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerGoogleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerGoogleForbidden creates a CheckerGoogleForbidden with default headers values
func NewCheckerGoogleForbidden() *CheckerGoogleForbidden {
	return &CheckerGoogleForbidden{}
}

/*
CheckerGoogleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerGoogleForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this checker google forbidden response has a 2xx status code
func (o *CheckerGoogleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker google forbidden response has a 3xx status code
func (o *CheckerGoogleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker google forbidden response has a 4xx status code
func (o *CheckerGoogleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker google forbidden response has a 5xx status code
func (o *CheckerGoogleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker google forbidden response a status code equal to that given
func (o *CheckerGoogleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerGoogleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleForbidden  %+v", 403, o.Payload)
}

func (o *CheckerGoogleForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleForbidden  %+v", 403, o.Payload)
}

func (o *CheckerGoogleForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerGoogleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerGoogleNotFound creates a CheckerGoogleNotFound with default headers values
func NewCheckerGoogleNotFound() *CheckerGoogleNotFound {
	return &CheckerGoogleNotFound{}
}

/*
CheckerGoogleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerGoogleNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this checker google not found response has a 2xx status code
func (o *CheckerGoogleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker google not found response has a 3xx status code
func (o *CheckerGoogleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker google not found response has a 4xx status code
func (o *CheckerGoogleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker google not found response has a 5xx status code
func (o *CheckerGoogleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker google not found response a status code equal to that given
func (o *CheckerGoogleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerGoogleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleNotFound  %+v", 404, o.Payload)
}

func (o *CheckerGoogleNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleNotFound  %+v", 404, o.Payload)
}

func (o *CheckerGoogleNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerGoogleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerGoogleInternalServerError creates a CheckerGoogleInternalServerError with default headers values
func NewCheckerGoogleInternalServerError() *CheckerGoogleInternalServerError {
	return &CheckerGoogleInternalServerError{}
}

/*
CheckerGoogleInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerGoogleInternalServerError struct {
}

// IsSuccess returns true when this checker google internal server error response has a 2xx status code
func (o *CheckerGoogleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker google internal server error response has a 3xx status code
func (o *CheckerGoogleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker google internal server error response has a 4xx status code
func (o *CheckerGoogleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker google internal server error response has a 5xx status code
func (o *CheckerGoogleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker google internal server error response a status code equal to that given
func (o *CheckerGoogleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerGoogleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleInternalServerError ", 500)
}

func (o *CheckerGoogleInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/google][%d] checkerGoogleInternalServerError ", 500)
}

func (o *CheckerGoogleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
