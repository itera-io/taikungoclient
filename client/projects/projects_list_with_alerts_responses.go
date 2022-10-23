// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// ProjectsListWithAlertsReader is a Reader for the ProjectsListWithAlerts structure.
type ProjectsListWithAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListWithAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListWithAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsListWithAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsListWithAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsListWithAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsListWithAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsListWithAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListWithAlertsOK creates a ProjectsListWithAlertsOK with default headers values
func NewProjectsListWithAlertsOK() *ProjectsListWithAlertsOK {
	return &ProjectsListWithAlertsOK{}
}

/*
ProjectsListWithAlertsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListWithAlertsOK struct {
	Payload []*ProjectsListWithAlertsOKBodyItems0
}

// IsSuccess returns true when this projects list with alerts o k response has a 2xx status code
func (o *ProjectsListWithAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects list with alerts o k response has a 3xx status code
func (o *ProjectsListWithAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts o k response has a 4xx status code
func (o *ProjectsListWithAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list with alerts o k response has a 5xx status code
func (o *ProjectsListWithAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts o k response a status code equal to that given
func (o *ProjectsListWithAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsListWithAlertsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsOK  %+v", 200, o.Payload)
}

func (o *ProjectsListWithAlertsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsOK  %+v", 200, o.Payload)
}

func (o *ProjectsListWithAlertsOK) GetPayload() []*ProjectsListWithAlertsOKBodyItems0 {
	return o.Payload
}

func (o *ProjectsListWithAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsBadRequest creates a ProjectsListWithAlertsBadRequest with default headers values
func NewProjectsListWithAlertsBadRequest() *ProjectsListWithAlertsBadRequest {
	return &ProjectsListWithAlertsBadRequest{}
}

/*
ProjectsListWithAlertsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsListWithAlertsBadRequest struct {
	Payload []*ProjectsListWithAlertsBadRequestBodyItems0
}

// IsSuccess returns true when this projects list with alerts bad request response has a 2xx status code
func (o *ProjectsListWithAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts bad request response has a 3xx status code
func (o *ProjectsListWithAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts bad request response has a 4xx status code
func (o *ProjectsListWithAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts bad request response has a 5xx status code
func (o *ProjectsListWithAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts bad request response a status code equal to that given
func (o *ProjectsListWithAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsListWithAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListWithAlertsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListWithAlertsBadRequest) GetPayload() []*ProjectsListWithAlertsBadRequestBodyItems0 {
	return o.Payload
}

func (o *ProjectsListWithAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsUnauthorized creates a ProjectsListWithAlertsUnauthorized with default headers values
func NewProjectsListWithAlertsUnauthorized() *ProjectsListWithAlertsUnauthorized {
	return &ProjectsListWithAlertsUnauthorized{}
}

/*
ProjectsListWithAlertsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsListWithAlertsUnauthorized struct {
	Payload *ProjectsListWithAlertsUnauthorizedBody
}

// IsSuccess returns true when this projects list with alerts unauthorized response has a 2xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts unauthorized response has a 3xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts unauthorized response has a 4xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts unauthorized response has a 5xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts unauthorized response a status code equal to that given
func (o *ProjectsListWithAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsListWithAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListWithAlertsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListWithAlertsUnauthorized) GetPayload() *ProjectsListWithAlertsUnauthorizedBody {
	return o.Payload
}

func (o *ProjectsListWithAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectsListWithAlertsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsForbidden creates a ProjectsListWithAlertsForbidden with default headers values
func NewProjectsListWithAlertsForbidden() *ProjectsListWithAlertsForbidden {
	return &ProjectsListWithAlertsForbidden{}
}

/*
ProjectsListWithAlertsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsListWithAlertsForbidden struct {
	Payload *ProjectsListWithAlertsForbiddenBody
}

// IsSuccess returns true when this projects list with alerts forbidden response has a 2xx status code
func (o *ProjectsListWithAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts forbidden response has a 3xx status code
func (o *ProjectsListWithAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts forbidden response has a 4xx status code
func (o *ProjectsListWithAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts forbidden response has a 5xx status code
func (o *ProjectsListWithAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts forbidden response a status code equal to that given
func (o *ProjectsListWithAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsListWithAlertsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListWithAlertsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListWithAlertsForbidden) GetPayload() *ProjectsListWithAlertsForbiddenBody {
	return o.Payload
}

func (o *ProjectsListWithAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectsListWithAlertsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsNotFound creates a ProjectsListWithAlertsNotFound with default headers values
func NewProjectsListWithAlertsNotFound() *ProjectsListWithAlertsNotFound {
	return &ProjectsListWithAlertsNotFound{}
}

/*
ProjectsListWithAlertsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsListWithAlertsNotFound struct {
	Payload *ProjectsListWithAlertsNotFoundBody
}

// IsSuccess returns true when this projects list with alerts not found response has a 2xx status code
func (o *ProjectsListWithAlertsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts not found response has a 3xx status code
func (o *ProjectsListWithAlertsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts not found response has a 4xx status code
func (o *ProjectsListWithAlertsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts not found response has a 5xx status code
func (o *ProjectsListWithAlertsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts not found response a status code equal to that given
func (o *ProjectsListWithAlertsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsListWithAlertsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListWithAlertsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListWithAlertsNotFound) GetPayload() *ProjectsListWithAlertsNotFoundBody {
	return o.Payload
}

func (o *ProjectsListWithAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProjectsListWithAlertsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsInternalServerError creates a ProjectsListWithAlertsInternalServerError with default headers values
func NewProjectsListWithAlertsInternalServerError() *ProjectsListWithAlertsInternalServerError {
	return &ProjectsListWithAlertsInternalServerError{}
}

/*
ProjectsListWithAlertsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsListWithAlertsInternalServerError struct {
}

// IsSuccess returns true when this projects list with alerts internal server error response has a 2xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts internal server error response has a 3xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts internal server error response has a 4xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list with alerts internal server error response has a 5xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects list with alerts internal server error response a status code equal to that given
func (o *ProjectsListWithAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsListWithAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsInternalServerError ", 500)
}

func (o *ProjectsListWithAlertsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsInternalServerError ", 500)
}

func (o *ProjectsListWithAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ProjectsListWithAlertsBadRequestBodyItems0 projects list with alerts bad request body items0
swagger:model ProjectsListWithAlertsBadRequestBodyItems0
*/
type ProjectsListWithAlertsBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this projects list with alerts bad request body items0
func (o *ProjectsListWithAlertsBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects list with alerts bad request body items0 based on context it is used
func (o *ProjectsListWithAlertsBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsListWithAlertsForbiddenBody projects list with alerts forbidden body
swagger:model ProjectsListWithAlertsForbiddenBody
*/
type ProjectsListWithAlertsForbiddenBody struct {

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

// Validate validates this projects list with alerts forbidden body
func (o *ProjectsListWithAlertsForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects list with alerts forbidden body based on context it is used
func (o *ProjectsListWithAlertsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsListWithAlertsNotFoundBody projects list with alerts not found body
swagger:model ProjectsListWithAlertsNotFoundBody
*/
type ProjectsListWithAlertsNotFoundBody struct {

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

// Validate validates this projects list with alerts not found body
func (o *ProjectsListWithAlertsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects list with alerts not found body based on context it is used
func (o *ProjectsListWithAlertsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsListWithAlertsOKBodyItems0 projects list with alerts o k body items0
swagger:model ProjectsListWithAlertsOKBodyItems0
*/
type ProjectsListWithAlertsOKBodyItems0 struct {

	// has kube config file
	HasKubeConfigFile bool `json:"hasKubeConfigFile"`

	// health
	Health string `json:"health,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// is monitoring enabled
	IsMonitoringEnabled bool `json:"isMonitoringEnabled"`

	// kubernetes alerts
	KubernetesAlerts []*ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0 `json:"kubernetesAlerts"`

	// monitoring credentials
	MonitoringCredentials []*ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0 `json:"monitoringCredentials"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this projects list with alerts o k body items0
func (o *ProjectsListWithAlertsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubernetesAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMonitoringCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsListWithAlertsOKBodyItems0) validateKubernetesAlerts(formats strfmt.Registry) error {
	if swag.IsZero(o.KubernetesAlerts) { // not required
		return nil
	}

	for i := 0; i < len(o.KubernetesAlerts); i++ {
		if swag.IsZero(o.KubernetesAlerts[i]) { // not required
			continue
		}

		if o.KubernetesAlerts[i] != nil {
			if err := o.KubernetesAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ProjectsListWithAlertsOKBodyItems0) validateMonitoringCredentials(formats strfmt.Registry) error {
	if swag.IsZero(o.MonitoringCredentials) { // not required
		return nil
	}

	for i := 0; i < len(o.MonitoringCredentials); i++ {
		if swag.IsZero(o.MonitoringCredentials[i]) { // not required
			continue
		}

		if o.MonitoringCredentials[i] != nil {
			if err := o.MonitoringCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this projects list with alerts o k body items0 based on the context it is used
func (o *ProjectsListWithAlertsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateKubernetesAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMonitoringCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ProjectsListWithAlertsOKBodyItems0) contextValidateKubernetesAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.KubernetesAlerts); i++ {

		if o.KubernetesAlerts[i] != nil {
			if err := o.KubernetesAlerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ProjectsListWithAlertsOKBodyItems0) contextValidateMonitoringCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.MonitoringCredentials); i++ {

		if o.MonitoringCredentials[i] != nil {
			if err := o.MonitoringCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0 projects list with alerts o k body items0 kubernetes alerts items0
swagger:model ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0
*/
type ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0 struct {

	// description
	Description string `json:"description,omitempty"`

	// end at
	EndAt string `json:"endAt,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is silenced
	IsSilenced bool `json:"isSilenced"`

	// is solved
	IsSolved bool `json:"isSolved"`

	// labels
	Labels interface{} `json:"labels,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// silence reason
	SilenceReason string `json:"silenceReason,omitempty"`

	// starts at
	StartsAt string `json:"startsAt,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this projects list with alerts o k body items0 kubernetes alerts items0
func (o *ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects list with alerts o k body items0 kubernetes alerts items0 based on context it is used
func (o *ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsOKBodyItems0KubernetesAlertsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0 projects list with alerts o k body items0 monitoring credentials items0
swagger:model ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0
*/
type ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0 struct {

	// alert manager Url
	AlertManagerURL string `json:"alertManagerUrl,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// loki Url
	LokiURL string `json:"lokiUrl,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// prometheus Url
	PrometheusURL string `json:"prometheusUrl,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this projects list with alerts o k body items0 monitoring credentials items0
func (o *ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects list with alerts o k body items0 monitoring credentials items0 based on context it is used
func (o *ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsOKBodyItems0MonitoringCredentialsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ProjectsListWithAlertsUnauthorizedBody projects list with alerts unauthorized body
swagger:model ProjectsListWithAlertsUnauthorizedBody
*/
type ProjectsListWithAlertsUnauthorizedBody struct {

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

// Validate validates this projects list with alerts unauthorized body
func (o *ProjectsListWithAlertsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this projects list with alerts unauthorized body based on context it is used
func (o *ProjectsListWithAlertsUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectsListWithAlertsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectsListWithAlertsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ProjectsListWithAlertsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
