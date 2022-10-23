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

// ImagesBindImagesToProjectReader is a Reader for the ImagesBindImagesToProject structure.
type ImagesBindImagesToProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesBindImagesToProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesBindImagesToProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesBindImagesToProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesBindImagesToProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesBindImagesToProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesBindImagesToProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesBindImagesToProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesBindImagesToProjectOK creates a ImagesBindImagesToProjectOK with default headers values
func NewImagesBindImagesToProjectOK() *ImagesBindImagesToProjectOK {
	return &ImagesBindImagesToProjectOK{}
}

/*
ImagesBindImagesToProjectOK describes a response with status code 200, with default header values.

Success
*/
type ImagesBindImagesToProjectOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this images bind images to project o k response has a 2xx status code
func (o *ImagesBindImagesToProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images bind images to project o k response has a 3xx status code
func (o *ImagesBindImagesToProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images bind images to project o k response has a 4xx status code
func (o *ImagesBindImagesToProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images bind images to project o k response has a 5xx status code
func (o *ImagesBindImagesToProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images bind images to project o k response a status code equal to that given
func (o *ImagesBindImagesToProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesBindImagesToProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectOK  %+v", 200, o.Payload)
}

func (o *ImagesBindImagesToProjectOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectOK  %+v", 200, o.Payload)
}

func (o *ImagesBindImagesToProjectOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ImagesBindImagesToProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesBindImagesToProjectBadRequest creates a ImagesBindImagesToProjectBadRequest with default headers values
func NewImagesBindImagesToProjectBadRequest() *ImagesBindImagesToProjectBadRequest {
	return &ImagesBindImagesToProjectBadRequest{}
}

/*
ImagesBindImagesToProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesBindImagesToProjectBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this images bind images to project bad request response has a 2xx status code
func (o *ImagesBindImagesToProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images bind images to project bad request response has a 3xx status code
func (o *ImagesBindImagesToProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images bind images to project bad request response has a 4xx status code
func (o *ImagesBindImagesToProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images bind images to project bad request response has a 5xx status code
func (o *ImagesBindImagesToProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images bind images to project bad request response a status code equal to that given
func (o *ImagesBindImagesToProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesBindImagesToProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesBindImagesToProjectBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesBindImagesToProjectBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ImagesBindImagesToProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesBindImagesToProjectUnauthorized creates a ImagesBindImagesToProjectUnauthorized with default headers values
func NewImagesBindImagesToProjectUnauthorized() *ImagesBindImagesToProjectUnauthorized {
	return &ImagesBindImagesToProjectUnauthorized{}
}

/*
ImagesBindImagesToProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesBindImagesToProjectUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images bind images to project unauthorized response has a 2xx status code
func (o *ImagesBindImagesToProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images bind images to project unauthorized response has a 3xx status code
func (o *ImagesBindImagesToProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images bind images to project unauthorized response has a 4xx status code
func (o *ImagesBindImagesToProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images bind images to project unauthorized response has a 5xx status code
func (o *ImagesBindImagesToProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images bind images to project unauthorized response a status code equal to that given
func (o *ImagesBindImagesToProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesBindImagesToProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesBindImagesToProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesBindImagesToProjectUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesBindImagesToProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesBindImagesToProjectForbidden creates a ImagesBindImagesToProjectForbidden with default headers values
func NewImagesBindImagesToProjectForbidden() *ImagesBindImagesToProjectForbidden {
	return &ImagesBindImagesToProjectForbidden{}
}

/*
ImagesBindImagesToProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesBindImagesToProjectForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images bind images to project forbidden response has a 2xx status code
func (o *ImagesBindImagesToProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images bind images to project forbidden response has a 3xx status code
func (o *ImagesBindImagesToProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images bind images to project forbidden response has a 4xx status code
func (o *ImagesBindImagesToProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images bind images to project forbidden response has a 5xx status code
func (o *ImagesBindImagesToProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images bind images to project forbidden response a status code equal to that given
func (o *ImagesBindImagesToProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesBindImagesToProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectForbidden  %+v", 403, o.Payload)
}

func (o *ImagesBindImagesToProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectForbidden  %+v", 403, o.Payload)
}

func (o *ImagesBindImagesToProjectForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesBindImagesToProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesBindImagesToProjectNotFound creates a ImagesBindImagesToProjectNotFound with default headers values
func NewImagesBindImagesToProjectNotFound() *ImagesBindImagesToProjectNotFound {
	return &ImagesBindImagesToProjectNotFound{}
}

/*
ImagesBindImagesToProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesBindImagesToProjectNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images bind images to project not found response has a 2xx status code
func (o *ImagesBindImagesToProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images bind images to project not found response has a 3xx status code
func (o *ImagesBindImagesToProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images bind images to project not found response has a 4xx status code
func (o *ImagesBindImagesToProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images bind images to project not found response has a 5xx status code
func (o *ImagesBindImagesToProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images bind images to project not found response a status code equal to that given
func (o *ImagesBindImagesToProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesBindImagesToProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectNotFound  %+v", 404, o.Payload)
}

func (o *ImagesBindImagesToProjectNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectNotFound  %+v", 404, o.Payload)
}

func (o *ImagesBindImagesToProjectNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesBindImagesToProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesBindImagesToProjectInternalServerError creates a ImagesBindImagesToProjectInternalServerError with default headers values
func NewImagesBindImagesToProjectInternalServerError() *ImagesBindImagesToProjectInternalServerError {
	return &ImagesBindImagesToProjectInternalServerError{}
}

/*
ImagesBindImagesToProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesBindImagesToProjectInternalServerError struct {
}

// IsSuccess returns true when this images bind images to project internal server error response has a 2xx status code
func (o *ImagesBindImagesToProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images bind images to project internal server error response has a 3xx status code
func (o *ImagesBindImagesToProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images bind images to project internal server error response has a 4xx status code
func (o *ImagesBindImagesToProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images bind images to project internal server error response has a 5xx status code
func (o *ImagesBindImagesToProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images bind images to project internal server error response a status code equal to that given
func (o *ImagesBindImagesToProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesBindImagesToProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectInternalServerError ", 500)
}

func (o *ImagesBindImagesToProjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Images/bind][%d] imagesBindImagesToProjectInternalServerError ", 500)
}

func (o *ImagesBindImagesToProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
