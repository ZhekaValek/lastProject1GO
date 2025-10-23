package WeatherParametresCalculator

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	coef := 1.0
	if weather.Condition == HeavyRain {
		coef += 0.2
	} else if weather.Condition == Snow {
		coef += 0.15
	} else if weather.Condition == Rain {
		coef += 0.125
	}
	if weather.WindSpeed > 15 {
		coef += 0.1
	}
	return coef
}
