package my_middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header)
		// Получаем токен из заголовков или cookies
		if strings.HasPrefix(r.URL.Path, "/scripts/") || strings.HasPrefix(r.URL.Path, "/styles/") || strings.HasPrefix(r.URL.Path, "/web/") || strings.HasPrefix(r.URL.Path, "/ping/") {
			next.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/auth/sing-up" || r.URL.Path == "/auth/sing-in" || r.URL.Path == "/" || r.URL.Path == "/sing-in" || r.URL.Path == "/sing-up" {
			// Если это маршрут для регистрации, пропускаем middleware
			next.ServeHTTP(w, r)
			return
		}
		token := extractToken(r)
		fmt.Println(token)
		// Проверяем токен
		if isValidToken(token) {
			// Если токен валиден, передаем управление следующему обработчику
			next.ServeHTTP(w, r)
		} else {
			// Если токен недействителен, отправляем ошибку Unauthorized
			//w.WriteHeader(http.StatusUnauthorized)
			//w.Write([]byte("Unauthorized"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	})
}

// Вспомогательная функция для извлечения токена из заголовков или cookies
func extractToken(r *http.Request) string {
	// Логика извлечения токена
	// Получаем значение заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Разделяем строку заголовка по пробелу
	// Формат: Bearer <token>
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ""
	}

	// Возвращаем токен
	return tokenParts[1]
}

// Вспомогательная функция для проверки валидности токена
func isValidToken(tokenStr string) bool {
	// Логика проверки валидности токена
	// Замени "your-secret-key" на свой секретный ключ, использованный при подписи токена
	secretKey := []byte("fgdsbjbsdfjgb#$#$#425243524352JJGJKGJGJ")

	// Пытаемся распарсить токен с использованием секретного ключа
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Проверяем ошибки
	if err != nil || !token.Valid {
		return false
	}

	return true
}
