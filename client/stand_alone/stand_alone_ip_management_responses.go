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

// StandAloneIPManagementReader is a Reader for the StandAloneIPManagement structure.
type StandAloneIPManagementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneIPManagementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneIPManagementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneIPManagementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneIPManagementUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneIPManagementForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneIPManagementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneIPManagementInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneIPManagementOK creates a StandAloneIPManagementOK with default headers values
func NewStandAloneIPManagementOK() *StandAloneIPManagementOK {
	return &StandAloneIPManagementOK{}
}

/*
StandAloneIPManagementOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneIPManagementOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone Ip management o k response has a 2xx status code
func (o *StandAloneIPManagementOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Ip management o k response has a 3xx status code
func (o *StandAloneIPManagementOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Ip management o k response has a 4xx status code
func (o *StandAloneIPManagementOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Ip management o k response has a 5xx status code
func (o *StandAloneIPManagementOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Ip management o k response a status code equal to that given
func (o *StandAloneIPManagementOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneIPManagementOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementOK  %+v", 200, o.Payload)
}

func (o *StandAloneIPManagementOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementOK  %+v", 200, o.Payload)
}

func (o *StandAloneIPManagementOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneIPManagementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneIPManagementBadRequest creates a StandAloneIPManagementBadRequest with default headers values
func NewStandAloneIPManagementBadRequest() *StandAloneIPManagementBadRequest {
	return &StandAloneIPManagementBadRequest{}
}

/*
StandAloneIPManagementBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneIPManagementBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone Ip management bad request response has a 2xx status code
func (o *StandAloneIPManagementBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Ip management bad request response has a 3xx status code
func (o *StandAloneIPManagementBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Ip management bad request response has a 4xx status code
func (o *StandAloneIPManagementBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Ip management bad request response has a 5xx status code
func (o *StandAloneIPManagementBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Ip management bad request response a status code equal to that given
func (o *StandAloneIPManagementBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneIPManagementBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneIPManagementBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneIPManagementBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneIPManagementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneIPManagementUnauthorized creates a StandAloneIPManagementUnauthorized with default headers values
func NewStandAloneIPManagementUnauthorized() *StandAloneIPManagementUnauthorized {
	return &StandAloneIPManagementUnauthorized{}
}

/*
StandAloneIPManagementUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneIPManagementUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone Ip management unauthorized response has a 2xx status code
func (o *StandAloneIPManagementUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Ip management unauthorized response has a 3xx status code
func (o *StandAloneIPManagementUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Ip management unauthorized response has a 4xx status code
func (o *StandAloneIPManagementUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Ip management unauthorized response has a 5xx status code
func (o *StandAloneIPManagementUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Ip management unauthorized response a status code equal to that given
func (o *StandAloneIPManagementUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneIPManagementUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneIPManagementUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneIPManagementUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneIPManagementUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneIPManagementForbidden creates a StandAloneIPManagementForbidden with default headers values
func NewStandAloneIPManagementForbidden() *StandAloneIPManagementForbidden {
	return &StandAloneIPManagementForbidden{}
}

/*
StandAloneIPManagementForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneIPManagementForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone Ip management forbidden response has a 2xx status code
func (o *StandAloneIPManagementForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Ip management forbidden response has a 3xx status code
func (o *StandAloneIPManagementForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Ip management forbidden response has a 4xx status code
func (o *StandAloneIPManagementForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Ip management forbidden response has a 5xx status code
func (o *StandAloneIPManagementForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Ip management forbidden response a status code equal to that given
func (o *StandAloneIPManagementForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneIPManagementForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneIPManagementForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneIPManagementForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneIPManagementForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneIPManagementNotFound creates a StandAloneIPManagementNotFound with default headers values
func NewStandAloneIPManagementNotFound() *StandAloneIPManagementNotFound {
	return &StandAloneIPManagementNotFound{}
}

/*
StandAloneIPManagementNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneIPManagementNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone Ip management not found response has a 2xx status code
func (o *StandAloneIPManagementNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Ip management not found response has a 3xx status code
func (o *StandAloneIPManagementNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Ip management not found response has a 4xx status code
func (o *StandAloneIPManagementNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Ip management not found response has a 5xx status code
func (o *StandAloneIPManagementNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Ip management not found response a status code equal to that given
func (o *StandAloneIPManagementNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneIPManagementNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneIPManagementNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneIPManagementNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneIPManagementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneIPManagementInternalServerError creates a StandAloneIPManagementInternalServerError with default headers values
func NewStandAloneIPManagementInternalServerError() *StandAloneIPManagementInternalServerError {
	return &StandAloneIPManagementInternalServerError{}
}

/*
StandAloneIPManagementInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneIPManagementInternalServerError struct {
}

// IsSuccess returns true when this stand alone Ip management internal server error response has a 2xx status code
func (o *StandAloneIPManagementInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Ip management internal server error response has a 3xx status code
func (o *StandAloneIPManagementInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Ip management internal server error response has a 4xx status code
func (o *StandAloneIPManagementInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Ip management internal server error response has a 5xx status code
func (o *StandAloneIPManagementInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Ip management internal server error response a status code equal to that given
func (o *StandAloneIPManagementInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneIPManagementInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementInternalServerError ", 500)
}

func (o *StandAloneIPManagementInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/ip/management][%d] standAloneIpManagementInternalServerError ", 500)
}

func (o *StandAloneIPManagementInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
