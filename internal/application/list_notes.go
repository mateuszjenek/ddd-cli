package application

import (
	"context"
	"fmt"
	"time"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/service"
)

type ListNotesResult struct {
	Notes []model.Note
}

type ListNotesInputPort interface {
	ListNotes() (*ListNotesResult, error)
}

var _ ListNotesInputPort = (*listNotesUseCase)(nil)

type listNotesUseCase struct {
	noteRepository service.NoteRepository
}

func NewListNotesUseCase(noteRepository service.NoteRepository) ListNotesInputPort {
	return &listNotesUseCase{
		noteRepository: noteRepository,
	}
}

func (l *listNotesUseCase) ListNotes() (*ListNotesResult, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	notes, err := l.noteRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list notes: %w", err)
	}
	return &ListNotesResult{ Notes: notes }, nil
}
