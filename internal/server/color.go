package server

import "fmt"

// Color придает строке цвет
// 91 - красный
// 92 - зеленый
// 93 - желтый
// 96 - бирюзовый
func Color(s string, colorStr string) string {
	var c int
	switch colorStr {
	case "green":
		c = 92
	case "red":
		c = 91
	case "yellow":
		c = 93
	case "blue":
		c = 96
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}
