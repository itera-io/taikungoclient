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

// SearchPrometheusRulesListReader is a Reader for the SearchPrometheusRulesList structure.
type SearchPrometheusRulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPrometheusRulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchPrometheusRulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchPrometheusRulesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchPrometheusRulesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchPrometheusRulesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchPrometheusRulesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchPrometheusRulesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchPrometheusRulesListOK creates a SearchPrometheusRulesListOK with default headers values
func NewSearchPrometheusRulesListOK() *SearchPrometheusRulesListOK {
	return &SearchPrometheusRulesListOK{}
}

/*
SearchPrometheusRulesListOK describes a response with status code 200, with default header values.

Success
*/
type SearchPrometheusRulesListOK struct {
	Payload *models.PrometheusRulesSearchList
}

// IsSuccess returns true when this search prometheus rules list o k response has a 2xx status code
func (o *SearchPrometheusRulesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search prometheus rules list o k response has a 3xx status code
func (o *SearchPrometheusRulesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search prometheus rules list o k response has a 4xx status code
func (o *SearchPrometheusRulesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search prometheus rules list o k response has a 5xx status code
func (o *SearchPrometheusRulesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search prometheus rules list o k response a status code equal to that given
func (o *SearchPrometheusRulesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchPrometheusRulesListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListOK  %+v", 200, o.Payload)
}

func (o *SearchPrometheusRulesListOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListOK  %+v", 200, o.Payload)
}

func (o *SearchPrometheusRulesListOK) GetPayload() *models.PrometheusRulesSearchList {
	return o.Payload
}

func (o *SearchPrometheusRulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrometheusRulesSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPrometheusRulesListBadRequest creates a SearchPrometheusRulesListBadRequest with default headers values
func NewSearchPrometheusRulesListBadRequest() *SearchPrometheusRulesListBadRequest {
	return &SearchPrometheusRulesListBadRequest{}
}

/*
SearchPrometheusRulesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchPrometheusRulesListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this search prometheus rules list bad request response has a 2xx status code
func (o *SearchPrometheusRulesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search prometheus rules list bad request response has a 3xx status code
func (o *SearchPrometheusRulesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search prometheus rules list bad request response has a 4xx status code
func (o *SearchPrometheusRulesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search prometheus rules list bad request response has a 5xx status code
func (o *SearchPrometheusRulesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search prometheus rules list bad request response a status code equal to that given
func (o *SearchPrometheusRulesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchPrometheusRulesListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPrometheusRulesListBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListBadRequest  %+v", 400, o.Payload)
}

func (o *SearchPrometheusRulesListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *SearchPrometheusRulesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPrometheusRulesListUnauthorized creates a SearchPrometheusRulesListUnauthorized with default headers values
func NewSearchPrometheusRulesListUnauthorized() *SearchPrometheusRulesListUnauthorized {
	return &SearchPrometheusRulesListUnauthorized{}
}

/*
SearchPrometheusRulesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchPrometheusRulesListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search prometheus rules list unauthorized response has a 2xx status code
func (o *SearchPrometheusRulesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search prometheus rules list unauthorized response has a 3xx status code
func (o *SearchPrometheusRulesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search prometheus rules list unauthorized response has a 4xx status code
func (o *SearchPrometheusRulesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search prometheus rules list unauthorized response has a 5xx status code
func (o *SearchPrometheusRulesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search prometheus rules list unauthorized response a status code equal to that given
func (o *SearchPrometheusRulesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchPrometheusRulesListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPrometheusRulesListUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchPrometheusRulesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPrometheusRulesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPrometheusRulesListForbidden creates a SearchPrometheusRulesListForbidden with default headers values
func NewSearchPrometheusRulesListForbidden() *SearchPrometheusRulesListForbidden {
	return &SearchPrometheusRulesListForbidden{}
}

/*
SearchPrometheusRulesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchPrometheusRulesListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search prometheus rules list forbidden response has a 2xx status code
func (o *SearchPrometheusRulesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search prometheus rules list forbidden response has a 3xx status code
func (o *SearchPrometheusRulesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search prometheus rules list forbidden response has a 4xx status code
func (o *SearchPrometheusRulesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search prometheus rules list forbidden response has a 5xx status code
func (o *SearchPrometheusRulesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search prometheus rules list forbidden response a status code equal to that given
func (o *SearchPrometheusRulesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchPrometheusRulesListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPrometheusRulesListForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListForbidden  %+v", 403, o.Payload)
}

func (o *SearchPrometheusRulesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPrometheusRulesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPrometheusRulesListNotFound creates a SearchPrometheusRulesListNotFound with default headers values
func NewSearchPrometheusRulesListNotFound() *SearchPrometheusRulesListNotFound {
	return &SearchPrometheusRulesListNotFound{}
}

/*
SearchPrometheusRulesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchPrometheusRulesListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this search prometheus rules list not found response has a 2xx status code
func (o *SearchPrometheusRulesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search prometheus rules list not found response has a 3xx status code
func (o *SearchPrometheusRulesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search prometheus rules list not found response has a 4xx status code
func (o *SearchPrometheusRulesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search prometheus rules list not found response has a 5xx status code
func (o *SearchPrometheusRulesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search prometheus rules list not found response a status code equal to that given
func (o *SearchPrometheusRulesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SearchPrometheusRulesListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPrometheusRulesListNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListNotFound  %+v", 404, o.Payload)
}

func (o *SearchPrometheusRulesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchPrometheusRulesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchPrometheusRulesListInternalServerError creates a SearchPrometheusRulesListInternalServerError with default headers values
func NewSearchPrometheusRulesListInternalServerError() *SearchPrometheusRulesListInternalServerError {
	return &SearchPrometheusRulesListInternalServerError{}
}

/*
SearchPrometheusRulesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchPrometheusRulesListInternalServerError struct {
}

// IsSuccess returns true when this search prometheus rules list internal server error response has a 2xx status code
func (o *SearchPrometheusRulesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search prometheus rules list internal server error response has a 3xx status code
func (o *SearchPrometheusRulesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search prometheus rules list internal server error response has a 4xx status code
func (o *SearchPrometheusRulesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search prometheus rules list internal server error response has a 5xx status code
func (o *SearchPrometheusRulesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search prometheus rules list internal server error response a status code equal to that given
func (o *SearchPrometheusRulesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchPrometheusRulesListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListInternalServerError ", 500)
}

func (o *SearchPrometheusRulesListInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/prometheus-rules][%d] searchPrometheusRulesListInternalServerError ", 500)
}

func (o *SearchPrometheusRulesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
