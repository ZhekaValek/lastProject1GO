package TimeMultiplyer

import "time"

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
