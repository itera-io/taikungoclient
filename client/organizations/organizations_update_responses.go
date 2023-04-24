// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrganizationsUpdateReader is a Reader for the OrganizationsUpdate structure.
type OrganizationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsUpdateOK creates a OrganizationsUpdateOK with default headers values
func NewOrganizationsUpdateOK() *OrganizationsUpdateOK {
	return &OrganizationsUpdateOK{}
}

/*
OrganizationsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsUpdateOK struct {
}

// IsSuccess returns true when this organizations update o k response has a 2xx status code
func (o *OrganizationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations update o k response has a 3xx status code
func (o *OrganizationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update o k response has a 4xx status code
func (o *OrganizationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations update o k response has a 5xx status code
func (o *OrganizationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update o k response a status code equal to that given
func (o *OrganizationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the organizations update o k response
func (o *OrganizationsUpdateOK) Code() int {
	return 200
}

func (o *OrganizationsUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateOK ", 200)
}

func (o *OrganizationsUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateOK ", 200)
}

func (o *OrganizationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrganizationsUpdateBadRequest creates a OrganizationsUpdateBadRequest with default headers values
func NewOrganizationsUpdateBadRequest() *OrganizationsUpdateBadRequest {
	return &OrganizationsUpdateBadRequest{}
}

/*
OrganizationsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations update bad request response has a 2xx status code
func (o *OrganizationsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update bad request response has a 3xx status code
func (o *OrganizationsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update bad request response has a 4xx status code
func (o *OrganizationsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update bad request response has a 5xx status code
func (o *OrganizationsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update bad request response a status code equal to that given
func (o *OrganizationsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the organizations update bad request response
func (o *OrganizationsUpdateBadRequest) Code() int {
	return 400
}

func (o *OrganizationsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdateUnauthorized creates a OrganizationsUpdateUnauthorized with default headers values
func NewOrganizationsUpdateUnauthorized() *OrganizationsUpdateUnauthorized {
	return &OrganizationsUpdateUnauthorized{}
}

/*
OrganizationsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsUpdateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations update unauthorized response has a 2xx status code
func (o *OrganizationsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update unauthorized response has a 3xx status code
func (o *OrganizationsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update unauthorized response has a 4xx status code
func (o *OrganizationsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update unauthorized response has a 5xx status code
func (o *OrganizationsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update unauthorized response a status code equal to that given
func (o *OrganizationsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the organizations update unauthorized response
func (o *OrganizationsUpdateUnauthorized) Code() int {
	return 401
}

func (o *OrganizationsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsUpdateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdateForbidden creates a OrganizationsUpdateForbidden with default headers values
func NewOrganizationsUpdateForbidden() *OrganizationsUpdateForbidden {
	return &OrganizationsUpdateForbidden{}
}

/*
OrganizationsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsUpdateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations update forbidden response has a 2xx status code
func (o *OrganizationsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update forbidden response has a 3xx status code
func (o *OrganizationsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update forbidden response has a 4xx status code
func (o *OrganizationsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update forbidden response has a 5xx status code
func (o *OrganizationsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update forbidden response a status code equal to that given
func (o *OrganizationsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the organizations update forbidden response
func (o *OrganizationsUpdateForbidden) Code() int {
	return 403
}

func (o *OrganizationsUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsUpdateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdateNotFound creates a OrganizationsUpdateNotFound with default headers values
func NewOrganizationsUpdateNotFound() *OrganizationsUpdateNotFound {
	return &OrganizationsUpdateNotFound{}
}

/*
OrganizationsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsUpdateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this organizations update not found response has a 2xx status code
func (o *OrganizationsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update not found response has a 3xx status code
func (o *OrganizationsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update not found response has a 4xx status code
func (o *OrganizationsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations update not found response has a 5xx status code
func (o *OrganizationsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations update not found response a status code equal to that given
func (o *OrganizationsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the organizations update not found response
func (o *OrganizationsUpdateNotFound) Code() int {
	return 404
}

func (o *OrganizationsUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsUpdateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsUpdateInternalServerError creates a OrganizationsUpdateInternalServerError with default headers values
func NewOrganizationsUpdateInternalServerError() *OrganizationsUpdateInternalServerError {
	return &OrganizationsUpdateInternalServerError{}
}

/*
OrganizationsUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsUpdateInternalServerError struct {
}

// IsSuccess returns true when this organizations update internal server error response has a 2xx status code
func (o *OrganizationsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations update internal server error response has a 3xx status code
func (o *OrganizationsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations update internal server error response has a 4xx status code
func (o *OrganizationsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations update internal server error response has a 5xx status code
func (o *OrganizationsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations update internal server error response a status code equal to that given
func (o *OrganizationsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the organizations update internal server error response
func (o *OrganizationsUpdateInternalServerError) Code() int {
	return 500
}

func (o *OrganizationsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateInternalServerError ", 500)
}

func (o *OrganizationsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/update][%d] organizationsUpdateInternalServerError ", 500)
}

func (o *OrganizationsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
