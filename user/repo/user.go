package repo

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/your/repo/user/model"
)

type UserRepo struct {
	db  *gorm.DB
	log *log.Logger
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db:  db,
		log: log.New(os.Stdout, "user-repo: ", log.LstdFlags),
	}
}

func (r *UserRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("phone = ?", phone).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindByID(id uint) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id uint) error {
	user := &model.User{}
	user.ID = id
	return r.db.Delete(user).Error
}

func (r *UserRepo) FindAll() ([]*model.User, error) {
	users := []*model.User{}
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
