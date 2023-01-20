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

// SubscriptionSubscriptionForLandingPageReader is a Reader for the SubscriptionSubscriptionForLandingPage structure.
type SubscriptionSubscriptionForLandingPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionSubscriptionForLandingPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionSubscriptionForLandingPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubscriptionSubscriptionForLandingPageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubscriptionSubscriptionForLandingPageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscriptionSubscriptionForLandingPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscriptionSubscriptionForLandingPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscriptionSubscriptionForLandingPageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionSubscriptionForLandingPageOK creates a SubscriptionSubscriptionForLandingPageOK with default headers values
func NewSubscriptionSubscriptionForLandingPageOK() *SubscriptionSubscriptionForLandingPageOK {
	return &SubscriptionSubscriptionForLandingPageOK{}
}

/*
SubscriptionSubscriptionForLandingPageOK describes a response with status code 200, with default header values.

Success
*/
type SubscriptionSubscriptionForLandingPageOK struct {
	Payload []*models.ListForLandingPageDto
}

// IsSuccess returns true when this subscription subscription for landing page o k response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription subscription for landing page o k response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page o k response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription subscription for landing page o k response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page o k response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the subscription subscription for landing page o k response
func (o *SubscriptionSubscriptionForLandingPageOK) Code() int {
	return 200
}

func (o *SubscriptionSubscriptionForLandingPageOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageOK  %+v", 200, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageOK  %+v", 200, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageOK) GetPayload() []*models.ListForLandingPageDto {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageBadRequest creates a SubscriptionSubscriptionForLandingPageBadRequest with default headers values
func NewSubscriptionSubscriptionForLandingPageBadRequest() *SubscriptionSubscriptionForLandingPageBadRequest {
	return &SubscriptionSubscriptionForLandingPageBadRequest{}
}

/*
SubscriptionSubscriptionForLandingPageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SubscriptionSubscriptionForLandingPageBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription subscription for landing page bad request response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page bad request response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page bad request response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page bad request response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page bad request response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the subscription subscription for landing page bad request response
func (o *SubscriptionSubscriptionForLandingPageBadRequest) Code() int {
	return 400
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageUnauthorized creates a SubscriptionSubscriptionForLandingPageUnauthorized with default headers values
func NewSubscriptionSubscriptionForLandingPageUnauthorized() *SubscriptionSubscriptionForLandingPageUnauthorized {
	return &SubscriptionSubscriptionForLandingPageUnauthorized{}
}

/*
SubscriptionSubscriptionForLandingPageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubscriptionSubscriptionForLandingPageUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription subscription for landing page unauthorized response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page unauthorized response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page unauthorized response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page unauthorized response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page unauthorized response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the subscription subscription for landing page unauthorized response
func (o *SubscriptionSubscriptionForLandingPageUnauthorized) Code() int {
	return 401
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageForbidden creates a SubscriptionSubscriptionForLandingPageForbidden with default headers values
func NewSubscriptionSubscriptionForLandingPageForbidden() *SubscriptionSubscriptionForLandingPageForbidden {
	return &SubscriptionSubscriptionForLandingPageForbidden{}
}

/*
SubscriptionSubscriptionForLandingPageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubscriptionSubscriptionForLandingPageForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription subscription for landing page forbidden response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page forbidden response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page forbidden response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page forbidden response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page forbidden response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the subscription subscription for landing page forbidden response
func (o *SubscriptionSubscriptionForLandingPageForbidden) Code() int {
	return 403
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageNotFound creates a SubscriptionSubscriptionForLandingPageNotFound with default headers values
func NewSubscriptionSubscriptionForLandingPageNotFound() *SubscriptionSubscriptionForLandingPageNotFound {
	return &SubscriptionSubscriptionForLandingPageNotFound{}
}

/*
SubscriptionSubscriptionForLandingPageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubscriptionSubscriptionForLandingPageNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this subscription subscription for landing page not found response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page not found response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page not found response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for landing page not found response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for landing page not found response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the subscription subscription for landing page not found response
func (o *SubscriptionSubscriptionForLandingPageNotFound) Code() int {
	return 404
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SubscriptionSubscriptionForLandingPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForLandingPageInternalServerError creates a SubscriptionSubscriptionForLandingPageInternalServerError with default headers values
func NewSubscriptionSubscriptionForLandingPageInternalServerError() *SubscriptionSubscriptionForLandingPageInternalServerError {
	return &SubscriptionSubscriptionForLandingPageInternalServerError{}
}

/*
SubscriptionSubscriptionForLandingPageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SubscriptionSubscriptionForLandingPageInternalServerError struct {
}

// IsSuccess returns true when this subscription subscription for landing page internal server error response has a 2xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for landing page internal server error response has a 3xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for landing page internal server error response has a 4xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription subscription for landing page internal server error response has a 5xx status code
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this subscription subscription for landing page internal server error response a status code equal to that given
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the subscription subscription for landing page internal server error response
func (o *SubscriptionSubscriptionForLandingPageInternalServerError) Code() int {
	return 500
}

func (o *SubscriptionSubscriptionForLandingPageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageInternalServerError ", 500)
}

func (o *SubscriptionSubscriptionForLandingPageInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/public][%d] subscriptionSubscriptionForLandingPageInternalServerError ", 500)
}

func (o *SubscriptionSubscriptionForLandingPageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
