package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FullName struct {
	firstName string
	lastName string
}

func NewFullName(firstName, lastName string) (FullName, error) {
	a := FullName{firstName, lastName}
	if err := a.validate(); err != nil {
		return a, fmt.Errorf("failed to validate full name: %w", err)
	}
	return a, nil
}

func (f FullName) GetFirstName() string {
	return f.firstName
}

func (f FullName) GetLastName() string {
	return f.lastName
}

func (f FullName) String() string {
	return fmt.Sprintf("%s %s", f.firstName, f.lastName)
}

func (f FullName) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
	}{
		FirstName: f.firstName,
		LastName: f.lastName,
	})
}

func (f FullName) UnmarshalJSON(data []byte) error {
	var unmarshaled struct{
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
	}
	err := json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return fmt.Errorf("failed to unmarshal full name: %w", err)
	}
	f.firstName = unmarshaled.FirstName
	f.lastName = unmarshaled.LastName
	return nil
}

func (f FullName) validate() error {
	if len(f.firstName) == 0 {
		return errors.New("first name cannot be empty")
	} else if len(f.firstName) > 128 {
		return errors.New("first name cannot have more than 256 characters")
	}

	if len(f.lastName) == 0 {
		return errors.New("last name cannot be empty")
	} else if len(f.lastName) > 128 {
		return errors.New("last name cannot have more than 256 characters")
	}
	return nil
}
