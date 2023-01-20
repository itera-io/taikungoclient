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

// FlavorsAwsFlavorsReader is a Reader for the FlavorsAwsFlavors structure.
type FlavorsAwsFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlavorsAwsFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlavorsAwsFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFlavorsAwsFlavorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFlavorsAwsFlavorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFlavorsAwsFlavorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFlavorsAwsFlavorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFlavorsAwsFlavorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFlavorsAwsFlavorsOK creates a FlavorsAwsFlavorsOK with default headers values
func NewFlavorsAwsFlavorsOK() *FlavorsAwsFlavorsOK {
	return &FlavorsAwsFlavorsOK{}
}

/*
FlavorsAwsFlavorsOK describes a response with status code 200, with default header values.

Success
*/
type FlavorsAwsFlavorsOK struct {
	Payload *models.AwsFlavorList
}

// IsSuccess returns true when this flavors aws flavors o k response has a 2xx status code
func (o *FlavorsAwsFlavorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flavors aws flavors o k response has a 3xx status code
func (o *FlavorsAwsFlavorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors aws flavors o k response has a 4xx status code
func (o *FlavorsAwsFlavorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors aws flavors o k response has a 5xx status code
func (o *FlavorsAwsFlavorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors aws flavors o k response a status code equal to that given
func (o *FlavorsAwsFlavorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the flavors aws flavors o k response
func (o *FlavorsAwsFlavorsOK) Code() int {
	return 200
}

func (o *FlavorsAwsFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsAwsFlavorsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsAwsFlavorsOK) GetPayload() *models.AwsFlavorList {
	return o.Payload
}

func (o *FlavorsAwsFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsFlavorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsAwsFlavorsBadRequest creates a FlavorsAwsFlavorsBadRequest with default headers values
func NewFlavorsAwsFlavorsBadRequest() *FlavorsAwsFlavorsBadRequest {
	return &FlavorsAwsFlavorsBadRequest{}
}

/*
FlavorsAwsFlavorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FlavorsAwsFlavorsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors aws flavors bad request response has a 2xx status code
func (o *FlavorsAwsFlavorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors aws flavors bad request response has a 3xx status code
func (o *FlavorsAwsFlavorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors aws flavors bad request response has a 4xx status code
func (o *FlavorsAwsFlavorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors aws flavors bad request response has a 5xx status code
func (o *FlavorsAwsFlavorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors aws flavors bad request response a status code equal to that given
func (o *FlavorsAwsFlavorsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the flavors aws flavors bad request response
func (o *FlavorsAwsFlavorsBadRequest) Code() int {
	return 400
}

func (o *FlavorsAwsFlavorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsAwsFlavorsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsAwsFlavorsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsAwsFlavorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsAwsFlavorsUnauthorized creates a FlavorsAwsFlavorsUnauthorized with default headers values
func NewFlavorsAwsFlavorsUnauthorized() *FlavorsAwsFlavorsUnauthorized {
	return &FlavorsAwsFlavorsUnauthorized{}
}

/*
FlavorsAwsFlavorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FlavorsAwsFlavorsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors aws flavors unauthorized response has a 2xx status code
func (o *FlavorsAwsFlavorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors aws flavors unauthorized response has a 3xx status code
func (o *FlavorsAwsFlavorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors aws flavors unauthorized response has a 4xx status code
func (o *FlavorsAwsFlavorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors aws flavors unauthorized response has a 5xx status code
func (o *FlavorsAwsFlavorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors aws flavors unauthorized response a status code equal to that given
func (o *FlavorsAwsFlavorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the flavors aws flavors unauthorized response
func (o *FlavorsAwsFlavorsUnauthorized) Code() int {
	return 401
}

func (o *FlavorsAwsFlavorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsAwsFlavorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsAwsFlavorsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsAwsFlavorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsAwsFlavorsForbidden creates a FlavorsAwsFlavorsForbidden with default headers values
func NewFlavorsAwsFlavorsForbidden() *FlavorsAwsFlavorsForbidden {
	return &FlavorsAwsFlavorsForbidden{}
}

/*
FlavorsAwsFlavorsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FlavorsAwsFlavorsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors aws flavors forbidden response has a 2xx status code
func (o *FlavorsAwsFlavorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors aws flavors forbidden response has a 3xx status code
func (o *FlavorsAwsFlavorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors aws flavors forbidden response has a 4xx status code
func (o *FlavorsAwsFlavorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors aws flavors forbidden response has a 5xx status code
func (o *FlavorsAwsFlavorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors aws flavors forbidden response a status code equal to that given
func (o *FlavorsAwsFlavorsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the flavors aws flavors forbidden response
func (o *FlavorsAwsFlavorsForbidden) Code() int {
	return 403
}

func (o *FlavorsAwsFlavorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsAwsFlavorsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsAwsFlavorsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsAwsFlavorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsAwsFlavorsNotFound creates a FlavorsAwsFlavorsNotFound with default headers values
func NewFlavorsAwsFlavorsNotFound() *FlavorsAwsFlavorsNotFound {
	return &FlavorsAwsFlavorsNotFound{}
}

/*
FlavorsAwsFlavorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FlavorsAwsFlavorsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors aws flavors not found response has a 2xx status code
func (o *FlavorsAwsFlavorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors aws flavors not found response has a 3xx status code
func (o *FlavorsAwsFlavorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors aws flavors not found response has a 4xx status code
func (o *FlavorsAwsFlavorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors aws flavors not found response has a 5xx status code
func (o *FlavorsAwsFlavorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors aws flavors not found response a status code equal to that given
func (o *FlavorsAwsFlavorsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the flavors aws flavors not found response
func (o *FlavorsAwsFlavorsNotFound) Code() int {
	return 404
}

func (o *FlavorsAwsFlavorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsAwsFlavorsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsAwsFlavorsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsAwsFlavorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsAwsFlavorsInternalServerError creates a FlavorsAwsFlavorsInternalServerError with default headers values
func NewFlavorsAwsFlavorsInternalServerError() *FlavorsAwsFlavorsInternalServerError {
	return &FlavorsAwsFlavorsInternalServerError{}
}

/*
FlavorsAwsFlavorsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type FlavorsAwsFlavorsInternalServerError struct {
}

// IsSuccess returns true when this flavors aws flavors internal server error response has a 2xx status code
func (o *FlavorsAwsFlavorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors aws flavors internal server error response has a 3xx status code
func (o *FlavorsAwsFlavorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors aws flavors internal server error response has a 4xx status code
func (o *FlavorsAwsFlavorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors aws flavors internal server error response has a 5xx status code
func (o *FlavorsAwsFlavorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this flavors aws flavors internal server error response a status code equal to that given
func (o *FlavorsAwsFlavorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the flavors aws flavors internal server error response
func (o *FlavorsAwsFlavorsInternalServerError) Code() int {
	return 500
}

func (o *FlavorsAwsFlavorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsInternalServerError ", 500)
}

func (o *FlavorsAwsFlavorsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/aws/{cloudId}][%d] flavorsAwsFlavorsInternalServerError ", 500)
}

func (o *FlavorsAwsFlavorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
