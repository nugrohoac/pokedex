package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_errors "github.com/pkg/errors"

	"github.com/nugrohoac/pokedex"
)

// ErrHandler .
func ErrHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		err := handler(c)
		if err == nil {
			return nil
		}

		switch errCause := _errors.Cause(err).(type) {
		case pokedex.ErrDuplicated: // error code 400
			return c.JSON(http.StatusConflict, errCause)
		case pokedex.ErrBindStruct:
			return c.JSON(http.StatusBadRequest, errCause)
		case pokedex.ErrorAuth: // error code 401
			return c.JSON(http.StatusUnauthorized, errCause)
		case pokedex.ErrInValid:
			return c.JSON(http.StatusUnauthorized, errCause)
		case pokedex.ErrNotFound: // error code 404
			return c.JSON(http.StatusNotFound, errCause)
		default: // default 500
			return c.JSON(http.StatusInternalServerError, errCause)
		}

	}
}
