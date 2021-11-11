// Code generated by go-swagger; DO NOT EDIT.

package opa_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpaProfilesDeleteReader is a Reader for the OpaProfilesDelete structure.
type OpaProfilesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesDeleteOK creates a OpaProfilesDeleteOK with default headers values
func NewOpaProfilesDeleteOK() *OpaProfilesDeleteOK {
	return &OpaProfilesDeleteOK{}
}

/* OpaProfilesDeleteOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesDeleteOK struct {
	Payload models.Unit
}

func (o *OpaProfilesDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteOK  %+v", 200, o.Payload)
}
func (o *OpaProfilesDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpaProfilesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteBadRequest creates a OpaProfilesDeleteBadRequest with default headers values
func NewOpaProfilesDeleteBadRequest() *OpaProfilesDeleteBadRequest {
	return &OpaProfilesDeleteBadRequest{}
}

/* OpaProfilesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OpaProfilesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *OpaProfilesDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteUnauthorized creates a OpaProfilesDeleteUnauthorized with default headers values
func NewOpaProfilesDeleteUnauthorized() *OpaProfilesDeleteUnauthorized {
	return &OpaProfilesDeleteUnauthorized{}
}

/* OpaProfilesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *OpaProfilesDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteForbidden creates a OpaProfilesDeleteForbidden with default headers values
func NewOpaProfilesDeleteForbidden() *OpaProfilesDeleteForbidden {
	return &OpaProfilesDeleteForbidden{}
}

/* OpaProfilesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesDeleteForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteForbidden  %+v", 403, o.Payload)
}
func (o *OpaProfilesDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteNotFound creates a OpaProfilesDeleteNotFound with default headers values
func NewOpaProfilesDeleteNotFound() *OpaProfilesDeleteNotFound {
	return &OpaProfilesDeleteNotFound{}
}

/* OpaProfilesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesDeleteNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteNotFound  %+v", 404, o.Payload)
}
func (o *OpaProfilesDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDeleteInternalServerError creates a OpaProfilesDeleteInternalServerError with default headers values
func NewOpaProfilesDeleteInternalServerError() *OpaProfilesDeleteInternalServerError {
	return &OpaProfilesDeleteInternalServerError{}
}

/* OpaProfilesDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesDeleteInternalServerError struct {
}

func (o *OpaProfilesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/delete][%d] opaProfilesDeleteInternalServerError ", 500)
}

func (o *OpaProfilesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}