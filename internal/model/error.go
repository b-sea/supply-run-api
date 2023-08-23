package model

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Code int      `json:"code"`
	ID   GlobalID `json:"message"`
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%d not found: %s", e.Code, e.ID.Key)
}

func NewNotFoundError(id GlobalID) *NotFoundError {
	return &NotFoundError{
		Code: http.StatusNotFound,
		ID:   id,
	}
}

type ConflictError struct {
	Code     int    `json:"code"`
	Revision string `json:"revision"`
}

func (e ConflictError) Error() string {
	return fmt.Sprintf("%d conflict error: %s", e.Code, e.Revision)
}

func NewConflictError(revision string) *ConflictError {
	return &ConflictError{
		Code:     http.StatusConflict,
		Revision: revision,
	}
}

type ServerError struct {
	Code       int   `json:"code"`
	InnerError error `json:"error"`
}

func (e ServerError) Error() string {
	return fmt.Sprintf("%d internal server error: %v", e.Code, e.InnerError)
}

func NewServerError(err error) *ServerError {
	return &ServerError{
		Code:       http.StatusInternalServerError,
		InnerError: err,
	}
}
