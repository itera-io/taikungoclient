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

// FlavorsOpenstackFlavorsReader is a Reader for the FlavorsOpenstackFlavors structure.
type FlavorsOpenstackFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlavorsOpenstackFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlavorsOpenstackFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFlavorsOpenstackFlavorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFlavorsOpenstackFlavorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFlavorsOpenstackFlavorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFlavorsOpenstackFlavorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFlavorsOpenstackFlavorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFlavorsOpenstackFlavorsOK creates a FlavorsOpenstackFlavorsOK with default headers values
func NewFlavorsOpenstackFlavorsOK() *FlavorsOpenstackFlavorsOK {
	return &FlavorsOpenstackFlavorsOK{}
}

/*
FlavorsOpenstackFlavorsOK describes a response with status code 200, with default header values.

Success
*/
type FlavorsOpenstackFlavorsOK struct {
	Payload *models.OpenstackFlavorList
}

// IsSuccess returns true when this flavors openstack flavors o k response has a 2xx status code
func (o *FlavorsOpenstackFlavorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flavors openstack flavors o k response has a 3xx status code
func (o *FlavorsOpenstackFlavorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors openstack flavors o k response has a 4xx status code
func (o *FlavorsOpenstackFlavorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors openstack flavors o k response has a 5xx status code
func (o *FlavorsOpenstackFlavorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors openstack flavors o k response a status code equal to that given
func (o *FlavorsOpenstackFlavorsOK) IsCode(code int) bool {
	return code == 200
}

func (o *FlavorsOpenstackFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsOpenstackFlavorsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsOK  %+v", 200, o.Payload)
}

func (o *FlavorsOpenstackFlavorsOK) GetPayload() *models.OpenstackFlavorList {
	return o.Payload
}

func (o *FlavorsOpenstackFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenstackFlavorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsOpenstackFlavorsBadRequest creates a FlavorsOpenstackFlavorsBadRequest with default headers values
func NewFlavorsOpenstackFlavorsBadRequest() *FlavorsOpenstackFlavorsBadRequest {
	return &FlavorsOpenstackFlavorsBadRequest{}
}

/*
FlavorsOpenstackFlavorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FlavorsOpenstackFlavorsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this flavors openstack flavors bad request response has a 2xx status code
func (o *FlavorsOpenstackFlavorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors openstack flavors bad request response has a 3xx status code
func (o *FlavorsOpenstackFlavorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors openstack flavors bad request response has a 4xx status code
func (o *FlavorsOpenstackFlavorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors openstack flavors bad request response has a 5xx status code
func (o *FlavorsOpenstackFlavorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors openstack flavors bad request response a status code equal to that given
func (o *FlavorsOpenstackFlavorsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *FlavorsOpenstackFlavorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsOpenstackFlavorsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsOpenstackFlavorsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *FlavorsOpenstackFlavorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsOpenstackFlavorsUnauthorized creates a FlavorsOpenstackFlavorsUnauthorized with default headers values
func NewFlavorsOpenstackFlavorsUnauthorized() *FlavorsOpenstackFlavorsUnauthorized {
	return &FlavorsOpenstackFlavorsUnauthorized{}
}

/*
FlavorsOpenstackFlavorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FlavorsOpenstackFlavorsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors openstack flavors unauthorized response has a 2xx status code
func (o *FlavorsOpenstackFlavorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors openstack flavors unauthorized response has a 3xx status code
func (o *FlavorsOpenstackFlavorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors openstack flavors unauthorized response has a 4xx status code
func (o *FlavorsOpenstackFlavorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors openstack flavors unauthorized response has a 5xx status code
func (o *FlavorsOpenstackFlavorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors openstack flavors unauthorized response a status code equal to that given
func (o *FlavorsOpenstackFlavorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *FlavorsOpenstackFlavorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsOpenstackFlavorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsOpenstackFlavorsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsOpenstackFlavorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsOpenstackFlavorsForbidden creates a FlavorsOpenstackFlavorsForbidden with default headers values
func NewFlavorsOpenstackFlavorsForbidden() *FlavorsOpenstackFlavorsForbidden {
	return &FlavorsOpenstackFlavorsForbidden{}
}

/*
FlavorsOpenstackFlavorsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FlavorsOpenstackFlavorsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors openstack flavors forbidden response has a 2xx status code
func (o *FlavorsOpenstackFlavorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors openstack flavors forbidden response has a 3xx status code
func (o *FlavorsOpenstackFlavorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors openstack flavors forbidden response has a 4xx status code
func (o *FlavorsOpenstackFlavorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors openstack flavors forbidden response has a 5xx status code
func (o *FlavorsOpenstackFlavorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors openstack flavors forbidden response a status code equal to that given
func (o *FlavorsOpenstackFlavorsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FlavorsOpenstackFlavorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsOpenstackFlavorsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsOpenstackFlavorsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsOpenstackFlavorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsOpenstackFlavorsNotFound creates a FlavorsOpenstackFlavorsNotFound with default headers values
func NewFlavorsOpenstackFlavorsNotFound() *FlavorsOpenstackFlavorsNotFound {
	return &FlavorsOpenstackFlavorsNotFound{}
}

/*
FlavorsOpenstackFlavorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FlavorsOpenstackFlavorsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors openstack flavors not found response has a 2xx status code
func (o *FlavorsOpenstackFlavorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors openstack flavors not found response has a 3xx status code
func (o *FlavorsOpenstackFlavorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors openstack flavors not found response has a 4xx status code
func (o *FlavorsOpenstackFlavorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors openstack flavors not found response has a 5xx status code
func (o *FlavorsOpenstackFlavorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors openstack flavors not found response a status code equal to that given
func (o *FlavorsOpenstackFlavorsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *FlavorsOpenstackFlavorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsOpenstackFlavorsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsOpenstackFlavorsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsOpenstackFlavorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsOpenstackFlavorsInternalServerError creates a FlavorsOpenstackFlavorsInternalServerError with default headers values
func NewFlavorsOpenstackFlavorsInternalServerError() *FlavorsOpenstackFlavorsInternalServerError {
	return &FlavorsOpenstackFlavorsInternalServerError{}
}

/*
FlavorsOpenstackFlavorsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type FlavorsOpenstackFlavorsInternalServerError struct {
}

// IsSuccess returns true when this flavors openstack flavors internal server error response has a 2xx status code
func (o *FlavorsOpenstackFlavorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors openstack flavors internal server error response has a 3xx status code
func (o *FlavorsOpenstackFlavorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors openstack flavors internal server error response has a 4xx status code
func (o *FlavorsOpenstackFlavorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors openstack flavors internal server error response has a 5xx status code
func (o *FlavorsOpenstackFlavorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this flavors openstack flavors internal server error response a status code equal to that given
func (o *FlavorsOpenstackFlavorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *FlavorsOpenstackFlavorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsInternalServerError ", 500)
}

func (o *FlavorsOpenstackFlavorsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/openstack/{cloudId}][%d] flavorsOpenstackFlavorsInternalServerError ", 500)
}

func (o *FlavorsOpenstackFlavorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
