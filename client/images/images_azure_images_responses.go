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

// ImagesAzureImagesReader is a Reader for the ImagesAzureImages structure.
type ImagesAzureImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesAzureImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesAzureImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesAzureImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesAzureImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesAzureImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesAzureImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesAzureImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesAzureImagesOK creates a ImagesAzureImagesOK with default headers values
func NewImagesAzureImagesOK() *ImagesAzureImagesOK {
	return &ImagesAzureImagesOK{}
}

/*
ImagesAzureImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesAzureImagesOK struct {
	Payload *models.AzureImageList
}

// IsSuccess returns true when this images azure images o k response has a 2xx status code
func (o *ImagesAzureImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images azure images o k response has a 3xx status code
func (o *ImagesAzureImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images azure images o k response has a 4xx status code
func (o *ImagesAzureImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images azure images o k response has a 5xx status code
func (o *ImagesAzureImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images azure images o k response a status code equal to that given
func (o *ImagesAzureImagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesAzureImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesAzureImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesAzureImagesOK) GetPayload() *models.AzureImageList {
	return o.Payload
}

func (o *ImagesAzureImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureImageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAzureImagesBadRequest creates a ImagesAzureImagesBadRequest with default headers values
func NewImagesAzureImagesBadRequest() *ImagesAzureImagesBadRequest {
	return &ImagesAzureImagesBadRequest{}
}

/*
ImagesAzureImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesAzureImagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this images azure images bad request response has a 2xx status code
func (o *ImagesAzureImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images azure images bad request response has a 3xx status code
func (o *ImagesAzureImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images azure images bad request response has a 4xx status code
func (o *ImagesAzureImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images azure images bad request response has a 5xx status code
func (o *ImagesAzureImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images azure images bad request response a status code equal to that given
func (o *ImagesAzureImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesAzureImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesAzureImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesAzureImagesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ImagesAzureImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAzureImagesUnauthorized creates a ImagesAzureImagesUnauthorized with default headers values
func NewImagesAzureImagesUnauthorized() *ImagesAzureImagesUnauthorized {
	return &ImagesAzureImagesUnauthorized{}
}

/*
ImagesAzureImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesAzureImagesUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this images azure images unauthorized response has a 2xx status code
func (o *ImagesAzureImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images azure images unauthorized response has a 3xx status code
func (o *ImagesAzureImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images azure images unauthorized response has a 4xx status code
func (o *ImagesAzureImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images azure images unauthorized response has a 5xx status code
func (o *ImagesAzureImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images azure images unauthorized response a status code equal to that given
func (o *ImagesAzureImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesAzureImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesAzureImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesAzureImagesUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAzureImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAzureImagesForbidden creates a ImagesAzureImagesForbidden with default headers values
func NewImagesAzureImagesForbidden() *ImagesAzureImagesForbidden {
	return &ImagesAzureImagesForbidden{}
}

/*
ImagesAzureImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesAzureImagesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this images azure images forbidden response has a 2xx status code
func (o *ImagesAzureImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images azure images forbidden response has a 3xx status code
func (o *ImagesAzureImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images azure images forbidden response has a 4xx status code
func (o *ImagesAzureImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images azure images forbidden response has a 5xx status code
func (o *ImagesAzureImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images azure images forbidden response a status code equal to that given
func (o *ImagesAzureImagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesAzureImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesAzureImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesAzureImagesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAzureImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAzureImagesNotFound creates a ImagesAzureImagesNotFound with default headers values
func NewImagesAzureImagesNotFound() *ImagesAzureImagesNotFound {
	return &ImagesAzureImagesNotFound{}
}

/*
ImagesAzureImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesAzureImagesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this images azure images not found response has a 2xx status code
func (o *ImagesAzureImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images azure images not found response has a 3xx status code
func (o *ImagesAzureImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images azure images not found response has a 4xx status code
func (o *ImagesAzureImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images azure images not found response has a 5xx status code
func (o *ImagesAzureImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images azure images not found response a status code equal to that given
func (o *ImagesAzureImagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesAzureImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesAzureImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesAzureImagesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAzureImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAzureImagesInternalServerError creates a ImagesAzureImagesInternalServerError with default headers values
func NewImagesAzureImagesInternalServerError() *ImagesAzureImagesInternalServerError {
	return &ImagesAzureImagesInternalServerError{}
}

/*
ImagesAzureImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesAzureImagesInternalServerError struct {
}

// IsSuccess returns true when this images azure images internal server error response has a 2xx status code
func (o *ImagesAzureImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images azure images internal server error response has a 3xx status code
func (o *ImagesAzureImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images azure images internal server error response has a 4xx status code
func (o *ImagesAzureImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images azure images internal server error response has a 5xx status code
func (o *ImagesAzureImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images azure images internal server error response a status code equal to that given
func (o *ImagesAzureImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesAzureImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesInternalServerError ", 500)
}

func (o *ImagesAzureImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}][%d] imagesAzureImagesInternalServerError ", 500)
}

func (o *ImagesAzureImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
