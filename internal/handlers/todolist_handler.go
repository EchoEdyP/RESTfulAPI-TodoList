package handlers

import (
	"LearnECHO/models/requestAndresponse"
	"github.com/labstack/echo/v4"
)

type TodoListHandler interface {
	Create(ctx echo.Context, request requestAndresponse.TodoListCreateRequest) error
	ReadAll(ctx echo.Context) error
	ReadById(ctx echo.Context, todolistId int) error
	UpdateTitleAndDescription(ctx echo.Context, todolistId int, request requestAndresponse.TodoListUpdateTitleDescription) error
	UpdateStatus(ctx echo.Context, todolistId int, request requestAndresponse.TodoListUpdateStatus) error
	Delete(ctx echo.Context, todolistId int) error
}
