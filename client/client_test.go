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
	ClientID     = "" //your client id here
	ClientSecret = "" //your client secret here
	CloudAccountID = "" //your cloud account id here
)

var (
	AccessToken *string
	DatabaseID  *strfmt.UUID
	ClusterID   = "" //your cluster id here

)

func TestOAuthToken(t *testing.T) {
	client := NewClient(BaseUrl, "")

	token, err := client.OAuthToken(context.Background(), ClientID, ClientSecret)
	if err == nil {
		AccessToken = &token.AccessToken
	}

	assert.Nil(t, err)
}

func TestGetDatabases(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetDatabases(context.Background())

	assert.Nil(t, err)
}

func TestCreateDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	request := &models.DatabaseCreationRequest{
		Name:      "test",
		ClusterID: ClusterID,
	}

	database, err := client.CreateDatabase(context.Background(), request)
	DatabaseID = &database.ID

	assert.Nil(t, err)
}

func TestGetDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}

func TestReplicateDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	_, err := client.ReplicateDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}

func TestDeleteDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	err := client.DeleteDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}

func TestGetAllClusters(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetAllClusters(context.Background())

	assert.Nil(t, err)
}

func TestCreateCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	var aws []*models.NodeGroup
	aws = append(aws, &models.NodeGroup{
		InstanceType: "t4g.small",
		Region: "us-east-1",
		Nodes: []*models.NodeGroupNodesItems0{
			{
				DisplayName: "Node1",
				IsActive:    true,
			},
		},
	})
	request := &models.ClusterCreationRequest{
		Name: "test8",
		NodeGroups: &models.ClusterCreationRequestNodeGroups{
			Aws: aws,
			Azure:  []*models.NodeGroup{},
			Google: []*models.NodeGroup{},
		},
		Firewall: &models.ClusterCreationRequestFirewall{
			Rules: []*models.ClusterCreationRequestFirewallRulesItems0{
				{
					Type:    "https",
					Port:    5432,
					Sources: []string{"0.0.0.0/0"},
				},
			},
		},

		CloudAccountID: CloudAccountID,
	}

	cluster, err := client.CreateCluster(context.Background(), request)
	ClusterID = cluster.ID

	assert.Nil(t, err)
}

func TestGetCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetCluster(context.Background(), strfmt.UUID(ClusterID))

	assert.Nil(t, err)
}

func TestDeleteCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	err := client.DeleteCluster(context.Background(), strfmt.UUID(ClusterID))

	assert.Nil(t, err)
}
