# Start one zookeeper and two broker along with UI for Kafka at localhost:8080
docker compose -f dockerfiles/kafka.base.yml -f dockerfiles/overrides/kafka.two.yml up --build