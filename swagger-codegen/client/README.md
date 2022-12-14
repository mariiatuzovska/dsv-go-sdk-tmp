# Go API client for client

The purpose of this application is to provide a simple service for storing and getting secrets

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.1
- Build date: 2022-10-03T09:25:20.141Z
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *https://secretsvaultcloud.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditApi* | [**DownloadAudit**](docs/AuditApi.md#downloadaudit) | **Get** /download/audit | Download Audit Records
*AuditApi* | [**FindAudit**](docs/AuditApi.md#findaudit) | **Get** /audit | Find Audit Records
*BreakGlassApi* | [**ApplyRequest**](docs/BreakGlassApi.md#applyrequest) | **Post** /breakglass/apply | Apply
*BreakGlassApi* | [**GenerateRequest**](docs/BreakGlassApi.md#generaterequest) | **Post** /breakglass/generate | Generate
*BreakGlassApi* | [**GetStatus**](docs/BreakGlassApi.md#getstatus) | **Get** /breakglass | Get Status
*ClientsApi* | [**CreateClientCredential**](docs/ClientsApi.md#createclientcredential) | **Post** /clients | Create a Client Credential
*ClientsApi* | [**DeleteClientCredential**](docs/ClientsApi.md#deleteclientcredential) | **Delete** /clients/{clientId} | Delete a Client Credential
*ClientsApi* | [**GetBootstrapClientCredential**](docs/ClientsApi.md#getbootstrapclientcredential) | **Get** /clients/bootstrap/{clientId} | Get a Client Bootstrap Credential including secret
*ClientsApi* | [**GetClientCredential**](docs/ClientsApi.md#getclientcredential) | **Get** /clients/{clientId} | Get a Client Credential
*ClientsApi* | [**RestoreClient**](docs/ClientsApi.md#restoreclient) | **Get** /clients/{clientId}/restore | Restore a Client
*ClientsApi* | [**SearchClients**](docs/ClientsApi.md#searchclients) | **Get** /clients | Search for Client Credentials
*ConfigApi* | [**GetConfig**](docs/ConfigApi.md#getconfig) | **Get** /config | Get Config
*ConfigApi* | [**GetConfigByVersion**](docs/ConfigApi.md#getconfigbyversion) | **Get** /config/version/{version} | Get Config By Version
*ConfigApi* | [**PostConfig**](docs/ConfigApi.md#postconfig) | **Post** /config | Create Config
*EaaSAutoApi* | [**CreateKey**](docs/EaaSAutoApi.md#createkey) | **Post** /crypto/key/{path} | Create Key
*EaaSAutoApi* | [**Decrypt**](docs/EaaSAutoApi.md#decrypt) | **Post** /crypto/decrypt | Decrypt
*EaaSAutoApi* | [**DeleteKey**](docs/EaaSAutoApi.md#deletekey) | **Delete** /crypto/key/{path} | Delete Key
*EaaSAutoApi* | [**Encrypt**](docs/EaaSAutoApi.md#encrypt) | **Post** /crypto/encrypt | Encrypt
*EaaSAutoApi* | [**GetKeyMetadata**](docs/EaaSAutoApi.md#getkeymetadata) | **Get** /crypto/key/{path} | Get Key Metadata
*EaaSAutoApi* | [**RestoreKey**](docs/EaaSAutoApi.md#restorekey) | **Put** /crypto/key/{path}/restore | Restore Key
*EaaSAutoApi* | [**RotationRequest**](docs/EaaSAutoApi.md#rotationrequest) | **Post** /crypto/rotate | Rotate Data and Key
*EaaSManualApi* | [**DecryptWithManualKey**](docs/EaaSManualApi.md#decryptwithmanualkey) | **Post** /crypto/manual/decrypt | Decrypt
*EaaSManualApi* | [**DeleteManualKey**](docs/EaaSManualApi.md#deletemanualkey) | **Delete** /crypto/manual/key/{path} | Delete Key
*EaaSManualApi* | [**EncryptWithManualKey**](docs/EaaSManualApi.md#encryptwithmanualkey) | **Post** /crypto/manual/encrypt | Encrypt
*EaaSManualApi* | [**ReadManualKey**](docs/EaaSManualApi.md#readmanualkey) | **Get** /crypto/manual/key/{path} | Read Key
*EaaSManualApi* | [**RestoreManualKey**](docs/EaaSManualApi.md#restoremanualkey) | **Put** /crypto/manual/key/{path}/restore | Restore Key
*EaaSManualApi* | [**UpdateKey**](docs/EaaSManualApi.md#updatekey) | **Put** /crypto/manual/key/{path} | Update Key
*EaaSManualApi* | [**UploadKey**](docs/EaaSManualApi.md#uploadkey) | **Post** /crypto/manual/key/{path} | Upload Key
*EnginesApi* | [**CreateEngine**](docs/EnginesApi.md#createengine) | **Post** /engines | Create Engine
*EnginesApi* | [**DeleteEngine**](docs/EnginesApi.md#deleteengine) | **Delete** /engines/{name} | Delete Engine
*EnginesApi* | [**GetEngine**](docs/EnginesApi.md#getengine) | **Get** /engines/{name} | Get Engine
*EnginesApi* | [**ListEngines**](docs/EnginesApi.md#listengines) | **Get** /engines | List Engines
*EnginesApi* | [**PingEngine**](docs/EnginesApi.md#pingengine) | **Post** /engines/{name}/ping | Ping Engine
*GroupsApi* | [**AddMember**](docs/GroupsApi.md#addmember) | **Post** /groups/{name}/members | Add Members To Group
*GroupsApi* | [**CreateGroup**](docs/GroupsApi.md#creategroup) | **Post** /groups/ | Create Group
*GroupsApi* | [**DeleteGroup**](docs/GroupsApi.md#deletegroup) | **Delete** /groups/{name} | Delete Group
*GroupsApi* | [**DeleteMember**](docs/GroupsApi.md#deletemember) | **Delete** /groups/{name}/members | Delete Members From Group
*GroupsApi* | [**GetGroup**](docs/GroupsApi.md#getgroup) | **Get** /groups/{name} | Get Group
*GroupsApi* | [**RestoreGroup**](docs/GroupsApi.md#restoregroup) | **Get** /groups/{name}/restore | Restore Group
*GroupsApi* | [**SearchGroups**](docs/GroupsApi.md#searchgroups) | **Get** /groups | Search Groups
*HomeApi* | [**CreateHomeSecret**](docs/HomeApi.md#createhomesecret) | **Post** /home/{principalName}/{path} | Create home secrets
*HomeApi* | [**DeleteHomeSecret**](docs/HomeApi.md#deletehomesecret) | **Delete** /home/{principalName}/{path} | Delete Home Secret
*HomeApi* | [**GetHomeSecret**](docs/HomeApi.md#gethomesecret) | **Get** /home/{principalName}/{path} | Get Home
*HomeApi* | [**GetHomeSecretByVersion**](docs/HomeApi.md#gethomesecretbyversion) | **Get** /home/{principalName}/{path}/version/{version} | Get Home Secret By Version
*HomeApi* | [**GetHomeSecretDescription**](docs/HomeApi.md#gethomesecretdescription) | **Get** /home/{principalName}/{path}::description | Get home Secret Description
*HomeApi* | [**RestoreHomeSecret**](docs/HomeApi.md#restorehomesecret) | **Get** /home/{principalName}/{path}/restore | Restore home Secret
*HomeApi* | [**RollbackHomeSecret**](docs/HomeApi.md#rollbackhomesecret) | **Put** /home/{principalName}/{path}/rollback/{version} | Rollback a Home Secret
*HomeApi* | [**SearchHomeSecrets**](docs/HomeApi.md#searchhomesecrets) | **Get** /home/{principalName} | Search for Home Secrets
*HomeApi* | [**UpdateHomeSecret**](docs/HomeApi.md#updatehomesecret) | **Put** /home/{principalName}/{path} | Update home Secret
*KeyApi* | [**Masterkeys**](docs/KeyApi.md#masterkeys) | **Put** /config/keys | Update Master Key
*PKIApi* | [**LeafParams**](docs/PKIApi.md#leafparams) | **Post** /pki/leaf | Create Leaf Certificate and Private Key
*PKIApi* | [**RegisterRoot**](docs/PKIApi.md#registerroot) | **Post** /pki/register | Register Root CA
*PKIApi* | [**RootCARegistration**](docs/PKIApi.md#rootcaregistration) | **Post** /pki/root | Generate Root Certificate
*PKIApi* | [**SSHCertParams**](docs/PKIApi.md#sshcertparams) | **Post** /pki/ssh-cert | Create SSH Certificate
*PKIApi* | [**SignCertificate**](docs/PKIApi.md#signcertificate) | **Post** /pki/sign | Create Signed Certificate
*PoliciesApi* | [**CreatePolicy**](docs/PoliciesApi.md#createpolicy) | **Post** /config/policies/ | Create Policy
*PoliciesApi* | [**DeletePolicy**](docs/PoliciesApi.md#deletepolicy) | **Delete** /config/policies/{path} | Delete Policy
*PoliciesApi* | [**GetPolicy**](docs/PoliciesApi.md#getpolicy) | **Get** /config/policies/{path} | Get Policy
*PoliciesApi* | [**GetPolicyByVersion**](docs/PoliciesApi.md#getpolicybyversion) | **Get** /config/policies/{path}/version/{version} | Get a list of policies by version
*PoliciesApi* | [**RestorePolicy**](docs/PoliciesApi.md#restorepolicy) | **Get** /config/policies/{path}/restore | Restore Policy
*PoliciesApi* | [**RollbackPolicy**](docs/PoliciesApi.md#rollbackpolicy) | **Put** /config/policies/{path}/rollback/{version} | Rollback Policy
*PoliciesApi* | [**SearchFilter**](docs/PoliciesApi.md#searchfilter) | **Get** /config/policies | Search Policies
*PoliciesApi* | [**UpdatePolicy**](docs/PoliciesApi.md#updatepolicy) | **Put** /config/policies/{path} | Update Policy
*PoolsApi* | [**CreatePool**](docs/PoolsApi.md#createpool) | **Post** /pools | Create Pool
*PoolsApi* | [**DeletePool**](docs/PoolsApi.md#deletepool) | **Delete** /pools/{name} | Delete Pool
*PoolsApi* | [**GetPool**](docs/PoolsApi.md#getpool) | **Get** /pools/{name} | Get Pool
*PoolsApi* | [**ListPools**](docs/PoolsApi.md#listpools) | **Get** /pools | List Pools
*RolesApi* | [**CreateRole**](docs/RolesApi.md#createrole) | **Post** /roles/ | Create a Role
*RolesApi* | [**DeleteRole**](docs/RolesApi.md#deleterole) | **Delete** /roles/{name} | Delete a Role
*RolesApi* | [**GetRole**](docs/RolesApi.md#getrole) | **Get** /roles/{name} | Get a Role
*RolesApi* | [**GetRoleByVersion**](docs/RolesApi.md#getrolebyversion) | **Get** /roles/{name}/version/{version} | Get a Role By Version
*RolesApi* | [**RestoreRole**](docs/RolesApi.md#restorerole) | **Get** /roles/{name}/restore | Restore a Role
*RolesApi* | [**SearchRoles**](docs/RolesApi.md#searchroles) | **Get** /roles | Search for Roles
*RolesApi* | [**UpdateRole**](docs/RolesApi.md#updaterole) | **Put** /roles/{name} | Update a Role
*SIEMApi* | [**SearchSiems**](docs/SIEMApi.md#searchsiems) | **Get** /config/siem | Search SIEM Endpoints
*SIEMApi* | [**SiemCreate**](docs/SIEMApi.md#siemcreate) | **Post** /config/siem | Create SIEM Endpoint
*SIEMApi* | [**SiemDelete**](docs/SIEMApi.md#siemdelete) | **Delete** /config/siem/{name} | Delete SIEM Endpoint
*SIEMApi* | [**SiemGet**](docs/SIEMApi.md#siemget) | **Get** /config/siem/{name} | Get SIEM Endpoint
*SIEMApi* | [**SiemUpdate**](docs/SIEMApi.md#siemupdate) | **Put** /config/siem/{name} | Update SIEM Endpoint
*SecretsApi* | [**CreateSecret**](docs/SecretsApi.md#createsecret) | **Post** /secrets/{path} | Create Secret
*SecretsApi* | [**DeleteSecret**](docs/SecretsApi.md#deletesecret) | **Delete** /secrets/{path} | Delete Secret
*SecretsApi* | [**GetSecret**](docs/SecretsApi.md#getsecret) | **Get** /secrets/{path} | Get Secret
*SecretsApi* | [**GetSecretByVersion**](docs/SecretsApi.md#getsecretbyversion) | **Get** /secrets/{path}/version/{version} | Get Secret By Version
*SecretsApi* | [**GetSecretDescription**](docs/SecretsApi.md#getsecretdescription) | **Get** /secrets/{path}::description | Get Secret Description
*SecretsApi* | [**ListSecretPaths**](docs/SecretsApi.md#listsecretpaths) | **Get** /secrets/{path}::listpaths | List Secret Paths
*SecretsApi* | [**RollbackSecret**](docs/SecretsApi.md#rollbacksecret) | **Put** /secrets/{path}/rollback/{version} | Rollback a Secret
*SecretsApi* | [**SearchSecrets**](docs/SecretsApi.md#searchsecrets) | **Get** /secrets | Search for Secrets
*SecretsApi* | [**UpdateSecret**](docs/SecretsApi.md#updatesecret) | **Put** /secrets/{path} | Update Secret
*SettingsApi* | [**CreateAuthSettings**](docs/SettingsApi.md#createauthsettings) | **Post** /config/auth/ | Create Authentication Provider
*SettingsApi* | [**DeleteAuthSettings**](docs/SettingsApi.md#deleteauthsettings) | **Delete** /config/auth/{name} | Delete Authentication Provider
*SettingsApi* | [**GetAuthSettings**](docs/SettingsApi.md#getauthsettings) | **Get** /config/auth/{name} | Get Authentication Provider
*SettingsApi* | [**GetAuthSettingsByVersion**](docs/SettingsApi.md#getauthsettingsbyversion) | **Get** /config/auth/{name}/version/{version} | Get a list of Authentication Settings by version
*SettingsApi* | [**RestoreAuthSettings**](docs/SettingsApi.md#restoreauthsettings) | **Get** /config/auth/{name}/restore | Restore Authentication Provider
*SettingsApi* | [**RollbackAuthSettings**](docs/SettingsApi.md#rollbackauthsettings) | **Put** /config/auth/{name}/rollback/{version} | Rollback Authentication Provider
*SettingsApi* | [**SearchSettings**](docs/SettingsApi.md#searchsettings) | **Get** /config/auth | Search Authentication Providers
*SettingsApi* | [**UpdateAuthSettings**](docs/SettingsApi.md#updateauthsettings) | **Put** /config/auth/{name} | Update Authentication Provider
*TasksApi* | [**GetTaskStatus**](docs/TasksApi.md#gettaskstatus) | **Get** /task/status/{id} | Get background task status
*TokensApi* | [**InitCertAuth**](docs/TokensApi.md#initcertauth) | **Post** /certificate/auth | Initiate authentication by certificate
*TokensApi* | [**Revoke**](docs/TokensApi.md#revoke) | **Post** /revoke/{refreshtoken} | Revoke Refresh Token
*TokensApi* | [**Token**](docs/TokensApi.md#token) | **Post** /token | Authenticate
*UsageApi* | [**GetUsage**](docs/UsageApi.md#getusage) | **Get** /usage | Get Usage
*UsersApi* | [**AddToGroups**](docs/UsersApi.md#addtogroups) | **Post** /users/{name}/groups | Add Member To Groups
*UsersApi* | [**ChangePassword**](docs/UsersApi.md#changepassword) | **Post** /users/{name}/password | Change Password
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /users/ | Create a User
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /users/{name} | Delete a User
*UsersApi* | [**GetMember**](docs/UsersApi.md#getmember) | **Get** /users/{name}/groups | Get Member Group
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /users/{name} | Get a User
*UsersApi* | [**GetUserByVersion**](docs/UsersApi.md#getuserbyversion) | **Get** /users/{name}/version/{version} | Get a User By Version
*UsersApi* | [**RestoreUser**](docs/UsersApi.md#restoreuser) | **Get** /users/{name}/restore | Restore a User
*UsersApi* | [**SearchUsers**](docs/UsersApi.md#searchusers) | **Get** /users | Search for Users
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /users/{name} | Update a User


## Documentation For Models

 - [DsvAccessTokenResponse](docs/DsvAccessTokenResponse.md)
 - [DsvAddMemberResponse](docs/DsvAddMemberResponse.md)
 - [DsvAddToGroupsRequest](docs/DsvAddToGroupsRequest.md)
 - [DsvAddToGroupsResponse](docs/DsvAddToGroupsResponse.md)
 - [DsvApplyResponse](docs/DsvApplyResponse.md)
 - [DsvAudit](docs/DsvAudit.md)
 - [DsvAuthProperties](docs/DsvAuthProperties.md)
 - [DsvAuthentication](docs/DsvAuthentication.md)
 - [DsvAuthenticationDetailsModel](docs/DsvAuthenticationDetailsModel.md)
 - [DsvAuthenticationProviderPropertiesModel](docs/DsvAuthenticationProviderPropertiesModel.md)
 - [DsvAuthenticationSettingsCreateModel](docs/DsvAuthenticationSettingsCreateModel.md)
 - [DsvAuthenticationSettingsResponse](docs/DsvAuthenticationSettingsResponse.md)
 - [DsvAuthenticationSettingsSearchResponse](docs/DsvAuthenticationSettingsSearchResponse.md)
 - [DsvAuthenticationSettingsVersionResponse](docs/DsvAuthenticationSettingsVersionResponse.md)
 - [DsvAutoKey](docs/DsvAutoKey.md)
 - [DsvAutoKeyResponse](docs/DsvAutoKeyResponse.md)
 - [DsvClientCreate](docs/DsvClientCreate.md)
 - [DsvClientCredentialsResponse](docs/DsvClientCredentialsResponse.md)
 - [DsvClientSearchModel](docs/DsvClientSearchModel.md)
 - [DsvCondition](docs/DsvCondition.md)
 - [DsvConditions](docs/DsvConditions.md)
 - [DsvConfigResponse](docs/DsvConfigResponse.md)
 - [DsvConfigVersionResponse](docs/DsvConfigVersionResponse.md)
 - [DsvCreateGroup](docs/DsvCreateGroup.md)
 - [DsvDecryptionResponse](docs/DsvDecryptionResponse.md)
 - [DsvDefaultPolicy](docs/DsvDefaultPolicy.md)
 - [DsvEncryptionResponse](docs/DsvEncryptionResponse.md)
 - [DsvEngineCreate](docs/DsvEngineCreate.md)
 - [DsvEngineCreateResponse](docs/DsvEngineCreateResponse.md)
 - [DsvEngineGetResponse](docs/DsvEngineGetResponse.md)
 - [DsvEngineListResult](docs/DsvEngineListResult.md)
 - [DsvEnginePingResponse](docs/DsvEnginePingResponse.md)
 - [DsvEngineSearchResponse](docs/DsvEngineSearchResponse.md)
 - [DsvGenerateResponse](docs/DsvGenerateResponse.md)
 - [DsvGroup](docs/DsvGroup.md)
 - [DsvGroupMemberInfo](docs/DsvGroupMemberInfo.md)
 - [DsvGroupResponse](docs/DsvGroupResponse.md)
 - [DsvGroupSearch](docs/DsvGroupSearch.md)
 - [DsvHistory](docs/DsvHistory.md)
 - [DsvHttpError](docs/DsvHttpError.md)
 - [DsvInitiateCertAuthResponse](docs/DsvInitiateCertAuthResponse.md)
 - [DsvKey](docs/DsvKey.md)
 - [DsvLogSearchResponse](docs/DsvLogSearchResponse.md)
 - [DsvManualKeyData](docs/DsvManualKeyData.md)
 - [DsvMasterkey](docs/DsvMasterkey.md)
 - [DsvMemberRequest](docs/DsvMemberRequest.md)
 - [DsvMemberResponse](docs/DsvMemberResponse.md)
 - [DsvMessageResponse](docs/DsvMessageResponse.md)
 - [DsvPageInfo](docs/DsvPageInfo.md)
 - [DsvPasswordChangeModel](docs/DsvPasswordChangeModel.md)
 - [DsvPolicyCreate](docs/DsvPolicyCreate.md)
 - [DsvPolicyResponse](docs/DsvPolicyResponse.md)
 - [DsvPolicySearchResponse](docs/DsvPolicySearchResponse.md)
 - [DsvPolicyUpdate](docs/DsvPolicyUpdate.md)
 - [DsvPolicyVersionResponse](docs/DsvPolicyVersionResponse.md)
 - [DsvPool](docs/DsvPool.md)
 - [DsvPoolCreate](docs/DsvPoolCreate.md)
 - [DsvPoolListResult](docs/DsvPoolListResult.md)
 - [DsvPostConfigModel](docs/DsvPostConfigModel.md)
 - [DsvResponseCertificate](docs/DsvResponseCertificate.md)
 - [DsvResponseRootCertificate](docs/DsvResponseRootCertificate.md)
 - [DsvResponseSshCertificate](docs/DsvResponseSshCertificate.md)
 - [DsvRoleCreate](docs/DsvRoleCreate.md)
 - [DsvRoleDetailsModel](docs/DsvRoleDetailsModel.md)
 - [DsvRoleResponse](docs/DsvRoleResponse.md)
 - [DsvRoleSearchResponse](docs/DsvRoleSearchResponse.md)
 - [DsvRoleVersionResponse](docs/DsvRoleVersionResponse.md)
 - [DsvRoles](docs/DsvRoles.md)
 - [DsvRootCaRegistration](docs/DsvRootCaRegistration.md)
 - [DsvRootCaSecret](docs/DsvRootCaSecret.md)
 - [DsvSecretCreate](docs/DsvSecretCreate.md)
 - [DsvSecretDescription](docs/DsvSecretDescription.md)
 - [DsvSecretListPathsResponse](docs/DsvSecretListPathsResponse.md)
 - [DsvSecretResponse](docs/DsvSecretResponse.md)
 - [DsvSecretSearch](docs/DsvSecretSearch.md)
 - [DsvSecretUpdate](docs/DsvSecretUpdate.md)
 - [DsvSecretVersionResponse](docs/DsvSecretVersionResponse.md)
 - [DsvSettings](docs/DsvSettings.md)
 - [DsvSiemCreateUpdateRequestModel](docs/DsvSiemCreateUpdateRequestModel.md)
 - [DsvSiemNoSensitiveResponseModel](docs/DsvSiemNoSensitiveResponseModel.md)
 - [DsvSiemResponse](docs/DsvSiemResponse.md)
 - [DsvSiemSearchResponse](docs/DsvSiemSearchResponse.md)
 - [DsvSignedLeafCertificate](docs/DsvSignedLeafCertificate.md)
 - [DsvSigningRequest](docs/DsvSigningRequest.md)
 - [DsvSigningRequestInformation](docs/DsvSigningRequestInformation.md)
 - [DsvSshCertInformation](docs/DsvSshCertInformation.md)
 - [DsvStatusResponse](docs/DsvStatusResponse.md)
 - [DsvTaskResult](docs/DsvTaskResult.md)
 - [DsvTaskState](docs/DsvTaskState.md)
 - [DsvUpdateKeyRequest](docs/DsvUpdateKeyRequest.md)
 - [DsvUsageResponseGeneral](docs/DsvUsageResponseGeneral.md)
 - [DsvUserCreateModel](docs/DsvUserCreateModel.md)
 - [DsvUserResponse](docs/DsvUserResponse.md)
 - [DsvUserSearch](docs/DsvUserSearch.md)
 - [DsvUserUpdateModel](docs/DsvUserUpdateModel.md)
 - [DsvUserVersionResponse](docs/DsvUserVersionResponse.md)


## Documentation For Authorization

## Bearer
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



