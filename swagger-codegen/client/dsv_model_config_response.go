/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Document is the per-tenant configuration store
type DsvConfigResponse struct {
	Created string `json:"created,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	PermissionDocument []DsvDefaultPolicy `json:"permissionDocument,omitempty"`
	RefreshTokenTTLHours int64 `json:"refreshTokenTTLHours,omitempty"`
	Settings *DsvSettings `json:"settings,omitempty"`
	TenantName string `json:"tenantName,omitempty"`
	Version string `json:"version,omitempty"`
}
