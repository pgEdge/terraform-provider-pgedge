package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// "github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/client/operations"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/pgEdge/terraform-provider-pgedge/client"
)

func main() {
	// create the transport
	transport := httptransport.New("devapi.pgedge.com:443", "", []string{"https"})

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)
	request := &operations.GetDatabasesParams{
		HTTPClient: &http.Client{},
		Context:    context.Background(),
	}

	request.SetAuthorization("Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODI1Njg3LCJleHAiOjE3MDQ5MTIwODcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.lxjj0MD44mgUl-PQonEyVSKB0qVdin__gUHoFgJJ1-gQsfmz86cyP0Iq7Rw1oFU4fRRonf9tOiu9uMSrKQC2sV-9ANt7nQWLSJuwVop1kxkCTUXJH2R6TglQ0kW-SsDNU1P-VezqOUhUD07LqvUytWYtJKKXilb_fWk6LnWBGyZ2Gmw9ioOdEuMjkJsuMc_HIiES28w1kA5NdTiayqo_Gdr434VT3UHW0kFD5ox9tIoR4JxIHk8cnlU7Oj7E3B-kTj8SJqCjEOweu2fw9bcfgDG2Dr1qLGtPij2RCU8iewYS9Kp7Nc0IzLFfLHJgh6UAEWBSomE8-q9_xzYxd1gi9Q")

	// make the request to get all items
	resp, err := client.Operations.GetDatabases(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Payload[0])
}
