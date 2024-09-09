package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/pgEdge/terraform-provider-pgedge/client/operations"
	"github.com/stretchr/testify/assert"
)

var (
	BaseUrl      = os.Getenv("PGEDGE_BASE_URL")      //your base url here
	ClientID     = os.Getenv("PGEDGE_CLIENT_ID")     //your client id here
	ClientSecret = os.Getenv("PGEDGE_CLIENT_SECRET") //your client secret here
)

var (
	AccessToken    *string
	DatabaseID     *strfmt.UUID
	ClusterID      *strfmt.UUID
	CloudAccountID *strfmt.UUID
)

func TestOAuthToken(t *testing.T) {
	client := NewClient(BaseUrl, "")

	token, err := client.OAuthToken(context.Background(), ClientID, ClientSecret, "")
	if err == nil {
		AccessToken = &token.AccessToken
	}

	assert.Nil(t, err)
}

func TestCreateCloudAccount(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	accountType := "aws"
	request := &models.CreateCloudAccountInput{
		Name: "TestAccount",
		Type: &accountType,
		Credentials: map[string]interface{}{
			"role_arn": os.Getenv("PGEDGE_ROLE_ARN"), //your role arn here
		},
	}

	account, err := client.CreateCloudAccount(context.Background(), request)
	CloudAccountID = account.ID

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "TestAccount", *account.Name)
}

func stringPtr(s string) *string {
	return &s
}
func int64Ptr(i int64) *int64 {
	return &i
}
func TestCreateCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	request := &models.CreateClusterInput{
		Name:           stringPtr("test-cluster1"),
		CloudAccountID: CloudAccountID.String(),
		Regions:        []string{"us-east-2", "us-west-2"},
		Nodes: []*models.ClusterNodeSettings{
			{
				Name:             "n1",
				Region:           stringPtr("us-east-2"),
				Image:            "postgres",
				InstanceType:     "t4g.small",
				AvailabilityZone: "us-east-2a",
				VolumeType:       "gp2",
			},
			{
				Name:             "n2",
				Region:           stringPtr("us-west-2"),
				Image:            "postgres",
				InstanceType:     "t4g.medium",
				AvailabilityZone: "us-west-2a",
				VolumeType:       "gp2",
			},
		},
		Networks: []*models.ClusterNetworkSettings{
			{
				Region:        stringPtr("us-east-2"),
				Cidr:          "10.1.0.0/16",
				PublicSubnets: []string{"10.1.0.0/24"},
			},
			{
				Region:        stringPtr("us-west-2"),
				Cidr:          "10.2.0.0/16",
				PublicSubnets: []string{"10.2.0.0/24"},
			},
		},
		FirewallRules: []*models.ClusterFirewallRuleSettings{
			{
				Name:    "postgres",
				Port:    int64Ptr(5432),
				Sources: []string{"0.0.0.0/0"},
			},
		},
		NodeLocation: stringPtr("public"),
		ResourceTags: map[string]string{
			"key": "value",
		},
	}

	cluster, err := client.CreateCluster(context.Background(), request)
	assert.Nil(t, err)
	assert.NotNil(t, cluster)
	assert.Equal(t, "available", *cluster.Status)
	ClusterID = cluster.ID
}

func TestGetCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	cluster, err := client.GetCluster(context.Background(), *ClusterID)

	assert.Nil(t, err)
	assert.NotNil(t, cluster)
	assert.Equal(t, ClusterID.String(), cluster.ID.String())
}

// func TestUpdateCluster(t *testing.T) {
// 	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

// 	request := &models.UpdateClusterInput{
// 		Regions: []string{"us-east-2", "us-west-2"},
// 		Nodes: []*models.ClusterNodeSettings{
// 			{
// 				Name:             "n1",
// 				Region:           stringPtr("us-east-2"),
// 				InstanceType:     "t4g.medium",
// 				AvailabilityZone: "us-east-2a",
// 				VolumeType:       "gp2",
// 			},
// 			{
// 				Name:             "n2",
// 				Region:           stringPtr("us-west-2"),
// 				InstanceType:     "t4g.medium",
// 				AvailabilityZone: "us-west-2a",
// 				VolumeType:       "gp2",
// 			},
// 		},
// 	}

// 	cluster, err := client.UpdateCluster(context.Background(), *ClusterID, request)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, cluster)
// 	assert.Equal(t, 2, len(cluster.Nodes))
// }

func TestGetClusterNodes(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	nodes, err := client.GetClusterNodes(context.Background(), *ClusterID, nil, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, nodes)
	assert.Equal(t, 2, len(nodes))
}

func TestGetClusterNodeLogs(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	// Assuming the first node in the cluster
	nodes, _ := client.GetClusterNodes(context.Background(), *ClusterID, nil, nil, nil)
	nodeID := nodes[0].ID

	logs, err := client.GetClusterNodeLogs(context.Background(), *ClusterID, strfmt.UUID(*nodeID), "postgresql", &operations.GetClustersIDNodesNodeIDLogsLogNameParams{
		Lines: int64Ptr(10),
	})

	assert.Nil(t, err)
	assert.NotNil(t, logs)
	assert.Greater(t, len(logs), 0)
}

func TestGetAllClusters(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetAllClusters(context.Background())

	assert.Nil(t, err)
}

func TestCreateDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	request := &models.DatabaseCreationRequest{
		Name:      "db5",
		ClusterID: *ClusterID,
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

func TestUpdateDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	request := &models.DatabaseUpdateRequest{
		Extensions: &models.DatabaseUpdateRequestExtensions{
			AutoManage: false,
		},
	}

	database, err := client.UpdateDatabase(context.Background(), *DatabaseID, request)
	DatabaseID = &database.ID

	assert.Nil(t, err)
}

func TestGetDatabases(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetDatabases(context.Background())

	assert.Nil(t, err)
}

// func TestReplicateDatabase(t *testing.T) {
// 	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

// 	_, err := client.ReplicateDatabase(context.Background(), *DatabaseID)

// 	assert.Nil(t, err)
// }

func TestDeleteDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	err := client.DeleteDatabase(context.Background(), *DatabaseID)

	assert.Nil(t, err)
}

func TestDeleteCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	err := client.DeleteCluster(context.Background(), *ClusterID)
	assert.Nil(t, err)

	// Verify that the cluster is deleted
	time.Sleep(30 * time.Second) // Give some time for deletion to propagate
	_, err = client.GetCluster(context.Background(), *ClusterID)
	assert.NotNil(t, err) // Expect an error as the cluster should not exist
}

func TestGetCloudAccounts(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	accounts, err := client.GetCloudAccounts(context.Background())

	assert.Nil(t, err)
	assert.NotEmpty(t, accounts)
}

func TestGetCloudAccount(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	account, err := client.GetCloudAccount(context.Background(), *CloudAccountID)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, CloudAccountID.String(), account.ID.String())
}

func TestDeleteCloudAccount(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	err := client.DeleteCloudAccount(context.Background(), *CloudAccountID)

	assert.Nil(t, err)

}
