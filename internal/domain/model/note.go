package model

import "github.com/mateuszjenek/ddd-cli/internal/domain/model/valueobject"

type Note struct {
	Id      valueobject.Id
	Author  valueobject.Author
	Message valueobject.Message
}
