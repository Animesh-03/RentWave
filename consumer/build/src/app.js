"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const kafkajs_1 = require("kafkajs");
const kafka = new kafkajs_1.Kafka({
    clientId: "consumer",
    brokers: ['localhost:9092', 'localhost:9093']
});
const consumer = kafka.consumer({ groupId: "group1" });
async function main() {
    await consumer.connect();
    await consumer.subscribe({ topic: 'test' });
    await consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
            console.log({
                partition,
                offset: message.offset,
                value: message.value.toString(),
            });
        },
    });
}
main();
//# sourceMappingURL=app.js.map