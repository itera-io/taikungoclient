// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SubscriptionListReader is a Reader for the SubscriptionList structure.
type SubscriptionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubscriptionListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubscriptionListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscriptionListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscriptionListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSubscriptionListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscriptionListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionListOK creates a SubscriptionListOK with default headers values
func NewSubscriptionListOK() *SubscriptionListOK {
	return &SubscriptionListOK{}
}

/* SubscriptionListOK describes a response with status code 200, with default header values.

Success
*/
type SubscriptionListOK struct {
	Payload *models.PrivateSubscriptionList
}

func (o *SubscriptionListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListOK  %+v", 200, o.Payload)
}
func (o *SubscriptionListOK) GetPayload() *models.PrivateSubscriptionList {
	return o.Payload
}

func (o *SubscriptionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSubscriptionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionListBadRequest creates a SubscriptionListBadRequest with default headers values
func NewSubscriptionListBadRequest() *SubscriptionListBadRequest {
	return &SubscriptionListBadRequest{}
}

/* SubscriptionListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SubscriptionListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SubscriptionListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListBadRequest  %+v", 400, o.Payload)
}
func (o *SubscriptionListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SubscriptionListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionListUnauthorized creates a SubscriptionListUnauthorized with default headers values
func NewSubscriptionListUnauthorized() *SubscriptionListUnauthorized {
	return &SubscriptionListUnauthorized{}
}

/* SubscriptionListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubscriptionListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SubscriptionListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListUnauthorized  %+v", 401, o.Payload)
}
func (o *SubscriptionListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionListForbidden creates a SubscriptionListForbidden with default headers values
func NewSubscriptionListForbidden() *SubscriptionListForbidden {
	return &SubscriptionListForbidden{}
}

/* SubscriptionListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubscriptionListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SubscriptionListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListForbidden  %+v", 403, o.Payload)
}
func (o *SubscriptionListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionListNotFound creates a SubscriptionListNotFound with default headers values
func NewSubscriptionListNotFound() *SubscriptionListNotFound {
	return &SubscriptionListNotFound{}
}

/* SubscriptionListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubscriptionListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SubscriptionListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListNotFound  %+v", 404, o.Payload)
}
func (o *SubscriptionListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionListTooManyRequests creates a SubscriptionListTooManyRequests with default headers values
func NewSubscriptionListTooManyRequests() *SubscriptionListTooManyRequests {
	return &SubscriptionListTooManyRequests{}
}

/* SubscriptionListTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type SubscriptionListTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *SubscriptionListTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListTooManyRequests  %+v", 429, o.Payload)
}
func (o *SubscriptionListTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionListInternalServerError creates a SubscriptionListInternalServerError with default headers values
func NewSubscriptionListInternalServerError() *SubscriptionListInternalServerError {
	return &SubscriptionListInternalServerError{}
}

/* SubscriptionListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SubscriptionListInternalServerError struct {
}

func (o *SubscriptionListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription][%d] subscriptionListInternalServerError ", 500)
}

func (o *SubscriptionListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
