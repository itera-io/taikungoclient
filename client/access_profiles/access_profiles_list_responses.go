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

// AccessProfilesListReader is a Reader for the AccessProfilesList structure.
type AccessProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccessProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccessProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAccessProfilesListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccessProfilesListOK creates a AccessProfilesListOK with default headers values
func NewAccessProfilesListOK() *AccessProfilesListOK {
	return &AccessProfilesListOK{}
}

/* AccessProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type AccessProfilesListOK struct {
	Payload *models.AccessProfilesList
}

func (o *AccessProfilesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListOK  %+v", 200, o.Payload)
}
func (o *AccessProfilesListOK) GetPayload() *models.AccessProfilesList {
	return o.Payload
}

func (o *AccessProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessProfilesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesListBadRequest creates a AccessProfilesListBadRequest with default headers values
func NewAccessProfilesListBadRequest() *AccessProfilesListBadRequest {
	return &AccessProfilesListBadRequest{}
}

/* AccessProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AccessProfilesListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AccessProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListBadRequest  %+v", 400, o.Payload)
}
func (o *AccessProfilesListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AccessProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesListUnauthorized creates a AccessProfilesListUnauthorized with default headers values
func NewAccessProfilesListUnauthorized() *AccessProfilesListUnauthorized {
	return &AccessProfilesListUnauthorized{}
}

/* AccessProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AccessProfilesListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AccessProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListUnauthorized  %+v", 401, o.Payload)
}
func (o *AccessProfilesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesListForbidden creates a AccessProfilesListForbidden with default headers values
func NewAccessProfilesListForbidden() *AccessProfilesListForbidden {
	return &AccessProfilesListForbidden{}
}

/* AccessProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccessProfilesListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AccessProfilesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListForbidden  %+v", 403, o.Payload)
}
func (o *AccessProfilesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesListNotFound creates a AccessProfilesListNotFound with default headers values
func NewAccessProfilesListNotFound() *AccessProfilesListNotFound {
	return &AccessProfilesListNotFound{}
}

/* AccessProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AccessProfilesListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AccessProfilesListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListNotFound  %+v", 404, o.Payload)
}
func (o *AccessProfilesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesListTooManyRequests creates a AccessProfilesListTooManyRequests with default headers values
func NewAccessProfilesListTooManyRequests() *AccessProfilesListTooManyRequests {
	return &AccessProfilesListTooManyRequests{}
}

/* AccessProfilesListTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AccessProfilesListTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AccessProfilesListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListTooManyRequests  %+v", 429, o.Payload)
}
func (o *AccessProfilesListTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesListInternalServerError creates a AccessProfilesListInternalServerError with default headers values
func NewAccessProfilesListInternalServerError() *AccessProfilesListInternalServerError {
	return &AccessProfilesListInternalServerError{}
}

/* AccessProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AccessProfilesListInternalServerError struct {
}

func (o *AccessProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles][%d] accessProfilesListInternalServerError ", 500)
}

func (o *AccessProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
