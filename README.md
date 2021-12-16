# A Small API

This simple service's purpose is to expose data regarding a vehicle fleet.
The fleet consists of several vehichle, of several brands and series / models.

## Endpoints
In its current state, the API exposes 3 endpoints : 
- One to get the whole fleet
- One to get one vehicle by its id
- One to get metrics related to one car brand within the fleet

## Objectives

- Implement the service responsible for retrieving a vehicle by its ID
- Implement the service responsible for retrieving metrics related to a specific brand
- Build the docker image of the API

## Some help

- Run the component locally : ```go run main.go```

- By default the server exposes :8080 port

- For requesting the fleet ```curl -i http://localhost:8080/getfleet```
- For requesting one vehicle ```curl -i http://localhost:8080/getvehicle/<vehicle_id>```
- For requesting metrics from a brand ```curl -i -X POST http://localhost:8080/metrics -d "{\"brand\":\"<car_brand>\"}"```

These are examples, you can of course use a more elaborate http client :)