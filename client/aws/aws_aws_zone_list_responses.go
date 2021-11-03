// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AwsAwsZoneListReader is a Reader for the AwsAwsZoneList structure.
type AwsAwsZoneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsAwsZoneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsAwsZoneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsAwsZoneListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsAwsZoneListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsAwsZoneListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsAwsZoneListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAwsAwsZoneListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsAwsZoneListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsAwsZoneListOK creates a AwsAwsZoneListOK with default headers values
func NewAwsAwsZoneListOK() *AwsAwsZoneListOK {
	return &AwsAwsZoneListOK{}
}

/* AwsAwsZoneListOK describes a response with status code 200, with default header values.

Success
*/
type AwsAwsZoneListOK struct {
	Payload []string
}

func (o *AwsAwsZoneListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListOK  %+v", 200, o.Payload)
}
func (o *AwsAwsZoneListOK) GetPayload() []string {
	return o.Payload
}

func (o *AwsAwsZoneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListBadRequest creates a AwsAwsZoneListBadRequest with default headers values
func NewAwsAwsZoneListBadRequest() *AwsAwsZoneListBadRequest {
	return &AwsAwsZoneListBadRequest{}
}

/* AwsAwsZoneListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsAwsZoneListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AwsAwsZoneListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListBadRequest  %+v", 400, o.Payload)
}
func (o *AwsAwsZoneListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListUnauthorized creates a AwsAwsZoneListUnauthorized with default headers values
func NewAwsAwsZoneListUnauthorized() *AwsAwsZoneListUnauthorized {
	return &AwsAwsZoneListUnauthorized{}
}

/* AwsAwsZoneListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsAwsZoneListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AwsAwsZoneListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListUnauthorized  %+v", 401, o.Payload)
}
func (o *AwsAwsZoneListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListForbidden creates a AwsAwsZoneListForbidden with default headers values
func NewAwsAwsZoneListForbidden() *AwsAwsZoneListForbidden {
	return &AwsAwsZoneListForbidden{}
}

/* AwsAwsZoneListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsAwsZoneListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AwsAwsZoneListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListForbidden  %+v", 403, o.Payload)
}
func (o *AwsAwsZoneListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListNotFound creates a AwsAwsZoneListNotFound with default headers values
func NewAwsAwsZoneListNotFound() *AwsAwsZoneListNotFound {
	return &AwsAwsZoneListNotFound{}
}

/* AwsAwsZoneListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsAwsZoneListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AwsAwsZoneListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListNotFound  %+v", 404, o.Payload)
}
func (o *AwsAwsZoneListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListTooManyRequests creates a AwsAwsZoneListTooManyRequests with default headers values
func NewAwsAwsZoneListTooManyRequests() *AwsAwsZoneListTooManyRequests {
	return &AwsAwsZoneListTooManyRequests{}
}

/* AwsAwsZoneListTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AwsAwsZoneListTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AwsAwsZoneListTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListTooManyRequests  %+v", 429, o.Payload)
}
func (o *AwsAwsZoneListTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsAwsZoneListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsAwsZoneListInternalServerError creates a AwsAwsZoneListInternalServerError with default headers values
func NewAwsAwsZoneListInternalServerError() *AwsAwsZoneListInternalServerError {
	return &AwsAwsZoneListInternalServerError{}
}

/* AwsAwsZoneListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsAwsZoneListInternalServerError struct {
}

func (o *AwsAwsZoneListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/zones][%d] awsAwsZoneListInternalServerError ", 500)
}

func (o *AwsAwsZoneListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
