package pokedex

import (
	"fmt"
)

// ErrNotFound .
type ErrNotFound struct {
	Message string `json:"Message" example:"item is not found"`
}

func (enf ErrNotFound) Error() string {
	return enf.Message
}

// ErrInValid .
type ErrInValid struct {
	Message string
}

func (ei ErrInValid) Error() string {
	return ei.Message
}

// ErrDuplicated .
type ErrDuplicated struct {
	Message string `json:"Message" example:"item already exist"`
}

func (ed ErrDuplicated) Error() string {
	return ed.Message
}

// ErrBindStruct .
type ErrBindStruct struct {
	Message string `json:"Message" example:"failed bind struct"`
}

func (ebs ErrBindStruct) Error() string {
	return ebs.Message
}

// ErrValidateStruct .
type ErrValidateStruct struct {
	Message string `json:"Message" example:"error validate struct"`
}

func (evs ErrValidateStruct) Error() string {
	return evs.Message
}

// InternalError .
type InternalError struct {
	Path string
}

func (i InternalError) Error() string {
	return fmt.Sprintf("internal server error at : %v", i.Path)
}

// SyntaxError .
type SyntaxError struct {
	Line int
	Col  int
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", s.Line, s.Col)
}

// ErrorAuth is used to return error auth
type ErrorAuth struct {
	Message string `json:"Message" example:"error auth"`
}

func (ea ErrorAuth) Error() string {
	return ea.Message
}

// ErrNoRowAffected .
type ErrNoRowAffected struct {
	Message string `json:"Message" example:"no row affect"`
}

func (enr ErrNoRowAffected) Error() string {
	return enr.Message
}
