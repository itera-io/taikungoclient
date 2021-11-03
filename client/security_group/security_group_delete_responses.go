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

// SecurityGroupDeleteReader is a Reader for the SecurityGroupDelete structure.
type SecurityGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityGroupDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSecurityGroupDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSecurityGroupDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSecurityGroupDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecurityGroupDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecurityGroupDeleteOK creates a SecurityGroupDeleteOK with default headers values
func NewSecurityGroupDeleteOK() *SecurityGroupDeleteOK {
	return &SecurityGroupDeleteOK{}
}

/* SecurityGroupDeleteOK describes a response with status code 200, with default header values.

Success
*/
type SecurityGroupDeleteOK struct {
	Payload models.Unit
}

func (o *SecurityGroupDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/delete][%d] securityGroupDeleteOK  %+v", 200, o.Payload)
}
func (o *SecurityGroupDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *SecurityGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteBadRequest creates a SecurityGroupDeleteBadRequest with default headers values
func NewSecurityGroupDeleteBadRequest() *SecurityGroupDeleteBadRequest {
	return &SecurityGroupDeleteBadRequest{}
}

/* SecurityGroupDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SecurityGroupDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SecurityGroupDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/delete][%d] securityGroupDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *SecurityGroupDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SecurityGroupDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteUnauthorized creates a SecurityGroupDeleteUnauthorized with default headers values
func NewSecurityGroupDeleteUnauthorized() *SecurityGroupDeleteUnauthorized {
	return &SecurityGroupDeleteUnauthorized{}
}

/* SecurityGroupDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SecurityGroupDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SecurityGroupDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/delete][%d] securityGroupDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *SecurityGroupDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SecurityGroupDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteForbidden creates a SecurityGroupDeleteForbidden with default headers values
func NewSecurityGroupDeleteForbidden() *SecurityGroupDeleteForbidden {
	return &SecurityGroupDeleteForbidden{}
}

/* SecurityGroupDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SecurityGroupDeleteForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SecurityGroupDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/delete][%d] securityGroupDeleteForbidden  %+v", 403, o.Payload)
}
func (o *SecurityGroupDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SecurityGroupDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteNotFound creates a SecurityGroupDeleteNotFound with default headers values
func NewSecurityGroupDeleteNotFound() *SecurityGroupDeleteNotFound {
	return &SecurityGroupDeleteNotFound{}
}

/* SecurityGroupDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SecurityGroupDeleteNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SecurityGroupDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/delete][%d] securityGroupDeleteNotFound  %+v", 404, o.Payload)
}
func (o *SecurityGroupDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SecurityGroupDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteInternalServerError creates a SecurityGroupDeleteInternalServerError with default headers values
func NewSecurityGroupDeleteInternalServerError() *SecurityGroupDeleteInternalServerError {
	return &SecurityGroupDeleteInternalServerError{}
}

/* SecurityGroupDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SecurityGroupDeleteInternalServerError struct {
}

func (o *SecurityGroupDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SecurityGroup/delete][%d] securityGroupDeleteInternalServerError ", 500)
}

func (o *SecurityGroupDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
