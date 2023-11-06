import axios, { AxiosError } from "axios";
import { Router } from "express";
import jwt from "jsonwebtoken"
import { bookServiceBaseUrl } from "../../constants";

const bookRouter = Router();

bookRouter.post("/add", async (req, res, next) => {
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

    const response = await axios.post(`${bookServiceBaseUrl}/add`, {
        isbn: req.body.isbn,
        title: req.body.title,
        author: req.body.author,
        image: req.body.image,
        description: req.body.description,
        ownerid: Number(authCookie.uid),
        price: Number(req.body.price),
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

bookRouter.get("/list/active", async (req, res, next) => {
    const response = await axios.get(`${bookServiceBaseUrl}/active_listings`)
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

bookRouter.patch("/toggle", async (req, res, next) => {
    const response = await axios.patch(`${bookServiceBaseUrl}/toggle_active`, {
        "_id": req.body._id
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
        res.sendStatus(200);
    }
});

bookRouter.get("/me", async (req, res, next) => {
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

    const vehicleList = await axios.get(`${bookServiceBaseUrl}/user_listings`, {
        params: {
            "uid": authCookie.uid
        }
    });

    res.send(vehicleList.data)
});

bookRouter.get("/search", async (req, res, next) => {
    try {
        const response = await axios.get(`${bookServiceBaseUrl}/search`, {
            params: req.query
        });
        res.send(response.data);
    }
    catch(err: unknown) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            // console.log(err);
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
    }
});

bookRouter.get("/isbn", async (req, res, next) => {
    try {
        const response = await axios.get(`${bookServiceBaseUrl}/isbn`, {
            params: req.query
        });
        res.send(response.data);
    }
    catch(err: unknown) {
        // Check for the type of error and handle accordingly
        if(axios.isAxiosError(err)) {
            // console.log(err);
            res.status(err.response?.status ?? 500).send(err.response?.data);
        }
        else {
            console.log("Error occured: ", err);
            res.status(500).send(err);
        }
    }
})

export default bookRouter
