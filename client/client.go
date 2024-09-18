package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sanity-io/litter"

	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/pgEdge/terraform-provider-pgedge/client/operations"

	httptransport "github.com/go-openapi/runtime/client"
)

type Client struct {
	AuthHeader      string
	HTTPClient      *http.Client
	PgEdgeAPIClient *PgEdgeAPI
}

func NewClient(baseUrl, authHeader string) *Client {
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
	if c.PgEdgeAPIClient == nil {
		return nil, fmt.Errorf("PgEdgeAPIClient is nil")
	}

	if c.PgEdgeAPIClient.Operations == nil {
		return nil, fmt.Errorf("Operations is nil")
	}

	request := &operations.GetDatabasesParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetDatabases(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetDatabase(ctx context.Context, id strfmt.UUID) (*models.DatabaseDetails, error) {
	request := &operations.GetDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetDatabasesID(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) CreateDatabase(ctx context.Context, database *models.DatabaseCreationRequest) (*models.DatabaseCreationResponse, error) {
	request := &operations.PostDatabasesParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:       database,
	}
	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostDatabases(request)
	if err != nil {
		return nil, err
	}

	for {
		databaseDetails, err := c.GetDatabase(ctx, resp.Payload.ID)
		if err != nil {
			return nil, err
		}

		switch databaseDetails.Status {
		case "available":
			litter.Dump(databaseDetails)
			return resp.Payload, nil
		case "failed":
			return nil, errors.New("database creation failed")
		case "creating":
			time.Sleep(5 * time.Second)
		default:
			return nil, errors.New("unexpected database status")
		}
	}
}

func (c *Client) UpdateDatabase(ctx context.Context, id strfmt.UUID, body *models.DatabaseUpdateRequest) (*models.DatabaseDetails, error) {
	request := &operations.PatchDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
		Body:       body,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PatchDatabasesID(request)
	if err != nil {
		return nil, err
	}

	// for {
	databaseDetails, err := c.GetDatabase(ctx, resp.Payload.ID)
	if err != nil {
		return nil, err
	}

	fmt.Println(databaseDetails, "databse details")

	// switch databaseDetails.Status {
	// case "available":
	return resp.Payload, nil
	// case "failed":
	// return nil, errors.New("database creation failed")
	// case "creating":
	// time.Sleep(5 * time.Second)
	// default:
	// return nil, errors.New("unexpected database status")
	// }
	// }
}

func (c *Client) DeleteDatabase(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteDatabasesID(request)
	if strings.Contains(err.Error(), "200") {
		for {
			_, err := c.GetDatabase(ctx, id)
			if err != nil {
				return nil
			}

			time.Sleep(5 * time.Second)
		}
	}

	return err
}

func (c *Client) ReplicateDatabase(ctx context.Context, id strfmt.UUID) (*models.ReplicationResponse, error) {
	request := &operations.PostDatabasesIDReplicateParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostDatabasesIDReplicate(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetAllClusters(ctx context.Context) ([]*models.Cluster, error) {
	request := &operations.GetClustersParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetClusters(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetCluster(ctx context.Context, id strfmt.UUID) (*models.Cluster, error) {
	request := &operations.GetClustersIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetClustersID(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) CreateCluster(ctx context.Context, cluster *models.CreateClusterInput) (*models.Cluster, error) {
	request := &operations.PostClustersParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:       cluster,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostClusters(request)
	if err != nil {
		return nil, err
	}

	for {
		clusterDetails, err := c.GetCluster(ctx, strfmt.UUID(*resp.Payload.ID))
		if err != nil {
			return nil, err
		}

		switch *clusterDetails.Status {
		case "available":
			return clusterDetails, nil
		case "failed":
			return nil, errors.New("cluster creation failed")
		case "creating":
			time.Sleep(5 * time.Second)
		default:
			return nil, fmt.Errorf("unexpected cluster status: %s", *clusterDetails.Status)
		}
	}
}

func (c *Client) UpdateCluster(ctx context.Context, id strfmt.UUID, body *models.UpdateClusterInput) (*models.Cluster, error) {
	request := &operations.PatchClustersIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
		Body:       body,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PatchClustersID(request)
	if err != nil {
		return nil, err
	}

	for {
		clusterDetails, err := c.GetCluster(ctx, strfmt.UUID(*resp.Payload.ID))
		if err != nil {
			return nil, err
		}

		switch *clusterDetails.Status {
		case "available":
			return clusterDetails, nil
		case "failed":
			return nil, errors.New("cluster creation failed")
		case "modifying":
			time.Sleep(5 * time.Second)
		default:
			return nil, fmt.Errorf("unexpected cluster status: %s", *clusterDetails.Status)
		}
	}
}

func (c *Client) DeleteCluster(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteClustersIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteClustersID(request)

	if strings.Contains(err.Error(), "200") {
		for {
			_, err := c.GetCluster(ctx, id)
			if err != nil {
				return nil
			}
			time.Sleep(5 * time.Second)

		}
	}

	return nil
}

func (c *Client) GetClusterNodes(ctx context.Context, id strfmt.UUID, nearLat, nearLon, orderBy *string) ([]*models.ClusterNode, error) {
	request := &operations.GetClustersIDNodesParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
		NearLat:    nearLat,
		NearLon:    nearLon,
		OrderBy:    orderBy,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetClustersIDNodes(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetClusterNodeLogs(ctx context.Context, clusterID, nodeID strfmt.UUID, logName string, params *operations.GetClustersIDNodesNodeIDLogsLogNameParams) ([]*models.ClusterNodeLogMessage, error) {
	request := &operations.GetClustersIDNodesNodeIDLogsLogNameParams{
		HTTPClient:    c.HTTPClient,
		Context:       ctx,
		ID:            clusterID.String(),
		NodeID:        nodeID.String(),
		LogName:       logName,
		Lines:         params.Lines,
		Since:         params.Since,
		Until:         params.Until,
		Priority:      params.Priority,
		Grep:          params.Grep,
		CaseSensitive: params.CaseSensitive,
		Reverse:       params.Reverse,
		Dmesg:         params.Dmesg,
		Output:        params.Output,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetClustersIDNodesNodeIDLogsLogName(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetCloudAccounts(ctx context.Context) ([]*models.CloudAccount, error) {
	if c.PgEdgeAPIClient == nil {
		return nil, fmt.Errorf("PgEdgeAPIClient is nil")
	}

	request := &operations.GetCloudAccountsParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetCloudAccounts(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetCloudAccount(ctx context.Context, id strfmt.UUID) (*models.CloudAccount, error) {
	request := &operations.GetCloudAccountsIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetCloudAccountsID(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) CreateCloudAccount(ctx context.Context, account *models.CreateCloudAccountInput) (*models.CloudAccount, error) {
	request := &operations.PostCloudAccountsParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:       account,
	}
	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostCloudAccounts(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) DeleteCloudAccount(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteCloudAccountsIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteCloudAccountsID(request)
	if err != nil {
		// The API returns 200 OK for successful deletion
		if strings.Contains(err.Error(), "200") {
			return nil
		}
		return err
	}

	return nil
}


func (c *Client) OAuthToken(ctx context.Context, clientId, clientSecret, grantType string) (*operations.PostOauthTokenOKBody, error) {
	request := &operations.PostOauthTokenParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body: operations.PostOauthTokenBody{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			GrantType:    grantType,
		},
	}

	resp, err := c.PgEdgeAPIClient.Operations.PostOauthToken(request)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
