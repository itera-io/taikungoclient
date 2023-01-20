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

// PreDefinedQueriesDeletePrometheusDashboardReader is a Reader for the PreDefinedQueriesDeletePrometheusDashboard structure.
type PreDefinedQueriesDeletePrometheusDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PreDefinedQueriesDeletePrometheusDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPreDefinedQueriesDeletePrometheusDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPreDefinedQueriesDeletePrometheusDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPreDefinedQueriesDeletePrometheusDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPreDefinedQueriesDeletePrometheusDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPreDefinedQueriesDeletePrometheusDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPreDefinedQueriesDeletePrometheusDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPreDefinedQueriesDeletePrometheusDashboardOK creates a PreDefinedQueriesDeletePrometheusDashboardOK with default headers values
func NewPreDefinedQueriesDeletePrometheusDashboardOK() *PreDefinedQueriesDeletePrometheusDashboardOK {
	return &PreDefinedQueriesDeletePrometheusDashboardOK{}
}

/*
PreDefinedQueriesDeletePrometheusDashboardOK describes a response with status code 200, with default header values.

Success
*/
type PreDefinedQueriesDeletePrometheusDashboardOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this pre defined queries delete prometheus dashboard o k response has a 2xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pre defined queries delete prometheus dashboard o k response has a 3xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries delete prometheus dashboard o k response has a 4xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pre defined queries delete prometheus dashboard o k response has a 5xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries delete prometheus dashboard o k response a status code equal to that given
func (o *PreDefinedQueriesDeletePrometheusDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pre defined queries delete prometheus dashboard o k response
func (o *PreDefinedQueriesDeletePrometheusDashboardOK) Code() int {
	return 200
}

func (o *PreDefinedQueriesDeletePrometheusDashboardOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardOK  %+v", 200, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardOK  %+v", 200, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PreDefinedQueriesDeletePrometheusDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesDeletePrometheusDashboardBadRequest creates a PreDefinedQueriesDeletePrometheusDashboardBadRequest with default headers values
func NewPreDefinedQueriesDeletePrometheusDashboardBadRequest() *PreDefinedQueriesDeletePrometheusDashboardBadRequest {
	return &PreDefinedQueriesDeletePrometheusDashboardBadRequest{}
}

/*
PreDefinedQueriesDeletePrometheusDashboardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PreDefinedQueriesDeletePrometheusDashboardBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this pre defined queries delete prometheus dashboard bad request response has a 2xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries delete prometheus dashboard bad request response has a 3xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries delete prometheus dashboard bad request response has a 4xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries delete prometheus dashboard bad request response has a 5xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries delete prometheus dashboard bad request response a status code equal to that given
func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pre defined queries delete prometheus dashboard bad request response
func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) Code() int {
	return 400
}

func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PreDefinedQueriesDeletePrometheusDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesDeletePrometheusDashboardUnauthorized creates a PreDefinedQueriesDeletePrometheusDashboardUnauthorized with default headers values
func NewPreDefinedQueriesDeletePrometheusDashboardUnauthorized() *PreDefinedQueriesDeletePrometheusDashboardUnauthorized {
	return &PreDefinedQueriesDeletePrometheusDashboardUnauthorized{}
}

/*
PreDefinedQueriesDeletePrometheusDashboardUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PreDefinedQueriesDeletePrometheusDashboardUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this pre defined queries delete prometheus dashboard unauthorized response has a 2xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries delete prometheus dashboard unauthorized response has a 3xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries delete prometheus dashboard unauthorized response has a 4xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries delete prometheus dashboard unauthorized response has a 5xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries delete prometheus dashboard unauthorized response a status code equal to that given
func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pre defined queries delete prometheus dashboard unauthorized response
func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) Code() int {
	return 401
}

func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PreDefinedQueriesDeletePrometheusDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesDeletePrometheusDashboardForbidden creates a PreDefinedQueriesDeletePrometheusDashboardForbidden with default headers values
func NewPreDefinedQueriesDeletePrometheusDashboardForbidden() *PreDefinedQueriesDeletePrometheusDashboardForbidden {
	return &PreDefinedQueriesDeletePrometheusDashboardForbidden{}
}

/*
PreDefinedQueriesDeletePrometheusDashboardForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PreDefinedQueriesDeletePrometheusDashboardForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this pre defined queries delete prometheus dashboard forbidden response has a 2xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries delete prometheus dashboard forbidden response has a 3xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries delete prometheus dashboard forbidden response has a 4xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries delete prometheus dashboard forbidden response has a 5xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries delete prometheus dashboard forbidden response a status code equal to that given
func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pre defined queries delete prometheus dashboard forbidden response
func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) Code() int {
	return 403
}

func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PreDefinedQueriesDeletePrometheusDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesDeletePrometheusDashboardNotFound creates a PreDefinedQueriesDeletePrometheusDashboardNotFound with default headers values
func NewPreDefinedQueriesDeletePrometheusDashboardNotFound() *PreDefinedQueriesDeletePrometheusDashboardNotFound {
	return &PreDefinedQueriesDeletePrometheusDashboardNotFound{}
}

/*
PreDefinedQueriesDeletePrometheusDashboardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PreDefinedQueriesDeletePrometheusDashboardNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this pre defined queries delete prometheus dashboard not found response has a 2xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries delete prometheus dashboard not found response has a 3xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries delete prometheus dashboard not found response has a 4xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pre defined queries delete prometheus dashboard not found response has a 5xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pre defined queries delete prometheus dashboard not found response a status code equal to that given
func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pre defined queries delete prometheus dashboard not found response
func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) Code() int {
	return 404
}

func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PreDefinedQueriesDeletePrometheusDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesDeletePrometheusDashboardInternalServerError creates a PreDefinedQueriesDeletePrometheusDashboardInternalServerError with default headers values
func NewPreDefinedQueriesDeletePrometheusDashboardInternalServerError() *PreDefinedQueriesDeletePrometheusDashboardInternalServerError {
	return &PreDefinedQueriesDeletePrometheusDashboardInternalServerError{}
}

/*
PreDefinedQueriesDeletePrometheusDashboardInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PreDefinedQueriesDeletePrometheusDashboardInternalServerError struct {
}

// IsSuccess returns true when this pre defined queries delete prometheus dashboard internal server error response has a 2xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pre defined queries delete prometheus dashboard internal server error response has a 3xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pre defined queries delete prometheus dashboard internal server error response has a 4xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pre defined queries delete prometheus dashboard internal server error response has a 5xx status code
func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pre defined queries delete prometheus dashboard internal server error response a status code equal to that given
func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pre defined queries delete prometheus dashboard internal server error response
func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) Code() int {
	return 500
}

func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardInternalServerError ", 500)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}][%d] preDefinedQueriesDeletePrometheusDashboardInternalServerError ", 500)
}

func (o *PreDefinedQueriesDeletePrometheusDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
