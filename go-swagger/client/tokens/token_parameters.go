// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTokenParams creates a new TokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTokenParams() *TokenParams {
	return &TokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTokenParamsWithTimeout creates a new TokenParams object
// with the ability to set a timeout on a request.
func NewTokenParamsWithTimeout(timeout time.Duration) *TokenParams {
	return &TokenParams{
		timeout: timeout,
	}
}

// NewTokenParamsWithContext creates a new TokenParams object
// with the ability to set a context for a request.
func NewTokenParamsWithContext(ctx context.Context) *TokenParams {
	return &TokenParams{
		Context: ctx,
	}
}

// NewTokenParamsWithHTTPClient creates a new TokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewTokenParamsWithHTTPClient(client *http.Client) *TokenParams {
	return &TokenParams{
		HTTPClient: client,
	}
}

/* TokenParams contains all the parameters to send to the API endpoint
   for the token operation.

   Typically these are written to a http.Request.
*/
type TokenParams struct {

	/* AwsBody.

	   Base64 encoded signed aws request body for aws_iam grant type
	*/
	AwsBody *string

	/* AwsHeaders.

	   Base64 encoded signed aws request headers for aws_iam grant type
	*/
	AwsHeaders *string

	/* CertChallengeID.

	   Challenge id for the certificate grant type
	*/
	CertChallengeID *string

	/* ClientID.

	   Client id for client_credentials grant type
	*/
	ClientID *string

	/* ClientSecret.

	   Client secret for client_credentials grant type
	*/
	ClientSecret *string

	/* DecryptedChallenge.

	   Decrypted and base64 encoded challenge for the certificate grant type
	*/
	DecryptedChallenge *string

	// GrantType.
	GrantType string

	/* Jwt.

	   JWT token for azure and gcp grant types
	*/
	Jwt *string

	/* Password.

	   Password for password grant type
	*/
	Password *string

	/* Provider.

	   Provider name for password grant flow for Thycotic One authentication
	*/
	Provider *string

	/* RefreshToken.

	   Previously issued refresh token for the refresh_token grant type
	*/
	RefreshToken *string

	/* Username.

	   Username for password grant type
	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenParams) WithDefaults() *TokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the token params
func (o *TokenParams) WithTimeout(timeout time.Duration) *TokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token params
func (o *TokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token params
func (o *TokenParams) WithContext(ctx context.Context) *TokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token params
func (o *TokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token params
func (o *TokenParams) WithHTTPClient(client *http.Client) *TokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token params
func (o *TokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAwsBody adds the awsBody to the token params
func (o *TokenParams) WithAwsBody(awsBody *string) *TokenParams {
	o.SetAwsBody(awsBody)
	return o
}

// SetAwsBody adds the awsBody to the token params
func (o *TokenParams) SetAwsBody(awsBody *string) {
	o.AwsBody = awsBody
}

// WithAwsHeaders adds the awsHeaders to the token params
func (o *TokenParams) WithAwsHeaders(awsHeaders *string) *TokenParams {
	o.SetAwsHeaders(awsHeaders)
	return o
}

// SetAwsHeaders adds the awsHeaders to the token params
func (o *TokenParams) SetAwsHeaders(awsHeaders *string) {
	o.AwsHeaders = awsHeaders
}

// WithCertChallengeID adds the certChallengeID to the token params
func (o *TokenParams) WithCertChallengeID(certChallengeID *string) *TokenParams {
	o.SetCertChallengeID(certChallengeID)
	return o
}

// SetCertChallengeID adds the certChallengeId to the token params
func (o *TokenParams) SetCertChallengeID(certChallengeID *string) {
	o.CertChallengeID = certChallengeID
}

// WithClientID adds the clientID to the token params
func (o *TokenParams) WithClientID(clientID *string) *TokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the token params
func (o *TokenParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the token params
func (o *TokenParams) WithClientSecret(clientSecret *string) *TokenParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the token params
func (o *TokenParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithDecryptedChallenge adds the decryptedChallenge to the token params
func (o *TokenParams) WithDecryptedChallenge(decryptedChallenge *string) *TokenParams {
	o.SetDecryptedChallenge(decryptedChallenge)
	return o
}

// SetDecryptedChallenge adds the decryptedChallenge to the token params
func (o *TokenParams) SetDecryptedChallenge(decryptedChallenge *string) {
	o.DecryptedChallenge = decryptedChallenge
}

// WithGrantType adds the grantType to the token params
func (o *TokenParams) WithGrantType(grantType string) *TokenParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the token params
func (o *TokenParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithJwt adds the jwt to the token params
func (o *TokenParams) WithJwt(jwt *string) *TokenParams {
	o.SetJwt(jwt)
	return o
}

// SetJwt adds the jwt to the token params
func (o *TokenParams) SetJwt(jwt *string) {
	o.Jwt = jwt
}

// WithPassword adds the password to the token params
func (o *TokenParams) WithPassword(password *string) *TokenParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the token params
func (o *TokenParams) SetPassword(password *string) {
	o.Password = password
}

// WithProvider adds the provider to the token params
func (o *TokenParams) WithProvider(provider *string) *TokenParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the token params
func (o *TokenParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithRefreshToken adds the refreshToken to the token params
func (o *TokenParams) WithRefreshToken(refreshToken *string) *TokenParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the token params
func (o *TokenParams) SetRefreshToken(refreshToken *string) {
	o.RefreshToken = refreshToken
}

// WithUsername adds the username to the token params
func (o *TokenParams) WithUsername(username *string) *TokenParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the token params
func (o *TokenParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *TokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AwsBody != nil {

		// form param aws_body
		var frAwsBody string
		if o.AwsBody != nil {
			frAwsBody = *o.AwsBody
		}
		fAwsBody := frAwsBody
		if fAwsBody != "" {
			if err := r.SetFormParam("aws_body", fAwsBody); err != nil {
				return err
			}
		}
	}

	if o.AwsHeaders != nil {

		// form param aws_headers
		var frAwsHeaders string
		if o.AwsHeaders != nil {
			frAwsHeaders = *o.AwsHeaders
		}
		fAwsHeaders := frAwsHeaders
		if fAwsHeaders != "" {
			if err := r.SetFormParam("aws_headers", fAwsHeaders); err != nil {
				return err
			}
		}
	}

	if o.CertChallengeID != nil {

		// form param cert_challenge_id
		var frCertChallengeID string
		if o.CertChallengeID != nil {
			frCertChallengeID = *o.CertChallengeID
		}
		fCertChallengeID := frCertChallengeID
		if fCertChallengeID != "" {
			if err := r.SetFormParam("cert_challenge_id", fCertChallengeID); err != nil {
				return err
			}
		}
	}

	if o.ClientID != nil {

		// form param client_id
		var frClientID string
		if o.ClientID != nil {
			frClientID = *o.ClientID
		}
		fClientID := frClientID
		if fClientID != "" {
			if err := r.SetFormParam("client_id", fClientID); err != nil {
				return err
			}
		}
	}

	if o.ClientSecret != nil {

		// form param client_secret
		var frClientSecret string
		if o.ClientSecret != nil {
			frClientSecret = *o.ClientSecret
		}
		fClientSecret := frClientSecret
		if fClientSecret != "" {
			if err := r.SetFormParam("client_secret", fClientSecret); err != nil {
				return err
			}
		}
	}

	if o.DecryptedChallenge != nil {

		// form param decrypted_challenge
		var frDecryptedChallenge string
		if o.DecryptedChallenge != nil {
			frDecryptedChallenge = *o.DecryptedChallenge
		}
		fDecryptedChallenge := frDecryptedChallenge
		if fDecryptedChallenge != "" {
			if err := r.SetFormParam("decrypted_challenge", fDecryptedChallenge); err != nil {
				return err
			}
		}
	}

	// form param grant_type
	frGrantType := o.GrantType
	fGrantType := frGrantType
	if fGrantType != "" {
		if err := r.SetFormParam("grant_type", fGrantType); err != nil {
			return err
		}
	}

	if o.Jwt != nil {

		// form param jwt
		var frJwt string
		if o.Jwt != nil {
			frJwt = *o.Jwt
		}
		fJwt := frJwt
		if fJwt != "" {
			if err := r.SetFormParam("jwt", fJwt); err != nil {
				return err
			}
		}
	}

	if o.Password != nil {

		// form param password
		var frPassword string
		if o.Password != nil {
			frPassword = *o.Password
		}
		fPassword := frPassword
		if fPassword != "" {
			if err := r.SetFormParam("password", fPassword); err != nil {
				return err
			}
		}
	}

	if o.Provider != nil {

		// form param provider
		var frProvider string
		if o.Provider != nil {
			frProvider = *o.Provider
		}
		fProvider := frProvider
		if fProvider != "" {
			if err := r.SetFormParam("provider", fProvider); err != nil {
				return err
			}
		}
	}

	if o.RefreshToken != nil {

		// form param refresh_token
		var frRefreshToken string
		if o.RefreshToken != nil {
			frRefreshToken = *o.RefreshToken
		}
		fRefreshToken := frRefreshToken
		if fRefreshToken != "" {
			if err := r.SetFormParam("refresh_token", fRefreshToken); err != nil {
				return err
			}
		}
	}

	if o.Username != nil {

		// form param username
		var frUsername string
		if o.Username != nil {
			frUsername = *o.Username
		}
		fUsername := frUsername
		if fUsername != "" {
			if err := r.SetFormParam("username", fUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
