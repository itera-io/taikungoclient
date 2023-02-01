// Code generated by go-swagger; DO NOT EDIT.

package pre_defined_queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// PreDefinedQueriesGetPrometheusDashboardListReader is a Reader for the PreDefinedQueriesGetPrometheusDashboardList structure.
type PreDefinedQueriesGetPrometheusDashboardListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PreDefinedQueriesGetPrometheusDashboardListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPreDefinedQueriesGetPrometheusDashboardListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPreDefinedQueriesGetPrometheusDashboardListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPreDefinedQueriesGetPrometheusDashboardListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPreDefinedQueriesGetPrometheusDashboardListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPreDefinedQueriesGetPrometheusDashboardListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPreDefinedQueriesGetPrometheusDashboardListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPreDefinedQueriesGetPrometheusDashboardListOK creates a PreDefinedQueriesGetPrometheusDashboardListOK with default headers values
func NewPreDefinedQueriesGetPrometheusDashboardListOK() *PreDefinedQueriesGetPrometheusDashboardListOK {
	return &PreDefinedQueriesGetPrometheusDashboardListOK{}
}

/*
PreDefinedQueriesGetPrometheusDashboardListOK describes a response with status code 200, with default header values.

Success
*/
type PreDefinedQueriesGetPrometheusDashboardListOK struct {
	Payload []*models.PrometheusDashboardListDto
}

// IsSuccess returns true when this pre defined queries get prometheus dashboard list o k response has a 2xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pre defined queries get prometheus dashboard list o k response has a 3xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries get prometheus dashboard list o k response has a 4xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pre defined queries get prometheus dashboard list o k response has a 5xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries get prometheus dashboard list o k response a status code equal to that given
func (o *PreDefinedQueriesGetPrometheusDashboardListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pre defined queries get prometheus dashboard list o k response
func (o *PreDefinedQueriesGetPrometheusDashboardListOK) Code() int {
	return 200
}

func (o *PreDefinedQueriesGetPrometheusDashboardListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListOK  %+v", 200, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListOK  %+v", 200, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListOK) GetPayload() []*models.PrometheusDashboardListDto {
	return o.Payload
}

func (o *PreDefinedQueriesGetPrometheusDashboardListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesGetPrometheusDashboardListBadRequest creates a PreDefinedQueriesGetPrometheusDashboardListBadRequest with default headers values
func NewPreDefinedQueriesGetPrometheusDashboardListBadRequest() *PreDefinedQueriesGetPrometheusDashboardListBadRequest {
	return &PreDefinedQueriesGetPrometheusDashboardListBadRequest{}
}

/*
PreDefinedQueriesGetPrometheusDashboardListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PreDefinedQueriesGetPrometheusDashboardListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this pre defined queries get prometheus dashboard list bad request response has a 2xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries get prometheus dashboard list bad request response has a 3xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries get prometheus dashboard list bad request response has a 4xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries get prometheus dashboard list bad request response has a 5xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries get prometheus dashboard list bad request response a status code equal to that given
func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pre defined queries get prometheus dashboard list bad request response
func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) Code() int {
	return 400
}

func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListBadRequest  %+v", 400, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListBadRequest  %+v", 400, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesGetPrometheusDashboardListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesGetPrometheusDashboardListUnauthorized creates a PreDefinedQueriesGetPrometheusDashboardListUnauthorized with default headers values
func NewPreDefinedQueriesGetPrometheusDashboardListUnauthorized() *PreDefinedQueriesGetPrometheusDashboardListUnauthorized {
	return &PreDefinedQueriesGetPrometheusDashboardListUnauthorized{}
}

/*
PreDefinedQueriesGetPrometheusDashboardListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PreDefinedQueriesGetPrometheusDashboardListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this pre defined queries get prometheus dashboard list unauthorized response has a 2xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries get prometheus dashboard list unauthorized response has a 3xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries get prometheus dashboard list unauthorized response has a 4xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries get prometheus dashboard list unauthorized response has a 5xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries get prometheus dashboard list unauthorized response a status code equal to that given
func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pre defined queries get prometheus dashboard list unauthorized response
func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) Code() int {
	return 401
}

func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListUnauthorized  %+v", 401, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListUnauthorized  %+v", 401, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesGetPrometheusDashboardListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesGetPrometheusDashboardListForbidden creates a PreDefinedQueriesGetPrometheusDashboardListForbidden with default headers values
func NewPreDefinedQueriesGetPrometheusDashboardListForbidden() *PreDefinedQueriesGetPrometheusDashboardListForbidden {
	return &PreDefinedQueriesGetPrometheusDashboardListForbidden{}
}

/*
PreDefinedQueriesGetPrometheusDashboardListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PreDefinedQueriesGetPrometheusDashboardListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this pre defined queries get prometheus dashboard list forbidden response has a 2xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries get prometheus dashboard list forbidden response has a 3xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries get prometheus dashboard list forbidden response has a 4xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries get prometheus dashboard list forbidden response has a 5xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries get prometheus dashboard list forbidden response a status code equal to that given
func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pre defined queries get prometheus dashboard list forbidden response
func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) Code() int {
	return 403
}

func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListForbidden  %+v", 403, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListForbidden  %+v", 403, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesGetPrometheusDashboardListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesGetPrometheusDashboardListNotFound creates a PreDefinedQueriesGetPrometheusDashboardListNotFound with default headers values
func NewPreDefinedQueriesGetPrometheusDashboardListNotFound() *PreDefinedQueriesGetPrometheusDashboardListNotFound {
	return &PreDefinedQueriesGetPrometheusDashboardListNotFound{}
}

/*
PreDefinedQueriesGetPrometheusDashboardListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PreDefinedQueriesGetPrometheusDashboardListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this pre defined queries get prometheus dashboard list not found response has a 2xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries get prometheus dashboard list not found response has a 3xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries get prometheus dashboard list not found response has a 4xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries get prometheus dashboard list not found response has a 5xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries get prometheus dashboard list not found response a status code equal to that given
func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pre defined queries get prometheus dashboard list not found response
func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) Code() int {
	return 404
}

func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListNotFound  %+v", 404, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListNotFound  %+v", 404, o.Payload)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesGetPrometheusDashboardListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesGetPrometheusDashboardListInternalServerError creates a PreDefinedQueriesGetPrometheusDashboardListInternalServerError with default headers values
func NewPreDefinedQueriesGetPrometheusDashboardListInternalServerError() *PreDefinedQueriesGetPrometheusDashboardListInternalServerError {
	return &PreDefinedQueriesGetPrometheusDashboardListInternalServerError{}
}

/*
PreDefinedQueriesGetPrometheusDashboardListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PreDefinedQueriesGetPrometheusDashboardListInternalServerError struct {
}

// IsSuccess returns true when this pre defined queries get prometheus dashboard list internal server error response has a 2xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries get prometheus dashboard list internal server error response has a 3xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries get prometheus dashboard list internal server error response has a 4xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pre defined queries get prometheus dashboard list internal server error response has a 5xx status code
func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pre defined queries get prometheus dashboard list internal server error response a status code equal to that given
func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pre defined queries get prometheus dashboard list internal server error response
func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) Code() int {
	return 500
}

func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListInternalServerError ", 500)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/PreDefinedQueries/prometheus/dashboard/list][%d] preDefinedQueriesGetPrometheusDashboardListInternalServerError ", 500)
}

func (o *PreDefinedQueriesGetPrometheusDashboardListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
