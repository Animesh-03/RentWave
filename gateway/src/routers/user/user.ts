import axios, { AxiosError } from "axios";
import { Router } from "express";
import { userServiceBaseUrl } from "../../constants";

const userRouter = Router();

userRouter.post("/register", async (req, res, next) => {
    // Redirect to User Service
    const registerResponse = await axios.post(`${userServiceBaseUrl}/register`, {
        "username": req.body.username,
        "password": req.body.password
    })
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            res.status(err.response?.status ?? 500).send(err.response?.data);
            next();
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
            next();
        }
    })


    if(registerResponse)
        res.sendStatus(200);
});

userRouter.post("/login", async (req, res, next) => {
    // Redirect to User Service
    const loginResponse = await axios.post(`${userServiceBaseUrl}/login`, {
        "username": req.body.username,
        "password": req.body.password
    })
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
        next();
    });

    if(loginResponse) {
        // If request is sucessful then set the auth cookie while sending response
        res.cookie("auth", loginResponse.data.access_token, {
            // Expire after 24hrs
            expires: new Date(Date.now() + 24*60*60*1000)
        });
        res.sendStatus(200);
    }

});


export default userRouter;