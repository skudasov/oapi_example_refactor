/*
 * Insolar Authorization Service API
 *
 * # Authorization service API  This document provides an API reference for the Insolar authentication service (auth-service).  Auth-service allows the Insolar Observer nodes to register and authenticate in the Insolar network.  ## Connect your Observer node to the Insolar network  To connect, complete the following steps:  1. Acquire a login and a registration code from Insolar.     The code is unique, lives for 24 hours, expires upon first use, and allows you to set your password.  2. [Set a password](#operation/set-password) with the provided code, login, and your password in parameters.  Observer node authenticates in the Insolar network in the following way:  1. Using basic authorization (your login and password), the Observer requests a new JSON Web Token (JWT) for every request.     The auth-service checks the login-password pair and, if valid, issues a JWT for authentication.  2. Then, the Observer substitutes the corresponding header of every request to the Insolar network with the acquired JWT.     Insolar's heavy node checks the JWT's validity and expiry timestamp.
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// ResponsesToken struct for ResponsesToken
type ResponsesToken struct {
	// JSON Web Token that grants a request invocation right in the Insolar network. Lives for 15 minutes and expires upon first use.
	AccessToken string `json:"access_token"`
	// Token expiration timestamp in Unix format.
	Expiration int64 `json:"expiration"`
}
