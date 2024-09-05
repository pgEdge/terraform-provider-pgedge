package client

import (
	"context"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/pgEdge/terraform-provider-pgedge/client/models"
	"github.com/stretchr/testify/assert"
)

const (
	BaseUrl        = "" //your base url here
	ClientID       = "" //your client id here
	ClientSecret   = "" //your client secret here
)

var (
	AccessToken *string
	DatabaseID  *strfmt.UUID
	ClusterID   *strfmt.UUID
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

func TestCreateCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	request := &models.ClusterCreationRequest{
		Name:           "n3",
		CloudAccountID: CloudAccountID.String(),
		Regions:        []string{"us-east-2"},
		Nodes: []*models.ClusterNode{
			{
				Name:             "n1",
				Region:           "us-east-2",
				Image:            "postgres",
				InstanceType:     "t4g.small",
				AvailabilityZone: "us-east-2a",
				VolumeType:       "gp2",
			},
		},
		Networks: []*models.Network{
			{
				Region:        "us-east-2",
				Cidr:          "10.1.0.0/16",
				PublicSubnets: []string{"10.1.0.0/24"},
			},
		},
		FirewallRules: []*models.FirewallRule{
			{
				Name:    "postgres",
				Port:    5432,
				Sources: []string{"0.0.0.0/0"},
			},
		},
		ResourceTags: map[string]string{
			"key": "value",
		},
	}

	cluster, err := client.CreateCluster(context.Background(), request)
	ClusterID = &cluster.ID

	assert.Nil(t, err)
}

func TestGetCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetCluster(context.Background(), strfmt.UUID(*ClusterID))

	assert.Nil(t, err)
}

func TestUpdateCluster(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	request := &models.ClusterUpdateRequest{
		Regions: []string{"us-east-2"},
		Nodes: []*models.ClusterNode{
			{
				Name:             "n1",
				Region:           "us-east-2",
				Image:            "postgres",
				InstanceType:     "t4g.small",
				AvailabilityZone: "us-east-2a",
				VolumeType:       "gp2",
			},
			{
				Name:             "n2",
				Region:           "us-east-1",
				Image:            "postgres",
				InstanceType:     "t4g.small",
				AvailabilityZone: "us-east-2a",
				VolumeType:       "gp2",
			},
		},
		Networks: []*models.Network{
			{
				Region:        "us-east-2",
				Cidr:          "10.1.0.0/16",
				PublicSubnets: []string{"10.1.0.0/24"},
			},
		},
		FirewallRules: []*models.FirewallRule{
			{
				Name:    "postgres",
				Port:    5432,
				Sources: []string{"0.0.0.0/0"},
			},
		},
	}

	cluster, err := client.UpdateCluster(context.Background(), *ClusterID, request)
	ClusterID = &cluster.ID

	assert.Nil(t, err)
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

	err := client.DeleteCluster(context.Background(), strfmt.UUID(*ClusterID))

	assert.Nil(t, err)
}


func TestGetCloudAccounts(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	accounts, err := client.GetCloudAccounts(context.Background())

	assert.Nil(t, err)
	assert.NotEmpty(t, accounts)
}

func TestCreateCloudAccount(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)

	accountType := "aws"
	request := &models.CreateCloudAccountInput{
		Name: "TestAccount",
		Type: &accountType,
		Credentials: map[string]interface{}{
			"role_arn": "",
		},
	}



	account, err := client.CreateCloudAccount(context.Background(), request)
	CloudAccountID = account.ID

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "TestAccount", *account.Name)
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

func TestGetDeletedCloudAccount(t *testing.T) {
	client := NewClient(BaseUrl, "Bearer "+*AccessToken)
	_, err := client.GetCloudAccount(context.Background(), *CloudAccountID)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "404")
}
