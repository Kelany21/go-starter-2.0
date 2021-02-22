package infrastructure

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/domian/repos"
	"golang-ddd-starter/helpers"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

//UserRepo implements the UserRepository interface
var _ repos.UserRepository = &UserRepo{}

func (r *UserRepo) CreateOrUpdate(user *models.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) FirstOrCreate(user *models.User) error {
	err := r.db.Where(user).FirstOrCreate(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Create(user *models.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) FindOrFail(id interface{}) (models.User, bool) {
	var user models.User
	err := r.db.Where("uuid = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, false
	}
	return user, true
}

func (r *UserRepo) Get(query interface{}, args ...interface{}) (models.User, error) {
	var user models.User
	err := r.db.Where(query, args...).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Paginate(p *helpers.Param) ([]models.User, *helpers.Paginator, error) {
	var users []models.User
	p.DB = DB
	paginator, err := Paging(p, &users)
	if err != nil {
		return nil, nil, err
	}
	paginator.RecordsCount = len(users)
	return users, paginator, nil
}

func (r *UserRepo) Update(user *models.User, id uuid.UUID) error {
	onlyAllowData := UpdateOnlyAllowColumns(user, models.UserFillAbleColumn())
	err := r.db.Model(&models.User{}).Where("uuid = ?", id).Updates(onlyAllowData).Error
	if err != nil {
		return err
	}
	err = r.db.Where("uuid = ?", id).Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateWhere(user *models.User, query interface{}, args ...interface{}) error {
	onlyAllowData := UpdateOnlyAllowColumns(user, models.UserFillAbleColumn())
	err := r.db.Model(&models.User{}).Where(query, args...).Updates(onlyAllowData).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Delete(id uuid.UUID) error {
	var user models.User
	err := r.db.Where("uuid = ?", id).Unscoped().Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
