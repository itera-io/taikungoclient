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

// SearchServiceListReader is a Reader for the SearchServiceList structure.
type SearchServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchServiceListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchServiceListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchServiceListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchServiceListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchServiceListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchServiceListOK creates a SearchServiceListOK with default headers values
func NewSearchServiceListOK() *SearchServiceListOK {
	return &SearchServiceListOK{}
}

/*
SearchServiceListOK describes a response with status code 200, with default header values.

Success
*/
type SearchServiceListOK struct {
	Payload *models.ServiceSearchList
}

// IsSuccess returns true when this search service list o k response has a 2xx status code
func (o *SearchServiceListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search service list o k response has a 3xx status code
func (o *SearchServiceListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service list o k response has a 4xx status code
func (o *SearchServiceListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search service list o k response has a 5xx status code
func (o *SearchServiceListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search service list o k response a status code equal to that given
func (o *SearchServiceListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchServiceListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListOK  %+v", 200, o.Payload)
}

func (o *SearchServiceListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListOK  %+v", 200, o.Payload)
}

func (o *SearchServiceListOK) GetPayload() *models.ServiceSearchList {
	return o.Payload
}

func (o *SearchServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServiceListBadRequest creates a SearchServiceListBadRequest with default headers values
func NewSearchServiceListBadRequest() *SearchServiceListBadRequest {
	return &SearchServiceListBadRequest{}
}

/*
SearchServiceListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchServiceListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this search service list bad request response has a 2xx status code
func (o *SearchServiceListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search service list bad request response has a 3xx status code
func (o *SearchServiceListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service list bad request response has a 4xx status code
func (o *SearchServiceListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search service list bad request response has a 5xx status code
func (o *SearchServiceListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search service list bad request response a status code equal to that given
func (o *SearchServiceListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchServiceListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchServiceListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchServiceListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *SearchServiceListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServiceListUnauthorized creates a SearchServiceListUnauthorized with default headers values
func NewSearchServiceListUnauthorized() *SearchServiceListUnauthorized {
	return &SearchServiceListUnauthorized{}
}

/*
SearchServiceListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchServiceListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search service list unauthorized response has a 2xx status code
func (o *SearchServiceListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search service list unauthorized response has a 3xx status code
func (o *SearchServiceListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service list unauthorized response has a 4xx status code
func (o *SearchServiceListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search service list unauthorized response has a 5xx status code
func (o *SearchServiceListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search service list unauthorized response a status code equal to that given
func (o *SearchServiceListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchServiceListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchServiceListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchServiceListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchServiceListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServiceListForbidden creates a SearchServiceListForbidden with default headers values
func NewSearchServiceListForbidden() *SearchServiceListForbidden {
	return &SearchServiceListForbidden{}
}

/*
SearchServiceListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchServiceListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search service list forbidden response has a 2xx status code
func (o *SearchServiceListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search service list forbidden response has a 3xx status code
func (o *SearchServiceListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service list forbidden response has a 4xx status code
func (o *SearchServiceListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search service list forbidden response has a 5xx status code
func (o *SearchServiceListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search service list forbidden response a status code equal to that given
func (o *SearchServiceListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchServiceListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListForbidden  %+v", 403, o.Payload)
}

func (o *SearchServiceListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListForbidden  %+v", 403, o.Payload)
}

func (o *SearchServiceListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchServiceListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServiceListNotFound creates a SearchServiceListNotFound with default headers values
func NewSearchServiceListNotFound() *SearchServiceListNotFound {
	return &SearchServiceListNotFound{}
}

/*
SearchServiceListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchServiceListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search service list not found response has a 2xx status code
func (o *SearchServiceListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search service list not found response has a 3xx status code
func (o *SearchServiceListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service list not found response has a 4xx status code
func (o *SearchServiceListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search service list not found response has a 5xx status code
func (o *SearchServiceListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search service list not found response a status code equal to that given
func (o *SearchServiceListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchServiceListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListNotFound  %+v", 404, o.Payload)
}

func (o *SearchServiceListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListNotFound  %+v", 404, o.Payload)
}

func (o *SearchServiceListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchServiceListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchServiceListInternalServerError creates a SearchServiceListInternalServerError with default headers values
func NewSearchServiceListInternalServerError() *SearchServiceListInternalServerError {
	return &SearchServiceListInternalServerError{}
}

/*
SearchServiceListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchServiceListInternalServerError struct {
}

// IsSuccess returns true when this search service list internal server error response has a 2xx status code
func (o *SearchServiceListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search service list internal server error response has a 3xx status code
func (o *SearchServiceListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service list internal server error response has a 4xx status code
func (o *SearchServiceListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search service list internal server error response has a 5xx status code
func (o *SearchServiceListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search service list internal server error response a status code equal to that given
func (o *SearchServiceListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchServiceListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListInternalServerError ", 500)
}

func (o *SearchServiceListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/services][%d] searchServiceListInternalServerError ", 500)
}

func (o *SearchServiceListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
