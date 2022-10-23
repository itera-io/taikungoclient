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

// CheckerOrganizationReader is a Reader for the CheckerOrganization structure.
type CheckerOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerOrganizationOK creates a CheckerOrganizationOK with default headers values
func NewCheckerOrganizationOK() *CheckerOrganizationOK {
	return &CheckerOrganizationOK{}
}

/*
CheckerOrganizationOK describes a response with status code 200, with default header values.

Success
*/
type CheckerOrganizationOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this checker organization o k response has a 2xx status code
func (o *CheckerOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker organization o k response has a 3xx status code
func (o *CheckerOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker organization o k response has a 4xx status code
func (o *CheckerOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker organization o k response has a 5xx status code
func (o *CheckerOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker organization o k response a status code equal to that given
func (o *CheckerOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationOK  %+v", 200, o.Payload)
}

func (o *CheckerOrganizationOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationOK  %+v", 200, o.Payload)
}

func (o *CheckerOrganizationOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationBadRequest creates a CheckerOrganizationBadRequest with default headers values
func NewCheckerOrganizationBadRequest() *CheckerOrganizationBadRequest {
	return &CheckerOrganizationBadRequest{}
}

/*
CheckerOrganizationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerOrganizationBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this checker organization bad request response has a 2xx status code
func (o *CheckerOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker organization bad request response has a 3xx status code
func (o *CheckerOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker organization bad request response has a 4xx status code
func (o *CheckerOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker organization bad request response has a 5xx status code
func (o *CheckerOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker organization bad request response a status code equal to that given
func (o *CheckerOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerOrganizationBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerOrganizationBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CheckerOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationUnauthorized creates a CheckerOrganizationUnauthorized with default headers values
func NewCheckerOrganizationUnauthorized() *CheckerOrganizationUnauthorized {
	return &CheckerOrganizationUnauthorized{}
}

/*
CheckerOrganizationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerOrganizationUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this checker organization unauthorized response has a 2xx status code
func (o *CheckerOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker organization unauthorized response has a 3xx status code
func (o *CheckerOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker organization unauthorized response has a 4xx status code
func (o *CheckerOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker organization unauthorized response has a 5xx status code
func (o *CheckerOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker organization unauthorized response a status code equal to that given
func (o *CheckerOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerOrganizationUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CheckerOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationForbidden creates a CheckerOrganizationForbidden with default headers values
func NewCheckerOrganizationForbidden() *CheckerOrganizationForbidden {
	return &CheckerOrganizationForbidden{}
}

/*
CheckerOrganizationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerOrganizationForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this checker organization forbidden response has a 2xx status code
func (o *CheckerOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker organization forbidden response has a 3xx status code
func (o *CheckerOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker organization forbidden response has a 4xx status code
func (o *CheckerOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker organization forbidden response has a 5xx status code
func (o *CheckerOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker organization forbidden response a status code equal to that given
func (o *CheckerOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *CheckerOrganizationForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *CheckerOrganizationForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CheckerOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationNotFound creates a CheckerOrganizationNotFound with default headers values
func NewCheckerOrganizationNotFound() *CheckerOrganizationNotFound {
	return &CheckerOrganizationNotFound{}
}

/*
CheckerOrganizationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerOrganizationNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this checker organization not found response has a 2xx status code
func (o *CheckerOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker organization not found response has a 3xx status code
func (o *CheckerOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker organization not found response has a 4xx status code
func (o *CheckerOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker organization not found response has a 5xx status code
func (o *CheckerOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker organization not found response a status code equal to that given
func (o *CheckerOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *CheckerOrganizationNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *CheckerOrganizationNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CheckerOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerOrganizationInternalServerError creates a CheckerOrganizationInternalServerError with default headers values
func NewCheckerOrganizationInternalServerError() *CheckerOrganizationInternalServerError {
	return &CheckerOrganizationInternalServerError{}
}

/*
CheckerOrganizationInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerOrganizationInternalServerError struct {
}

// IsSuccess returns true when this checker organization internal server error response has a 2xx status code
func (o *CheckerOrganizationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker organization internal server error response has a 3xx status code
func (o *CheckerOrganizationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker organization internal server error response has a 4xx status code
func (o *CheckerOrganizationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker organization internal server error response has a 5xx status code
func (o *CheckerOrganizationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker organization internal server error response a status code equal to that given
func (o *CheckerOrganizationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationInternalServerError ", 500)
}

func (o *CheckerOrganizationInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/organization][%d] checkerOrganizationInternalServerError ", 500)
}

func (o *CheckerOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
