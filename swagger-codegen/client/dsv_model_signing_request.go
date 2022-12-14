/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type DsvSigningRequest struct {
	// Boolean indicating whether to return a root certificate
	Chain bool `json:"chain,omitempty"`
	// Certificate Signing Request
	Csr string `json:"csr"`
	// Path to secret - registered root CA
	RootCAPath string `json:"rootCAPath"`
	// A list of Subject Alternative Names for a certificate (each domain must be present in the list of allowed domains)
	SubjectAltNames []string `json:"subjectAltNames,omitempty"`
	// TTL for a generated certificate (in hours, cannot exceed the maximum TTL specified in root CA secret)
	Ttl int64 `json:"ttl,omitempty"`
}
