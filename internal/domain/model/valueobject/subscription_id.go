package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SubscriptionId struct {
	id string
}

func NewSubscriptionId(id string) (SubscriptionId, error) {
	i := SubscriptionId{id}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("failed to validate subscription id: %w", err)
	}
	return i, nil
}

func (s SubscriptionId) GetId() string {
	return s.id
}

func (s SubscriptionId) String() string {
	return s.id
}

func (s SubscriptionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func (s SubscriptionId) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.id)
}

func (s SubscriptionId) validate() error {
	if len(s.id) == 0 {
		return errors.New("subscription id cannot be empty")
	} else if len(s.id) > 256 {
		return errors.New("subscription id cannot have more than 256 characters")
	}
	return nil
}
