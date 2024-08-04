package kafka

import (
	"context"
	"time"
	"vilka/internal/models"
	"vilka/pkg/logger"

	"github.com/IBM/sarama"
)

func StartConsumer(ctx context.Context, stop <-chan struct{}) {
	brokerList := []string{"localhost:9092"}

	// Создание контекста с таймаутом для подключения к Kafka
	connectCtx, cancel := context.WithTimeout(ctx, models.ConnectionTimeoutDuration)
	defer cancel()

	// Проверяем подключение к Kafka
	if err := ConnectToKafka(connectCtx, brokerList); err != nil {
		logger.GetLogger().Fatalf("Не удалось подключиться к Kafka: %v", err)
	}

	// Дальнейшая логика вашего кода может быть добавлена здесь...
}

// ConnectToKafka проверяет подключение к Kafka
func ConnectToKafka(ctx context.Context, brokerList []string) error {
	config := sarama.NewConfig()
	var client sarama.Client
	var err error

	// Получаем экземпляр логгера
	sugar := logger.GetLogger()

	// Попытка подключения в течение времени, заданного в контексте
	for {
		select {
		case <-ctx.Done():
			return ctx.Err() // Возвращаем ошибку контекста, если он завершён
		default:
			// Пытаемся создать нового клиента Kafka
			client, err = sarama.NewClient(brokerList, config)
			if err == nil {
				// Успешное подключение, закрываем клиента и возвращаем nil
				client.Close()
				sugar.Info("Успешно подключено к Kafka")
				return nil
			}
			sugar.Errorf("Не удалось подключиться к Kafka: %v", err)
			time.Sleep(2 * time.Second) // Ожидание 2 секунды перед следующей попыткой
		}
	}
}
