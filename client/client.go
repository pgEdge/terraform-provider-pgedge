package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
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
	var schemes []string

	if baseUrl == "" {
		url = "https://api.pgedge.com/v1"
	} else {
		url = baseUrl
		url = strings.TrimSuffix(url, "/")

		if !strings.HasSuffix(url, "/v1") {
			url += "/v1"
		}
	}

	if strings.Contains(url, "://") {
		parts := strings.SplitN(url, "://", 2)
		schemes = []string{parts[0]}
		url = parts[1]
	} else {
		schemes = []string{"https"}
	}

	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")

	hostAndPath := strings.SplitN(url, "/", 2)
	host := hostAndPath[0]
	path := ""
	if len(hostAndPath) > 1 {
		path = "/" + hostAndPath[1]
	}

	transport := httptransport.New(host, path, schemes)
	client := New(transport, strfmt.Default)

	return &Client{
		AuthHeader: authHeader,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		PgEdgeAPIClient: client,
	}
}

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

type GeneratedAPIError interface {
	Code() int
	Message() string
}

type KnownError interface {
	GetPayload() *models.Error
}

func handleAPIError(err error) error {
	var knownErr KnownError
	if errors.As(err, &knownErr) {
		payload := knownErr.GetPayload()
		return &APIError{
			StatusCode: int(payload.Code),
			Message:    payload.Message,
		}
	}

	var runtimeErr *runtime.APIError
	if errors.As(err, &runtimeErr) {
		return &APIError{
			StatusCode: runtimeErr.Code,
			Message:    runtimeErr.Error(),
		}
	}

	return &APIError{
		StatusCode: 500,
		Message:    err.Error(),
	}
}

type TaskPollingConfig struct {
	SubjectID   string
	SubjectKind string
	MaxAttempts int
	Interval    time.Duration
}

