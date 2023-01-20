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

// OrganizationsListReader is a Reader for the OrganizationsList structure.
type OrganizationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsListOK creates a OrganizationsListOK with default headers values
func NewOrganizationsListOK() *OrganizationsListOK {
	return &OrganizationsListOK{}
}

/*
OrganizationsListOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsListOK struct {
	Payload *models.OrganizationsList
}

// IsSuccess returns true when this organizations list o k response has a 2xx status code
func (o *OrganizationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations list o k response has a 3xx status code
func (o *OrganizationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations list o k response has a 4xx status code
func (o *OrganizationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations list o k response has a 5xx status code
func (o *OrganizationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations list o k response a status code equal to that given
func (o *OrganizationsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the organizations list o k response
func (o *OrganizationsListOK) Code() int {
	return 200
}

func (o *OrganizationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListOK  %+v", 200, o.Payload)
}

func (o *OrganizationsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListOK  %+v", 200, o.Payload)
}

func (o *OrganizationsListOK) GetPayload() *models.OrganizationsList {
	return o.Payload
}

func (o *OrganizationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListBadRequest creates a OrganizationsListBadRequest with default headers values
func NewOrganizationsListBadRequest() *OrganizationsListBadRequest {
	return &OrganizationsListBadRequest{}
}

/*
OrganizationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations list bad request response has a 2xx status code
func (o *OrganizationsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations list bad request response has a 3xx status code
func (o *OrganizationsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations list bad request response has a 4xx status code
func (o *OrganizationsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations list bad request response has a 5xx status code
func (o *OrganizationsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations list bad request response a status code equal to that given
func (o *OrganizationsListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the organizations list bad request response
func (o *OrganizationsListBadRequest) Code() int {
	return 400
}

func (o *OrganizationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListUnauthorized creates a OrganizationsListUnauthorized with default headers values
func NewOrganizationsListUnauthorized() *OrganizationsListUnauthorized {
	return &OrganizationsListUnauthorized{}
}

/*
OrganizationsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations list unauthorized response has a 2xx status code
func (o *OrganizationsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations list unauthorized response has a 3xx status code
func (o *OrganizationsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations list unauthorized response has a 4xx status code
func (o *OrganizationsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations list unauthorized response has a 5xx status code
func (o *OrganizationsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations list unauthorized response a status code equal to that given
func (o *OrganizationsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the organizations list unauthorized response
func (o *OrganizationsListUnauthorized) Code() int {
	return 401
}

func (o *OrganizationsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListForbidden creates a OrganizationsListForbidden with default headers values
func NewOrganizationsListForbidden() *OrganizationsListForbidden {
	return &OrganizationsListForbidden{}
}

/*
OrganizationsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations list forbidden response has a 2xx status code
func (o *OrganizationsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations list forbidden response has a 3xx status code
func (o *OrganizationsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations list forbidden response has a 4xx status code
func (o *OrganizationsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations list forbidden response has a 5xx status code
func (o *OrganizationsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations list forbidden response a status code equal to that given
func (o *OrganizationsListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the organizations list forbidden response
func (o *OrganizationsListForbidden) Code() int {
	return 403
}

func (o *OrganizationsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListNotFound creates a OrganizationsListNotFound with default headers values
func NewOrganizationsListNotFound() *OrganizationsListNotFound {
	return &OrganizationsListNotFound{}
}

/*
OrganizationsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations list not found response has a 2xx status code
func (o *OrganizationsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations list not found response has a 3xx status code
func (o *OrganizationsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations list not found response has a 4xx status code
func (o *OrganizationsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations list not found response has a 5xx status code
func (o *OrganizationsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations list not found response a status code equal to that given
func (o *OrganizationsListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the organizations list not found response
func (o *OrganizationsListNotFound) Code() int {
	return 404
}

func (o *OrganizationsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsListInternalServerError creates a OrganizationsListInternalServerError with default headers values
func NewOrganizationsListInternalServerError() *OrganizationsListInternalServerError {
	return &OrganizationsListInternalServerError{}
}

/*
OrganizationsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsListInternalServerError struct {
}

// IsSuccess returns true when this organizations list internal server error response has a 2xx status code
func (o *OrganizationsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations list internal server error response has a 3xx status code
func (o *OrganizationsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations list internal server error response has a 4xx status code
func (o *OrganizationsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations list internal server error response has a 5xx status code
func (o *OrganizationsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations list internal server error response a status code equal to that given
func (o *OrganizationsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the organizations list internal server error response
func (o *OrganizationsListInternalServerError) Code() int {
	return 500
}

func (o *OrganizationsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListInternalServerError ", 500)
}

func (o *OrganizationsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations][%d] organizationsListInternalServerError ", 500)
}

func (o *OrganizationsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
