package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
	// "github.com/lib/pq"
)

func main() {
    // Загрузите переменные окружения из .env файла
    if err := godotenv.Load(); err != nil {
        fmt.Println("Ошибка при загрузке .env файла:", err)
        return
    }

    // Подключение к Kafka
    broker := os.Getenv("KAFKA_BROKER")
    producer, err := sarama.NewSyncProducer([]string{broker}, nil)
    if err != nil {
        fmt.Println("Ошибка при создании продюсера Kafka:", err)
        return
    }
    defer producer.Close()

    topic := "my-topic"
    message := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder("Hello, Kafka!"),
    }
    _, _, err = producer.SendMessage(message)
    if err != nil {
        fmt.Println("Ошибка при отправке сообщения в Kafka:", err)
        return
    }
    fmt.Println("Сообщение успешно отправлено в Kafka.")

    // Подключение к PostgreSQL
    dbURL := os.Getenv("DATABASE_URL")
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        fmt.Println("Ошибка при подключении к PostgreSQL:", err)
        return
    }
    defer db.Close()

    // Создание таблицы
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id serial primary key, content text)")
    if err != nil {
        fmt.Println("Ошибка при создании таблицы:", err)
        return
    }
    fmt.Println("Таблица 'messages' создана или уже существует.")

    // Вставка данных в PostgreSQL
    content := "Hello, PostgreSQL!"
    _, err = db.Exec("INSERT INTO messages (content) VALUES ($1)", content)
    if err != nil {
        fmt.Println("Ошибка при вставке данных в PostgreSQL:", err)
        return
    }
    fmt.Println("Данные успешно вставлены в PostgreSQL.")

    // Чтение данных из PostgreSQL
    rows, err := db.Query("SELECT content FROM messages")
    if err != nil {
        fmt.Println("Ошибка при чтении данных из PostgreSQL:", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var message string
        if err := rows.Scan(&message); err != nil {
            fmt.Println("Ошибка при сканировании строки:", err)
            return
        }
        fmt.Printf("Сообщение из PostgreSQL: %s\n", message)
    }
}
