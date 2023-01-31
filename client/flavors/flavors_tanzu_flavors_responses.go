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

// FlavorsTanzuFlavorsReader is a Reader for the FlavorsTanzuFlavors structure.
type FlavorsTanzuFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlavorsTanzuFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlavorsTanzuFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFlavorsTanzuFlavorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFlavorsTanzuFlavorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFlavorsTanzuFlavorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFlavorsTanzuFlavorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFlavorsTanzuFlavorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFlavorsTanzuFlavorsOK creates a FlavorsTanzuFlavorsOK with default headers values
func NewFlavorsTanzuFlavorsOK() *FlavorsTanzuFlavorsOK {
	return &FlavorsTanzuFlavorsOK{}
}

/*
FlavorsTanzuFlavorsOK describes a response with status code 200, with default header values.

Success
*/
type FlavorsTanzuFlavorsOK struct {
	Payload *models.TanzuFlavorList
}

// IsSuccess returns true when this flavors tanzu flavors o k response has a 2xx status code
func (o *FlavorsTanzuFlavorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flavors tanzu flavors o k response has a 3xx status code
func (o *FlavorsTanzuFlavorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors tanzu flavors o k response has a 4xx status code
func (o *FlavorsTanzuFlavorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors tanzu flavors o k response has a 5xx status code
func (o *FlavorsTanzuFlavorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors tanzu flavors o k response a status code equal to that given
func (o *FlavorsTanzuFlavorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the flavors tanzu flavors o k response
func (o *FlavorsTanzuFlavorsOK) Code() int {
	return 200
}

func (o *FlavorsTanzuFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsTanzuFlavorsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsTanzuFlavorsOK) GetPayload() *models.TanzuFlavorList {
	return o.Payload
}

func (o *FlavorsTanzuFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TanzuFlavorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsTanzuFlavorsBadRequest creates a FlavorsTanzuFlavorsBadRequest with default headers values
func NewFlavorsTanzuFlavorsBadRequest() *FlavorsTanzuFlavorsBadRequest {
	return &FlavorsTanzuFlavorsBadRequest{}
}

/*
FlavorsTanzuFlavorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FlavorsTanzuFlavorsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors tanzu flavors bad request response has a 2xx status code
func (o *FlavorsTanzuFlavorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors tanzu flavors bad request response has a 3xx status code
func (o *FlavorsTanzuFlavorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors tanzu flavors bad request response has a 4xx status code
func (o *FlavorsTanzuFlavorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors tanzu flavors bad request response has a 5xx status code
func (o *FlavorsTanzuFlavorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors tanzu flavors bad request response a status code equal to that given
func (o *FlavorsTanzuFlavorsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the flavors tanzu flavors bad request response
func (o *FlavorsTanzuFlavorsBadRequest) Code() int {
	return 400
}

func (o *FlavorsTanzuFlavorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsTanzuFlavorsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsTanzuFlavorsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsTanzuFlavorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsTanzuFlavorsUnauthorized creates a FlavorsTanzuFlavorsUnauthorized with default headers values
func NewFlavorsTanzuFlavorsUnauthorized() *FlavorsTanzuFlavorsUnauthorized {
	return &FlavorsTanzuFlavorsUnauthorized{}
}

/*
FlavorsTanzuFlavorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FlavorsTanzuFlavorsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors tanzu flavors unauthorized response has a 2xx status code
func (o *FlavorsTanzuFlavorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors tanzu flavors unauthorized response has a 3xx status code
func (o *FlavorsTanzuFlavorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors tanzu flavors unauthorized response has a 4xx status code
func (o *FlavorsTanzuFlavorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors tanzu flavors unauthorized response has a 5xx status code
func (o *FlavorsTanzuFlavorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors tanzu flavors unauthorized response a status code equal to that given
func (o *FlavorsTanzuFlavorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the flavors tanzu flavors unauthorized response
func (o *FlavorsTanzuFlavorsUnauthorized) Code() int {
	return 401
}

func (o *FlavorsTanzuFlavorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsTanzuFlavorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsTanzuFlavorsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsTanzuFlavorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsTanzuFlavorsForbidden creates a FlavorsTanzuFlavorsForbidden with default headers values
func NewFlavorsTanzuFlavorsForbidden() *FlavorsTanzuFlavorsForbidden {
	return &FlavorsTanzuFlavorsForbidden{}
}

/*
FlavorsTanzuFlavorsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FlavorsTanzuFlavorsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors tanzu flavors forbidden response has a 2xx status code
func (o *FlavorsTanzuFlavorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors tanzu flavors forbidden response has a 3xx status code
func (o *FlavorsTanzuFlavorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors tanzu flavors forbidden response has a 4xx status code
func (o *FlavorsTanzuFlavorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors tanzu flavors forbidden response has a 5xx status code
func (o *FlavorsTanzuFlavorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors tanzu flavors forbidden response a status code equal to that given
func (o *FlavorsTanzuFlavorsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the flavors tanzu flavors forbidden response
func (o *FlavorsTanzuFlavorsForbidden) Code() int {
	return 403
}

func (o *FlavorsTanzuFlavorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsTanzuFlavorsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsTanzuFlavorsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsTanzuFlavorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsTanzuFlavorsNotFound creates a FlavorsTanzuFlavorsNotFound with default headers values
func NewFlavorsTanzuFlavorsNotFound() *FlavorsTanzuFlavorsNotFound {
	return &FlavorsTanzuFlavorsNotFound{}
}

/*
FlavorsTanzuFlavorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FlavorsTanzuFlavorsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this flavors tanzu flavors not found response has a 2xx status code
func (o *FlavorsTanzuFlavorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors tanzu flavors not found response has a 3xx status code
func (o *FlavorsTanzuFlavorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors tanzu flavors not found response has a 4xx status code
func (o *FlavorsTanzuFlavorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors tanzu flavors not found response has a 5xx status code
func (o *FlavorsTanzuFlavorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors tanzu flavors not found response a status code equal to that given
func (o *FlavorsTanzuFlavorsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the flavors tanzu flavors not found response
func (o *FlavorsTanzuFlavorsNotFound) Code() int {
	return 404
}

func (o *FlavorsTanzuFlavorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsTanzuFlavorsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsTanzuFlavorsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *FlavorsTanzuFlavorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsTanzuFlavorsInternalServerError creates a FlavorsTanzuFlavorsInternalServerError with default headers values
func NewFlavorsTanzuFlavorsInternalServerError() *FlavorsTanzuFlavorsInternalServerError {
	return &FlavorsTanzuFlavorsInternalServerError{}
}

/*
FlavorsTanzuFlavorsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type FlavorsTanzuFlavorsInternalServerError struct {
}

// IsSuccess returns true when this flavors tanzu flavors internal server error response has a 2xx status code
func (o *FlavorsTanzuFlavorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors tanzu flavors internal server error response has a 3xx status code
func (o *FlavorsTanzuFlavorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors tanzu flavors internal server error response has a 4xx status code
func (o *FlavorsTanzuFlavorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors tanzu flavors internal server error response has a 5xx status code
func (o *FlavorsTanzuFlavorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this flavors tanzu flavors internal server error response a status code equal to that given
func (o *FlavorsTanzuFlavorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the flavors tanzu flavors internal server error response
func (o *FlavorsTanzuFlavorsInternalServerError) Code() int {
	return 500
}

func (o *FlavorsTanzuFlavorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsInternalServerError ", 500)
}

func (o *FlavorsTanzuFlavorsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/tanzu/{cloudId}][%d] flavorsTanzuFlavorsInternalServerError ", 500)
}

func (o *FlavorsTanzuFlavorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
