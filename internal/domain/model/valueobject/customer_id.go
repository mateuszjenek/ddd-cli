package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CustomerId struct {
	id string
}

func NewCustomerId(id string) (CustomerId, error) {
	i := CustomerId{id}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("failed to validate customer id: %w", err)
	}
	return i, nil
}

func (c CustomerId) GetId() string {
	return c.id
}

func (c CustomerId) String() string {
	return c.id
}

func (c CustomerId) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.id)
}

func (c CustomerId) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.id)
}

func (c CustomerId) validate() error {
	if len(c.id) == 0 {
		return errors.New("customer id cannot be empty")
	} else if len(c.id) > 256 {
		return errors.New("customer id cannot have more than 256 characters")
	}
	return nil
}
