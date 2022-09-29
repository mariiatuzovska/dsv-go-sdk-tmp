/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

// ResponseModelDescription contains metadata but not the sensitive secret data
type SecretDescription struct {
	// The user defined metadata
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Created date
	Created string `json:"created,omitempty"`
	// Who created
	CreatedBy string `json:"createdBy,omitempty"`
	// Description of secret
	Description string `json:"description,omitempty"`
	// The unique id for this item
	Id string `json:"id,omitempty"`
	// Last updated date
	LastModified string `json:"lastModified,omitempty"`
	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// The path the secret is located at
	Path string `json:"path,omitempty"`
	Version string `json:"version,omitempty"`
}
