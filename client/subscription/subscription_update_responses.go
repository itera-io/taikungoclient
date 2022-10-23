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

// SubscriptionUpdateReader is a Reader for the SubscriptionUpdate structure.
type SubscriptionUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubscriptionUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubscriptionUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscriptionUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscriptionUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscriptionUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionUpdateOK creates a SubscriptionUpdateOK with default headers values
func NewSubscriptionUpdateOK() *SubscriptionUpdateOK {
	return &SubscriptionUpdateOK{}
}

/*
SubscriptionUpdateOK describes a response with status code 200, with default header values.

Success
*/
type SubscriptionUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this subscription update o k response has a 2xx status code
func (o *SubscriptionUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription update o k response has a 3xx status code
func (o *SubscriptionUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription update o k response has a 4xx status code
func (o *SubscriptionUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription update o k response has a 5xx status code
func (o *SubscriptionUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription update o k response a status code equal to that given
func (o *SubscriptionUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SubscriptionUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateOK  %+v", 200, o.Payload)
}

func (o *SubscriptionUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateOK  %+v", 200, o.Payload)
}

func (o *SubscriptionUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *SubscriptionUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionUpdateBadRequest creates a SubscriptionUpdateBadRequest with default headers values
func NewSubscriptionUpdateBadRequest() *SubscriptionUpdateBadRequest {
	return &SubscriptionUpdateBadRequest{}
}

/*
SubscriptionUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SubscriptionUpdateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this subscription update bad request response has a 2xx status code
func (o *SubscriptionUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription update bad request response has a 3xx status code
func (o *SubscriptionUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription update bad request response has a 4xx status code
func (o *SubscriptionUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription update bad request response has a 5xx status code
func (o *SubscriptionUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription update bad request response a status code equal to that given
func (o *SubscriptionUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SubscriptionUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionUpdateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *SubscriptionUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionUpdateUnauthorized creates a SubscriptionUpdateUnauthorized with default headers values
func NewSubscriptionUpdateUnauthorized() *SubscriptionUpdateUnauthorized {
	return &SubscriptionUpdateUnauthorized{}
}

/*
SubscriptionUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubscriptionUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription update unauthorized response has a 2xx status code
func (o *SubscriptionUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription update unauthorized response has a 3xx status code
func (o *SubscriptionUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription update unauthorized response has a 4xx status code
func (o *SubscriptionUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription update unauthorized response has a 5xx status code
func (o *SubscriptionUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription update unauthorized response a status code equal to that given
func (o *SubscriptionUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SubscriptionUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionUpdateForbidden creates a SubscriptionUpdateForbidden with default headers values
func NewSubscriptionUpdateForbidden() *SubscriptionUpdateForbidden {
	return &SubscriptionUpdateForbidden{}
}

/*
SubscriptionUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubscriptionUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription update forbidden response has a 2xx status code
func (o *SubscriptionUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription update forbidden response has a 3xx status code
func (o *SubscriptionUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription update forbidden response has a 4xx status code
func (o *SubscriptionUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription update forbidden response has a 5xx status code
func (o *SubscriptionUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription update forbidden response a status code equal to that given
func (o *SubscriptionUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubscriptionUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionUpdateNotFound creates a SubscriptionUpdateNotFound with default headers values
func NewSubscriptionUpdateNotFound() *SubscriptionUpdateNotFound {
	return &SubscriptionUpdateNotFound{}
}

/*
SubscriptionUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubscriptionUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription update not found response has a 2xx status code
func (o *SubscriptionUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription update not found response has a 3xx status code
func (o *SubscriptionUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription update not found response has a 4xx status code
func (o *SubscriptionUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription update not found response has a 5xx status code
func (o *SubscriptionUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription update not found response a status code equal to that given
func (o *SubscriptionUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubscriptionUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionUpdateInternalServerError creates a SubscriptionUpdateInternalServerError with default headers values
func NewSubscriptionUpdateInternalServerError() *SubscriptionUpdateInternalServerError {
	return &SubscriptionUpdateInternalServerError{}
}

/*
SubscriptionUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SubscriptionUpdateInternalServerError struct {
}

// IsSuccess returns true when this subscription update internal server error response has a 2xx status code
func (o *SubscriptionUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription update internal server error response has a 3xx status code
func (o *SubscriptionUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription update internal server error response has a 4xx status code
func (o *SubscriptionUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription update internal server error response has a 5xx status code
func (o *SubscriptionUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this subscription update internal server error response a status code equal to that given
func (o *SubscriptionUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SubscriptionUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateInternalServerError ", 500)
}

func (o *SubscriptionUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Subscription/update][%d] subscriptionUpdateInternalServerError ", 500)
}

func (o *SubscriptionUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
