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

// SubscriptionCreateReader is a Reader for the SubscriptionCreate structure.
type SubscriptionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubscriptionCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubscriptionCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscriptionCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscriptionCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscriptionCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionCreateOK creates a SubscriptionCreateOK with default headers values
func NewSubscriptionCreateOK() *SubscriptionCreateOK {
	return &SubscriptionCreateOK{}
}

/*
SubscriptionCreateOK describes a response with status code 200, with default header values.

Success
*/
type SubscriptionCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this subscription create o k response has a 2xx status code
func (o *SubscriptionCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription create o k response has a 3xx status code
func (o *SubscriptionCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create o k response has a 4xx status code
func (o *SubscriptionCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription create o k response has a 5xx status code
func (o *SubscriptionCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription create o k response a status code equal to that given
func (o *SubscriptionCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateOK  %+v", 200, o.Payload)
}

func (o *SubscriptionCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateOK  %+v", 200, o.Payload)
}

func (o *SubscriptionCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *SubscriptionCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionCreateBadRequest creates a SubscriptionCreateBadRequest with default headers values
func NewSubscriptionCreateBadRequest() *SubscriptionCreateBadRequest {
	return &SubscriptionCreateBadRequest{}
}

/*
SubscriptionCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SubscriptionCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription create bad request response has a 2xx status code
func (o *SubscriptionCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription create bad request response has a 3xx status code
func (o *SubscriptionCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create bad request response has a 4xx status code
func (o *SubscriptionCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription create bad request response has a 5xx status code
func (o *SubscriptionCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription create bad request response a status code equal to that given
func (o *SubscriptionCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubscriptionCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionCreateUnauthorized creates a SubscriptionCreateUnauthorized with default headers values
func NewSubscriptionCreateUnauthorized() *SubscriptionCreateUnauthorized {
	return &SubscriptionCreateUnauthorized{}
}

/*
SubscriptionCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubscriptionCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription create unauthorized response has a 2xx status code
func (o *SubscriptionCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription create unauthorized response has a 3xx status code
func (o *SubscriptionCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create unauthorized response has a 4xx status code
func (o *SubscriptionCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription create unauthorized response has a 5xx status code
func (o *SubscriptionCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription create unauthorized response a status code equal to that given
func (o *SubscriptionCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SubscriptionCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionCreateForbidden creates a SubscriptionCreateForbidden with default headers values
func NewSubscriptionCreateForbidden() *SubscriptionCreateForbidden {
	return &SubscriptionCreateForbidden{}
}

/*
SubscriptionCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubscriptionCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription create forbidden response has a 2xx status code
func (o *SubscriptionCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription create forbidden response has a 3xx status code
func (o *SubscriptionCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create forbidden response has a 4xx status code
func (o *SubscriptionCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription create forbidden response has a 5xx status code
func (o *SubscriptionCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription create forbidden response a status code equal to that given
func (o *SubscriptionCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubscriptionCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionCreateNotFound creates a SubscriptionCreateNotFound with default headers values
func NewSubscriptionCreateNotFound() *SubscriptionCreateNotFound {
	return &SubscriptionCreateNotFound{}
}

/*
SubscriptionCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubscriptionCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription create not found response has a 2xx status code
func (o *SubscriptionCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription create not found response has a 3xx status code
func (o *SubscriptionCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create not found response has a 4xx status code
func (o *SubscriptionCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription create not found response has a 5xx status code
func (o *SubscriptionCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription create not found response a status code equal to that given
func (o *SubscriptionCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubscriptionCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionCreateInternalServerError creates a SubscriptionCreateInternalServerError with default headers values
func NewSubscriptionCreateInternalServerError() *SubscriptionCreateInternalServerError {
	return &SubscriptionCreateInternalServerError{}
}

/*
SubscriptionCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SubscriptionCreateInternalServerError struct {
}

// IsSuccess returns true when this subscription create internal server error response has a 2xx status code
func (o *SubscriptionCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription create internal server error response has a 3xx status code
func (o *SubscriptionCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create internal server error response has a 4xx status code
func (o *SubscriptionCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription create internal server error response has a 5xx status code
func (o *SubscriptionCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this subscription create internal server error response a status code equal to that given
func (o *SubscriptionCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubscriptionCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateInternalServerError ", 500)
}

func (o *SubscriptionCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/create][%d] subscriptionCreateInternalServerError ", 500)
}

func (o *SubscriptionCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
