// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OrganizationsListReader is a Reader for the OrganizationsList structure.
type OrganizationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsListOK creates a OrganizationsListOK with default headers values
func NewOrganizationsListOK() *OrganizationsListOK {
	return &OrganizationsListOK{}
}

/* OrganizationsListOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsListOK struct {
	Payload *models.OrganizationsList
}

func (o *OrganizationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListOK  %+v", 200, o.Payload)
}
func (o *OrganizationsListOK) GetPayload() *models.OrganizationsList {
	return o.Payload
}

func (o *OrganizationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListBadRequest creates a OrganizationsListBadRequest with default headers values
func NewOrganizationsListBadRequest() *OrganizationsListBadRequest {
	return &OrganizationsListBadRequest{}
}

/* OrganizationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OrganizationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListBadRequest  %+v", 400, o.Payload)
}
func (o *OrganizationsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListUnauthorized creates a OrganizationsListUnauthorized with default headers values
func NewOrganizationsListUnauthorized() *OrganizationsListUnauthorized {
	return &OrganizationsListUnauthorized{}
}

/* OrganizationsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListUnauthorized  %+v", 401, o.Payload)
}
func (o *OrganizationsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListForbidden creates a OrganizationsListForbidden with default headers values
func NewOrganizationsListForbidden() *OrganizationsListForbidden {
	return &OrganizationsListForbidden{}
}

/* OrganizationsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListForbidden  %+v", 403, o.Payload)
}
func (o *OrganizationsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListNotFound creates a OrganizationsListNotFound with default headers values
func NewOrganizationsListNotFound() *OrganizationsListNotFound {
	return &OrganizationsListNotFound{}
}

/* OrganizationsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListNotFound  %+v", 404, o.Payload)
}
func (o *OrganizationsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListInternalServerError creates a OrganizationsListInternalServerError with default headers values
func NewOrganizationsListInternalServerError() *OrganizationsListInternalServerError {
	return &OrganizationsListInternalServerError{}
}

/* OrganizationsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsListInternalServerError struct {
}

func (o *OrganizationsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListInternalServerError ", 500)
}

func (o *OrganizationsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
