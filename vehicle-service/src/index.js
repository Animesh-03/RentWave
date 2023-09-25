import express from "express";
import mongoose from "mongoose";

const app = express();
const port = 3000;
const uri =
  "mongodb+srv://admin:uTSWnw9VC3LTydfG@microservicesproject.u5sbgzf.mongodb.net/?retryWrites=true&w=majority";

app.use(express.json());

mongoose.connect(uri);

const VehicleListing = mongoose.model("Vehicle_Listing", {
  uid: { type: Number },
  name: { type: String },
  type: { type: String },
  desc: { type: String },
  price: { type: Number },
  active: { type: Boolean },
});

app.post("/upload_listing", (req, res) => {
  var new_listing = new VehicleListing({
    uid: req.body.uid,
    name: req.body.name,
    type: req.body.type,
    desc: req.body.desc,
    price: req.body.price,
    active: req.body.active,
  });

  new_listing.save();
  res.end(new_listing._id.toString());
});

app.patch("/toggle_active", async (req, res) => {
  const activeness = await VehicleListing.findById(req.body._id).active;
  await VehicleListing.findByIdAndUpdate(req.body._id, { active: !activeness });
  res.end();
});

app.get("/active_listings", async (req, res) => {
  const listings = await VehicleListing.find({ active: true });
  res.end(listings.toString());
});

app.get("/user_listings", async (req, res) => {
  const listings = await VehicleListing.find({ uid: req.query.uid });
  res.end(listings.toString());
});

app.listen(port);