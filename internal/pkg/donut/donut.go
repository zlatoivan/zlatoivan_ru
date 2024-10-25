package donut

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

const (
	delay      = 8 * time.Millisecond
	coreString = ".,-~:;=!*#$@"

	width        = 80
	height       = 22
	bufferSize   = width * height
	thetaSpacing = 0.07
	phiSpacing   = 0.02
)

func genFrame(A float64, B float64) string {
	b := make([]byte, bufferSize)
	z := make([]float64, bufferSize)

	for t := range b {
		b[t] = ' '
		z[t] = 0
	}

	sa := math.Sin(A)
	ca := math.Cos(A)
	sb := math.Sin(B)
	cb := math.Cos(B)

	// theta goes around the cross-sectional circle of a torus
	for theta := float64(0); theta < 2*math.Pi; theta += thetaSpacing {
		st := math.Sin(theta)
		ct := math.Cos(theta)

		// phi goes around the center of revolution of a torus
		for phi := float64(0); phi < 2*math.Pi; phi += phiSpacing {
			sp := math.Sin(phi)
			cp := math.Cos(phi)

			h := ct + 2
			D := 1 / (sp*h*sa + st*ca + 5)
			t := sp*h*ca - st*sa

			x := int(width/2 + 30*D*(cp*h*cb-t*sb))
			y := int(height/2 + 15*D*(cp*h*sb+t*cb))

			o := x + width*y

			N := int(8 * ((st*sa-sp*ct*ca)*cb - sp*ct*sa - st*ca - cp*ct*sb))

			if 0 < y && y < height && 0 < x && x < width && z[o] < D {
				z[o] = D
				b[o] = coreString[max(N, 0)]
			}
		}
	}

	for k := 0; k < bufferSize; k++ {
		if k%width == 0 {
			b[k] = '\n'
		}
	}

	return string(b)
}

func SendDonutToConsole(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprint(w, "\u001B[H\u001B[2J") // clear previous stdout
	if err != nil {
		return fmt.Errorf("fmt.Fprintf: %w", err)
	}

	var A, B float64
	for {
		select {
		case <-r.Context().Done():
			log.Print("Curl client disconnected")
			return nil
		default:
			frame := genFrame(A, B)

			A += 0.071
			B += 0.035

			_, err = fmt.Fprint(w, "\u001B[H") // bring cursor to "home" location
			if err != nil {
				return fmt.Errorf("fmt.Fprintf: %w", err)
			}

			_, err = fmt.Fprint(w, frame)
			if err != nil {
				return fmt.Errorf("fmt.Fprintf: %w", err)
			}

			time.Sleep(delay)
		}
	}
}
