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

// FlavorsGetSelectedFlavorsForProjectReader is a Reader for the FlavorsGetSelectedFlavorsForProject structure.
type FlavorsGetSelectedFlavorsForProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FlavorsGetSelectedFlavorsForProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFlavorsGetSelectedFlavorsForProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFlavorsGetSelectedFlavorsForProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFlavorsGetSelectedFlavorsForProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFlavorsGetSelectedFlavorsForProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFlavorsGetSelectedFlavorsForProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFlavorsGetSelectedFlavorsForProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFlavorsGetSelectedFlavorsForProjectOK creates a FlavorsGetSelectedFlavorsForProjectOK with default headers values
func NewFlavorsGetSelectedFlavorsForProjectOK() *FlavorsGetSelectedFlavorsForProjectOK {
	return &FlavorsGetSelectedFlavorsForProjectOK{}
}

/*
FlavorsGetSelectedFlavorsForProjectOK describes a response with status code 200, with default header values.

Success
*/
type FlavorsGetSelectedFlavorsForProjectOK struct {
	Payload *models.BoundFlavorsForProjectsList
}

// IsSuccess returns true when this flavors get selected flavors for project o k response has a 2xx status code
func (o *FlavorsGetSelectedFlavorsForProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this flavors get selected flavors for project o k response has a 3xx status code
func (o *FlavorsGetSelectedFlavorsForProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors get selected flavors for project o k response has a 4xx status code
func (o *FlavorsGetSelectedFlavorsForProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors get selected flavors for project o k response has a 5xx status code
func (o *FlavorsGetSelectedFlavorsForProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors get selected flavors for project o k response a status code equal to that given
func (o *FlavorsGetSelectedFlavorsForProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *FlavorsGetSelectedFlavorsForProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectOK  %+v", 200, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectOK  %+v", 200, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectOK) GetPayload() *models.BoundFlavorsForProjectsList {
	return o.Payload
}

func (o *FlavorsGetSelectedFlavorsForProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BoundFlavorsForProjectsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGetSelectedFlavorsForProjectBadRequest creates a FlavorsGetSelectedFlavorsForProjectBadRequest with default headers values
func NewFlavorsGetSelectedFlavorsForProjectBadRequest() *FlavorsGetSelectedFlavorsForProjectBadRequest {
	return &FlavorsGetSelectedFlavorsForProjectBadRequest{}
}

/*
FlavorsGetSelectedFlavorsForProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FlavorsGetSelectedFlavorsForProjectBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors get selected flavors for project bad request response has a 2xx status code
func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors get selected flavors for project bad request response has a 3xx status code
func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors get selected flavors for project bad request response has a 4xx status code
func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors get selected flavors for project bad request response has a 5xx status code
func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors get selected flavors for project bad request response a status code equal to that given
func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectBadRequest  %+v", 400, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsGetSelectedFlavorsForProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGetSelectedFlavorsForProjectUnauthorized creates a FlavorsGetSelectedFlavorsForProjectUnauthorized with default headers values
func NewFlavorsGetSelectedFlavorsForProjectUnauthorized() *FlavorsGetSelectedFlavorsForProjectUnauthorized {
	return &FlavorsGetSelectedFlavorsForProjectUnauthorized{}
}

/*
FlavorsGetSelectedFlavorsForProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FlavorsGetSelectedFlavorsForProjectUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors get selected flavors for project unauthorized response has a 2xx status code
func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors get selected flavors for project unauthorized response has a 3xx status code
func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors get selected flavors for project unauthorized response has a 4xx status code
func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors get selected flavors for project unauthorized response has a 5xx status code
func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors get selected flavors for project unauthorized response a status code equal to that given
func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsGetSelectedFlavorsForProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGetSelectedFlavorsForProjectForbidden creates a FlavorsGetSelectedFlavorsForProjectForbidden with default headers values
func NewFlavorsGetSelectedFlavorsForProjectForbidden() *FlavorsGetSelectedFlavorsForProjectForbidden {
	return &FlavorsGetSelectedFlavorsForProjectForbidden{}
}

/*
FlavorsGetSelectedFlavorsForProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FlavorsGetSelectedFlavorsForProjectForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors get selected flavors for project forbidden response has a 2xx status code
func (o *FlavorsGetSelectedFlavorsForProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors get selected flavors for project forbidden response has a 3xx status code
func (o *FlavorsGetSelectedFlavorsForProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors get selected flavors for project forbidden response has a 4xx status code
func (o *FlavorsGetSelectedFlavorsForProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors get selected flavors for project forbidden response has a 5xx status code
func (o *FlavorsGetSelectedFlavorsForProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors get selected flavors for project forbidden response a status code equal to that given
func (o *FlavorsGetSelectedFlavorsForProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FlavorsGetSelectedFlavorsForProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectForbidden  %+v", 403, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsGetSelectedFlavorsForProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGetSelectedFlavorsForProjectNotFound creates a FlavorsGetSelectedFlavorsForProjectNotFound with default headers values
func NewFlavorsGetSelectedFlavorsForProjectNotFound() *FlavorsGetSelectedFlavorsForProjectNotFound {
	return &FlavorsGetSelectedFlavorsForProjectNotFound{}
}

/*
FlavorsGetSelectedFlavorsForProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FlavorsGetSelectedFlavorsForProjectNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this flavors get selected flavors for project not found response has a 2xx status code
func (o *FlavorsGetSelectedFlavorsForProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors get selected flavors for project not found response has a 3xx status code
func (o *FlavorsGetSelectedFlavorsForProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors get selected flavors for project not found response has a 4xx status code
func (o *FlavorsGetSelectedFlavorsForProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this flavors get selected flavors for project not found response has a 5xx status code
func (o *FlavorsGetSelectedFlavorsForProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this flavors get selected flavors for project not found response a status code equal to that given
func (o *FlavorsGetSelectedFlavorsForProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *FlavorsGetSelectedFlavorsForProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectNotFound  %+v", 404, o.Payload)
}

func (o *FlavorsGetSelectedFlavorsForProjectNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *FlavorsGetSelectedFlavorsForProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFlavorsGetSelectedFlavorsForProjectInternalServerError creates a FlavorsGetSelectedFlavorsForProjectInternalServerError with default headers values
func NewFlavorsGetSelectedFlavorsForProjectInternalServerError() *FlavorsGetSelectedFlavorsForProjectInternalServerError {
	return &FlavorsGetSelectedFlavorsForProjectInternalServerError{}
}

/*
FlavorsGetSelectedFlavorsForProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type FlavorsGetSelectedFlavorsForProjectInternalServerError struct {
}

// IsSuccess returns true when this flavors get selected flavors for project internal server error response has a 2xx status code
func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this flavors get selected flavors for project internal server error response has a 3xx status code
func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this flavors get selected flavors for project internal server error response has a 4xx status code
func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this flavors get selected flavors for project internal server error response has a 5xx status code
func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this flavors get selected flavors for project internal server error response a status code equal to that given
func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectInternalServerError ", 500)
}

func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Flavors/projects/list][%d] flavorsGetSelectedFlavorsForProjectInternalServerError ", 500)
}

func (o *FlavorsGetSelectedFlavorsForProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
