package model

import "github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"

type SubscriptionPlan struct {
	Id valueobject.SubscriptionPlanId
	Name valueobject.SubscriptionPlanName
}