/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client
import (
	"time"
)

// TaskState represents a state of a task
type TaskState struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Results []TaskResult `json:"results,omitempty"`
	State string `json:"state,omitempty"`
	TaskName string `json:"taskName,omitempty"`
}
