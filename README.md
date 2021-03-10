The repository contains 2 golang application:
- `port-client-api` provides REST API to accept a list of items in JSON format, which is forwarded to another service via gRPC.
- `port-domain-service` listens on gRPC endpoint to receive individual items and store them in a data layer.

To start the application, run the following commands from the root folder:
```
docker-compose build
```
and
```
docker-compose up
```

This makes a REST endpoint available on `localhost:8080`, to send items for processing.

To invoke the endpoint, the following curl command can be used:
```
curl -X POST 'localhost:8080/ports' -H 'Content-Type: application/json' --data-raw '{}'
```
where `--data-raw '{}'` contains the json payload of the item list. 
