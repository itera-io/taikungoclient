// Code generated by go-swagger; DO NOT EDIT.

package partner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// PartnerDeleteWhiteListDomainReader is a Reader for the PartnerDeleteWhiteListDomain structure.
type PartnerDeleteWhiteListDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerDeleteWhiteListDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerDeleteWhiteListDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerDeleteWhiteListDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerDeleteWhiteListDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerDeleteWhiteListDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerDeleteWhiteListDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerDeleteWhiteListDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerDeleteWhiteListDomainOK creates a PartnerDeleteWhiteListDomainOK with default headers values
func NewPartnerDeleteWhiteListDomainOK() *PartnerDeleteWhiteListDomainOK {
	return &PartnerDeleteWhiteListDomainOK{}
}

/* PartnerDeleteWhiteListDomainOK describes a response with status code 200, with default header values.

Success
*/
type PartnerDeleteWhiteListDomainOK struct {
	Payload models.Unit
}

func (o *PartnerDeleteWhiteListDomainOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/delete/whitelist/domain][%d] partnerDeleteWhiteListDomainOK  %+v", 200, o.Payload)
}
func (o *PartnerDeleteWhiteListDomainOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PartnerDeleteWhiteListDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDeleteWhiteListDomainBadRequest creates a PartnerDeleteWhiteListDomainBadRequest with default headers values
func NewPartnerDeleteWhiteListDomainBadRequest() *PartnerDeleteWhiteListDomainBadRequest {
	return &PartnerDeleteWhiteListDomainBadRequest{}
}

/* PartnerDeleteWhiteListDomainBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerDeleteWhiteListDomainBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *PartnerDeleteWhiteListDomainBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/delete/whitelist/domain][%d] partnerDeleteWhiteListDomainBadRequest  %+v", 400, o.Payload)
}
func (o *PartnerDeleteWhiteListDomainBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PartnerDeleteWhiteListDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDeleteWhiteListDomainUnauthorized creates a PartnerDeleteWhiteListDomainUnauthorized with default headers values
func NewPartnerDeleteWhiteListDomainUnauthorized() *PartnerDeleteWhiteListDomainUnauthorized {
	return &PartnerDeleteWhiteListDomainUnauthorized{}
}

/* PartnerDeleteWhiteListDomainUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerDeleteWhiteListDomainUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *PartnerDeleteWhiteListDomainUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/delete/whitelist/domain][%d] partnerDeleteWhiteListDomainUnauthorized  %+v", 401, o.Payload)
}
func (o *PartnerDeleteWhiteListDomainUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerDeleteWhiteListDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDeleteWhiteListDomainForbidden creates a PartnerDeleteWhiteListDomainForbidden with default headers values
func NewPartnerDeleteWhiteListDomainForbidden() *PartnerDeleteWhiteListDomainForbidden {
	return &PartnerDeleteWhiteListDomainForbidden{}
}

/* PartnerDeleteWhiteListDomainForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerDeleteWhiteListDomainForbidden struct {
	Payload *models.ProblemDetails
}

func (o *PartnerDeleteWhiteListDomainForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/delete/whitelist/domain][%d] partnerDeleteWhiteListDomainForbidden  %+v", 403, o.Payload)
}
func (o *PartnerDeleteWhiteListDomainForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerDeleteWhiteListDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDeleteWhiteListDomainNotFound creates a PartnerDeleteWhiteListDomainNotFound with default headers values
func NewPartnerDeleteWhiteListDomainNotFound() *PartnerDeleteWhiteListDomainNotFound {
	return &PartnerDeleteWhiteListDomainNotFound{}
}

/* PartnerDeleteWhiteListDomainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerDeleteWhiteListDomainNotFound struct {
	Payload *models.ProblemDetails
}

func (o *PartnerDeleteWhiteListDomainNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/delete/whitelist/domain][%d] partnerDeleteWhiteListDomainNotFound  %+v", 404, o.Payload)
}
func (o *PartnerDeleteWhiteListDomainNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PartnerDeleteWhiteListDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerDeleteWhiteListDomainInternalServerError creates a PartnerDeleteWhiteListDomainInternalServerError with default headers values
func NewPartnerDeleteWhiteListDomainInternalServerError() *PartnerDeleteWhiteListDomainInternalServerError {
	return &PartnerDeleteWhiteListDomainInternalServerError{}
}

/* PartnerDeleteWhiteListDomainInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerDeleteWhiteListDomainInternalServerError struct {
}

func (o *PartnerDeleteWhiteListDomainInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/delete/whitelist/domain][%d] partnerDeleteWhiteListDomainInternalServerError ", 500)
}

func (o *PartnerDeleteWhiteListDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
