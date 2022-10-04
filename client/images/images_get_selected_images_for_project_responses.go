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

// ImagesGetSelectedImagesForProjectReader is a Reader for the ImagesGetSelectedImagesForProject structure.
type ImagesGetSelectedImagesForProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagesGetSelectedImagesForProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagesGetSelectedImagesForProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImagesGetSelectedImagesForProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImagesGetSelectedImagesForProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImagesGetSelectedImagesForProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImagesGetSelectedImagesForProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagesGetSelectedImagesForProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagesGetSelectedImagesForProjectOK creates a ImagesGetSelectedImagesForProjectOK with default headers values
func NewImagesGetSelectedImagesForProjectOK() *ImagesGetSelectedImagesForProjectOK {
	return &ImagesGetSelectedImagesForProjectOK{}
}

/*
ImagesGetSelectedImagesForProjectOK describes a response with status code 200, with default header values.

Success
*/
type ImagesGetSelectedImagesForProjectOK struct {
	Payload *models.BoundImagesForProjectsList
}

// IsSuccess returns true when this images get selected images for project o k response has a 2xx status code
func (o *ImagesGetSelectedImagesForProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this images get selected images for project o k response has a 3xx status code
func (o *ImagesGetSelectedImagesForProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get selected images for project o k response has a 4xx status code
func (o *ImagesGetSelectedImagesForProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this images get selected images for project o k response has a 5xx status code
func (o *ImagesGetSelectedImagesForProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this images get selected images for project o k response a status code equal to that given
func (o *ImagesGetSelectedImagesForProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagesGetSelectedImagesForProjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectOK  %+v", 200, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectOK  %+v", 200, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectOK) GetPayload() *models.BoundImagesForProjectsList {
	return o.Payload
}

func (o *ImagesGetSelectedImagesForProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BoundImagesForProjectsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetSelectedImagesForProjectBadRequest creates a ImagesGetSelectedImagesForProjectBadRequest with default headers values
func NewImagesGetSelectedImagesForProjectBadRequest() *ImagesGetSelectedImagesForProjectBadRequest {
	return &ImagesGetSelectedImagesForProjectBadRequest{}
}

/*
ImagesGetSelectedImagesForProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImagesGetSelectedImagesForProjectBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this images get selected images for project bad request response has a 2xx status code
func (o *ImagesGetSelectedImagesForProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get selected images for project bad request response has a 3xx status code
func (o *ImagesGetSelectedImagesForProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get selected images for project bad request response has a 4xx status code
func (o *ImagesGetSelectedImagesForProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get selected images for project bad request response has a 5xx status code
func (o *ImagesGetSelectedImagesForProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this images get selected images for project bad request response a status code equal to that given
func (o *ImagesGetSelectedImagesForProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImagesGetSelectedImagesForProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ImagesGetSelectedImagesForProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetSelectedImagesForProjectUnauthorized creates a ImagesGetSelectedImagesForProjectUnauthorized with default headers values
func NewImagesGetSelectedImagesForProjectUnauthorized() *ImagesGetSelectedImagesForProjectUnauthorized {
	return &ImagesGetSelectedImagesForProjectUnauthorized{}
}

/*
ImagesGetSelectedImagesForProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ImagesGetSelectedImagesForProjectUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images get selected images for project unauthorized response has a 2xx status code
func (o *ImagesGetSelectedImagesForProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get selected images for project unauthorized response has a 3xx status code
func (o *ImagesGetSelectedImagesForProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get selected images for project unauthorized response has a 4xx status code
func (o *ImagesGetSelectedImagesForProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get selected images for project unauthorized response has a 5xx status code
func (o *ImagesGetSelectedImagesForProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this images get selected images for project unauthorized response a status code equal to that given
func (o *ImagesGetSelectedImagesForProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ImagesGetSelectedImagesForProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesGetSelectedImagesForProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetSelectedImagesForProjectForbidden creates a ImagesGetSelectedImagesForProjectForbidden with default headers values
func NewImagesGetSelectedImagesForProjectForbidden() *ImagesGetSelectedImagesForProjectForbidden {
	return &ImagesGetSelectedImagesForProjectForbidden{}
}

/*
ImagesGetSelectedImagesForProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ImagesGetSelectedImagesForProjectForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images get selected images for project forbidden response has a 2xx status code
func (o *ImagesGetSelectedImagesForProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get selected images for project forbidden response has a 3xx status code
func (o *ImagesGetSelectedImagesForProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get selected images for project forbidden response has a 4xx status code
func (o *ImagesGetSelectedImagesForProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get selected images for project forbidden response has a 5xx status code
func (o *ImagesGetSelectedImagesForProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this images get selected images for project forbidden response a status code equal to that given
func (o *ImagesGetSelectedImagesForProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ImagesGetSelectedImagesForProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectForbidden  %+v", 403, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectForbidden  %+v", 403, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesGetSelectedImagesForProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetSelectedImagesForProjectNotFound creates a ImagesGetSelectedImagesForProjectNotFound with default headers values
func NewImagesGetSelectedImagesForProjectNotFound() *ImagesGetSelectedImagesForProjectNotFound {
	return &ImagesGetSelectedImagesForProjectNotFound{}
}

/*
ImagesGetSelectedImagesForProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImagesGetSelectedImagesForProjectNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this images get selected images for project not found response has a 2xx status code
func (o *ImagesGetSelectedImagesForProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get selected images for project not found response has a 3xx status code
func (o *ImagesGetSelectedImagesForProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get selected images for project not found response has a 4xx status code
func (o *ImagesGetSelectedImagesForProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this images get selected images for project not found response has a 5xx status code
func (o *ImagesGetSelectedImagesForProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this images get selected images for project not found response a status code equal to that given
func (o *ImagesGetSelectedImagesForProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagesGetSelectedImagesForProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectNotFound  %+v", 404, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectNotFound  %+v", 404, o.Payload)
}

func (o *ImagesGetSelectedImagesForProjectNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ImagesGetSelectedImagesForProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagesGetSelectedImagesForProjectInternalServerError creates a ImagesGetSelectedImagesForProjectInternalServerError with default headers values
func NewImagesGetSelectedImagesForProjectInternalServerError() *ImagesGetSelectedImagesForProjectInternalServerError {
	return &ImagesGetSelectedImagesForProjectInternalServerError{}
}

/*
ImagesGetSelectedImagesForProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ImagesGetSelectedImagesForProjectInternalServerError struct {
}

// IsSuccess returns true when this images get selected images for project internal server error response has a 2xx status code
func (o *ImagesGetSelectedImagesForProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this images get selected images for project internal server error response has a 3xx status code
func (o *ImagesGetSelectedImagesForProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this images get selected images for project internal server error response has a 4xx status code
func (o *ImagesGetSelectedImagesForProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this images get selected images for project internal server error response has a 5xx status code
func (o *ImagesGetSelectedImagesForProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this images get selected images for project internal server error response a status code equal to that given
func (o *ImagesGetSelectedImagesForProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagesGetSelectedImagesForProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectInternalServerError ", 500)
}

func (o *ImagesGetSelectedImagesForProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Images/projects/list][%d] imagesGetSelectedImagesForProjectInternalServerError ", 500)
}

func (o *ImagesGetSelectedImagesForProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
