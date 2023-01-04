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

// StandAloneDetailsReader is a Reader for the StandAloneDetails structure.
type StandAloneDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneDetailsOK creates a StandAloneDetailsOK with default headers values
func NewStandAloneDetailsOK() *StandAloneDetailsOK {
	return &StandAloneDetailsOK{}
}

/*
StandAloneDetailsOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneDetailsOK struct {
	Payload *models.StandAloneVMListForDetails
}

// IsSuccess returns true when this stand alone details o k response has a 2xx status code
func (o *StandAloneDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone details o k response has a 3xx status code
func (o *StandAloneDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone details o k response has a 4xx status code
func (o *StandAloneDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone details o k response has a 5xx status code
func (o *StandAloneDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone details o k response a status code equal to that given
func (o *StandAloneDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsOK  %+v", 200, o.Payload)
}

func (o *StandAloneDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsOK  %+v", 200, o.Payload)
}

func (o *StandAloneDetailsOK) GetPayload() *models.StandAloneVMListForDetails {
	return o.Payload
}

func (o *StandAloneDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandAloneVMListForDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDetailsBadRequest creates a StandAloneDetailsBadRequest with default headers values
func NewStandAloneDetailsBadRequest() *StandAloneDetailsBadRequest {
	return &StandAloneDetailsBadRequest{}
}

/*
StandAloneDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneDetailsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone details bad request response has a 2xx status code
func (o *StandAloneDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone details bad request response has a 3xx status code
func (o *StandAloneDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone details bad request response has a 4xx status code
func (o *StandAloneDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone details bad request response has a 5xx status code
func (o *StandAloneDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone details bad request response a status code equal to that given
func (o *StandAloneDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneDetailsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDetailsUnauthorized creates a StandAloneDetailsUnauthorized with default headers values
func NewStandAloneDetailsUnauthorized() *StandAloneDetailsUnauthorized {
	return &StandAloneDetailsUnauthorized{}
}

/*
StandAloneDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone details unauthorized response has a 2xx status code
func (o *StandAloneDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone details unauthorized response has a 3xx status code
func (o *StandAloneDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone details unauthorized response has a 4xx status code
func (o *StandAloneDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone details unauthorized response has a 5xx status code
func (o *StandAloneDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone details unauthorized response a status code equal to that given
func (o *StandAloneDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDetailsForbidden creates a StandAloneDetailsForbidden with default headers values
func NewStandAloneDetailsForbidden() *StandAloneDetailsForbidden {
	return &StandAloneDetailsForbidden{}
}

/*
StandAloneDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneDetailsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone details forbidden response has a 2xx status code
func (o *StandAloneDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone details forbidden response has a 3xx status code
func (o *StandAloneDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone details forbidden response has a 4xx status code
func (o *StandAloneDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone details forbidden response has a 5xx status code
func (o *StandAloneDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone details forbidden response a status code equal to that given
func (o *StandAloneDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDetailsNotFound creates a StandAloneDetailsNotFound with default headers values
func NewStandAloneDetailsNotFound() *StandAloneDetailsNotFound {
	return &StandAloneDetailsNotFound{}
}

/*
StandAloneDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneDetailsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone details not found response has a 2xx status code
func (o *StandAloneDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone details not found response has a 3xx status code
func (o *StandAloneDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone details not found response has a 4xx status code
func (o *StandAloneDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone details not found response has a 5xx status code
func (o *StandAloneDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone details not found response a status code equal to that given
func (o *StandAloneDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDetailsInternalServerError creates a StandAloneDetailsInternalServerError with default headers values
func NewStandAloneDetailsInternalServerError() *StandAloneDetailsInternalServerError {
	return &StandAloneDetailsInternalServerError{}
}

/*
StandAloneDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneDetailsInternalServerError struct {
}

// IsSuccess returns true when this stand alone details internal server error response has a 2xx status code
func (o *StandAloneDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone details internal server error response has a 3xx status code
func (o *StandAloneDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone details internal server error response has a 4xx status code
func (o *StandAloneDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone details internal server error response has a 5xx status code
func (o *StandAloneDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone details internal server error response a status code equal to that given
func (o *StandAloneDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsInternalServerError ", 500)
}

func (o *StandAloneDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAlone/{projectId}][%d] standAloneDetailsInternalServerError ", 500)
}

func (o *StandAloneDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
