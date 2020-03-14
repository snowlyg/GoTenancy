package database

import (
	"errors"

	"github.com/snowlyg/go-tenancy/models"
)

type Engine uint32

const (
	Memory Engine = iota
	Gorm
)

func LoadUsers(engine Engine) (map[int64]models.User, error) {
	if engine == Gorm {
		return nil, errors.New("for the shake of simplicity we're using a simple map as the data source")
	}

	return make(map[int64]models.User), nil
}
