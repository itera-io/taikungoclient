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

// AccessProfilesUpdateReader is a Reader for the AccessProfilesUpdate structure.
type AccessProfilesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessProfilesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessProfilesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessProfilesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccessProfilesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccessProfilesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessProfilesUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessProfilesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccessProfilesUpdateOK creates a AccessProfilesUpdateOK with default headers values
func NewAccessProfilesUpdateOK() *AccessProfilesUpdateOK {
	return &AccessProfilesUpdateOK{}
}

/*
AccessProfilesUpdateOK describes a response with status code 200, with default header values.

Success
*/
type AccessProfilesUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this access profiles update o k response has a 2xx status code
func (o *AccessProfilesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access profiles update o k response has a 3xx status code
func (o *AccessProfilesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update o k response has a 4xx status code
func (o *AccessProfilesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles update o k response has a 5xx status code
func (o *AccessProfilesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update o k response a status code equal to that given
func (o *AccessProfilesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AccessProfilesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AccessProfilesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateBadRequest creates a AccessProfilesUpdateBadRequest with default headers values
func NewAccessProfilesUpdateBadRequest() *AccessProfilesUpdateBadRequest {
	return &AccessProfilesUpdateBadRequest{}
}

/*
AccessProfilesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AccessProfilesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this access profiles update bad request response has a 2xx status code
func (o *AccessProfilesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update bad request response has a 3xx status code
func (o *AccessProfilesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update bad request response has a 4xx status code
func (o *AccessProfilesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update bad request response has a 5xx status code
func (o *AccessProfilesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update bad request response a status code equal to that given
func (o *AccessProfilesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AccessProfilesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AccessProfilesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateUnauthorized creates a AccessProfilesUpdateUnauthorized with default headers values
func NewAccessProfilesUpdateUnauthorized() *AccessProfilesUpdateUnauthorized {
	return &AccessProfilesUpdateUnauthorized{}
}

/*
AccessProfilesUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AccessProfilesUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles update unauthorized response has a 2xx status code
func (o *AccessProfilesUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update unauthorized response has a 3xx status code
func (o *AccessProfilesUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update unauthorized response has a 4xx status code
func (o *AccessProfilesUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update unauthorized response has a 5xx status code
func (o *AccessProfilesUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update unauthorized response a status code equal to that given
func (o *AccessProfilesUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AccessProfilesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateForbidden creates a AccessProfilesUpdateForbidden with default headers values
func NewAccessProfilesUpdateForbidden() *AccessProfilesUpdateForbidden {
	return &AccessProfilesUpdateForbidden{}
}

/*
AccessProfilesUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccessProfilesUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles update forbidden response has a 2xx status code
func (o *AccessProfilesUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update forbidden response has a 3xx status code
func (o *AccessProfilesUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update forbidden response has a 4xx status code
func (o *AccessProfilesUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update forbidden response has a 5xx status code
func (o *AccessProfilesUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update forbidden response a status code equal to that given
func (o *AccessProfilesUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AccessProfilesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateNotFound creates a AccessProfilesUpdateNotFound with default headers values
func NewAccessProfilesUpdateNotFound() *AccessProfilesUpdateNotFound {
	return &AccessProfilesUpdateNotFound{}
}

/*
AccessProfilesUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AccessProfilesUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles update not found response has a 2xx status code
func (o *AccessProfilesUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update not found response has a 3xx status code
func (o *AccessProfilesUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update not found response has a 4xx status code
func (o *AccessProfilesUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles update not found response has a 5xx status code
func (o *AccessProfilesUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles update not found response a status code equal to that given
func (o *AccessProfilesUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AccessProfilesUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesUpdateInternalServerError creates a AccessProfilesUpdateInternalServerError with default headers values
func NewAccessProfilesUpdateInternalServerError() *AccessProfilesUpdateInternalServerError {
	return &AccessProfilesUpdateInternalServerError{}
}

/*
AccessProfilesUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AccessProfilesUpdateInternalServerError struct {
}

// IsSuccess returns true when this access profiles update internal server error response has a 2xx status code
func (o *AccessProfilesUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles update internal server error response has a 3xx status code
func (o *AccessProfilesUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles update internal server error response has a 4xx status code
func (o *AccessProfilesUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles update internal server error response has a 5xx status code
func (o *AccessProfilesUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access profiles update internal server error response a status code equal to that given
func (o *AccessProfilesUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AccessProfilesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateInternalServerError ", 500)
}

func (o *AccessProfilesUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AccessProfiles/update/{id}][%d] accessProfilesUpdateInternalServerError ", 500)
}

func (o *AccessProfilesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
