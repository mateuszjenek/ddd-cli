package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Id struct {
	id string
}

func NewId(id string) (Id, error) {
	i := Id{id}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("author is invalid: %w", err)
	}
	return i, nil
}

func (i Id) GetId() string {
	return i.id
}

func (i Id) String() string {
	return i.id
}

func (i Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.id)
}

func (i Id) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.id)
}

func (i Id) validate() error {
	if len(i.id) == 0 {
		return errors.New("author name cannot be empty")
	} else if len(i.id) > 256 {
		return errors.New("author name cannot have more than 256 characters")
	}
	return nil
}
