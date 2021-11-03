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

// AdminAddBalanceReader is a Reader for the AdminAddBalance structure.
type AdminAddBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminAddBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminAddBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminAddBalanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminAddBalanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminAddBalanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminAddBalanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAdminAddBalanceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminAddBalanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminAddBalanceOK creates a AdminAddBalanceOK with default headers values
func NewAdminAddBalanceOK() *AdminAddBalanceOK {
	return &AdminAddBalanceOK{}
}

/* AdminAddBalanceOK describes a response with status code 200, with default header values.

Success
*/
type AdminAddBalanceOK struct {
	Payload models.Unit
}

func (o *AdminAddBalanceOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceOK  %+v", 200, o.Payload)
}
func (o *AdminAddBalanceOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminAddBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddBalanceBadRequest creates a AdminAddBalanceBadRequest with default headers values
func NewAdminAddBalanceBadRequest() *AdminAddBalanceBadRequest {
	return &AdminAddBalanceBadRequest{}
}

/* AdminAddBalanceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminAddBalanceBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AdminAddBalanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceBadRequest  %+v", 400, o.Payload)
}
func (o *AdminAddBalanceBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminAddBalanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddBalanceUnauthorized creates a AdminAddBalanceUnauthorized with default headers values
func NewAdminAddBalanceUnauthorized() *AdminAddBalanceUnauthorized {
	return &AdminAddBalanceUnauthorized{}
}

/* AdminAddBalanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminAddBalanceUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AdminAddBalanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminAddBalanceUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminAddBalanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddBalanceForbidden creates a AdminAddBalanceForbidden with default headers values
func NewAdminAddBalanceForbidden() *AdminAddBalanceForbidden {
	return &AdminAddBalanceForbidden{}
}

/* AdminAddBalanceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminAddBalanceForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AdminAddBalanceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceForbidden  %+v", 403, o.Payload)
}
func (o *AdminAddBalanceForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminAddBalanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddBalanceNotFound creates a AdminAddBalanceNotFound with default headers values
func NewAdminAddBalanceNotFound() *AdminAddBalanceNotFound {
	return &AdminAddBalanceNotFound{}
}

/* AdminAddBalanceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminAddBalanceNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AdminAddBalanceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceNotFound  %+v", 404, o.Payload)
}
func (o *AdminAddBalanceNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminAddBalanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddBalanceTooManyRequests creates a AdminAddBalanceTooManyRequests with default headers values
func NewAdminAddBalanceTooManyRequests() *AdminAddBalanceTooManyRequests {
	return &AdminAddBalanceTooManyRequests{}
}

/* AdminAddBalanceTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AdminAddBalanceTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AdminAddBalanceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceTooManyRequests  %+v", 429, o.Payload)
}
func (o *AdminAddBalanceTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminAddBalanceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddBalanceInternalServerError creates a AdminAddBalanceInternalServerError with default headers values
func NewAdminAddBalanceInternalServerError() *AdminAddBalanceInternalServerError {
	return &AdminAddBalanceInternalServerError{}
}

/* AdminAddBalanceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminAddBalanceInternalServerError struct {
}

func (o *AdminAddBalanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/organizations/add/balance][%d] adminAddBalanceInternalServerError ", 500)
}

func (o *AdminAddBalanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
