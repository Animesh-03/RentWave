import express from "express";
import cors from "cors"
import cookieParser from 'cookie-parser';

import userRouter from "./routers/user/user";
import vehicleRouter from "./routers/vehicle/vehicle";

const app = express();

app.use(express.json());
app.use(cors());
app.use(cookieParser());

app.use("/user", userRouter);
app.use("/vehicle",vehicleRouter);

app.get("/", (_, res) => {
    res.send("Hello World");
})

app.listen(8082, () => console.log("Started API Gateway on PORT 8082"));