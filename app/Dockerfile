# Используйте официальный образ Golang
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем только файлы с зависимостями, чтобы ускорить сборку
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код приложения в контейнер
COPY . .
COPY . /app/.env

# Собираем Go-приложение
RUN go build -o myapp

# Команда для запуска приложения при старте контейнера
CMD ["./myapp"]
