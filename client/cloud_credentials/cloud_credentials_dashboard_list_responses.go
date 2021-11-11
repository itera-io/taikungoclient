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

/* CloudCredentialsDashboardListOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsDashboardListOK struct {
	Payload *models.CredentialsChart
}

func (o *CloudCredentialsDashboardListOK) Error() string {
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

/* CloudCredentialsDashboardListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsDashboardListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CloudCredentialsDashboardListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListBadRequest  %+v", 400, o.Payload)
}
func (o *CloudCredentialsDashboardListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsDashboardListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsDashboardListUnauthorized creates a CloudCredentialsDashboardListUnauthorized with default headers values
func NewCloudCredentialsDashboardListUnauthorized() *CloudCredentialsDashboardListUnauthorized {
	return &CloudCredentialsDashboardListUnauthorized{}
}

/* CloudCredentialsDashboardListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsDashboardListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsDashboardListUnauthorized) Error() string {
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

/* CloudCredentialsDashboardListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsDashboardListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsDashboardListForbidden) Error() string {
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

/* CloudCredentialsDashboardListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsDashboardListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsDashboardListNotFound) Error() string {
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

/* CloudCredentialsDashboardListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsDashboardListInternalServerError struct {
}

func (o *CloudCredentialsDashboardListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/list][%d] cloudCredentialsDashboardListInternalServerError ", 500)
}

func (o *CloudCredentialsDashboardListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
