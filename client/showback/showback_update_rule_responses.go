// Code generated by go-swagger; DO NOT EDIT.

package showback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackUpdateRuleReader is a Reader for the ShowbackUpdateRule structure.
type ShowbackUpdateRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackUpdateRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackUpdateRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackUpdateRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackUpdateRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackUpdateRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackUpdateRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackUpdateRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackUpdateRuleOK creates a ShowbackUpdateRuleOK with default headers values
func NewShowbackUpdateRuleOK() *ShowbackUpdateRuleOK {
	return &ShowbackUpdateRuleOK{}
}

/* ShowbackUpdateRuleOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackUpdateRuleOK struct {
	Payload models.Unit
}

func (o *ShowbackUpdateRuleOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/rule/update][%d] showbackUpdateRuleOK  %+v", 200, o.Payload)
}
func (o *ShowbackUpdateRuleOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ShowbackUpdateRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackUpdateRuleBadRequest creates a ShowbackUpdateRuleBadRequest with default headers values
func NewShowbackUpdateRuleBadRequest() *ShowbackUpdateRuleBadRequest {
	return &ShowbackUpdateRuleBadRequest{}
}

/* ShowbackUpdateRuleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackUpdateRuleBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackUpdateRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/rule/update][%d] showbackUpdateRuleBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackUpdateRuleBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackUpdateRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackUpdateRuleUnauthorized creates a ShowbackUpdateRuleUnauthorized with default headers values
func NewShowbackUpdateRuleUnauthorized() *ShowbackUpdateRuleUnauthorized {
	return &ShowbackUpdateRuleUnauthorized{}
}

/* ShowbackUpdateRuleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackUpdateRuleUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackUpdateRuleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/rule/update][%d] showbackUpdateRuleUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackUpdateRuleUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackUpdateRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackUpdateRuleForbidden creates a ShowbackUpdateRuleForbidden with default headers values
func NewShowbackUpdateRuleForbidden() *ShowbackUpdateRuleForbidden {
	return &ShowbackUpdateRuleForbidden{}
}

/* ShowbackUpdateRuleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackUpdateRuleForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackUpdateRuleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/rule/update][%d] showbackUpdateRuleForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackUpdateRuleForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackUpdateRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackUpdateRuleNotFound creates a ShowbackUpdateRuleNotFound with default headers values
func NewShowbackUpdateRuleNotFound() *ShowbackUpdateRuleNotFound {
	return &ShowbackUpdateRuleNotFound{}
}

/* ShowbackUpdateRuleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackUpdateRuleNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackUpdateRuleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/rule/update][%d] showbackUpdateRuleNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackUpdateRuleNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackUpdateRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackUpdateRuleInternalServerError creates a ShowbackUpdateRuleInternalServerError with default headers values
func NewShowbackUpdateRuleInternalServerError() *ShowbackUpdateRuleInternalServerError {
	return &ShowbackUpdateRuleInternalServerError{}
}

/* ShowbackUpdateRuleInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackUpdateRuleInternalServerError struct {
}

func (o *ShowbackUpdateRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Showback/rule/update][%d] showbackUpdateRuleInternalServerError ", 500)
}

func (o *ShowbackUpdateRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
