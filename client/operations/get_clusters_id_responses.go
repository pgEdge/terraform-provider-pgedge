// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
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

Response containing the cluster definition.
*/
type GetClustersIDOK struct {
	Payload *models.Cluster
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
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdOK %s", 200, payload)
}

func (o *GetClustersIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdOK %s", 200, payload)
}

func (o *GetClustersIDOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GetClustersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

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

Bad request.
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
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdBadRequest %s", 400, payload)
}

func (o *GetClustersIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdBadRequest %s", 400, payload)
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

Authorization information is missing or invalid.
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
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdUnauthorized %s", 401, payload)
}

func (o *GetClustersIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}][%d] getClustersIdUnauthorized %s", 401, payload)
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
