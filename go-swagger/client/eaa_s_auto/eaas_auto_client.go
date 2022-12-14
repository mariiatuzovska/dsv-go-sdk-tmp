// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_auto

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new eaa s auto API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for eaa s auto API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RotationRequest(params *RotationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RotationRequestCreated, error)

	CreateKey(params *CreateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateKeyCreated, error)

	Decrypt(params *DecryptParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecryptOK, error)

	DeleteKey(params *DeleteKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKeyOK, *DeleteKeyNoContent, error)

	Encrypt(params *EncryptParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EncryptCreated, error)

	GetKeyMetadata(params *GetKeyMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKeyMetadataOK, error)

	RestoreKey(params *RestoreKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreKeyNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  RotationRequest rotates data and key

  Rotate data and optionally an existing encryption/decryption key.

If the starting version is the current version of the key, then DSV will rotate the key (create a new version of it)
and re-encrypt the data using this new version.

If the starting version is NOT the current version of the key, and the ending version is not provided, then DSV will
only re-encrypt the data using the current latest version of the key.

The starting and ending versions can also be below the latest one, so long as the starting is below the ending.
In this case, DSV will re-encrypt the data using the version of the key specified by the ending version.
*/
func (a *Client) RotationRequest(params *RotationRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RotationRequestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRotationRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RotationRequest",
		Method:             "POST",
		PathPattern:        "/crypto/rotate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RotationRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RotationRequestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RotationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateKey creates key

  Create a new encryption/decryption key.
*/
func (a *Client) CreateKey(params *CreateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createKey",
		Method:             "POST",
		PathPattern:        "/crypto/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Decrypt decrypts

  Decrypt ciphertext with a key.
*/
func (a *Client) Decrypt(params *DecryptParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecryptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecryptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "decrypt",
		Method:             "POST",
		PathPattern:        "/crypto/decrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DecryptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DecryptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for decrypt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteKey deletes key

  Delete an existing encryption/decryption key.
*/
func (a *Client) DeleteKey(params *DeleteKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKeyOK, *DeleteKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteKey",
		Method:             "DELETE",
		PathPattern:        "/crypto/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteKeyOK:
		return value, nil, nil
	case *DeleteKeyNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for eaa_s_auto: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Encrypt encrypts

  Encrypt plaintext with a key.
*/
func (a *Client) Encrypt(params *EncryptParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EncryptCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEncryptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "encrypt",
		Method:             "POST",
		PathPattern:        "/crypto/encrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EncryptReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EncryptCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for encrypt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetKeyMetadata gets key metadata

  Get metadata of an existing encryption/decryption key.
*/
func (a *Client) GetKeyMetadata(params *GetKeyMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKeyMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKeyMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getKeyMetadata",
		Method:             "GET",
		PathPattern:        "/crypto/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKeyMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetKeyMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getKeyMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestoreKey restores key

  Restore a soft-deleted encryption/decryption key.
*/
func (a *Client) RestoreKey(params *RestoreKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restoreKey",
		Method:             "PUT",
		PathPattern:        "/crypto/key/{path}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreKeyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
