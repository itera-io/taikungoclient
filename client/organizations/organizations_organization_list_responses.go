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

// OrganizationsOrganizationListReader is a Reader for the OrganizationsOrganizationList structure.
type OrganizationsOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsOrganizationListOK creates a OrganizationsOrganizationListOK with default headers values
func NewOrganizationsOrganizationListOK() *OrganizationsOrganizationListOK {
	return &OrganizationsOrganizationListOK{}
}

/* OrganizationsOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsOrganizationListOK struct {
	Payload []*models.CommonDropdownDto
}

func (o *OrganizationsOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/list][%d] organizationsOrganizationListOK  %+v", 200, o.Payload)
}
func (o *OrganizationsOrganizationListOK) GetPayload() []*models.CommonDropdownDto {
	return o.Payload
}

func (o *OrganizationsOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsOrganizationListBadRequest creates a OrganizationsOrganizationListBadRequest with default headers values
func NewOrganizationsOrganizationListBadRequest() *OrganizationsOrganizationListBadRequest {
	return &OrganizationsOrganizationListBadRequest{}
}

/* OrganizationsOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsOrganizationListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OrganizationsOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/list][%d] organizationsOrganizationListBadRequest  %+v", 400, o.Payload)
}
func (o *OrganizationsOrganizationListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationsOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsOrganizationListUnauthorized creates a OrganizationsOrganizationListUnauthorized with default headers values
func NewOrganizationsOrganizationListUnauthorized() *OrganizationsOrganizationListUnauthorized {
	return &OrganizationsOrganizationListUnauthorized{}
}

/* OrganizationsOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsOrganizationListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/list][%d] organizationsOrganizationListUnauthorized  %+v", 401, o.Payload)
}
func (o *OrganizationsOrganizationListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsOrganizationListForbidden creates a OrganizationsOrganizationListForbidden with default headers values
func NewOrganizationsOrganizationListForbidden() *OrganizationsOrganizationListForbidden {
	return &OrganizationsOrganizationListForbidden{}
}

/* OrganizationsOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsOrganizationListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/list][%d] organizationsOrganizationListForbidden  %+v", 403, o.Payload)
}
func (o *OrganizationsOrganizationListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsOrganizationListNotFound creates a OrganizationsOrganizationListNotFound with default headers values
func NewOrganizationsOrganizationListNotFound() *OrganizationsOrganizationListNotFound {
	return &OrganizationsOrganizationListNotFound{}
}

/* OrganizationsOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsOrganizationListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/list][%d] organizationsOrganizationListNotFound  %+v", 404, o.Payload)
}
func (o *OrganizationsOrganizationListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsOrganizationListInternalServerError creates a OrganizationsOrganizationListInternalServerError with default headers values
func NewOrganizationsOrganizationListInternalServerError() *OrganizationsOrganizationListInternalServerError {
	return &OrganizationsOrganizationListInternalServerError{}
}

/* OrganizationsOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsOrganizationListInternalServerError struct {
}

func (o *OrganizationsOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/list][%d] organizationsOrganizationListInternalServerError ", 500)
}

func (o *OrganizationsOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
