nano ~/.bashrc

export PATH=$PATH:/usr/local/go/bin

source ~/.bashrc


Запуск Кафки

kafka-server-start.sh /usr/local/kafka/kafka_2.13-3.8.0/config/kraft/server.properties

java -jar /usr/local/kafka/kafdrop/target/kafdrop-4.0.3-SNAPSHOT.jar --kafka.brokerConnect=localhost:9092

Консольный потребитель

kafka-console-producer.sh --topic test-topic --bootstrap-server localhost:9092


Остановить Кафку

kafka-server-stop.sh


Просмотр логов

tail -f /usr/local/kafka/kafka_2.13-3.8.0/logs/kafka.log

Инициализация нового Maven проект:

mvn archetype:generate -DgroupId=com.example.kafka -DartifactId=my-connector -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false



