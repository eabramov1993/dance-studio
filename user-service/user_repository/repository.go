package user_repository

import (
	"database/sql"
	"fmt"

	"github.com/eabramov1993/dance-studio/tree/main/user-service/models"

	"golang.org/x/crypto/bcrypt"
)

// Функция для создания нового пользователясв юю
func CreateUser(db *sql.DB, user *models.User) error {
	// Хэшируем пароль перед сохранением в базу данных
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка при хэшировании пароля: %v", err)
	}

	// Выполняем SQL-запрос для вставки данных в таблицу users
	_, err = db.Exec("INSERT INTO users (name, email, password, role, contact_info) VALUES ($1, $2, $3, $4, $5)",
		user.Name, user.Email, hashedPassword, user.Role, user.ContactInfo)
	if err != nil {
		return fmt.Errorf("ошибка при вставке пользователя в базу данных: %v", err)
	}

	return nil
}
