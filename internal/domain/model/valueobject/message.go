package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Message struct {
	msg string
}

func NewMessage(msg string) (Message, error) {
	m := Message{msg}
	if err := m.validate(); err != nil {
		return m, fmt.Errorf("author is invalid: %w", err)
	}
	return m, nil
}

func (m Message) GetMessage() string {
	return m.msg
}

func (m Message) String() string {
	return m.msg
}

func (m Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.msg)
}

func (m Message) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.msg)
}

func (m Message) validate() error {
	if len(m.msg) == 0 {
		return errors.New("author name cannot be empty")
	} else if len(m.msg) > 256 {
		return errors.New("author name cannot have more than 256 characters")
	}
	return nil
}
