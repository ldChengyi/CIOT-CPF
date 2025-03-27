package mapper

import (
	"github.com/ldChengyi/CIOT-CPF/internal/domain/entity"
	"github.com/ldChengyi/CIOT-CPF/internal/pkg/db"
)

func GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func CreateUser(user *entity.User) error {
	return db.DB.Create(user).Error
}

func UpdateUser(user *entity.User) error {
	return db.DB.Save(user).Error
}
