/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type DsvRootCaRegistration struct {
	CommonName string `json:"commonName"`
	Country string `json:"country,omitempty"`
	// URL of the CRL from which the revocation of leaf certificates can be checked
	Crl string `json:"crl,omitempty"`
	Description string `json:"description,omitempty"`
	// List of domains for which certificate signing is allowed
	Domains []string `json:"domains"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Locality string `json:"locality,omitempty"`
	// Maximum TTL of a signed certificate issued from a given root CA (in hours)
	MaxTTL int64 `json:"maxTTL"`
	Organization string `json:"organization,omitempty"`
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`
	// The name of the secret containing the root CA certificate
	RootCAPath string `json:"rootCAPath"`
	State string `json:"state,omitempty"`
	// The name of the secret in which to store the generated certificate and private key
	StorePath string `json:"storePath,omitempty"`
}
