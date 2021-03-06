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

// SearchServersListReader is a Reader for the SearchServersList structure.
type SearchServersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchServersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchServersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchServersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchServersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchServersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchServersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchServersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchServersListOK creates a SearchServersListOK with default headers values
func NewSearchServersListOK() *SearchServersListOK {
	return &SearchServersListOK{}
}

/* SearchServersListOK describes a response with status code 200, with default header values.

Success
*/
type SearchServersListOK struct {
	Payload *models.ServersSearchList
}

func (o *SearchServersListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/servers][%d] searchServersListOK  %+v", 200, o.Payload)
}
func (o *SearchServersListOK) GetPayload() *models.ServersSearchList {
	return o.Payload
}

func (o *SearchServersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServersSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServersListBadRequest creates a SearchServersListBadRequest with default headers values
func NewSearchServersListBadRequest() *SearchServersListBadRequest {
	return &SearchServersListBadRequest{}
}

/* SearchServersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchServersListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SearchServersListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/servers][%d] searchServersListBadRequest  %+v", 400, o.Payload)
}
func (o *SearchServersListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchServersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServersListUnauthorized creates a SearchServersListUnauthorized with default headers values
func NewSearchServersListUnauthorized() *SearchServersListUnauthorized {
	return &SearchServersListUnauthorized{}
}

/* SearchServersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchServersListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SearchServersListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/servers][%d] searchServersListUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchServersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchServersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServersListForbidden creates a SearchServersListForbidden with default headers values
func NewSearchServersListForbidden() *SearchServersListForbidden {
	return &SearchServersListForbidden{}
}

/* SearchServersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchServersListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SearchServersListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/servers][%d] searchServersListForbidden  %+v", 403, o.Payload)
}
func (o *SearchServersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchServersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServersListNotFound creates a SearchServersListNotFound with default headers values
func NewSearchServersListNotFound() *SearchServersListNotFound {
	return &SearchServersListNotFound{}
}

/* SearchServersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchServersListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SearchServersListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/servers][%d] searchServersListNotFound  %+v", 404, o.Payload)
}
func (o *SearchServersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchServersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServersListInternalServerError creates a SearchServersListInternalServerError with default headers values
func NewSearchServersListInternalServerError() *SearchServersListInternalServerError {
	return &SearchServersListInternalServerError{}
}

/* SearchServersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchServersListInternalServerError struct {
}

func (o *SearchServersListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/servers][%d] searchServersListInternalServerError ", 500)
}

func (o *SearchServersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
