// Code generated by go-swagger; DO NOT EDIT.

package eaa_s_manual

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new eaa s manual API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for eaa s manual API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DecryptWithManualKey(params *DecryptWithManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecryptWithManualKeyOK, error)

	DeleteManualKey(params *DeleteManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteManualKeyOK, *DeleteManualKeyNoContent, error)

	EncryptWithManualKey(params *EncryptWithManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EncryptWithManualKeyCreated, error)

	ReadManualKey(params *ReadManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadManualKeyOK, error)

	RestoreManualKey(params *RestoreManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreManualKeyNoContent, error)

	UpdateKey(params *UpdateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateKeyOK, error)

	UploadKey(params *UploadKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadKeyCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DecryptWithManualKey decrypts

  Decrypt ciphertext with a key.
*/
func (a *Client) DecryptWithManualKey(params *DecryptWithManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DecryptWithManualKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecryptWithManualKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "decryptWithManualKey",
		Method:             "POST",
		PathPattern:        "/crypto/manual/decrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DecryptWithManualKeyReader{formats: a.formats},
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
	success, ok := result.(*DecryptWithManualKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for decryptWithManualKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteManualKey deletes key

  Delete an existing encryption/decryption key.
*/
func (a *Client) DeleteManualKey(params *DeleteManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteManualKeyOK, *DeleteManualKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManualKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteManualKey",
		Method:             "DELETE",
		PathPattern:        "/crypto/manual/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteManualKeyReader{formats: a.formats},
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
	case *DeleteManualKeyOK:
		return value, nil, nil
	case *DeleteManualKeyNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for eaa_s_manual: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EncryptWithManualKey encrypts

  Encrypt plaintext with a key.
*/
func (a *Client) EncryptWithManualKey(params *EncryptWithManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EncryptWithManualKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEncryptWithManualKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "encryptWithManualKey",
		Method:             "POST",
		PathPattern:        "/crypto/manual/encrypt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EncryptWithManualKeyReader{formats: a.formats},
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
	success, ok := result.(*EncryptWithManualKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for encryptWithManualKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadManualKey reads key

  Read existing encryption/decryption key.
*/
func (a *Client) ReadManualKey(params *ReadManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReadManualKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadManualKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "readManualKey",
		Method:             "GET",
		PathPattern:        "/crypto/manual/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadManualKeyReader{formats: a.formats},
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
	success, ok := result.(*ReadManualKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readManualKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestoreManualKey restores key

  Restore a soft-deleted encryption/decryption key.
*/
func (a *Client) RestoreManualKey(params *RestoreManualKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreManualKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreManualKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restoreManualKey",
		Method:             "PUT",
		PathPattern:        "/crypto/manual/key/{path}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreManualKeyReader{formats: a.formats},
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
	success, ok := result.(*RestoreManualKeyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreManualKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateKey updates key

  Update an existing encryption/decryption key.
*/
func (a *Client) UpdateKey(params *UpdateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateKey",
		Method:             "PUT",
		PathPattern:        "/crypto/manual/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateKeyReader{formats: a.formats},
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
	success, ok := result.(*UpdateKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadKey uploads key

  Upload a new encryption/decryption key.
*/
func (a *Client) UploadKey(params *UploadKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadKey",
		Method:             "POST",
		PathPattern:        "/crypto/manual/key/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadKeyReader{formats: a.formats},
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
	success, ok := result.(*UploadKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
