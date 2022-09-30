/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Role defines the role security principal and any mappings to 3rd party providers
type DsvRoleResponse struct {
	// Created date
	Created string `json:"created,omitempty"`
	// Who created
	CreatedBy string `json:"createdBy,omitempty"`
	// Role description
	Description string `json:"description,omitempty"`
	// External identifier, such as an AWS arn for 3rd party authentication
	ExternalId string `json:"externalId,omitempty"`
	Groups []string `json:"groups,omitempty"`
	// the id for this item
	Id string `json:"id,omitempty"`
	// Last updated date
	LastModified string `json:"lastModified,omitempty"`
	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// Name of role
	Name string `json:"name"`
	// Provider name defined in the authentication settings section of configuration
	Provider string `json:"provider,omitempty"`
	Version string `json:"version,omitempty"`
}
