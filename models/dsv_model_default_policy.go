/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DsvDefaultPolicy struct {
	Actions []string `json:"actions,omitempty"`
	Conditions *map[string]DsvCondition `json:"conditions,omitempty"`
	Description string `json:"description,omitempty"`
	Effect string `json:"effect,omitempty"`
	Id string `json:"id,omitempty"`
	Meta []int32 `json:"meta,omitempty"`
	Resources []string `json:"resources,omitempty"`
	Subjects []string `json:"subjects,omitempty"`
}