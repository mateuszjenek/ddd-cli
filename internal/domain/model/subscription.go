package model

import "github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"

type Subscription struct {
	Id valueobject.SubscriptionId
	ExpireAt valueobject.SubscriptionExpireAt

	SubscriptionPlanId valueobject.SubscriptionPlanId
	CustomerId valueobject.CustomerId
}