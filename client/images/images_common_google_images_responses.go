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

// ImagesCommonGoogleImagesReader is a Reader for the ImagesCommonGoogleImages structure.
type ImagesCommonGoogleImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesCommonGoogleImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesCommonGoogleImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesCommonGoogleImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesCommonGoogleImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesCommonGoogleImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesCommonGoogleImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesCommonGoogleImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesCommonGoogleImagesOK creates a ImagesCommonGoogleImagesOK with default headers values
func NewImagesCommonGoogleImagesOK() *ImagesCommonGoogleImagesOK {
	return &ImagesCommonGoogleImagesOK{}
}

/*
ImagesCommonGoogleImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesCommonGoogleImagesOK struct {
	Payload []*models.GoogleOwnerDetails
}

// IsSuccess returns true when this images common google images o k response has a 2xx status code
func (o *ImagesCommonGoogleImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images common google images o k response has a 3xx status code
func (o *ImagesCommonGoogleImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common google images o k response has a 4xx status code
func (o *ImagesCommonGoogleImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common google images o k response has a 5xx status code
func (o *ImagesCommonGoogleImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images common google images o k response a status code equal to that given
func (o *ImagesCommonGoogleImagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesCommonGoogleImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonGoogleImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonGoogleImagesOK) GetPayload() []*models.GoogleOwnerDetails {
	return o.Payload
}

func (o *ImagesCommonGoogleImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonGoogleImagesBadRequest creates a ImagesCommonGoogleImagesBadRequest with default headers values
func NewImagesCommonGoogleImagesBadRequest() *ImagesCommonGoogleImagesBadRequest {
	return &ImagesCommonGoogleImagesBadRequest{}
}

/*
ImagesCommonGoogleImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesCommonGoogleImagesBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this images common google images bad request response has a 2xx status code
func (o *ImagesCommonGoogleImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common google images bad request response has a 3xx status code
func (o *ImagesCommonGoogleImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common google images bad request response has a 4xx status code
func (o *ImagesCommonGoogleImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common google images bad request response has a 5xx status code
func (o *ImagesCommonGoogleImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images common google images bad request response a status code equal to that given
func (o *ImagesCommonGoogleImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesCommonGoogleImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonGoogleImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonGoogleImagesBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ImagesCommonGoogleImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonGoogleImagesUnauthorized creates a ImagesCommonGoogleImagesUnauthorized with default headers values
func NewImagesCommonGoogleImagesUnauthorized() *ImagesCommonGoogleImagesUnauthorized {
	return &ImagesCommonGoogleImagesUnauthorized{}
}

/*
ImagesCommonGoogleImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesCommonGoogleImagesUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images common google images unauthorized response has a 2xx status code
func (o *ImagesCommonGoogleImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common google images unauthorized response has a 3xx status code
func (o *ImagesCommonGoogleImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common google images unauthorized response has a 4xx status code
func (o *ImagesCommonGoogleImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common google images unauthorized response has a 5xx status code
func (o *ImagesCommonGoogleImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images common google images unauthorized response a status code equal to that given
func (o *ImagesCommonGoogleImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesCommonGoogleImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonGoogleImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonGoogleImagesUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesCommonGoogleImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonGoogleImagesForbidden creates a ImagesCommonGoogleImagesForbidden with default headers values
func NewImagesCommonGoogleImagesForbidden() *ImagesCommonGoogleImagesForbidden {
	return &ImagesCommonGoogleImagesForbidden{}
}

/*
ImagesCommonGoogleImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesCommonGoogleImagesForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images common google images forbidden response has a 2xx status code
func (o *ImagesCommonGoogleImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common google images forbidden response has a 3xx status code
func (o *ImagesCommonGoogleImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common google images forbidden response has a 4xx status code
func (o *ImagesCommonGoogleImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common google images forbidden response has a 5xx status code
func (o *ImagesCommonGoogleImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images common google images forbidden response a status code equal to that given
func (o *ImagesCommonGoogleImagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesCommonGoogleImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonGoogleImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonGoogleImagesForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesCommonGoogleImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonGoogleImagesNotFound creates a ImagesCommonGoogleImagesNotFound with default headers values
func NewImagesCommonGoogleImagesNotFound() *ImagesCommonGoogleImagesNotFound {
	return &ImagesCommonGoogleImagesNotFound{}
}

/*
ImagesCommonGoogleImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesCommonGoogleImagesNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images common google images not found response has a 2xx status code
func (o *ImagesCommonGoogleImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common google images not found response has a 3xx status code
func (o *ImagesCommonGoogleImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common google images not found response has a 4xx status code
func (o *ImagesCommonGoogleImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common google images not found response has a 5xx status code
func (o *ImagesCommonGoogleImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images common google images not found response a status code equal to that given
func (o *ImagesCommonGoogleImagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesCommonGoogleImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonGoogleImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonGoogleImagesNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesCommonGoogleImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonGoogleImagesInternalServerError creates a ImagesCommonGoogleImagesInternalServerError with default headers values
func NewImagesCommonGoogleImagesInternalServerError() *ImagesCommonGoogleImagesInternalServerError {
	return &ImagesCommonGoogleImagesInternalServerError{}
}

/*
ImagesCommonGoogleImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesCommonGoogleImagesInternalServerError struct {
}

// IsSuccess returns true when this images common google images internal server error response has a 2xx status code
func (o *ImagesCommonGoogleImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common google images internal server error response has a 3xx status code
func (o *ImagesCommonGoogleImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common google images internal server error response has a 4xx status code
func (o *ImagesCommonGoogleImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common google images internal server error response has a 5xx status code
func (o *ImagesCommonGoogleImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images common google images internal server error response a status code equal to that given
func (o *ImagesCommonGoogleImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesCommonGoogleImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesInternalServerError ", 500)
}

func (o *ImagesCommonGoogleImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/google/common/{cloudId}][%d] imagesCommonGoogleImagesInternalServerError ", 500)
}

func (o *ImagesCommonGoogleImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
