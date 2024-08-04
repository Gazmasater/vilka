CLUSTER_ID=$(uuidgen)

echo $CLUSTER_ID

bin/kafka-storage.sh format -t $CLUSTER_ID -c config/kraft/server.properties

Топики

1xbet



/etc/kafka/server.properties

/etc/kafka/kraft/server.properties

log.dirs=/tmp/kraft-combined-logs

confluent-hub install confluentinc/kafka-connect-http-source:latest

Версия Kafka 

kafka-topics --version



