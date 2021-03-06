// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SecurityGroupEditReader is a Reader for the SecurityGroupEdit structure.
type SecurityGroupEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityGroupEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSecurityGroupEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecurityGroupEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecurityGroupEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecurityGroupEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecurityGroupEditOK creates a SecurityGroupEditOK with default headers values
func NewSecurityGroupEditOK() *SecurityGroupEditOK {
	return &SecurityGroupEditOK{}
}

/* SecurityGroupEditOK describes a response with status code 200, with default header values.

Success
*/
type SecurityGroupEditOK struct {
	Payload models.Unit
}

func (o *SecurityGroupEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/edit][%d] securityGroupEditOK  %+v", 200, o.Payload)
}
func (o *SecurityGroupEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *SecurityGroupEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupEditBadRequest creates a SecurityGroupEditBadRequest with default headers values
func NewSecurityGroupEditBadRequest() *SecurityGroupEditBadRequest {
	return &SecurityGroupEditBadRequest{}
}

/* SecurityGroupEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SecurityGroupEditBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SecurityGroupEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/edit][%d] securityGroupEditBadRequest  %+v", 400, o.Payload)
}
func (o *SecurityGroupEditBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SecurityGroupEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupEditUnauthorized creates a SecurityGroupEditUnauthorized with default headers values
func NewSecurityGroupEditUnauthorized() *SecurityGroupEditUnauthorized {
	return &SecurityGroupEditUnauthorized{}
}

/* SecurityGroupEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SecurityGroupEditUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SecurityGroupEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/edit][%d] securityGroupEditUnauthorized  %+v", 401, o.Payload)
}
func (o *SecurityGroupEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SecurityGroupEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupEditForbidden creates a SecurityGroupEditForbidden with default headers values
func NewSecurityGroupEditForbidden() *SecurityGroupEditForbidden {
	return &SecurityGroupEditForbidden{}
}

/* SecurityGroupEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SecurityGroupEditForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SecurityGroupEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/edit][%d] securityGroupEditForbidden  %+v", 403, o.Payload)
}
func (o *SecurityGroupEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SecurityGroupEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupEditNotFound creates a SecurityGroupEditNotFound with default headers values
func NewSecurityGroupEditNotFound() *SecurityGroupEditNotFound {
	return &SecurityGroupEditNotFound{}
}

/* SecurityGroupEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SecurityGroupEditNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SecurityGroupEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/edit][%d] securityGroupEditNotFound  %+v", 404, o.Payload)
}
func (o *SecurityGroupEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SecurityGroupEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupEditInternalServerError creates a SecurityGroupEditInternalServerError with default headers values
func NewSecurityGroupEditInternalServerError() *SecurityGroupEditInternalServerError {
	return &SecurityGroupEditInternalServerError{}
}

/* SecurityGroupEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SecurityGroupEditInternalServerError struct {
}

func (o *SecurityGroupEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/edit][%d] securityGroupEditInternalServerError ", 500)
}

func (o *SecurityGroupEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
