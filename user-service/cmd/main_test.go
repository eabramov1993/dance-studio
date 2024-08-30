package main

import (
	"database/sql"
	"testing"

	"github.com/eabramov1993/dance-studio/tree/main/user-service/models"

	"github.com/eabramov1993/dance-studio/tree/main/user-service/user_repository"

	_ "github.com/lib/pq"
)

func TestCreateUser(t *testing.T) {
	// Настройка подключения к базе данных (замените на ваши данные)
	db, err := sql.Open("postgres", "postgres://postgres:Aspirine1@localhost:5432/dance_studio?sslmode=disable")
	if err != nil {
		t.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Создание тестового пользователя
	newUser := &models.User{
		Name:        "Тестовый Пользователь",
		Email:       "test@example.com",
		Password:    "testpassword",
		Role:        "родитель",
		ContactInfo: "1234567890",
	}

	// Вызов функции CreateUser
	err = user_repository.CreateUser(db, newUser)
	if err != nil {
		t.Fatalf("Ошибка при создании пользователя: %v", err)
	}

	// Проверка, что пользователь был создан
	var user models.User
	err = db.QueryRow("SELECT id, name, email, password, role, contact_info FROM users WHERE email = $1", newUser.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.ContactInfo)
	if err != nil {
		t.Fatalf("Ошибка при получении пользователя из базы данных: %v", err)
	}

	if user.Name != newUser.Name || user.Email != newUser.Email || user.Role != newUser.Role || user.ContactInfo != newUser.ContactInfo {
		t.Errorf("Данные созданного пользователя не совпадают с ожидаемыми")
	}
}
