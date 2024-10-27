package color

import (
	"fmt"
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

// Тускловатые

func BNRed(s string) string {
	return fmt.Sprintf("\x1b[31;1m%s\x1b[0m", s)
}

func BNGreen(s string) string {
	return fmt.Sprintf("\x1b[32;1m%s\x1b[0m", s)
}

func BNYellow(s string) string {
	return fmt.Sprintf("\x1b[33;1m%s\x1b[0m", s)
}

func BNBlue(s string) string {
	return fmt.Sprintf("\x1b[34;1m%s\x1b[0m", s)
}

func BNMagenta(s string) string {
	return fmt.Sprintf("\x1b[35;1m%s\x1b[0m", s)
}

func BNCyan(s string) string {
	return fmt.Sprintf("\x1b[36;1m%s\x1b[0m", s)
}

// Жирные яркие

func BRed(s string) string {
	return fmt.Sprintf("\x1b[91;1m%s\x1b[0m", s)
}

func BGreen(s string) string {
	return fmt.Sprintf("\x1b[92;1m%s\x1b[0m", s)
}

func BYellow(s string) string {
	return fmt.Sprintf("\x1b[93;1m%s\x1b[0m", s)
}

func BBlue(s string) string {
	return fmt.Sprintf("\x1b[94;1m%s\x1b[0m", s)
}

func BMagenta(s string) string {
	return fmt.Sprintf("\x1b[95;1m%s\x1b[0m", s)
}

func BCyan(s string) string {
	return fmt.Sprintf("\x1b[96;1m%s\x1b[0m", s)
}
