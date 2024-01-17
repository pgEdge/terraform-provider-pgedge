package client

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/models"
	"github.com/stretchr/testify/assert"
)

const (
	BaseUrl      = "" //your base url here
	ClusterID    = "" //your cluster id here
	ClientID     = "" //your client id here
	ClientSecret = "" //your client secret here
)

var (
	AccessToken *string
	DatabaseID  *strfmt.UUID
)

func TestOAuthToken(t *testing.T) {
	client := NewClient(BaseUrl, "", ClusterID)

	token, err := client.OAuthToken(context.Background(), ClientID, ClientSecret)
	if err == nil {
		AccessToken = &token.AccessToken
	}

	assert.Nil(t, err)
}

func TestGetDatabases(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken, ClusterID)
	_, err := client.GetDatabases(context.Background())

	assert.Nil(t, err)
}

func TestCreateDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken, ClusterID)

	request := &models.DatabaseCreationRequest{
		Name: "test",
	}

	database, err := client.CreateDatabase(context.Background(), request)
	DatabaseID = &database.ID

	assert.Nil(t, err)
}

func TestGetDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken, ClusterID)
	_, err := client.GetDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}

func TestReplicateDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken, ClusterID)

	_, err := client.ReplicateDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}

func TestDeleteDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken, ClusterID)

	err := client.DeleteDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}
