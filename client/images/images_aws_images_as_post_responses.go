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

// ImagesAwsImagesAsPostReader is a Reader for the ImagesAwsImagesAsPost structure.
type ImagesAwsImagesAsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesAwsImagesAsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesAwsImagesAsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesAwsImagesAsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesAwsImagesAsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesAwsImagesAsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesAwsImagesAsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesAwsImagesAsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesAwsImagesAsPostOK creates a ImagesAwsImagesAsPostOK with default headers values
func NewImagesAwsImagesAsPostOK() *ImagesAwsImagesAsPostOK {
	return &ImagesAwsImagesAsPostOK{}
}

/*
ImagesAwsImagesAsPostOK describes a response with status code 200, with default header values.

Success
*/
type ImagesAwsImagesAsPostOK struct {
	Payload *models.AwsImagesPostList
}

// IsSuccess returns true when this images aws images as post o k response has a 2xx status code
func (o *ImagesAwsImagesAsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images aws images as post o k response has a 3xx status code
func (o *ImagesAwsImagesAsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images aws images as post o k response has a 4xx status code
func (o *ImagesAwsImagesAsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images aws images as post o k response has a 5xx status code
func (o *ImagesAwsImagesAsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images aws images as post o k response a status code equal to that given
func (o *ImagesAwsImagesAsPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the images aws images as post o k response
func (o *ImagesAwsImagesAsPostOK) Code() int {
	return 200
}

func (o *ImagesAwsImagesAsPostOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostOK  %+v", 200, o.Payload)
}

func (o *ImagesAwsImagesAsPostOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostOK  %+v", 200, o.Payload)
}

func (o *ImagesAwsImagesAsPostOK) GetPayload() *models.AwsImagesPostList {
	return o.Payload
}

func (o *ImagesAwsImagesAsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsImagesPostList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAwsImagesAsPostBadRequest creates a ImagesAwsImagesAsPostBadRequest with default headers values
func NewImagesAwsImagesAsPostBadRequest() *ImagesAwsImagesAsPostBadRequest {
	return &ImagesAwsImagesAsPostBadRequest{}
}

/*
ImagesAwsImagesAsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesAwsImagesAsPostBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this images aws images as post bad request response has a 2xx status code
func (o *ImagesAwsImagesAsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images aws images as post bad request response has a 3xx status code
func (o *ImagesAwsImagesAsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images aws images as post bad request response has a 4xx status code
func (o *ImagesAwsImagesAsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images aws images as post bad request response has a 5xx status code
func (o *ImagesAwsImagesAsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images aws images as post bad request response a status code equal to that given
func (o *ImagesAwsImagesAsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the images aws images as post bad request response
func (o *ImagesAwsImagesAsPostBadRequest) Code() int {
	return 400
}

func (o *ImagesAwsImagesAsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesAwsImagesAsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesAwsImagesAsPostBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAwsImagesAsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAwsImagesAsPostUnauthorized creates a ImagesAwsImagesAsPostUnauthorized with default headers values
func NewImagesAwsImagesAsPostUnauthorized() *ImagesAwsImagesAsPostUnauthorized {
	return &ImagesAwsImagesAsPostUnauthorized{}
}

/*
ImagesAwsImagesAsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesAwsImagesAsPostUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this images aws images as post unauthorized response has a 2xx status code
func (o *ImagesAwsImagesAsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images aws images as post unauthorized response has a 3xx status code
func (o *ImagesAwsImagesAsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images aws images as post unauthorized response has a 4xx status code
func (o *ImagesAwsImagesAsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images aws images as post unauthorized response has a 5xx status code
func (o *ImagesAwsImagesAsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images aws images as post unauthorized response a status code equal to that given
func (o *ImagesAwsImagesAsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the images aws images as post unauthorized response
func (o *ImagesAwsImagesAsPostUnauthorized) Code() int {
	return 401
}

func (o *ImagesAwsImagesAsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesAwsImagesAsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesAwsImagesAsPostUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAwsImagesAsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAwsImagesAsPostForbidden creates a ImagesAwsImagesAsPostForbidden with default headers values
func NewImagesAwsImagesAsPostForbidden() *ImagesAwsImagesAsPostForbidden {
	return &ImagesAwsImagesAsPostForbidden{}
}

/*
ImagesAwsImagesAsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesAwsImagesAsPostForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this images aws images as post forbidden response has a 2xx status code
func (o *ImagesAwsImagesAsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images aws images as post forbidden response has a 3xx status code
func (o *ImagesAwsImagesAsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images aws images as post forbidden response has a 4xx status code
func (o *ImagesAwsImagesAsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images aws images as post forbidden response has a 5xx status code
func (o *ImagesAwsImagesAsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images aws images as post forbidden response a status code equal to that given
func (o *ImagesAwsImagesAsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the images aws images as post forbidden response
func (o *ImagesAwsImagesAsPostForbidden) Code() int {
	return 403
}

func (o *ImagesAwsImagesAsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostForbidden  %+v", 403, o.Payload)
}

func (o *ImagesAwsImagesAsPostForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostForbidden  %+v", 403, o.Payload)
}

func (o *ImagesAwsImagesAsPostForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAwsImagesAsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAwsImagesAsPostNotFound creates a ImagesAwsImagesAsPostNotFound with default headers values
func NewImagesAwsImagesAsPostNotFound() *ImagesAwsImagesAsPostNotFound {
	return &ImagesAwsImagesAsPostNotFound{}
}

/*
ImagesAwsImagesAsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesAwsImagesAsPostNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this images aws images as post not found response has a 2xx status code
func (o *ImagesAwsImagesAsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images aws images as post not found response has a 3xx status code
func (o *ImagesAwsImagesAsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images aws images as post not found response has a 4xx status code
func (o *ImagesAwsImagesAsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images aws images as post not found response has a 5xx status code
func (o *ImagesAwsImagesAsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images aws images as post not found response a status code equal to that given
func (o *ImagesAwsImagesAsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the images aws images as post not found response
func (o *ImagesAwsImagesAsPostNotFound) Code() int {
	return 404
}

func (o *ImagesAwsImagesAsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostNotFound  %+v", 404, o.Payload)
}

func (o *ImagesAwsImagesAsPostNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostNotFound  %+v", 404, o.Payload)
}

func (o *ImagesAwsImagesAsPostNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ImagesAwsImagesAsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesAwsImagesAsPostInternalServerError creates a ImagesAwsImagesAsPostInternalServerError with default headers values
func NewImagesAwsImagesAsPostInternalServerError() *ImagesAwsImagesAsPostInternalServerError {
	return &ImagesAwsImagesAsPostInternalServerError{}
}

/*
ImagesAwsImagesAsPostInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesAwsImagesAsPostInternalServerError struct {
}

// IsSuccess returns true when this images aws images as post internal server error response has a 2xx status code
func (o *ImagesAwsImagesAsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images aws images as post internal server error response has a 3xx status code
func (o *ImagesAwsImagesAsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images aws images as post internal server error response has a 4xx status code
func (o *ImagesAwsImagesAsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images aws images as post internal server error response has a 5xx status code
func (o *ImagesAwsImagesAsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images aws images as post internal server error response a status code equal to that given
func (o *ImagesAwsImagesAsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the images aws images as post internal server error response
func (o *ImagesAwsImagesAsPostInternalServerError) Code() int {
	return 500
}

func (o *ImagesAwsImagesAsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostInternalServerError ", 500)
}

func (o *ImagesAwsImagesAsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/aws][%d] imagesAwsImagesAsPostInternalServerError ", 500)
}

func (o *ImagesAwsImagesAsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
