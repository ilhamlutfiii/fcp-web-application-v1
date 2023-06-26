package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var userTaskCategories []model.UserTaskCategory
	err := r.db.Table("users").
		Select("users.id, users.fullname, users.email, users.password, tasks.title as task, tasks.deadline, tasks.priority, categories.name as category, tasks.status").
		Joins("JOIN tasks ON tasks.user_id = users.id").
		Joins("JOIN categories ON categories.id = tasks.category_id").
		Scan(&userTaskCategories).Error
	if err != nil {
		return nil, err
	}
	return userTaskCategories, nil
}
