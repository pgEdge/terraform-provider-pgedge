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

// GetClustersIDNodesNodeIDLogsLogNameReader is a Reader for the GetClustersIDNodesNodeIDLogsLogName structure.
type GetClustersIDNodesNodeIDLogsLogNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersIDNodesNodeIDLogsLogNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersIDNodesNodeIDLogsLogNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClustersIDNodesNodeIDLogsLogNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetClustersIDNodesNodeIDLogsLogNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}] GetClustersIDNodesNodeIDLogsLogName", response, response.Code())
	}
}

// NewGetClustersIDNodesNodeIDLogsLogNameOK creates a GetClustersIDNodesNodeIDLogsLogNameOK with default headers values
func NewGetClustersIDNodesNodeIDLogsLogNameOK() *GetClustersIDNodesNodeIDLogsLogNameOK {
	return &GetClustersIDNodesNodeIDLogsLogNameOK{}
}

/*
GetClustersIDNodesNodeIDLogsLogNameOK describes a response with status code 200, with default header values.

Response containing log file messages
*/
type GetClustersIDNodesNodeIDLogsLogNameOK struct {
	Payload []*models.ClusterNodeLogMessage
}

// IsSuccess returns true when this get clusters Id nodes node Id logs log name o k response has a 2xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get clusters Id nodes node Id logs log name o k response has a 3xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id nodes node Id logs log name o k response has a 4xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clusters Id nodes node Id logs log name o k response has a 5xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters Id nodes node Id logs log name o k response a status code equal to that given
func (o *GetClustersIDNodesNodeIDLogsLogNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get clusters Id nodes node Id logs log name o k response
func (o *GetClustersIDNodesNodeIDLogsLogNameOK) Code() int {
	return 200
}

func (o *GetClustersIDNodesNodeIDLogsLogNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}][%d] getClustersIdNodesNodeIdLogsLogNameOK %s", 200, payload)
}

func (o *GetClustersIDNodesNodeIDLogsLogNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}][%d] getClustersIdNodesNodeIdLogsLogNameOK %s", 200, payload)
}

func (o *GetClustersIDNodesNodeIDLogsLogNameOK) GetPayload() []*models.ClusterNodeLogMessage {
	return o.Payload
}

func (o *GetClustersIDNodesNodeIDLogsLogNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersIDNodesNodeIDLogsLogNameBadRequest creates a GetClustersIDNodesNodeIDLogsLogNameBadRequest with default headers values
func NewGetClustersIDNodesNodeIDLogsLogNameBadRequest() *GetClustersIDNodesNodeIDLogsLogNameBadRequest {
	return &GetClustersIDNodesNodeIDLogsLogNameBadRequest{}
}

/*
GetClustersIDNodesNodeIDLogsLogNameBadRequest describes a response with status code 400, with default header values.

Bad request.
*/
type GetClustersIDNodesNodeIDLogsLogNameBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters Id nodes node Id logs log name bad request response has a 2xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters Id nodes node Id logs log name bad request response has a 3xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id nodes node Id logs log name bad request response has a 4xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters Id nodes node Id logs log name bad request response has a 5xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters Id nodes node Id logs log name bad request response a status code equal to that given
func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get clusters Id nodes node Id logs log name bad request response
func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) Code() int {
	return 400
}

func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}][%d] getClustersIdNodesNodeIdLogsLogNameBadRequest %s", 400, payload)
}

func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}][%d] getClustersIdNodesNodeIdLogsLogNameBadRequest %s", 400, payload)
}

func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersIDNodesNodeIDLogsLogNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersIDNodesNodeIDLogsLogNameUnauthorized creates a GetClustersIDNodesNodeIDLogsLogNameUnauthorized with default headers values
func NewGetClustersIDNodesNodeIDLogsLogNameUnauthorized() *GetClustersIDNodesNodeIDLogsLogNameUnauthorized {
	return &GetClustersIDNodesNodeIDLogsLogNameUnauthorized{}
}

/*
GetClustersIDNodesNodeIDLogsLogNameUnauthorized describes a response with status code 401, with default header values.

Authorization information is missing or invalid.
*/
type GetClustersIDNodesNodeIDLogsLogNameUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clusters Id nodes node Id logs log name unauthorized response has a 2xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters Id nodes node Id logs log name unauthorized response has a 3xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters Id nodes node Id logs log name unauthorized response has a 4xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters Id nodes node Id logs log name unauthorized response has a 5xx status code
func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters Id nodes node Id logs log name unauthorized response a status code equal to that given
func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get clusters Id nodes node Id logs log name unauthorized response
func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) Code() int {
	return 401
}

func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}][%d] getClustersIdNodesNodeIdLogsLogNameUnauthorized %s", 401, payload)
}

func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/{id}/nodes/{node_id}/logs/{log_name}][%d] getClustersIdNodesNodeIdLogsLogNameUnauthorized %s", 401, payload)
}

func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersIDNodesNodeIDLogsLogNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
