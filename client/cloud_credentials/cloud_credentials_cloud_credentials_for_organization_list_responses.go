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

// CloudCredentialsCloudCredentialsForOrganizationListReader is a Reader for the CloudCredentialsCloudCredentialsForOrganizationList structure.
type CloudCredentialsCloudCredentialsForOrganizationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsCloudCredentialsForOrganizationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsCloudCredentialsForOrganizationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsCloudCredentialsForOrganizationListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsCloudCredentialsForOrganizationListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsCloudCredentialsForOrganizationListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsCloudCredentialsForOrganizationListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsCloudCredentialsForOrganizationListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsCloudCredentialsForOrganizationListOK creates a CloudCredentialsCloudCredentialsForOrganizationListOK with default headers values
func NewCloudCredentialsCloudCredentialsForOrganizationListOK() *CloudCredentialsCloudCredentialsForOrganizationListOK {
	return &CloudCredentialsCloudCredentialsForOrganizationListOK{}
}

/* CloudCredentialsCloudCredentialsForOrganizationListOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsCloudCredentialsForOrganizationListOK struct {
	Payload []*models.CloudCredentialsForOrganizationEntity
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials][%d] cloudCredentialsCloudCredentialsForOrganizationListOK  %+v", 200, o.Payload)
}
func (o *CloudCredentialsCloudCredentialsForOrganizationListOK) GetPayload() []*models.CloudCredentialsForOrganizationEntity {
	return o.Payload
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsCloudCredentialsForOrganizationListBadRequest creates a CloudCredentialsCloudCredentialsForOrganizationListBadRequest with default headers values
func NewCloudCredentialsCloudCredentialsForOrganizationListBadRequest() *CloudCredentialsCloudCredentialsForOrganizationListBadRequest {
	return &CloudCredentialsCloudCredentialsForOrganizationListBadRequest{}
}

/* CloudCredentialsCloudCredentialsForOrganizationListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsCloudCredentialsForOrganizationListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials][%d] cloudCredentialsCloudCredentialsForOrganizationListBadRequest  %+v", 400, o.Payload)
}
func (o *CloudCredentialsCloudCredentialsForOrganizationListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsCloudCredentialsForOrganizationListUnauthorized creates a CloudCredentialsCloudCredentialsForOrganizationListUnauthorized with default headers values
func NewCloudCredentialsCloudCredentialsForOrganizationListUnauthorized() *CloudCredentialsCloudCredentialsForOrganizationListUnauthorized {
	return &CloudCredentialsCloudCredentialsForOrganizationListUnauthorized{}
}

/* CloudCredentialsCloudCredentialsForOrganizationListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsCloudCredentialsForOrganizationListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials][%d] cloudCredentialsCloudCredentialsForOrganizationListUnauthorized  %+v", 401, o.Payload)
}
func (o *CloudCredentialsCloudCredentialsForOrganizationListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsCloudCredentialsForOrganizationListForbidden creates a CloudCredentialsCloudCredentialsForOrganizationListForbidden with default headers values
func NewCloudCredentialsCloudCredentialsForOrganizationListForbidden() *CloudCredentialsCloudCredentialsForOrganizationListForbidden {
	return &CloudCredentialsCloudCredentialsForOrganizationListForbidden{}
}

/* CloudCredentialsCloudCredentialsForOrganizationListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsCloudCredentialsForOrganizationListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials][%d] cloudCredentialsCloudCredentialsForOrganizationListForbidden  %+v", 403, o.Payload)
}
func (o *CloudCredentialsCloudCredentialsForOrganizationListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsCloudCredentialsForOrganizationListNotFound creates a CloudCredentialsCloudCredentialsForOrganizationListNotFound with default headers values
func NewCloudCredentialsCloudCredentialsForOrganizationListNotFound() *CloudCredentialsCloudCredentialsForOrganizationListNotFound {
	return &CloudCredentialsCloudCredentialsForOrganizationListNotFound{}
}

/* CloudCredentialsCloudCredentialsForOrganizationListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsCloudCredentialsForOrganizationListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials][%d] cloudCredentialsCloudCredentialsForOrganizationListNotFound  %+v", 404, o.Payload)
}
func (o *CloudCredentialsCloudCredentialsForOrganizationListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsCloudCredentialsForOrganizationListInternalServerError creates a CloudCredentialsCloudCredentialsForOrganizationListInternalServerError with default headers values
func NewCloudCredentialsCloudCredentialsForOrganizationListInternalServerError() *CloudCredentialsCloudCredentialsForOrganizationListInternalServerError {
	return &CloudCredentialsCloudCredentialsForOrganizationListInternalServerError{}
}

/* CloudCredentialsCloudCredentialsForOrganizationListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsCloudCredentialsForOrganizationListInternalServerError struct {
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials][%d] cloudCredentialsCloudCredentialsForOrganizationListInternalServerError ", 500)
}

func (o *CloudCredentialsCloudCredentialsForOrganizationListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
