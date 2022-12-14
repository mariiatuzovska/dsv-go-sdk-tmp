/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type DsvEngineCreate struct {
	// Name of engine
	Name string `json:"name"`
	// Name of pool to assign engine to
	PoolName string `json:"poolName"`
}
