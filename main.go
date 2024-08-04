package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"vilka/kafka"
	"vilka/pkg/logger"
)

func main() {
	// Инициализация логгера
	logger.Init()
	sugar := logger.GetLogger()

	// Создаем контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Канал для сигналов остановки
	stopKafka := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	// Запуск Kafka consumer в отдельной горутине
	go func() {
		defer wg.Done()
		kafka.StartConsumer(ctx, stopKafka) // Передаем контекст и сигнал остановки
	}()

	// Обработка корректного завершения работы
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Ожидание сигнала для корректного завершения
	<-stop
	sugar.Info("Остановка Kafka consumer...")

	// Отправляем сигнал остановки Kafka consumer
	close(stopKafka)

	// Отменяем контекст для Kafka consumer
	cancel()

	// Ждем завершения всех горутин
	wg.Wait()
	sugar.Info("Kafka consumer корректно остановлен")
}
