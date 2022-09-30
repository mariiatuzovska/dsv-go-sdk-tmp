/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DsvRootCaSecret struct {
	// Certificate of the root CA that contains information about it and public key
	Certificate string `json:"certificate"`
	// URL of the CRL from which the revocation of leaf certificates can be checked
	Crl string `json:"crl,omitempty"`
	// List of domains for which certificate signing is allowed
	Domains []string `json:"domains"`
	// Maximum TTL of a signed certificate issued from a given root CA (in hours)
	MaxTTL int64 `json:"maxTTL"`
	// Private key of the root CA
	PrivateKey string `json:"privateKey"`
	// RootCAPath to secret, which also serves as an identifier of the root CA
	RootCAPath string `json:"rootCAPath"`
}
