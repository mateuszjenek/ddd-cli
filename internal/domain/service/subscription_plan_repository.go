package service

import (
	"context"

	"github.com/mateuszjenek/ddd-cli/internal/domain/model"
	"github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"
)

type SubscriptionPlanRepository interface {
	Create(context.Context, model.SubscriptionPlan) (model.SubscriptionPlan, error)
	Get(context.Context, valueobject.SubscriptionPlanId) (model.SubscriptionPlan, error)
	List(context.Context) ([]model.SubscriptionPlan, error)
	Update(context.Context, model.SubscriptionPlan) (model.SubscriptionPlan, error)
	Withdraw(context.Context, valueobject.SubscriptionPlanId) (model.SubscriptionPlan, error)
}