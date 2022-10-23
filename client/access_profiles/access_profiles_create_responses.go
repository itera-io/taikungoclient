// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AccessProfilesCreateReader is a Reader for the AccessProfilesCreate structure.
type AccessProfilesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessProfilesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessProfilesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessProfilesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccessProfilesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccessProfilesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessProfilesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessProfilesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccessProfilesCreateOK creates a AccessProfilesCreateOK with default headers values
func NewAccessProfilesCreateOK() *AccessProfilesCreateOK {
	return &AccessProfilesCreateOK{}
}

/*
AccessProfilesCreateOK describes a response with status code 200, with default header values.

Success
*/
type AccessProfilesCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this access profiles create o k response has a 2xx status code
func (o *AccessProfilesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access profiles create o k response has a 3xx status code
func (o *AccessProfilesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles create o k response has a 4xx status code
func (o *AccessProfilesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles create o k response has a 5xx status code
func (o *AccessProfilesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles create o k response a status code equal to that given
func (o *AccessProfilesCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AccessProfilesCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *AccessProfilesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesCreateBadRequest creates a AccessProfilesCreateBadRequest with default headers values
func NewAccessProfilesCreateBadRequest() *AccessProfilesCreateBadRequest {
	return &AccessProfilesCreateBadRequest{}
}

/*
AccessProfilesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AccessProfilesCreateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this access profiles create bad request response has a 2xx status code
func (o *AccessProfilesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles create bad request response has a 3xx status code
func (o *AccessProfilesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles create bad request response has a 4xx status code
func (o *AccessProfilesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles create bad request response has a 5xx status code
func (o *AccessProfilesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles create bad request response a status code equal to that given
func (o *AccessProfilesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AccessProfilesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesCreateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AccessProfilesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesCreateUnauthorized creates a AccessProfilesCreateUnauthorized with default headers values
func NewAccessProfilesCreateUnauthorized() *AccessProfilesCreateUnauthorized {
	return &AccessProfilesCreateUnauthorized{}
}

/*
AccessProfilesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AccessProfilesCreateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this access profiles create unauthorized response has a 2xx status code
func (o *AccessProfilesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles create unauthorized response has a 3xx status code
func (o *AccessProfilesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles create unauthorized response has a 4xx status code
func (o *AccessProfilesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles create unauthorized response has a 5xx status code
func (o *AccessProfilesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles create unauthorized response a status code equal to that given
func (o *AccessProfilesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AccessProfilesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesCreateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AccessProfilesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesCreateForbidden creates a AccessProfilesCreateForbidden with default headers values
func NewAccessProfilesCreateForbidden() *AccessProfilesCreateForbidden {
	return &AccessProfilesCreateForbidden{}
}

/*
AccessProfilesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccessProfilesCreateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this access profiles create forbidden response has a 2xx status code
func (o *AccessProfilesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles create forbidden response has a 3xx status code
func (o *AccessProfilesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles create forbidden response has a 4xx status code
func (o *AccessProfilesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles create forbidden response has a 5xx status code
func (o *AccessProfilesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles create forbidden response a status code equal to that given
func (o *AccessProfilesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AccessProfilesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesCreateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AccessProfilesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesCreateNotFound creates a AccessProfilesCreateNotFound with default headers values
func NewAccessProfilesCreateNotFound() *AccessProfilesCreateNotFound {
	return &AccessProfilesCreateNotFound{}
}

/*
AccessProfilesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AccessProfilesCreateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this access profiles create not found response has a 2xx status code
func (o *AccessProfilesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles create not found response has a 3xx status code
func (o *AccessProfilesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles create not found response has a 4xx status code
func (o *AccessProfilesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles create not found response has a 5xx status code
func (o *AccessProfilesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles create not found response a status code equal to that given
func (o *AccessProfilesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AccessProfilesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesCreateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AccessProfilesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesCreateInternalServerError creates a AccessProfilesCreateInternalServerError with default headers values
func NewAccessProfilesCreateInternalServerError() *AccessProfilesCreateInternalServerError {
	return &AccessProfilesCreateInternalServerError{}
}

/*
AccessProfilesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AccessProfilesCreateInternalServerError struct {
}

// IsSuccess returns true when this access profiles create internal server error response has a 2xx status code
func (o *AccessProfilesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles create internal server error response has a 3xx status code
func (o *AccessProfilesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles create internal server error response has a 4xx status code
func (o *AccessProfilesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles create internal server error response has a 5xx status code
func (o *AccessProfilesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access profiles create internal server error response a status code equal to that given
func (o *AccessProfilesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AccessProfilesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateInternalServerError ", 500)
}

func (o *AccessProfilesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AccessProfiles/create][%d] accessProfilesCreateInternalServerError ", 500)
}

func (o *AccessProfilesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
