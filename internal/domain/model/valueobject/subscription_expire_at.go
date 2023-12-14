package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SubscriptionExpireAt struct {
	time time.Time
}

func NewSubscriptionExpireAt(time time.Time) (SubscriptionExpireAt, error) {
	i := SubscriptionExpireAt{time}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("failed to validate subscription id: %w", err)
	}
	return i, nil
}

func (s SubscriptionExpireAt) GetTime() time.Time {
	return s.time
}

func (s SubscriptionExpireAt) String() string {
	return s.time.String()
}

func (s SubscriptionExpireAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.time)
}

func (s SubscriptionExpireAt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.time)
}

func (s SubscriptionExpireAt) validate() error {
	if time.Now().After(s.time) {
		return errors.New("expire at cannot be earlier than now")
	}
	return nil
}
