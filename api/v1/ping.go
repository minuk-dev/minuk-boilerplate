// Package v1 provides the API version 1.
// It contains the request and response structures for the API endpoints.
package v1

// PingResponse represents the response structure for the ping endpoint.
type PingResponse struct {
	Message string `json:"message"`
}
