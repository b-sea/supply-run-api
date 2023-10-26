// Package service implements all business logic for the API.
package service

import "fmt"

var (
	ErrAuthentication = fmt.Errorf("authentication error")
	ErrAuthorization  = fmt.Errorf("authorization error")
)
