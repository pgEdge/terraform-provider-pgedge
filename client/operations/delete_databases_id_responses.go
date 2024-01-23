// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteDatabasesIDReader is a Reader for the DeleteDatabasesID structure.
type DeleteDatabasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatabasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDatabasesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDatabasesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDatabasesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDatabasesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /databases/{id}] DeleteDatabasesID", response, response.Code())
	}
}

// NewDeleteDatabasesIDNoContent creates a DeleteDatabasesIDNoContent with default headers values
func NewDeleteDatabasesIDNoContent() *DeleteDatabasesIDNoContent {
	return &DeleteDatabasesIDNoContent{}
}

/*
DeleteDatabasesIDNoContent describes a response with status code 204, with default header values.

No Content (successful deletion)
*/
type DeleteDatabasesIDNoContent struct {
}

// IsSuccess returns true when this delete databases Id no content response has a 2xx status code
func (o *DeleteDatabasesIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete databases Id no content response has a 3xx status code
func (o *DeleteDatabasesIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete databases Id no content response has a 4xx status code
func (o *DeleteDatabasesIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete databases Id no content response has a 5xx status code
func (o *DeleteDatabasesIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete databases Id no content response a status code equal to that given
func (o *DeleteDatabasesIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete databases Id no content response
func (o *DeleteDatabasesIDNoContent) Code() int {
	return 204
}

func (o *DeleteDatabasesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdNoContent ", 204)
}

func (o *DeleteDatabasesIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdNoContent ", 204)
}

func (o *DeleteDatabasesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDatabasesIDUnauthorized creates a DeleteDatabasesIDUnauthorized with default headers values
func NewDeleteDatabasesIDUnauthorized() *DeleteDatabasesIDUnauthorized {
	return &DeleteDatabasesIDUnauthorized{}
}

/*
DeleteDatabasesIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDatabasesIDUnauthorized struct {
}

// IsSuccess returns true when this delete databases Id unauthorized response has a 2xx status code
func (o *DeleteDatabasesIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete databases Id unauthorized response has a 3xx status code
func (o *DeleteDatabasesIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete databases Id unauthorized response has a 4xx status code
func (o *DeleteDatabasesIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete databases Id unauthorized response has a 5xx status code
func (o *DeleteDatabasesIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete databases Id unauthorized response a status code equal to that given
func (o *DeleteDatabasesIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete databases Id unauthorized response
func (o *DeleteDatabasesIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteDatabasesIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdUnauthorized ", 401)
}

func (o *DeleteDatabasesIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdUnauthorized ", 401)
}

func (o *DeleteDatabasesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDatabasesIDNotFound creates a DeleteDatabasesIDNotFound with default headers values
func NewDeleteDatabasesIDNotFound() *DeleteDatabasesIDNotFound {
	return &DeleteDatabasesIDNotFound{}
}

/*
DeleteDatabasesIDNotFound describes a response with status code 404, with default header values.

Database not found
*/
type DeleteDatabasesIDNotFound struct {
}

// IsSuccess returns true when this delete databases Id not found response has a 2xx status code
func (o *DeleteDatabasesIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete databases Id not found response has a 3xx status code
func (o *DeleteDatabasesIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete databases Id not found response has a 4xx status code
func (o *DeleteDatabasesIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete databases Id not found response has a 5xx status code
func (o *DeleteDatabasesIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete databases Id not found response a status code equal to that given
func (o *DeleteDatabasesIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete databases Id not found response
func (o *DeleteDatabasesIDNotFound) Code() int {
	return 404
}

func (o *DeleteDatabasesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdNotFound ", 404)
}

func (o *DeleteDatabasesIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdNotFound ", 404)
}

func (o *DeleteDatabasesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDatabasesIDInternalServerError creates a DeleteDatabasesIDInternalServerError with default headers values
func NewDeleteDatabasesIDInternalServerError() *DeleteDatabasesIDInternalServerError {
	return &DeleteDatabasesIDInternalServerError{}
}

/*
DeleteDatabasesIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteDatabasesIDInternalServerError struct {
}

// IsSuccess returns true when this delete databases Id internal server error response has a 2xx status code
func (o *DeleteDatabasesIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete databases Id internal server error response has a 3xx status code
func (o *DeleteDatabasesIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete databases Id internal server error response has a 4xx status code
func (o *DeleteDatabasesIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete databases Id internal server error response has a 5xx status code
func (o *DeleteDatabasesIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete databases Id internal server error response a status code equal to that given
func (o *DeleteDatabasesIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete databases Id internal server error response
func (o *DeleteDatabasesIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteDatabasesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdInternalServerError ", 500)
}

func (o *DeleteDatabasesIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /databases/{id}][%d] deleteDatabasesIdInternalServerError ", 500)
}

func (o *DeleteDatabasesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}