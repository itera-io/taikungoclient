// Code generated by go-swagger; DO NOT EDIT.

package ticket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TicketArticleListReader is a Reader for the TicketArticleList structure.
type TicketArticleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketArticleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketArticleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketArticleListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketArticleListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketArticleListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketArticleListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketArticleListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketArticleListOK creates a TicketArticleListOK with default headers values
func NewTicketArticleListOK() *TicketArticleListOK {
	return &TicketArticleListOK{}
}

/*
TicketArticleListOK describes a response with status code 200, with default header values.

Success
*/
type TicketArticleListOK struct {
	Payload *TicketArticleListOKBody
}

// IsSuccess returns true when this ticket article list o k response has a 2xx status code
func (o *TicketArticleListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket article list o k response has a 3xx status code
func (o *TicketArticleListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket article list o k response has a 4xx status code
func (o *TicketArticleListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket article list o k response has a 5xx status code
func (o *TicketArticleListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket article list o k response a status code equal to that given
func (o *TicketArticleListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketArticleListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListOK  %+v", 200, o.Payload)
}

func (o *TicketArticleListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListOK  %+v", 200, o.Payload)
}

func (o *TicketArticleListOK) GetPayload() *TicketArticleListOKBody {
	return o.Payload
}

func (o *TicketArticleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketArticleListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListBadRequest creates a TicketArticleListBadRequest with default headers values
func NewTicketArticleListBadRequest() *TicketArticleListBadRequest {
	return &TicketArticleListBadRequest{}
}

/*
TicketArticleListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketArticleListBadRequest struct {
	Payload []*TicketArticleListBadRequestBodyItems0
}

// IsSuccess returns true when this ticket article list bad request response has a 2xx status code
func (o *TicketArticleListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket article list bad request response has a 3xx status code
func (o *TicketArticleListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket article list bad request response has a 4xx status code
func (o *TicketArticleListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket article list bad request response has a 5xx status code
func (o *TicketArticleListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket article list bad request response a status code equal to that given
func (o *TicketArticleListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketArticleListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListBadRequest  %+v", 400, o.Payload)
}

func (o *TicketArticleListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListBadRequest  %+v", 400, o.Payload)
}

func (o *TicketArticleListBadRequest) GetPayload() []*TicketArticleListBadRequestBodyItems0 {
	return o.Payload
}

func (o *TicketArticleListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListUnauthorized creates a TicketArticleListUnauthorized with default headers values
func NewTicketArticleListUnauthorized() *TicketArticleListUnauthorized {
	return &TicketArticleListUnauthorized{}
}

/*
TicketArticleListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketArticleListUnauthorized struct {
	Payload *TicketArticleListUnauthorizedBody
}

// IsSuccess returns true when this ticket article list unauthorized response has a 2xx status code
func (o *TicketArticleListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket article list unauthorized response has a 3xx status code
func (o *TicketArticleListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket article list unauthorized response has a 4xx status code
func (o *TicketArticleListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket article list unauthorized response has a 5xx status code
func (o *TicketArticleListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket article list unauthorized response a status code equal to that given
func (o *TicketArticleListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketArticleListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketArticleListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketArticleListUnauthorized) GetPayload() *TicketArticleListUnauthorizedBody {
	return o.Payload
}

func (o *TicketArticleListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketArticleListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListForbidden creates a TicketArticleListForbidden with default headers values
func NewTicketArticleListForbidden() *TicketArticleListForbidden {
	return &TicketArticleListForbidden{}
}

/*
TicketArticleListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketArticleListForbidden struct {
	Payload *TicketArticleListForbiddenBody
}

// IsSuccess returns true when this ticket article list forbidden response has a 2xx status code
func (o *TicketArticleListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket article list forbidden response has a 3xx status code
func (o *TicketArticleListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket article list forbidden response has a 4xx status code
func (o *TicketArticleListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket article list forbidden response has a 5xx status code
func (o *TicketArticleListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket article list forbidden response a status code equal to that given
func (o *TicketArticleListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketArticleListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListForbidden  %+v", 403, o.Payload)
}

func (o *TicketArticleListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListForbidden  %+v", 403, o.Payload)
}

func (o *TicketArticleListForbidden) GetPayload() *TicketArticleListForbiddenBody {
	return o.Payload
}

func (o *TicketArticleListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketArticleListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListNotFound creates a TicketArticleListNotFound with default headers values
func NewTicketArticleListNotFound() *TicketArticleListNotFound {
	return &TicketArticleListNotFound{}
}

/*
TicketArticleListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketArticleListNotFound struct {
	Payload *TicketArticleListNotFoundBody
}

// IsSuccess returns true when this ticket article list not found response has a 2xx status code
func (o *TicketArticleListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket article list not found response has a 3xx status code
func (o *TicketArticleListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket article list not found response has a 4xx status code
func (o *TicketArticleListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket article list not found response has a 5xx status code
func (o *TicketArticleListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket article list not found response a status code equal to that given
func (o *TicketArticleListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketArticleListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListNotFound  %+v", 404, o.Payload)
}

func (o *TicketArticleListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListNotFound  %+v", 404, o.Payload)
}

func (o *TicketArticleListNotFound) GetPayload() *TicketArticleListNotFoundBody {
	return o.Payload
}

func (o *TicketArticleListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketArticleListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListInternalServerError creates a TicketArticleListInternalServerError with default headers values
func NewTicketArticleListInternalServerError() *TicketArticleListInternalServerError {
	return &TicketArticleListInternalServerError{}
}

/*
TicketArticleListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketArticleListInternalServerError struct {
}

// IsSuccess returns true when this ticket article list internal server error response has a 2xx status code
func (o *TicketArticleListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket article list internal server error response has a 3xx status code
func (o *TicketArticleListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket article list internal server error response has a 4xx status code
func (o *TicketArticleListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket article list internal server error response has a 5xx status code
func (o *TicketArticleListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket article list internal server error response a status code equal to that given
func (o *TicketArticleListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketArticleListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListInternalServerError ", 500)
}

func (o *TicketArticleListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListInternalServerError ", 500)
}

func (o *TicketArticleListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
TicketArticleListBadRequestBodyItems0 ticket article list bad request body items0
swagger:model TicketArticleListBadRequestBodyItems0
*/
type TicketArticleListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this ticket article list bad request body items0
func (o *TicketArticleListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket article list bad request body items0 based on context it is used
func (o *TicketArticleListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketArticleListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketArticleListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res TicketArticleListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketArticleListForbiddenBody ticket article list forbidden body
swagger:model TicketArticleListForbiddenBody
*/
type TicketArticleListForbiddenBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this ticket article list forbidden body
func (o *TicketArticleListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket article list forbidden body based on context it is used
func (o *TicketArticleListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketArticleListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketArticleListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res TicketArticleListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketArticleListNotFoundBody ticket article list not found body
swagger:model TicketArticleListNotFoundBody
*/
type TicketArticleListNotFoundBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this ticket article list not found body
func (o *TicketArticleListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket article list not found body based on context it is used
func (o *TicketArticleListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketArticleListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketArticleListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res TicketArticleListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketArticleListOKBody ticket article list o k body
swagger:model TicketArticleListOKBody
*/
type TicketArticleListOKBody struct {

	// data
	Data []*TicketArticleListOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this ticket article list o k body
func (o *TicketArticleListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TicketArticleListOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ticketArticleListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ticketArticleListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ticket article list o k body based on the context it is used
func (o *TicketArticleListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TicketArticleListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ticketArticleListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ticketArticleListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TicketArticleListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketArticleListOKBody) UnmarshalBinary(b []byte) error {
	var res TicketArticleListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketArticleListOKBodyDataItems0 ticket article list o k body data items0
swagger:model TicketArticleListOKBodyDataItems0
*/
type TicketArticleListOKBodyDataItems0 struct {

	// body
	Body string `json:"body,omitempty"`

	// create at
	CreateAt string `json:"createAt,omitempty"`

	// is csm
	IsCsm bool `json:"isCsm"`

	// message Id
	MessageID string `json:"messageId,omitempty"`

	// sender name
	SenderName string `json:"senderName,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this ticket article list o k body data items0
func (o *TicketArticleListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket article list o k body data items0 based on context it is used
func (o *TicketArticleListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketArticleListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketArticleListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res TicketArticleListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketArticleListUnauthorizedBody ticket article list unauthorized body
swagger:model TicketArticleListUnauthorizedBody
*/
type TicketArticleListUnauthorizedBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this ticket article list unauthorized body
func (o *TicketArticleListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket article list unauthorized body based on context it is used
func (o *TicketArticleListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketArticleListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketArticleListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res TicketArticleListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
