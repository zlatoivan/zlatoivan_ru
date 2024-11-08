package color

import (
	"fmt"
)

// Тусклые

// NRed - тусклый красный
func NRed(s string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", s)
}

// NGreen - тусклый зеленый
func NGreen(s string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", s)
}

// NYellow - тусклый желтый
func NYellow(s string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", s)
}

// NBlue - тусклый голубой
func NBlue(s string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", s)
}

// NMagenta - тусклый фиолетовый
func NMagenta(s string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", s)
}

// NCyan - тусклый циан
func NCyan(s string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", s)
}

// Яркие

// Red - яркий красный
func Red(s string) string {
	return fmt.Sprintf("\x1b[91m%s\x1b[0m", s)
}

// Green - яркий зеленый
func Green(s string) string {
	return fmt.Sprintf("\x1b[92m%s\x1b[0m", s)
}

// Yellow - яркий желтый
func Yellow(s string) string {
	return fmt.Sprintf("\x1b[93m%s\x1b[0m", s)
}

// Blue - яркий голубой
func Blue(s string) string {
	return fmt.Sprintf("\x1b[94m%s\x1b[0m", s)
}

// Magenta - яркий фиолетовый
func Magenta(s string) string {
	return fmt.Sprintf("\x1b[95m%s\x1b[0m", s)
}

// Cyan - яркий циан
func Cyan(s string) string {
	return fmt.Sprintf("\x1b[96m%s\x1b[0m", s)
}

// Тусклые жирные

// BNRed - тусклый жирный красный
func BNRed(s string) string {
	return fmt.Sprintf("\x1b[31;1m%s\x1b[0m", s)
}

// BNGreen - тусклый жирный зеленый
func BNGreen(s string) string {
	return fmt.Sprintf("\x1b[32;1m%s\x1b[0m", s)
}

// BNYellow - тусклый жирный желтый
func BNYellow(s string) string {
	return fmt.Sprintf("\x1b[33;1m%s\x1b[0m", s)
}

// BNBlue - тусклый жирный голубой
func BNBlue(s string) string {
	return fmt.Sprintf("\x1b[34;1m%s\x1b[0m", s)
}

// BNMagenta - тусклый жирный фиолетовый
func BNMagenta(s string) string {
	return fmt.Sprintf("\x1b[35;1m%s\x1b[0m", s)
}

// BNCyan - тусклый жирный циан
func BNCyan(s string) string {
	return fmt.Sprintf("\x1b[36;1m%s\x1b[0m", s)
}

// Яркие жирные

// BRed - яркий жирный красный
func BRed(s string) string {
	return fmt.Sprintf("\x1b[91;1m%s\x1b[0m", s)
}

// BGreen - яркий жирный зеленый
func BGreen(s string) string {
	return fmt.Sprintf("\x1b[92;1m%s\x1b[0m", s)
}

// BYellow - яркий жирный желтый
func BYellow(s string) string {
	return fmt.Sprintf("\x1b[93;1m%s\x1b[0m", s)
}

// BBlue - яркий жирный голубой
func BBlue(s string) string {
	return fmt.Sprintf("\x1b[94;1m%s\x1b[0m", s)
}

// BMagenta - яркий жирный фиолетовый
func BMagenta(s string) string {
	return fmt.Sprintf("\x1b[95;1m%s\x1b[0m", s)
}

// BCyan - яркий жирный циан
func BCyan(s string) string {
	return fmt.Sprintf("\x1b[96;1m%s\x1b[0m", s)
}
