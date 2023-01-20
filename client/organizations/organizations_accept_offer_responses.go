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

// OrganizationsAcceptOfferReader is a Reader for the OrganizationsAcceptOffer structure.
type OrganizationsAcceptOfferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsAcceptOfferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsAcceptOfferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsAcceptOfferBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsAcceptOfferUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsAcceptOfferForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsAcceptOfferNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsAcceptOfferInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsAcceptOfferOK creates a OrganizationsAcceptOfferOK with default headers values
func NewOrganizationsAcceptOfferOK() *OrganizationsAcceptOfferOK {
	return &OrganizationsAcceptOfferOK{}
}

/*
OrganizationsAcceptOfferOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsAcceptOfferOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this organizations accept offer o k response has a 2xx status code
func (o *OrganizationsAcceptOfferOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations accept offer o k response has a 3xx status code
func (o *OrganizationsAcceptOfferOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations accept offer o k response has a 4xx status code
func (o *OrganizationsAcceptOfferOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations accept offer o k response has a 5xx status code
func (o *OrganizationsAcceptOfferOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations accept offer o k response a status code equal to that given
func (o *OrganizationsAcceptOfferOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the organizations accept offer o k response
func (o *OrganizationsAcceptOfferOK) Code() int {
	return 200
}

func (o *OrganizationsAcceptOfferOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferOK  %+v", 200, o.Payload)
}

func (o *OrganizationsAcceptOfferOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferOK  %+v", 200, o.Payload)
}

func (o *OrganizationsAcceptOfferOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OrganizationsAcceptOfferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsAcceptOfferBadRequest creates a OrganizationsAcceptOfferBadRequest with default headers values
func NewOrganizationsAcceptOfferBadRequest() *OrganizationsAcceptOfferBadRequest {
	return &OrganizationsAcceptOfferBadRequest{}
}

/*
OrganizationsAcceptOfferBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsAcceptOfferBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations accept offer bad request response has a 2xx status code
func (o *OrganizationsAcceptOfferBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations accept offer bad request response has a 3xx status code
func (o *OrganizationsAcceptOfferBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations accept offer bad request response has a 4xx status code
func (o *OrganizationsAcceptOfferBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations accept offer bad request response has a 5xx status code
func (o *OrganizationsAcceptOfferBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations accept offer bad request response a status code equal to that given
func (o *OrganizationsAcceptOfferBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the organizations accept offer bad request response
func (o *OrganizationsAcceptOfferBadRequest) Code() int {
	return 400
}

func (o *OrganizationsAcceptOfferBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsAcceptOfferBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsAcceptOfferBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsAcceptOfferBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsAcceptOfferUnauthorized creates a OrganizationsAcceptOfferUnauthorized with default headers values
func NewOrganizationsAcceptOfferUnauthorized() *OrganizationsAcceptOfferUnauthorized {
	return &OrganizationsAcceptOfferUnauthorized{}
}

/*
OrganizationsAcceptOfferUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsAcceptOfferUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations accept offer unauthorized response has a 2xx status code
func (o *OrganizationsAcceptOfferUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations accept offer unauthorized response has a 3xx status code
func (o *OrganizationsAcceptOfferUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations accept offer unauthorized response has a 4xx status code
func (o *OrganizationsAcceptOfferUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations accept offer unauthorized response has a 5xx status code
func (o *OrganizationsAcceptOfferUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations accept offer unauthorized response a status code equal to that given
func (o *OrganizationsAcceptOfferUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the organizations accept offer unauthorized response
func (o *OrganizationsAcceptOfferUnauthorized) Code() int {
	return 401
}

func (o *OrganizationsAcceptOfferUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsAcceptOfferUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsAcceptOfferUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsAcceptOfferUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsAcceptOfferForbidden creates a OrganizationsAcceptOfferForbidden with default headers values
func NewOrganizationsAcceptOfferForbidden() *OrganizationsAcceptOfferForbidden {
	return &OrganizationsAcceptOfferForbidden{}
}

/*
OrganizationsAcceptOfferForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsAcceptOfferForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations accept offer forbidden response has a 2xx status code
func (o *OrganizationsAcceptOfferForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations accept offer forbidden response has a 3xx status code
func (o *OrganizationsAcceptOfferForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations accept offer forbidden response has a 4xx status code
func (o *OrganizationsAcceptOfferForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations accept offer forbidden response has a 5xx status code
func (o *OrganizationsAcceptOfferForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations accept offer forbidden response a status code equal to that given
func (o *OrganizationsAcceptOfferForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the organizations accept offer forbidden response
func (o *OrganizationsAcceptOfferForbidden) Code() int {
	return 403
}

func (o *OrganizationsAcceptOfferForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsAcceptOfferForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsAcceptOfferForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsAcceptOfferForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsAcceptOfferNotFound creates a OrganizationsAcceptOfferNotFound with default headers values
func NewOrganizationsAcceptOfferNotFound() *OrganizationsAcceptOfferNotFound {
	return &OrganizationsAcceptOfferNotFound{}
}

/*
OrganizationsAcceptOfferNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsAcceptOfferNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations accept offer not found response has a 2xx status code
func (o *OrganizationsAcceptOfferNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations accept offer not found response has a 3xx status code
func (o *OrganizationsAcceptOfferNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations accept offer not found response has a 4xx status code
func (o *OrganizationsAcceptOfferNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations accept offer not found response has a 5xx status code
func (o *OrganizationsAcceptOfferNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations accept offer not found response a status code equal to that given
func (o *OrganizationsAcceptOfferNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the organizations accept offer not found response
func (o *OrganizationsAcceptOfferNotFound) Code() int {
	return 404
}

func (o *OrganizationsAcceptOfferNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsAcceptOfferNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsAcceptOfferNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsAcceptOfferNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsAcceptOfferInternalServerError creates a OrganizationsAcceptOfferInternalServerError with default headers values
func NewOrganizationsAcceptOfferInternalServerError() *OrganizationsAcceptOfferInternalServerError {
	return &OrganizationsAcceptOfferInternalServerError{}
}

/*
OrganizationsAcceptOfferInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsAcceptOfferInternalServerError struct {
}

// IsSuccess returns true when this organizations accept offer internal server error response has a 2xx status code
func (o *OrganizationsAcceptOfferInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations accept offer internal server error response has a 3xx status code
func (o *OrganizationsAcceptOfferInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations accept offer internal server error response has a 4xx status code
func (o *OrganizationsAcceptOfferInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations accept offer internal server error response has a 5xx status code
func (o *OrganizationsAcceptOfferInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations accept offer internal server error response a status code equal to that given
func (o *OrganizationsAcceptOfferInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the organizations accept offer internal server error response
func (o *OrganizationsAcceptOfferInternalServerError) Code() int {
	return 500
}

func (o *OrganizationsAcceptOfferInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferInternalServerError ", 500)
}

func (o *OrganizationsAcceptOfferInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/accept-offer][%d] organizationsAcceptOfferInternalServerError ", 500)
}

func (o *OrganizationsAcceptOfferInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
