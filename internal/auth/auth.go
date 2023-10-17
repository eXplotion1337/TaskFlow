package auth

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword сравнивает введенный пароль с хешированным паролем
func CheckPassword(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}

// GenerateToken генерирует JWT токен на основе имени пользователя
func GenerateToken(username string) (string, error) {
	// Создаем новый токен с именем пользователя в качестве идентификатора
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Токен действителен 24 часа
	})

	// Подписываем токен с использованием секретного ключа
	secretKey := []byte("fgdsbjbsdfjgb#$#$#425243524352JJGJKGJGJ") // Замени на свой секретный ключ
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
