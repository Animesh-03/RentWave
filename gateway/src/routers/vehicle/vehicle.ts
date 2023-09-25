import axios, { AxiosError } from "axios";
import { Router } from "express";
import jwt from "jsonwebtoken"
import { vehicleServiceBaseUrl } from "../../constants";

const vehicleRouter = Router();

vehicleRouter.all("/", (req, res) => res.send("Hello World"));

vehicleRouter.post("/add", async (req, res, next) => {
    let authCookie: jwt.JwtPayload;
    try {
        authCookie = jwt.verify(req.cookies.auth, 'secret') as jwt.JwtPayload;
    }
    catch(e) {
        res.status(401).send({
            "error": "Invalid user credentials"
        });
        next();
        return;
    }

    const response = await axios.post(`${vehicleServiceBaseUrl}/upload_listing`, {
        uid: authCookie.uid,
        name: req.body.name,
        type: req.body.type,
        desc: req.body.desc,
        price: req.body.price,
        active: req.body.active,
    })
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            console.log(err);
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
        next();
    });

    if(response)
    {
        res.send(response.data);
    }
});

vehicleRouter.get("/list/active", async (req,res, next) => {
    const response = await axios.get(`${vehicleServiceBaseUrl}/active_listings`)
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            console.log(err);
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
        next();
    });

    if(response)
    {
        res.send(response.data);
    }
});

vehicleRouter.patch("/toggle", async (req, res, next) => {
    let authCookie: jwt.JwtPayload;
    try {
        authCookie = jwt.verify(req.cookies.auth, 'secret') as jwt.JwtPayload;
    }
    catch(e) {
        res.status(401).send({
            "error": "Invalid user credentials"
        });
        next();
        return;
    }

    const response = await axios.patch(`${vehicleServiceBaseUrl}/toggle_active`, {
        "_id": authCookie.uid
    })
    .catch(function (err : Error | AxiosError) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            console.log(err);
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
        next();
    });

    if(response) 
    {
        // console.log(response.status)
        res.sendStatus(200);
    }
});

vehicleRouter.get("/me", async (req, res, next) => {
    let authCookie: jwt.JwtPayload;
    try {
        authCookie = jwt.verify(req.cookies.auth, 'secret') as jwt.JwtPayload;
    }
    catch(e) {
        res.status(401).send({
            "error": "Invalid user credentials"
        });
        next();
        return;
    }

    const vehicleList = await axios.get(`${vehicleServiceBaseUrl}/user_listings`, {
        params: {
            "uid": authCookie.uid
        }
    });

    res.send(vehicleList.data)
})


export default vehicleRouter;