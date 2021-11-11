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

// OpaProfilesDisableGatekeeperReader is a Reader for the OpaProfilesDisableGatekeeper structure.
type OpaProfilesDisableGatekeeperReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesDisableGatekeeperReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesDisableGatekeeperOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesDisableGatekeeperBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesDisableGatekeeperUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesDisableGatekeeperForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesDisableGatekeeperNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesDisableGatekeeperInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesDisableGatekeeperOK creates a OpaProfilesDisableGatekeeperOK with default headers values
func NewOpaProfilesDisableGatekeeperOK() *OpaProfilesDisableGatekeeperOK {
	return &OpaProfilesDisableGatekeeperOK{}
}

/* OpaProfilesDisableGatekeeperOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesDisableGatekeeperOK struct {
	Payload models.Unit
}

func (o *OpaProfilesDisableGatekeeperOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/disablegatekeeper][%d] opaProfilesDisableGatekeeperOK  %+v", 200, o.Payload)
}
func (o *OpaProfilesDisableGatekeeperOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OpaProfilesDisableGatekeeperOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDisableGatekeeperBadRequest creates a OpaProfilesDisableGatekeeperBadRequest with default headers values
func NewOpaProfilesDisableGatekeeperBadRequest() *OpaProfilesDisableGatekeeperBadRequest {
	return &OpaProfilesDisableGatekeeperBadRequest{}
}

/* OpaProfilesDisableGatekeeperBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesDisableGatekeeperBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OpaProfilesDisableGatekeeperBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/disablegatekeeper][%d] opaProfilesDisableGatekeeperBadRequest  %+v", 400, o.Payload)
}
func (o *OpaProfilesDisableGatekeeperBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDisableGatekeeperBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDisableGatekeeperUnauthorized creates a OpaProfilesDisableGatekeeperUnauthorized with default headers values
func NewOpaProfilesDisableGatekeeperUnauthorized() *OpaProfilesDisableGatekeeperUnauthorized {
	return &OpaProfilesDisableGatekeeperUnauthorized{}
}

/* OpaProfilesDisableGatekeeperUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesDisableGatekeeperUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesDisableGatekeeperUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/disablegatekeeper][%d] opaProfilesDisableGatekeeperUnauthorized  %+v", 401, o.Payload)
}
func (o *OpaProfilesDisableGatekeeperUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDisableGatekeeperUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDisableGatekeeperForbidden creates a OpaProfilesDisableGatekeeperForbidden with default headers values
func NewOpaProfilesDisableGatekeeperForbidden() *OpaProfilesDisableGatekeeperForbidden {
	return &OpaProfilesDisableGatekeeperForbidden{}
}

/* OpaProfilesDisableGatekeeperForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesDisableGatekeeperForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesDisableGatekeeperForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/disablegatekeeper][%d] opaProfilesDisableGatekeeperForbidden  %+v", 403, o.Payload)
}
func (o *OpaProfilesDisableGatekeeperForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDisableGatekeeperForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDisableGatekeeperNotFound creates a OpaProfilesDisableGatekeeperNotFound with default headers values
func NewOpaProfilesDisableGatekeeperNotFound() *OpaProfilesDisableGatekeeperNotFound {
	return &OpaProfilesDisableGatekeeperNotFound{}
}

/* OpaProfilesDisableGatekeeperNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesDisableGatekeeperNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OpaProfilesDisableGatekeeperNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/disablegatekeeper][%d] opaProfilesDisableGatekeeperNotFound  %+v", 404, o.Payload)
}
func (o *OpaProfilesDisableGatekeeperNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpaProfilesDisableGatekeeperNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesDisableGatekeeperInternalServerError creates a OpaProfilesDisableGatekeeperInternalServerError with default headers values
func NewOpaProfilesDisableGatekeeperInternalServerError() *OpaProfilesDisableGatekeeperInternalServerError {
	return &OpaProfilesDisableGatekeeperInternalServerError{}
}

/* OpaProfilesDisableGatekeeperInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesDisableGatekeeperInternalServerError struct {
}

func (o *OpaProfilesDisableGatekeeperInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/OpaProfiles/disablegatekeeper][%d] opaProfilesDisableGatekeeperInternalServerError ", 500)
}

func (o *OpaProfilesDisableGatekeeperInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}