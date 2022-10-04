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

// UserTokenDeleteReader is a Reader for the UserTokenDelete structure.
type UserTokenDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserTokenDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserTokenDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserTokenDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserTokenDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserTokenDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserTokenDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserTokenDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserTokenDeleteOK creates a UserTokenDeleteOK with default headers values
func NewUserTokenDeleteOK() *UserTokenDeleteOK {
	return &UserTokenDeleteOK{}
}

/*
UserTokenDeleteOK describes a response with status code 200, with default header values.

Success
*/
type UserTokenDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this user token delete o k response has a 2xx status code
func (o *UserTokenDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user token delete o k response has a 3xx status code
func (o *UserTokenDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token delete o k response has a 4xx status code
func (o *UserTokenDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token delete o k response has a 5xx status code
func (o *UserTokenDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user token delete o k response a status code equal to that given
func (o *UserTokenDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserTokenDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteOK  %+v", 200, o.Payload)
}

func (o *UserTokenDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteOK  %+v", 200, o.Payload)
}

func (o *UserTokenDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserTokenDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenDeleteBadRequest creates a UserTokenDeleteBadRequest with default headers values
func NewUserTokenDeleteBadRequest() *UserTokenDeleteBadRequest {
	return &UserTokenDeleteBadRequest{}
}

/*
UserTokenDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserTokenDeleteBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token delete bad request response has a 2xx status code
func (o *UserTokenDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token delete bad request response has a 3xx status code
func (o *UserTokenDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token delete bad request response has a 4xx status code
func (o *UserTokenDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token delete bad request response has a 5xx status code
func (o *UserTokenDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user token delete bad request response a status code equal to that given
func (o *UserTokenDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserTokenDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenDeleteBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenDeleteUnauthorized creates a UserTokenDeleteUnauthorized with default headers values
func NewUserTokenDeleteUnauthorized() *UserTokenDeleteUnauthorized {
	return &UserTokenDeleteUnauthorized{}
}

/*
UserTokenDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserTokenDeleteUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token delete unauthorized response has a 2xx status code
func (o *UserTokenDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token delete unauthorized response has a 3xx status code
func (o *UserTokenDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token delete unauthorized response has a 4xx status code
func (o *UserTokenDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token delete unauthorized response has a 5xx status code
func (o *UserTokenDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user token delete unauthorized response a status code equal to that given
func (o *UserTokenDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserTokenDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenDeleteUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenDeleteForbidden creates a UserTokenDeleteForbidden with default headers values
func NewUserTokenDeleteForbidden() *UserTokenDeleteForbidden {
	return &UserTokenDeleteForbidden{}
}

/*
UserTokenDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserTokenDeleteForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token delete forbidden response has a 2xx status code
func (o *UserTokenDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token delete forbidden response has a 3xx status code
func (o *UserTokenDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token delete forbidden response has a 4xx status code
func (o *UserTokenDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token delete forbidden response has a 5xx status code
func (o *UserTokenDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user token delete forbidden response a status code equal to that given
func (o *UserTokenDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserTokenDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenDeleteForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenDeleteNotFound creates a UserTokenDeleteNotFound with default headers values
func NewUserTokenDeleteNotFound() *UserTokenDeleteNotFound {
	return &UserTokenDeleteNotFound{}
}

/*
UserTokenDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserTokenDeleteNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token delete not found response has a 2xx status code
func (o *UserTokenDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token delete not found response has a 3xx status code
func (o *UserTokenDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token delete not found response has a 4xx status code
func (o *UserTokenDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token delete not found response has a 5xx status code
func (o *UserTokenDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user token delete not found response a status code equal to that given
func (o *UserTokenDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserTokenDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenDeleteNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenDeleteInternalServerError creates a UserTokenDeleteInternalServerError with default headers values
func NewUserTokenDeleteInternalServerError() *UserTokenDeleteInternalServerError {
	return &UserTokenDeleteInternalServerError{}
}

/*
UserTokenDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserTokenDeleteInternalServerError struct {
}

// IsSuccess returns true when this user token delete internal server error response has a 2xx status code
func (o *UserTokenDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token delete internal server error response has a 3xx status code
func (o *UserTokenDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token delete internal server error response has a 4xx status code
func (o *UserTokenDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token delete internal server error response has a 5xx status code
func (o *UserTokenDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user token delete internal server error response a status code equal to that given
func (o *UserTokenDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserTokenDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteInternalServerError ", 500)
}

func (o *UserTokenDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserToken/delete/{id}][%d] userTokenDeleteInternalServerError ", 500)
}

func (o *UserTokenDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
