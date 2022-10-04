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

// SearchAccessProfilesListReader is a Reader for the SearchAccessProfilesList structure.
type SearchAccessProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAccessProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAccessProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchAccessProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchAccessProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchAccessProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchAccessProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchAccessProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchAccessProfilesListOK creates a SearchAccessProfilesListOK with default headers values
func NewSearchAccessProfilesListOK() *SearchAccessProfilesListOK {
	return &SearchAccessProfilesListOK{}
}

/*
SearchAccessProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type SearchAccessProfilesListOK struct {
	Payload *models.AccessProfilesSearchList
}

// IsSuccess returns true when this search access profiles list o k response has a 2xx status code
func (o *SearchAccessProfilesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search access profiles list o k response has a 3xx status code
func (o *SearchAccessProfilesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search access profiles list o k response has a 4xx status code
func (o *SearchAccessProfilesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search access profiles list o k response has a 5xx status code
func (o *SearchAccessProfilesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search access profiles list o k response a status code equal to that given
func (o *SearchAccessProfilesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchAccessProfilesListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListOK  %+v", 200, o.Payload)
}

func (o *SearchAccessProfilesListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListOK  %+v", 200, o.Payload)
}

func (o *SearchAccessProfilesListOK) GetPayload() *models.AccessProfilesSearchList {
	return o.Payload
}

func (o *SearchAccessProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessProfilesSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAccessProfilesListBadRequest creates a SearchAccessProfilesListBadRequest with default headers values
func NewSearchAccessProfilesListBadRequest() *SearchAccessProfilesListBadRequest {
	return &SearchAccessProfilesListBadRequest{}
}

/*
SearchAccessProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchAccessProfilesListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this search access profiles list bad request response has a 2xx status code
func (o *SearchAccessProfilesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search access profiles list bad request response has a 3xx status code
func (o *SearchAccessProfilesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search access profiles list bad request response has a 4xx status code
func (o *SearchAccessProfilesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search access profiles list bad request response has a 5xx status code
func (o *SearchAccessProfilesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search access profiles list bad request response a status code equal to that given
func (o *SearchAccessProfilesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchAccessProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchAccessProfilesListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchAccessProfilesListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchAccessProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAccessProfilesListUnauthorized creates a SearchAccessProfilesListUnauthorized with default headers values
func NewSearchAccessProfilesListUnauthorized() *SearchAccessProfilesListUnauthorized {
	return &SearchAccessProfilesListUnauthorized{}
}

/*
SearchAccessProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchAccessProfilesListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search access profiles list unauthorized response has a 2xx status code
func (o *SearchAccessProfilesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search access profiles list unauthorized response has a 3xx status code
func (o *SearchAccessProfilesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search access profiles list unauthorized response has a 4xx status code
func (o *SearchAccessProfilesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search access profiles list unauthorized response has a 5xx status code
func (o *SearchAccessProfilesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search access profiles list unauthorized response a status code equal to that given
func (o *SearchAccessProfilesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchAccessProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchAccessProfilesListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchAccessProfilesListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchAccessProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAccessProfilesListForbidden creates a SearchAccessProfilesListForbidden with default headers values
func NewSearchAccessProfilesListForbidden() *SearchAccessProfilesListForbidden {
	return &SearchAccessProfilesListForbidden{}
}

/*
SearchAccessProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchAccessProfilesListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search access profiles list forbidden response has a 2xx status code
func (o *SearchAccessProfilesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search access profiles list forbidden response has a 3xx status code
func (o *SearchAccessProfilesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search access profiles list forbidden response has a 4xx status code
func (o *SearchAccessProfilesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search access profiles list forbidden response has a 5xx status code
func (o *SearchAccessProfilesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search access profiles list forbidden response a status code equal to that given
func (o *SearchAccessProfilesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchAccessProfilesListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchAccessProfilesListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchAccessProfilesListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchAccessProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAccessProfilesListNotFound creates a SearchAccessProfilesListNotFound with default headers values
func NewSearchAccessProfilesListNotFound() *SearchAccessProfilesListNotFound {
	return &SearchAccessProfilesListNotFound{}
}

/*
SearchAccessProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchAccessProfilesListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search access profiles list not found response has a 2xx status code
func (o *SearchAccessProfilesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search access profiles list not found response has a 3xx status code
func (o *SearchAccessProfilesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search access profiles list not found response has a 4xx status code
func (o *SearchAccessProfilesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search access profiles list not found response has a 5xx status code
func (o *SearchAccessProfilesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search access profiles list not found response a status code equal to that given
func (o *SearchAccessProfilesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchAccessProfilesListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchAccessProfilesListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchAccessProfilesListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchAccessProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAccessProfilesListInternalServerError creates a SearchAccessProfilesListInternalServerError with default headers values
func NewSearchAccessProfilesListInternalServerError() *SearchAccessProfilesListInternalServerError {
	return &SearchAccessProfilesListInternalServerError{}
}

/*
SearchAccessProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchAccessProfilesListInternalServerError struct {
}

// IsSuccess returns true when this search access profiles list internal server error response has a 2xx status code
func (o *SearchAccessProfilesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search access profiles list internal server error response has a 3xx status code
func (o *SearchAccessProfilesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search access profiles list internal server error response has a 4xx status code
func (o *SearchAccessProfilesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search access profiles list internal server error response has a 5xx status code
func (o *SearchAccessProfilesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search access profiles list internal server error response a status code equal to that given
func (o *SearchAccessProfilesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchAccessProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListInternalServerError ", 500)
}

func (o *SearchAccessProfilesListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/access-profiles][%d] searchAccessProfilesListInternalServerError ", 500)
}

func (o *SearchAccessProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
