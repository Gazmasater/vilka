nano ~/.bashrc

export PATH=$PATH:/usr/local/go/bin

source ~/.bashrc


Запуск в фоновом режиме

kafka-server-start.sh /usr/local/kafka/kafka_2.13-3.8.0/config/kraft/server.properties > kafka.log 2>&1 &

Остановить Кафку

kafka-server-stop.sh


Просмотр логов

tail -f kafka.log
