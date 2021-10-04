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

// OrganizationsDetailsReader is a Reader for the OrganizationsDetails structure.
type OrganizationsDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsDetailsOK creates a OrganizationsDetailsOK with default headers values
func NewOrganizationsDetailsOK() *OrganizationsDetailsOK {
	return &OrganizationsDetailsOK{}
}

/* OrganizationsDetailsOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsDetailsOK struct {
	Payload interface{}
}

func (o *OrganizationsDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsOK  %+v", 200, o.Payload)
}
func (o *OrganizationsDetailsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsBadRequest creates a OrganizationsDetailsBadRequest with default headers values
func NewOrganizationsDetailsBadRequest() *OrganizationsDetailsBadRequest {
	return &OrganizationsDetailsBadRequest{}
}

/* OrganizationsDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsDetailsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OrganizationsDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *OrganizationsDetailsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationsDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsUnauthorized creates a OrganizationsDetailsUnauthorized with default headers values
func NewOrganizationsDetailsUnauthorized() *OrganizationsDetailsUnauthorized {
	return &OrganizationsDetailsUnauthorized{}
}

/* OrganizationsDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *OrganizationsDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsForbidden creates a OrganizationsDetailsForbidden with default headers values
func NewOrganizationsDetailsForbidden() *OrganizationsDetailsForbidden {
	return &OrganizationsDetailsForbidden{}
}

/* OrganizationsDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsDetailsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsForbidden  %+v", 403, o.Payload)
}
func (o *OrganizationsDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsNotFound creates a OrganizationsDetailsNotFound with default headers values
func NewOrganizationsDetailsNotFound() *OrganizationsDetailsNotFound {
	return &OrganizationsDetailsNotFound{}
}

/* OrganizationsDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsDetailsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationsDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsNotFound  %+v", 404, o.Payload)
}
func (o *OrganizationsDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsInternalServerError creates a OrganizationsDetailsInternalServerError with default headers values
func NewOrganizationsDetailsInternalServerError() *OrganizationsDetailsInternalServerError {
	return &OrganizationsDetailsInternalServerError{}
}

/* OrganizationsDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsDetailsInternalServerError struct {
}

func (o *OrganizationsDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsInternalServerError ", 500)
}

func (o *OrganizationsDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
