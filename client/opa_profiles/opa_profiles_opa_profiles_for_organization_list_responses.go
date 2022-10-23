// Code generated by go-swagger; DO NOT EDIT.

package opa_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpaProfilesOpaProfilesForOrganizationListReader is a Reader for the OpaProfilesOpaProfilesForOrganizationList structure.
type OpaProfilesOpaProfilesForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpaProfilesOpaProfilesForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpaProfilesOpaProfilesForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpaProfilesOpaProfilesForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpaProfilesOpaProfilesForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpaProfilesOpaProfilesForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpaProfilesOpaProfilesForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpaProfilesOpaProfilesForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpaProfilesOpaProfilesForOrganizationListOK creates a OpaProfilesOpaProfilesForOrganizationListOK with default headers values
func NewOpaProfilesOpaProfilesForOrganizationListOK() *OpaProfilesOpaProfilesForOrganizationListOK {
	return &OpaProfilesOpaProfilesForOrganizationListOK{}
}

/*
OpaProfilesOpaProfilesForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type OpaProfilesOpaProfilesForOrganizationListOK struct {
	Payload []*models.CommonExtendedDropdownDto
}

// IsSuccess returns true when this opa profiles opa profiles for organization list o k response has a 2xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this opa profiles opa profiles for organization list o k response has a 3xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles opa profiles for organization list o k response has a 4xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles opa profiles for organization list o k response has a 5xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles opa profiles for organization list o k response a status code equal to that given
func (o *OpaProfilesOpaProfilesForOrganizationListOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpaProfilesOpaProfilesForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListOK) GetPayload() []*models.CommonExtendedDropdownDto {
	return o.Payload
}

func (o *OpaProfilesOpaProfilesForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesOpaProfilesForOrganizationListBadRequest creates a OpaProfilesOpaProfilesForOrganizationListBadRequest with default headers values
func NewOpaProfilesOpaProfilesForOrganizationListBadRequest() *OpaProfilesOpaProfilesForOrganizationListBadRequest {
	return &OpaProfilesOpaProfilesForOrganizationListBadRequest{}
}

/*
OpaProfilesOpaProfilesForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpaProfilesOpaProfilesForOrganizationListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this opa profiles opa profiles for organization list bad request response has a 2xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles opa profiles for organization list bad request response has a 3xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles opa profiles for organization list bad request response has a 4xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles opa profiles for organization list bad request response has a 5xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles opa profiles for organization list bad request response a status code equal to that given
func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *OpaProfilesOpaProfilesForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesOpaProfilesForOrganizationListUnauthorized creates a OpaProfilesOpaProfilesForOrganizationListUnauthorized with default headers values
func NewOpaProfilesOpaProfilesForOrganizationListUnauthorized() *OpaProfilesOpaProfilesForOrganizationListUnauthorized {
	return &OpaProfilesOpaProfilesForOrganizationListUnauthorized{}
}

/*
OpaProfilesOpaProfilesForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpaProfilesOpaProfilesForOrganizationListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this opa profiles opa profiles for organization list unauthorized response has a 2xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles opa profiles for organization list unauthorized response has a 3xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles opa profiles for organization list unauthorized response has a 4xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles opa profiles for organization list unauthorized response has a 5xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles opa profiles for organization list unauthorized response a status code equal to that given
func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpaProfilesOpaProfilesForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesOpaProfilesForOrganizationListForbidden creates a OpaProfilesOpaProfilesForOrganizationListForbidden with default headers values
func NewOpaProfilesOpaProfilesForOrganizationListForbidden() *OpaProfilesOpaProfilesForOrganizationListForbidden {
	return &OpaProfilesOpaProfilesForOrganizationListForbidden{}
}

/*
OpaProfilesOpaProfilesForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpaProfilesOpaProfilesForOrganizationListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this opa profiles opa profiles for organization list forbidden response has a 2xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles opa profiles for organization list forbidden response has a 3xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles opa profiles for organization list forbidden response has a 4xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles opa profiles for organization list forbidden response has a 5xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles opa profiles for organization list forbidden response a status code equal to that given
func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpaProfilesOpaProfilesForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesOpaProfilesForOrganizationListNotFound creates a OpaProfilesOpaProfilesForOrganizationListNotFound with default headers values
func NewOpaProfilesOpaProfilesForOrganizationListNotFound() *OpaProfilesOpaProfilesForOrganizationListNotFound {
	return &OpaProfilesOpaProfilesForOrganizationListNotFound{}
}

/*
OpaProfilesOpaProfilesForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpaProfilesOpaProfilesForOrganizationListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this opa profiles opa profiles for organization list not found response has a 2xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles opa profiles for organization list not found response has a 3xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles opa profiles for organization list not found response has a 4xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this opa profiles opa profiles for organization list not found response has a 5xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this opa profiles opa profiles for organization list not found response a status code equal to that given
func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OpaProfilesOpaProfilesForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpaProfilesOpaProfilesForOrganizationListInternalServerError creates a OpaProfilesOpaProfilesForOrganizationListInternalServerError with default headers values
func NewOpaProfilesOpaProfilesForOrganizationListInternalServerError() *OpaProfilesOpaProfilesForOrganizationListInternalServerError {
	return &OpaProfilesOpaProfilesForOrganizationListInternalServerError{}
}

/*
OpaProfilesOpaProfilesForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpaProfilesOpaProfilesForOrganizationListInternalServerError struct {
}

// IsSuccess returns true when this opa profiles opa profiles for organization list internal server error response has a 2xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this opa profiles opa profiles for organization list internal server error response has a 3xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this opa profiles opa profiles for organization list internal server error response has a 4xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this opa profiles opa profiles for organization list internal server error response has a 5xx status code
func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this opa profiles opa profiles for organization list internal server error response a status code equal to that given
func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListInternalServerError ", 500)
}

func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/OpaProfiles/list][%d] opaProfilesOpaProfilesForOrganizationListInternalServerError ", 500)
}

func (o *OpaProfilesOpaProfilesForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
