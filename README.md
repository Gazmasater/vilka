nano ~/.bashrc

export PATH=$PATH:/usr/local/go/bin

source ~/.bashrc


Запуск в фоновом режиме

kafka-server-start.sh /usr/local/kafka/kafka_2.13-3.8.0/config/kraft/server.properties > kafka.log 2>&1 &

java -jar target/kafdrop-4.0.3-SNAPSHOT.jar --kafka.brokerConnect=localhost:9092


Консольный потребитель

kafka-console-producer.sh --topic test-topic --bootstrap-server localhost:9092


Остановить Кафку

kafka-server-stop.sh


Просмотр логов

tail -f /usr/local/kafka/kafka_2.13-3.8.0/logs/kafka.log


