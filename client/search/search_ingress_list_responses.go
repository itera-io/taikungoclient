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

// SearchIngressListReader is a Reader for the SearchIngressList structure.
type SearchIngressListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchIngressListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchIngressListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchIngressListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchIngressListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchIngressListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchIngressListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchIngressListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchIngressListOK creates a SearchIngressListOK with default headers values
func NewSearchIngressListOK() *SearchIngressListOK {
	return &SearchIngressListOK{}
}

/*
SearchIngressListOK describes a response with status code 200, with default header values.

Success
*/
type SearchIngressListOK struct {
	Payload *models.IngressSearchList
}

// IsSuccess returns true when this search ingress list o k response has a 2xx status code
func (o *SearchIngressListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search ingress list o k response has a 3xx status code
func (o *SearchIngressListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ingress list o k response has a 4xx status code
func (o *SearchIngressListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search ingress list o k response has a 5xx status code
func (o *SearchIngressListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search ingress list o k response a status code equal to that given
func (o *SearchIngressListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchIngressListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListOK  %+v", 200, o.Payload)
}

func (o *SearchIngressListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListOK  %+v", 200, o.Payload)
}

func (o *SearchIngressListOK) GetPayload() *models.IngressSearchList {
	return o.Payload
}

func (o *SearchIngressListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IngressSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIngressListBadRequest creates a SearchIngressListBadRequest with default headers values
func NewSearchIngressListBadRequest() *SearchIngressListBadRequest {
	return &SearchIngressListBadRequest{}
}

/*
SearchIngressListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchIngressListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this search ingress list bad request response has a 2xx status code
func (o *SearchIngressListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ingress list bad request response has a 3xx status code
func (o *SearchIngressListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ingress list bad request response has a 4xx status code
func (o *SearchIngressListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ingress list bad request response has a 5xx status code
func (o *SearchIngressListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search ingress list bad request response a status code equal to that given
func (o *SearchIngressListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchIngressListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchIngressListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchIngressListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchIngressListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIngressListUnauthorized creates a SearchIngressListUnauthorized with default headers values
func NewSearchIngressListUnauthorized() *SearchIngressListUnauthorized {
	return &SearchIngressListUnauthorized{}
}

/*
SearchIngressListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchIngressListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search ingress list unauthorized response has a 2xx status code
func (o *SearchIngressListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ingress list unauthorized response has a 3xx status code
func (o *SearchIngressListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ingress list unauthorized response has a 4xx status code
func (o *SearchIngressListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ingress list unauthorized response has a 5xx status code
func (o *SearchIngressListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search ingress list unauthorized response a status code equal to that given
func (o *SearchIngressListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchIngressListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchIngressListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchIngressListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchIngressListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIngressListForbidden creates a SearchIngressListForbidden with default headers values
func NewSearchIngressListForbidden() *SearchIngressListForbidden {
	return &SearchIngressListForbidden{}
}

/*
SearchIngressListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchIngressListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search ingress list forbidden response has a 2xx status code
func (o *SearchIngressListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ingress list forbidden response has a 3xx status code
func (o *SearchIngressListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ingress list forbidden response has a 4xx status code
func (o *SearchIngressListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ingress list forbidden response has a 5xx status code
func (o *SearchIngressListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search ingress list forbidden response a status code equal to that given
func (o *SearchIngressListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchIngressListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListForbidden  %+v", 403, o.Payload)
}

func (o *SearchIngressListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListForbidden  %+v", 403, o.Payload)
}

func (o *SearchIngressListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchIngressListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIngressListNotFound creates a SearchIngressListNotFound with default headers values
func NewSearchIngressListNotFound() *SearchIngressListNotFound {
	return &SearchIngressListNotFound{}
}

/*
SearchIngressListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchIngressListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this search ingress list not found response has a 2xx status code
func (o *SearchIngressListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ingress list not found response has a 3xx status code
func (o *SearchIngressListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ingress list not found response has a 4xx status code
func (o *SearchIngressListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ingress list not found response has a 5xx status code
func (o *SearchIngressListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search ingress list not found response a status code equal to that given
func (o *SearchIngressListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchIngressListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListNotFound  %+v", 404, o.Payload)
}

func (o *SearchIngressListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListNotFound  %+v", 404, o.Payload)
}

func (o *SearchIngressListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *SearchIngressListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchIngressListInternalServerError creates a SearchIngressListInternalServerError with default headers values
func NewSearchIngressListInternalServerError() *SearchIngressListInternalServerError {
	return &SearchIngressListInternalServerError{}
}

/*
SearchIngressListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchIngressListInternalServerError struct {
}

// IsSuccess returns true when this search ingress list internal server error response has a 2xx status code
func (o *SearchIngressListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ingress list internal server error response has a 3xx status code
func (o *SearchIngressListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ingress list internal server error response has a 4xx status code
func (o *SearchIngressListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search ingress list internal server error response has a 5xx status code
func (o *SearchIngressListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search ingress list internal server error response a status code equal to that given
func (o *SearchIngressListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchIngressListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListInternalServerError ", 500)
}

func (o *SearchIngressListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/ingress][%d] searchIngressListInternalServerError ", 500)
}

func (o *SearchIngressListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
