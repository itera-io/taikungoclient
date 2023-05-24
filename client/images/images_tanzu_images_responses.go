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

// ImagesTanzuImagesReader is a Reader for the ImagesTanzuImages structure.
type ImagesTanzuImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesTanzuImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesTanzuImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesTanzuImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesTanzuImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesTanzuImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesTanzuImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesTanzuImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesTanzuImagesOK creates a ImagesTanzuImagesOK with default headers values
func NewImagesTanzuImagesOK() *ImagesTanzuImagesOK {
	return &ImagesTanzuImagesOK{}
}

/*
ImagesTanzuImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesTanzuImagesOK struct {
	Payload *models.TanzuImageList
}

// IsSuccess returns true when this images tanzu images o k response has a 2xx status code
func (o *ImagesTanzuImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images tanzu images o k response has a 3xx status code
func (o *ImagesTanzuImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images tanzu images o k response has a 4xx status code
func (o *ImagesTanzuImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images tanzu images o k response has a 5xx status code
func (o *ImagesTanzuImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images tanzu images o k response a status code equal to that given
func (o *ImagesTanzuImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the images tanzu images o k response
func (o *ImagesTanzuImagesOK) Code() int {
	return 200
}

func (o *ImagesTanzuImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesTanzuImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesTanzuImagesOK) GetPayload() *models.TanzuImageList {
	return o.Payload
}

func (o *ImagesTanzuImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TanzuImageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesTanzuImagesBadRequest creates a ImagesTanzuImagesBadRequest with default headers values
func NewImagesTanzuImagesBadRequest() *ImagesTanzuImagesBadRequest {
	return &ImagesTanzuImagesBadRequest{}
}

/*
ImagesTanzuImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesTanzuImagesBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this images tanzu images bad request response has a 2xx status code
func (o *ImagesTanzuImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images tanzu images bad request response has a 3xx status code
func (o *ImagesTanzuImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images tanzu images bad request response has a 4xx status code
func (o *ImagesTanzuImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images tanzu images bad request response has a 5xx status code
func (o *ImagesTanzuImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images tanzu images bad request response a status code equal to that given
func (o *ImagesTanzuImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the images tanzu images bad request response
func (o *ImagesTanzuImagesBadRequest) Code() int {
	return 400
}

func (o *ImagesTanzuImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesTanzuImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesTanzuImagesBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesTanzuImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesTanzuImagesUnauthorized creates a ImagesTanzuImagesUnauthorized with default headers values
func NewImagesTanzuImagesUnauthorized() *ImagesTanzuImagesUnauthorized {
	return &ImagesTanzuImagesUnauthorized{}
}

/*
ImagesTanzuImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesTanzuImagesUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this images tanzu images unauthorized response has a 2xx status code
func (o *ImagesTanzuImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images tanzu images unauthorized response has a 3xx status code
func (o *ImagesTanzuImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images tanzu images unauthorized response has a 4xx status code
func (o *ImagesTanzuImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images tanzu images unauthorized response has a 5xx status code
func (o *ImagesTanzuImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images tanzu images unauthorized response a status code equal to that given
func (o *ImagesTanzuImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the images tanzu images unauthorized response
func (o *ImagesTanzuImagesUnauthorized) Code() int {
	return 401
}

func (o *ImagesTanzuImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesTanzuImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesTanzuImagesUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesTanzuImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesTanzuImagesForbidden creates a ImagesTanzuImagesForbidden with default headers values
func NewImagesTanzuImagesForbidden() *ImagesTanzuImagesForbidden {
	return &ImagesTanzuImagesForbidden{}
}

/*
ImagesTanzuImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesTanzuImagesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this images tanzu images forbidden response has a 2xx status code
func (o *ImagesTanzuImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images tanzu images forbidden response has a 3xx status code
func (o *ImagesTanzuImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images tanzu images forbidden response has a 4xx status code
func (o *ImagesTanzuImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images tanzu images forbidden response has a 5xx status code
func (o *ImagesTanzuImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images tanzu images forbidden response a status code equal to that given
func (o *ImagesTanzuImagesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the images tanzu images forbidden response
func (o *ImagesTanzuImagesForbidden) Code() int {
	return 403
}

func (o *ImagesTanzuImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesTanzuImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesTanzuImagesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesTanzuImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesTanzuImagesNotFound creates a ImagesTanzuImagesNotFound with default headers values
func NewImagesTanzuImagesNotFound() *ImagesTanzuImagesNotFound {
	return &ImagesTanzuImagesNotFound{}
}

/*
ImagesTanzuImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesTanzuImagesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this images tanzu images not found response has a 2xx status code
func (o *ImagesTanzuImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images tanzu images not found response has a 3xx status code
func (o *ImagesTanzuImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images tanzu images not found response has a 4xx status code
func (o *ImagesTanzuImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images tanzu images not found response has a 5xx status code
func (o *ImagesTanzuImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images tanzu images not found response a status code equal to that given
func (o *ImagesTanzuImagesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the images tanzu images not found response
func (o *ImagesTanzuImagesNotFound) Code() int {
	return 404
}

func (o *ImagesTanzuImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesTanzuImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesTanzuImagesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesTanzuImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesTanzuImagesInternalServerError creates a ImagesTanzuImagesInternalServerError with default headers values
func NewImagesTanzuImagesInternalServerError() *ImagesTanzuImagesInternalServerError {
	return &ImagesTanzuImagesInternalServerError{}
}

/*
ImagesTanzuImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesTanzuImagesInternalServerError struct {
}

// IsSuccess returns true when this images tanzu images internal server error response has a 2xx status code
func (o *ImagesTanzuImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images tanzu images internal server error response has a 3xx status code
func (o *ImagesTanzuImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images tanzu images internal server error response has a 4xx status code
func (o *ImagesTanzuImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images tanzu images internal server error response has a 5xx status code
func (o *ImagesTanzuImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images tanzu images internal server error response a status code equal to that given
func (o *ImagesTanzuImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the images tanzu images internal server error response
func (o *ImagesTanzuImagesInternalServerError) Code() int {
	return 500
}

func (o *ImagesTanzuImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesInternalServerError ", 500)
}

func (o *ImagesTanzuImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/tanzu/{cloudId}][%d] imagesTanzuImagesInternalServerError ", 500)
}

func (o *ImagesTanzuImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}