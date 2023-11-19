package infrastructure

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
	"github.com/mateuszjenek/ddd-cli/internal/domain/service"
	"github.com/mateuszjenek/ddd-cli/internal/infrastructure/port"
)

type noteLocalRepository struct{
	dbPort port.Database
}

func NewNoteLocalRepository(dbPort port.Database) service.NoteRepository {
	return &noteLocalRepository{ dbPort: dbPort }
}

func (n *noteLocalRepository) Create(ctx context.Context, note model.Note) (model.Note, error) {
	idValue, err := valueobject.NewId(uuid.New().String())
	if err != nil {
		return note, fmt.Errorf("failed to create a note id value object")
	}
	note.Id = idValue

	_, err = n.dbPort.GetDB().ExecContext(ctx, "INSERT INTO notes VALUES($1, $2, $3)", note.Id.String(), note.Author.String(), note.Message.String())
	if err != nil {
		return note, fmt.Errorf("failed to insert a note: %w", err)
	}

	return note, nil
}

func (*noteLocalRepository) Get(context.Context, valueobject.Id) (model.Note, error) {
	panic("unimplemented")
}

func (n *noteLocalRepository) List(ctx context.Context) ([]model.Note, error) {
	rows, err := n.dbPort.GetDB().QueryContext(ctx, "SELECT * FROM notes")
	if err != nil {
		return nil, fmt.Errorf("failed to select notes: %w", err)
	}
	defer rows.Close()

	notes := make([]model.Note, 0)
	for rows.Next() {
		var id, author, message string
		err := rows.Scan(&id, &author, &message)
		if err != nil {
			return nil, fmt.Errorf("failed to scan a note row: %w", err)
		}
		idValue, err := valueobject.NewId(id)
		if err != nil {
			return nil, fmt.Errorf("failed to create a id value object: %w", err)
		}
		authorValue, err := valueobject.NewAuthor(author)
		if err != nil {
			return nil, fmt.Errorf("failed to create an author value object: %w", err)
		}
		messageValue, err := valueobject.NewMessage(message)
		if err != nil {
			return nil, fmt.Errorf("failed to create a message value object: %w", err)
		}
		note := model.Note{
			Id: idValue,
			Author: authorValue,
			Message: messageValue,
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (*noteLocalRepository) Remove(context.Context, valueobject.Id) error {
	panic("unimplemented")
}

func (*noteLocalRepository) Update(context.Context, model.Note) (model.Note, error) {
	panic("unimplemented")
}
