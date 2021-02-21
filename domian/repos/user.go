package repos

import (
	"github.com/google/uuid"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
)

type UserRepository interface {
	CreateOrUpdate(*models.User) error
	FirstOrCreate(*models.User) error
	Create(*models.User) error
	FindOrFail(interface{}) (models.User, bool)
	Get(interface{}, ...interface{}) (models.User, error)
	GetAll() ([]models.User, error)
	Paginate(*helpers.Param) ([]models.User, *helpers.Paginator, error)
	Update(*models.User, uuid.UUID) error
	UpdateWhere(*models.User, interface{}, ...interface{}) error
	Delete(uuid.UUID) error
}
