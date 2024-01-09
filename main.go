package main

// "context"
// "fmt"
// "io/ioutil"
// "log"
// "net/http"
// "net/url"

// type DatabaseClient struct {
// 	BaseURL string
// 	Client *http.Client
//  }

// func NewDatabaseClient(baseURL string) (*DatabaseClient, error) {
// 	if _, err := url.ParseRequestURI(baseURL); err != nil {
// 		return nil, err
// 	}

// 	return &DatabaseClient{
// 		BaseURL: baseURL,
// 		Client: &http.Client{},
// 	}, nil
//  }

// func (client *DatabaseClient) GetDatabases(ctx context.Context) (*http.Response, error) {
// 	req, err := http.NewRequestWithContext(ctx, "GET", client.BaseURL+"/databases", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Set headers
// 	req.Header.Set("Accept", "application/json, text/plain, */*")
// 	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
// 	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0Mzk1MTIzLCJleHAiOjE3MDQ0ODE1MjMsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.Vx5cKi6327braC34br-zauQF1no_pI6sJjrBqdVRJdS6_S7nulCooVrhxNI8VGrkIJYstpOWNLwDDIn31zJbN4L30W5XpL0G80F87rWsRiLKrsXEelZzNtJ9UXg5Mt9Xv8kTo97v2FLI2I4W-9mW2B8Qmy0WLrkmnfNFsNxGThGiCF8ywnsUmWg8kewRfUMgsEesAEMSKa0KNNYG1ykk0pW5rq7aGDrr8p67-T-czANb1kWGv7RztmSv8nEclbtwaeoVjnAuWQvjUTbvVQEADSdasBuqU9AYUjM9Cp6SlZB3UmYPUyov1JlEm9HrFxDMt2QyGaIcMsilR_uHhjQPMg")

// 	resp, err := client.Client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
//  }

//  func main() {
// 	// Create a new database client
// 	client, err := NewDatabaseClient("https://devapi.pgedge.com")
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// Create a new context
// 	ctx := context.Background()

// 	// Make a request to get databases
// 	resp, err := client.GetDatabases(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to get databases: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Failed to read response body: %v", err)
// 	}

// 	// Convert the response body to a string and print it
// 	bodyString := string(bodyBytes)
// 	fmt.Println("Response body:", bodyString)

// 	// Print the response status
// 	fmt.Println("Response status:", resp.Status)
// }

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
	HTTPClient:   &http.Client{},
	Context: context.Background(),
  }

  request.SetAuthorization("Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODI1Njg3LCJleHAiOjE3MDQ5MTIwODcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.lxjj0MD44mgUl-PQonEyVSKB0qVdin__gUHoFgJJ1-gQsfmz86cyP0Iq7Rw1oFU4fRRonf9tOiu9uMSrKQC2sV-9ANt7nQWLSJuwVop1kxkCTUXJH2R6TglQ0kW-SsDNU1P-VezqOUhUD07LqvUytWYtJKKXilb_fWk6LnWBGyZ2Gmw9ioOdEuMjkJsuMc_HIiES28w1kA5NdTiayqo_Gdr434VT3UHW0kFD5ox9tIoR4JxIHk8cnlU7Oj7E3B-kTj8SJqCjEOweu2fw9bcfgDG2Dr1qLGtPij2RCU8iewYS9Kp7Nc0IzLFfLHJgh6UAEWBSomE8-q9_xzYxd1gi9Q")
	
	// make the request to get all items
	resp, err := client.Operations.GetDatabases(request)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(resp.Payload[0])
  }