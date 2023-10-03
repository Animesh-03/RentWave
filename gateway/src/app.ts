import express from "express";
import cors from "cors"
import cookieParser from 'cookie-parser';

import userRouter from "./routers/user/user";
import vehicleRouter from "./routers/vehicle/vehicle";
import bookRouter from "./routers/book/book";

const app = express();

app.use(express.json());
app.use(cors({
    origin: "http://localhost:8080",
    credentials: true,
    exposedHeaders: ["set-cookie"],
}));
app.use(cookieParser());

app.use("/user", userRouter);
app.use("/vehicle",vehicleRouter);
app.use("/book", bookRouter)

app.get("/", (_, res) => {
    res.send("Hello World");
})

app.listen(8082, () => console.log("Started API Gateway on PORT 8082"));