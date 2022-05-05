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

// ImagesCommonAzureImagesReader is a Reader for the ImagesCommonAzureImages structure.
type ImagesCommonAzureImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesCommonAzureImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesCommonAzureImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesCommonAzureImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesCommonAzureImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesCommonAzureImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesCommonAzureImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesCommonAzureImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesCommonAzureImagesOK creates a ImagesCommonAzureImagesOK with default headers values
func NewImagesCommonAzureImagesOK() *ImagesCommonAzureImagesOK {
	return &ImagesCommonAzureImagesOK{}
}

/* ImagesCommonAzureImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesCommonAzureImagesOK struct {
	Payload []*models.AzurePublisherDetails
}

func (o *ImagesCommonAzureImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesOK  %+v", 200, o.Payload)
}
func (o *ImagesCommonAzureImagesOK) GetPayload() []*models.AzurePublisherDetails {
	return o.Payload
}

func (o *ImagesCommonAzureImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesBadRequest creates a ImagesCommonAzureImagesBadRequest with default headers values
func NewImagesCommonAzureImagesBadRequest() *ImagesCommonAzureImagesBadRequest {
	return &ImagesCommonAzureImagesBadRequest{}
}

/* ImagesCommonAzureImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesCommonAzureImagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ImagesCommonAzureImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesBadRequest  %+v", 400, o.Payload)
}
func (o *ImagesCommonAzureImagesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ImagesCommonAzureImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesUnauthorized creates a ImagesCommonAzureImagesUnauthorized with default headers values
func NewImagesCommonAzureImagesUnauthorized() *ImagesCommonAzureImagesUnauthorized {
	return &ImagesCommonAzureImagesUnauthorized{}
}

/* ImagesCommonAzureImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesCommonAzureImagesUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ImagesCommonAzureImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesUnauthorized  %+v", 401, o.Payload)
}
func (o *ImagesCommonAzureImagesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ImagesCommonAzureImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesForbidden creates a ImagesCommonAzureImagesForbidden with default headers values
func NewImagesCommonAzureImagesForbidden() *ImagesCommonAzureImagesForbidden {
	return &ImagesCommonAzureImagesForbidden{}
}

/* ImagesCommonAzureImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesCommonAzureImagesForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ImagesCommonAzureImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesForbidden  %+v", 403, o.Payload)
}
func (o *ImagesCommonAzureImagesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ImagesCommonAzureImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesNotFound creates a ImagesCommonAzureImagesNotFound with default headers values
func NewImagesCommonAzureImagesNotFound() *ImagesCommonAzureImagesNotFound {
	return &ImagesCommonAzureImagesNotFound{}
}

/* ImagesCommonAzureImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesCommonAzureImagesNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ImagesCommonAzureImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesNotFound  %+v", 404, o.Payload)
}
func (o *ImagesCommonAzureImagesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ImagesCommonAzureImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesInternalServerError creates a ImagesCommonAzureImagesInternalServerError with default headers values
func NewImagesCommonAzureImagesInternalServerError() *ImagesCommonAzureImagesInternalServerError {
	return &ImagesCommonAzureImagesInternalServerError{}
}

/* ImagesCommonAzureImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesCommonAzureImagesInternalServerError struct {
}

func (o *ImagesCommonAzureImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesInternalServerError ", 500)
}

func (o *ImagesCommonAzureImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
