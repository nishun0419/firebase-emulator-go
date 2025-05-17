package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()
	createSample(ctx, client)
}

func createClient(ctx context.Context) *firestore.Client {
	projectID := "test-app-123456"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("Failed to create client: %v")
	}
	return client
}

func createSample(ctx context.Context, client *firestore.Client) {
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "test",
		"last":  "test",
		"born":  1995,
	})
	if err != nil {
		fmt.Println("Failed adding alovelace: %v")
	}
}
