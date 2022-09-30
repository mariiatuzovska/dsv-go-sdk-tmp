/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ClientCredentials are used in the client_credentials authentication flow
type DsvClientCredentialsResponse struct {
	// TTL expiration in seconds
	TTL int64 `json:"TTL,omitempty"`
	// Url already used or not
	Accessed string `json:"accessed,omitempty"`
	// Unique uuid of client credentials
	ClientId string `json:"clientId,omitempty"`
	// Secret key returned on create
	ClientSecret string `json:"clientSecret,omitempty"`
	// Created date
	Created string `json:"created,omitempty"`
	// Who created
	CreatedBy string `json:"createdBy,omitempty"`
	Description string `json:"description,omitempty"`
	// ExpiredAt expiration time
	ExpiredAt string `json:"expiredAt,omitempty"`
	// the id for this item
	Id string `json:"id,omitempty"`
	// Assigned role for determining access
	Role string `json:"role,omitempty"`
	Status string `json:"status,omitempty"`
	// If Url requested
	Url bool `json:"url,omitempty"`
	// Url expiration time
	UrlExpires string `json:"urlExpires,omitempty"`
	// URL Path
	UrlPath string `json:"urlPath,omitempty"`
	// Url expiration in seconds
	UrlTTL int64 `json:"urlTTL,omitempty"`
	UsedCount int64 `json:"usedCount,omitempty"`
	// Uses  the number of times the client credential can be read.  if set to 0, it can be used infinitely.  default is 0.
	UsesLimit int64 `json:"usesLimit,omitempty"`
}