func (c *Client) GetTasks(ctx context.Context, subjectID, subjectKind string, id, name *string, status *string, limit, offset *int64) ([]*models.Task, error) {
	request := &operations.GetTasksParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	if subjectID != "" {
		request.SubjectID = &subjectID
	}
	if subjectKind != "" {
		request.SubjectKind = &subjectKind
	}
	if id != nil {
		request.ID = id
	}
	if name != nil {
		request.Name = name
	}
	if status != nil {
		request.Status = status
	}
	if limit != nil {
		request.Limit = limit
	}
	if offset != nil {
		request.Offset = offset
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetTasks(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) PollTaskStatus(ctx context.Context, config TaskPollingConfig) error {
	var taskID *string
	attempt := 0

	for {
		if attempt >= config.MaxAttempts {
			if taskID == nil {
				return fmt.Errorf("no task found for %s %s after %d attempts",
					config.SubjectKind, config.SubjectID, attempt)
			}
			return fmt.Errorf("timeout waiting for task %s to complete", *taskID)
		}

		// Locate the most recent task if we don't have one
		if taskID == nil {
			tasks, err := c.GetTasks(ctx, config.SubjectID, config.SubjectKind, nil, nil, nil, nil, nil)

			if err != nil {
				return fmt.Errorf("error checking task status: %w", err)
			}

			var latestTime time.Time
			for _, task := range tasks {
				if task.Status != nil && (*task.Status == "running" || *task.Status == "queued") {
					taskTime, err := time.Parse(time.RFC3339, *task.CreatedAt)
					if err != nil {
						continue
					}
					if taskID == nil || taskTime.After(latestTime) {
						taskID = task.ID
						latestTime = taskTime
					}
				}
			}
			if taskID == nil {
				if attempt < 4 {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(2 * time.Second):
					}
					attempt++
					continue
				}
				return fmt.Errorf("no active task found for %s %s",
					config.SubjectKind, config.SubjectID)
			}
			continue
		}

		// Poll the task by ID
		tasks, err := c.GetTasks(ctx, config.SubjectID, config.SubjectKind, taskID, nil, nil, nil, nil)
		if len(tasks) == 0 {
			return fmt.Errorf("task %s not found", *taskID)
		}

		if err != nil {
			return fmt.Errorf("error checking task status: %w", err)
		}

		task := tasks[0]
		if task.Status == nil {
			return fmt.Errorf("task %s has no status", *taskID)
		}

		switch *task.Status {
		case "succeeded":
			return nil
		case "failed":
			if task.Error != "" {
				return fmt.Errorf("task failed: %s", task.Error)
			}
			return fmt.Errorf("task failed without error message")
		case "running", "queued":
			// Continue polling
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(config.Interval):
		}
		attempt++
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
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) CreateDatabase(ctx context.Context, database *models.CreateDatabaseInput) (*models.Database, error) {
	request := &operations.PostDatabasesParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:       database,
	}
	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostDatabases(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   resp.Payload.ID.String(),
		SubjectKind: "database",
		MaxAttempts: 360, // 30 minutes
		Interval:    5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return c.GetDatabase(ctx, *resp.Payload.ID)
}

func (c *Client) GetDatabase(ctx context.Context, id strfmt.UUID) (*models.Database, error) {
	request := &operations.GetDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetDatabasesID(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) UpdateDatabase(ctx context.Context, id strfmt.UUID, body *models.UpdateDatabaseInput) (*models.Database, error) {
	request := &operations.PatchDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
		Body:       body,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PatchDatabasesID(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   resp.Payload.ID.String(),
		SubjectKind: "database",
		MaxAttempts: 360, // 30 minutes
		Interval:    5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return c.GetDatabase(ctx, id)
}

func (c *Client) DeleteDatabase(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteDatabasesIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteDatabasesID(request)
	if err != nil {
		return handleAPIError(err)
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   id.String(),
		SubjectKind: "database",
		MaxAttempts: 360, // 30 minutes
		Interval:    5 * time.Second,
	})

	return err
}

func (c *Client) GetAllClusters(ctx context.Context) ([]*models.Cluster, error) {
	request := &operations.GetClustersParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetClusters(request)
	if err != nil {
		return nil, handleAPIError(err)
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
		return nil, handleAPIError(err)
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
		return nil, handleAPIError(err)
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   resp.Payload.ID.String(),
		SubjectKind: "cluster",
		MaxAttempts: 540, // 45 minutes
		Interval:    5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return c.GetCluster(ctx, *resp.Payload.ID)
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
		return nil, handleAPIError(err)
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   resp.Payload.ID.String(),
		SubjectKind: "cluster",
		MaxAttempts: 540, // 45 minutes
		Interval:    5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return c.GetCluster(ctx, strfmt.UUID(*resp.Payload.ID))
}

func (c *Client) DeleteCluster(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteClustersIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteClustersID(request)
	if err != nil {
		return handleAPIError(err)
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   id.String(),
		SubjectKind: "cluster",
		MaxAttempts: 540, // 45 minutes
		Interval:    5 * time.Second,
	})

	return err
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
		return nil, handleAPIError(err)
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
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) GetCloudAccounts(ctx context.Context) ([]*models.CloudAccount, error) {
	request := &operations.GetCloudAccountsParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetCloudAccounts(request)
	if err != nil {
		return nil, handleAPIError(err)
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
		return nil, handleAPIError(err)
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
		return nil, handleAPIError(err)
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
		return handleAPIError(err)
	}

	return nil
}

func (c *Client) GetSSHKeys(ctx context.Context) ([]*models.SSHKey, error) {
	request := &operations.GetSSHKeysParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetSSHKeys(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) CreateSSHKey(ctx context.Context, sshKey *models.CreateSSHKeyInput) (*models.SSHKey, error) {
	request := &operations.PostSSHKeysParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:       sshKey,
	}
	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostSSHKeys(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) GetSSHKey(ctx context.Context, id strfmt.UUID) (*models.SSHKey, error) {
	request := &operations.GetSSHKeysIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetSSHKeysID(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) DeleteSSHKey(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteSSHKeysIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteSSHKeysID(request)
	if err != nil {
		return handleAPIError(err)
	}
	return nil
}

func (c *Client) GetBackupStores(ctx context.Context, createdAfter, createdBefore *string, limit, offset *int64, descending *bool) ([]*models.BackupStore, error) {
	request := &operations.GetBackupStoresParams{
		HTTPClient:    c.HTTPClient,
		Context:       ctx,
		CreatedAfter:  createdAfter,
		CreatedBefore: createdBefore,
		Limit:         limit,
		Offset:        offset,
		Descending:    descending,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetBackupStores(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) CreateBackupStore(ctx context.Context, input *models.CreateBackupStoreInput) (*models.BackupStore, error) {
	request := &operations.PostBackupStoresParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		Body:       input,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.PostBackupStores(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	if resp == nil || resp.Payload == nil {
		return nil, fmt.Errorf("received nil response or payload")
	}

	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   resp.Payload.ID.String(),
		SubjectKind: "backup_store",
		MaxAttempts: 360, // 30 minutes
		Interval:    5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return c.GetBackupStore(ctx, *resp.Payload.ID)
}

func (c *Client) GetBackupStore(ctx context.Context, id strfmt.UUID) (*models.BackupStore, error) {
	request := &operations.GetBackupStoresIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	resp, err := c.PgEdgeAPIClient.Operations.GetBackupStoresID(request)
	if err != nil {
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}

func (c *Client) DeleteBackupStore(ctx context.Context, id strfmt.UUID) error {
	request := &operations.DeleteBackupStoresIDParams{
		HTTPClient: c.HTTPClient,
		Context:    ctx,
		ID:         id,
	}

	request.SetAuthorization(c.AuthHeader)

	_, err := c.PgEdgeAPIClient.Operations.DeleteBackupStoresID(request)
	if err != nil {
		return handleAPIError(err)
	}

	// Poll for task completion
	err = c.PollTaskStatus(ctx, TaskPollingConfig{
		SubjectID:   id.String(),
		SubjectKind: "backup_store",
		MaxAttempts: 360, // 30 minutes
		Interval:    5 * time.Second,
	})

	return err
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
		return nil, handleAPIError(err)
	}

	return resp.Payload, nil
}
