// Code generated by go-swagger; DO NOT EDIT.

package organization_subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OrganizationSubscriptionsListReader is a Reader for the OrganizationSubscriptionsList structure.
type OrganizationSubscriptionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationSubscriptionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationSubscriptionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationSubscriptionsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationSubscriptionsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationSubscriptionsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationSubscriptionsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationSubscriptionsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationSubscriptionsListOK creates a OrganizationSubscriptionsListOK with default headers values
func NewOrganizationSubscriptionsListOK() *OrganizationSubscriptionsListOK {
	return &OrganizationSubscriptionsListOK{}
}

/* OrganizationSubscriptionsListOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationSubscriptionsListOK struct {
	Payload *models.OrganizationSubscriptionList
}

func (o *OrganizationSubscriptionsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OrganizationSubscriptions][%d] organizationSubscriptionsListOK  %+v", 200, o.Payload)
}
func (o *OrganizationSubscriptionsListOK) GetPayload() *models.OrganizationSubscriptionList {
	return o.Payload
}

func (o *OrganizationSubscriptionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationSubscriptionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationSubscriptionsListBadRequest creates a OrganizationSubscriptionsListBadRequest with default headers values
func NewOrganizationSubscriptionsListBadRequest() *OrganizationSubscriptionsListBadRequest {
	return &OrganizationSubscriptionsListBadRequest{}
}

/* OrganizationSubscriptionsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationSubscriptionsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *OrganizationSubscriptionsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OrganizationSubscriptions][%d] organizationSubscriptionsListBadRequest  %+v", 400, o.Payload)
}
func (o *OrganizationSubscriptionsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationSubscriptionsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationSubscriptionsListUnauthorized creates a OrganizationSubscriptionsListUnauthorized with default headers values
func NewOrganizationSubscriptionsListUnauthorized() *OrganizationSubscriptionsListUnauthorized {
	return &OrganizationSubscriptionsListUnauthorized{}
}

/* OrganizationSubscriptionsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationSubscriptionsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationSubscriptionsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OrganizationSubscriptions][%d] organizationSubscriptionsListUnauthorized  %+v", 401, o.Payload)
}
func (o *OrganizationSubscriptionsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationSubscriptionsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationSubscriptionsListForbidden creates a OrganizationSubscriptionsListForbidden with default headers values
func NewOrganizationSubscriptionsListForbidden() *OrganizationSubscriptionsListForbidden {
	return &OrganizationSubscriptionsListForbidden{}
}

/* OrganizationSubscriptionsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationSubscriptionsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationSubscriptionsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OrganizationSubscriptions][%d] organizationSubscriptionsListForbidden  %+v", 403, o.Payload)
}
func (o *OrganizationSubscriptionsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationSubscriptionsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationSubscriptionsListNotFound creates a OrganizationSubscriptionsListNotFound with default headers values
func NewOrganizationSubscriptionsListNotFound() *OrganizationSubscriptionsListNotFound {
	return &OrganizationSubscriptionsListNotFound{}
}

/* OrganizationSubscriptionsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationSubscriptionsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *OrganizationSubscriptionsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OrganizationSubscriptions][%d] organizationSubscriptionsListNotFound  %+v", 404, o.Payload)
}
func (o *OrganizationSubscriptionsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationSubscriptionsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationSubscriptionsListInternalServerError creates a OrganizationSubscriptionsListInternalServerError with default headers values
func NewOrganizationSubscriptionsListInternalServerError() *OrganizationSubscriptionsListInternalServerError {
	return &OrganizationSubscriptionsListInternalServerError{}
}

/* OrganizationSubscriptionsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationSubscriptionsListInternalServerError struct {
}

func (o *OrganizationSubscriptionsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OrganizationSubscriptions][%d] organizationSubscriptionsListInternalServerError ", 500)
}

func (o *OrganizationSubscriptionsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
