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

// SearchProjectsListReader is a Reader for the SearchProjectsList structure.
type SearchProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchProjectsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchProjectsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchProjectsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchProjectsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchProjectsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchProjectsListOK creates a SearchProjectsListOK with default headers values
func NewSearchProjectsListOK() *SearchProjectsListOK {
	return &SearchProjectsListOK{}
}

/*
SearchProjectsListOK describes a response with status code 200, with default header values.

Success
*/
type SearchProjectsListOK struct {
	Payload *models.ProjectsSearchList
}

// IsSuccess returns true when this search projects list o k response has a 2xx status code
func (o *SearchProjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search projects list o k response has a 3xx status code
func (o *SearchProjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search projects list o k response has a 4xx status code
func (o *SearchProjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search projects list o k response has a 5xx status code
func (o *SearchProjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search projects list o k response a status code equal to that given
func (o *SearchProjectsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchProjectsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListOK  %+v", 200, o.Payload)
}

func (o *SearchProjectsListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListOK  %+v", 200, o.Payload)
}

func (o *SearchProjectsListOK) GetPayload() *models.ProjectsSearchList {
	return o.Payload
}

func (o *SearchProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectsSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProjectsListBadRequest creates a SearchProjectsListBadRequest with default headers values
func NewSearchProjectsListBadRequest() *SearchProjectsListBadRequest {
	return &SearchProjectsListBadRequest{}
}

/*
SearchProjectsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchProjectsListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this search projects list bad request response has a 2xx status code
func (o *SearchProjectsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search projects list bad request response has a 3xx status code
func (o *SearchProjectsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search projects list bad request response has a 4xx status code
func (o *SearchProjectsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search projects list bad request response has a 5xx status code
func (o *SearchProjectsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search projects list bad request response a status code equal to that given
func (o *SearchProjectsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchProjectsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchProjectsListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchProjectsListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *SearchProjectsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProjectsListUnauthorized creates a SearchProjectsListUnauthorized with default headers values
func NewSearchProjectsListUnauthorized() *SearchProjectsListUnauthorized {
	return &SearchProjectsListUnauthorized{}
}

/*
SearchProjectsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchProjectsListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search projects list unauthorized response has a 2xx status code
func (o *SearchProjectsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search projects list unauthorized response has a 3xx status code
func (o *SearchProjectsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search projects list unauthorized response has a 4xx status code
func (o *SearchProjectsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search projects list unauthorized response has a 5xx status code
func (o *SearchProjectsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search projects list unauthorized response a status code equal to that given
func (o *SearchProjectsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchProjectsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchProjectsListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchProjectsListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchProjectsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProjectsListForbidden creates a SearchProjectsListForbidden with default headers values
func NewSearchProjectsListForbidden() *SearchProjectsListForbidden {
	return &SearchProjectsListForbidden{}
}

/*
SearchProjectsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchProjectsListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search projects list forbidden response has a 2xx status code
func (o *SearchProjectsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search projects list forbidden response has a 3xx status code
func (o *SearchProjectsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search projects list forbidden response has a 4xx status code
func (o *SearchProjectsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search projects list forbidden response has a 5xx status code
func (o *SearchProjectsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search projects list forbidden response a status code equal to that given
func (o *SearchProjectsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchProjectsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListForbidden  %+v", 403, o.Payload)
}

func (o *SearchProjectsListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListForbidden  %+v", 403, o.Payload)
}

func (o *SearchProjectsListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchProjectsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProjectsListNotFound creates a SearchProjectsListNotFound with default headers values
func NewSearchProjectsListNotFound() *SearchProjectsListNotFound {
	return &SearchProjectsListNotFound{}
}

/*
SearchProjectsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchProjectsListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search projects list not found response has a 2xx status code
func (o *SearchProjectsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search projects list not found response has a 3xx status code
func (o *SearchProjectsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search projects list not found response has a 4xx status code
func (o *SearchProjectsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search projects list not found response has a 5xx status code
func (o *SearchProjectsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search projects list not found response a status code equal to that given
func (o *SearchProjectsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchProjectsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListNotFound  %+v", 404, o.Payload)
}

func (o *SearchProjectsListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListNotFound  %+v", 404, o.Payload)
}

func (o *SearchProjectsListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchProjectsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchProjectsListInternalServerError creates a SearchProjectsListInternalServerError with default headers values
func NewSearchProjectsListInternalServerError() *SearchProjectsListInternalServerError {
	return &SearchProjectsListInternalServerError{}
}

/*
SearchProjectsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchProjectsListInternalServerError struct {
}

// IsSuccess returns true when this search projects list internal server error response has a 2xx status code
func (o *SearchProjectsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search projects list internal server error response has a 3xx status code
func (o *SearchProjectsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search projects list internal server error response has a 4xx status code
func (o *SearchProjectsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search projects list internal server error response has a 5xx status code
func (o *SearchProjectsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search projects list internal server error response a status code equal to that given
func (o *SearchProjectsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchProjectsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListInternalServerError ", 500)
}

func (o *SearchProjectsListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/projects][%d] searchProjectsListInternalServerError ", 500)
}

func (o *SearchProjectsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
