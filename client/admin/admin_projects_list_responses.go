// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminProjectsListReader is a Reader for the AdminProjectsList structure.
type AdminProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminProjectsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminProjectsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminProjectsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminProjectsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminProjectsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminProjectsListOK creates a AdminProjectsListOK with default headers values
func NewAdminProjectsListOK() *AdminProjectsListOK {
	return &AdminProjectsListOK{}
}

/*
AdminProjectsListOK describes a response with status code 200, with default header values.

Success
*/
type AdminProjectsListOK struct {
	Payload *models.AdminProjectsList
}

// IsSuccess returns true when this admin projects list o k response has a 2xx status code
func (o *AdminProjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin projects list o k response has a 3xx status code
func (o *AdminProjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin projects list o k response has a 4xx status code
func (o *AdminProjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin projects list o k response has a 5xx status code
func (o *AdminProjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin projects list o k response a status code equal to that given
func (o *AdminProjectsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin projects list o k response
func (o *AdminProjectsListOK) Code() int {
	return 200
}

func (o *AdminProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListOK  %+v", 200, o.Payload)
}

func (o *AdminProjectsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListOK  %+v", 200, o.Payload)
}

func (o *AdminProjectsListOK) GetPayload() *models.AdminProjectsList {
	return o.Payload
}

func (o *AdminProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminProjectsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProjectsListBadRequest creates a AdminProjectsListBadRequest with default headers values
func NewAdminProjectsListBadRequest() *AdminProjectsListBadRequest {
	return &AdminProjectsListBadRequest{}
}

/*
AdminProjectsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminProjectsListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin projects list bad request response has a 2xx status code
func (o *AdminProjectsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin projects list bad request response has a 3xx status code
func (o *AdminProjectsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin projects list bad request response has a 4xx status code
func (o *AdminProjectsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin projects list bad request response has a 5xx status code
func (o *AdminProjectsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin projects list bad request response a status code equal to that given
func (o *AdminProjectsListBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin projects list bad request response
func (o *AdminProjectsListBadRequest) Code() int {
	return 400
}

func (o *AdminProjectsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListBadRequest  %+v", 400, o.Payload)
}

func (o *AdminProjectsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListBadRequest  %+v", 400, o.Payload)
}

func (o *AdminProjectsListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminProjectsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProjectsListUnauthorized creates a AdminProjectsListUnauthorized with default headers values
func NewAdminProjectsListUnauthorized() *AdminProjectsListUnauthorized {
	return &AdminProjectsListUnauthorized{}
}

/*
AdminProjectsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminProjectsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin projects list unauthorized response has a 2xx status code
func (o *AdminProjectsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin projects list unauthorized response has a 3xx status code
func (o *AdminProjectsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin projects list unauthorized response has a 4xx status code
func (o *AdminProjectsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin projects list unauthorized response has a 5xx status code
func (o *AdminProjectsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin projects list unauthorized response a status code equal to that given
func (o *AdminProjectsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin projects list unauthorized response
func (o *AdminProjectsListUnauthorized) Code() int {
	return 401
}

func (o *AdminProjectsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminProjectsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminProjectsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminProjectsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProjectsListForbidden creates a AdminProjectsListForbidden with default headers values
func NewAdminProjectsListForbidden() *AdminProjectsListForbidden {
	return &AdminProjectsListForbidden{}
}

/*
AdminProjectsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminProjectsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin projects list forbidden response has a 2xx status code
func (o *AdminProjectsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin projects list forbidden response has a 3xx status code
func (o *AdminProjectsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin projects list forbidden response has a 4xx status code
func (o *AdminProjectsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin projects list forbidden response has a 5xx status code
func (o *AdminProjectsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin projects list forbidden response a status code equal to that given
func (o *AdminProjectsListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin projects list forbidden response
func (o *AdminProjectsListForbidden) Code() int {
	return 403
}

func (o *AdminProjectsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListForbidden  %+v", 403, o.Payload)
}

func (o *AdminProjectsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListForbidden  %+v", 403, o.Payload)
}

func (o *AdminProjectsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminProjectsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProjectsListNotFound creates a AdminProjectsListNotFound with default headers values
func NewAdminProjectsListNotFound() *AdminProjectsListNotFound {
	return &AdminProjectsListNotFound{}
}

/*
AdminProjectsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminProjectsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin projects list not found response has a 2xx status code
func (o *AdminProjectsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin projects list not found response has a 3xx status code
func (o *AdminProjectsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin projects list not found response has a 4xx status code
func (o *AdminProjectsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin projects list not found response has a 5xx status code
func (o *AdminProjectsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin projects list not found response a status code equal to that given
func (o *AdminProjectsListNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the admin projects list not found response
func (o *AdminProjectsListNotFound) Code() int {
	return 404
}

func (o *AdminProjectsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListNotFound  %+v", 404, o.Payload)
}

func (o *AdminProjectsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListNotFound  %+v", 404, o.Payload)
}

func (o *AdminProjectsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminProjectsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminProjectsListInternalServerError creates a AdminProjectsListInternalServerError with default headers values
func NewAdminProjectsListInternalServerError() *AdminProjectsListInternalServerError {
	return &AdminProjectsListInternalServerError{}
}

/*
AdminProjectsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminProjectsListInternalServerError struct {
}

// IsSuccess returns true when this admin projects list internal server error response has a 2xx status code
func (o *AdminProjectsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin projects list internal server error response has a 3xx status code
func (o *AdminProjectsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin projects list internal server error response has a 4xx status code
func (o *AdminProjectsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin projects list internal server error response has a 5xx status code
func (o *AdminProjectsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin projects list internal server error response a status code equal to that given
func (o *AdminProjectsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin projects list internal server error response
func (o *AdminProjectsListInternalServerError) Code() int {
	return 500
}

func (o *AdminProjectsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListInternalServerError ", 500)
}

func (o *AdminProjectsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/projects/list][%d] adminProjectsListInternalServerError ", 500)
}

func (o *AdminProjectsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
