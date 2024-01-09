package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/client/operations"
	"github.com/pgEdge/terraform-provider-pgedge/models"

	httptransport "github.com/go-openapi/runtime/client"
)

type Client struct {
	AuthHeader      string
	HTTPClient      *http.Client
	PgEdgeAPIClient *PgEdgeAPI
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewClient(baseUrl,authHeader string) *Client {
	var url string
    var schemas []string
	if baseUrl == "" {
		url = "localhost"
	} else {
		url = baseUrl
        schemas = strings.Split(url, "://")
	}

    if strings.HasPrefix(url, "https") {
		url += ":443"
	}

	url = strings.ReplaceAll(url, "http://", "")
	url = strings.ReplaceAll(url, "https://", "")

    fmt.Println("url: ", url)

	transport := httptransport.New(url, "", schemas)
	client := New(transport, strfmt.Default)

	return &Client{
		AuthHeader: authHeader,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		PgEdgeAPIClient: client,
	}
}

func (c *Client) GetDatabases(ctx context.Context) ([]*models.Database, error) {
	request := &operations.GetDatabasesParams{
        HTTPClient: c.HTTPClient,
        Context:    ctx,
    }

    fmt.Println("c.AuthHeader: ", request.Authorization)

    request.SetAuthorization(c.AuthHeader)

    fmt.Println("c.AuthHeader: ", request.Authorization)


    resp, err := c.PgEdgeAPIClient.Operations.GetDatabases(request)
    if err != nil {
        fmt.Println("err: ", err)
        return nil, err
    }

    return resp.Payload, nil
}
