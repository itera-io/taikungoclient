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

// SearchStsListReader is a Reader for the SearchStsList structure.
type SearchStsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchStsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchStsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchStsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchStsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchStsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchStsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchStsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchStsListOK creates a SearchStsListOK with default headers values
func NewSearchStsListOK() *SearchStsListOK {
	return &SearchStsListOK{}
}

/*
SearchStsListOK describes a response with status code 200, with default header values.

Success
*/
type SearchStsListOK struct {
	Payload *models.StsSearchList
}

// IsSuccess returns true when this search sts list o k response has a 2xx status code
func (o *SearchStsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search sts list o k response has a 3xx status code
func (o *SearchStsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search sts list o k response has a 4xx status code
func (o *SearchStsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search sts list o k response has a 5xx status code
func (o *SearchStsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search sts list o k response a status code equal to that given
func (o *SearchStsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchStsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListOK  %+v", 200, o.Payload)
}

func (o *SearchStsListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListOK  %+v", 200, o.Payload)
}

func (o *SearchStsListOK) GetPayload() *models.StsSearchList {
	return o.Payload
}

func (o *SearchStsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StsSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStsListBadRequest creates a SearchStsListBadRequest with default headers values
func NewSearchStsListBadRequest() *SearchStsListBadRequest {
	return &SearchStsListBadRequest{}
}

/*
SearchStsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchStsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this search sts list bad request response has a 2xx status code
func (o *SearchStsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search sts list bad request response has a 3xx status code
func (o *SearchStsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search sts list bad request response has a 4xx status code
func (o *SearchStsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search sts list bad request response has a 5xx status code
func (o *SearchStsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search sts list bad request response a status code equal to that given
func (o *SearchStsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchStsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchStsListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchStsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchStsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStsListUnauthorized creates a SearchStsListUnauthorized with default headers values
func NewSearchStsListUnauthorized() *SearchStsListUnauthorized {
	return &SearchStsListUnauthorized{}
}

/*
SearchStsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchStsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search sts list unauthorized response has a 2xx status code
func (o *SearchStsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search sts list unauthorized response has a 3xx status code
func (o *SearchStsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search sts list unauthorized response has a 4xx status code
func (o *SearchStsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search sts list unauthorized response has a 5xx status code
func (o *SearchStsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search sts list unauthorized response a status code equal to that given
func (o *SearchStsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchStsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchStsListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchStsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchStsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStsListForbidden creates a SearchStsListForbidden with default headers values
func NewSearchStsListForbidden() *SearchStsListForbidden {
	return &SearchStsListForbidden{}
}

/*
SearchStsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchStsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search sts list forbidden response has a 2xx status code
func (o *SearchStsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search sts list forbidden response has a 3xx status code
func (o *SearchStsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search sts list forbidden response has a 4xx status code
func (o *SearchStsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search sts list forbidden response has a 5xx status code
func (o *SearchStsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search sts list forbidden response a status code equal to that given
func (o *SearchStsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchStsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListForbidden  %+v", 403, o.Payload)
}

func (o *SearchStsListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListForbidden  %+v", 403, o.Payload)
}

func (o *SearchStsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchStsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStsListNotFound creates a SearchStsListNotFound with default headers values
func NewSearchStsListNotFound() *SearchStsListNotFound {
	return &SearchStsListNotFound{}
}

/*
SearchStsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchStsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search sts list not found response has a 2xx status code
func (o *SearchStsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search sts list not found response has a 3xx status code
func (o *SearchStsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search sts list not found response has a 4xx status code
func (o *SearchStsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search sts list not found response has a 5xx status code
func (o *SearchStsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search sts list not found response a status code equal to that given
func (o *SearchStsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchStsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListNotFound  %+v", 404, o.Payload)
}

func (o *SearchStsListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListNotFound  %+v", 404, o.Payload)
}

func (o *SearchStsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchStsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchStsListInternalServerError creates a SearchStsListInternalServerError with default headers values
func NewSearchStsListInternalServerError() *SearchStsListInternalServerError {
	return &SearchStsListInternalServerError{}
}

/*
SearchStsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchStsListInternalServerError struct {
}

// IsSuccess returns true when this search sts list internal server error response has a 2xx status code
func (o *SearchStsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search sts list internal server error response has a 3xx status code
func (o *SearchStsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search sts list internal server error response has a 4xx status code
func (o *SearchStsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search sts list internal server error response has a 5xx status code
func (o *SearchStsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search sts list internal server error response a status code equal to that given
func (o *SearchStsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchStsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListInternalServerError ", 500)
}

func (o *SearchStsListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/sts][%d] searchStsListInternalServerError ", 500)
}

func (o *SearchStsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
