package elasticsearch

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatalf("error when creating elasticsearch client: %v", err)
	}

	res, err := client.Info()
	if err != nil {
		t.Fatalf("error when connecting to Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	fmt.Printf("Response status from Elasticsearch: %s\n", res.Status())
}
