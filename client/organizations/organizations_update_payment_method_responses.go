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

// OrganizationsUpdatePaymentMethodReader is a Reader for the OrganizationsUpdatePaymentMethod structure.
type OrganizationsUpdatePaymentMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsUpdatePaymentMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsUpdatePaymentMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsUpdatePaymentMethodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsUpdatePaymentMethodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsUpdatePaymentMethodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsUpdatePaymentMethodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsUpdatePaymentMethodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsUpdatePaymentMethodOK creates a OrganizationsUpdatePaymentMethodOK with default headers values
func NewOrganizationsUpdatePaymentMethodOK() *OrganizationsUpdatePaymentMethodOK {
	return &OrganizationsUpdatePaymentMethodOK{}
}

/*
OrganizationsUpdatePaymentMethodOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsUpdatePaymentMethodOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this organizations update payment method o k response has a 2xx status code
func (o *OrganizationsUpdatePaymentMethodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations update payment method o k response has a 3xx status code
func (o *OrganizationsUpdatePaymentMethodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update payment method o k response has a 4xx status code
func (o *OrganizationsUpdatePaymentMethodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations update payment method o k response has a 5xx status code
func (o *OrganizationsUpdatePaymentMethodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update payment method o k response a status code equal to that given
func (o *OrganizationsUpdatePaymentMethodOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsUpdatePaymentMethodOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodOK  %+v", 200, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodOK  %+v", 200, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OrganizationsUpdatePaymentMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdatePaymentMethodBadRequest creates a OrganizationsUpdatePaymentMethodBadRequest with default headers values
func NewOrganizationsUpdatePaymentMethodBadRequest() *OrganizationsUpdatePaymentMethodBadRequest {
	return &OrganizationsUpdatePaymentMethodBadRequest{}
}

/*
OrganizationsUpdatePaymentMethodBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsUpdatePaymentMethodBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this organizations update payment method bad request response has a 2xx status code
func (o *OrganizationsUpdatePaymentMethodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update payment method bad request response has a 3xx status code
func (o *OrganizationsUpdatePaymentMethodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update payment method bad request response has a 4xx status code
func (o *OrganizationsUpdatePaymentMethodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update payment method bad request response has a 5xx status code
func (o *OrganizationsUpdatePaymentMethodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update payment method bad request response a status code equal to that given
func (o *OrganizationsUpdatePaymentMethodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OrganizationsUpdatePaymentMethodBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *OrganizationsUpdatePaymentMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdatePaymentMethodUnauthorized creates a OrganizationsUpdatePaymentMethodUnauthorized with default headers values
func NewOrganizationsUpdatePaymentMethodUnauthorized() *OrganizationsUpdatePaymentMethodUnauthorized {
	return &OrganizationsUpdatePaymentMethodUnauthorized{}
}

/*
OrganizationsUpdatePaymentMethodUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsUpdatePaymentMethodUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations update payment method unauthorized response has a 2xx status code
func (o *OrganizationsUpdatePaymentMethodUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update payment method unauthorized response has a 3xx status code
func (o *OrganizationsUpdatePaymentMethodUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update payment method unauthorized response has a 4xx status code
func (o *OrganizationsUpdatePaymentMethodUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update payment method unauthorized response has a 5xx status code
func (o *OrganizationsUpdatePaymentMethodUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update payment method unauthorized response a status code equal to that given
func (o *OrganizationsUpdatePaymentMethodUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OrganizationsUpdatePaymentMethodUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsUpdatePaymentMethodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdatePaymentMethodForbidden creates a OrganizationsUpdatePaymentMethodForbidden with default headers values
func NewOrganizationsUpdatePaymentMethodForbidden() *OrganizationsUpdatePaymentMethodForbidden {
	return &OrganizationsUpdatePaymentMethodForbidden{}
}

/*
OrganizationsUpdatePaymentMethodForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsUpdatePaymentMethodForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations update payment method forbidden response has a 2xx status code
func (o *OrganizationsUpdatePaymentMethodForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update payment method forbidden response has a 3xx status code
func (o *OrganizationsUpdatePaymentMethodForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update payment method forbidden response has a 4xx status code
func (o *OrganizationsUpdatePaymentMethodForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update payment method forbidden response has a 5xx status code
func (o *OrganizationsUpdatePaymentMethodForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update payment method forbidden response a status code equal to that given
func (o *OrganizationsUpdatePaymentMethodForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsUpdatePaymentMethodForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsUpdatePaymentMethodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdatePaymentMethodNotFound creates a OrganizationsUpdatePaymentMethodNotFound with default headers values
func NewOrganizationsUpdatePaymentMethodNotFound() *OrganizationsUpdatePaymentMethodNotFound {
	return &OrganizationsUpdatePaymentMethodNotFound{}
}

/*
OrganizationsUpdatePaymentMethodNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsUpdatePaymentMethodNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations update payment method not found response has a 2xx status code
func (o *OrganizationsUpdatePaymentMethodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update payment method not found response has a 3xx status code
func (o *OrganizationsUpdatePaymentMethodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update payment method not found response has a 4xx status code
func (o *OrganizationsUpdatePaymentMethodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update payment method not found response has a 5xx status code
func (o *OrganizationsUpdatePaymentMethodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update payment method not found response a status code equal to that given
func (o *OrganizationsUpdatePaymentMethodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OrganizationsUpdatePaymentMethodNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsUpdatePaymentMethodNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsUpdatePaymentMethodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdatePaymentMethodInternalServerError creates a OrganizationsUpdatePaymentMethodInternalServerError with default headers values
func NewOrganizationsUpdatePaymentMethodInternalServerError() *OrganizationsUpdatePaymentMethodInternalServerError {
	return &OrganizationsUpdatePaymentMethodInternalServerError{}
}

/*
OrganizationsUpdatePaymentMethodInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsUpdatePaymentMethodInternalServerError struct {
}

// IsSuccess returns true when this organizations update payment method internal server error response has a 2xx status code
func (o *OrganizationsUpdatePaymentMethodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update payment method internal server error response has a 3xx status code
func (o *OrganizationsUpdatePaymentMethodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update payment method internal server error response has a 4xx status code
func (o *OrganizationsUpdatePaymentMethodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations update payment method internal server error response has a 5xx status code
func (o *OrganizationsUpdatePaymentMethodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations update payment method internal server error response a status code equal to that given
func (o *OrganizationsUpdatePaymentMethodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OrganizationsUpdatePaymentMethodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodInternalServerError ", 500)
}

func (o *OrganizationsUpdatePaymentMethodInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/updatepaymentmethod][%d] organizationsUpdatePaymentMethodInternalServerError ", 500)
}

func (o *OrganizationsUpdatePaymentMethodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
