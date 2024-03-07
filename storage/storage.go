package storage

import (
	// "github.com/gofrs/uuid"
)

type DocumentActions interface {
	AddNews() error
	GetNews() error
	GetNewsByID() error
}
