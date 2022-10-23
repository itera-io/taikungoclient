// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CloudCredentialsDashboardListReader is a Reader for the CloudCredentialsDashboardList structure.
type CloudCredentialsDashboardListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsDashboardListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsDashboardListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsDashboardListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsDashboardListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsDashboardListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsDashboardListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsDashboardListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsDashboardListOK creates a CloudCredentialsDashboardListOK with default headers values
func NewCloudCredentialsDashboardListOK() *CloudCredentialsDashboardListOK {
	return &CloudCredentialsDashboardListOK{}
}

/*
CloudCredentialsDashboardListOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsDashboardListOK struct {
	Payload *models.CredentialsChart
}

// IsSuccess returns true when this cloud credentials dashboard list o k response has a 2xx status code
func (o *CloudCredentialsDashboardListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud credentials dashboard list o k response has a 3xx status code
func (o *CloudCredentialsDashboardListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials dashboard list o k response has a 4xx status code
func (o *CloudCredentialsDashboardListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials dashboard list o k response has a 5xx status code
func (o *CloudCredentialsDashboardListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials dashboard list o k response a status code equal to that given
func (o *CloudCredentialsDashboardListOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudCredentialsDashboardListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsDashboardListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListOK  %+v", 200, o.Payload)
}

func (o *CloudCredentialsDashboardListOK) GetPayload() *models.CredentialsChart {
	return o.Payload
}

func (o *CloudCredentialsDashboardListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsChart)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsDashboardListBadRequest creates a CloudCredentialsDashboardListBadRequest with default headers values
func NewCloudCredentialsDashboardListBadRequest() *CloudCredentialsDashboardListBadRequest {
	return &CloudCredentialsDashboardListBadRequest{}
}

/*
CloudCredentialsDashboardListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsDashboardListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this cloud credentials dashboard list bad request response has a 2xx status code
func (o *CloudCredentialsDashboardListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials dashboard list bad request response has a 3xx status code
func (o *CloudCredentialsDashboardListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials dashboard list bad request response has a 4xx status code
func (o *CloudCredentialsDashboardListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials dashboard list bad request response has a 5xx status code
func (o *CloudCredentialsDashboardListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials dashboard list bad request response a status code equal to that given
func (o *CloudCredentialsDashboardListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CloudCredentialsDashboardListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsDashboardListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListBadRequest  %+v", 400, o.Payload)
}

func (o *CloudCredentialsDashboardListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CloudCredentialsDashboardListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsDashboardListUnauthorized creates a CloudCredentialsDashboardListUnauthorized with default headers values
func NewCloudCredentialsDashboardListUnauthorized() *CloudCredentialsDashboardListUnauthorized {
	return &CloudCredentialsDashboardListUnauthorized{}
}

/*
CloudCredentialsDashboardListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsDashboardListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cloud credentials dashboard list unauthorized response has a 2xx status code
func (o *CloudCredentialsDashboardListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials dashboard list unauthorized response has a 3xx status code
func (o *CloudCredentialsDashboardListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials dashboard list unauthorized response has a 4xx status code
func (o *CloudCredentialsDashboardListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials dashboard list unauthorized response has a 5xx status code
func (o *CloudCredentialsDashboardListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials dashboard list unauthorized response a status code equal to that given
func (o *CloudCredentialsDashboardListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloudCredentialsDashboardListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsDashboardListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListUnauthorized  %+v", 401, o.Payload)
}

func (o *CloudCredentialsDashboardListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsDashboardListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsDashboardListForbidden creates a CloudCredentialsDashboardListForbidden with default headers values
func NewCloudCredentialsDashboardListForbidden() *CloudCredentialsDashboardListForbidden {
	return &CloudCredentialsDashboardListForbidden{}
}

/*
CloudCredentialsDashboardListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsDashboardListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cloud credentials dashboard list forbidden response has a 2xx status code
func (o *CloudCredentialsDashboardListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials dashboard list forbidden response has a 3xx status code
func (o *CloudCredentialsDashboardListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials dashboard list forbidden response has a 4xx status code
func (o *CloudCredentialsDashboardListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials dashboard list forbidden response has a 5xx status code
func (o *CloudCredentialsDashboardListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials dashboard list forbidden response a status code equal to that given
func (o *CloudCredentialsDashboardListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloudCredentialsDashboardListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsDashboardListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListForbidden  %+v", 403, o.Payload)
}

func (o *CloudCredentialsDashboardListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsDashboardListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsDashboardListNotFound creates a CloudCredentialsDashboardListNotFound with default headers values
func NewCloudCredentialsDashboardListNotFound() *CloudCredentialsDashboardListNotFound {
	return &CloudCredentialsDashboardListNotFound{}
}

/*
CloudCredentialsDashboardListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsDashboardListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cloud credentials dashboard list not found response has a 2xx status code
func (o *CloudCredentialsDashboardListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials dashboard list not found response has a 3xx status code
func (o *CloudCredentialsDashboardListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials dashboard list not found response has a 4xx status code
func (o *CloudCredentialsDashboardListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud credentials dashboard list not found response has a 5xx status code
func (o *CloudCredentialsDashboardListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud credentials dashboard list not found response a status code equal to that given
func (o *CloudCredentialsDashboardListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloudCredentialsDashboardListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsDashboardListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListNotFound  %+v", 404, o.Payload)
}

func (o *CloudCredentialsDashboardListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsDashboardListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsDashboardListInternalServerError creates a CloudCredentialsDashboardListInternalServerError with default headers values
func NewCloudCredentialsDashboardListInternalServerError() *CloudCredentialsDashboardListInternalServerError {
	return &CloudCredentialsDashboardListInternalServerError{}
}

/*
CloudCredentialsDashboardListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsDashboardListInternalServerError struct {
}

// IsSuccess returns true when this cloud credentials dashboard list internal server error response has a 2xx status code
func (o *CloudCredentialsDashboardListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud credentials dashboard list internal server error response has a 3xx status code
func (o *CloudCredentialsDashboardListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud credentials dashboard list internal server error response has a 4xx status code
func (o *CloudCredentialsDashboardListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud credentials dashboard list internal server error response has a 5xx status code
func (o *CloudCredentialsDashboardListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud credentials dashboard list internal server error response a status code equal to that given
func (o *CloudCredentialsDashboardListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloudCredentialsDashboardListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListInternalServerError ", 500)
}

func (o *CloudCredentialsDashboardListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListInternalServerError ", 500)
}

func (o *CloudCredentialsDashboardListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
