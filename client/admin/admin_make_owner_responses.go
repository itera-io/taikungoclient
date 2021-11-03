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

// AdminMakeOwnerReader is a Reader for the AdminMakeOwner structure.
type AdminMakeOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminMakeOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminMakeOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminMakeOwnerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminMakeOwnerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminMakeOwnerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminMakeOwnerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAdminMakeOwnerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminMakeOwnerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminMakeOwnerOK creates a AdminMakeOwnerOK with default headers values
func NewAdminMakeOwnerOK() *AdminMakeOwnerOK {
	return &AdminMakeOwnerOK{}
}

/* AdminMakeOwnerOK describes a response with status code 200, with default header values.

Success
*/
type AdminMakeOwnerOK struct {
	Payload models.Unit
}

func (o *AdminMakeOwnerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerOK  %+v", 200, o.Payload)
}
func (o *AdminMakeOwnerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminMakeOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeOwnerBadRequest creates a AdminMakeOwnerBadRequest with default headers values
func NewAdminMakeOwnerBadRequest() *AdminMakeOwnerBadRequest {
	return &AdminMakeOwnerBadRequest{}
}

/* AdminMakeOwnerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminMakeOwnerBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AdminMakeOwnerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerBadRequest  %+v", 400, o.Payload)
}
func (o *AdminMakeOwnerBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminMakeOwnerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeOwnerUnauthorized creates a AdminMakeOwnerUnauthorized with default headers values
func NewAdminMakeOwnerUnauthorized() *AdminMakeOwnerUnauthorized {
	return &AdminMakeOwnerUnauthorized{}
}

/* AdminMakeOwnerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminMakeOwnerUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AdminMakeOwnerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminMakeOwnerUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminMakeOwnerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeOwnerForbidden creates a AdminMakeOwnerForbidden with default headers values
func NewAdminMakeOwnerForbidden() *AdminMakeOwnerForbidden {
	return &AdminMakeOwnerForbidden{}
}

/* AdminMakeOwnerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminMakeOwnerForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AdminMakeOwnerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerForbidden  %+v", 403, o.Payload)
}
func (o *AdminMakeOwnerForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminMakeOwnerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeOwnerNotFound creates a AdminMakeOwnerNotFound with default headers values
func NewAdminMakeOwnerNotFound() *AdminMakeOwnerNotFound {
	return &AdminMakeOwnerNotFound{}
}

/* AdminMakeOwnerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminMakeOwnerNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AdminMakeOwnerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerNotFound  %+v", 404, o.Payload)
}
func (o *AdminMakeOwnerNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminMakeOwnerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeOwnerTooManyRequests creates a AdminMakeOwnerTooManyRequests with default headers values
func NewAdminMakeOwnerTooManyRequests() *AdminMakeOwnerTooManyRequests {
	return &AdminMakeOwnerTooManyRequests{}
}

/* AdminMakeOwnerTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AdminMakeOwnerTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AdminMakeOwnerTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerTooManyRequests  %+v", 429, o.Payload)
}
func (o *AdminMakeOwnerTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminMakeOwnerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeOwnerInternalServerError creates a AdminMakeOwnerInternalServerError with default headers values
func NewAdminMakeOwnerInternalServerError() *AdminMakeOwnerInternalServerError {
	return &AdminMakeOwnerInternalServerError{}
}

/* AdminMakeOwnerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminMakeOwnerInternalServerError struct {
}

func (o *AdminMakeOwnerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/owner][%d] adminMakeOwnerInternalServerError ", 500)
}

func (o *AdminMakeOwnerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
