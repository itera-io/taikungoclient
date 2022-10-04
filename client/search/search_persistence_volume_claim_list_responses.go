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

// SearchPersistenceVolumeClaimListReader is a Reader for the SearchPersistenceVolumeClaimList structure.
type SearchPersistenceVolumeClaimListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPersistenceVolumeClaimListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPersistenceVolumeClaimListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchPersistenceVolumeClaimListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchPersistenceVolumeClaimListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchPersistenceVolumeClaimListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchPersistenceVolumeClaimListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchPersistenceVolumeClaimListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchPersistenceVolumeClaimListOK creates a SearchPersistenceVolumeClaimListOK with default headers values
func NewSearchPersistenceVolumeClaimListOK() *SearchPersistenceVolumeClaimListOK {
	return &SearchPersistenceVolumeClaimListOK{}
}

/*
SearchPersistenceVolumeClaimListOK describes a response with status code 200, with default header values.

Success
*/
type SearchPersistenceVolumeClaimListOK struct {
	Payload *models.PvcSearchList
}

// IsSuccess returns true when this search persistence volume claim list o k response has a 2xx status code
func (o *SearchPersistenceVolumeClaimListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search persistence volume claim list o k response has a 3xx status code
func (o *SearchPersistenceVolumeClaimListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search persistence volume claim list o k response has a 4xx status code
func (o *SearchPersistenceVolumeClaimListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search persistence volume claim list o k response has a 5xx status code
func (o *SearchPersistenceVolumeClaimListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search persistence volume claim list o k response a status code equal to that given
func (o *SearchPersistenceVolumeClaimListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchPersistenceVolumeClaimListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListOK  %+v", 200, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListOK  %+v", 200, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListOK) GetPayload() *models.PvcSearchList {
	return o.Payload
}

func (o *SearchPersistenceVolumeClaimListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PvcSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPersistenceVolumeClaimListBadRequest creates a SearchPersistenceVolumeClaimListBadRequest with default headers values
func NewSearchPersistenceVolumeClaimListBadRequest() *SearchPersistenceVolumeClaimListBadRequest {
	return &SearchPersistenceVolumeClaimListBadRequest{}
}

/*
SearchPersistenceVolumeClaimListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchPersistenceVolumeClaimListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search persistence volume claim list bad request response has a 2xx status code
func (o *SearchPersistenceVolumeClaimListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search persistence volume claim list bad request response has a 3xx status code
func (o *SearchPersistenceVolumeClaimListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search persistence volume claim list bad request response has a 4xx status code
func (o *SearchPersistenceVolumeClaimListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search persistence volume claim list bad request response has a 5xx status code
func (o *SearchPersistenceVolumeClaimListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search persistence volume claim list bad request response a status code equal to that given
func (o *SearchPersistenceVolumeClaimListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchPersistenceVolumeClaimListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchPersistenceVolumeClaimListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPersistenceVolumeClaimListUnauthorized creates a SearchPersistenceVolumeClaimListUnauthorized with default headers values
func NewSearchPersistenceVolumeClaimListUnauthorized() *SearchPersistenceVolumeClaimListUnauthorized {
	return &SearchPersistenceVolumeClaimListUnauthorized{}
}

/*
SearchPersistenceVolumeClaimListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchPersistenceVolumeClaimListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search persistence volume claim list unauthorized response has a 2xx status code
func (o *SearchPersistenceVolumeClaimListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search persistence volume claim list unauthorized response has a 3xx status code
func (o *SearchPersistenceVolumeClaimListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search persistence volume claim list unauthorized response has a 4xx status code
func (o *SearchPersistenceVolumeClaimListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search persistence volume claim list unauthorized response has a 5xx status code
func (o *SearchPersistenceVolumeClaimListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search persistence volume claim list unauthorized response a status code equal to that given
func (o *SearchPersistenceVolumeClaimListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchPersistenceVolumeClaimListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchPersistenceVolumeClaimListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPersistenceVolumeClaimListForbidden creates a SearchPersistenceVolumeClaimListForbidden with default headers values
func NewSearchPersistenceVolumeClaimListForbidden() *SearchPersistenceVolumeClaimListForbidden {
	return &SearchPersistenceVolumeClaimListForbidden{}
}

/*
SearchPersistenceVolumeClaimListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchPersistenceVolumeClaimListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search persistence volume claim list forbidden response has a 2xx status code
func (o *SearchPersistenceVolumeClaimListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search persistence volume claim list forbidden response has a 3xx status code
func (o *SearchPersistenceVolumeClaimListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search persistence volume claim list forbidden response has a 4xx status code
func (o *SearchPersistenceVolumeClaimListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search persistence volume claim list forbidden response has a 5xx status code
func (o *SearchPersistenceVolumeClaimListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search persistence volume claim list forbidden response a status code equal to that given
func (o *SearchPersistenceVolumeClaimListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchPersistenceVolumeClaimListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchPersistenceVolumeClaimListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPersistenceVolumeClaimListNotFound creates a SearchPersistenceVolumeClaimListNotFound with default headers values
func NewSearchPersistenceVolumeClaimListNotFound() *SearchPersistenceVolumeClaimListNotFound {
	return &SearchPersistenceVolumeClaimListNotFound{}
}

/*
SearchPersistenceVolumeClaimListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchPersistenceVolumeClaimListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search persistence volume claim list not found response has a 2xx status code
func (o *SearchPersistenceVolumeClaimListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search persistence volume claim list not found response has a 3xx status code
func (o *SearchPersistenceVolumeClaimListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search persistence volume claim list not found response has a 4xx status code
func (o *SearchPersistenceVolumeClaimListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search persistence volume claim list not found response has a 5xx status code
func (o *SearchPersistenceVolumeClaimListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search persistence volume claim list not found response a status code equal to that given
func (o *SearchPersistenceVolumeClaimListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchPersistenceVolumeClaimListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPersistenceVolumeClaimListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchPersistenceVolumeClaimListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPersistenceVolumeClaimListInternalServerError creates a SearchPersistenceVolumeClaimListInternalServerError with default headers values
func NewSearchPersistenceVolumeClaimListInternalServerError() *SearchPersistenceVolumeClaimListInternalServerError {
	return &SearchPersistenceVolumeClaimListInternalServerError{}
}

/*
SearchPersistenceVolumeClaimListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchPersistenceVolumeClaimListInternalServerError struct {
}

// IsSuccess returns true when this search persistence volume claim list internal server error response has a 2xx status code
func (o *SearchPersistenceVolumeClaimListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search persistence volume claim list internal server error response has a 3xx status code
func (o *SearchPersistenceVolumeClaimListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search persistence volume claim list internal server error response has a 4xx status code
func (o *SearchPersistenceVolumeClaimListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search persistence volume claim list internal server error response has a 5xx status code
func (o *SearchPersistenceVolumeClaimListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search persistence volume claim list internal server error response a status code equal to that given
func (o *SearchPersistenceVolumeClaimListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchPersistenceVolumeClaimListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListInternalServerError ", 500)
}

func (o *SearchPersistenceVolumeClaimListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/pvcs][%d] searchPersistenceVolumeClaimListInternalServerError ", 500)
}

func (o *SearchPersistenceVolumeClaimListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
