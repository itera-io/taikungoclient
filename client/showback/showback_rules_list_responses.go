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

// ShowbackRulesListReader is a Reader for the ShowbackRulesList structure.
type ShowbackRulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackRulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackRulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackRulesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackRulesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackRulesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackRulesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackRulesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackRulesListOK creates a ShowbackRulesListOK with default headers values
func NewShowbackRulesListOK() *ShowbackRulesListOK {
	return &ShowbackRulesListOK{}
}

/* ShowbackRulesListOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackRulesListOK struct {
	Payload *models.ShowbackRuleList
}

func (o *ShowbackRulesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/rules][%d] showbackRulesListOK  %+v", 200, o.Payload)
}
func (o *ShowbackRulesListOK) GetPayload() *models.ShowbackRuleList {
	return o.Payload
}

func (o *ShowbackRulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShowbackRuleList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListBadRequest creates a ShowbackRulesListBadRequest with default headers values
func NewShowbackRulesListBadRequest() *ShowbackRulesListBadRequest {
	return &ShowbackRulesListBadRequest{}
}

/* ShowbackRulesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackRulesListBadRequest struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackRulesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/rules][%d] showbackRulesListBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackRulesListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackRulesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListUnauthorized creates a ShowbackRulesListUnauthorized with default headers values
func NewShowbackRulesListUnauthorized() *ShowbackRulesListUnauthorized {
	return &ShowbackRulesListUnauthorized{}
}

/* ShowbackRulesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackRulesListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackRulesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/rules][%d] showbackRulesListUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackRulesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackRulesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListForbidden creates a ShowbackRulesListForbidden with default headers values
func NewShowbackRulesListForbidden() *ShowbackRulesListForbidden {
	return &ShowbackRulesListForbidden{}
}

/* ShowbackRulesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackRulesListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackRulesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/rules][%d] showbackRulesListForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackRulesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackRulesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListNotFound creates a ShowbackRulesListNotFound with default headers values
func NewShowbackRulesListNotFound() *ShowbackRulesListNotFound {
	return &ShowbackRulesListNotFound{}
}

/* ShowbackRulesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackRulesListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackRulesListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/rules][%d] showbackRulesListNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackRulesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackRulesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackRulesListInternalServerError creates a ShowbackRulesListInternalServerError with default headers values
func NewShowbackRulesListInternalServerError() *ShowbackRulesListInternalServerError {
	return &ShowbackRulesListInternalServerError{}
}

/* ShowbackRulesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackRulesListInternalServerError struct {
}

func (o *ShowbackRulesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/rules][%d] showbackRulesListInternalServerError ", 500)
}

func (o *ShowbackRulesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
