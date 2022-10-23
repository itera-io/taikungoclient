// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneProjectDetailsReader is a Reader for the StandAloneProjectDetails structure.
type StandAloneProjectDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneProjectDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneProjectDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneProjectDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneProjectDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneProjectDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneProjectDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneProjectDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneProjectDetailsOK creates a StandAloneProjectDetailsOK with default headers values
func NewStandAloneProjectDetailsOK() *StandAloneProjectDetailsOK {
	return &StandAloneProjectDetailsOK{}
}

/*
StandAloneProjectDetailsOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneProjectDetailsOK struct {
	Payload *models.ProjectFullListDto
}

// IsSuccess returns true when this stand alone project details o k response has a 2xx status code
func (o *StandAloneProjectDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone project details o k response has a 3xx status code
func (o *StandAloneProjectDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone project details o k response has a 4xx status code
func (o *StandAloneProjectDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone project details o k response has a 5xx status code
func (o *StandAloneProjectDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone project details o k response a status code equal to that given
func (o *StandAloneProjectDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneProjectDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsOK  %+v", 200, o.Payload)
}

func (o *StandAloneProjectDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsOK  %+v", 200, o.Payload)
}

func (o *StandAloneProjectDetailsOK) GetPayload() *models.ProjectFullListDto {
	return o.Payload
}

func (o *StandAloneProjectDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectFullListDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProjectDetailsBadRequest creates a StandAloneProjectDetailsBadRequest with default headers values
func NewStandAloneProjectDetailsBadRequest() *StandAloneProjectDetailsBadRequest {
	return &StandAloneProjectDetailsBadRequest{}
}

/*
StandAloneProjectDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneProjectDetailsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone project details bad request response has a 2xx status code
func (o *StandAloneProjectDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone project details bad request response has a 3xx status code
func (o *StandAloneProjectDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone project details bad request response has a 4xx status code
func (o *StandAloneProjectDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone project details bad request response has a 5xx status code
func (o *StandAloneProjectDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone project details bad request response a status code equal to that given
func (o *StandAloneProjectDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneProjectDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProjectDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProjectDetailsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneProjectDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProjectDetailsUnauthorized creates a StandAloneProjectDetailsUnauthorized with default headers values
func NewStandAloneProjectDetailsUnauthorized() *StandAloneProjectDetailsUnauthorized {
	return &StandAloneProjectDetailsUnauthorized{}
}

/*
StandAloneProjectDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneProjectDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone project details unauthorized response has a 2xx status code
func (o *StandAloneProjectDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone project details unauthorized response has a 3xx status code
func (o *StandAloneProjectDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone project details unauthorized response has a 4xx status code
func (o *StandAloneProjectDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone project details unauthorized response has a 5xx status code
func (o *StandAloneProjectDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone project details unauthorized response a status code equal to that given
func (o *StandAloneProjectDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneProjectDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProjectDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProjectDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProjectDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProjectDetailsForbidden creates a StandAloneProjectDetailsForbidden with default headers values
func NewStandAloneProjectDetailsForbidden() *StandAloneProjectDetailsForbidden {
	return &StandAloneProjectDetailsForbidden{}
}

/*
StandAloneProjectDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneProjectDetailsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone project details forbidden response has a 2xx status code
func (o *StandAloneProjectDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone project details forbidden response has a 3xx status code
func (o *StandAloneProjectDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone project details forbidden response has a 4xx status code
func (o *StandAloneProjectDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone project details forbidden response has a 5xx status code
func (o *StandAloneProjectDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone project details forbidden response a status code equal to that given
func (o *StandAloneProjectDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneProjectDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProjectDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProjectDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProjectDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProjectDetailsNotFound creates a StandAloneProjectDetailsNotFound with default headers values
func NewStandAloneProjectDetailsNotFound() *StandAloneProjectDetailsNotFound {
	return &StandAloneProjectDetailsNotFound{}
}

/*
StandAloneProjectDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneProjectDetailsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone project details not found response has a 2xx status code
func (o *StandAloneProjectDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone project details not found response has a 3xx status code
func (o *StandAloneProjectDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone project details not found response has a 4xx status code
func (o *StandAloneProjectDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone project details not found response has a 5xx status code
func (o *StandAloneProjectDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone project details not found response a status code equal to that given
func (o *StandAloneProjectDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneProjectDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProjectDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProjectDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProjectDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProjectDetailsInternalServerError creates a StandAloneProjectDetailsInternalServerError with default headers values
func NewStandAloneProjectDetailsInternalServerError() *StandAloneProjectDetailsInternalServerError {
	return &StandAloneProjectDetailsInternalServerError{}
}

/*
StandAloneProjectDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneProjectDetailsInternalServerError struct {
}

// IsSuccess returns true when this stand alone project details internal server error response has a 2xx status code
func (o *StandAloneProjectDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone project details internal server error response has a 3xx status code
func (o *StandAloneProjectDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone project details internal server error response has a 4xx status code
func (o *StandAloneProjectDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone project details internal server error response has a 5xx status code
func (o *StandAloneProjectDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone project details internal server error response a status code equal to that given
func (o *StandAloneProjectDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneProjectDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsInternalServerError ", 500)
}

func (o *StandAloneProjectDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/project/{projectId}][%d] standAloneProjectDetailsInternalServerError ", 500)
}

func (o *StandAloneProjectDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
