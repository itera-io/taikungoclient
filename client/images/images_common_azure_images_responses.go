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

/*
ImagesCommonAzureImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesCommonAzureImagesOK struct {
	Payload []*models.AzurePublisherDetails
}

// IsSuccess returns true when this images common azure images o k response has a 2xx status code
func (o *ImagesCommonAzureImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images common azure images o k response has a 3xx status code
func (o *ImagesCommonAzureImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images o k response has a 4xx status code
func (o *ImagesCommonAzureImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common azure images o k response has a 5xx status code
func (o *ImagesCommonAzureImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images o k response a status code equal to that given
func (o *ImagesCommonAzureImagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the images common azure images o k response
func (o *ImagesCommonAzureImagesOK) Code() int {
	return 200
}

func (o *ImagesCommonAzureImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonAzureImagesOK) String() string {
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

/*
ImagesCommonAzureImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesCommonAzureImagesBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this images common azure images bad request response has a 2xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images bad request response has a 3xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images bad request response has a 4xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images bad request response has a 5xx status code
func (o *ImagesCommonAzureImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images bad request response a status code equal to that given
func (o *ImagesCommonAzureImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the images common azure images bad request response
func (o *ImagesCommonAzureImagesBadRequest) Code() int {
	return 400
}

func (o *ImagesCommonAzureImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonAzureImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonAzureImagesBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesCommonAzureImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesUnauthorized creates a ImagesCommonAzureImagesUnauthorized with default headers values
func NewImagesCommonAzureImagesUnauthorized() *ImagesCommonAzureImagesUnauthorized {
	return &ImagesCommonAzureImagesUnauthorized{}
}

/*
ImagesCommonAzureImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesCommonAzureImagesUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this images common azure images unauthorized response has a 2xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images unauthorized response has a 3xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images unauthorized response has a 4xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images unauthorized response has a 5xx status code
func (o *ImagesCommonAzureImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images unauthorized response a status code equal to that given
func (o *ImagesCommonAzureImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the images common azure images unauthorized response
func (o *ImagesCommonAzureImagesUnauthorized) Code() int {
	return 401
}

func (o *ImagesCommonAzureImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonAzureImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonAzureImagesUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesCommonAzureImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesForbidden creates a ImagesCommonAzureImagesForbidden with default headers values
func NewImagesCommonAzureImagesForbidden() *ImagesCommonAzureImagesForbidden {
	return &ImagesCommonAzureImagesForbidden{}
}

/*
ImagesCommonAzureImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesCommonAzureImagesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this images common azure images forbidden response has a 2xx status code
func (o *ImagesCommonAzureImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images forbidden response has a 3xx status code
func (o *ImagesCommonAzureImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images forbidden response has a 4xx status code
func (o *ImagesCommonAzureImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images forbidden response has a 5xx status code
func (o *ImagesCommonAzureImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images forbidden response a status code equal to that given
func (o *ImagesCommonAzureImagesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the images common azure images forbidden response
func (o *ImagesCommonAzureImagesForbidden) Code() int {
	return 403
}

func (o *ImagesCommonAzureImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonAzureImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonAzureImagesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesCommonAzureImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesNotFound creates a ImagesCommonAzureImagesNotFound with default headers values
func NewImagesCommonAzureImagesNotFound() *ImagesCommonAzureImagesNotFound {
	return &ImagesCommonAzureImagesNotFound{}
}

/*
ImagesCommonAzureImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesCommonAzureImagesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this images common azure images not found response has a 2xx status code
func (o *ImagesCommonAzureImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images not found response has a 3xx status code
func (o *ImagesCommonAzureImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images not found response has a 4xx status code
func (o *ImagesCommonAzureImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common azure images not found response has a 5xx status code
func (o *ImagesCommonAzureImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images common azure images not found response a status code equal to that given
func (o *ImagesCommonAzureImagesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the images common azure images not found response
func (o *ImagesCommonAzureImagesNotFound) Code() int {
	return 404
}

func (o *ImagesCommonAzureImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonAzureImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonAzureImagesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesCommonAzureImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAzureImagesInternalServerError creates a ImagesCommonAzureImagesInternalServerError with default headers values
func NewImagesCommonAzureImagesInternalServerError() *ImagesCommonAzureImagesInternalServerError {
	return &ImagesCommonAzureImagesInternalServerError{}
}

/*
ImagesCommonAzureImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesCommonAzureImagesInternalServerError struct {
}

// IsSuccess returns true when this images common azure images internal server error response has a 2xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common azure images internal server error response has a 3xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common azure images internal server error response has a 4xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common azure images internal server error response has a 5xx status code
func (o *ImagesCommonAzureImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images common azure images internal server error response a status code equal to that given
func (o *ImagesCommonAzureImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the images common azure images internal server error response
func (o *ImagesCommonAzureImagesInternalServerError) Code() int {
	return 500
}

func (o *ImagesCommonAzureImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesInternalServerError ", 500)
}

func (o *ImagesCommonAzureImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/azure/common/{cloudId}][%d] imagesCommonAzureImagesInternalServerError ", 500)
}

func (o *ImagesCommonAzureImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
