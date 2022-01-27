// Code generated by go-swagger; DO NOT EDIT.

package google_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// GoogleCloudCreateReader is a Reader for the GoogleCloudCreate structure.
type GoogleCloudCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GoogleCloudCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGoogleCloudCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGoogleCloudCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGoogleCloudCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGoogleCloudCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGoogleCloudCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGoogleCloudCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGoogleCloudCreateOK creates a GoogleCloudCreateOK with default headers values
func NewGoogleCloudCreateOK() *GoogleCloudCreateOK {
	return &GoogleCloudCreateOK{}
}

/* GoogleCloudCreateOK describes a response with status code 200, with default header values.

Success
*/
type GoogleCloudCreateOK struct {
	Payload models.Unit
}

func (o *GoogleCloudCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/create][%d] googleCloudCreateOK  %+v", 200, o.Payload)
}
func (o *GoogleCloudCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *GoogleCloudCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudCreateBadRequest creates a GoogleCloudCreateBadRequest with default headers values
func NewGoogleCloudCreateBadRequest() *GoogleCloudCreateBadRequest {
	return &GoogleCloudCreateBadRequest{}
}

/* GoogleCloudCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GoogleCloudCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *GoogleCloudCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/create][%d] googleCloudCreateBadRequest  %+v", 400, o.Payload)
}
func (o *GoogleCloudCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *GoogleCloudCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudCreateUnauthorized creates a GoogleCloudCreateUnauthorized with default headers values
func NewGoogleCloudCreateUnauthorized() *GoogleCloudCreateUnauthorized {
	return &GoogleCloudCreateUnauthorized{}
}

/* GoogleCloudCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GoogleCloudCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *GoogleCloudCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/create][%d] googleCloudCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *GoogleCloudCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudCreateForbidden creates a GoogleCloudCreateForbidden with default headers values
func NewGoogleCloudCreateForbidden() *GoogleCloudCreateForbidden {
	return &GoogleCloudCreateForbidden{}
}

/* GoogleCloudCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GoogleCloudCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *GoogleCloudCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/create][%d] googleCloudCreateForbidden  %+v", 403, o.Payload)
}
func (o *GoogleCloudCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudCreateNotFound creates a GoogleCloudCreateNotFound with default headers values
func NewGoogleCloudCreateNotFound() *GoogleCloudCreateNotFound {
	return &GoogleCloudCreateNotFound{}
}

/* GoogleCloudCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GoogleCloudCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *GoogleCloudCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/create][%d] googleCloudCreateNotFound  %+v", 404, o.Payload)
}
func (o *GoogleCloudCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudCreateInternalServerError creates a GoogleCloudCreateInternalServerError with default headers values
func NewGoogleCloudCreateInternalServerError() *GoogleCloudCreateInternalServerError {
	return &GoogleCloudCreateInternalServerError{}
}

/* GoogleCloudCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GoogleCloudCreateInternalServerError struct {
}

func (o *GoogleCloudCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/create][%d] googleCloudCreateInternalServerError ", 500)
}

func (o *GoogleCloudCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
