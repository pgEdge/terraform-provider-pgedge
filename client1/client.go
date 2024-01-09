package client1

import (
	// "context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/client/operations"
)

const (
    BaseURLV1 = "https://api.facest.io/v1"
)

type Client struct {
    BaseURL    string
    apiKey     string
    HTTPClient *http.Client
}

type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type successResponse struct {
    Code int         `json:"code"`
    Data interface{} `json:"data"`
}

func NewClient(apiKey string) *Client {
    return &Client{
        BaseURL: BaseURLV1,
        apiKey:  apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
    }
}

func GetDatabases() {
	client := operations.New(nil, strfmt.Default)

	_, err := client.GetDatabases(nil)
	if err != nil {
		log.Fatalf("Failed to get databases: %v", err)
	}

}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    req.Header.Set("Accept", "application/json; charset=utf-8")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }

    defer res.Body.Close()

    if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
            return errors.New(errRes.Message)
        }

        return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
    }

    fullResponse := successResponse{
        Data: v,
    }
    if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
        return err
    }

    return nil
}