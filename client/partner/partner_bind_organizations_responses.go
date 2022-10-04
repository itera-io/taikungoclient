// Code generated by go-swagger; DO NOT EDIT.

package partner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// PartnerBindOrganizationsReader is a Reader for the PartnerBindOrganizations structure.
type PartnerBindOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartnerBindOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartnerBindOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartnerBindOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartnerBindOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPartnerBindOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartnerBindOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPartnerBindOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartnerBindOrganizationsOK creates a PartnerBindOrganizationsOK with default headers values
func NewPartnerBindOrganizationsOK() *PartnerBindOrganizationsOK {
	return &PartnerBindOrganizationsOK{}
}

/*
PartnerBindOrganizationsOK describes a response with status code 200, with default header values.

Success
*/
type PartnerBindOrganizationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this partner bind organizations o k response has a 2xx status code
func (o *PartnerBindOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this partner bind organizations o k response has a 3xx status code
func (o *PartnerBindOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner bind organizations o k response has a 4xx status code
func (o *PartnerBindOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner bind organizations o k response has a 5xx status code
func (o *PartnerBindOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this partner bind organizations o k response a status code equal to that given
func (o *PartnerBindOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PartnerBindOrganizationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsOK  %+v", 200, o.Payload)
}

func (o *PartnerBindOrganizationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsOK  %+v", 200, o.Payload)
}

func (o *PartnerBindOrganizationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PartnerBindOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerBindOrganizationsBadRequest creates a PartnerBindOrganizationsBadRequest with default headers values
func NewPartnerBindOrganizationsBadRequest() *PartnerBindOrganizationsBadRequest {
	return &PartnerBindOrganizationsBadRequest{}
}

/*
PartnerBindOrganizationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PartnerBindOrganizationsBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner bind organizations bad request response has a 2xx status code
func (o *PartnerBindOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner bind organizations bad request response has a 3xx status code
func (o *PartnerBindOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner bind organizations bad request response has a 4xx status code
func (o *PartnerBindOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner bind organizations bad request response has a 5xx status code
func (o *PartnerBindOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this partner bind organizations bad request response a status code equal to that given
func (o *PartnerBindOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PartnerBindOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerBindOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *PartnerBindOrganizationsBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerBindOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerBindOrganizationsUnauthorized creates a PartnerBindOrganizationsUnauthorized with default headers values
func NewPartnerBindOrganizationsUnauthorized() *PartnerBindOrganizationsUnauthorized {
	return &PartnerBindOrganizationsUnauthorized{}
}

/*
PartnerBindOrganizationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PartnerBindOrganizationsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner bind organizations unauthorized response has a 2xx status code
func (o *PartnerBindOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner bind organizations unauthorized response has a 3xx status code
func (o *PartnerBindOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner bind organizations unauthorized response has a 4xx status code
func (o *PartnerBindOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner bind organizations unauthorized response has a 5xx status code
func (o *PartnerBindOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this partner bind organizations unauthorized response a status code equal to that given
func (o *PartnerBindOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PartnerBindOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerBindOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PartnerBindOrganizationsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerBindOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerBindOrganizationsForbidden creates a PartnerBindOrganizationsForbidden with default headers values
func NewPartnerBindOrganizationsForbidden() *PartnerBindOrganizationsForbidden {
	return &PartnerBindOrganizationsForbidden{}
}

/*
PartnerBindOrganizationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PartnerBindOrganizationsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner bind organizations forbidden response has a 2xx status code
func (o *PartnerBindOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner bind organizations forbidden response has a 3xx status code
func (o *PartnerBindOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner bind organizations forbidden response has a 4xx status code
func (o *PartnerBindOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner bind organizations forbidden response has a 5xx status code
func (o *PartnerBindOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this partner bind organizations forbidden response a status code equal to that given
func (o *PartnerBindOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PartnerBindOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *PartnerBindOrganizationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *PartnerBindOrganizationsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerBindOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerBindOrganizationsNotFound creates a PartnerBindOrganizationsNotFound with default headers values
func NewPartnerBindOrganizationsNotFound() *PartnerBindOrganizationsNotFound {
	return &PartnerBindOrganizationsNotFound{}
}

/*
PartnerBindOrganizationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PartnerBindOrganizationsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this partner bind organizations not found response has a 2xx status code
func (o *PartnerBindOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner bind organizations not found response has a 3xx status code
func (o *PartnerBindOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner bind organizations not found response has a 4xx status code
func (o *PartnerBindOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this partner bind organizations not found response has a 5xx status code
func (o *PartnerBindOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this partner bind organizations not found response a status code equal to that given
func (o *PartnerBindOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PartnerBindOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *PartnerBindOrganizationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *PartnerBindOrganizationsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PartnerBindOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartnerBindOrganizationsInternalServerError creates a PartnerBindOrganizationsInternalServerError with default headers values
func NewPartnerBindOrganizationsInternalServerError() *PartnerBindOrganizationsInternalServerError {
	return &PartnerBindOrganizationsInternalServerError{}
}

/*
PartnerBindOrganizationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PartnerBindOrganizationsInternalServerError struct {
}

// IsSuccess returns true when this partner bind organizations internal server error response has a 2xx status code
func (o *PartnerBindOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this partner bind organizations internal server error response has a 3xx status code
func (o *PartnerBindOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this partner bind organizations internal server error response has a 4xx status code
func (o *PartnerBindOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this partner bind organizations internal server error response has a 5xx status code
func (o *PartnerBindOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this partner bind organizations internal server error response a status code equal to that given
func (o *PartnerBindOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PartnerBindOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsInternalServerError ", 500)
}

func (o *PartnerBindOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Partner/bindorganizations][%d] partnerBindOrganizationsInternalServerError ", 500)
}

func (o *PartnerBindOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
