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

// SearchConfigMapListReader is a Reader for the SearchConfigMapList structure.
type SearchConfigMapListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchConfigMapListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchConfigMapListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchConfigMapListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchConfigMapListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchConfigMapListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchConfigMapListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchConfigMapListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchConfigMapListOK creates a SearchConfigMapListOK with default headers values
func NewSearchConfigMapListOK() *SearchConfigMapListOK {
	return &SearchConfigMapListOK{}
}

/*
SearchConfigMapListOK describes a response with status code 200, with default header values.

Success
*/
type SearchConfigMapListOK struct {
	Payload *models.ConfigMapSearchList
}

// IsSuccess returns true when this search config map list o k response has a 2xx status code
func (o *SearchConfigMapListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search config map list o k response has a 3xx status code
func (o *SearchConfigMapListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search config map list o k response has a 4xx status code
func (o *SearchConfigMapListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search config map list o k response has a 5xx status code
func (o *SearchConfigMapListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search config map list o k response a status code equal to that given
func (o *SearchConfigMapListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchConfigMapListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListOK  %+v", 200, o.Payload)
}

func (o *SearchConfigMapListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListOK  %+v", 200, o.Payload)
}

func (o *SearchConfigMapListOK) GetPayload() *models.ConfigMapSearchList {
	return o.Payload
}

func (o *SearchConfigMapListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigMapSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchConfigMapListBadRequest creates a SearchConfigMapListBadRequest with default headers values
func NewSearchConfigMapListBadRequest() *SearchConfigMapListBadRequest {
	return &SearchConfigMapListBadRequest{}
}

/*
SearchConfigMapListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchConfigMapListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search config map list bad request response has a 2xx status code
func (o *SearchConfigMapListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search config map list bad request response has a 3xx status code
func (o *SearchConfigMapListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search config map list bad request response has a 4xx status code
func (o *SearchConfigMapListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search config map list bad request response has a 5xx status code
func (o *SearchConfigMapListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search config map list bad request response a status code equal to that given
func (o *SearchConfigMapListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchConfigMapListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchConfigMapListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchConfigMapListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchConfigMapListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchConfigMapListUnauthorized creates a SearchConfigMapListUnauthorized with default headers values
func NewSearchConfigMapListUnauthorized() *SearchConfigMapListUnauthorized {
	return &SearchConfigMapListUnauthorized{}
}

/*
SearchConfigMapListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchConfigMapListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search config map list unauthorized response has a 2xx status code
func (o *SearchConfigMapListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search config map list unauthorized response has a 3xx status code
func (o *SearchConfigMapListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search config map list unauthorized response has a 4xx status code
func (o *SearchConfigMapListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search config map list unauthorized response has a 5xx status code
func (o *SearchConfigMapListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search config map list unauthorized response a status code equal to that given
func (o *SearchConfigMapListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchConfigMapListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchConfigMapListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchConfigMapListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchConfigMapListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchConfigMapListForbidden creates a SearchConfigMapListForbidden with default headers values
func NewSearchConfigMapListForbidden() *SearchConfigMapListForbidden {
	return &SearchConfigMapListForbidden{}
}

/*
SearchConfigMapListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchConfigMapListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search config map list forbidden response has a 2xx status code
func (o *SearchConfigMapListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search config map list forbidden response has a 3xx status code
func (o *SearchConfigMapListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search config map list forbidden response has a 4xx status code
func (o *SearchConfigMapListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search config map list forbidden response has a 5xx status code
func (o *SearchConfigMapListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search config map list forbidden response a status code equal to that given
func (o *SearchConfigMapListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchConfigMapListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListForbidden  %+v", 403, o.Payload)
}

func (o *SearchConfigMapListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListForbidden  %+v", 403, o.Payload)
}

func (o *SearchConfigMapListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchConfigMapListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchConfigMapListNotFound creates a SearchConfigMapListNotFound with default headers values
func NewSearchConfigMapListNotFound() *SearchConfigMapListNotFound {
	return &SearchConfigMapListNotFound{}
}

/*
SearchConfigMapListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchConfigMapListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search config map list not found response has a 2xx status code
func (o *SearchConfigMapListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search config map list not found response has a 3xx status code
func (o *SearchConfigMapListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search config map list not found response has a 4xx status code
func (o *SearchConfigMapListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search config map list not found response has a 5xx status code
func (o *SearchConfigMapListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search config map list not found response a status code equal to that given
func (o *SearchConfigMapListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchConfigMapListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListNotFound  %+v", 404, o.Payload)
}

func (o *SearchConfigMapListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListNotFound  %+v", 404, o.Payload)
}

func (o *SearchConfigMapListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchConfigMapListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchConfigMapListInternalServerError creates a SearchConfigMapListInternalServerError with default headers values
func NewSearchConfigMapListInternalServerError() *SearchConfigMapListInternalServerError {
	return &SearchConfigMapListInternalServerError{}
}

/*
SearchConfigMapListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchConfigMapListInternalServerError struct {
}

// IsSuccess returns true when this search config map list internal server error response has a 2xx status code
func (o *SearchConfigMapListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search config map list internal server error response has a 3xx status code
func (o *SearchConfigMapListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search config map list internal server error response has a 4xx status code
func (o *SearchConfigMapListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search config map list internal server error response has a 5xx status code
func (o *SearchConfigMapListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search config map list internal server error response a status code equal to that given
func (o *SearchConfigMapListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchConfigMapListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListInternalServerError ", 500)
}

func (o *SearchConfigMapListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/config-maps][%d] searchConfigMapListInternalServerError ", 500)
}

func (o *SearchConfigMapListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
