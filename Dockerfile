# Используем официальный образ Golang
FROM golang:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы в текущую директорию
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/main.go

# Экспортируем переменные окружения
ENV HTTP_PORT=:8080
ENV HTTP_HOST=localhost
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_DRIVER=postgres
ENV SSL_MODE=disable
ENV DSN="postgresql://postgres:123123@postgres:5432/postgres?sslmode=disable"


# Открываем порт 8080
EXPOSE 8080
EXPOSE 5432

# Запускаем приложение
CMD ["./main"]