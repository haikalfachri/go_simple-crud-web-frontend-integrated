package models

import "github.com/labstack/echo"

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    T      `json:"data"`
}

func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
	return c.JSON(statusCode, Response[T]{
		Status:  statusMessage,
		Message: message,
		Data:    data,
	})
}