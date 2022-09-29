/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AuthenticationDetailsModel struct {
	Properties *AuthenticationProviderPropertiesModel `json:"properties,omitempty"`
	// AuthenticationSettings provider system type.
	Type_ string `json:"type,omitempty"`
}
