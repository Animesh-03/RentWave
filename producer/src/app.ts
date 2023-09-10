import {Kafka} from "kafkajs";

const kafka = new Kafka({
    clientId: "producer",
    brokers: ['localhost:9092','localhost:9093']
});

const producer = kafka.producer();

const main = async () => {
    await producer.connect();
    await producer.send({
        topic: "test",
        messages: [
            {
                // key: "m1",
                value: "Obj1"
            },
            {
                // key: "m2",
                value: "Obj2"
            },
            {
                // key: "m3",
                value: "Obj3"
            },
            {
                // key: "m1",
                value: "Obj4"
            }
        ]
    });
    console.log("Disconnecting....")
    await producer.disconnect();
}

main();