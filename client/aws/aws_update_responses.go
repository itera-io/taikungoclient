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

// AwsUpdateReader is a Reader for the AwsUpdate structure.
type AwsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAwsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAwsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAwsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAwsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAwsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAwsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAwsUpdateOK creates a AwsUpdateOK with default headers values
func NewAwsUpdateOK() *AwsUpdateOK {
	return &AwsUpdateOK{}
}

/* AwsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type AwsUpdateOK struct {
	Payload models.Unit
}

func (o *AwsUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/update][%d] awsUpdateOK  %+v", 200, o.Payload)
}
func (o *AwsUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AwsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsUpdateBadRequest creates a AwsUpdateBadRequest with default headers values
func NewAwsUpdateBadRequest() *AwsUpdateBadRequest {
	return &AwsUpdateBadRequest{}
}

/* AwsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AwsUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AwsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/update][%d] awsUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *AwsUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AwsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsUpdateUnauthorized creates a AwsUpdateUnauthorized with default headers values
func NewAwsUpdateUnauthorized() *AwsUpdateUnauthorized {
	return &AwsUpdateUnauthorized{}
}

/* AwsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AwsUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AwsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/update][%d] awsUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *AwsUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsUpdateForbidden creates a AwsUpdateForbidden with default headers values
func NewAwsUpdateForbidden() *AwsUpdateForbidden {
	return &AwsUpdateForbidden{}
}

/* AwsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AwsUpdateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AwsUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/update][%d] awsUpdateForbidden  %+v", 403, o.Payload)
}
func (o *AwsUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsUpdateNotFound creates a AwsUpdateNotFound with default headers values
func NewAwsUpdateNotFound() *AwsUpdateNotFound {
	return &AwsUpdateNotFound{}
}

/* AwsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AwsUpdateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AwsUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/update][%d] awsUpdateNotFound  %+v", 404, o.Payload)
}
func (o *AwsUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AwsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsUpdateInternalServerError creates a AwsUpdateInternalServerError with default headers values
func NewAwsUpdateInternalServerError() *AwsUpdateInternalServerError {
	return &AwsUpdateInternalServerError{}
}

/* AwsUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AwsUpdateInternalServerError struct {
}

func (o *AwsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Aws/update][%d] awsUpdateInternalServerError ", 500)
}

func (o *AwsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
