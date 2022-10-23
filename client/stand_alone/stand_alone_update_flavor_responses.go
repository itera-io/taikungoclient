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

// StandAloneUpdateFlavorReader is a Reader for the StandAloneUpdateFlavor structure.
type StandAloneUpdateFlavorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneUpdateFlavorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneUpdateFlavorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneUpdateFlavorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneUpdateFlavorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneUpdateFlavorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneUpdateFlavorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneUpdateFlavorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneUpdateFlavorOK creates a StandAloneUpdateFlavorOK with default headers values
func NewStandAloneUpdateFlavorOK() *StandAloneUpdateFlavorOK {
	return &StandAloneUpdateFlavorOK{}
}

/*
StandAloneUpdateFlavorOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneUpdateFlavorOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone update flavor o k response has a 2xx status code
func (o *StandAloneUpdateFlavorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone update flavor o k response has a 3xx status code
func (o *StandAloneUpdateFlavorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update flavor o k response has a 4xx status code
func (o *StandAloneUpdateFlavorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone update flavor o k response has a 5xx status code
func (o *StandAloneUpdateFlavorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update flavor o k response a status code equal to that given
func (o *StandAloneUpdateFlavorOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneUpdateFlavorOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorOK  %+v", 200, o.Payload)
}

func (o *StandAloneUpdateFlavorOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorOK  %+v", 200, o.Payload)
}

func (o *StandAloneUpdateFlavorOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneUpdateFlavorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateFlavorBadRequest creates a StandAloneUpdateFlavorBadRequest with default headers values
func NewStandAloneUpdateFlavorBadRequest() *StandAloneUpdateFlavorBadRequest {
	return &StandAloneUpdateFlavorBadRequest{}
}

/*
StandAloneUpdateFlavorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneUpdateFlavorBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone update flavor bad request response has a 2xx status code
func (o *StandAloneUpdateFlavorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update flavor bad request response has a 3xx status code
func (o *StandAloneUpdateFlavorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update flavor bad request response has a 4xx status code
func (o *StandAloneUpdateFlavorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update flavor bad request response has a 5xx status code
func (o *StandAloneUpdateFlavorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update flavor bad request response a status code equal to that given
func (o *StandAloneUpdateFlavorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneUpdateFlavorBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneUpdateFlavorBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneUpdateFlavorBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneUpdateFlavorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateFlavorUnauthorized creates a StandAloneUpdateFlavorUnauthorized with default headers values
func NewStandAloneUpdateFlavorUnauthorized() *StandAloneUpdateFlavorUnauthorized {
	return &StandAloneUpdateFlavorUnauthorized{}
}

/*
StandAloneUpdateFlavorUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneUpdateFlavorUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone update flavor unauthorized response has a 2xx status code
func (o *StandAloneUpdateFlavorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update flavor unauthorized response has a 3xx status code
func (o *StandAloneUpdateFlavorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update flavor unauthorized response has a 4xx status code
func (o *StandAloneUpdateFlavorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update flavor unauthorized response has a 5xx status code
func (o *StandAloneUpdateFlavorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update flavor unauthorized response a status code equal to that given
func (o *StandAloneUpdateFlavorUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneUpdateFlavorUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneUpdateFlavorUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneUpdateFlavorUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneUpdateFlavorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateFlavorForbidden creates a StandAloneUpdateFlavorForbidden with default headers values
func NewStandAloneUpdateFlavorForbidden() *StandAloneUpdateFlavorForbidden {
	return &StandAloneUpdateFlavorForbidden{}
}

/*
StandAloneUpdateFlavorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneUpdateFlavorForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone update flavor forbidden response has a 2xx status code
func (o *StandAloneUpdateFlavorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update flavor forbidden response has a 3xx status code
func (o *StandAloneUpdateFlavorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update flavor forbidden response has a 4xx status code
func (o *StandAloneUpdateFlavorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update flavor forbidden response has a 5xx status code
func (o *StandAloneUpdateFlavorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update flavor forbidden response a status code equal to that given
func (o *StandAloneUpdateFlavorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneUpdateFlavorForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneUpdateFlavorForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneUpdateFlavorForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneUpdateFlavorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateFlavorNotFound creates a StandAloneUpdateFlavorNotFound with default headers values
func NewStandAloneUpdateFlavorNotFound() *StandAloneUpdateFlavorNotFound {
	return &StandAloneUpdateFlavorNotFound{}
}

/*
StandAloneUpdateFlavorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneUpdateFlavorNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone update flavor not found response has a 2xx status code
func (o *StandAloneUpdateFlavorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update flavor not found response has a 3xx status code
func (o *StandAloneUpdateFlavorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update flavor not found response has a 4xx status code
func (o *StandAloneUpdateFlavorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update flavor not found response has a 5xx status code
func (o *StandAloneUpdateFlavorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update flavor not found response a status code equal to that given
func (o *StandAloneUpdateFlavorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneUpdateFlavorNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneUpdateFlavorNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneUpdateFlavorNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneUpdateFlavorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateFlavorInternalServerError creates a StandAloneUpdateFlavorInternalServerError with default headers values
func NewStandAloneUpdateFlavorInternalServerError() *StandAloneUpdateFlavorInternalServerError {
	return &StandAloneUpdateFlavorInternalServerError{}
}

/*
StandAloneUpdateFlavorInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneUpdateFlavorInternalServerError struct {
}

// IsSuccess returns true when this stand alone update flavor internal server error response has a 2xx status code
func (o *StandAloneUpdateFlavorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update flavor internal server error response has a 3xx status code
func (o *StandAloneUpdateFlavorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update flavor internal server error response has a 4xx status code
func (o *StandAloneUpdateFlavorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone update flavor internal server error response has a 5xx status code
func (o *StandAloneUpdateFlavorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone update flavor internal server error response a status code equal to that given
func (o *StandAloneUpdateFlavorInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneUpdateFlavorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorInternalServerError ", 500)
}

func (o *StandAloneUpdateFlavorInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update/flavor][%d] standAloneUpdateFlavorInternalServerError ", 500)
}

func (o *StandAloneUpdateFlavorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
