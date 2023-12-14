package service

import (
	"context"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
)

type SubscriptionRepository interface {
	Create(context.Context, model.Subscription) (model.Subscription, error)
	Get(context.Context, valueobject.SubscriptionId) (model.Subscription, error)
	List(context.Context) ([]model.Subscription, error)
	
	Expire(context.Context, valueobject.SubscriptionId) (model.Subscription, error)

	// Update(context.Context, model.Subscription) (model.Subscription, error)
	// Remove(context.Context, valueobject.SubscriptionId) error
}