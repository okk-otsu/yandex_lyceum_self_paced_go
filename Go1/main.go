package main

import (
	"fmt"
	"reflect"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	var res = map[string]float64{}

	for _, tr := range vehicles {
		typeName := reflect.TypeOf(tr).String()
		res[typeName] = tr.CalculateTravelTime(distance)
	}

	return res
}

func main() {
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	vehicles := []Vehicle{car, motorcycle}

	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)

	fmt.Printf("Ожидается время для автомобиля %.2f часа\n", travelTimes["main.Car"])
	fmt.Printf("Ожидается время для мотоцикла %.2f часа", travelTimes["main.Motorcycle"])
}
