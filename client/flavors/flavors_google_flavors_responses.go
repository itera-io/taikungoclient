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

// FlavorsGoogleFlavorsReader is a Reader for the FlavorsGoogleFlavors structure.
type FlavorsGoogleFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlavorsGoogleFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlavorsGoogleFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFlavorsGoogleFlavorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFlavorsGoogleFlavorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFlavorsGoogleFlavorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFlavorsGoogleFlavorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFlavorsGoogleFlavorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFlavorsGoogleFlavorsOK creates a FlavorsGoogleFlavorsOK with default headers values
func NewFlavorsGoogleFlavorsOK() *FlavorsGoogleFlavorsOK {
	return &FlavorsGoogleFlavorsOK{}
}

/*
FlavorsGoogleFlavorsOK describes a response with status code 200, with default header values.

Success
*/
type FlavorsGoogleFlavorsOK struct {
	Payload *models.GoogleFlavorList
}

// IsSuccess returns true when this flavors google flavors o k response has a 2xx status code
func (o *FlavorsGoogleFlavorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flavors google flavors o k response has a 3xx status code
func (o *FlavorsGoogleFlavorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors google flavors o k response has a 4xx status code
func (o *FlavorsGoogleFlavorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors google flavors o k response has a 5xx status code
func (o *FlavorsGoogleFlavorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors google flavors o k response a status code equal to that given
func (o *FlavorsGoogleFlavorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the flavors google flavors o k response
func (o *FlavorsGoogleFlavorsOK) Code() int {
	return 200
}

func (o *FlavorsGoogleFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsGoogleFlavorsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsGoogleFlavorsOK) GetPayload() *models.GoogleFlavorList {
	return o.Payload
}

func (o *FlavorsGoogleFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoogleFlavorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGoogleFlavorsBadRequest creates a FlavorsGoogleFlavorsBadRequest with default headers values
func NewFlavorsGoogleFlavorsBadRequest() *FlavorsGoogleFlavorsBadRequest {
	return &FlavorsGoogleFlavorsBadRequest{}
}

/*
FlavorsGoogleFlavorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FlavorsGoogleFlavorsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors google flavors bad request response has a 2xx status code
func (o *FlavorsGoogleFlavorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors google flavors bad request response has a 3xx status code
func (o *FlavorsGoogleFlavorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors google flavors bad request response has a 4xx status code
func (o *FlavorsGoogleFlavorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors google flavors bad request response has a 5xx status code
func (o *FlavorsGoogleFlavorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors google flavors bad request response a status code equal to that given
func (o *FlavorsGoogleFlavorsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the flavors google flavors bad request response
func (o *FlavorsGoogleFlavorsBadRequest) Code() int {
	return 400
}

func (o *FlavorsGoogleFlavorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsGoogleFlavorsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsGoogleFlavorsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsGoogleFlavorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGoogleFlavorsUnauthorized creates a FlavorsGoogleFlavorsUnauthorized with default headers values
func NewFlavorsGoogleFlavorsUnauthorized() *FlavorsGoogleFlavorsUnauthorized {
	return &FlavorsGoogleFlavorsUnauthorized{}
}

/*
FlavorsGoogleFlavorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FlavorsGoogleFlavorsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors google flavors unauthorized response has a 2xx status code
func (o *FlavorsGoogleFlavorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors google flavors unauthorized response has a 3xx status code
func (o *FlavorsGoogleFlavorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors google flavors unauthorized response has a 4xx status code
func (o *FlavorsGoogleFlavorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors google flavors unauthorized response has a 5xx status code
func (o *FlavorsGoogleFlavorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors google flavors unauthorized response a status code equal to that given
func (o *FlavorsGoogleFlavorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the flavors google flavors unauthorized response
func (o *FlavorsGoogleFlavorsUnauthorized) Code() int {
	return 401
}

func (o *FlavorsGoogleFlavorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsGoogleFlavorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsGoogleFlavorsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsGoogleFlavorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGoogleFlavorsForbidden creates a FlavorsGoogleFlavorsForbidden with default headers values
func NewFlavorsGoogleFlavorsForbidden() *FlavorsGoogleFlavorsForbidden {
	return &FlavorsGoogleFlavorsForbidden{}
}

/*
FlavorsGoogleFlavorsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FlavorsGoogleFlavorsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors google flavors forbidden response has a 2xx status code
func (o *FlavorsGoogleFlavorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors google flavors forbidden response has a 3xx status code
func (o *FlavorsGoogleFlavorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors google flavors forbidden response has a 4xx status code
func (o *FlavorsGoogleFlavorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors google flavors forbidden response has a 5xx status code
func (o *FlavorsGoogleFlavorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors google flavors forbidden response a status code equal to that given
func (o *FlavorsGoogleFlavorsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the flavors google flavors forbidden response
func (o *FlavorsGoogleFlavorsForbidden) Code() int {
	return 403
}

func (o *FlavorsGoogleFlavorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsGoogleFlavorsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsGoogleFlavorsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsGoogleFlavorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGoogleFlavorsNotFound creates a FlavorsGoogleFlavorsNotFound with default headers values
func NewFlavorsGoogleFlavorsNotFound() *FlavorsGoogleFlavorsNotFound {
	return &FlavorsGoogleFlavorsNotFound{}
}

/*
FlavorsGoogleFlavorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FlavorsGoogleFlavorsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors google flavors not found response has a 2xx status code
func (o *FlavorsGoogleFlavorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors google flavors not found response has a 3xx status code
func (o *FlavorsGoogleFlavorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors google flavors not found response has a 4xx status code
func (o *FlavorsGoogleFlavorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors google flavors not found response has a 5xx status code
func (o *FlavorsGoogleFlavorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors google flavors not found response a status code equal to that given
func (o *FlavorsGoogleFlavorsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the flavors google flavors not found response
func (o *FlavorsGoogleFlavorsNotFound) Code() int {
	return 404
}

func (o *FlavorsGoogleFlavorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsGoogleFlavorsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsGoogleFlavorsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsGoogleFlavorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGoogleFlavorsInternalServerError creates a FlavorsGoogleFlavorsInternalServerError with default headers values
func NewFlavorsGoogleFlavorsInternalServerError() *FlavorsGoogleFlavorsInternalServerError {
	return &FlavorsGoogleFlavorsInternalServerError{}
}

/*
FlavorsGoogleFlavorsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type FlavorsGoogleFlavorsInternalServerError struct {
}

// IsSuccess returns true when this flavors google flavors internal server error response has a 2xx status code
func (o *FlavorsGoogleFlavorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors google flavors internal server error response has a 3xx status code
func (o *FlavorsGoogleFlavorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors google flavors internal server error response has a 4xx status code
func (o *FlavorsGoogleFlavorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors google flavors internal server error response has a 5xx status code
func (o *FlavorsGoogleFlavorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this flavors google flavors internal server error response a status code equal to that given
func (o *FlavorsGoogleFlavorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the flavors google flavors internal server error response
func (o *FlavorsGoogleFlavorsInternalServerError) Code() int {
	return 500
}

func (o *FlavorsGoogleFlavorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsInternalServerError ", 500)
}

func (o *FlavorsGoogleFlavorsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/google/{cloudId}][%d] flavorsGoogleFlavorsInternalServerError ", 500)
}

func (o *FlavorsGoogleFlavorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
