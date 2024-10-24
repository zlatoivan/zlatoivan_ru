package color

import (
	"fmt"
	"strconv"
)

// Тускловатые

func NRed(s string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", s)
}

func NGreen(s string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", s)
}

func NYellow(s string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", s)
}

func NBlue(s string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", s)
}

func NMagenta(s string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", s)
}

func NCyan(s string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", s)
}

// Яркие

func Red(s string) string {
	return fmt.Sprintf("\x1b[91m%s\x1b[0m", s)
}

func Green(s string) string {
	return fmt.Sprintf("\x1b[92m%s\x1b[0m", s)
}

func Yellow(s string) string {
	return fmt.Sprintf("\x1b[93m%s\x1b[0m", s)
}

func Blue(s string) string {
	return fmt.Sprintf("\x1b[94m%s\x1b[0m", s)
}

func Magenta(s string) string {
	return fmt.Sprintf("\x1b[95m%s\x1b[0m", s)
}

func Cyan(s string) string {
	return fmt.Sprintf("\x1b[96m%s\x1b[0m", s)
}

func GetColoredStatus(status int) string {
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
