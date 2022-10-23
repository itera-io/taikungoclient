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

// UserTokenListReader is a Reader for the UserTokenList structure.
type UserTokenListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserTokenListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserTokenListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserTokenListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserTokenListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserTokenListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserTokenListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserTokenListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserTokenListOK creates a UserTokenListOK with default headers values
func NewUserTokenListOK() *UserTokenListOK {
	return &UserTokenListOK{}
}

/*
UserTokenListOK describes a response with status code 200, with default header values.

Success
*/
type UserTokenListOK struct {
	Payload []*models.UserTokensListDto
}

// IsSuccess returns true when this user token list o k response has a 2xx status code
func (o *UserTokenListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user token list o k response has a 3xx status code
func (o *UserTokenListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token list o k response has a 4xx status code
func (o *UserTokenListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token list o k response has a 5xx status code
func (o *UserTokenListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user token list o k response a status code equal to that given
func (o *UserTokenListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserTokenListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListOK  %+v", 200, o.Payload)
}

func (o *UserTokenListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListOK  %+v", 200, o.Payload)
}

func (o *UserTokenListOK) GetPayload() []*models.UserTokensListDto {
	return o.Payload
}

func (o *UserTokenListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenListBadRequest creates a UserTokenListBadRequest with default headers values
func NewUserTokenListBadRequest() *UserTokenListBadRequest {
	return &UserTokenListBadRequest{}
}

/*
UserTokenListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserTokenListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this user token list bad request response has a 2xx status code
func (o *UserTokenListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token list bad request response has a 3xx status code
func (o *UserTokenListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token list bad request response has a 4xx status code
func (o *UserTokenListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token list bad request response has a 5xx status code
func (o *UserTokenListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user token list bad request response a status code equal to that given
func (o *UserTokenListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserTokenListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *UserTokenListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenListUnauthorized creates a UserTokenListUnauthorized with default headers values
func NewUserTokenListUnauthorized() *UserTokenListUnauthorized {
	return &UserTokenListUnauthorized{}
}

/*
UserTokenListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserTokenListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token list unauthorized response has a 2xx status code
func (o *UserTokenListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token list unauthorized response has a 3xx status code
func (o *UserTokenListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token list unauthorized response has a 4xx status code
func (o *UserTokenListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token list unauthorized response has a 5xx status code
func (o *UserTokenListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user token list unauthorized response a status code equal to that given
func (o *UserTokenListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserTokenListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenListForbidden creates a UserTokenListForbidden with default headers values
func NewUserTokenListForbidden() *UserTokenListForbidden {
	return &UserTokenListForbidden{}
}

/*
UserTokenListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserTokenListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token list forbidden response has a 2xx status code
func (o *UserTokenListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token list forbidden response has a 3xx status code
func (o *UserTokenListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token list forbidden response has a 4xx status code
func (o *UserTokenListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token list forbidden response has a 5xx status code
func (o *UserTokenListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user token list forbidden response a status code equal to that given
func (o *UserTokenListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserTokenListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenListNotFound creates a UserTokenListNotFound with default headers values
func NewUserTokenListNotFound() *UserTokenListNotFound {
	return &UserTokenListNotFound{}
}

/*
UserTokenListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserTokenListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this user token list not found response has a 2xx status code
func (o *UserTokenListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token list not found response has a 3xx status code
func (o *UserTokenListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token list not found response has a 4xx status code
func (o *UserTokenListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token list not found response has a 5xx status code
func (o *UserTokenListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user token list not found response a status code equal to that given
func (o *UserTokenListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserTokenListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UserTokenListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenListInternalServerError creates a UserTokenListInternalServerError with default headers values
func NewUserTokenListInternalServerError() *UserTokenListInternalServerError {
	return &UserTokenListInternalServerError{}
}

/*
UserTokenListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserTokenListInternalServerError struct {
}

// IsSuccess returns true when this user token list internal server error response has a 2xx status code
func (o *UserTokenListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token list internal server error response has a 3xx status code
func (o *UserTokenListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token list internal server error response has a 4xx status code
func (o *UserTokenListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token list internal server error response has a 5xx status code
func (o *UserTokenListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user token list internal server error response a status code equal to that given
func (o *UserTokenListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserTokenListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListInternalServerError ", 500)
}

func (o *UserTokenListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/list][%d] userTokenListInternalServerError ", 500)
}

func (o *UserTokenListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
