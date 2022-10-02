/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// GroupResponse  response with metadata
type DsvGroupResponse struct {
	// Created date
	Created string `json:"created,omitempty"`
	// Who created
	CreatedBy string `json:"createdBy,omitempty"`
	// Name
	GroupName string `json:"groupName,omitempty"`
	// the id for this item
	Id string `json:"id,omitempty"`
	// Last updated date
	LastModified string `json:"lastModified,omitempty"`
	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	// Members
	Members []string `json:"members,omitempty"`
	// MetaData
	MetaData []map[string]string `json:"metaData,omitempty"`
	// Total number of members
	Total int64 `json:"total,omitempty"`
	Version string `json:"version,omitempty"`
}