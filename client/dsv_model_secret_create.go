/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DsvSecretCreate struct {
	// The user defined metadata
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The secret data
	Data map[string]interface{} `json:"data,omitempty"`
	// The secret's description
	Description string `json:"description,omitempty"`
}
