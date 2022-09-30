/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type DsvManualKeyData struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Base64 encoded nonce to be used with key. If not provided, DSV generates it for the user.
	Nonce string `json:"nonce,omitempty"`
	// Base64 encoded private key
	PrivateKey string `json:"privateKey"`
	// Encryption scheme to be used.
	Scheme string `json:"scheme"`
}
