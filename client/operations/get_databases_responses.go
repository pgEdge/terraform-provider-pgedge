// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/models"
)

// GetDatabasesReader is a Reader for the GetDatabases structure.
type GetDatabasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatabasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatabasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDatabasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDatabasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /databases] GetDatabases", response, response.Code())
	}
}

// NewGetDatabasesOK creates a GetDatabasesOK with default headers values
func NewGetDatabasesOK() *GetDatabasesOK {
	return &GetDatabasesOK{}
}

/*
GetDatabasesOK describes a response with status code 200, with default header values.

Successful response
*/
type GetDatabasesOK struct {
	Payload []*models.Database
}

// IsSuccess returns true when this get databases o k response has a 2xx status code
func (o *GetDatabasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get databases o k response has a 3xx status code
func (o *GetDatabasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get databases o k response has a 4xx status code
func (o *GetDatabasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get databases o k response has a 5xx status code
func (o *GetDatabasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get databases o k response a status code equal to that given
func (o *GetDatabasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get databases o k response
func (o *GetDatabasesOK) Code() int {
	return 200
}

func (o *GetDatabasesOK) Error() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesOK  %+v", 200, o.Payload)
}

func (o *GetDatabasesOK) String() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesOK  %+v", 200, o.Payload)
}

func (o *GetDatabasesOK) GetPayload() []*models.Database {
	return o.Payload
}

func (o *GetDatabasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatabasesUnauthorized creates a GetDatabasesUnauthorized with default headers values
func NewGetDatabasesUnauthorized() *GetDatabasesUnauthorized {
	return &GetDatabasesUnauthorized{}
}

/*
GetDatabasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDatabasesUnauthorized struct {
}

// IsSuccess returns true when this get databases unauthorized response has a 2xx status code
func (o *GetDatabasesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get databases unauthorized response has a 3xx status code
func (o *GetDatabasesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get databases unauthorized response has a 4xx status code
func (o *GetDatabasesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get databases unauthorized response has a 5xx status code
func (o *GetDatabasesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get databases unauthorized response a status code equal to that given
func (o *GetDatabasesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get databases unauthorized response
func (o *GetDatabasesUnauthorized) Code() int {
	return 401
}

func (o *GetDatabasesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesUnauthorized ", 401)
}

func (o *GetDatabasesUnauthorized) String() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesUnauthorized ", 401)
}

func (o *GetDatabasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDatabasesInternalServerError creates a GetDatabasesInternalServerError with default headers values
func NewGetDatabasesInternalServerError() *GetDatabasesInternalServerError {
	return &GetDatabasesInternalServerError{}
}

/*
GetDatabasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDatabasesInternalServerError struct {
}

// IsSuccess returns true when this get databases internal server error response has a 2xx status code
func (o *GetDatabasesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get databases internal server error response has a 3xx status code
func (o *GetDatabasesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get databases internal server error response has a 4xx status code
func (o *GetDatabasesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get databases internal server error response has a 5xx status code
func (o *GetDatabasesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get databases internal server error response a status code equal to that given
func (o *GetDatabasesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get databases internal server error response
func (o *GetDatabasesInternalServerError) Code() int {
	return 500
}

func (o *GetDatabasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesInternalServerError ", 500)
}

func (o *GetDatabasesInternalServerError) String() string {
	return fmt.Sprintf("[GET /databases][%d] getDatabasesInternalServerError ", 500)
}

func (o *GetDatabasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
