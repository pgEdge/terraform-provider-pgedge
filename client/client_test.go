package client

import (
	"context"
	"fmt"
	"testing"

	// "github.com/pgEdge/terraform-provider-pgedge/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	mock.Mock
}

// func TestGetDatabases(t *testing.T) {
// 	client := NewClient("https://devapi.pgedge.com","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODg4Njc3LCJleHAiOjE3MDQ5NzUwNzcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.aJcnvr7OuAhafGN0sf784cGFYdg16confNXN0tOSBAQrorD-ywFYxSTU0nI4lZwJudJfPsBEeIjbt-Rv6IxG3XNs6MpZCd-uMQEmAirkKHcMd978vhP5nYlZhIAD7aRm5VxTalrlJphTi3hUs7yUhGUINzE9N3SgFe32x5ZdhMaRnDXvIX64vTOJfZbyMCcBQq_-QZG7Iqsfn4KISGPCia5Q-9CaSfViInMgJm0WZb_AT-NHNqt8ic7QYjP8EVlsCpC-eN1HVepL1AynxkPz7J-cSq21ycZTRH8AJ0EHca8HJ3ke5Skq8xovJxbMOZ1ofwH4jvLxSis8cKgdLYH8Yg")
// 	databases, err := client.GetDatabases(context.Background())
// 	fmt.Println("databases: ", databases)

// 	assert.Nil(t, err)
// }


// func TestGetDatabase(t *testing.T){
// 	client := NewClient("https://devapi.pgedge.com","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODg4Njc3LCJleHAiOjE3MDQ5NzUwNzcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.aJcnvr7OuAhafGN0sf784cGFYdg16confNXN0tOSBAQrorD-ywFYxSTU0nI4lZwJudJfPsBEeIjbt-Rv6IxG3XNs6MpZCd-uMQEmAirkKHcMd978vhP5nYlZhIAD7aRm5VxTalrlJphTi3hUs7yUhGUINzE9N3SgFe32x5ZdhMaRnDXvIX64vTOJfZbyMCcBQq_-QZG7Iqsfn4KISGPCia5Q-9CaSfViInMgJm0WZb_AT-NHNqt8ic7QYjP8EVlsCpC-eN1HVepL1AynxkPz7J-cSq21ycZTRH8AJ0EHca8HJ3ke5Skq8xovJxbMOZ1ofwH4jvLxSis8cKgdLYH8Yg")
// 	database, err := client.GetDatabase(context.Background(),"f556e712-696f-41d7-9191-d5c337c348ea")
// 	fmt.Println("database: ", database)

// 	assert.Nil(t, err)
// }

// func TestCreateDatabase(t *testing.T){
// 	client := NewClient("https://devapi.pgedge.com","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODg4Njc3LCJleHAiOjE3MDQ5NzUwNzcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.aJcnvr7OuAhafGN0sf784cGFYdg16confNXN0tOSBAQrorD-ywFYxSTU0nI4lZwJudJfPsBEeIjbt-Rv6IxG3XNs6MpZCd-uMQEmAirkKHcMd978vhP5nYlZhIAD7aRm5VxTalrlJphTi3hUs7yUhGUINzE9N3SgFe32x5ZdhMaRnDXvIX64vTOJfZbyMCcBQq_-QZG7Iqsfn4KISGPCia5Q-9CaSfViInMgJm0WZb_AT-NHNqt8ic7QYjP8EVlsCpC-eN1HVepL1AynxkPz7J-cSq21ycZTRH8AJ0EHca8HJ3ke5Skq8xovJxbMOZ1ofwH4jvLxSis8cKgdLYH8Yg")
	
// 	request := &models.DatabaseCreationRequest{
// 		Name: "test",
// 		ClusterID: "5e7478e5-4e68-464b-902d-747db528eccc",
// 	}
	
// 	database, err := client.CreateDatabase(context.Background(),request)
// 	fmt.Println("database: ", database)

// 	assert.Nil(t, err)
// }

// func TestDeleteDatabase(t *testing.T){
// 	client := NewClient("https://devapi.pgedge.com","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODg4Njc3LCJleHAiOjE3MDQ5NzUwNzcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.aJcnvr7OuAhafGN0sf784cGFYdg16confNXN0tOSBAQrorD-ywFYxSTU0nI4lZwJudJfPsBEeIjbt-Rv6IxG3XNs6MpZCd-uMQEmAirkKHcMd978vhP5nYlZhIAD7aRm5VxTalrlJphTi3hUs7yUhGUINzE9N3SgFe32x5ZdhMaRnDXvIX64vTOJfZbyMCcBQq_-QZG7Iqsfn4KISGPCia5Q-9CaSfViInMgJm0WZb_AT-NHNqt8ic7QYjP8EVlsCpC-eN1HVepL1AynxkPz7J-cSq21ycZTRH8AJ0EHca8HJ3ke5Skq8xovJxbMOZ1ofwH4jvLxSis8cKgdLYH8Yg")

// 	err := client.DeleteDatabase(context.Background(),"3055443b-c930-41e4-8e1e-d252f2ad992d")
// 	fmt.Println("err: ", err)

// 	assert.Contains(t, err.Error(), "200")
// }

// func TestReplicateDatabase(t *testing.T){
// 	client := NewClient("https://devapi.pgedge.com","Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Il9zQ19yVTFoTFRxTUgwajh4a04zMiJ9.eyJwZ2VkZ2VfdGVuYW50X2lkIjoiYmI3NWI2NzMtMjkzMS00N2E3LTk1NjAtYmZiMGVmOGNlOWViIiwicGdlZGdlX2NsaWVudF9pZCI6ImEzYTc0ZmI3LTY4YjItNDY0ZC1hMjFjLWM2MDYxZGVjYjllNCIsImlzcyI6Imh0dHBzOi8vZGV2YXV0aC5wZ2VkZ2UuY29tLyIsInN1YiI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2RldmFwaS5wZ2VkZ2UuY29tIiwiaWF0IjoxNzA0ODg4Njc3LCJleHAiOjE3MDQ5NzUwNzcsImF6cCI6ImkycG5kN0wxNTIzNFN0eThGSXdUQ0ZUT2JxRDFSWFpsIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.aJcnvr7OuAhafGN0sf784cGFYdg16confNXN0tOSBAQrorD-ywFYxSTU0nI4lZwJudJfPsBEeIjbt-Rv6IxG3XNs6MpZCd-uMQEmAirkKHcMd978vhP5nYlZhIAD7aRm5VxTalrlJphTi3hUs7yUhGUINzE9N3SgFe32x5ZdhMaRnDXvIX64vTOJfZbyMCcBQq_-QZG7Iqsfn4KISGPCia5Q-9CaSfViInMgJm0WZb_AT-NHNqt8ic7QYjP8EVlsCpC-eN1HVepL1AynxkPz7J-cSq21ycZTRH8AJ0EHca8HJ3ke5Skq8xovJxbMOZ1ofwH4jvLxSis8cKgdLYH8Yg")

// 	database, err := client.ReplicateDatabase(context.Background(),"9eecf93c-dd34-444d-8a0e-c8f82f4cdffb")
// 	fmt.Println("database: ", database)
// 	fmt.Println("err: ", err)

// 	assert.Nil(t, err)
// }


func TestOAuthToken(t *testing.T){
	client := NewClient("https://devapi.pgedge.com","")

	token, err := client.OAuthToken(context.Background())
	fmt.Println("token: ", token)
	fmt.Println("err: ", err)

	assert.Nil(t, err)
}