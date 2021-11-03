// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SearchUsersListReader is a Reader for the SearchUsersList structure.
type SearchUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchUsersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchUsersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchUsersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchUsersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUsersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchUsersListOK creates a SearchUsersListOK with default headers values
func NewSearchUsersListOK() *SearchUsersListOK {
	return &SearchUsersListOK{}
}

/* SearchUsersListOK describes a response with status code 200, with default header values.

Success
*/
type SearchUsersListOK struct {
	Payload *models.UsersSearchList
}

func (o *SearchUsersListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/users][%d] searchUsersListOK  %+v", 200, o.Payload)
}
func (o *SearchUsersListOK) GetPayload() *models.UsersSearchList {
	return o.Payload
}

func (o *SearchUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersListBadRequest creates a SearchUsersListBadRequest with default headers values
func NewSearchUsersListBadRequest() *SearchUsersListBadRequest {
	return &SearchUsersListBadRequest{}
}

/* SearchUsersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchUsersListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SearchUsersListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/users][%d] searchUsersListBadRequest  %+v", 400, o.Payload)
}
func (o *SearchUsersListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchUsersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersListUnauthorized creates a SearchUsersListUnauthorized with default headers values
func NewSearchUsersListUnauthorized() *SearchUsersListUnauthorized {
	return &SearchUsersListUnauthorized{}
}

/* SearchUsersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchUsersListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SearchUsersListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/users][%d] searchUsersListUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchUsersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchUsersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersListForbidden creates a SearchUsersListForbidden with default headers values
func NewSearchUsersListForbidden() *SearchUsersListForbidden {
	return &SearchUsersListForbidden{}
}

/* SearchUsersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchUsersListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SearchUsersListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/users][%d] searchUsersListForbidden  %+v", 403, o.Payload)
}
func (o *SearchUsersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchUsersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersListNotFound creates a SearchUsersListNotFound with default headers values
func NewSearchUsersListNotFound() *SearchUsersListNotFound {
	return &SearchUsersListNotFound{}
}

/* SearchUsersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchUsersListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SearchUsersListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/users][%d] searchUsersListNotFound  %+v", 404, o.Payload)
}
func (o *SearchUsersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchUsersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersListInternalServerError creates a SearchUsersListInternalServerError with default headers values
func NewSearchUsersListInternalServerError() *SearchUsersListInternalServerError {
	return &SearchUsersListInternalServerError{}
}

/* SearchUsersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchUsersListInternalServerError struct {
}

func (o *SearchUsersListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/users][%d] searchUsersListInternalServerError ", 500)
}

func (o *SearchUsersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
