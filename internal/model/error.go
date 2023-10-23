package model

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	Code() int
}

type NotFoundError struct {
	ID ID `json:"id"`
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%d Not Found: %s", e.Code(), e.ID)
}

func (e NotFoundError) Code() int {
	return http.StatusNotFound
}

type BadRequestError struct {
	Reason interface{} `json:"message"`
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("%d Bad Request: %v", e.Code(), e.Reason)
}

func (e BadRequestError) Code() int {
	return http.StatusBadRequest
}

type AuthenticationError struct{}

func (e AuthenticationError) Error() string {
	return fmt.Sprintf("%d Unauthorized", e.Code())
}

func (e AuthenticationError) Code() int {
	return http.StatusUnauthorized
}

type AuthorizationError struct{}

func (e AuthorizationError) Error() string {
	return fmt.Sprintf("%d Forbidden", e.Code())
}

func (e AuthorizationError) Code() int {
	return http.StatusForbidden
}
