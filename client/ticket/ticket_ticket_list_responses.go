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

// TicketTicketListReader is a Reader for the TicketTicketList structure.
type TicketTicketListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketTicketListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketTicketListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketTicketListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketTicketListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketTicketListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketTicketListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketTicketListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketTicketListOK creates a TicketTicketListOK with default headers values
func NewTicketTicketListOK() *TicketTicketListOK {
	return &TicketTicketListOK{}
}

/*
TicketTicketListOK describes a response with status code 200, with default header values.

Success
*/
type TicketTicketListOK struct {
	Payload *TicketTicketListOKBody
}

// IsSuccess returns true when this ticket ticket list o k response has a 2xx status code
func (o *TicketTicketListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket ticket list o k response has a 3xx status code
func (o *TicketTicketListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket ticket list o k response has a 4xx status code
func (o *TicketTicketListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket ticket list o k response has a 5xx status code
func (o *TicketTicketListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket ticket list o k response a status code equal to that given
func (o *TicketTicketListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketTicketListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListOK  %+v", 200, o.Payload)
}

func (o *TicketTicketListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListOK  %+v", 200, o.Payload)
}

func (o *TicketTicketListOK) GetPayload() *TicketTicketListOKBody {
	return o.Payload
}

func (o *TicketTicketListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketTicketListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTicketListBadRequest creates a TicketTicketListBadRequest with default headers values
func NewTicketTicketListBadRequest() *TicketTicketListBadRequest {
	return &TicketTicketListBadRequest{}
}

/*
TicketTicketListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketTicketListBadRequest struct {
	Payload []*TicketTicketListBadRequestBodyItems0
}

// IsSuccess returns true when this ticket ticket list bad request response has a 2xx status code
func (o *TicketTicketListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket ticket list bad request response has a 3xx status code
func (o *TicketTicketListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket ticket list bad request response has a 4xx status code
func (o *TicketTicketListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket ticket list bad request response has a 5xx status code
func (o *TicketTicketListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket ticket list bad request response a status code equal to that given
func (o *TicketTicketListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketTicketListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListBadRequest  %+v", 400, o.Payload)
}

func (o *TicketTicketListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListBadRequest  %+v", 400, o.Payload)
}

func (o *TicketTicketListBadRequest) GetPayload() []*TicketTicketListBadRequestBodyItems0 {
	return o.Payload
}

func (o *TicketTicketListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTicketListUnauthorized creates a TicketTicketListUnauthorized with default headers values
func NewTicketTicketListUnauthorized() *TicketTicketListUnauthorized {
	return &TicketTicketListUnauthorized{}
}

/*
TicketTicketListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketTicketListUnauthorized struct {
	Payload *TicketTicketListUnauthorizedBody
}

// IsSuccess returns true when this ticket ticket list unauthorized response has a 2xx status code
func (o *TicketTicketListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket ticket list unauthorized response has a 3xx status code
func (o *TicketTicketListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket ticket list unauthorized response has a 4xx status code
func (o *TicketTicketListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket ticket list unauthorized response has a 5xx status code
func (o *TicketTicketListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket ticket list unauthorized response a status code equal to that given
func (o *TicketTicketListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketTicketListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketTicketListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketTicketListUnauthorized) GetPayload() *TicketTicketListUnauthorizedBody {
	return o.Payload
}

func (o *TicketTicketListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketTicketListUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTicketListForbidden creates a TicketTicketListForbidden with default headers values
func NewTicketTicketListForbidden() *TicketTicketListForbidden {
	return &TicketTicketListForbidden{}
}

/*
TicketTicketListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketTicketListForbidden struct {
	Payload *TicketTicketListForbiddenBody
}

// IsSuccess returns true when this ticket ticket list forbidden response has a 2xx status code
func (o *TicketTicketListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket ticket list forbidden response has a 3xx status code
func (o *TicketTicketListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket ticket list forbidden response has a 4xx status code
func (o *TicketTicketListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket ticket list forbidden response has a 5xx status code
func (o *TicketTicketListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket ticket list forbidden response a status code equal to that given
func (o *TicketTicketListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketTicketListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListForbidden  %+v", 403, o.Payload)
}

func (o *TicketTicketListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListForbidden  %+v", 403, o.Payload)
}

func (o *TicketTicketListForbidden) GetPayload() *TicketTicketListForbiddenBody {
	return o.Payload
}

func (o *TicketTicketListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketTicketListForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTicketListNotFound creates a TicketTicketListNotFound with default headers values
func NewTicketTicketListNotFound() *TicketTicketListNotFound {
	return &TicketTicketListNotFound{}
}

/*
TicketTicketListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketTicketListNotFound struct {
	Payload *TicketTicketListNotFoundBody
}

// IsSuccess returns true when this ticket ticket list not found response has a 2xx status code
func (o *TicketTicketListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket ticket list not found response has a 3xx status code
func (o *TicketTicketListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket ticket list not found response has a 4xx status code
func (o *TicketTicketListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket ticket list not found response has a 5xx status code
func (o *TicketTicketListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket ticket list not found response a status code equal to that given
func (o *TicketTicketListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketTicketListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListNotFound  %+v", 404, o.Payload)
}

func (o *TicketTicketListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListNotFound  %+v", 404, o.Payload)
}

func (o *TicketTicketListNotFound) GetPayload() *TicketTicketListNotFoundBody {
	return o.Payload
}

func (o *TicketTicketListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TicketTicketListNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTicketListInternalServerError creates a TicketTicketListInternalServerError with default headers values
func NewTicketTicketListInternalServerError() *TicketTicketListInternalServerError {
	return &TicketTicketListInternalServerError{}
}

/*
TicketTicketListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketTicketListInternalServerError struct {
}

// IsSuccess returns true when this ticket ticket list internal server error response has a 2xx status code
func (o *TicketTicketListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket ticket list internal server error response has a 3xx status code
func (o *TicketTicketListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket ticket list internal server error response has a 4xx status code
func (o *TicketTicketListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket ticket list internal server error response has a 5xx status code
func (o *TicketTicketListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket ticket list internal server error response a status code equal to that given
func (o *TicketTicketListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketTicketListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListInternalServerError ", 500)
}

func (o *TicketTicketListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/list][%d] ticketTicketListInternalServerError ", 500)
}

func (o *TicketTicketListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
TicketTicketListBadRequestBodyItems0 ticket ticket list bad request body items0
swagger:model TicketTicketListBadRequestBodyItems0
*/
type TicketTicketListBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this ticket ticket list bad request body items0
func (o *TicketTicketListBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket ticket list bad request body items0 based on context it is used
func (o *TicketTicketListBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketTicketListBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketTicketListBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res TicketTicketListBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketTicketListForbiddenBody ticket ticket list forbidden body
swagger:model TicketTicketListForbiddenBody
*/
type TicketTicketListForbiddenBody struct {

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

// Validate validates this ticket ticket list forbidden body
func (o *TicketTicketListForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket ticket list forbidden body based on context it is used
func (o *TicketTicketListForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketTicketListForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketTicketListForbiddenBody) UnmarshalBinary(b []byte) error {
	var res TicketTicketListForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketTicketListNotFoundBody ticket ticket list not found body
swagger:model TicketTicketListNotFoundBody
*/
type TicketTicketListNotFoundBody struct {

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

// Validate validates this ticket ticket list not found body
func (o *TicketTicketListNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket ticket list not found body based on context it is used
func (o *TicketTicketListNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketTicketListNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketTicketListNotFoundBody) UnmarshalBinary(b []byte) error {
	var res TicketTicketListNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketTicketListOKBody ticket ticket list o k body
swagger:model TicketTicketListOKBody
*/
type TicketTicketListOKBody struct {

	// data
	Data []*TicketTicketListOKBodyDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this ticket ticket list o k body
func (o *TicketTicketListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TicketTicketListOKBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("ticketTicketListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ticketTicketListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ticket ticket list o k body based on the context it is used
func (o *TicketTicketListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TicketTicketListOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ticketTicketListOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ticketTicketListOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TicketTicketListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketTicketListOKBody) UnmarshalBinary(b []byte) error {
	var res TicketTicketListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketTicketListOKBodyDataItems0 ticket ticket list o k body data items0
swagger:model TicketTicketListOKBodyDataItems0
*/
type TicketTicketListOKBodyDataItems0 struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// current status date
	CurrentStatusDate string `json:"currentStatusDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last modifier
	LastModifier string `json:"lastModifier,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number
	Number int32 `json:"number,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// partner logo
	PartnerLogo string `json:"partnerLogo,omitempty"`

	// partner name
	PartnerName string `json:"partnerName,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this ticket ticket list o k body data items0
func (o *TicketTicketListOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket ticket list o k body data items0 based on context it is used
func (o *TicketTicketListOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketTicketListOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketTicketListOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res TicketTicketListOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
TicketTicketListUnauthorizedBody ticket ticket list unauthorized body
swagger:model TicketTicketListUnauthorizedBody
*/
type TicketTicketListUnauthorizedBody struct {

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

// Validate validates this ticket ticket list unauthorized body
func (o *TicketTicketListUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ticket ticket list unauthorized body based on context it is used
func (o *TicketTicketListUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TicketTicketListUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TicketTicketListUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res TicketTicketListUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
