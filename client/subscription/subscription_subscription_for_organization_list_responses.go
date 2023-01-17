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

// SubscriptionSubscriptionForOrganizationListReader is a Reader for the SubscriptionSubscriptionForOrganizationList structure.
type SubscriptionSubscriptionForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionSubscriptionForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionSubscriptionForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubscriptionSubscriptionForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubscriptionSubscriptionForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubscriptionSubscriptionForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubscriptionSubscriptionForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubscriptionSubscriptionForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionSubscriptionForOrganizationListOK creates a SubscriptionSubscriptionForOrganizationListOK with default headers values
func NewSubscriptionSubscriptionForOrganizationListOK() *SubscriptionSubscriptionForOrganizationListOK {
	return &SubscriptionSubscriptionForOrganizationListOK{}
}

/*
SubscriptionSubscriptionForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type SubscriptionSubscriptionForOrganizationListOK struct {
	Payload []*models.ListForOrganizationEditDto
}

// IsSuccess returns true when this subscription subscription for organization list o k response has a 2xx status code
func (o *SubscriptionSubscriptionForOrganizationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription subscription for organization list o k response has a 3xx status code
func (o *SubscriptionSubscriptionForOrganizationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for organization list o k response has a 4xx status code
func (o *SubscriptionSubscriptionForOrganizationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription subscription for organization list o k response has a 5xx status code
func (o *SubscriptionSubscriptionForOrganizationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for organization list o k response a status code equal to that given
func (o *SubscriptionSubscriptionForOrganizationListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the subscription subscription for organization list o k response
func (o *SubscriptionSubscriptionForOrganizationListOK) Code() int {
	return 200
}

func (o *SubscriptionSubscriptionForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListOK) GetPayload() []*models.ListForOrganizationEditDto {
	return o.Payload
}

func (o *SubscriptionSubscriptionForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForOrganizationListBadRequest creates a SubscriptionSubscriptionForOrganizationListBadRequest with default headers values
func NewSubscriptionSubscriptionForOrganizationListBadRequest() *SubscriptionSubscriptionForOrganizationListBadRequest {
	return &SubscriptionSubscriptionForOrganizationListBadRequest{}
}

/*
SubscriptionSubscriptionForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SubscriptionSubscriptionForOrganizationListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription subscription for organization list bad request response has a 2xx status code
func (o *SubscriptionSubscriptionForOrganizationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for organization list bad request response has a 3xx status code
func (o *SubscriptionSubscriptionForOrganizationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for organization list bad request response has a 4xx status code
func (o *SubscriptionSubscriptionForOrganizationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for organization list bad request response has a 5xx status code
func (o *SubscriptionSubscriptionForOrganizationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for organization list bad request response a status code equal to that given
func (o *SubscriptionSubscriptionForOrganizationListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the subscription subscription for organization list bad request response
func (o *SubscriptionSubscriptionForOrganizationListBadRequest) Code() int {
	return 400
}

func (o *SubscriptionSubscriptionForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionSubscriptionForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForOrganizationListUnauthorized creates a SubscriptionSubscriptionForOrganizationListUnauthorized with default headers values
func NewSubscriptionSubscriptionForOrganizationListUnauthorized() *SubscriptionSubscriptionForOrganizationListUnauthorized {
	return &SubscriptionSubscriptionForOrganizationListUnauthorized{}
}

/*
SubscriptionSubscriptionForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubscriptionSubscriptionForOrganizationListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription subscription for organization list unauthorized response has a 2xx status code
func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for organization list unauthorized response has a 3xx status code
func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for organization list unauthorized response has a 4xx status code
func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for organization list unauthorized response has a 5xx status code
func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for organization list unauthorized response a status code equal to that given
func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the subscription subscription for organization list unauthorized response
func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) Code() int {
	return 401
}

func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionSubscriptionForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForOrganizationListForbidden creates a SubscriptionSubscriptionForOrganizationListForbidden with default headers values
func NewSubscriptionSubscriptionForOrganizationListForbidden() *SubscriptionSubscriptionForOrganizationListForbidden {
	return &SubscriptionSubscriptionForOrganizationListForbidden{}
}

/*
SubscriptionSubscriptionForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubscriptionSubscriptionForOrganizationListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription subscription for organization list forbidden response has a 2xx status code
func (o *SubscriptionSubscriptionForOrganizationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for organization list forbidden response has a 3xx status code
func (o *SubscriptionSubscriptionForOrganizationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for organization list forbidden response has a 4xx status code
func (o *SubscriptionSubscriptionForOrganizationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for organization list forbidden response has a 5xx status code
func (o *SubscriptionSubscriptionForOrganizationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for organization list forbidden response a status code equal to that given
func (o *SubscriptionSubscriptionForOrganizationListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the subscription subscription for organization list forbidden response
func (o *SubscriptionSubscriptionForOrganizationListForbidden) Code() int {
	return 403
}

func (o *SubscriptionSubscriptionForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionSubscriptionForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForOrganizationListNotFound creates a SubscriptionSubscriptionForOrganizationListNotFound with default headers values
func NewSubscriptionSubscriptionForOrganizationListNotFound() *SubscriptionSubscriptionForOrganizationListNotFound {
	return &SubscriptionSubscriptionForOrganizationListNotFound{}
}

/*
SubscriptionSubscriptionForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubscriptionSubscriptionForOrganizationListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this subscription subscription for organization list not found response has a 2xx status code
func (o *SubscriptionSubscriptionForOrganizationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for organization list not found response has a 3xx status code
func (o *SubscriptionSubscriptionForOrganizationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for organization list not found response has a 4xx status code
func (o *SubscriptionSubscriptionForOrganizationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this subscription subscription for organization list not found response has a 5xx status code
func (o *SubscriptionSubscriptionForOrganizationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription subscription for organization list not found response a status code equal to that given
func (o *SubscriptionSubscriptionForOrganizationListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the subscription subscription for organization list not found response
func (o *SubscriptionSubscriptionForOrganizationListNotFound) Code() int {
	return 404
}

func (o *SubscriptionSubscriptionForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *SubscriptionSubscriptionForOrganizationListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SubscriptionSubscriptionForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscriptionSubscriptionForOrganizationListInternalServerError creates a SubscriptionSubscriptionForOrganizationListInternalServerError with default headers values
func NewSubscriptionSubscriptionForOrganizationListInternalServerError() *SubscriptionSubscriptionForOrganizationListInternalServerError {
	return &SubscriptionSubscriptionForOrganizationListInternalServerError{}
}

/*
SubscriptionSubscriptionForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SubscriptionSubscriptionForOrganizationListInternalServerError struct {
}

// IsSuccess returns true when this subscription subscription for organization list internal server error response has a 2xx status code
func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this subscription subscription for organization list internal server error response has a 3xx status code
func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription subscription for organization list internal server error response has a 4xx status code
func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription subscription for organization list internal server error response has a 5xx status code
func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this subscription subscription for organization list internal server error response a status code equal to that given
func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the subscription subscription for organization list internal server error response
func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) Code() int {
	return 500
}

func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListInternalServerError ", 500)
}

func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Subscription/boundlist][%d] subscriptionSubscriptionForOrganizationListInternalServerError ", 500)
}

func (o *SubscriptionSubscriptionForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
