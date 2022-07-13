// Code generated by go-swagger; DO NOT EDIT.

package allowed_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AllowedHostDeleteReader is a Reader for the AllowedHostDelete structure.
type AllowedHostDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllowedHostDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllowedHostDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAllowedHostDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllowedHostDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAllowedHostDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAllowedHostDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllowedHostDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAllowedHostDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllowedHostDeleteOK creates a AllowedHostDeleteOK with default headers values
func NewAllowedHostDeleteOK() *AllowedHostDeleteOK {
	return &AllowedHostDeleteOK{}
}

/* AllowedHostDeleteOK describes a response with status code 200, with default header values.

Success
*/
type AllowedHostDeleteOK struct {
	Payload models.Unit
}

func (o *AllowedHostDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteOK  %+v", 200, o.Payload)
}
func (o *AllowedHostDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AllowedHostDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostDeleteNoContent creates a AllowedHostDeleteNoContent with default headers values
func NewAllowedHostDeleteNoContent() *AllowedHostDeleteNoContent {
	return &AllowedHostDeleteNoContent{}
}

/* AllowedHostDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type AllowedHostDeleteNoContent struct {
}

func (o *AllowedHostDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteNoContent ", 204)
}

func (o *AllowedHostDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAllowedHostDeleteBadRequest creates a AllowedHostDeleteBadRequest with default headers values
func NewAllowedHostDeleteBadRequest() *AllowedHostDeleteBadRequest {
	return &AllowedHostDeleteBadRequest{}
}

/* AllowedHostDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllowedHostDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AllowedHostDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *AllowedHostDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AllowedHostDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostDeleteUnauthorized creates a AllowedHostDeleteUnauthorized with default headers values
func NewAllowedHostDeleteUnauthorized() *AllowedHostDeleteUnauthorized {
	return &AllowedHostDeleteUnauthorized{}
}

/* AllowedHostDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AllowedHostDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AllowedHostDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *AllowedHostDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AllowedHostDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostDeleteForbidden creates a AllowedHostDeleteForbidden with default headers values
func NewAllowedHostDeleteForbidden() *AllowedHostDeleteForbidden {
	return &AllowedHostDeleteForbidden{}
}

/* AllowedHostDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AllowedHostDeleteForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AllowedHostDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteForbidden  %+v", 403, o.Payload)
}
func (o *AllowedHostDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AllowedHostDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostDeleteNotFound creates a AllowedHostDeleteNotFound with default headers values
func NewAllowedHostDeleteNotFound() *AllowedHostDeleteNotFound {
	return &AllowedHostDeleteNotFound{}
}

/* AllowedHostDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllowedHostDeleteNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AllowedHostDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteNotFound  %+v", 404, o.Payload)
}
func (o *AllowedHostDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AllowedHostDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostDeleteInternalServerError creates a AllowedHostDeleteInternalServerError with default headers values
func NewAllowedHostDeleteInternalServerError() *AllowedHostDeleteInternalServerError {
	return &AllowedHostDeleteInternalServerError{}
}

/* AllowedHostDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AllowedHostDeleteInternalServerError struct {
}

func (o *AllowedHostDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/AllowedHost/{id}][%d] allowedHostDeleteInternalServerError ", 500)
}

func (o *AllowedHostDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
