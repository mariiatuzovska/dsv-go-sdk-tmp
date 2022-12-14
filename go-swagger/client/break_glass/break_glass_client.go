// Code generated by go-swagger; DO NOT EDIT.

package break_glass

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new break glass API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for break glass API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ApplyRequest(params *ApplyRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ApplyRequestCreated, error)

	GenerateRequest(params *GenerateRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GenerateRequestCreated, error)

	GetStatus(params *GetStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ApplyRequest applies

  Apply secret shares to break glass and give users admin rights
*/
func (a *Client) ApplyRequest(params *ApplyRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ApplyRequestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplyRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ApplyRequest",
		Method:             "POST",
		PathPattern:        "/breakglass/apply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplyRequestReader{formats: a.formats},
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
	success, ok := result.(*ApplyRequestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApplyRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GenerateRequest generates

  Generate a new break glass secret and shares
*/
func (a *Client) GenerateRequest(params *GenerateRequestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GenerateRequestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GenerateRequest",
		Method:             "POST",
		PathPattern:        "/breakglass/generate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenerateRequestReader{formats: a.formats},
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
	success, ok := result.(*GenerateRequestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GenerateRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStatus gets status

  Get break glass status
*/
func (a *Client) GetStatus(params *GetStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStatus",
		Method:             "GET",
		PathPattern:        "/breakglass",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStatusReader{formats: a.formats},
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
	success, ok := result.(*GetStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
