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
	"encoding/json"
	"net/http"
	"strings"
)

// A AuthAPIApiController binds http requests to an api service and writes the service results to the http response
type AuthAPIApiController struct {
	service AuthAPIApiServicer
}

// NewAuthAPIApiController creates a default api controller
func NewAuthAPIApiController(s AuthAPIApiServicer) Router {
	return &AuthAPIApiController{service: s}
}

// Routes returns all of the api route for the AuthAPIApiController
func (c *AuthAPIApiController) Routes() Routes {
	return Routes{
		{
			"AddAccount",
			strings.ToUpper("Put"),
			"/addAccount",
			c.AddAccount,
		},
		{
			"AuthenticateUser",
			strings.ToUpper("Post"),
			"/authenticate",
			c.AuthenticateUser,
		},
		{
			"Ping",
			strings.ToUpper("Get"),
			"/ping",
			c.Ping,
		},
	}
}

// AddAccount - adds a new account
func (c *AuthAPIApiController) AddAccount(w http.ResponseWriter, r *http.Request) {
	userAccountDetails := &UserAccountDetails{}
	if err := json.NewDecoder(r.Body).Decode(&userAccountDetails); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.AddAccount(r.Context(), *userAccountDetails)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AuthenticateUser - authenticates an existing user
func (c *AuthAPIApiController) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	userCredentials := &UserCredentials{}
	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.AuthenticateUser(r.Context(), *userCredentials)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Ping - tests this api
func (c *AuthAPIApiController) Ping(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Ping(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}