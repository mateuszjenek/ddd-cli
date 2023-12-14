package model

import "github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"

type Customer struct {
	Id valueobject.CustomerId
	FullName valueobject.FullName
	Email valueobject.Email
}
