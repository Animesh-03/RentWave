# Project Structure

## Shell Scripts

The shell scripts are just variations of docker compose commands to acheive variaton in the number of brokers in the kafka cluster.

To run
```sh
./kafka-{value}.sh
```

## Consumer

This is the consumer of kafka events.

To run
```sh
# Only needed if code is changed
npx tsc
node build/src/app.js
```

You can run multiple consumers on different terminals and the partitions will be allotted automatically

## Dockerfiles

Contains the dockerfiles and the various overrides needed for variation like running one or multiple brokers.

`kafka-base.yaml` runs one zookeeper and one broker along with the UI exposed on port `8080`

All the overrides are built on top of this and are used to run multiple nodes.

## Kafka-docker

This contains the setup and the dockerfile needed to build a kafka broker.

All the brokers are built using the doockerfile in this folder.

## Producer

This is used to generate the kafka events. Currently generates 4 events with different keys and values.

To run
```sh
# Only needed if code is changed
npx tsc
node build/src/app.js
```