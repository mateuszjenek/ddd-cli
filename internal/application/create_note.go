package application

import (
	"context"
	"fmt"
	"time"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
	"github.com/mateuszjenek/ddd-cli/internal/domain/service"
)

type CreateNoteResult struct {
	Note model.Note
}

type CreateNoteInputPort interface {
	CreateNote(message, author string) (*CreateNoteResult, error)
}

var _ CreateNoteInputPort = (*createNoteUseCase)(nil)

type createNoteUseCase struct {
	noteRepository service.NoteRepository
}

func NewCreateNoteUseCase(noteRepository service.NoteRepository) CreateNoteInputPort {
	return &createNoteUseCase{
		noteRepository: noteRepository,
	}
}

func (c *createNoteUseCase) CreateNote(message, author string) (*CreateNoteResult, error) {
	authorValue, err := valueobject.NewAuthor(author)
	if err != nil {
		return nil, fmt.Errorf("failed to create an author value object: %v", err)
	}

	messageValue, err := valueobject.NewMessage(message)
	if err != nil {
		return nil, fmt.Errorf("failed to create a message value object: %v", err)
	}

	note := model.Note{
		Author:  authorValue,
		Message: messageValue,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	note, err = c.noteRepository.Create(ctx, note)
	if err != nil {
		return nil, fmt.Errorf("failed to create a note: %v", err)
	}

	return &CreateNoteResult{Note: note}, nil
}
