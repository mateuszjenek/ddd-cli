package service

import (
	"context"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
)

type NoteRepository interface {
	Create(context.Context, model.Note) (model.Note, error)
	Get(context.Context, valueobject.Id) (model.Note, error)
	List(context.Context) ([]model.Note, error)
	Update(context.Context, model.Note) (model.Note, error)
	Remove(context.Context, valueobject.Id) error

}