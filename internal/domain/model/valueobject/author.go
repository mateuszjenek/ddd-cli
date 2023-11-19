package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Author struct {
	name string
}

func NewAuthor(name string) (Author, error) {
	a := Author{name}
	if err := a.validate(); err != nil {
		return a, fmt.Errorf("author is invalid: %w", err)
	}
	return a, nil
}

func (a Author) GetName() string {
	return a.name
}

func (a Author) String() string {
	return a.name
}

func (a Author) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.name)
}

func (a Author) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.name)
}

func (a Author) validate() error {
	if len(a.name) == 0 {
		return errors.New("author name cannot be empty")
	} else if len(a.name) > 50 {
		return errors.New("author name cannot have more than 256 characters")
	}
	return nil
}
