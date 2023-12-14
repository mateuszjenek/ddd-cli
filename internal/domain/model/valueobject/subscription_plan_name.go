package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SubscriptionPlanName struct {
	name string
}

func NewSubscriptionPlanName(name string) (SubscriptionPlanName, error) {
	i := SubscriptionPlanName{name}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("failed to validate subscription plan name: %w", err)
	}
	return i, nil
}

func (s SubscriptionPlanName) GetId() string {
	return s.name
}

func (s SubscriptionPlanName) String() string {
	return s.name
}

func (s SubscriptionPlanName) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.name)
}

func (s SubscriptionPlanName) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.name)
}

func (s SubscriptionPlanName) validate() error {
	if len(s.name) == 0 {
		return errors.New("subscription plan name cannot be empty")
	} else if len(s.name) > 256 {
		return errors.New("subscription plan name cannot have more than 256 characters")
	}
	return nil
}
