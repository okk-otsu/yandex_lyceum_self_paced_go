package main

import (
	"fmt"
	"math"
	"time"
)

type WeatherCondition int

const (
	minPrice                  = 99.0
	maxPrice                  = 20_000.0
	Clear    WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
	Rain
	HeavyRain
	Snow
)

func ApplyPriceLimits(price float64) float64 {
	return min(max(price, minPrice), maxPrice)
}

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	res := 1.0
	if weather.WindSpeed > 15 {
		res += 0.1
	}
	switch weather.Condition {
	case Rain:
		res += 0.125
	case HeavyRain:
		res += 0.2
	case Snow:
		res += 0.15
	default:
	}

	return res
}

var pricePerKm = 10.0
var pricePerMinute = 2.0

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(tp TripParameters) float64 {
	return tp.Distance*pricePerKm + tp.Duration*pricePerMinute
}

type TrafficClient interface {
	GetTrafficLevel(lat, lng float64) int // 1–5
}

func GetTrafficMultiplier(trafficLevel int) float64 {
	return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
	TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
	return 3 // Константное значение в нашем примере, в реальности оно будет вычисляться сервисом Яндекс Карты
}

func GetTimeMultiplier(t time.Time) float64 {
	hour := t.Hour()
	isWeekend := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday // Проверка, что сегодня суббота или воскресенье (выходные)

	switch {
	case hour >= 0 && hour < 5:
		return 1.5 // Ночной тариф
	case hour >= 7 && hour < 10 && !isWeekend:
		return 1.3 // Утренний час пик
	case isWeekend:
		return 1.2 // Выходные
	default:
		return 1.0
	}
}

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}

func main() {
	calculator := PriceCalculator{
		TrafficClient: &RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		WeatherData{HeavyRain, 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
