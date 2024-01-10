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

    request.SetAuthorization(c.AuthHeader)

    resp, err := c.PgEdgeAPIClient.Operations.GetDatabases(request)
    if err != nil {
        fmt.Println("err: ", err)
        return nil, err
    }

    return resp.Payload, nil
}


func (c *Client) GetDatabase(ctx context.Context, id strfmt.UUID) (*models.DatabaseDetails, error) {
	fmt.Println("id: ", id)
	request := &operations.GetDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID: id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetDatabasesID(request)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) CreateDatabase(ctx context.Context, database *models.DatabaseCreationRequest) (*models.DatabaseCreationResponse, error) {
	request := &operations.PostDatabasesParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:   database,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostDatabases(request)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) DeleteDatabase(ctx context.Context, id strfmt.UUID) (error) {
	request := &operations.DeleteDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID: id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteDatabasesID(request)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	return nil
}

func (c *Client) ReplicateDatabase(ctx context.Context, id strfmt.UUID) (*models.ReplicationResponse, error) {
	request := &operations.PostDatabasesIDReplicateParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID: id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostDatabasesIDReplicate(request)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) OAuthToken(ctx context.Context) (*operations.PostOauthTokenOKBody, error) {
	request := &operations.PostOauthTokenParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ClientID: "CIzx5xcvt9MFRYVIoFl7Bz9Kl8ryNSdh",
		ClientSecret: "XqRDtkdyyVKNjjT-NiDXdP-ovAJMEmTqKlbMD89WonZhRLyQocKA11rddxw85H8r",
	}

	fmt.Println("clientId: ", request.ClientID)
	fmt.Println("clientId: ", request.ClientSecret)

	resp, err := c.PgEdgeAPIClient.Operations.PostOauthToken(request)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	return resp.Payload, nil
}
