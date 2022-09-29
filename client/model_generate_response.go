/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type GenerateResponse struct {
	MinNumberOfShares int64 `json:"minNumberOfShares,omitempty"`
	NewAdmins []string `json:"newAdmins,omitempty"`
	Status string `json:"status,omitempty"`
	TotalNumberOfShares int64 `json:"totalNumberOfShares,omitempty"`
}
