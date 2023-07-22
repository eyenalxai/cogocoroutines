package functions

import (
	"fmt"
	"time"
)

func CountUp(upTo int) func() bool {
	state := 0

	return func() bool {
		if state > upTo {
			return true
		}
		fmt.Println("count up:", state)
		state++
		return false
	}
}

func CountDown(from int) func() bool {
	state := from

	return func() bool {
		if state < 0 {
			return true
		}
		fmt.Println("count dn:", state)
		state--
		return false
	}
}

func Sleep(duration int) func() bool {
	until := time.Now().Add(time.Duration(duration) * time.Second)

	return func() bool {
		if time.Now().After(until) {
			return true
		}
		return false
	}
}
