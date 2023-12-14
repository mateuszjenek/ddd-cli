package service

import (
	"context"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
)

type CustomerRepository interface {
	Create(context.Context, model.Customer) (model.Customer, error)
	Get(context.Context, valueobject.CustomerId) (model.Customer, error)
	List(context.Context) ([]model.Customer, error)
	Update(context.Context, model.Customer) (model.Customer, error)
	Remove(context.Context, valueobject.CustomerId) error
}