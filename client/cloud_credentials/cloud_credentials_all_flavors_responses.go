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

// CloudCredentialsAllFlavorsReader is a Reader for the CloudCredentialsAllFlavors structure.
type CloudCredentialsAllFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudCredentialsAllFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudCredentialsAllFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudCredentialsAllFlavorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCloudCredentialsAllFlavorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudCredentialsAllFlavorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloudCredentialsAllFlavorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCloudCredentialsAllFlavorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudCredentialsAllFlavorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloudCredentialsAllFlavorsOK creates a CloudCredentialsAllFlavorsOK with default headers values
func NewCloudCredentialsAllFlavorsOK() *CloudCredentialsAllFlavorsOK {
	return &CloudCredentialsAllFlavorsOK{}
}

/* CloudCredentialsAllFlavorsOK describes a response with status code 200, with default header values.

Success
*/
type CloudCredentialsAllFlavorsOK struct {
	Payload *models.AllFlavorsList
}

func (o *CloudCredentialsAllFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsOK  %+v", 200, o.Payload)
}
func (o *CloudCredentialsAllFlavorsOK) GetPayload() *models.AllFlavorsList {
	return o.Payload
}

func (o *CloudCredentialsAllFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllFlavorsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsAllFlavorsBadRequest creates a CloudCredentialsAllFlavorsBadRequest with default headers values
func NewCloudCredentialsAllFlavorsBadRequest() *CloudCredentialsAllFlavorsBadRequest {
	return &CloudCredentialsAllFlavorsBadRequest{}
}

/* CloudCredentialsAllFlavorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudCredentialsAllFlavorsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CloudCredentialsAllFlavorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsBadRequest  %+v", 400, o.Payload)
}
func (o *CloudCredentialsAllFlavorsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsAllFlavorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsAllFlavorsUnauthorized creates a CloudCredentialsAllFlavorsUnauthorized with default headers values
func NewCloudCredentialsAllFlavorsUnauthorized() *CloudCredentialsAllFlavorsUnauthorized {
	return &CloudCredentialsAllFlavorsUnauthorized{}
}

/* CloudCredentialsAllFlavorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CloudCredentialsAllFlavorsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsAllFlavorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsUnauthorized  %+v", 401, o.Payload)
}
func (o *CloudCredentialsAllFlavorsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsAllFlavorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsAllFlavorsForbidden creates a CloudCredentialsAllFlavorsForbidden with default headers values
func NewCloudCredentialsAllFlavorsForbidden() *CloudCredentialsAllFlavorsForbidden {
	return &CloudCredentialsAllFlavorsForbidden{}
}

/* CloudCredentialsAllFlavorsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudCredentialsAllFlavorsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsAllFlavorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsForbidden  %+v", 403, o.Payload)
}
func (o *CloudCredentialsAllFlavorsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsAllFlavorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsAllFlavorsNotFound creates a CloudCredentialsAllFlavorsNotFound with default headers values
func NewCloudCredentialsAllFlavorsNotFound() *CloudCredentialsAllFlavorsNotFound {
	return &CloudCredentialsAllFlavorsNotFound{}
}

/* CloudCredentialsAllFlavorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloudCredentialsAllFlavorsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsAllFlavorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsNotFound  %+v", 404, o.Payload)
}
func (o *CloudCredentialsAllFlavorsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsAllFlavorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsAllFlavorsTooManyRequests creates a CloudCredentialsAllFlavorsTooManyRequests with default headers values
func NewCloudCredentialsAllFlavorsTooManyRequests() *CloudCredentialsAllFlavorsTooManyRequests {
	return &CloudCredentialsAllFlavorsTooManyRequests{}
}

/* CloudCredentialsAllFlavorsTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type CloudCredentialsAllFlavorsTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *CloudCredentialsAllFlavorsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsTooManyRequests  %+v", 429, o.Payload)
}
func (o *CloudCredentialsAllFlavorsTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CloudCredentialsAllFlavorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudCredentialsAllFlavorsInternalServerError creates a CloudCredentialsAllFlavorsInternalServerError with default headers values
func NewCloudCredentialsAllFlavorsInternalServerError() *CloudCredentialsAllFlavorsInternalServerError {
	return &CloudCredentialsAllFlavorsInternalServerError{}
}

/* CloudCredentialsAllFlavorsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloudCredentialsAllFlavorsInternalServerError struct {
}

func (o *CloudCredentialsAllFlavorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/CloudCredentials/flavors/{cloudId}][%d] cloudCredentialsAllFlavorsInternalServerError ", 500)
}

func (o *CloudCredentialsAllFlavorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
