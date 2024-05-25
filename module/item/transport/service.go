package transport

import (
	"context"
	"social-todo-list/module/item/model"
)

type ItemUseCase interface {
	createItem(ctx context.Context, data *model.TodoItemCreation) error
	GetItemById(ctx context.Context, id int) (*model.TodoItem, error)
	UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error
}

type itemService struct {
	useCase ItemUseCase
}
