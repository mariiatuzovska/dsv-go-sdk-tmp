/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DsvSiemCreateUpdateRequestModel struct {
	// Authentication token
	Auth string `json:"auth"`
	// Authentication method (token)
	AuthMethod string `json:"authMethod"`
	// Endpoint
	Endpoint string `json:"endpoint,omitempty"`
	// Collect Server IP/FQDN
	Host string `json:"host"`
	// Logging Format (i.e. syslog (RFC 5424))
	LoggingFormat string `json:"loggingFormat"`
	// Name of registered SIEM endpoint, similar to path
	Name string `json:"name"`
	// Engine pool name, used when sending request to a DSV engine instance
	Pool string `json:"pool,omitempty"`
	// Port
	Port int64 `json:"port"`
	// Type of protocol (i.e. TCP, UDP)
	Protocol string `json:"protocol"`
	// Denotes whether the endpoint should be accessed through a DSV engine instance
	SendToEngine bool `json:"sendToEngine,omitempty"`
	// Type of endpoint (\"syslog\", \"cef\", \"json\", \"splunk\")
	SiemType string `json:"siemType"`
}