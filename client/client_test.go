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
	SSHKeyID *strfmt.UUID

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

func TestCreateSSHKey(t *testing.T) {
    client := NewClient(BaseUrl, "Bearer "+*AccessToken)

    request := &models.CreateSSHKeyInput{
        Name:      stringPtr("TestSSHKey"),
        PublicKey: stringPtr("ssh-rsa AAAA..."),
    }

    sshKey, err := client.CreateSSHKey(context.Background(), request)
    assert.Nil(t, err)
    assert.NotNil(t, sshKey)
    assert.Equal(t, "TestSSHKey", *sshKey.Name)

    // Store the SSH key ID for use in other tests
    SSHKeyID = sshKey.ID
}

func TestGetSSHKey(t *testing.T) {
    client := NewClient(BaseUrl, "Bearer "+*AccessToken)

    sshKey, err := client.GetSSHKey(context.Background(), *SSHKeyID)
    assert.Nil(t, err)
    assert.NotNil(t, sshKey)
    assert.Equal(t, SSHKeyID.String(), sshKey.ID.String())
}

func TestGetSSHKeys(t *testing.T) {
    client := NewClient(BaseUrl, "Bearer "+*AccessToken)

    sshKeys, err := client.GetSSHKeys(context.Background())
    assert.Nil(t, err)
    assert.NotEmpty(t, sshKeys)

    // Check if our created SSH key is in the list
    found := false
    for _, key := range sshKeys {
        if key.ID == SSHKeyID {
            found = true
            break
        }
    }
    assert.True(t, found, "Created SSH key not found in the list")
}

func TestDeleteSSHKey(t *testing.T) {
    client := NewClient(BaseUrl, "Bearer "+*AccessToken)

    err := client.DeleteSSHKey(context.Background(), *SSHKeyID)
    assert.Nil(t, err)

    // Verify that the SSH key is deleted
    _, err = client.GetSSHKey(context.Background(), *SSHKeyID)
    assert.NotNil(t, err) // Expect an error as the SSH key should not exist
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
		Regions:        []string{"ap-northeast-1", "ap-northeast-3"},
		Nodes: []*models.ClusterNodeSettings{
			{
				Name:             "n1",
				Region:           stringPtr("ap-northeast-1"),
				Image:            "postgres",
				InstanceType:     "t4g.small",
				VolumeType:       "gp2",
			},
			{
				Name:             "n2",
				Region:           stringPtr("ap-northeast-3"),
				Image:            "postgres",
				InstanceType:     "t4g.medium",
				VolumeType:       "gp2",
			},
		},
		Networks: []*models.ClusterNetworkSettings{
			{
				Region:        stringPtr("ap-northeast-1"),
				Cidr:          "10.1.0.0/16",
				PublicSubnets: []string{"10.1.0.0/24"},
			},
			{
				Region:        stringPtr("ap-northeast-3"),
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
// 		Regions: []string{"ap-northeast-1", "ap-northeast-3"},
// 		Nodes: []*models.ClusterNodeSettings{
// 			{
// 				Name:             "n1",
// 				Region:           stringPtr("ap-northeast-1"),
// 				InstanceType:     "t4g.medium",
// 				AvailabilityZone: "ap-northeast-1a",
// 				VolumeType:       "gp2",
// 			},
// 			{
// 				Name:             "n2",
// 				Region:           stringPtr("ap-northeast-3"),
// 				InstanceType:     "t4g.medium",
// 				AvailabilityZone: "ap-northeast-3a",
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

	request := &models.CreateDatabaseInput{
		Name:      stringPtr("testdb"),
		ClusterID: *ClusterID,
	}

	database, err := client.CreateDatabase(context.Background(), request)
	assert.Nil(t, err)
	assert.NotNil(t, database)
	assert.Equal(t, "available", *database.Status)

	// Store the database ID for use in other tests
	DatabaseID = database.ID
}

func TestGetDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	database, err := client.GetDatabase(context.Background(), *DatabaseID)
	assert.Nil(t, err)
	assert.NotNil(t, database)
	assert.Equal(t, DatabaseID.String(), database.ID.String())
}

func TestGetDatabases(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	databases, err := client.GetDatabases(context.Background())
	assert.Nil(t, err)
	assert.NotEmpty(t, databases)

	// Check if our created database is in the list
	found := false
	for _, db := range databases {
		if db.ID == DatabaseID {
			found = true
			break
		}
	}
	assert.True(t, found, "Created database not found in the list")
}

// func TestUpdateDatabase(t *testing.T) {
// 	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

// 	updateRequest := &models.UpdateDatabaseInput{
// 		Options: []string{"new_option"},
// 	}

// 	updatedDatabase, err := client.UpdateDatabase(context.Background(), *DatabaseID, updateRequest)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, updatedDatabase)
// 	assert.Contains(t, updatedDatabase.Options, "new_option")
// }

func TestDeleteDatabase(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	err := client.DeleteDatabase(context.Background(), *DatabaseID)
	assert.Nil(t, err)

	// Verify that the database is deleted
	time.Sleep(5 * time.Second) // Give some time for deletion to propagate
	_, err = client.GetDatabase(context.Background(), *DatabaseID)
	assert.NotNil(t, err) // Expect an error as the database should not exist
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
