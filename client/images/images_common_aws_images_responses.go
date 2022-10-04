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

// ImagesCommonAwsImagesReader is a Reader for the ImagesCommonAwsImages structure.
type ImagesCommonAwsImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesCommonAwsImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesCommonAwsImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesCommonAwsImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesCommonAwsImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesCommonAwsImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesCommonAwsImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesCommonAwsImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesCommonAwsImagesOK creates a ImagesCommonAwsImagesOK with default headers values
func NewImagesCommonAwsImagesOK() *ImagesCommonAwsImagesOK {
	return &ImagesCommonAwsImagesOK{}
}

/*
ImagesCommonAwsImagesOK describes a response with status code 200, with default header values.

Success
*/
type ImagesCommonAwsImagesOK struct {
	Payload []*models.AwsOwnerDetails
}

// IsSuccess returns true when this images common aws images o k response has a 2xx status code
func (o *ImagesCommonAwsImagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images common aws images o k response has a 3xx status code
func (o *ImagesCommonAwsImagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common aws images o k response has a 4xx status code
func (o *ImagesCommonAwsImagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common aws images o k response has a 5xx status code
func (o *ImagesCommonAwsImagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images common aws images o k response a status code equal to that given
func (o *ImagesCommonAwsImagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesCommonAwsImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonAwsImagesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesOK  %+v", 200, o.Payload)
}

func (o *ImagesCommonAwsImagesOK) GetPayload() []*models.AwsOwnerDetails {
	return o.Payload
}

func (o *ImagesCommonAwsImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAwsImagesBadRequest creates a ImagesCommonAwsImagesBadRequest with default headers values
func NewImagesCommonAwsImagesBadRequest() *ImagesCommonAwsImagesBadRequest {
	return &ImagesCommonAwsImagesBadRequest{}
}

/*
ImagesCommonAwsImagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesCommonAwsImagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this images common aws images bad request response has a 2xx status code
func (o *ImagesCommonAwsImagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common aws images bad request response has a 3xx status code
func (o *ImagesCommonAwsImagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common aws images bad request response has a 4xx status code
func (o *ImagesCommonAwsImagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common aws images bad request response has a 5xx status code
func (o *ImagesCommonAwsImagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images common aws images bad request response a status code equal to that given
func (o *ImagesCommonAwsImagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesCommonAwsImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonAwsImagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesCommonAwsImagesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ImagesCommonAwsImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAwsImagesUnauthorized creates a ImagesCommonAwsImagesUnauthorized with default headers values
func NewImagesCommonAwsImagesUnauthorized() *ImagesCommonAwsImagesUnauthorized {
	return &ImagesCommonAwsImagesUnauthorized{}
}

/*
ImagesCommonAwsImagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesCommonAwsImagesUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images common aws images unauthorized response has a 2xx status code
func (o *ImagesCommonAwsImagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common aws images unauthorized response has a 3xx status code
func (o *ImagesCommonAwsImagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common aws images unauthorized response has a 4xx status code
func (o *ImagesCommonAwsImagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common aws images unauthorized response has a 5xx status code
func (o *ImagesCommonAwsImagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images common aws images unauthorized response a status code equal to that given
func (o *ImagesCommonAwsImagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesCommonAwsImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonAwsImagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesCommonAwsImagesUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesCommonAwsImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAwsImagesForbidden creates a ImagesCommonAwsImagesForbidden with default headers values
func NewImagesCommonAwsImagesForbidden() *ImagesCommonAwsImagesForbidden {
	return &ImagesCommonAwsImagesForbidden{}
}

/*
ImagesCommonAwsImagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesCommonAwsImagesForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images common aws images forbidden response has a 2xx status code
func (o *ImagesCommonAwsImagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common aws images forbidden response has a 3xx status code
func (o *ImagesCommonAwsImagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common aws images forbidden response has a 4xx status code
func (o *ImagesCommonAwsImagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common aws images forbidden response has a 5xx status code
func (o *ImagesCommonAwsImagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images common aws images forbidden response a status code equal to that given
func (o *ImagesCommonAwsImagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesCommonAwsImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonAwsImagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesForbidden  %+v", 403, o.Payload)
}

func (o *ImagesCommonAwsImagesForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesCommonAwsImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAwsImagesNotFound creates a ImagesCommonAwsImagesNotFound with default headers values
func NewImagesCommonAwsImagesNotFound() *ImagesCommonAwsImagesNotFound {
	return &ImagesCommonAwsImagesNotFound{}
}

/*
ImagesCommonAwsImagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesCommonAwsImagesNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images common aws images not found response has a 2xx status code
func (o *ImagesCommonAwsImagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common aws images not found response has a 3xx status code
func (o *ImagesCommonAwsImagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common aws images not found response has a 4xx status code
func (o *ImagesCommonAwsImagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images common aws images not found response has a 5xx status code
func (o *ImagesCommonAwsImagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images common aws images not found response a status code equal to that given
func (o *ImagesCommonAwsImagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesCommonAwsImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonAwsImagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesNotFound  %+v", 404, o.Payload)
}

func (o *ImagesCommonAwsImagesNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesCommonAwsImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesCommonAwsImagesInternalServerError creates a ImagesCommonAwsImagesInternalServerError with default headers values
func NewImagesCommonAwsImagesInternalServerError() *ImagesCommonAwsImagesInternalServerError {
	return &ImagesCommonAwsImagesInternalServerError{}
}

/*
ImagesCommonAwsImagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesCommonAwsImagesInternalServerError struct {
}

// IsSuccess returns true when this images common aws images internal server error response has a 2xx status code
func (o *ImagesCommonAwsImagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images common aws images internal server error response has a 3xx status code
func (o *ImagesCommonAwsImagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images common aws images internal server error response has a 4xx status code
func (o *ImagesCommonAwsImagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images common aws images internal server error response has a 5xx status code
func (o *ImagesCommonAwsImagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images common aws images internal server error response a status code equal to that given
func (o *ImagesCommonAwsImagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesCommonAwsImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesInternalServerError ", 500)
}

func (o *ImagesCommonAwsImagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/aws/common/{cloudId}][%d] imagesCommonAwsImagesInternalServerError ", 500)
}

func (o *ImagesCommonAwsImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
