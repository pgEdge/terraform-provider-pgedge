package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	mock.Mock
}

func TestGetDatabases(t *testing.T) {
	client := NewClient("https://devapi.pgedge.com","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODM1NzUzLCJleHAiOjE3MDQ5MjIxNTMsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.DshrEfoBaXTTr-XzHpoIHK-FI0ZmWmUNx-9hycauPliagQeJQs8WYPf-nIlZE8s9ofTyvlhivVFhjr22oTD8sg-bWWK7yLoSCxc2BUlb4WFOu75pa8N3vkcUcbP7Zvq5GcdoS314tHf0WAwCSA81_Af33JKbudX05QWeYUnibzwfXfbd9emCQPDMIfy2xz2u8LwGFA06U1T8ZN-uKrABMozi2EeenPZtuduUwZ7CZd2mRHp5-xm4OmLq2xhTZELPnv3aF1_LZFzv9LcloVxLNHpKgOVfmmHOl5ZgMu46KjLI2Z0ALU25X2QTFyYYFY0Wfq1cAQW1DGMVUudp-Yc25g")
	databases, err := client.GetDatabases(context.Background())
	fmt.Println("databases: ", databases)

	assert.Nil(t, err)
	// assert.Equal(t, len(databases), 2)
	// assert.Equal(t, databases[0].ID, "1")
	// assert.Equal(t, databases[0].Name, "database1")
	// assert.Equal(t, databases[1].ID, "2")
	// assert.Equal(t, databases[1].Name, "database2")
}
