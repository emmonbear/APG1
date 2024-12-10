// Copyright 2024 Moskalev Ilya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package elasticsearch

import (
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/go-elasticsearch/v8"
)

// NewClient creates and returns a new Elasticsearch client.
// The client is configured to handle retries for transient errors such as 502 (Bad Gateway),
// 503 (Service Unavailable), 504 (Gateway Timeout), and 429 (Too Many Requests), using exponential backoff.
//
// Retry settings:
//  - RetryOnStatus: A list of HTTP statuses for which the request will be retried.
//  - RetryBackoff: A function to determine the backoff duration between retry attempts using exponential backoff.
//  - MaxRetries: The maximum number of retry attempts (default is 5).
//
// An exponential backoff instance is created and reset at the first retry attempt. After that, the intervals between retries increase.
// 
// Returns:
//  - *elasticsearch.Client: The Elasticsearch client to interact with the server.
//  - error: An error if the client creation fails (e.g., invalid configuration).
//
// Example usage:
// 	client, err := NewClient()
// 	if err != nil {
// 		log.Fatalf("Error creating client: %v", err)
// 	}
// 	// Use the client to perform requests to Elasticsearch.
func NewClient() (*elasticsearch.Client, error) {
	retryBackoff := backoff.NewExponentialBackOff()
	return elasticsearch.NewClient(elasticsearch.Config{
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff: func(attempt int) time.Duration {
			if attempt == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		MaxRetries: 5,
	})
}
