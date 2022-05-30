// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ImagesPersonalAzureImagesReader is a Reader for the ImagesPersonalAzureImages structure.
type ImagesPersonalAzureImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesPersonalAzureImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesPersonalAzureImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesPersonalAzureImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesPersonalAzureImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesPersonalAzureImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesPersonalAzureImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesPersonalAzureImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesPersonalAzureImagesOK creates a ImagesPersonalAzureImagesOK with default headers values
func NewImagesPersonalAzureImagesOK() *ImagesPersonalAzureImagesOK {
	return &ImagesPersonalAzureImagesOK{}
}

/* ImagesPersonalAzureImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesPersonalAzureImagesOK struct {
	Payload []*models.CommonStringBasedDropdownDto
}

func (o *ImagesPersonalAzureImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/personal/{cloudId}][%d] imagesPersonalAzureImagesOK  %+v", 200, o.Payload)
}
func (o *ImagesPersonalAzureImagesOK) GetPayload() []*models.CommonStringBasedDropdownDto {
	return o.Payload
}

func (o *ImagesPersonalAzureImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesPersonalAzureImagesBadRequest creates a ImagesPersonalAzureImagesBadRequest with default headers values
func NewImagesPersonalAzureImagesBadRequest() *ImagesPersonalAzureImagesBadRequest {
	return &ImagesPersonalAzureImagesBadRequest{}
}

/* ImagesPersonalAzureImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesPersonalAzureImagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ImagesPersonalAzureImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/personal/{cloudId}][%d] imagesPersonalAzureImagesBadRequest  %+v", 400, o.Payload)
}
func (o *ImagesPersonalAzureImagesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ImagesPersonalAzureImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesPersonalAzureImagesUnauthorized creates a ImagesPersonalAzureImagesUnauthorized with default headers values
func NewImagesPersonalAzureImagesUnauthorized() *ImagesPersonalAzureImagesUnauthorized {
	return &ImagesPersonalAzureImagesUnauthorized{}
}

/* ImagesPersonalAzureImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesPersonalAzureImagesUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ImagesPersonalAzureImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/personal/{cloudId}][%d] imagesPersonalAzureImagesUnauthorized  %+v", 401, o.Payload)
}
func (o *ImagesPersonalAzureImagesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ImagesPersonalAzureImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesPersonalAzureImagesForbidden creates a ImagesPersonalAzureImagesForbidden with default headers values
func NewImagesPersonalAzureImagesForbidden() *ImagesPersonalAzureImagesForbidden {
	return &ImagesPersonalAzureImagesForbidden{}
}

/* ImagesPersonalAzureImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesPersonalAzureImagesForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ImagesPersonalAzureImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/personal/{cloudId}][%d] imagesPersonalAzureImagesForbidden  %+v", 403, o.Payload)
}
func (o *ImagesPersonalAzureImagesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ImagesPersonalAzureImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesPersonalAzureImagesNotFound creates a ImagesPersonalAzureImagesNotFound with default headers values
func NewImagesPersonalAzureImagesNotFound() *ImagesPersonalAzureImagesNotFound {
	return &ImagesPersonalAzureImagesNotFound{}
}

/* ImagesPersonalAzureImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesPersonalAzureImagesNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ImagesPersonalAzureImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/personal/{cloudId}][%d] imagesPersonalAzureImagesNotFound  %+v", 404, o.Payload)
}
func (o *ImagesPersonalAzureImagesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ImagesPersonalAzureImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesPersonalAzureImagesInternalServerError creates a ImagesPersonalAzureImagesInternalServerError with default headers values
func NewImagesPersonalAzureImagesInternalServerError() *ImagesPersonalAzureImagesInternalServerError {
	return &ImagesPersonalAzureImagesInternalServerError{}
}

/* ImagesPersonalAzureImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesPersonalAzureImagesInternalServerError struct {
}

func (o *ImagesPersonalAzureImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/personal/{cloudId}][%d] imagesPersonalAzureImagesInternalServerError ", 500)
}

func (o *ImagesPersonalAzureImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}