/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// AccessTokenResponse contains the bearer access token for accessing authorized endpoints
type DsvAccessTokenResponse struct {
	// JWT access token for authorized requests
	AccessToken string `json:"accessToken,omitempty"`
	// Seconds until access token expires
	ExpiresIn int64 `json:"expiresIn,omitempty"`
	// Refresh token that can be used to get a new access token
	RefreshToken string `json:"refreshToken,omitempty"`
	// Token type
	TokenType string `json:"tokenType,omitempty"`
}
