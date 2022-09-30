/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// User search response with metadata
type DsvUserSearch struct {
	// Cursor to next batch of results
	Cursor string `json:"cursor,omitempty"`
	// Users that match the search term
	Data []DsvUserResponse `json:"data,omitempty"`
	// The number of results in this response
	Length int64 `json:"length,omitempty"`
	// The maximum number of results per cursor
	Limit int64 `json:"limit,omitempty"`
	// Total number of items
	Total int64 `json:"total,omitempty"`
}
