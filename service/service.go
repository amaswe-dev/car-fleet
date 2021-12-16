package service

import (
	"amaswe-be/domain"
)

const (
	resource = "resource.json"
)

func GetFleet() ([]domain.Vehicle, error) {
	vehicles, err := domain.ParseInput(resource)
	return vehicles, err
}

/**
Write a service that fetch a vehicle by its uuid
*/
func GetVehicle(id string) (*domain.Vehicle, error) {
	_, _ = domain.ParseInput(resource)

	return nil, nil
}

/**
Write a service that gather metrics from the fleet
The metric must return the number of vehicles from a specific brand
Along with the overall vehicles, the metric must include the number of vehicle per model
e.g :
For instance, a request such as {"brand":"Polestar"} should return :
{"brand":"Volvo","overall":7,"models":{"C40":2,"S60":1,"S90":2,"V90":1,"XC60":1}}
*/
func GetMetric(body []byte) (*domain.Metrics, error) {
	_, _ = domain.ParseInput("resource.json")

	return &domain.Metrics{}, nil
}
