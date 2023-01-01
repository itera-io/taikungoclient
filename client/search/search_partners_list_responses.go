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

// SearchPartnersListReader is a Reader for the SearchPartnersList structure.
type SearchPartnersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPartnersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPartnersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchPartnersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchPartnersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchPartnersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchPartnersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchPartnersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchPartnersListOK creates a SearchPartnersListOK with default headers values
func NewSearchPartnersListOK() *SearchPartnersListOK {
	return &SearchPartnersListOK{}
}

/*
SearchPartnersListOK describes a response with status code 200, with default header values.

Success
*/
type SearchPartnersListOK struct {
	Payload *models.PartnersSearchList
}

// IsSuccess returns true when this search partners list o k response has a 2xx status code
func (o *SearchPartnersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search partners list o k response has a 3xx status code
func (o *SearchPartnersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search partners list o k response has a 4xx status code
func (o *SearchPartnersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search partners list o k response has a 5xx status code
func (o *SearchPartnersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search partners list o k response a status code equal to that given
func (o *SearchPartnersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchPartnersListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListOK  %+v", 200, o.Payload)
}

func (o *SearchPartnersListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListOK  %+v", 200, o.Payload)
}

func (o *SearchPartnersListOK) GetPayload() *models.PartnersSearchList {
	return o.Payload
}

func (o *SearchPartnersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartnersSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPartnersListBadRequest creates a SearchPartnersListBadRequest with default headers values
func NewSearchPartnersListBadRequest() *SearchPartnersListBadRequest {
	return &SearchPartnersListBadRequest{}
}

/*
SearchPartnersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchPartnersListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search partners list bad request response has a 2xx status code
func (o *SearchPartnersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search partners list bad request response has a 3xx status code
func (o *SearchPartnersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search partners list bad request response has a 4xx status code
func (o *SearchPartnersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search partners list bad request response has a 5xx status code
func (o *SearchPartnersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search partners list bad request response a status code equal to that given
func (o *SearchPartnersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchPartnersListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPartnersListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPartnersListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPartnersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPartnersListUnauthorized creates a SearchPartnersListUnauthorized with default headers values
func NewSearchPartnersListUnauthorized() *SearchPartnersListUnauthorized {
	return &SearchPartnersListUnauthorized{}
}

/*
SearchPartnersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchPartnersListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search partners list unauthorized response has a 2xx status code
func (o *SearchPartnersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search partners list unauthorized response has a 3xx status code
func (o *SearchPartnersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search partners list unauthorized response has a 4xx status code
func (o *SearchPartnersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search partners list unauthorized response has a 5xx status code
func (o *SearchPartnersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search partners list unauthorized response a status code equal to that given
func (o *SearchPartnersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchPartnersListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPartnersListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPartnersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPartnersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPartnersListForbidden creates a SearchPartnersListForbidden with default headers values
func NewSearchPartnersListForbidden() *SearchPartnersListForbidden {
	return &SearchPartnersListForbidden{}
}

/*
SearchPartnersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchPartnersListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search partners list forbidden response has a 2xx status code
func (o *SearchPartnersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search partners list forbidden response has a 3xx status code
func (o *SearchPartnersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search partners list forbidden response has a 4xx status code
func (o *SearchPartnersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search partners list forbidden response has a 5xx status code
func (o *SearchPartnersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search partners list forbidden response a status code equal to that given
func (o *SearchPartnersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchPartnersListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPartnersListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPartnersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPartnersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPartnersListNotFound creates a SearchPartnersListNotFound with default headers values
func NewSearchPartnersListNotFound() *SearchPartnersListNotFound {
	return &SearchPartnersListNotFound{}
}

/*
SearchPartnersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchPartnersListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search partners list not found response has a 2xx status code
func (o *SearchPartnersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search partners list not found response has a 3xx status code
func (o *SearchPartnersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search partners list not found response has a 4xx status code
func (o *SearchPartnersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search partners list not found response has a 5xx status code
func (o *SearchPartnersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search partners list not found response a status code equal to that given
func (o *SearchPartnersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchPartnersListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPartnersListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPartnersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPartnersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPartnersListInternalServerError creates a SearchPartnersListInternalServerError with default headers values
func NewSearchPartnersListInternalServerError() *SearchPartnersListInternalServerError {
	return &SearchPartnersListInternalServerError{}
}

/*
SearchPartnersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchPartnersListInternalServerError struct {
}

// IsSuccess returns true when this search partners list internal server error response has a 2xx status code
func (o *SearchPartnersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search partners list internal server error response has a 3xx status code
func (o *SearchPartnersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search partners list internal server error response has a 4xx status code
func (o *SearchPartnersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search partners list internal server error response has a 5xx status code
func (o *SearchPartnersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search partners list internal server error response a status code equal to that given
func (o *SearchPartnersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchPartnersListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListInternalServerError ", 500)
}

func (o *SearchPartnersListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/partners][%d] searchPartnersListInternalServerError ", 500)
}

func (o *SearchPartnersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
