// Code generated by go-swagger; DO NOT EDIT.

package allowed_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AllowedHostEditReader is a Reader for the AllowedHostEdit structure.
type AllowedHostEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllowedHostEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllowedHostEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllowedHostEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAllowedHostEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAllowedHostEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllowedHostEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAllowedHostEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllowedHostEditOK creates a AllowedHostEditOK with default headers values
func NewAllowedHostEditOK() *AllowedHostEditOK {
	return &AllowedHostEditOK{}
}

/*
AllowedHostEditOK describes a response with status code 200, with default header values.

Success
*/
type AllowedHostEditOK struct {
}

// IsSuccess returns true when this allowed host edit o k response has a 2xx status code
func (o *AllowedHostEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this allowed host edit o k response has a 3xx status code
func (o *AllowedHostEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host edit o k response has a 4xx status code
func (o *AllowedHostEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this allowed host edit o k response has a 5xx status code
func (o *AllowedHostEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host edit o k response a status code equal to that given
func (o *AllowedHostEditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the allowed host edit o k response
func (o *AllowedHostEditOK) Code() int {
	return 200
}

func (o *AllowedHostEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditOK ", 200)
}

func (o *AllowedHostEditOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditOK ", 200)
}

func (o *AllowedHostEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAllowedHostEditBadRequest creates a AllowedHostEditBadRequest with default headers values
func NewAllowedHostEditBadRequest() *AllowedHostEditBadRequest {
	return &AllowedHostEditBadRequest{}
}

/*
AllowedHostEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllowedHostEditBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this allowed host edit bad request response has a 2xx status code
func (o *AllowedHostEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host edit bad request response has a 3xx status code
func (o *AllowedHostEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host edit bad request response has a 4xx status code
func (o *AllowedHostEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host edit bad request response has a 5xx status code
func (o *AllowedHostEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host edit bad request response a status code equal to that given
func (o *AllowedHostEditBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the allowed host edit bad request response
func (o *AllowedHostEditBadRequest) Code() int {
	return 400
}

func (o *AllowedHostEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditBadRequest  %+v", 400, o.Payload)
}

func (o *AllowedHostEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditBadRequest  %+v", 400, o.Payload)
}

func (o *AllowedHostEditBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *AllowedHostEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostEditUnauthorized creates a AllowedHostEditUnauthorized with default headers values
func NewAllowedHostEditUnauthorized() *AllowedHostEditUnauthorized {
	return &AllowedHostEditUnauthorized{}
}

/*
AllowedHostEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AllowedHostEditUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this allowed host edit unauthorized response has a 2xx status code
func (o *AllowedHostEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host edit unauthorized response has a 3xx status code
func (o *AllowedHostEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host edit unauthorized response has a 4xx status code
func (o *AllowedHostEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host edit unauthorized response has a 5xx status code
func (o *AllowedHostEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host edit unauthorized response a status code equal to that given
func (o *AllowedHostEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the allowed host edit unauthorized response
func (o *AllowedHostEditUnauthorized) Code() int {
	return 401
}

func (o *AllowedHostEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditUnauthorized  %+v", 401, o.Payload)
}

func (o *AllowedHostEditUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditUnauthorized  %+v", 401, o.Payload)
}

func (o *AllowedHostEditUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AllowedHostEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostEditForbidden creates a AllowedHostEditForbidden with default headers values
func NewAllowedHostEditForbidden() *AllowedHostEditForbidden {
	return &AllowedHostEditForbidden{}
}

/*
AllowedHostEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AllowedHostEditForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this allowed host edit forbidden response has a 2xx status code
func (o *AllowedHostEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host edit forbidden response has a 3xx status code
func (o *AllowedHostEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host edit forbidden response has a 4xx status code
func (o *AllowedHostEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host edit forbidden response has a 5xx status code
func (o *AllowedHostEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host edit forbidden response a status code equal to that given
func (o *AllowedHostEditForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the allowed host edit forbidden response
func (o *AllowedHostEditForbidden) Code() int {
	return 403
}

func (o *AllowedHostEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditForbidden  %+v", 403, o.Payload)
}

func (o *AllowedHostEditForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditForbidden  %+v", 403, o.Payload)
}

func (o *AllowedHostEditForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AllowedHostEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostEditNotFound creates a AllowedHostEditNotFound with default headers values
func NewAllowedHostEditNotFound() *AllowedHostEditNotFound {
	return &AllowedHostEditNotFound{}
}

/*
AllowedHostEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllowedHostEditNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this allowed host edit not found response has a 2xx status code
func (o *AllowedHostEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host edit not found response has a 3xx status code
func (o *AllowedHostEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host edit not found response has a 4xx status code
func (o *AllowedHostEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this allowed host edit not found response has a 5xx status code
func (o *AllowedHostEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this allowed host edit not found response a status code equal to that given
func (o *AllowedHostEditNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the allowed host edit not found response
func (o *AllowedHostEditNotFound) Code() int {
	return 404
}

func (o *AllowedHostEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditNotFound  %+v", 404, o.Payload)
}

func (o *AllowedHostEditNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditNotFound  %+v", 404, o.Payload)
}

func (o *AllowedHostEditNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AllowedHostEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllowedHostEditInternalServerError creates a AllowedHostEditInternalServerError with default headers values
func NewAllowedHostEditInternalServerError() *AllowedHostEditInternalServerError {
	return &AllowedHostEditInternalServerError{}
}

/*
AllowedHostEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AllowedHostEditInternalServerError struct {
}

// IsSuccess returns true when this allowed host edit internal server error response has a 2xx status code
func (o *AllowedHostEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this allowed host edit internal server error response has a 3xx status code
func (o *AllowedHostEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allowed host edit internal server error response has a 4xx status code
func (o *AllowedHostEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this allowed host edit internal server error response has a 5xx status code
func (o *AllowedHostEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this allowed host edit internal server error response a status code equal to that given
func (o *AllowedHostEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the allowed host edit internal server error response
func (o *AllowedHostEditInternalServerError) Code() int {
	return 500
}

func (o *AllowedHostEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditInternalServerError ", 500)
}

func (o *AllowedHostEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AllowedHost/edit/{id}][%d] allowedHostEditInternalServerError ", 500)
}

func (o *AllowedHostEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
