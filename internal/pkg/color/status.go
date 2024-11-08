package color

import "strconv"

// Status - возвращает код статуса в соответствующем цвете
func Status(status int) string {
	switch {
	case status < 200:
		return Blue(strconv.Itoa(status))
	case status < 300:
		return Green(strconv.Itoa(status))
	case status < 400:
		return Cyan(strconv.Itoa(status))
	case status < 500:
		return Yellow(strconv.Itoa(status))
	default:
		return Red(strconv.Itoa(status))
	}
}
