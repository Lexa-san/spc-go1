package main

import (
	"fmt"
)

func main() {
	var (
		pulse   float32
		min     float32
		max     float32
		counter int32
	)

	for {
		fmt.Scan(&pulse)

		if pulse < 0 {
			fmt.Println(counter)
			fmt.Printf("%.1f %.1f\n", min, max)
			return
		}

		if pulse >= 100 && pulse <= 140 {
			counter++
			if pulse < min || min == 0 {
				min = pulse
			}
			if pulse > max || max == 0 {
				max = pulse
			}
		}
	}
}
