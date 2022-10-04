// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminUpdateUserPasswordReader is a Reader for the AdminUpdateUserPassword structure.
type AdminUpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUpdateUserPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateUserPasswordOK creates a AdminUpdateUserPasswordOK with default headers values
func NewAdminUpdateUserPasswordOK() *AdminUpdateUserPasswordOK {
	return &AdminUpdateUserPasswordOK{}
}

/*
AdminUpdateUserPasswordOK describes a response with status code 200, with default header values.

Success
*/
type AdminUpdateUserPasswordOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this admin update user password o k response has a 2xx status code
func (o *AdminUpdateUserPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin update user password o k response has a 3xx status code
func (o *AdminUpdateUserPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password o k response has a 4xx status code
func (o *AdminUpdateUserPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user password o k response has a 5xx status code
func (o *AdminUpdateUserPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password o k response a status code equal to that given
func (o *AdminUpdateUserPasswordOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminUpdateUserPasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordOK  %+v", 200, o.Payload)
}

func (o *AdminUpdateUserPasswordOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordOK  %+v", 200, o.Payload)
}

func (o *AdminUpdateUserPasswordOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminUpdateUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordBadRequest creates a AdminUpdateUserPasswordBadRequest with default headers values
func NewAdminUpdateUserPasswordBadRequest() *AdminUpdateUserPasswordBadRequest {
	return &AdminUpdateUserPasswordBadRequest{}
}

/*
AdminUpdateUserPasswordBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUpdateUserPasswordBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this admin update user password bad request response has a 2xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password bad request response has a 3xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password bad request response has a 4xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password bad request response has a 5xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password bad request response a status code equal to that given
func (o *AdminUpdateUserPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminUpdateUserPasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateUserPasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateUserPasswordBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminUpdateUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordUnauthorized creates a AdminUpdateUserPasswordUnauthorized with default headers values
func NewAdminUpdateUserPasswordUnauthorized() *AdminUpdateUserPasswordUnauthorized {
	return &AdminUpdateUserPasswordUnauthorized{}
}

/*
AdminUpdateUserPasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUpdateUserPasswordUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this admin update user password unauthorized response has a 2xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password unauthorized response has a 3xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password unauthorized response has a 4xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password unauthorized response has a 5xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password unauthorized response a status code equal to that given
func (o *AdminUpdateUserPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminUpdateUserPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateUserPasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateUserPasswordUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AdminUpdateUserPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordForbidden creates a AdminUpdateUserPasswordForbidden with default headers values
func NewAdminUpdateUserPasswordForbidden() *AdminUpdateUserPasswordForbidden {
	return &AdminUpdateUserPasswordForbidden{}
}

/*
AdminUpdateUserPasswordForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUpdateUserPasswordForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this admin update user password forbidden response has a 2xx status code
func (o *AdminUpdateUserPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password forbidden response has a 3xx status code
func (o *AdminUpdateUserPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password forbidden response has a 4xx status code
func (o *AdminUpdateUserPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password forbidden response has a 5xx status code
func (o *AdminUpdateUserPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password forbidden response a status code equal to that given
func (o *AdminUpdateUserPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminUpdateUserPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordForbidden  %+v", 403, o.Payload)
}

func (o *AdminUpdateUserPasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordForbidden  %+v", 403, o.Payload)
}

func (o *AdminUpdateUserPasswordForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AdminUpdateUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordNotFound creates a AdminUpdateUserPasswordNotFound with default headers values
func NewAdminUpdateUserPasswordNotFound() *AdminUpdateUserPasswordNotFound {
	return &AdminUpdateUserPasswordNotFound{}
}

/*
AdminUpdateUserPasswordNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUpdateUserPasswordNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this admin update user password not found response has a 2xx status code
func (o *AdminUpdateUserPasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password not found response has a 3xx status code
func (o *AdminUpdateUserPasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password not found response has a 4xx status code
func (o *AdminUpdateUserPasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password not found response has a 5xx status code
func (o *AdminUpdateUserPasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password not found response a status code equal to that given
func (o *AdminUpdateUserPasswordNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminUpdateUserPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordNotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateUserPasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordNotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateUserPasswordNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AdminUpdateUserPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordInternalServerError creates a AdminUpdateUserPasswordInternalServerError with default headers values
func NewAdminUpdateUserPasswordInternalServerError() *AdminUpdateUserPasswordInternalServerError {
	return &AdminUpdateUserPasswordInternalServerError{}
}

/*
AdminUpdateUserPasswordInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUpdateUserPasswordInternalServerError struct {
}

// IsSuccess returns true when this admin update user password internal server error response has a 2xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password internal server error response has a 3xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password internal server error response has a 4xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user password internal server error response has a 5xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin update user password internal server error response a status code equal to that given
func (o *AdminUpdateUserPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminUpdateUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordInternalServerError ", 500)
}

func (o *AdminUpdateUserPasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/password][%d] adminUpdateUserPasswordInternalServerError ", 500)
}

func (o *AdminUpdateUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
