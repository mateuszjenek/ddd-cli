package application

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
	"github.com/mateuszjenek/ddd-cli/internal/domain/service"
)

type RegisterNewCustomerInputPort interface {
	RegisterNewCustomer(firstName, lastName, email string) (*model.Customer, error)
}

type registerNewCustomerUseCase struct {
	customers service.CustomerRepository
}

func NewRegisterNewCustomer(customers service.CustomerRepository) RegisterNewCustomerInputPort {
	return &registerNewCustomerUseCase{customers}
}

func (r *registerNewCustomerUseCase) RegisterNewCustomer(firstName string, lastName string, email string) (*model.Customer, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to create an uuid: %w", err)
	}
	
	customerIdValue, err := valueobject.NewCustomerId(id.String())
	if err != nil {
		return nil, fmt.Errorf("failed to create customer id: %w", err)
	}

	fullNameValue, err := valueobject.NewFullName(firstName, lastName)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer full name: %w", err)
	}

	emailValue, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer email: %w", err)
	}

	customer := model.Customer{
		Id:       customerIdValue,
		FullName: fullNameValue,
		Email:    emailValue,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	customer, err = r.customers.Create(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("failed to create a customer: %w", err)
	}

	return &customer, err
}
