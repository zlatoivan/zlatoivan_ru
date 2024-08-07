package utils

import "fmt"

// Color придает строке цвет
// 91 - красный
// 92 - зеленый
// 93 - желтый
// 96 - бирюзовый
func Color(s string, colorStr string) string {
	var c int
	switch colorStr {
	// Тускловатые
	case "nRed":
		c = 31
	case "nGreen":
		c = 32
	case "nYellow":
		c = 33
	case "nBlue":
		c = 34
	case "nMagenta":
		c = 35
	case "nCyan":
		c = 36

	// Яркие
	case "red":
		c = 91
	case "green":
		c = 92
	case "yellow":
		c = 93
	case "blue":
		c = 94
	case "magenta":
		c = 95
	case "cyan":
		c = 96
	}

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

func GetColorByStatusCode(status int) string {
	switch {
	case status < 200:
		return "blue"
	case status < 300:
		return "green"
	case status < 400:
		return "cyan"
	case status < 500:
		return "yellow"
	default:
		return "red"
	}
}
