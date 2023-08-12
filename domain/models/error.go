package models

import "fmt"

type ErrNotFound struct {
	value string
}

type ErrInternalServer struct{}

const (
	teamID    = "teamID"
	articleID = "articleID"
)

var (
	ErrTeamIDNotFound    = NewErrNotFound(teamID)
	ErrArticleIDNotFound = NewErrNotFound(articleID)
)

func NewErrNotFound(value string) *ErrNotFound {
	return &ErrNotFound{value}
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", err.value)
}

func (err ErrInternalServer) Error() string {
	return "internal server error"
}
