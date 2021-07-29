/*
 * User API
 *
 * An API for creating and authenticating system users
 *
 * API version: 0.1.0
 * Contact: joe@bebo.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"net/http"
)

// AuthAPIApiRouter defines the required methods for binding the api requests to a responses for the AuthAPIApi
// The AuthAPIApiRouter implementation should parse necessary information from the http request,
// pass the data to a AuthAPIApiServicer to perform the required actions, then write the service results to the http response.
type AuthAPIApiRouter interface {
	AddAccount(http.ResponseWriter, *http.Request)
	AuthenticateUser(http.ResponseWriter, *http.Request)
	Ping(http.ResponseWriter, *http.Request)
}

// AuthAPIApiServicer defines the api actions for the AuthAPIApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AuthAPIApiServicer interface {
	AddAccount(context.Context, UserAccountDetails) (ImplResponse, error)
	AuthenticateUser(context.Context, UserCredentials) (ImplResponse, error)
	Ping(context.Context) (ImplResponse, error)
}