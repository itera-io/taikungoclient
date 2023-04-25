// Code generated by go-swagger; DO NOT EDIT.

package user_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UserTokenAvailableEndpointListReader is a Reader for the UserTokenAvailableEndpointList structure.
type UserTokenAvailableEndpointListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserTokenAvailableEndpointListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserTokenAvailableEndpointListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserTokenAvailableEndpointListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserTokenAvailableEndpointListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserTokenAvailableEndpointListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserTokenAvailableEndpointListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserTokenAvailableEndpointListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserTokenAvailableEndpointListOK creates a UserTokenAvailableEndpointListOK with default headers values
func NewUserTokenAvailableEndpointListOK() *UserTokenAvailableEndpointListOK {
	return &UserTokenAvailableEndpointListOK{}
}

/*
UserTokenAvailableEndpointListOK describes a response with status code 200, with default header values.

Success
*/
type UserTokenAvailableEndpointListOK struct {
	Payload *models.AvailableEndpointsList
}

// IsSuccess returns true when this user token available endpoint list o k response has a 2xx status code
func (o *UserTokenAvailableEndpointListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user token available endpoint list o k response has a 3xx status code
func (o *UserTokenAvailableEndpointListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token available endpoint list o k response has a 4xx status code
func (o *UserTokenAvailableEndpointListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token available endpoint list o k response has a 5xx status code
func (o *UserTokenAvailableEndpointListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user token available endpoint list o k response a status code equal to that given
func (o *UserTokenAvailableEndpointListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user token available endpoint list o k response
func (o *UserTokenAvailableEndpointListOK) Code() int {
	return 200
}

func (o *UserTokenAvailableEndpointListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListOK  %+v", 200, o.Payload)
}

func (o *UserTokenAvailableEndpointListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListOK  %+v", 200, o.Payload)
}

func (o *UserTokenAvailableEndpointListOK) GetPayload() *models.AvailableEndpointsList {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailableEndpointsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListBadRequest creates a UserTokenAvailableEndpointListBadRequest with default headers values
func NewUserTokenAvailableEndpointListBadRequest() *UserTokenAvailableEndpointListBadRequest {
	return &UserTokenAvailableEndpointListBadRequest{}
}

/*
UserTokenAvailableEndpointListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserTokenAvailableEndpointListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this user token available endpoint list bad request response has a 2xx status code
func (o *UserTokenAvailableEndpointListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token available endpoint list bad request response has a 3xx status code
func (o *UserTokenAvailableEndpointListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token available endpoint list bad request response has a 4xx status code
func (o *UserTokenAvailableEndpointListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token available endpoint list bad request response has a 5xx status code
func (o *UserTokenAvailableEndpointListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user token available endpoint list bad request response a status code equal to that given
func (o *UserTokenAvailableEndpointListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user token available endpoint list bad request response
func (o *UserTokenAvailableEndpointListBadRequest) Code() int {
	return 400
}

func (o *UserTokenAvailableEndpointListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenAvailableEndpointListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenAvailableEndpointListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListUnauthorized creates a UserTokenAvailableEndpointListUnauthorized with default headers values
func NewUserTokenAvailableEndpointListUnauthorized() *UserTokenAvailableEndpointListUnauthorized {
	return &UserTokenAvailableEndpointListUnauthorized{}
}

/*
UserTokenAvailableEndpointListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserTokenAvailableEndpointListUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this user token available endpoint list unauthorized response has a 2xx status code
func (o *UserTokenAvailableEndpointListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token available endpoint list unauthorized response has a 3xx status code
func (o *UserTokenAvailableEndpointListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token available endpoint list unauthorized response has a 4xx status code
func (o *UserTokenAvailableEndpointListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token available endpoint list unauthorized response has a 5xx status code
func (o *UserTokenAvailableEndpointListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user token available endpoint list unauthorized response a status code equal to that given
func (o *UserTokenAvailableEndpointListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user token available endpoint list unauthorized response
func (o *UserTokenAvailableEndpointListUnauthorized) Code() int {
	return 401
}

func (o *UserTokenAvailableEndpointListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenAvailableEndpointListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenAvailableEndpointListUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListForbidden creates a UserTokenAvailableEndpointListForbidden with default headers values
func NewUserTokenAvailableEndpointListForbidden() *UserTokenAvailableEndpointListForbidden {
	return &UserTokenAvailableEndpointListForbidden{}
}

/*
UserTokenAvailableEndpointListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserTokenAvailableEndpointListForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this user token available endpoint list forbidden response has a 2xx status code
func (o *UserTokenAvailableEndpointListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token available endpoint list forbidden response has a 3xx status code
func (o *UserTokenAvailableEndpointListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token available endpoint list forbidden response has a 4xx status code
func (o *UserTokenAvailableEndpointListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token available endpoint list forbidden response has a 5xx status code
func (o *UserTokenAvailableEndpointListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user token available endpoint list forbidden response a status code equal to that given
func (o *UserTokenAvailableEndpointListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user token available endpoint list forbidden response
func (o *UserTokenAvailableEndpointListForbidden) Code() int {
	return 403
}

func (o *UserTokenAvailableEndpointListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenAvailableEndpointListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenAvailableEndpointListForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListNotFound creates a UserTokenAvailableEndpointListNotFound with default headers values
func NewUserTokenAvailableEndpointListNotFound() *UserTokenAvailableEndpointListNotFound {
	return &UserTokenAvailableEndpointListNotFound{}
}

/*
UserTokenAvailableEndpointListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserTokenAvailableEndpointListNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this user token available endpoint list not found response has a 2xx status code
func (o *UserTokenAvailableEndpointListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token available endpoint list not found response has a 3xx status code
func (o *UserTokenAvailableEndpointListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token available endpoint list not found response has a 4xx status code
func (o *UserTokenAvailableEndpointListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token available endpoint list not found response has a 5xx status code
func (o *UserTokenAvailableEndpointListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user token available endpoint list not found response a status code equal to that given
func (o *UserTokenAvailableEndpointListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user token available endpoint list not found response
func (o *UserTokenAvailableEndpointListNotFound) Code() int {
	return 404
}

func (o *UserTokenAvailableEndpointListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenAvailableEndpointListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenAvailableEndpointListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListInternalServerError creates a UserTokenAvailableEndpointListInternalServerError with default headers values
func NewUserTokenAvailableEndpointListInternalServerError() *UserTokenAvailableEndpointListInternalServerError {
	return &UserTokenAvailableEndpointListInternalServerError{}
}

/*
UserTokenAvailableEndpointListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserTokenAvailableEndpointListInternalServerError struct {
}

// IsSuccess returns true when this user token available endpoint list internal server error response has a 2xx status code
func (o *UserTokenAvailableEndpointListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token available endpoint list internal server error response has a 3xx status code
func (o *UserTokenAvailableEndpointListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token available endpoint list internal server error response has a 4xx status code
func (o *UserTokenAvailableEndpointListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token available endpoint list internal server error response has a 5xx status code
func (o *UserTokenAvailableEndpointListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user token available endpoint list internal server error response a status code equal to that given
func (o *UserTokenAvailableEndpointListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user token available endpoint list internal server error response
func (o *UserTokenAvailableEndpointListInternalServerError) Code() int {
	return 500
}

func (o *UserTokenAvailableEndpointListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListInternalServerError ", 500)
}

func (o *UserTokenAvailableEndpointListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListInternalServerError ", 500)
}

func (o *UserTokenAvailableEndpointListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
