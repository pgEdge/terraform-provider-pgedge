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

// GetClustersIDReader is a Reader for the GetClustersID structure.
type GetClustersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClustersIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetClustersIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClustersIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /clusters/{id}] GetClustersID", response, response.Code())
	}
}

// NewGetClustersIDOK creates a GetClustersIDOK with default headers values
func NewGetClustersIDOK() *GetClustersIDOK {
	return &GetClustersIDOK{}
}

/*
GetClustersIDOK describes a response with status code 200, with default header values.

Successful response
*/
type GetClustersIDOK struct {
	Payload *models.ClusterDetails
}

// IsSuccess returns true when this get clusters Id o k response has a 2xx status code
func (o *GetClustersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get clusters Id o k response has a 3xx status code
func (o *GetClustersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id o k response has a 4xx status code
func (o *GetClustersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clusters Id o k response has a 5xx status code
func (o *GetClustersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters Id o k response a status code equal to that given
func (o *GetClustersIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get clusters Id o k response
func (o *GetClustersIDOK) Code() int {
	return 200
}

func (o *GetClustersIDOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdOK  %+v", 200, o.Payload)
}

func (o *GetClustersIDOK) String() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdOK  %+v", 200, o.Payload)
}

func (o *GetClustersIDOK) GetPayload() *models.ClusterDetails {
	return o.Payload
}

func (o *GetClustersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersIDBadRequest creates a GetClustersIDBadRequest with default headers values
func NewGetClustersIDBadRequest() *GetClustersIDBadRequest {
	return &GetClustersIDBadRequest{}
}

/*
GetClustersIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetClustersIDBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters Id bad request response has a 2xx status code
func (o *GetClustersIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters Id bad request response has a 3xx status code
func (o *GetClustersIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id bad request response has a 4xx status code
func (o *GetClustersIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters Id bad request response has a 5xx status code
func (o *GetClustersIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters Id bad request response a status code equal to that given
func (o *GetClustersIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get clusters Id bad request response
func (o *GetClustersIDBadRequest) Code() int {
	return 400
}

func (o *GetClustersIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetClustersIDBadRequest) String() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetClustersIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersIDUnauthorized creates a GetClustersIDUnauthorized with default headers values
func NewGetClustersIDUnauthorized() *GetClustersIDUnauthorized {
	return &GetClustersIDUnauthorized{}
}

/*
GetClustersIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetClustersIDUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters Id unauthorized response has a 2xx status code
func (o *GetClustersIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters Id unauthorized response has a 3xx status code
func (o *GetClustersIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id unauthorized response has a 4xx status code
func (o *GetClustersIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters Id unauthorized response has a 5xx status code
func (o *GetClustersIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters Id unauthorized response a status code equal to that given
func (o *GetClustersIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get clusters Id unauthorized response
func (o *GetClustersIDUnauthorized) Code() int {
	return 401
}

func (o *GetClustersIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetClustersIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetClustersIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersIDInternalServerError creates a GetClustersIDInternalServerError with default headers values
func NewGetClustersIDInternalServerError() *GetClustersIDInternalServerError {
	return &GetClustersIDInternalServerError{}
}

/*
GetClustersIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetClustersIDInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters Id internal server error response has a 2xx status code
func (o *GetClustersIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters Id internal server error response has a 3xx status code
func (o *GetClustersIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id internal server error response has a 4xx status code
func (o *GetClustersIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clusters Id internal server error response has a 5xx status code
func (o *GetClustersIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get clusters Id internal server error response a status code equal to that given
func (o *GetClustersIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get clusters Id internal server error response
func (o *GetClustersIDInternalServerError) Code() int {
	return 500
}

func (o *GetClustersIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClustersIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClustersIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
