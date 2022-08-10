// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectsListReader is a Reader for the ProjectsList structure.
type ProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListOK creates a ProjectsListOK with default headers values
func NewProjectsListOK() *ProjectsListOK {
	return &ProjectsListOK{}
}

/* ProjectsListOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListOK struct {
	Payload *models.ShowbackProjectsList
}

func (o *ProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/Projects][%d] projectsListOK  %+v", 200, o.Payload)
}
func (o *ProjectsListOK) GetPayload() *models.ShowbackProjectsList {
	return o.Payload
}

func (o *ProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShowbackProjectsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}