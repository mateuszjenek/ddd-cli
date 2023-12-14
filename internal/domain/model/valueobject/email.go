package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

const emailRegex = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`

type Email struct {
	email string
}

func NewEmail(email string) (Email, error) {
	i := Email{email}
	if err := i.validate(); err != nil {
		return i, fmt.Errorf("failed to validate email: %w", err)
	}
	return i, nil
}

func (e Email) GetEmail() string {
	return e.email
}

func (e Email) String() string {
	return e.email
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.email)
}

func (e Email) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.email)
}

func (e Email) validate() error {
	if len(e.email) == 0 {
		return errors.New("email cannot be empty")
	} else if len(e.email) > 256 {
		return errors.New("email cannot have more than 256 characters")
	}

	regex, err := regexp.Compile(emailRegex)
	if err != nil {
		return fmt.Errorf("failed to compile email regex: %w", err)
	}
	if !regex.Match([]byte(e.email)) {
		return errors.New("email do not match standard email format")
	}
	return nil
}
