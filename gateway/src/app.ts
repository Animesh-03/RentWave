import express from "express";
import cookieParser from 'cookie-parser';

import userRouter from "./routers/user/user";
import vehicleRouter from "./routers/vehicle/vehicle";
import bookRouter from "./routers/book/book";

const app = express();

app.use(express.json());

app.all("*", (req, res, next) => {
    const origin = req.header('origin');
    res.header("Access-Control-Allow-Origin", origin ?? "*");
    res.header("Access-Control-Allow-Credentials", "true");
    res.header("Access-Control-Expose-Headers", "set-cookie");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
})

app.use(cookieParser());



app.use("/user", userRouter);
app.use("/vehicle",vehicleRouter);
app.use("/book", bookRouter)



app.get("/", (_, res) => {
    res.send("Hello World");
})

app.listen(8082, () => console.log("Started API Gateway on PORT 8082"));