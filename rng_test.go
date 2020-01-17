package xoshiro128

import (
	"log"
	"testing"
)

func Test_outputs(t *testing.T) {
	// Seeds generated from random.org
	gen := CreateXoshiro128(34993*14379, 7707*27782, 40232*65345, 34415*28170)

	min := 1.0
	max := 0.0

	for i := 0; i < 1024*1024; i++ {
		v := gen.Random()

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	log.Print("min: ", min, " max: ", max, " next: ", gen.Next())
}
