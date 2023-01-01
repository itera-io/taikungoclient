// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AccessProfilesAccessProfilesForOrganizationListReader is a Reader for the AccessProfilesAccessProfilesForOrganizationList structure.
type AccessProfilesAccessProfilesForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccessProfilesAccessProfilesForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccessProfilesAccessProfilesForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAccessProfilesAccessProfilesForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAccessProfilesAccessProfilesForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAccessProfilesAccessProfilesForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAccessProfilesAccessProfilesForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAccessProfilesAccessProfilesForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccessProfilesAccessProfilesForOrganizationListOK creates a AccessProfilesAccessProfilesForOrganizationListOK with default headers values
func NewAccessProfilesAccessProfilesForOrganizationListOK() *AccessProfilesAccessProfilesForOrganizationListOK {
	return &AccessProfilesAccessProfilesForOrganizationListOK{}
}

/*
AccessProfilesAccessProfilesForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type AccessProfilesAccessProfilesForOrganizationListOK struct {
	Payload []*models.CommonDropdownDto
}

// IsSuccess returns true when this access profiles access profiles for organization list o k response has a 2xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this access profiles access profiles for organization list o k response has a 3xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles access profiles for organization list o k response has a 4xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles access profiles for organization list o k response has a 5xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles access profiles for organization list o k response a status code equal to that given
func (o *AccessProfilesAccessProfilesForOrganizationListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AccessProfilesAccessProfilesForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListOK) GetPayload() []*models.CommonDropdownDto {
	return o.Payload
}

func (o *AccessProfilesAccessProfilesForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesAccessProfilesForOrganizationListBadRequest creates a AccessProfilesAccessProfilesForOrganizationListBadRequest with default headers values
func NewAccessProfilesAccessProfilesForOrganizationListBadRequest() *AccessProfilesAccessProfilesForOrganizationListBadRequest {
	return &AccessProfilesAccessProfilesForOrganizationListBadRequest{}
}

/*
AccessProfilesAccessProfilesForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AccessProfilesAccessProfilesForOrganizationListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles access profiles for organization list bad request response has a 2xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles access profiles for organization list bad request response has a 3xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles access profiles for organization list bad request response has a 4xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles access profiles for organization list bad request response has a 5xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles access profiles for organization list bad request response a status code equal to that given
func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesAccessProfilesForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesAccessProfilesForOrganizationListUnauthorized creates a AccessProfilesAccessProfilesForOrganizationListUnauthorized with default headers values
func NewAccessProfilesAccessProfilesForOrganizationListUnauthorized() *AccessProfilesAccessProfilesForOrganizationListUnauthorized {
	return &AccessProfilesAccessProfilesForOrganizationListUnauthorized{}
}

/*
AccessProfilesAccessProfilesForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AccessProfilesAccessProfilesForOrganizationListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles access profiles for organization list unauthorized response has a 2xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles access profiles for organization list unauthorized response has a 3xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles access profiles for organization list unauthorized response has a 4xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles access profiles for organization list unauthorized response has a 5xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles access profiles for organization list unauthorized response a status code equal to that given
func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesAccessProfilesForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesAccessProfilesForOrganizationListForbidden creates a AccessProfilesAccessProfilesForOrganizationListForbidden with default headers values
func NewAccessProfilesAccessProfilesForOrganizationListForbidden() *AccessProfilesAccessProfilesForOrganizationListForbidden {
	return &AccessProfilesAccessProfilesForOrganizationListForbidden{}
}

/*
AccessProfilesAccessProfilesForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AccessProfilesAccessProfilesForOrganizationListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles access profiles for organization list forbidden response has a 2xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles access profiles for organization list forbidden response has a 3xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles access profiles for organization list forbidden response has a 4xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles access profiles for organization list forbidden response has a 5xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles access profiles for organization list forbidden response a status code equal to that given
func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesAccessProfilesForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesAccessProfilesForOrganizationListNotFound creates a AccessProfilesAccessProfilesForOrganizationListNotFound with default headers values
func NewAccessProfilesAccessProfilesForOrganizationListNotFound() *AccessProfilesAccessProfilesForOrganizationListNotFound {
	return &AccessProfilesAccessProfilesForOrganizationListNotFound{}
}

/*
AccessProfilesAccessProfilesForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AccessProfilesAccessProfilesForOrganizationListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this access profiles access profiles for organization list not found response has a 2xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles access profiles for organization list not found response has a 3xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles access profiles for organization list not found response has a 4xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this access profiles access profiles for organization list not found response has a 5xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this access profiles access profiles for organization list not found response a status code equal to that given
func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AccessProfilesAccessProfilesForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccessProfilesAccessProfilesForOrganizationListInternalServerError creates a AccessProfilesAccessProfilesForOrganizationListInternalServerError with default headers values
func NewAccessProfilesAccessProfilesForOrganizationListInternalServerError() *AccessProfilesAccessProfilesForOrganizationListInternalServerError {
	return &AccessProfilesAccessProfilesForOrganizationListInternalServerError{}
}

/*
AccessProfilesAccessProfilesForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AccessProfilesAccessProfilesForOrganizationListInternalServerError struct {
}

// IsSuccess returns true when this access profiles access profiles for organization list internal server error response has a 2xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this access profiles access profiles for organization list internal server error response has a 3xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this access profiles access profiles for organization list internal server error response has a 4xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this access profiles access profiles for organization list internal server error response has a 5xx status code
func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this access profiles access profiles for organization list internal server error response a status code equal to that given
func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListInternalServerError ", 500)
}

func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AccessProfiles/list][%d] accessProfilesAccessProfilesForOrganizationListInternalServerError ", 500)
}

func (o *AccessProfilesAccessProfilesForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
