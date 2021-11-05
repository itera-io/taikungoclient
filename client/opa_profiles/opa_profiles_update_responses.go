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

// OpaProfilesUpdateReader is a Reader for the OpaProfilesUpdate structure.
type OpaProfilesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesUpdateOK creates a OpaProfilesUpdateOK with default headers values
func NewOpaProfilesUpdateOK() *OpaProfilesUpdateOK {
	return &OpaProfilesUpdateOK{}
}

/* OpaProfilesUpdateOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesUpdateOK struct {
	Payload models.Unit
}

func (o *OpaProfilesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/OpaProfiles][%d] opaProfilesUpdateOK  %+v", 200, o.Payload)
}
func (o *OpaProfilesUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpaProfilesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesUpdateBadRequest creates a OpaProfilesUpdateBadRequest with default headers values
func NewOpaProfilesUpdateBadRequest() *OpaProfilesUpdateBadRequest {
	return &OpaProfilesUpdateBadRequest{}
}

/* OpaProfilesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OpaProfilesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/OpaProfiles][%d] opaProfilesUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *OpaProfilesUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpaProfilesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesUpdateUnauthorized creates a OpaProfilesUpdateUnauthorized with default headers values
func NewOpaProfilesUpdateUnauthorized() *OpaProfilesUpdateUnauthorized {
	return &OpaProfilesUpdateUnauthorized{}
}

/* OpaProfilesUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/OpaProfiles][%d] opaProfilesUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *OpaProfilesUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesUpdateForbidden creates a OpaProfilesUpdateForbidden with default headers values
func NewOpaProfilesUpdateForbidden() *OpaProfilesUpdateForbidden {
	return &OpaProfilesUpdateForbidden{}
}

/* OpaProfilesUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesUpdateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/OpaProfiles][%d] opaProfilesUpdateForbidden  %+v", 403, o.Payload)
}
func (o *OpaProfilesUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesUpdateNotFound creates a OpaProfilesUpdateNotFound with default headers values
func NewOpaProfilesUpdateNotFound() *OpaProfilesUpdateNotFound {
	return &OpaProfilesUpdateNotFound{}
}

/* OpaProfilesUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesUpdateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/OpaProfiles][%d] opaProfilesUpdateNotFound  %+v", 404, o.Payload)
}
func (o *OpaProfilesUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesUpdateInternalServerError creates a OpaProfilesUpdateInternalServerError with default headers values
func NewOpaProfilesUpdateInternalServerError() *OpaProfilesUpdateInternalServerError {
	return &OpaProfilesUpdateInternalServerError{}
}

/* OpaProfilesUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesUpdateInternalServerError struct {
}

func (o *OpaProfilesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/OpaProfiles][%d] opaProfilesUpdateInternalServerError ", 500)
}

func (o *OpaProfilesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
