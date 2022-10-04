// Code generated by go-swagger; DO NOT EDIT.

package flavors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// FlavorsUnbindFromProjectReader is a Reader for the FlavorsUnbindFromProject structure.
type FlavorsUnbindFromProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlavorsUnbindFromProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlavorsUnbindFromProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFlavorsUnbindFromProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFlavorsUnbindFromProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFlavorsUnbindFromProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFlavorsUnbindFromProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFlavorsUnbindFromProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFlavorsUnbindFromProjectOK creates a FlavorsUnbindFromProjectOK with default headers values
func NewFlavorsUnbindFromProjectOK() *FlavorsUnbindFromProjectOK {
	return &FlavorsUnbindFromProjectOK{}
}

/*
FlavorsUnbindFromProjectOK describes a response with status code 200, with default header values.

Success
*/
type FlavorsUnbindFromProjectOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this flavors unbind from project o k response has a 2xx status code
func (o *FlavorsUnbindFromProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flavors unbind from project o k response has a 3xx status code
func (o *FlavorsUnbindFromProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors unbind from project o k response has a 4xx status code
func (o *FlavorsUnbindFromProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors unbind from project o k response has a 5xx status code
func (o *FlavorsUnbindFromProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors unbind from project o k response a status code equal to that given
func (o *FlavorsUnbindFromProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *FlavorsUnbindFromProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectOK  %+v", 200, o.Payload)
}

func (o *FlavorsUnbindFromProjectOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectOK  %+v", 200, o.Payload)
}

func (o *FlavorsUnbindFromProjectOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *FlavorsUnbindFromProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsUnbindFromProjectBadRequest creates a FlavorsUnbindFromProjectBadRequest with default headers values
func NewFlavorsUnbindFromProjectBadRequest() *FlavorsUnbindFromProjectBadRequest {
	return &FlavorsUnbindFromProjectBadRequest{}
}

/*
FlavorsUnbindFromProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FlavorsUnbindFromProjectBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this flavors unbind from project bad request response has a 2xx status code
func (o *FlavorsUnbindFromProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors unbind from project bad request response has a 3xx status code
func (o *FlavorsUnbindFromProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors unbind from project bad request response has a 4xx status code
func (o *FlavorsUnbindFromProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors unbind from project bad request response has a 5xx status code
func (o *FlavorsUnbindFromProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors unbind from project bad request response a status code equal to that given
func (o *FlavorsUnbindFromProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *FlavorsUnbindFromProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsUnbindFromProjectBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsUnbindFromProjectBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *FlavorsUnbindFromProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsUnbindFromProjectUnauthorized creates a FlavorsUnbindFromProjectUnauthorized with default headers values
func NewFlavorsUnbindFromProjectUnauthorized() *FlavorsUnbindFromProjectUnauthorized {
	return &FlavorsUnbindFromProjectUnauthorized{}
}

/*
FlavorsUnbindFromProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FlavorsUnbindFromProjectUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors unbind from project unauthorized response has a 2xx status code
func (o *FlavorsUnbindFromProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors unbind from project unauthorized response has a 3xx status code
func (o *FlavorsUnbindFromProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors unbind from project unauthorized response has a 4xx status code
func (o *FlavorsUnbindFromProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors unbind from project unauthorized response has a 5xx status code
func (o *FlavorsUnbindFromProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors unbind from project unauthorized response a status code equal to that given
func (o *FlavorsUnbindFromProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *FlavorsUnbindFromProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsUnbindFromProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsUnbindFromProjectUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsUnbindFromProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsUnbindFromProjectForbidden creates a FlavorsUnbindFromProjectForbidden with default headers values
func NewFlavorsUnbindFromProjectForbidden() *FlavorsUnbindFromProjectForbidden {
	return &FlavorsUnbindFromProjectForbidden{}
}

/*
FlavorsUnbindFromProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FlavorsUnbindFromProjectForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors unbind from project forbidden response has a 2xx status code
func (o *FlavorsUnbindFromProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors unbind from project forbidden response has a 3xx status code
func (o *FlavorsUnbindFromProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors unbind from project forbidden response has a 4xx status code
func (o *FlavorsUnbindFromProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors unbind from project forbidden response has a 5xx status code
func (o *FlavorsUnbindFromProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors unbind from project forbidden response a status code equal to that given
func (o *FlavorsUnbindFromProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FlavorsUnbindFromProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsUnbindFromProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsUnbindFromProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsUnbindFromProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsUnbindFromProjectNotFound creates a FlavorsUnbindFromProjectNotFound with default headers values
func NewFlavorsUnbindFromProjectNotFound() *FlavorsUnbindFromProjectNotFound {
	return &FlavorsUnbindFromProjectNotFound{}
}

/*
FlavorsUnbindFromProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FlavorsUnbindFromProjectNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors unbind from project not found response has a 2xx status code
func (o *FlavorsUnbindFromProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors unbind from project not found response has a 3xx status code
func (o *FlavorsUnbindFromProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors unbind from project not found response has a 4xx status code
func (o *FlavorsUnbindFromProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors unbind from project not found response has a 5xx status code
func (o *FlavorsUnbindFromProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors unbind from project not found response a status code equal to that given
func (o *FlavorsUnbindFromProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *FlavorsUnbindFromProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsUnbindFromProjectNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsUnbindFromProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsUnbindFromProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsUnbindFromProjectInternalServerError creates a FlavorsUnbindFromProjectInternalServerError with default headers values
func NewFlavorsUnbindFromProjectInternalServerError() *FlavorsUnbindFromProjectInternalServerError {
	return &FlavorsUnbindFromProjectInternalServerError{}
}

/*
FlavorsUnbindFromProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type FlavorsUnbindFromProjectInternalServerError struct {
}

// IsSuccess returns true when this flavors unbind from project internal server error response has a 2xx status code
func (o *FlavorsUnbindFromProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors unbind from project internal server error response has a 3xx status code
func (o *FlavorsUnbindFromProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors unbind from project internal server error response has a 4xx status code
func (o *FlavorsUnbindFromProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors unbind from project internal server error response has a 5xx status code
func (o *FlavorsUnbindFromProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this flavors unbind from project internal server error response a status code equal to that given
func (o *FlavorsUnbindFromProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *FlavorsUnbindFromProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectInternalServerError ", 500)
}

func (o *FlavorsUnbindFromProjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Flavors/unbind][%d] flavorsUnbindFromProjectInternalServerError ", 500)
}

func (o *FlavorsUnbindFromProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
