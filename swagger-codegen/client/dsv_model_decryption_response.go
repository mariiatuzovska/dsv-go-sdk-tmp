/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// DecryptionResponse contains data decrypted from ciphertext
type DsvDecryptionResponse struct {
	Data string `json:"data,omitempty"`
	// Path of the key with which decryption was performed
	Path string `json:"path,omitempty"`
	// Version of the key with which decryption was performed
	Version string `json:"version,omitempty"`
}
