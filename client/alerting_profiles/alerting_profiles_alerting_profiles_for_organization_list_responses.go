// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AlertingProfilesAlertingProfilesForOrganizationListReader is a Reader for the AlertingProfilesAlertingProfilesForOrganizationList structure.
type AlertingProfilesAlertingProfilesForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesAlertingProfilesForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesAlertingProfilesForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesAlertingProfilesForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesAlertingProfilesForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesAlertingProfilesForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesAlertingProfilesForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesAlertingProfilesForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesAlertingProfilesForOrganizationListOK creates a AlertingProfilesAlertingProfilesForOrganizationListOK with default headers values
func NewAlertingProfilesAlertingProfilesForOrganizationListOK() *AlertingProfilesAlertingProfilesForOrganizationListOK {
	return &AlertingProfilesAlertingProfilesForOrganizationListOK{}
}

/*
AlertingProfilesAlertingProfilesForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesAlertingProfilesForOrganizationListOK struct {
	Payload []*models.CommonDropdownDto
}

// IsSuccess returns true when this alerting profiles alerting profiles for organization list o k response has a 2xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles alerting profiles for organization list o k response has a 3xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles alerting profiles for organization list o k response has a 4xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles alerting profiles for organization list o k response has a 5xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles alerting profiles for organization list o k response a status code equal to that given
func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) GetPayload() []*models.CommonDropdownDto {
	return o.Payload
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAlertingProfilesForOrganizationListBadRequest creates a AlertingProfilesAlertingProfilesForOrganizationListBadRequest with default headers values
func NewAlertingProfilesAlertingProfilesForOrganizationListBadRequest() *AlertingProfilesAlertingProfilesForOrganizationListBadRequest {
	return &AlertingProfilesAlertingProfilesForOrganizationListBadRequest{}
}

/*
AlertingProfilesAlertingProfilesForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesAlertingProfilesForOrganizationListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this alerting profiles alerting profiles for organization list bad request response has a 2xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles alerting profiles for organization list bad request response has a 3xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles alerting profiles for organization list bad request response has a 4xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles alerting profiles for organization list bad request response has a 5xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles alerting profiles for organization list bad request response a status code equal to that given
func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAlertingProfilesForOrganizationListUnauthorized creates a AlertingProfilesAlertingProfilesForOrganizationListUnauthorized with default headers values
func NewAlertingProfilesAlertingProfilesForOrganizationListUnauthorized() *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized {
	return &AlertingProfilesAlertingProfilesForOrganizationListUnauthorized{}
}

/*
AlertingProfilesAlertingProfilesForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesAlertingProfilesForOrganizationListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles alerting profiles for organization list unauthorized response has a 2xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles alerting profiles for organization list unauthorized response has a 3xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles alerting profiles for organization list unauthorized response has a 4xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles alerting profiles for organization list unauthorized response has a 5xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles alerting profiles for organization list unauthorized response a status code equal to that given
func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAlertingProfilesForOrganizationListForbidden creates a AlertingProfilesAlertingProfilesForOrganizationListForbidden with default headers values
func NewAlertingProfilesAlertingProfilesForOrganizationListForbidden() *AlertingProfilesAlertingProfilesForOrganizationListForbidden {
	return &AlertingProfilesAlertingProfilesForOrganizationListForbidden{}
}

/*
AlertingProfilesAlertingProfilesForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesAlertingProfilesForOrganizationListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles alerting profiles for organization list forbidden response has a 2xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles alerting profiles for organization list forbidden response has a 3xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles alerting profiles for organization list forbidden response has a 4xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles alerting profiles for organization list forbidden response has a 5xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles alerting profiles for organization list forbidden response a status code equal to that given
func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAlertingProfilesForOrganizationListNotFound creates a AlertingProfilesAlertingProfilesForOrganizationListNotFound with default headers values
func NewAlertingProfilesAlertingProfilesForOrganizationListNotFound() *AlertingProfilesAlertingProfilesForOrganizationListNotFound {
	return &AlertingProfilesAlertingProfilesForOrganizationListNotFound{}
}

/*
AlertingProfilesAlertingProfilesForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesAlertingProfilesForOrganizationListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles alerting profiles for organization list not found response has a 2xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles alerting profiles for organization list not found response has a 3xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles alerting profiles for organization list not found response has a 4xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles alerting profiles for organization list not found response has a 5xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles alerting profiles for organization list not found response a status code equal to that given
func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAlertingProfilesForOrganizationListInternalServerError creates a AlertingProfilesAlertingProfilesForOrganizationListInternalServerError with default headers values
func NewAlertingProfilesAlertingProfilesForOrganizationListInternalServerError() *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError {
	return &AlertingProfilesAlertingProfilesForOrganizationListInternalServerError{}
}

/*
AlertingProfilesAlertingProfilesForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesAlertingProfilesForOrganizationListInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles alerting profiles for organization list internal server error response has a 2xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles alerting profiles for organization list internal server error response has a 3xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles alerting profiles for organization list internal server error response has a 4xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles alerting profiles for organization list internal server error response has a 5xx status code
func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles alerting profiles for organization list internal server error response a status code equal to that given
func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListInternalServerError ", 500)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles/list][%d] alertingProfilesAlertingProfilesForOrganizationListInternalServerError ", 500)
}

func (o *AlertingProfilesAlertingProfilesForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
