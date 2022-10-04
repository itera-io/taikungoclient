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

// ImagesOpenstackImagesReader is a Reader for the ImagesOpenstackImages structure.
type ImagesOpenstackImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesOpenstackImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesOpenstackImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesOpenstackImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesOpenstackImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesOpenstackImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesOpenstackImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesOpenstackImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesOpenstackImagesOK creates a ImagesOpenstackImagesOK with default headers values
func NewImagesOpenstackImagesOK() *ImagesOpenstackImagesOK {
	return &ImagesOpenstackImagesOK{}
}

/*
ImagesOpenstackImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesOpenstackImagesOK struct {
	Payload *models.OpenstackImageList
}

// IsSuccess returns true when this images openstack images o k response has a 2xx status code
func (o *ImagesOpenstackImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images openstack images o k response has a 3xx status code
func (o *ImagesOpenstackImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images openstack images o k response has a 4xx status code
func (o *ImagesOpenstackImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images openstack images o k response has a 5xx status code
func (o *ImagesOpenstackImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images openstack images o k response a status code equal to that given
func (o *ImagesOpenstackImagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesOpenstackImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesOpenstackImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesOpenstackImagesOK) GetPayload() *models.OpenstackImageList {
	return o.Payload
}

func (o *ImagesOpenstackImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenstackImageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesOpenstackImagesBadRequest creates a ImagesOpenstackImagesBadRequest with default headers values
func NewImagesOpenstackImagesBadRequest() *ImagesOpenstackImagesBadRequest {
	return &ImagesOpenstackImagesBadRequest{}
}

/*
ImagesOpenstackImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesOpenstackImagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this images openstack images bad request response has a 2xx status code
func (o *ImagesOpenstackImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images openstack images bad request response has a 3xx status code
func (o *ImagesOpenstackImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images openstack images bad request response has a 4xx status code
func (o *ImagesOpenstackImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images openstack images bad request response has a 5xx status code
func (o *ImagesOpenstackImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images openstack images bad request response a status code equal to that given
func (o *ImagesOpenstackImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesOpenstackImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesOpenstackImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesOpenstackImagesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ImagesOpenstackImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesOpenstackImagesUnauthorized creates a ImagesOpenstackImagesUnauthorized with default headers values
func NewImagesOpenstackImagesUnauthorized() *ImagesOpenstackImagesUnauthorized {
	return &ImagesOpenstackImagesUnauthorized{}
}

/*
ImagesOpenstackImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesOpenstackImagesUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images openstack images unauthorized response has a 2xx status code
func (o *ImagesOpenstackImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images openstack images unauthorized response has a 3xx status code
func (o *ImagesOpenstackImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images openstack images unauthorized response has a 4xx status code
func (o *ImagesOpenstackImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images openstack images unauthorized response has a 5xx status code
func (o *ImagesOpenstackImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images openstack images unauthorized response a status code equal to that given
func (o *ImagesOpenstackImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesOpenstackImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesOpenstackImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesOpenstackImagesUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesOpenstackImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesOpenstackImagesForbidden creates a ImagesOpenstackImagesForbidden with default headers values
func NewImagesOpenstackImagesForbidden() *ImagesOpenstackImagesForbidden {
	return &ImagesOpenstackImagesForbidden{}
}

/*
ImagesOpenstackImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesOpenstackImagesForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images openstack images forbidden response has a 2xx status code
func (o *ImagesOpenstackImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images openstack images forbidden response has a 3xx status code
func (o *ImagesOpenstackImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images openstack images forbidden response has a 4xx status code
func (o *ImagesOpenstackImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images openstack images forbidden response has a 5xx status code
func (o *ImagesOpenstackImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images openstack images forbidden response a status code equal to that given
func (o *ImagesOpenstackImagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesOpenstackImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesOpenstackImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesOpenstackImagesForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesOpenstackImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesOpenstackImagesNotFound creates a ImagesOpenstackImagesNotFound with default headers values
func NewImagesOpenstackImagesNotFound() *ImagesOpenstackImagesNotFound {
	return &ImagesOpenstackImagesNotFound{}
}

/*
ImagesOpenstackImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesOpenstackImagesNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images openstack images not found response has a 2xx status code
func (o *ImagesOpenstackImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images openstack images not found response has a 3xx status code
func (o *ImagesOpenstackImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images openstack images not found response has a 4xx status code
func (o *ImagesOpenstackImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images openstack images not found response has a 5xx status code
func (o *ImagesOpenstackImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images openstack images not found response a status code equal to that given
func (o *ImagesOpenstackImagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesOpenstackImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesOpenstackImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesOpenstackImagesNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesOpenstackImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesOpenstackImagesInternalServerError creates a ImagesOpenstackImagesInternalServerError with default headers values
func NewImagesOpenstackImagesInternalServerError() *ImagesOpenstackImagesInternalServerError {
	return &ImagesOpenstackImagesInternalServerError{}
}

/*
ImagesOpenstackImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesOpenstackImagesInternalServerError struct {
}

// IsSuccess returns true when this images openstack images internal server error response has a 2xx status code
func (o *ImagesOpenstackImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images openstack images internal server error response has a 3xx status code
func (o *ImagesOpenstackImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images openstack images internal server error response has a 4xx status code
func (o *ImagesOpenstackImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images openstack images internal server error response has a 5xx status code
func (o *ImagesOpenstackImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images openstack images internal server error response a status code equal to that given
func (o *ImagesOpenstackImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesOpenstackImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesInternalServerError ", 500)
}

func (o *ImagesOpenstackImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/openstack/{cloudId}][%d] imagesOpenstackImagesInternalServerError ", 500)
}

func (o *ImagesOpenstackImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
