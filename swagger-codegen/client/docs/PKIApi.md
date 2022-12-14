# \PKIApi

All URIs are relative to *https://secretsvaultcloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeafParams**](PKIApi.md#LeafParams) | **Post** /pki/leaf | Create Leaf Certificate and Private Key
[**RegisterRoot**](PKIApi.md#RegisterRoot) | **Post** /pki/register | Register Root CA
[**RootCARegistration**](PKIApi.md#RootCARegistration) | **Post** /pki/root | Generate Root Certificate
[**SSHCertParams**](PKIApi.md#SSHCertParams) | **Post** /pki/ssh-cert | Create SSH Certificate
[**SignCertificate**](PKIApi.md#SignCertificate) | **Post** /pki/sign | Create Signed Certificate


# **LeafParams**
> DsvResponseCertificate LeafParams(ctx, body)
Create Leaf Certificate and Private Key

Create and return a signed certificate with a private key based on a registered root CA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvSigningRequestInformation**](DsvSigningRequestInformation.md)|  | 

### Return type

[**DsvResponseCertificate**](ResponseCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterRoot**
> DsvResponseRootCertificate RegisterRoot(ctx, body)
Register Root CA

Register a root CA as a secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvRootCaSecret**](DsvRootCaSecret.md)|  | 

### Return type

[**DsvResponseRootCertificate**](ResponseRootCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RootCARegistration**
> DsvResponseCertificate RootCARegistration(ctx, body)
Generate Root Certificate

Create and return a new root certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvRootCaRegistration**](DsvRootCaRegistration.md)|  | 

### Return type

[**DsvResponseCertificate**](ResponseCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SSHCertParams**
> DsvResponseSshCertificate SSHCertParams(ctx, body)
Create SSH Certificate

Create, store and return a signed SSH certificate using a root private key and SHH-compatible leaf public key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvSshCertInformation**](DsvSshCertInformation.md)|  | 

### Return type

[**DsvResponseSshCertificate**](ResponseSSHCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignCertificate**
> DsvSignedLeafCertificate SignCertificate(ctx, body)
Create Signed Certificate

Create and return a signed certificate based on a registered root CA with a given CSR.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DsvSigningRequest**](DsvSigningRequest.md)|  | 

### Return type

[**DsvSignedLeafCertificate**](SignedLeafCertificate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

