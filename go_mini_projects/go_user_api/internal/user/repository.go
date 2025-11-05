package user

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(u *User) error {
	return r.DB.Create(u).Error

}

func (r *Repository) GetAll() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err

}
