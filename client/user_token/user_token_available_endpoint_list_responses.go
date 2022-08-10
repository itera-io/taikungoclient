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

/* UserTokenAvailableEndpointListOK describes a response with status code 200, with default header values.

Success
*/
type UserTokenAvailableEndpointListOK struct {
	Payload *models.AvailableEndpointsList
}

func (o *UserTokenAvailableEndpointListOK) Error() string {
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

/* UserTokenAvailableEndpointListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserTokenAvailableEndpointListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *UserTokenAvailableEndpointListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListBadRequest  %+v", 400, o.Payload)
}
func (o *UserTokenAvailableEndpointListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListUnauthorized creates a UserTokenAvailableEndpointListUnauthorized with default headers values
func NewUserTokenAvailableEndpointListUnauthorized() *UserTokenAvailableEndpointListUnauthorized {
	return &UserTokenAvailableEndpointListUnauthorized{}
}

/* UserTokenAvailableEndpointListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserTokenAvailableEndpointListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *UserTokenAvailableEndpointListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListUnauthorized  %+v", 401, o.Payload)
}
func (o *UserTokenAvailableEndpointListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListForbidden creates a UserTokenAvailableEndpointListForbidden with default headers values
func NewUserTokenAvailableEndpointListForbidden() *UserTokenAvailableEndpointListForbidden {
	return &UserTokenAvailableEndpointListForbidden{}
}

/* UserTokenAvailableEndpointListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserTokenAvailableEndpointListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *UserTokenAvailableEndpointListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListForbidden  %+v", 403, o.Payload)
}
func (o *UserTokenAvailableEndpointListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListNotFound creates a UserTokenAvailableEndpointListNotFound with default headers values
func NewUserTokenAvailableEndpointListNotFound() *UserTokenAvailableEndpointListNotFound {
	return &UserTokenAvailableEndpointListNotFound{}
}

/* UserTokenAvailableEndpointListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserTokenAvailableEndpointListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *UserTokenAvailableEndpointListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListNotFound  %+v", 404, o.Payload)
}
func (o *UserTokenAvailableEndpointListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserTokenAvailableEndpointListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenAvailableEndpointListInternalServerError creates a UserTokenAvailableEndpointListInternalServerError with default headers values
func NewUserTokenAvailableEndpointListInternalServerError() *UserTokenAvailableEndpointListInternalServerError {
	return &UserTokenAvailableEndpointListInternalServerError{}
}

/* UserTokenAvailableEndpointListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserTokenAvailableEndpointListInternalServerError struct {
}

func (o *UserTokenAvailableEndpointListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserToken/available-endpoints][%d] userTokenAvailableEndpointListInternalServerError ", 500)
}

func (o *UserTokenAvailableEndpointListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}