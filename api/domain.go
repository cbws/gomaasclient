package api

import (
	"github.com/maas/gomaasclient/entity"
)

type Domain interface {
	Get(id int) (*entity.Domain, error)
	SetDefault(id int) (*entity.Domain, error)
	Update(id int, params *entity.DomainParams) (*entity.Domain, error)
	Delete(id int) error
}
