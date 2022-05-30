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

// GoogleCloudZoneListReader is a Reader for the GoogleCloudZoneList structure.
type GoogleCloudZoneListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GoogleCloudZoneListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGoogleCloudZoneListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGoogleCloudZoneListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGoogleCloudZoneListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGoogleCloudZoneListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGoogleCloudZoneListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGoogleCloudZoneListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGoogleCloudZoneListOK creates a GoogleCloudZoneListOK with default headers values
func NewGoogleCloudZoneListOK() *GoogleCloudZoneListOK {
	return &GoogleCloudZoneListOK{}
}

/* GoogleCloudZoneListOK describes a response with status code 200, with default header values.

Success
*/
type GoogleCloudZoneListOK struct {
	Payload []string
}

func (o *GoogleCloudZoneListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListOK  %+v", 200, o.Payload)
}
func (o *GoogleCloudZoneListOK) GetPayload() []string {
	return o.Payload
}

func (o *GoogleCloudZoneListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListBadRequest creates a GoogleCloudZoneListBadRequest with default headers values
func NewGoogleCloudZoneListBadRequest() *GoogleCloudZoneListBadRequest {
	return &GoogleCloudZoneListBadRequest{}
}

/* GoogleCloudZoneListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GoogleCloudZoneListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *GoogleCloudZoneListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListBadRequest  %+v", 400, o.Payload)
}
func (o *GoogleCloudZoneListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *GoogleCloudZoneListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListUnauthorized creates a GoogleCloudZoneListUnauthorized with default headers values
func NewGoogleCloudZoneListUnauthorized() *GoogleCloudZoneListUnauthorized {
	return &GoogleCloudZoneListUnauthorized{}
}

/* GoogleCloudZoneListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GoogleCloudZoneListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *GoogleCloudZoneListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListUnauthorized  %+v", 401, o.Payload)
}
func (o *GoogleCloudZoneListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudZoneListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListForbidden creates a GoogleCloudZoneListForbidden with default headers values
func NewGoogleCloudZoneListForbidden() *GoogleCloudZoneListForbidden {
	return &GoogleCloudZoneListForbidden{}
}

/* GoogleCloudZoneListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GoogleCloudZoneListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *GoogleCloudZoneListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListForbidden  %+v", 403, o.Payload)
}
func (o *GoogleCloudZoneListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudZoneListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListNotFound creates a GoogleCloudZoneListNotFound with default headers values
func NewGoogleCloudZoneListNotFound() *GoogleCloudZoneListNotFound {
	return &GoogleCloudZoneListNotFound{}
}

/* GoogleCloudZoneListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GoogleCloudZoneListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *GoogleCloudZoneListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListNotFound  %+v", 404, o.Payload)
}
func (o *GoogleCloudZoneListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GoogleCloudZoneListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGoogleCloudZoneListInternalServerError creates a GoogleCloudZoneListInternalServerError with default headers values
func NewGoogleCloudZoneListInternalServerError() *GoogleCloudZoneListInternalServerError {
	return &GoogleCloudZoneListInternalServerError{}
}

/* GoogleCloudZoneListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GoogleCloudZoneListInternalServerError struct {
}

func (o *GoogleCloudZoneListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/GoogleCloud/zones][%d] googleCloudZoneListInternalServerError ", 500)
}

func (o *GoogleCloudZoneListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}