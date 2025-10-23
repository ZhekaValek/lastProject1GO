package main

import (
	"fmt"
	. "lastProject1/TrafficMultiplier"
	. "lastProject1/TripParametresCalculator"
	. "lastProject1/WeatherParametresCalculator"
	"time"
)

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
