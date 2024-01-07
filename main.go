package main

import (
	// "context"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "net/url"
)


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