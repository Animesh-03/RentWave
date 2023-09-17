import express from "express";
import cors from "cors"
import userRouter from "./routers/user/user";

const app = express();

app.use(express.json());
app.use(cors());

app.use("/user", userRouter);

app.get("/", (_, res) => {
    res.send("Hello World");
})

app.listen(8082, () => console.log("Started API Gateway on PORT 8082"));