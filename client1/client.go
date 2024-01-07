package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
    BaseURLV1 = "https://api.facest.io/v1"
)

type Client struct {
    BaseURL    string
    apiKey     string
    HTTPClient *http.Client
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

type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type successResponse struct {
    Code int         `json:"code"`
    Data interface{} `json:"data"`
}

type ApiResponse []struct {
	CreatedAt string `json:"created_at"`
	Domain  string `json:"domain"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Nodes   []Nodes `json:"nodes"`
	Status   string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type Nodes struct {
	Connection Connection `json:"connection"`
	Location  Location  `json:"location"`
	Name      string    `json:"name"`
}

type Connection struct {
	Database string `json:"database"`
	Host   string `json:"host"`
	Password string `json:"password"`
	Port   int   `json:"port"`
	Username string `json:"username"`
}

type Location struct {
	Code     string `json:"code"`
	Country  string `json:"country"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	RegionCode string `json:"region_code"`
}


type FacesList struct {
    Count      int    `json:"count"`
    PagesCount int    `json:"pages_count"`
    Faces      []Face `json:"faces"`
}

type Face struct {
    FaceToken  string      `json:"face_token"`
    FaceID     string      `json:"face_id"`
    FaceImages []FaceImage `json:"face_images"`
    CreatedAt  time.Time   `json:"created_at"`
}

type FaceImage struct {
    ImageToken string    `json:"image_token"`
    ImageURL   string    `json:"image_url"`
    CreatedAt  time.Time `json:"created_at"`
}

type FacesListOptions struct {
    Limit int `json:"limit"`
    Page  int `json:"page"`
}


func (c *Client) GetFaces(ctx context.Context, options *FacesListOptions) (*FacesList, error) {
    limit := 100
    page := 1
    if options != nil {
        limit = options.Limit
        page = options.Page
    }

    req, err := http.NewRequest("GET", fmt.Sprintf("%s/faces?limit=%d&page=%d", c.BaseURL, limit, page), nil)
    if err != nil {
        return nil, err
    }

    req = req.WithContext(ctx)

    res := FacesList{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
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
