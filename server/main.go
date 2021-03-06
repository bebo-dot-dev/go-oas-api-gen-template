/*
 * User API
 *
 * An API for creating and authenticating system users
 *
 * API version: 0.1.0
 * Contact: joe@bebo.dev
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	api "github.com/bebo-dot-dev/go-oas-api-gen-template/server/api"
)

func main() {
	log.Printf("Server started")

	AuthAPIApiService := api.NewAuthAPIApiService()
	AuthAPIApiController := api.NewAuthAPIApiController(AuthAPIApiService)

	router := api.NewRouter(AuthAPIApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
