// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminUpdateUserEmailReader is a Reader for the AdminUpdateUserEmail structure.
type AdminUpdateUserEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUpdateUserEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAdminUpdateUserEmailTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateUserEmailOK creates a AdminUpdateUserEmailOK with default headers values
func NewAdminUpdateUserEmailOK() *AdminUpdateUserEmailOK {
	return &AdminUpdateUserEmailOK{}
}

/* AdminUpdateUserEmailOK describes a response with status code 200, with default header values.

Success
*/
type AdminUpdateUserEmailOK struct {
	Payload models.Unit
}

func (o *AdminUpdateUserEmailOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailOK  %+v", 200, o.Payload)
}
func (o *AdminUpdateUserEmailOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminUpdateUserEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailBadRequest creates a AdminUpdateUserEmailBadRequest with default headers values
func NewAdminUpdateUserEmailBadRequest() *AdminUpdateUserEmailBadRequest {
	return &AdminUpdateUserEmailBadRequest{}
}

/* AdminUpdateUserEmailBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUpdateUserEmailBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AdminUpdateUserEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailBadRequest  %+v", 400, o.Payload)
}
func (o *AdminUpdateUserEmailBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailUnauthorized creates a AdminUpdateUserEmailUnauthorized with default headers values
func NewAdminUpdateUserEmailUnauthorized() *AdminUpdateUserEmailUnauthorized {
	return &AdminUpdateUserEmailUnauthorized{}
}

/* AdminUpdateUserEmailUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUpdateUserEmailUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserEmailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminUpdateUserEmailUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserEmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailForbidden creates a AdminUpdateUserEmailForbidden with default headers values
func NewAdminUpdateUserEmailForbidden() *AdminUpdateUserEmailForbidden {
	return &AdminUpdateUserEmailForbidden{}
}

/* AdminUpdateUserEmailForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUpdateUserEmailForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserEmailForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailForbidden  %+v", 403, o.Payload)
}
func (o *AdminUpdateUserEmailForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailNotFound creates a AdminUpdateUserEmailNotFound with default headers values
func NewAdminUpdateUserEmailNotFound() *AdminUpdateUserEmailNotFound {
	return &AdminUpdateUserEmailNotFound{}
}

/* AdminUpdateUserEmailNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUpdateUserEmailNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailNotFound  %+v", 404, o.Payload)
}
func (o *AdminUpdateUserEmailNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailTooManyRequests creates a AdminUpdateUserEmailTooManyRequests with default headers values
func NewAdminUpdateUserEmailTooManyRequests() *AdminUpdateUserEmailTooManyRequests {
	return &AdminUpdateUserEmailTooManyRequests{}
}

/* AdminUpdateUserEmailTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AdminUpdateUserEmailTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AdminUpdateUserEmailTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailTooManyRequests  %+v", 429, o.Payload)
}
func (o *AdminUpdateUserEmailTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserEmailTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailInternalServerError creates a AdminUpdateUserEmailInternalServerError with default headers values
func NewAdminUpdateUserEmailInternalServerError() *AdminUpdateUserEmailInternalServerError {
	return &AdminUpdateUserEmailInternalServerError{}
}

/* AdminUpdateUserEmailInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUpdateUserEmailInternalServerError struct {
}

func (o *AdminUpdateUserEmailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailInternalServerError ", 500)
}

func (o *AdminUpdateUserEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
