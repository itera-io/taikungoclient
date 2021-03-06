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

// PreDefinedQueriesUpdatePrometheusDashboardReader is a Reader for the PreDefinedQueriesUpdatePrometheusDashboard structure.
type PreDefinedQueriesUpdatePrometheusDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PreDefinedQueriesUpdatePrometheusDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPreDefinedQueriesUpdatePrometheusDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPreDefinedQueriesUpdatePrometheusDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPreDefinedQueriesUpdatePrometheusDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPreDefinedQueriesUpdatePrometheusDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPreDefinedQueriesUpdatePrometheusDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPreDefinedQueriesUpdatePrometheusDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPreDefinedQueriesUpdatePrometheusDashboardOK creates a PreDefinedQueriesUpdatePrometheusDashboardOK with default headers values
func NewPreDefinedQueriesUpdatePrometheusDashboardOK() *PreDefinedQueriesUpdatePrometheusDashboardOK {
	return &PreDefinedQueriesUpdatePrometheusDashboardOK{}
}

/* PreDefinedQueriesUpdatePrometheusDashboardOK describes a response with status code 200, with default header values.

Success
*/
type PreDefinedQueriesUpdatePrometheusDashboardOK struct {
	Payload models.Unit
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/PreDefinedQueries/prometheus/dashboard/update][%d] preDefinedQueriesUpdatePrometheusDashboardOK  %+v", 200, o.Payload)
}
func (o *PreDefinedQueriesUpdatePrometheusDashboardOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesUpdatePrometheusDashboardBadRequest creates a PreDefinedQueriesUpdatePrometheusDashboardBadRequest with default headers values
func NewPreDefinedQueriesUpdatePrometheusDashboardBadRequest() *PreDefinedQueriesUpdatePrometheusDashboardBadRequest {
	return &PreDefinedQueriesUpdatePrometheusDashboardBadRequest{}
}

/* PreDefinedQueriesUpdatePrometheusDashboardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PreDefinedQueriesUpdatePrometheusDashboardBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/PreDefinedQueries/prometheus/dashboard/update][%d] preDefinedQueriesUpdatePrometheusDashboardBadRequest  %+v", 400, o.Payload)
}
func (o *PreDefinedQueriesUpdatePrometheusDashboardBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesUpdatePrometheusDashboardUnauthorized creates a PreDefinedQueriesUpdatePrometheusDashboardUnauthorized with default headers values
func NewPreDefinedQueriesUpdatePrometheusDashboardUnauthorized() *PreDefinedQueriesUpdatePrometheusDashboardUnauthorized {
	return &PreDefinedQueriesUpdatePrometheusDashboardUnauthorized{}
}

/* PreDefinedQueriesUpdatePrometheusDashboardUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PreDefinedQueriesUpdatePrometheusDashboardUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/PreDefinedQueries/prometheus/dashboard/update][%d] preDefinedQueriesUpdatePrometheusDashboardUnauthorized  %+v", 401, o.Payload)
}
func (o *PreDefinedQueriesUpdatePrometheusDashboardUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesUpdatePrometheusDashboardForbidden creates a PreDefinedQueriesUpdatePrometheusDashboardForbidden with default headers values
func NewPreDefinedQueriesUpdatePrometheusDashboardForbidden() *PreDefinedQueriesUpdatePrometheusDashboardForbidden {
	return &PreDefinedQueriesUpdatePrometheusDashboardForbidden{}
}

/* PreDefinedQueriesUpdatePrometheusDashboardForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PreDefinedQueriesUpdatePrometheusDashboardForbidden struct {
	Payload *models.ProblemDetails
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/PreDefinedQueries/prometheus/dashboard/update][%d] preDefinedQueriesUpdatePrometheusDashboardForbidden  %+v", 403, o.Payload)
}
func (o *PreDefinedQueriesUpdatePrometheusDashboardForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesUpdatePrometheusDashboardNotFound creates a PreDefinedQueriesUpdatePrometheusDashboardNotFound with default headers values
func NewPreDefinedQueriesUpdatePrometheusDashboardNotFound() *PreDefinedQueriesUpdatePrometheusDashboardNotFound {
	return &PreDefinedQueriesUpdatePrometheusDashboardNotFound{}
}

/* PreDefinedQueriesUpdatePrometheusDashboardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PreDefinedQueriesUpdatePrometheusDashboardNotFound struct {
	Payload *models.ProblemDetails
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/PreDefinedQueries/prometheus/dashboard/update][%d] preDefinedQueriesUpdatePrometheusDashboardNotFound  %+v", 404, o.Payload)
}
func (o *PreDefinedQueriesUpdatePrometheusDashboardNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPreDefinedQueriesUpdatePrometheusDashboardInternalServerError creates a PreDefinedQueriesUpdatePrometheusDashboardInternalServerError with default headers values
func NewPreDefinedQueriesUpdatePrometheusDashboardInternalServerError() *PreDefinedQueriesUpdatePrometheusDashboardInternalServerError {
	return &PreDefinedQueriesUpdatePrometheusDashboardInternalServerError{}
}

/* PreDefinedQueriesUpdatePrometheusDashboardInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PreDefinedQueriesUpdatePrometheusDashboardInternalServerError struct {
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/PreDefinedQueries/prometheus/dashboard/update][%d] preDefinedQueriesUpdatePrometheusDashboardInternalServerError ", 500)
}

func (o *PreDefinedQueriesUpdatePrometheusDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
