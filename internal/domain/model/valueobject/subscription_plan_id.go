package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SubscriptionPlanId struct {
	id string
}

func NewSubscriptionPlanId(id string) (SubscriptionPlanId, error) {
	i := SubscriptionPlanId{id}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("failed to validate subscription plan id: %w", err)
	}
	return i, nil
}

func (s SubscriptionPlanId) GetId() string {
	return s.id
}

func (s SubscriptionPlanId) String() string {
	return s.id
}

func (s SubscriptionPlanId) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.id)
}

func (s SubscriptionPlanId) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.id)
}

func (s SubscriptionPlanId) validate() error {
	if len(s.id) == 0 {
		return errors.New("subscription plan id cannot be empty")
	} else if len(s.id) > 256 {
		return errors.New("subscription plan id cannot have more than 256 characters")
	}
	return nil
}
