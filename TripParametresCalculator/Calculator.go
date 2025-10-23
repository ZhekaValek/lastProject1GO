package TripParametresCalculator

type TripParameters struct {
	Distance float64
	Duration float64
}

const (
	pricePerKm, pricePerMinute float64 = 10.0, 2.0
)

func CalculateBasePrice(t TripParameters) float64 {
	return t.Distance*pricePerKm + t.Duration*pricePerMinute
}
