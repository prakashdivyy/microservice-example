const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");
const asyncHandler = require("express-async-handler");
const express = require("express");
const bodyParser = require("body-parser");
const amqp = require("amqplib/callback_api");

const app = express();
app.use(bodyParser.json());
const host = process.env.APP_HOST || "localhost";
const port = process.env.APP_PORT ? parseInt(process.env.APP_PORT, 10) : 3000;

const protoLoc = process.env.PROTO_LOCATION || "../proto/products.proto";
const productServiceHost = process.env.PRODUCT_SERVICE_HOST || "localhost";
const productServicePort = process.env.PRODUCT_SERVICE_PORT || "50050";
const rabbitHost = process.env.RABBIT_HOST || "localhost";

// Load the protobuf
const proto = grpc.loadPackageDefinition(
  protoLoader.loadSync(protoLoc, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  })
);

// Create a new client instance that binds to the IP and port of the grpc server.
const client = new proto.rpc.ProductService(
  `${productServiceHost}:${productServicePort}`,
  grpc.credentials.createInsecure()
);

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.post("/login", (req, res) => {
  // TODO: Check email and password on auth-service
  const { email, password } = req.body;

  amqp.connect(`amqp://${rabbitHost}`, function (error0, connection) {
    if (error0) {
      throw error0;
    }
    connection.createChannel(function (error1, channel) {
      if (error1) {
        throw error1;
      }
      var queue = "notification";
      var msg = JSON.stringify({ name: "Name", email });

      channel.assertQueue(queue, {
        durable: false,
      });
      channel.sendToQueue(queue, Buffer.from(msg));

      console.log(" [x] Sent %s", msg);
    });
    setTimeout(function () {
      connection.close();
    }, 500);
  });

  res.send({ status: "ok" });
});

app.get("/products/:id", (req, res) => {
  const id = parseInt(req.params.id, 10);
  console.log(id);
  client.getProduct({ id }, (error, response) => {
    if (!error) {
      if (response.id === 0) {
        res.sendStatus(404);
      } else {
        res.send(response);
      }
    } else {
      res.sendStatus(500);
    }
  });
});

app.listen(port, host, () => {
  console.log(`gateway-service listening at http://${host}:${port}`);
});
