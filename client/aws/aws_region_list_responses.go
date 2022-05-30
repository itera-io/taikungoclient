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

// AwsRegionListReader is a Reader for the AwsRegionList structure.
type AwsRegionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsRegionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsRegionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsRegionListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsRegionListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsRegionListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsRegionListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsRegionListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsRegionListOK creates a AwsRegionListOK with default headers values
func NewAwsRegionListOK() *AwsRegionListOK {
	return &AwsRegionListOK{}
}

/* AwsRegionListOK describes a response with status code 200, with default header values.

Success
*/
type AwsRegionListOK struct {
	Payload []*models.AwsRegionDto
}

func (o *AwsRegionListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/regions][%d] awsRegionListOK  %+v", 200, o.Payload)
}
func (o *AwsRegionListOK) GetPayload() []*models.AwsRegionDto {
	return o.Payload
}

func (o *AwsRegionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsRegionListBadRequest creates a AwsRegionListBadRequest with default headers values
func NewAwsRegionListBadRequest() *AwsRegionListBadRequest {
	return &AwsRegionListBadRequest{}
}

/* AwsRegionListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsRegionListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AwsRegionListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/regions][%d] awsRegionListBadRequest  %+v", 400, o.Payload)
}
func (o *AwsRegionListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AwsRegionListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsRegionListUnauthorized creates a AwsRegionListUnauthorized with default headers values
func NewAwsRegionListUnauthorized() *AwsRegionListUnauthorized {
	return &AwsRegionListUnauthorized{}
}

/* AwsRegionListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsRegionListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AwsRegionListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/regions][%d] awsRegionListUnauthorized  %+v", 401, o.Payload)
}
func (o *AwsRegionListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsRegionListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsRegionListForbidden creates a AwsRegionListForbidden with default headers values
func NewAwsRegionListForbidden() *AwsRegionListForbidden {
	return &AwsRegionListForbidden{}
}

/* AwsRegionListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsRegionListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AwsRegionListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/regions][%d] awsRegionListForbidden  %+v", 403, o.Payload)
}
func (o *AwsRegionListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsRegionListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsRegionListNotFound creates a AwsRegionListNotFound with default headers values
func NewAwsRegionListNotFound() *AwsRegionListNotFound {
	return &AwsRegionListNotFound{}
}

/* AwsRegionListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsRegionListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AwsRegionListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/regions][%d] awsRegionListNotFound  %+v", 404, o.Payload)
}
func (o *AwsRegionListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsRegionListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsRegionListInternalServerError creates a AwsRegionListInternalServerError with default headers values
func NewAwsRegionListInternalServerError() *AwsRegionListInternalServerError {
	return &AwsRegionListInternalServerError{}
}

/* AwsRegionListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsRegionListInternalServerError struct {
}

func (o *AwsRegionListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/regions][%d] awsRegionListInternalServerError ", 500)
}

func (o *AwsRegionListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}