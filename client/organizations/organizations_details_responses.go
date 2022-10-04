// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OrganizationsDetailsReader is a Reader for the OrganizationsDetails structure.
type OrganizationsDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsDetailsOK creates a OrganizationsDetailsOK with default headers values
func NewOrganizationsDetailsOK() *OrganizationsDetailsOK {
	return &OrganizationsDetailsOK{}
}

/*
OrganizationsDetailsOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsDetailsOK struct {
	Payload *models.DashboardChart
}

// IsSuccess returns true when this organizations details o k response has a 2xx status code
func (o *OrganizationsDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations details o k response has a 3xx status code
func (o *OrganizationsDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations details o k response has a 4xx status code
func (o *OrganizationsDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations details o k response has a 5xx status code
func (o *OrganizationsDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations details o k response a status code equal to that given
func (o *OrganizationsDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsOK  %+v", 200, o.Payload)
}

func (o *OrganizationsDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsOK  %+v", 200, o.Payload)
}

func (o *OrganizationsDetailsOK) GetPayload() *models.DashboardChart {
	return o.Payload
}

func (o *OrganizationsDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardChart)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsBadRequest creates a OrganizationsDetailsBadRequest with default headers values
func NewOrganizationsDetailsBadRequest() *OrganizationsDetailsBadRequest {
	return &OrganizationsDetailsBadRequest{}
}

/*
OrganizationsDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsDetailsBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations details bad request response has a 2xx status code
func (o *OrganizationsDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations details bad request response has a 3xx status code
func (o *OrganizationsDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations details bad request response has a 4xx status code
func (o *OrganizationsDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations details bad request response has a 5xx status code
func (o *OrganizationsDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations details bad request response a status code equal to that given
func (o *OrganizationsDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OrganizationsDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsDetailsBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsUnauthorized creates a OrganizationsDetailsUnauthorized with default headers values
func NewOrganizationsDetailsUnauthorized() *OrganizationsDetailsUnauthorized {
	return &OrganizationsDetailsUnauthorized{}
}

/*
OrganizationsDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsDetailsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations details unauthorized response has a 2xx status code
func (o *OrganizationsDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations details unauthorized response has a 3xx status code
func (o *OrganizationsDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations details unauthorized response has a 4xx status code
func (o *OrganizationsDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations details unauthorized response has a 5xx status code
func (o *OrganizationsDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations details unauthorized response a status code equal to that given
func (o *OrganizationsDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OrganizationsDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsDetailsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsForbidden creates a OrganizationsDetailsForbidden with default headers values
func NewOrganizationsDetailsForbidden() *OrganizationsDetailsForbidden {
	return &OrganizationsDetailsForbidden{}
}

/*
OrganizationsDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsDetailsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations details forbidden response has a 2xx status code
func (o *OrganizationsDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations details forbidden response has a 3xx status code
func (o *OrganizationsDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations details forbidden response has a 4xx status code
func (o *OrganizationsDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations details forbidden response has a 5xx status code
func (o *OrganizationsDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations details forbidden response a status code equal to that given
func (o *OrganizationsDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsDetailsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsNotFound creates a OrganizationsDetailsNotFound with default headers values
func NewOrganizationsDetailsNotFound() *OrganizationsDetailsNotFound {
	return &OrganizationsDetailsNotFound{}
}

/*
OrganizationsDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsDetailsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this organizations details not found response has a 2xx status code
func (o *OrganizationsDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations details not found response has a 3xx status code
func (o *OrganizationsDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations details not found response has a 4xx status code
func (o *OrganizationsDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations details not found response has a 5xx status code
func (o *OrganizationsDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations details not found response a status code equal to that given
func (o *OrganizationsDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OrganizationsDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsDetailsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *OrganizationsDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsDetailsInternalServerError creates a OrganizationsDetailsInternalServerError with default headers values
func NewOrganizationsDetailsInternalServerError() *OrganizationsDetailsInternalServerError {
	return &OrganizationsDetailsInternalServerError{}
}

/*
OrganizationsDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsDetailsInternalServerError struct {
}

// IsSuccess returns true when this organizations details internal server error response has a 2xx status code
func (o *OrganizationsDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations details internal server error response has a 3xx status code
func (o *OrganizationsDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations details internal server error response has a 4xx status code
func (o *OrganizationsDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations details internal server error response has a 5xx status code
func (o *OrganizationsDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations details internal server error response a status code equal to that given
func (o *OrganizationsDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OrganizationsDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsInternalServerError ", 500)
}

func (o *OrganizationsDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Organizations/details/{organizationId}][%d] organizationsDetailsInternalServerError ", 500)
}

func (o *OrganizationsDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
